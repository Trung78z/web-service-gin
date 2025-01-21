# Web Service with Gin

This is a web service built using the Gin framework in Go. It includes user management functionalities with PostgreSQL and Redis integration.

## Project Structure
. ├── cmd/ │ └── main.go ├── config/ │ └── config.go ├── db/ │ ├── migration/ │ │ ├── 000001_init_schema.down.sql │ │ └── 000001_init_schema.up.sql │ └── query/ │ └── user.sql ├── internal/ │ ├── controllers/ │ │ └── userController.go │ ├── middlewares/ │ ├── repositories/ │ │ ├── queries/ │ │ │ ├── db.go │ │ │ ├── models.go │ │ │ └── user.sql.go │ │ └── repository.go │ ├── routes/ │ │ ├── main.go │ │ └── userRoutes.go │ └── services/ │ └── userService.go ├── pkg/ │ ├── cache/ │ │ └── rd.go │ ├── database/ │ │ └── db.go │ ├── logger/ │ └── utils/ │ ├── helpers.go │ └── response.go ├── test/ │ └── test.http ├── .air.toml ├── .env ├── .env.example ├── .gitignore ├── docker-compose.yml ├── go.mod ├── go.sum ├── main_test.go ├── Makefile ├── README.md └── sqlc.yaml

## Getting Started

### Prerequisites

- Go 1.23.3 or later
- Docker and Docker Compose
- PostgreSQL
- Redis

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/Trung78z/web-service-gin.git
    cd web-service-gin
    ```

2. Copy the example environment file and update the values:
    ```sh
    cp .env.example .env
    ```

3. Start the services using Docker Compose:
    ```sh
    docker-compose up -d
    ```

4. Run the application:
    ```sh
    go run cmd/main.go
    ```

### Running Tests

To run the tests, use the following command:
```sh
go test ./...
API Endpoints
Users
GET /api/v1/users - List all users
GET /api/v1/users/:id - Get user by ID
POST /api/v1/users - Create a new user
License
This project is licensed under the MIT License.
Feel free to customize this README.md to better fit your project's needs.
Feel free to customize this README.md to better fit your project's needs.