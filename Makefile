run:
	go run cmd/main.go

tidy:
	go mod tidy

up:
	docker-compose up -d

# Stop services
down:
	docker-compose down

# Rebuild services
build:
	docker build -t service-task-list:1.0.0 .

# Show logs
logs:
	docker-compose logs -f

# Connect to the MySQL service
connect:
	docker-compose exec db mysql -uroot -proot

# Help
help:
	@echo "Commands:"
	@echo "make up      - Start services"
	@echo "make down    - Stop services"
	@echo "make build   - Rebuild services"
	@echo "make logs    - Show logs"
	@echo "make exec    - Run a command inside the service"
	@echo "make help    - Show this help"
