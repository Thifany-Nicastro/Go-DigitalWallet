FROM golang:1.21.5

WORKDIR /app

COPY . .

RUN go mod download

RUN go install github.com/cosmtrek/air@latest

# ENTRYPOINT go run main.go


# RUN go build -o main ./cmd/server/main.go

# EXPOSE 8080

# CMD [ "/app/main" ]