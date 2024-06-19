FROM golang:1.22.1

WORKDIR /app

COPY . .

RUN go build -o math

CMD ["./math"]
