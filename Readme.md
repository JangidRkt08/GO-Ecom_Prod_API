# GO Ecom Prod API

![Go Version](https://img.shields.io/badge/go-1.18%2B-blue.svg)
![Build Status](https://img.shields.io/github/actions/workflow/status/JangidRkt08/GO-Ecom_Prod_API/go.yml?branch=main)
![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)
![GitHub stars](https://img.shields.io/github/stars/JangidRkt08/GO-Ecom_Prod_API?style=social)

A production-ready e-commerce RESTful API built with **Go (Golang)**. This project provides a robust backend foundation for an e-commerce platform, focusing on clean architecture, type-safety, and performance. It utilizes **`sqlc`** for generating type-safe Go code from raw SQL and **`pgx`** for high-performance PostgreSQL operations.

## ✨ Core Features

* **Type-Safe Database Queries:** Uses **`sqlc`** to generate fully type-safe Go code from SQL queries, eliminating SQL injection vulnerabilities and runtime errors.
* **High-Performance DB Driver:** Employs **`pgx/v5`** as the PostgreSQL driver, known for its speed and native support for PostgreSQL features.
* **Structured Logging:** Integrated with Go's modern **`log/slog`** package for efficient, structured, and leveled JSON logging.
* **Robust Routing:** Built with **`go-chi/v5`**, a lightweight and idiomatic Go router, including essential middleware like `Logger`, `Recoverer`, `RequestID`, `RealIP`, and `Timeout`.
* **Transactional Integrity:** Order placement logic is wrapped in a **SQL transaction** (`BEGIN`, `COMMIT`, `ROLLBACK`) to ensure data consistency.
* **Graceful Shutdown:** Implements a standard `http.Server` with configured timeouts (`WriteTimeout`, `ReadTimeout`, `IdleTimeout`) for graceful shutdown.
* **Clean Architecture:** Organizes code by feature (`products`, `orders`) and separates concerns (`handler`, `service`, `repository`).
* **Containerized:** Includes a `docker-compose.yml` for easy setup of the development environment (PostgreSQL).

<img src="https://github.com/user-attachments/assets/c38dbab8-db1b-485a-8b81-5fea2e01f42e"
     width="1920">


## 🛠️ Tech Stack

* **Language:** [**Go (Golang)**](https://golang.org/)
* **Router:** [**`go-chi/v5`**](https://github.com/go-chi/chi)
* **Database:** [**PostgreSQL**](https://www.postgresql.org/)
* **Database Driver:** [**`pgx/v5`**](https://github.com/jackc/pgx)
* **Query Generation:** [**`sqlc`**](https://sqlc.dev/)
* **Database Migrations:** [**`goose`**](https://github.com/pressly/goose) (inferred from `GOOSE_DBSTRING`)
* **Logging:** [**`log/slog`**](https://pkg.go.dev/log/slog) (Native Go structured logger)
* **Containerization:** [**Docker**](https://www.docker.com/) / `docker-compose`

## 🚀 Getting Started

Follow these instructions to get a copy of the project up and running on your local machine.

### Prerequisites

* [Go](https://golang.org/doc/install) (version 1.18 or higher)
* [Docker](https://www.docker.com/products/docker-desktop)
* [goose](https://github.com/pressly/goose#install) (for database migrations)
* [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html) (for generating Go code from SQL)

### 1. Clone the Repository

```sh
git clone [https://github.com/JangidRkt08/GO-Ecom_Prod_API.git](https://github.com/JangidRkt08/GO-Ecom_Prod_API.git)
cd GO-Ecom_Prod_API
```

### 2. Set Up Environment

Create a `.env` file in the root directory. The `main.go` file provides a default, but you can override it here.

```env
# .env
# This is the connection string for goose and the application
GOOSE_DBSTRING="host=localhost user=userName password=your_Password dbname=dbName port=5432 sslmode=disable"
```

### 3. Start the Database
You can use the provided Docker Compose file to start a PostgreSQL instance.

```Bash

docker-compose up -d
```
This will start a PostgreSQL server on localhost:5432 with the credentials specified in docker-compose.yml.

### 4.Run Database Migrations
With the database running, apply the SQL migrations using goose.

```Bash

# Ensure goose is installed
# go install github.com/pressly/goose/v3/cmd/goose@latest

# Run migrations
goose -dir ./migrations/adapters/postgresql up

```
### 5. Generate sqlc Code
If you modify any SQL queries in the internal/adapters/postgresql/queries/ directory, you must regenerate the Go code.

```Bash

# Ensure sqlc is installed
sqlc generate
```
### 6. Install Dependencies & Run
```Bash

# Install Go modules
go mod tidy


# Run the application
go run ./cmd/main.go
```
The server will start, and you should see a log message: INFO "Connected to database" INFO "Listening on :8080"

## 📦 API Endpoints

<img width="1920" height="1080" alt="image" src="https://github.com/user-attachments/assets/75053a1c-9f11-4c06-8980-a21cc31bd05a" />


Here are the primary endpoints implemented so far:

### Health Check

* `GET /health`
    * **Description:** A simple health check endpoint.
    * **Success Response (200 OK):**
        ```sh
        All Good for now...
        ```

### Products

* `GET /products`
    * **Description:** Retrieves a list of all available products.
    * **Success Response (200 OK):**
        ```json
        [
          {
            "ID": 1,
            "Name": "Example Product",
            "PriceInCents": 1999,
            "Quantity": 100,
            "CreatedAt": "2023-10-27T10:00:00Z"
          }
        ]
        ```

### Orders

* `POST /orders`
    * **Description:** Places a new order. This operation is transactional. It validates product existence and stock (stock check is implemented, but decrementing is a future task).
    * **Request Body:**
        ```json
        {
          "customerId": 1,
          "items": [
            { "productId": 1, "quantity": 2 },
            { "productId": 2, "quantity": 1 }
          ]
        }
        ```
    * **Success Response (201 Created):**
        ```json
        {
          "ID": 123,
          "CustomerID": 1,
          "Status": "pending",
          "CreatedAt": "2023-10-27T10:30:00Z"
        }
        ```
    * **Error Responses:**
        * `404 Not Found`: If a `productId` in the items list does not exist (`"product not found"`).
        * `500 Internal Server Error`: If stock is insufficient (`"product has not enough stock"`) or for other database errors.
## 📁 Project Structure
```
.
├── cmd/
│   └── main.go
├── internal/
│   ├── adapters/
│   │   └── postgresql/
│   │       ├── queries/
│   │       └── sqlc/
│   ├── env/
│   ├── json/
│   ├── orders/
│   │   ├── handler.go
│   │   ├── service.go
│   │   └── types.go
│   ├── products/
│   │   ├── handler.go
│   │   └── service.go
│   └── api.go
├── migrations/
│   └── adapters/
│       └── postgresql/
├── .gitignore
├── docker-compose.yml
├── go.mod
├── go.sum
└── sqlc.yaml

```

## 🤝 Contributing

1. **Fork the repository**

2. **Create a feature branch**
   ```sh
   git checkout -b feature/AmazingFeature
   ```
3. **Commit your changes**
    ```sh
    git commit -m "Add some AmazingFeature"
    ```
4. **Push YOur Changes**
    ```sh
    git push origin feature/AmazingFeature
    ```
5. Open a Pull Request  
