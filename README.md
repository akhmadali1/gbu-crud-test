# GBU CRUD Test API

A simple **User CRUD API** implementation using two different backend stacks:

* **PHP Slim Framework + MongoDB**
* **Go Fiber Framework + PostgresSQL**

This project demonstrates how the same business logic (User CRUD) can be implemented using two different backend technologies.

---

# Architecture Overview
```
project-root
│
├── php-user-service
│   ├── src
│   │   ├── Application
│   │   ├── Domain
│   │   └── Infrastructure
│   └── public
│
├── go-user-service
│   ├── cmd/app
│   │   └── main.go
│   │
│   ├── internal
│   │   ├── app
│   │   └── user
│   │       ├── controller
│   │       ├── entity
│   │       ├── repo
│   │       └── usecase
│   │
│   └── migrations
│
└── postman
    └── gbu-crud-test.postman_collection.json
```

---

# Features

* Create User
* List Users
* View User
* Update User
* Delete User (Soft delete)
* Postman collection for testing

---

# PHP Slim API

### Health Check

```
GET /
```

### List Users

```
GET /users
```

### View User

```
GET /users/{id}
```

Example:

```
GET /users/69b0e7e42dc4ac4e400b7b92
```

### Create User

```
POST /users
```

Body:

```json
{
  "user_name": "john.doe",
  "first_name": "John",
  "last_name": "Doe"
}
```

### Update User

```
PUT /users/{id}
```

Body:

```json
{
  "user_name": "john.updated",
  "first_name": "John",
  "last_name": "Updated"
}
```

### Delete User

```
DELETE /users/{id}
```

---

# Go Fiber API

### List Users

```
GET /v1/user
```

### Create User

```
POST /v1/user
```

Body:

```json
{
  "data": {
    "user_name": "john.doe",
    "first_name": "John",
    "last_name": "Doe"
  }
}
```

### Update User

```
PUT /v1/user
```

Body:

```json
{
  "data": {
    "id": 1,
    "user_name": "john.updated",
    "first_name": "John",
    "last_name": "Updated"
  }
}
```

### Delete User

```
DELETE /v1/user
```

Body:

```json
{
  "data": {
    "id": 1
  }
}
```

---

# Postman Collection

The repository includes a **Postman Collection** for easy testing.

Import:

```
postman/gbu-crud-test.postman_collection.json
```
---

# Setup Instructions

## 1. Clone Repository

```
git clone https://github.com/akhmadali1/gbu-crud-test.git
cd gbu-crud-test
```

---

# Run PHP Slim API

Requirements:

* PHP 8+
* Composer
* MongoDB

Install dependencies:

```
composer install
```

Run server:

```
php -S localhost:8081 -t public
```

Test:

```
http://localhost:8081/users
```

---

# Run Go Fiber API

Requirements:

* Go 1.24+

Run server:

```
make run
```

Test:

```
http://localhost:8080/v1/user
```
# Run Go + PHP with Docker 🐳
Stop existing containers (optional):

```
docker compose down
```

Build all services:

```
docker compose build
```

Start the full stack:

```
docker compose up
```

Run in background (recommended):

```
docker compose up -d
```
---

# Example Response

```
GET /users
```

```json
[
  {
    "_id": "69b0e7e42dc4ac4e400b7b92",
    "user_name": "john.doe",
    "first_name": "John",
    "last_name": "Doe",
    "created_at": "2026-03-11T14:12:30Z"
  }
]
```

---

# Error Format

Example:

```json
{
  "statusCode": 500,
  "error": {
    "type": "SERVER_ERROR",
    "description": "User not found"
  }
}
```

---

# Technologies Used

### Backend

* PHP 8
* Slim Framework
* Go
* Fiber Framework

### Database

* MongoDB (PHP API)
* PostgresSQL (Go API)

### Tools

* Postman
* Composer
* Go Modules

---