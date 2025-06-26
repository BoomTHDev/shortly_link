package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ShortlyLink struct {
	gorm.Model
	OriginalURL string `gorm:"unique"`
	ShortURL    string `gorm:"unique"`
}

var (
	port = ":8080"
	db   *gorm.DB
	err  error
)

func main() {
	// Load environment variables
	if err = godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get environment variables
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Connect to database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Bangkok", dbHost, dbPort, dbUsername, dbPassword, dbName)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Connected to database")

	db.AutoMigrate(&ShortlyLink{})

	// Initialize router
	router := http.NewServeMux()

	// Setup API
	router.HandleFunc("/shorten", newShorten)
	router.HandleFunc("/{shortURL}", redirect)

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	}).Handler(router)

	// Initialize server with CORS middleware
	server := &http.Server{
		Addr:    port,
		Handler: handler,
	}

	// Start server
	fmt.Printf("Server running on port %s\n", port)
	if err = server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func newShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// Set Headers
	w.Header().Set("Content-Type", "application/json")

	// Get request body
	var requestBody struct {
		URL string `json:"url" binding:"required"`
	}

	if err = json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
		})
		return
	}

	fmt.Println("Received request to /shorten")
	fmt.Println("Payload:", requestBody.URL)

	link := ShortlyLink{}
	if err = db.Where("original_url = ?", requestBody.URL).First(&link).Error; err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"shortLink": link.ShortURL,
		})
		return
	} else if err != gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"messaeg": err.Error(),
		})
	}

	shortURL := generateShortURL()
	link = ShortlyLink{
		OriginalURL: requestBody.URL,
		ShortURL:    shortURL,
	}

	if err = db.Create(&link).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"shortLink": shortURL,
	})
}

func redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	shortURL := r.PathValue("shortURL")

	link := &ShortlyLink{}
	if err := db.Where("short_url = ?", shortURL).First(link).Error; err != nil {
		http.Error(w, "Link not found", http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]any{
			"error": err.Error(),
		})
		return
	}

	http.Redirect(w, r, link.OriginalURL, http.StatusMovedPermanently)
}

func generateShortURL() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6
	rand.New(rand.NewSource(time.Now().UnixNano()))

	shortURL := ""
	for range length {
		shortURL += string(chars[rand.Intn(len(chars))])
	}

	return shortURL
}
