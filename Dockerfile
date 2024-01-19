FROM golang:1.20.12-alpine3.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o etov /app/main.go

EXPOSE 8181

CMD ["./etov"]