FROM golang:1.21.5

WORKDIR /app

COPY ./ /app

RUN go mod download

# ENTRYPOINT go run main.go


# RUN go build -o main ./cmd/server/main.go

# EXPOSE 8080

# CMD [ "/app/main" ]