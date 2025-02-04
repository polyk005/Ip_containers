FROM golang:1.23

WORKDIR /app

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN go mod download
RUN go build -o pinger ./cmd/main.go

CMD ["./pinger"]
