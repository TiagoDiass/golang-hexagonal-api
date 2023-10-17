# Go Hexagonal API

This is a simple REST API built with the Go programming language, Hexagonal Architecture, MySQL and Docker.

## Endpoints

In this API you'll find endpoints related to a CRUD of **products**

- **GET** - GetProducts - Path: "_/products_"
- **GET** - GetProductById - Path: "_/products/{id}_"
- **POST** - CreateProduct - Path: "_/products_"
<!-- - **PUT** - UpdateProduct - Path: "_/products/{id}_" -->
- **DELETE** - DeleteProduct - Path: "_/products/{id}_"

## :arrow_forward: How to run

### First steps

If you want to run the application in your computer, follow these steps;

First of all, you'll need to have these tools installed on your computer

- [Git](https://git-scm.com/)
- [Go](https://go.dev/)
- [Docker](https://www.docker.com/)

Assuming you have those tools installed on your computer, you'll first need to clone the repository and install the dependencies,

Follow the steps bellow:

```
# Clone the repo
$ git clone https://github.com/TiagoDiass/golang-hexagonal-api.git

# Enter the repo's folder
$ cd golang-hexagonal-api

# Install the dependencies
$ go mod download
```

### Application

To run the application, just follow the steps below (assuming you have followed the first steps above)

```
# Start the Docker container
$ docker compose up -d

# Enter in the app directory inside Docker container
$ docker compose exec goapp bash

# Start the app
$ go run cmd/app/main.go
```

After followings these steps, you'll have the backend server running in your `localhost:8000`. If you want to stop the application you can just press `CTRL+C` in your terminal where you started the app.

### Tasks

- [ ] Make the documentation of the endpoints a table in this README
- [ ] Create a swagger for the API endpoints
- [ ] Add unit tests
