# Library System REST API

This project is a library management system built with Go. It provides a RESTful API to perform operations such as adding, updating, and deleting books, authors, and reviews. The application also includes a built-in authorization mechanism using PostgreSQL.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running in Docker](#running-in-docker)
- [API Documentation](#api-documentation)
- [Monitoring & Health Checks](#monitoring--health-checks)
- [Technologies Used](#technologies-used)
- [Contributing](#contributing)
- [License](#license)

## Overview

The Library System REST API is designed to run in a Docker environment, making it easily deployable on any local machine. The API handles standard CRUD operations for library-related entities through dedicated handlers (e.g., `authors.go`, `books.go`, `reviews.go`). The route definitions in `routes.go` integrate these functionalities. Additionally, the project includes built-in support for Swagger documentation, Prometheus metrics, and various health checks to ensure system reliability.

## Features

- **RESTful API:** Implements CRUD operations for managing authors, books, and reviews.
- **Dockerized:** Runs seamlessly on any local machine using Docker.
- **Swagger Documentation:** Accessible at [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) for interactive API exploration.
- **Monitoring:** Prometheus metrics endpoint available at `/metrics`.
- **Health Checks:** Integrated health checks in Docker Compose for PostgreSQL, Redis, and the API.
- **Rate Limiting:** Built-in middleware to limit excessive requests.
- **Authorization:** PostgreSQL-backed authorization system.

## Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- Go 1.16+ (if running locally outside of Docker)

## Installation

1. **Clone the Repository**

   ```sh
   git clone https://github.com/mehmetpala/go-rest-api-mehmet-pala.git
   cd go-rest-api-mehmet-pala
   ```

2. **Set Up Environment Variables**

   Create a `.env` file in the project root with the following content:

   ```ini
   DB_HOST=postgres
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=postgres
   REDIS_HOST=redis
   REDIS_PORT=6379
   REDIS_PASSWORD=password
   PORT=8080
   ```

   *Note: The project loads these variables automatically via Docker Compose using the `env_file` directive.*

3. **Download Dependencies**

   If running locally:

   ```sh
   go mod tidy
   ```

## Running in Docker

The application is fully containerized. You can run all services (API, PostgreSQL, Redis, Prometheus, Grafana, etc.) using Docker Compose.

1. **Start the Services**

   ```sh
   docker-compose up --build
   ```

2. **Access the API and Tools**

   - **API Endpoints:** [http://localhost:8080](http://localhost:8080)
   - **Swagger UI:** [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
   - **Prometheus:** [http://localhost:9090](http://localhost:9090)
   - **Grafana:** [http://localhost:3000](http://localhost:3000)

## API Documentation

The project uses Swagger for API documentation. Once the containers are running, navigate to [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) to explore and test the API endpoints interactively.

## Monitoring & Health Checks

- **Health Check Endpoint:**  
  The API exposes a `/health` endpoint that returns the current status.
  
- **Prometheus Metrics:**  
  Metrics are available at `/metrics` for monitoring application performance.
  
- **Docker Health Checks:**  
  Docker Compose is configured with health checks for PostgreSQL, Redis, and the application to ensure all services are running correctly.

## Technologies Used

- **Go & Gin:** For building the REST API.
- **GORM:** ORM for database interactions.
- **PostgreSQL:** Primary database with authorization support.
- **Redis:** Used for caching and rate limiting.
- **Docker & Docker Compose:** Containerization and orchestration.
- **Swagger:** API documentation.
- **Prometheus & Grafana:** Monitoring and visualization.
- **Additional Tools:** Various middleware (e.g., rate limiting) to ensure robust functionality.

## Contributing

Contributions are welcome! If you would like to contribute, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.