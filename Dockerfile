# Используем образ golang для сборки
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Собираем приложение
RUN go build -o /app/whats-mass-ditrib ./.

# Запускаем приложение при запуске контейнера
CMD ["./whats-mass-ditrib"]

EXPOSE 8080