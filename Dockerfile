FROM golang:1.21.6

WORKDIR /app

COPY . .
RUN go mod download github.com/lib/pq
RUN go mod tidy
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g ./cmd/main.go
RUN go build -o main ./cmd/main.go

CMD ["./main"]