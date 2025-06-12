# Digital Wallet API

A simple Digital Wallet API built using **Go**, **Gin**, **GORM**, **MySQL**, and **JWT** for authentication. This API supports user registration, deposit, withdrawal, balance inquiry, and login functionality.

## Features

- **Login**: Generate JWT token for authenticated users.
- **Add User**: Register new users in the system.
- **Deposit**: Deposit funds into the user's wallet.
- **Withdraw**: Withdraw funds from the user's wallet.
- **Balance Inquiry**: Retrieve the current wallet balance.
- **JWT Authentication**: Secure routes with JWT tokens.

## Prerequisites

- **Go**: Version 1.16 or higher
- **MySQL**: Installed and running
- **Postman** (Optional): For API testing
- **Gin**: Go web framework
- **GORM**: ORM for Go
- **JWT**: JSON Web Token for authentication

## Installation

Follow these steps to set up and run the project locally:

### 1. Clone the repository
```bash
git clone https://github.com/fajarpratamaputra/digital_wallet_api
cd digital-wallet-api
```
### 2. Set up the environment

#### Create a .env file in the root of the project to store your configuration.
```
echo "DB_USER=root" > .env
echo "DB_PASSWORD=password" >> .env
echo "DB_HOST=localhost" >> .env
echo "DB_PORT=3306" >> .env
echo "DB_NAME=wallet_db" >> .env
echo "JWT_SECRET_KEY=secret_key" >> .env
```

### 3. Install dependencies

#### Install the necessary dependencies using Go modules.
``` go mod tidy```

### 4. Set up the database

#### Create a MySQL database:
```mysql -u root -p -e "CREATE DATABASE wallet_db;"```

#### (Optional) You can use the following command to run migrations if necessary
#### GORM will automatically migrate the schema when the app starts.

### 5. Run the application

#### Start the Go server by running:
```go run cmd/server.go```

#### By default, the API will run on http://localhost:8080

## API Endpoints

### 1. Login

- **Endpoint**: `POST /login`
- **Description**: Login a user and generate a JWT token.
- **Request Body**:
  ```json
  {
    "name": "user1",
    "password": "password123"
  }
  ```
- **Response**:

    ```{
    "token": "your-jwt-token-here"
    }
  ```

- **Curl Command**:
    ```
    curl -X POST -H "Content-Type: application/json" -d '{"name": "user1", "password": "password123"}' http://localhost:8080/login
    ```

### 2. Add User

   - **Endpoint:** 
        ```POST /user```

   - **Description: Register a new user.**

     - **Request Body:**
      ```
      {
      "name": "newuser",
      "password": "newpassword123"
      }
      ```
     - **Response:**
    ```
    {
    "message": "User created successfully",
    "user": {
        "id": 1,
        "name": "newuser",
        "balance": 0
        }
    }
    ```
    - **Curl Command:**
    ```
        curl -X POST -H "Content-Type: application/json" -d '{"name": "newuser", "password": "newpassword123"}' http://localhost:8080/user
    ```
   ### 3. Deposit
- **Endpoint:** 
``` POST /deposit ```
- **Description: Deposit funds into the user's wallet.**
- **Authentication: JWT token is required in the Authorization header (Bearer <your-token-here>).**
- **Request Body:**
```
{
"amount": 100
}
```
- **Response:**
```
{
"message": "Deposit successful",
"balance": 100
}
```
- **Curl Command:**
```
    curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer your-jwt-token-here" -d '{"amount": 100}' http://localhost:8080/deposit
```

### 4.Withdraw

   - **Endpoint:** ```POST /withdraw```

   - **Description: Withdraw funds from the user's wallet.**

   - **Authentication: JWT token is required in the Authorization header (Bearer <your-token-here>).**

   - **Request Body:**
```
{
"amount": 50
}
```
 - **Response:**
```
{
"message": "Withdrawal successful",
"balance": 50
}
```
- **Curl Command:**
```
    curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer your-jwt-token-here" -d '{"amount": 50}' http://localhost:8080/withdraw
```

### 5.Balance Inquiry

   - **Endpoint:** ```GET /balance```

   - **Description: Retrieve the user's current balance.**

   - **Authentication: JWT token is required in the Authorization header (Bearer <your-token-here>).**

   - **Response:**
```
{
"balance": 50
}
```
   - **Curl Command:**
```
curl -X GET -H "Content-Type: application/json" -H "Authorization: Bearer your-jwt-token-here" http://localhost:8080/balance
```

## Directory Structure
```
/digital-wallet
│
├── /cmd
│   └── /server.go               # Entry point for the application
│
├── /internal
│   ├── /db                      # Database connection and migration
│   └── /middleware              # JWT middleware
│
├── /domain
│   ├── /models                  # Model definitions (User, etc.)
│   ├── /interfaces              # Repository interfaces
│   └── /usecases                # Business logic (e.g., Withdraw, Deposit)
│
├── /infrastructure
│   ├── /repository              # Repository implementations
│   ├── /service                 # Services like JWT generation
│   └── /auth                    # Auth service
│
└── /api                         # API handlers (Login, Withdraw, etc.)
├── /auth.go                 # Handler for login and JWT
├── /withdraw.go             # Withdraw handler
├── /deposit.go              # Deposit handler
└── /balance.go              # Balance inquiry handler

```

## Environment Variables

- **DB_USER: MySQL database username (default: root)**
- **DB_PASSWORD: MySQL database password (default: password)**
- **DB_HOST: MySQL database host (default: localhost)**
- **DB_PORT: MySQL database port (default: 3306)**
- **DB_NAME: MySQL database name (default: wallet_db)**
- **JWT_SECRET_KEY: Secret key used for signing JWT tokens (default: secret_key)**