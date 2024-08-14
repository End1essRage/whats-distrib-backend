# Используем образ golang для сборки
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Собираем приложение
RUN go build -o /app/whats-distrib-backend ./.

# Запускаем приложение при запуске контейнера
CMD ["./whats-distrib-backend"]

EXPOSE 8080