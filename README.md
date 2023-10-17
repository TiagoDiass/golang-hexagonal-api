# Go Hexagonal API

This is a simple REST API built with the Go programming language, Hexagonal Architecture, MySQL and Docker.

<p align="left">
  <img alt="Repo's top languages" src="https://img.shields.io/static/v1?label=Main%20technology&message=Go&style=for-the-badge&color=007D9C&labelColor=000000">
</p>

## Endpoints

In this API you'll find endpoints related to a CRUD of **products**

<table>
  <tr>
    <th>Request name</th>
    <th>Method</th>
    <th>Endpoint</th>
    <th>Request body</th>
    <th>Returns</th>
  </tr>
  
  <tr>
    <td>GetProducts</td>
    <td>GET</td>
    <td><i>/products</i></td>
    <td>No body</td>
    <td>All products</td>
  </tr>

  <tr>
    <td>GetProductById</td>
    <td>GET</td>
    <td><i>/products/{productId}</i></td>
    <td>No body</td>
    <td>A single product</td>
  </tr>

   <tr>
    <td>CreateProduct</td>
    <td>POST</td>
    <td><i>/products</i></td>
    <td>JSON with the properties <br /> <code>name: string, price: integer</code></td>
    <td>The created product</td>
  </tr>

  <tr>
    <td>DeleteProduct</td>
    <td>DELETE</td>
    <td><i>/products/{productId}</i></td>
    <td>No body</td>
    <td>Nothing</td>
  </tr>
</table>

<!-- ## Things I used

- Go
- MySQL
- Docker
- Chi (a router for building Go HTTP services)
- Hexagonal Architecture -->

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
