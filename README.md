<div align="center">

<img src="https://www.goravel.dev/logo.png" width="300" alt="Logo">

# Goravel CRUD Example
This is example project of goravel crud, this created with MVP pattern and planning to upgrade to use service and repository

## What's inside?

This project using Goravel

### Built With

- [Goravel](https://www.goravel.dev/)

### Want to running the project localy?

Run the following command:

1. Clone the repo into a public GitHub repository (or fork https://github.com/aditya-wiguna/goravel-crud-example/fork).

   ```sh
   git clone https://github.com/aditya-wiguna/goravel-crud-example
   ```

2. Go to the project folder

   ```sh
   cd goravel-crud-example
   ```

3. Install packages

   ```sh
    go mod tidy

    // Create .env environment configuration file
    cp .env.example .env

    // Generate application key
    go run . artisan key:generate
   ```

4. Migrate Database

   ```sh
    go run . artisan migrate
   ```

4. Run development mode on root folder

   ```sh
   go run .
   ```

Access the application in `localhost:3000`, sometimes when this port used it will be using other available port.
