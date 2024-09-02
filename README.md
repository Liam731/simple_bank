# SimpleBank RESTful API

## Project Overview

This project is a RESTful API built with Gin and GORM, designed to manage bank account operations. The project follows a modular structure, ensuring code maintainability and scalability.

## Project Structure

- **config**: Contains configuration for database connections and other global settings.

- **controller**: Receives requests from the router, calls the corresponding service methods, processes data, and returns the appropriate response. The controller layer is responsible for assembling the data and passing it to the service layer.

- **dao/model**: Defines the database table models (struct) generated using gentool, which GORM uses for database interactions.

- **helper**: Includes utility functions and error handling logic.

- **repository**: Handles CRUD operations on the database. Complex query logic is assembled in the service layer and then passed to the repository for execution.

- **request**: Defines the JSON request structure (struct) used in API calls. These structures are used in the controller layer to parse incoming client requests.

- **response**: Defines the JSON response structure (struct) used by the API. After processing requests, the controller formats the data into the corresponding response and returns it to the client.

- **router**: Sets up the API routes and maps the paths to the corresponding controller methods.

- **service**: Contains the business logic, integrating the repository and validator, processing the data, and passing it to the controller layer.

- **main.go**: The entry point of the project, responsible for connecting to the database, performing database migrations, initializing various modules (validator, repository, service, controller, router), and starting the server.

## Getting Started

### 1. Clone the Project

```bash
git clone https://github.com/yourusername/simplebank.git
cd simplebank
