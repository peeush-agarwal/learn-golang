# REST API using `gin` framework

## Setup

- Initialize the module: `go mod init example/web-service-gin`
- Download and install the dependency `gin`: `go get -u github.com/gin-gonic/gin`
- Download and install `swag` for Swagger documentation: `go install github.com/swaggo/swag/cmd/swag@latest`
  - `go get -u github.com/swaggo/files`
  - `go get -u github.com/swaggo/gin-swagger`
- Run `swag init`
- Run `swag fmt`

## Run the server

```
go run .
```

## Requests to the server

- Open a new Terminal
- GET /albums request
  ```shell
  curl http://localhost:5000/albums
  ```
- POST /albums request
  ```shell
  curl http://localhost:5000/albums \
  --include \
  --header "Content-Type: application/json" \
  --request "POST" \
  --data '{"id": "4", "title": "The Modern Sound of Betty Carter", "artist": "Betty Carter", "price": 49.99}'
  ```
- GET /albums/1 request
  ```shell
  curl http://localhost:5000/albums/1
  ```
- GET /albums/100 request
  ```shell
  curl http://localhost:5000/albums/100
  ```

## References

- [Web service using Gin](https://go.dev/doc/tutorial/web-service-gin)