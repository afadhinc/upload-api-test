# Upload API

Upload API is a test project using Golang and GIN as Framework.

## Installation

Download and install the required modules.

```bash
go get
```

## Usage

```
# Run app without Docker
go run main.go

# Build app with Docker
docker build -t upload-api  .

# Run app with Docker
docker run -d -p 8080:8080 upload-api
```

## Documentation

You can test and learn more about the REST API with following
[this link](https://documenter.getpostman.com/view/13019052/2s93kz5k3b)
