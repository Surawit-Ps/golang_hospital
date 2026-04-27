.PHONY: help build up down logs restart clean ps migrate test test-verbose test-coverage

help:
	@echo "Hospital Management System - Docker Commands"
	@echo ""
	@echo "Available targets:"
	@echo "  make build       - Build Docker images"
	@echo "  make up          - Start all services (build if needed)"
	@echo "  make down        - Stop all services"
	@echo "  make logs        - View logs from all services"
	@echo "  make ps          - Show running containers"
	@echo "  make clean       - Stop services and remove volumes (CAREFUL: deletes data)"
	@echo "  make restart     - Restart all services"
	@echo "  make backend-log - View backend logs"
	@echo "  make db-log      - View database logs"
	@echo "  make nginx-log   - View nginx logs"
	@echo "  make test        - Run all unit tests"
	@echo "  make test-verbose - Run tests with verbose output"
	@echo "  make test-coverage - Generate coverage report"

build:
	docker-compose build

up:
	docker-compose up -d
	@echo "Services started. Waiting for database to be ready..."
	@sleep 5
	@echo "Services are running!"

down:
	docker-compose down

logs:
	docker-compose logs -f

ps:
	docker-compose ps

clean:
	docker-compose down -v
	@echo "All services and volumes removed!"

restart:
	docker-compose restart

backend-log:
	docker-compose logs -f backend

db-log:
	docker-compose logs -f postgres

nginx-log:
	docker-compose logs -f nginx

# Backend specific commands
backend-build:
	docker-compose build backend

backend-restart:
	docker-compose down backend
	docker-compose up -d backend

backend-shell:
	docker exec -it hospital_backend /bin/sh

# Database commands
db-shell:
	docker exec -it hospital_db psql -U postgres -d hospital

db-backup:
	docker exec hospital_db pg_dump -U postgres -d hospital > backup_$(shell date +%Y%m%d_%H%M%S).sql

db-restore:
	@read -p "Enter backup file name: " file; \
	docker exec -i hospital_db psql -U postgres -d hospital < $$file

# Test endpoints
test-login:
	curl -X POST http://localhost/staff/login \
	  -H "Content-Type: application/json" \
	  -d '{"username":"staff1","password":"123456","hospitalId":"HOSP0001"}'

test-hospitals:
	curl -X GET http://localhost/hospitals \
	  -H "Authorization: Bearer YOUR_TOKEN_HERE"

test-health:
	curl http://localhost/health

# Development commands
dev-up:
	docker-compose up -d postgres
	cd backend && go run .

deps:
	cd backend && go mod download && go mod tidy

fmt:
	cd backend && go fmt ./...

# Testing commands
test:
	cd backend && go test -v ./...

test-verbose:
	cd backend && go test -v -race ./...

test-coverage:
	cd backend && go test -v -coverprofile=coverage.out ./... && go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

test-handler:
	cd backend && go test -v ./adapter/handler

test-service:
	cd backend && go test -v ./core/service

test-short:
	cd backend && go test -short ./...
