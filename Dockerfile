# Build stage
FROM golang:1.25-alpine AS builder

# Устанавливаем зависимости для goose (если нужно)
RUN apk add --no-cache git

WORKDIR /app

# Копируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main ./cmd/server

# Final stage
FROM alpine:3.19

WORKDIR /app

# Устанавливаем только необходимые пакеты
RUN apk --no-cache add ca-certificates tzdata

# Копируем бинарник
COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./migrations

# Создаем не-root пользователя для безопасности
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup
USER appuser

# Открываем порт
EXPOSE 8512

# Запускаем приложение
CMD ["./main"]