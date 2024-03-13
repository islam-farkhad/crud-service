FROM golang:1.19-alpine

# Set destination for COPY
WORKDIR /app

COPY . .

RUN apk add --no-cache ca-certificates && update-ca-certificates

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o crud-service cmd/posts_service/main.go

EXPOSE 8080

# Run
CMD ["./crud-service"]
