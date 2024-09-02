# SimpleBank RESTful API

## Project Overview

This project is a RESTful API built with Gin and GORM, designed to manage bank account operations. The project follows a modular structure, ensuring code maintainability and scalability.

## Project Structure

- **config**: Contains configuration for database connections and other global settings.

- **controller**: Receives requests from the router, calls the corresponding service methods, processes data, and returns the appropriate response. The controller layer assembles the data and passes it to the service layer.

- **dao/model**: Defines the database table models (struct) generated using gentool, which GORM uses for database interactions.

- **helper**: Includes utility functions and error-handling logic.

- **repository**: Handles CRUD operations on the database. Complex query logic is assembled in the service layer and passed to the repository for execution.

- **request**: Defines the JSON request structure used in API calls. These structures are used in the controller layer to parse incoming client requests.

- **response**: Defines the JSON response structure the API uses. After processing requests, the controller formats the data into the corresponding response and returns it to the client.

- **router**: Sets up the API routes and maps the paths to the corresponding controller methods.

- **service**: Contains the business logic, integrating the repository and validator, processing the data, and passing it to the controller layer.

- **main.go**: The entry point of the project, responsible for connecting to the database, performing database migrations, initializing various modules (validator, repository, service, controller, router), and starting the server.

## Getting Started

### 1. Set up docker-compose.yml

```bash
services:
  postgres:
    image: postgres:16-alpine
    environment:
      POSTGRES_USER: root
      POSTGRES_PASSWORD: secret
      POSTGRES_DB: root
      TZ: Asia/Taipei
    ports:
      - "5432:5432"

  api:
    image: liam731/simplebank:latest
    environment:
      DB_HOST: postgres
      DB_USER: root
      DB_PASSWORD: secret
      DB_NAME: root
    ports:
      - "8080:8080"
    depends_on:
      - postgres
```
### 2. Quick start

```bash
docker-compose up postgres -d
docker-compose up api -d
```
### 3. Test API

```bash
curl -X PUT http://localhost:8080/api/accounts -H "Content-Type: application/json" -d '{"owner": "Liam","currency": "USD","balance": 1000}'
curl -X PATCH http://localhost:8080/api/accounts/1 -H "Content-Type: application/json" -d '{"balance": 10}'
curl -X GET http://localhost:8080/api/accounts/1
```

