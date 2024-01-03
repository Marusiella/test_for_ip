FROM golang:latest
LABEL authors="marek"

WORKDIR /app

COPY . /app

RUN go mod download

RUN go build -o main .

EXPOSE 5454

CMD ["./main"]