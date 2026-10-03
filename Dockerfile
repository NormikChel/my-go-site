# Сборка приложения
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Копируем зависимости
COPY go.mod ./
# Если есть go.sum, раскомментируй строчку ниже:
# COPY go.sum ./
# RUN go mod download

# Копируем весь исходный код
COPY . .

# Собираем статический бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Финальный легкий образ
FROM alpine:latest

WORKDIR /app

# Копируем скомпилированный бинарник
COPY --from=builder /app/server /app/server

# Копируем шаблоны и папки с ассетами, иначе шаблоны не найдутся!
COPY templates/ ./templates/
COPY assets/ ./assets/

EXPOSE 8080

CMD ["/app/server"]