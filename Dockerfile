# Stage 1: compile the program
FROM golang:1.16 as build-stage
WORKDIR /app
COPY go.* .
RUN go mod download
COPY . .
RUN go build -o main main.go

# Stage 2: build the image
FROM alpine:latest  
RUN apk --no-cache add ca-certificates libc6-compat
WORKDIR /app/
COPY --from=build-stage /app/main .
EXPOSE 8080
CMD ["./main"]  