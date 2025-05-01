# User Management API

A RESTful API for managing users with PostgreSQL storage and RabbitMQ event notifications.

---

## Features

- CRUD operations for user management
- Filtering and pagination for user listings
- Event-driven architecture with RabbitMQ
- Containerized with Docker
- Health checks and graceful shutdown
- Unit tests and mock implementations

---

## System Architecture

```
USERMANAGEMENT/
├── api/               # API entry point and routing
├── internal/          # Core application components
│   ├── config/        # Configuration utilities
│   ├── handler/       # HTTP request handlers
│   ├── mocks/         # Mock implementations for testing
│   ├── models/        # Data structures
│   ├── repository/    # Database operations
│   └── service/       # Business logic and event notifications
├── postgres-data/     # Persistent database storage
├── tests/             # Integration tests
├── docker-compose.yml # Docker compose file
├── Dockerfile         # Dockerfile
└── ...                # Configuration files
```

---

## API Endpoints

| Method | Endpoint   | Description             |
| ------ | ---------- | ----------------------- |
| GET    | /health    | Service health check    |
| GET    | /users     | List users (filterable) |
| POST   | /users     | Create a new user       |
| PUT    | /users/:id | Update an existing user |
| DELETE | /users/:id | Delete a user           |

---

## Prerequisites

- Docker 20.10+
- Docker Compose 2.0+
- Go 1.24+ (for development)

---

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/izzzicos/UserManagement.git
cd UserManagement
```

### 2. Set up environment variables

Edit the example environment file to your needs located in (`.env`):

| Variable                | Description              | Default                           |
| ----------------------- | ------------------------ | --------------------------------- |
| `API_PORT`              | Port for the API service | 8080                              |
| `DB_HOST`               | PostgreSQL host          | db                                |
| `DB_PORT`               | PostgreSQL port          | 5432                              |
| `DB_USER`               | Database username        | postgres                          |
| `DB_PASSWORD`           | Database password        | changeme                          |
| `DB_NAME`               | Database name            | userdb                            |
| `AMQP_URL`              | RabbitMQ connection URL  | amqp://guest:guest@rabbitmq:5672/ |
| `RABBITMQ_DEFAULT_USER` | RabbitMQ username        | guest                             |
| `RABBITMQ_DEFAULT_PASS` | RabbitMQ password        | guest                             |

---

### 3. Build and start services

```bash
docker-compose up --build
```

### 4. Verify the service is running

```bash
curl http://localhost:8080/health
```

---

## Development

### Running Tests

```bash
# Unit tests
make test

# Test coverage
make test-cover
```

## Deployment

The service is containerized and deployable to any orchestration platform.

### Production build

```bash
docker-compose -f docker-compose.yml up --build -d
```

## Monitoring

- **RabbitMQ Management UI**: [http://localhost:15672](http://localhost:15672)  
  _(Use credentials from `RABBITMQ_DEFAULT_USER` / `RABBITMQ_DEFAULT_PASS`)_

---

## Technical Details

### Database Schema

Users are stored with:

- ID (UUID)
- First Name
- Last Name
- Nickname
- Email
- Password (bcrypt hashed)
- Country
- Timestamps

### Event Notifications

Events published to RabbitMQ:

- `user.created`
- `user.updated`
- `user.deleted`

### Error Handling

- `400` for invalid requests
- `500` for server errors
- Consistent error response format

---
