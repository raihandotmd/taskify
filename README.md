# Taskify

A modern task management API built with Go, Gin, GORM, and PostgreSQL. Taskify allows users to manage projects and tasks with full CRUD operations, JWT authentication, and a clean hexagonal architecture.

## 🚀 Features

- **User Management**: Registration, authentication, and logout with JWT
- **Project Management**: Create, read, update, and delete projects
- **Task Management**: Full CRUD operations for tasks within projects
- **Security**: JWT-based authentication and authorization with token revocation
- **Caching**: Redis integration for token blacklisting and session management
- **Database**: PostgreSQL with GORM ORM
- **Architecture**: Clean hexagonal architecture with clear separation of concerns
- **Containerization**: Docker and Docker Compose support
- **Migrations**: Database migration support

## 🏗️ Architecture

This project follows the hexagonal architecture (ports and adapters) pattern:

```
├── cmd/http/                    # Application entry point
├── internal/
│   ├── adapters/
│   │   ├── inbound/            # HTTP handlers and middleware
│   │   └── outbound/           # Database models and repositories
│   ├── config/                 # Configuration management
│   └── usecase/                # Business logic
├── migrations/                 # Database migrations
└── pkg/                        # Shared packages
```

## 📋 Prerequisites

- Go 1.24.4 or higher
- Docker and Docker Compose
- PostgreSQL (if running without Docker)
- Redis (if running without Docker)

## ⚙️ Installation

### 1. Clone the repository

```bash
git clone https://github.com/raihandotmd/taskify.git
cd taskify
```

### 2. Set up environment variables

```bash
cp .env.example .env
```

Edit the `.env` file with your configuration:

```env
# Application Configuration
APP_NAME=Taskify
APP_ENV=development
APP_HOST=localhost
APP_PORT=3000
APP_MODE=debug

# Database Configuration
POSTGRES_HOST=taskify_postgres
POSTGRES_PORT=5432
POSTGRES_USER=taskify
POSTGRES_PASSWORD=taskify
POSTGRES_DB=taskify_db

# Database Connection (for GORM)
GORM_HOST=localhost
GORM_PORT=5432
GORM_USER=taskify
GORM_PASSWORD=taskify
GORM_DB=taskify_db

# Adminer Configuration
ADMINER_DEFAULT_SERVER=taskify_postgres
ADMINER_DEFAULT_DB_DRIVER=pgsql
ADMINER_PORT=8080

# JWT Configuration
TOKEN_SECRET=your-super-secret-jwt-key
TOKEN_EXPIRATION=24h

# Redis Configuration
REDIS_ADDR=localhost:6379
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0
```

### 3. Install dependencies

```bash
go mod tidy
```

## 🐳 Running with Docker

### Start the infrastructure services

```bash
make init.up
```

This will start:
- PostgreSQL database
- Redis cache server
- Adminer (database administration tool)

### Run database migrations

```bash
make migrate.up
```

### Start the application

```bash
make run
```

The API will be available at `http://localhost:3000`

## 🛠️ Development Commands

| Command | Description |
|---------|-------------|
| `make run` | Start the application |
| `make init.up` | Start PostgreSQL, Redis, and Adminer |
| `make init.down` | Stop all infrastructure services |
| `make pg.up` | Start PostgreSQL only |
| `make pg.down` | Stop PostgreSQL only |
| `make redis.up` | Start Redis only |
| `make redis.down` | Stop Redis only |
| `make adminer.up` | Start Adminer only |
| `make adminer.down` | Stop Adminer only |
| `make migrate.up` | Run database migrations |
| `make migrate.down` | Rollback database migrations |
| `make init.down` | Stop all services |
| `make pg.up` | Start only PostgreSQL |
| `make pg.down` | Stop PostgreSQL |
| `make adminer.up` | Start only Adminer |
| `make adminer.down` | Stop Adminer |
| `make migrate.up` | Run database migrations |
| `make migrate.down` | Rollback database migrations |

## 📚 API Documentation

### Base URL
```
http://localhost:3000
```

### Health Check
```http
GET /health
```

### Authentication Endpoints

#### Register User
```http
POST /auth/register
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

#### Login User
```http
POST /auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "access_token": "jwt-token-here"
}
```

#### Login User
```http
GET /auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "password123"
}
```
---
### Protected Endpoints
All endpoints under `/api/v1/` require JWT authentication. Include the token in the Authorization header:

```
Authorization: Bearer <your-jwt-token>
```

### Authentication Endpoints

#### Logout User
Revoking the access token, by inserting into redis cache. 
```http
GET /auth/logout
```

### Project Endpoints

#### Create Project
```http
POST /api/v1/projects
Content-Type: application/json

{
  "name": "My Project",
  "description": "Project description"
}
```

#### Get All Projects
```http
GET /api/v1/projects
```

#### Update Project
```http
PUT /api/v1/projects/{id}
Content-Type: application/json

{
  "name": "Updated Project Name",
  "description": "Updated description"
}
```

#### Delete Project
```http
DELETE /api/v1/projects/{id}
```

### Task Endpoints

#### Create Task
```http
POST /api/v1/tasks
Content-Type: application/json

{
  "project_id": "project-uuid",
  "title": "Task Title",
  "description": "Task description",
  "status": "todo",
  "deadline": "2025-12-31"
}
```

**Task Status Options:** `todo`, `in_progress`, `done`

#### Get Tasks by Project
```http
GET /api/v1/projects/{project_id}/tasks
```

#### Update Task
```http
PUT /api/v1/tasks/{id}
Content-Type: application/json

{
  "title": "Updated Task Title",
  "description": "Updated description",
  "status": "in_progress",
  "deadline": "2025-12-31"
}
```

#### Delete Task
```http
DELETE /api/v1/tasks/{id}
```

## 🗄️ Database Schema

### Users Table
- `id` (UUID, Primary Key)
- `name` (String, Required)
- `email` (String, Required, Unique)
- `password` (String, Required)
- `created_at` (Timestamp)
- `updated_at` (Timestamp)

### Projects Table
- `id` (UUID, Primary Key)
- `name` (String, Required)
- `description` (Text)
- `created_by` (UUID, Foreign Key to Users)
- `created_at` (Timestamp)
- `updated_at` (Timestamp)

### Tasks Table
- `id` (UUID, Primary Key)
- `project_id` (UUID, Foreign Key to Projects)
- `title` (String, Required)
- `description` (Text)
- `status` (Enum: todo, in_progress, done)
- `deadline` (Date)
- `created_at` (Timestamp)
- `updated_at` (Timestamp)

## 🔐 Security

- **JWT Authentication**: All API endpoints (except auth) require valid JWT tokens
- **Password Hashing**: User passwords are securely hashed
- **Input Validation**: Request validation using struct tags

## 🧪 Testing

The project includes comprehensive validation:
- Request validation using Go validator
- Database constraints and foreign keys
- JWT token validation

## 📁 Project Structure

```
taskify/
├── cmd/http/                           # Application entry point
│   └── main.go
├── internal/
│   ├── adapters/
│   │   ├── inbound/
│   │   │   ├── handlers/http/          # HTTP request handlers
│   │   │   │   ├── model/              # Request/Response models
│   │   │   │   ├── projects/           # Project handlers
│   │   │   │   ├── tasks/              # Task handlers
│   │   │   │   ├── users/              # User handlers
│   │   │   │   └── routes.go           # Route definitions
│   │   │   └── middleware/             # Middleware (JWT auth)
│   │   └── outbound/
│   │       ├── models/                 # Database models
│   │       └── repositories/           # Database repositories
│   ├── config/                         # Configuration management
│   └── usecase/                        # Business logic layer
├── migrations/sql/                     # Database migrations
├── pkg/                               # Shared packages
├── docker-compose.yml                 # Docker services
├── Makefile                          # Development commands
└── .env.example                      # Environment variables template
```

## 📄 License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.

## 👤 Author

**Raihan**
- GitHub: [@raihandotmd](https://github.com/raihandotmd)

## 🙏 Acknowledgments

- Built with [Gin](https://gin-gonic.com/) web framework
- Database ORM with [GORM](https://gorm.io/)
- JWT implementation using [golang-jwt](https://github.com/golang-jwt/jwt)
- PostgreSQL database

---

**Made With Love! 🚀**
