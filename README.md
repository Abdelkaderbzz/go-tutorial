
# Go Tutorial API (Gin + MongoDB)
This is a sample REST API built with Go, Gin, and MongoDB. It provides basic user management functionalities, including creating and greeting a user. The project also includes input validation and automatically generated API documentation.

## Project Structure

The project follows a standard structure for a Go web application:

- **main.go**: The entry point of the application. It initializes the Gin router, connects to the database, and registers the routes.
- **controllers/**: Contains the request handlers (controllers) for different API endpoints.
- **database/**: Manages the database connection.
- **docs/**: Contains the auto-generated Swagger documentation.
- **models/**: Defines the data structures and validation rules.
- **routes/**: Defines the API routes and maps them to the corresponding controllers.
- **services/**: Contains the business logic of the application.

## Getting Started

- **main.go**: The entry point of the application. It initializes the Gin router, connects to the database, and registers the routes.
- **controllers/**: Contains the request handlers (controllers) for different API endpoints.
- **database/**: Manages the database connection.
- **docs/**: Contains the auto-generated Swagger documentation.
- **models/**: Defines the data structures and validation rules.
- **routes/**: Defines the API routes and maps them to the corresponding controllers.
- **services/**: Contains the business logic of the application.


### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or higher)
- [MongoDB](https://docs.mongodb.com/manual/installation/)

### Installation

1.  Clone the repository:
    ```bash
    git clone https://github.com/your-username/go_tutorial.git
    cd go_tutorial
    ```

2.  Create a `.env` file in the root directory and add your MongoDB connection string:
    ```
    MONGO_URI=mongodb://localhost:27017
    PORT=8080
    ```

3.  Install the dependencies:
    ```bash
    go mod tidy
    ```

4.  Run the application:
    ```bash
    go run main.go
    ```

The server will start on `http://localhost:8080`.

## API Endpoints

The API provides the following endpoints:

- `GET /user/{name}`: Greets a user by name.
- `POST /users`: Creates a new user.

### `GET /user/{name}`

- **Description**: Returns a greeting message to the user.
- **URL Params**: `name=[string]` (required)
- **Success Response**:
  - **Code**: 200 OK
  - **Content**: `{ "message": "Hello, {name}!" }`

### `POST /users`

- **Description**: Creates a new user.
- **Body**:
  ```json
  {
    "name": "John Doe",
    "email": "john.doe@example.com",
    "password": "password123",
    "age": 30
  }
  ```
- **Success Response**:
  - **Code**: 201 Created
  - **Content**:
    ```json
    {
      "inserted_id": "60f1b3b3b3b3b3b3b3b3b3b3",
      "success": true,
      "message": "user created successfully"
    }
    ```
- **Error Responses**:
  - **Code**: 400 Bad Request (for validation errors)
  - **Code**: 409 Conflict (if the email is already registered)
  - **Code**: 500 Internal Server Error

## Validation

Request body validation is handled by the [go-playground/validator](https://github.com/go-playground/validator) library. Validation rules are defined in the `models/user.go` file using struct tags.

For example, the `User` model has the following validation rules:
- `Name`: required, min 3 characters, max 20 characters
- `Email`: required, must be a valid email address
- `Password`: required
- `Age`: greater than or equal to 0, less than or equal to 120

## Swagger Documentation

The API documentation is available at `http://localhost:8080/swagger/index.html`. It is automatically generated from the code comments using [swaggo/swag](https://github.com/swaggo/swag).

To regenerate the documentation, run the following command:

```bash
swag init -g main.go -o docs
```

## Tech Stack

- **Go**: The programming language used to build the API.
- **Gin**: A web framework for Go.
- **MongoDB Go Driver**: The official MongoDB driver for Go.
- **godotenv**: A library to load environment variables from a `.env` file.
- **go-playground/validator**: A library for request body validation.
- **swaggo/swag**: A tool to automatically generate RESTful API documentation with Swagger 2.0.
