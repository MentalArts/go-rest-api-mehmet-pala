# go-rest-api-mehmet-pala

This project is a Go-based REST API for managing a book library. It provides endpoints for managing authors, books, and reviews.

## Features

- RESTful API for managing authors, books, and reviews
- Swagger documentation for API endpoints
- Docker support for containerization
- Environment variable configuration
- Continuous Integration and Deployment with GitLab CI

## Project Structure

```
go-rest-api-mehmet-pala/
├── .git/                   # Git repository files
├── .gitlab-ci.yml          # GitLab CI configuration
├── Dockerfile              # Dockerfile for building the application
├── docker-compose.yaml     # Docker Compose configuration
├── go.mod                  # Go module file
├── go.sum                  # Go dependencies file
├── main.go                 # Main application entry point
├── internal/               # Internal packages
│   ├── db/                 # Database initialization and connection
│   └── routes/             # API route definitions
├── .env                    # Environment variables file
└── README.md               # Project README file
```


## Getting Started

### Prerequisites

- Go 1.16+
- Docker
- Docker Compose

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/MentalArts/go-rest-api-mehmet-pala.git
    cd go-rest-api-mehmet-pala
    ```

2. Copy the `.env.example` file to `.env` and update the environment variables as needed. The `.env.example` file contains placeholder values for the environment variables required by the application. You need to replace these placeholders with actual values.

    ```sh
    cp .env.example .env
    ```

    Open the `.env` file in a text editor and update the values as needed. For example:

    ```env
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_NAME=your_db_name
    ```

3. Build and run the Docker containers:

    ```sh
    docker-compose up --build
    ```

4. The API will be available at `http://localhost:8080`.

### API Documentation

Swagger documentation is available at `http://localhost:8080/swagger/index.html`.

### Running Tests

To run tests, use the following command:

```sh
go test ./...

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.