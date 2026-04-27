# Windows batch file for running Docker Compose commands
# Usage: docker-commands.bat <command>
# Example: docker-commands.bat up

@echo off
setlocal enabledelayedexpansion

if "%1"=="" (
    echo Hospital Management System - Docker Commands
    echo.
    echo Usage: docker-commands.bat [command]
    echo.
    echo Available commands:
    echo   up         - Start all services
    echo   down       - Stop all services
    echo   build      - Build Docker images
    echo   logs       - View logs
    echo   ps         - Show running containers
    echo   clean      - Stop services and remove volumes (CAREFUL!)
    echo   backend    - View backend logs
    echo   db         - View database logs
    echo   nginx      - View nginx logs
    echo   shell      - Open backend shell
    echo   restart    - Restart all services
    exit /b 0
)

if "%1"=="up" (
    echo Starting services...
    docker-compose up -d
    timeout /t 5 /nobreak
    echo Services started!
    exit /b 0
)

if "%1"=="down" (
    docker-compose down
    exit /b 0
)

if "%1"=="build" (
    docker-compose build
    exit /b 0
)

if "%1"=="logs" (
    docker-compose logs -f
    exit /b 0
)

if "%1"=="ps" (
    docker-compose ps
    exit /b 0
)

if "%1"=="clean" (
    echo WARNING: This will delete all data!
    pause
    docker-compose down -v
    echo All services and volumes removed!
    exit /b 0
)

if "%1"=="backend" (
    docker-compose logs -f backend
    exit /b 0
)

if "%1"=="db" (
    docker-compose logs -f postgres
    exit /b 0
)

if "%1"=="nginx" (
    docker-compose logs -f nginx
    exit /b 0
)

if "%1"=="shell" (
    docker exec -it hospital_backend /bin/sh
    exit /b 0
)

if "%1"=="restart" (
    docker-compose restart
    exit /b 0
)

echo Unknown command: %1
exit /b 1
