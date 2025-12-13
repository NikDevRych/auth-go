# Auth-go - Go JWT Authentication Service

A minimal Go (`net/http`) application that implements JWT authentication with access and refresh tokens.

> ⚠️ This project is under active development and currently supports only basic functionality.

---

## Features

- User authentication
- JWT access token
- Refresh token
- HTTP API built with `net/http`

---

## Requirements

- Go (version 1.25.5+)
- Docker
- Docker Compose

---

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/NikDevRych/auth-go.git
```
```bash
cd auth-go
```

### 2. Create docker-compose.yml

Create a file named docker-compose.yml in the root directory of the project and paste your Docker Compose configuration into it.

```docker-compose.yml
services:
  postgres-db:
    image: postgres:latest
    restart: always
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: auth_go
    ports:
      - "5432:5432"

  auth-go:
    build:
      context: .
      dockerfile: Dockerfile
    restart: always
    ports:
      - "8080:8080"
    environment:
      connection_string: "postgres://postgres:postgres@postgres-db:5432/auth_go"
      jwt_key: "some-jwt-token-key-only-for-test"
    depends_on:
      - postgres-db
```

### 3. Run the application

Build and start the application using Docker Compose:

```bash
docker-compose up --build
```
