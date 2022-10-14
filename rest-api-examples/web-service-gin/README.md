# REST API using `gin` framework

## Setup

- Initialize the module: `go mod init example/web-service-gin`
- Download and install the dependency `gin`: `go get -u github.com/gin-gonic/gin`

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
