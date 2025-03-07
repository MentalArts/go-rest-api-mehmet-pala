# go-rest-api-mehmet-pala

This project is a Go-based REST API for managing a book library. It provides endpoints for managing authors, books, and reviews. It integrates Swagger for API documentation, supports Docker for containerization, and uses PostgreSQL and Redis for data storage.

## Features

- **RESTful API** for managing authors, books, and reviews.
- **Swagger Documentation**: Auto-generated API documentation.
- **CRUD Operations**: Perform CRUD operations on books, authors, and reviews.
- **Health Check Endpoint**: API includes a `/health` endpoint to check the status of the server.
- **Prometheus Monitoring**: Includes Prometheus metrics at `/metrics`.
- **Docker Support**: For containerization with PostgreSQL, Redis, Prometheus, Grafana, and the application.
- **Environment Variable Configuration**: `.env` file for easy configuration.
- **Continuous Integration and Deployment** with GitLab CI.


## Prerequisites

- **Go**: Version 1.16 or higher.
- **Docker**: For containerizing the application and running services like PostgreSQL, Redis, Prometheus, and Grafana.
- **Swagger**: To auto-generate API documentation.
- **PostgreSQL**: For the relational database.
- **Redis**: For caching and session storage.
- **GitLab CI**: For Continuous Integration and Deployment.

## Getting Started

### Installation

1. **Clone the repository**:
    ```sh
    git clone https://github.com/MentalArts/go-rest-api-mehmet-pala.git
    cd go-rest-api-mehmet-pala
    ```

2. **Set up `.env` file**:
   Copy the `.env.example` file to `.env` and update the environment variables as needed. The `.env.example` file contains placeholder values for the environment variables required by the application.

   ```sh
   cp .env.example .env
