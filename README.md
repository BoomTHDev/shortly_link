# Shortly Link

A modern URL shortener service built with Next.js, TypeScript, and Go. This application allows users to create short, memorable links from long URLs, making them easier to share.

## Features

- 🔗 Convert long URLs into short, shareable links
- ⚡ Fast redirection to original URLs
- 🎨 Modern, responsive user interface
- 🐳 Docker support for easy deployment
- 🔄 Real-time link generation
- 📱 Mobile-friendly design

## Tech Stack

### Frontend

- Next.js 15 (App Router)
- TypeScript
- Tailwind CSS for styling
- React Server Components

### Backend

- Go 1.24
- PostgreSQL database
- GORM ORM

## Prerequisites

Before you begin, ensure you have the following installed:

- Docker and Docker Compose
- Node.js (v20 or later)
- Go (1.24 or later)
- PostgreSQL (or use the Docker setup)

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/BoomTHDev/shortly_link.git
cd shortly_link
```

### 2. Set up environment variables

Create a `.env` file in the root directory with the following variables:

```env
# Backend
DB_USERNAME=postgres
DB_PASSWORD=your_secure_password
DB_HOST=db
DB_PORT=5432
DB_NAME=shortly

# Frontend (optional)
NEXT_PUBLIC_API_URL=http://localhost:8080
```

### 3. Run with Docker Compose (Recommended)

The easiest way to run the application is using Docker Compose. This will set up all the necessary services with a single command.

#### Prerequisites

- Docker Engine 20.10.0+
- Docker Compose 2.0.0+

#### Environment Setup

1. Create a `.env` file in the project root with the following variables:

```env
# Database
POSTGRES_USER=postgres
POSTGRES_PASSWORD=your_secure_password
POSTGRES_DB=shortly

# Backend
DB_USERNAME=postgres
DB_PASSWORD=your_secure_password
DB_HOST=db
DB_PORT=5432
DB_NAME=shortly

# Frontend
NEXT_PUBLIC_API_URL=http://localhost:8080
```

#### Starting the Application

To build and start all services in detached mode:

```bash
docker-compose up -d --build
```

This will start the following services:

- Frontend (Next.js) on http://localhost:3000
- Backend (Go) on http://localhost:8080
- PostgreSQL database on port 5432

#### Viewing Logs

To view logs from all services:

```bash
docker-compose logs -f
```

View logs for a specific service:

```bash
docker-compose logs -f frontend
docker-compose logs -f backend
docker-compose logs -f db
```

#### Common Commands

Stop all services:

```bash
docker-compose down
```

Stop and remove all containers, networks, and volumes:

```bash
docker-compose down -v
```

Rebuild and restart a specific service (e.g., after making changes):

```bash
docker-compose up -d --build frontend
```

#### Development with Hot Reloading

For development with hot reloading, you can use the following command to start the frontend in development mode:

```bash
# In one terminal
cd frontend
npm run dev

# In another terminal
docker-compose up -d db backend
```

#### Database Management

To access the PostgreSQL database:

```bash
docker-compose exec db psql -U postgres -d shortly
```

#### Cleanup

To remove all containers, networks, and volumes (including database data):

```bash
docker-compose down -v --rmi all
```

### 4. Run without Docker

#### Backend

```bash
cd backend
go mod download
go run main.go
```

#### Frontend

In a new terminal:

```bash
cd frontend
npm install
npm run dev
```

## API Endpoints

- `POST /shorten` - Create a new short URL

  - Request body: `{ "url": "https://example.com" }`
  - Response: `{ "shortUrl": "http://localhost:8080/abc123" }`

- `GET /:shortURL` - Redirect to the original URL

## Development

### Frontend Development

```bash
cd frontend
npm install
npm run dev
```

### Backend Development

```bash
cd backend
go run main.go
```

### Running Tests

Frontend tests:

```bash
cd frontend
npm test
```

Backend tests:

```bash
cd backend
go test ./...
```

## Environment Variables

### Backend

- `DB_USERNAME`: Database username
- `DB_PASSWORD`: Database password
- `DB_HOST`: Database host
- `DB_PORT`: Database port
- `DB_NAME`: Database name

### Frontend

- `NEXT_PUBLIC_API_URL`: URL of the backend API

## Contributing

1. Fork the repository
2. Create a new branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with ❤️ using Next.js and Go
- Inspired by popular URL shorteners like Bitly and TinyURL
