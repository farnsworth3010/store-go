FROM golang:1.21.6

WORKDIR /app

COPY . .
RUN go mod download github.com/lib/pq
RUN go mod tidy
RUN go build -o main ./cmd/main.go

CMD ["./main"]