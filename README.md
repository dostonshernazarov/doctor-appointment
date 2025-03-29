# Doctor Appointment System

A RESTful API service for managing doctor appointments, built with Go, Fiber, and PostgreSQL.

## Features

- User authentication (signup/signin)
- Doctor management (CRUD operations)
- Appointment scheduling
- Swagger documentation
- PostgreSQL database
- Docker support

## Tech Stack

- Go 1.21
- Fiber (web framework)
- PostgreSQL (database)
- Docker & Docker Compose
- Swagger (API documentation)

## Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- PostgreSQL (if running locally)

## Getting Started

### Using Docker

1. Clone the repository:
```bash
git clone https://github.com/dostonshernazarov/doctor-appointment.git
cd doctor-appointment
```

2. Create a `.env` file in the root directory:
```env
APP_NAME=doctor-appointment
APP_VERSION=1.0.0
HTTP_PORT=8070
HTTP_USE_PREFORK_MODE=false
LOG_LEVEL=debug
PG_POOL_MAX=10
PG_URL=postgres://postgres:root@doctor_postgres:5432/doctor_appointment?sslmode=disable
METRICS_ENABLED=true
SWAGGER_ENABLED=true
JWT_SECRET=your-secret-key
JWT_EXPIRES_AT=3600
ROLE_ADMIN=admin
ROLE_USER=user
```

3. Run the application using Docker Compose:
```bash
docker compose up --build
```

The application will be available at `http://localhost:8070`


### Running Locally

1. Install dependencies:
```bash
go mod download
```

2. Set up the database:
```bash
psql -U postgres -c "CREATE DATABASE doctor_appointment;"
```

3. Run migrations:
```bash
migrate -path migrations -database "postgres://postgres:root@localhost:5432/doctor_appointment?sslmode=disable" up
```

4. Run the application:
```bash
go run cmd/app/main.go
```

## API Documentation

Once the application is running, you can access the Swagger documentation at:
- `http://localhost:8070/swagger`

## API Endpoints

### Authentication
- `POST /auth/signup` - Register a new user
- `POST /auth/signin` - Login user

### Users
- `GET /users` - Get all users
- `GET /users/:id` - Get user by ID
- `POST /users` - Create new user
- `PUT /users/:id` - Update user
- `DELETE /users/:id` - Delete user

### Doctors
- `GET /doctors` - Get all doctors
- `GET /doctors/:id` - Get doctor by ID
- `POST /doctors` - Create new doctor
- `PUT /doctors/:id` - Update doctor
- `DELETE /doctors/:id` - Delete doctor
- `GET /doctors/specializations` - List all specializations
- `GET /doctors/specialization/:specialization` - Get doctors by specialization

### Appointments
- `GET /appointments` - Get all appointments
- `GET /appointments/:id` - Get appointment by ID
- `POST /appointments` - Create new appointment
- `PUT /appointments/:id` - Update appointment
- `DELETE /appointments/:id` - Delete appointment
- `GET /appointments/doctor/:doctor_id` - Get appointments by doctor ID
- `GET /appointments/user/:user_id` - Get appointments by user ID
- `GET /appointments/doctor/:doctor_id/booked-schedules` - Get booked schedules by doctor ID
- `GET /appointments/user/:user_id/booked-schedules` - Get booked schedules by user ID

## Project Structure

```
.
├── cmd/
│   └── app/
│       └── main.go
├── config/
│   └── config.go
├── internal/
│   ├── controller/
│   │   └── http/
│   │       ├── v1/
│   │       ├── middleware/
│   │       └── models/
│   ├── entity/
│   ├── repo/
│   │   └── persistent/
│   └── usecase/
│       └── common/
├── migrations/
├── pkg/
│   ├── etc/
│   ├── logger/
│   ├── postgres/
│   └── token/
├── .env
├── docker-compose.yaml
├── Dockerfile
└── go.mod
```

## Testing

Run tests:
```bash
go test ./...
```
