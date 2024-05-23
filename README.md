# Запуск
Сборка и запуск контейнеров Postgres, Nats-streaming, Redis :
```
docker-compose up -build
```
Запуск приложения:
```
go mod download
go mod tidy
go run cmd/main.go
```
