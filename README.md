# Forum Site API

A robust and scalable API for managing forum platforms, built to handle user authentication, post creation, and user interactions like commenting and liking posts.

## Features

-   **User Authentication**: Signup, login, and token refresh functionality.
-   **Post Management**: Create, retrieve, and manage forum posts.
-   **User Activity**: Add comments and likes to posts.

## Installation

Clone the repository and navigate to the project directory:

```bash
# Clone the repository
git clone https://github.com/ilhamrdh/forum-site-api.git
cd forum-site-api

# Install dependencies
go mod tidy

# Build the application
go build -o forum-site-api

# Run the application
go run cmd/main.go
```

## Configuration

Before running the application, you need to configure the environment variables in the `intenal/configs/config.yaml` file.

```bash
# Create the config.yaml file
touch config.yaml

# Add the required environment variables
service:
    port: "your-port"
    secret_jwt: "your-secret-jwt"
database:
    db_source_name: "username:password@tcp(host:port)/database?parseTime=true"
```

## Migrations

To apply database migrations, run the following command:

```bash
# Set the database URL
MYSQL_URL="mysql://username:password@host:port/database"

# Create database migrations
migrate create -ext sql -dir scripts/migrations -seq

# Apply database migrations
migrate -database ${MYSQL_URL} -path scripts/migrations up
```

## Usage

To use the API, make requests to the appropriate endpoints and handle the responses as per your requirements.

## Endpoints

-   **Signup**: Endpoint to register a new user.
-   **Login**: Endpoint to authenticate a user and generate an access token.
-   **Refresh**: Endpoint to generate a new access token using a refresh token.
-   **Create Post**: Endpoint to create a new post.
-   **Get Post**: Endpoint to retrieve a post by ID.
-   **Add Comment**: Endpoint to add a comment to a post.
-   **Like Post**: Endpoint to like a post.

## Documentation

[Documentation](https://github.com/ilhamrdh/forum-site-api/blob/main/docs/swagger.yml)

## Acknowledgements

-   [Gin](https://github.com/gin-gonic/gin)
-   [Go](https://golang.org/)
-   [Migrate](https://github.com/golang-migrate/migrate)
-   [JWT](https://github.com/golang-jwt/jwt)
-   [MySQL](https://www.mysql.com/)
-   [Viper](https://github.com/spf13/viper)
-   [Swagger](https://swagger.io/)

## License

This project is licensed under the [MIT License](LICENSE).
