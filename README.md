# Go Gin CRUD API

## 📦 Project Overview

This project is a simple backend application written in **Go** that provides a REST API for managing users.  
It demonstrates how to build a basic CRUD service using **Gin**, **GORM**, and **PostgreSQL**, with **Docker** used to run the database.

The project is intended as a learning and portfolio example for backend development.

---

## 🧩 Core Components

### Main Application
**`main.go`**

Initializes the Gin web server, connects to the PostgreSQL database, and registers the application routes.

---

### Controllers
**`controller/user.go`**

Handles the HTTP logic for user operations, including:

- Create user
- Get users
- Get user by ID
- Update user
- Delete user

Controllers act as the bridge between the HTTP layer and the database models.

---

### Routes
**`routes/user.go`**

Defines the API endpoints and maps them to their corresponding controller functions.

Example routes:

- `GET /users`
- `GET /users/:id`
- `POST /users`
- `PUT /users/:id`
- `DELETE /users/:id`

---

### Configuration
**`config/db.go`**

Responsible for:

- Establishing the PostgreSQL database connection
- Running automatic migrations with GORM
- Managing database initialization

---

### Models
**`models/user.go`**

Defines the **User** structure using GORM fields.

Example fields may include:

- ID
- Name
- Email
- CreatedAt
- UpdatedAt

---

### Docker Setup
**`docker-compose.yml`**

Provides a PostgreSQL container used by the application.

This allows the project to run without installing PostgreSQL locally.

---

## 🧪 Technologies Used

- **Go** – Programming language used for the backend
- **Gin** – High-performance HTTP web framework for Go
- **GORM** – ORM library for database operations
- **PostgreSQL** – Relational database system
- **Docker** – Containerization platform used to run the database

---

## 🚀 Running the Project

Start the PostgreSQL container:

```bash
docker-compose up -d
