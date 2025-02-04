FROM golang:1.19-alpine

RUN apk add --no-cache gcc musl-dev

WORKDIR /app
COPY . .

RUN go build -o pinger ./cmd/main.go

CMD ["./pinger"]