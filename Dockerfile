# Build stage
FROM golang:1.23.4 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod tidy
RUN go mod download

COPY . .

#RUN go build -o go-app main.go
RUN GOOS=linux GOARCH=amd64 go build -o go-app main.go


# Final minimal image
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/go-app .

EXPOSE 8090

CMD ["./go-app"]
