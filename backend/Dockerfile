FROM golang:latest as dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD ["go", "run", "cmd/main/main.go"]