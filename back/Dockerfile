FROM golang:1.23

# Установите рабочую директорию
WORKDIR /app

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o back ./cmd/main.go

CMD ["./back"]