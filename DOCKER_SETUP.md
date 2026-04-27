# Hospital Management System - Docker Setup

This Docker Compose setup includes:
- **PostgreSQL** - Database service
- **Go Backend** - Hospital management API service
- **Nginx** - Reverse proxy and load balancer

## Prerequisites

- Docker
- Docker Compose

## Quick Start

### 1. Clone or navigate to the project directory

```bash
cd golang_hospital
```

### 2. Start all services

```bash
docker-compose up -d
```

This will:
- Create and start PostgreSQL container
- Build and start the Go backend service
- Start Nginx reverse proxy

### 3. Verify services are running

```bash
docker-compose ps
```

### 4. Access the application

- **API**: http://localhost (through Nginx)
- **Direct Backend**: http://localhost:8080
- **Database**: localhost:5432

## Configuration

Environment variables are defined in `.env` file:

```env
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=hospital
```

You can modify these values before running `docker-compose up`.

## Common Commands

### View logs
```bash
docker-compose logs -f backend
docker-compose logs -f postgres
docker-compose logs -f nginx
```

### Stop services
```bash
docker-compose down
```

### Stop services and remove volumes (WARNING: deletes database data)
```bash
docker-compose down -v
```

### Rebuild the backend service
```bash
docker-compose build --no-cache backend
docker-compose up -d backend
```

## API Endpoints

### Public Endpoints
- `POST /staff/login` - Staff login
- `POST /staff` - Register new staff

### Protected Endpoints (require JWT token)
- `POST /hospitals` - Create hospital
- `GET /hospitals` - Get all hospitals
- `GET /hospitals/:id` - Get hospital by ID
- `POST /patients` - Register patient
- `GET /patients` - Get patients list
- `GET /patients/:id` - Get patient by ID

## Database Connection

The database is automatically created on first run. You can connect directly to PostgreSQL:

```bash
psql -h localhost -U postgres -d hospital
```

Password: `password` (or as configured in .env)

## SSL/HTTPS Setup

To enable HTTPS:

1. Generate SSL certificates:
```bash
mkdir -p ssl
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout ssl/key.pem -out ssl/cert.pem
```

2. Uncomment HTTPS sections in `nginx.conf`

3. Restart Nginx:
```bash
docker-compose restart nginx
```

## Troubleshooting

### Backend cannot connect to database
- Check if PostgreSQL container is healthy: `docker-compose ps`
- View PostgreSQL logs: `docker-compose logs postgres`
- Ensure `.env` variables match docker-compose.yml settings

### Port already in use
Change port mappings in `docker-compose.yml`:
```yaml
ports:
  - "8081:8080"  # Maps host:8081 to container:8080
```

### Database connection refused
Wait for PostgreSQL to be ready (~10 seconds), then restart backend:
```bash
docker-compose restart backend
```

## Development

To rebuild and restart a service after code changes:

```bash
docker-compose down backend
docker-compose build backend
docker-compose up -d backend
```

## Production Notes

For production deployment:

1. Update `.env` with strong passwords
2. Enable HTTPS in nginx.conf
3. Set `GIN_MODE=release`
4. Use external managed PostgreSQL service
5. Configure proper logging and monitoring
6. Use Docker secrets for sensitive data
7. Implement health checks and auto-recovery
8. Set resource limits in docker-compose.yml

## Support

For issues, check:
1. Docker Compose logs
2. Service health status
3. Network connectivity between containers
4. Environment variable configuration
