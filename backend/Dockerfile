FROM golang:1.24-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

RUN apk add --no-cache git

COPY go.* ./

RUN go mod download

COPY . .

EXPOSE 8080

CMD ["air"]
