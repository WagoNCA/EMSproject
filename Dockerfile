
FROM golang:1.26.4

WORKDIR /app

COPY . .

RUN go mod tidy

CMD ["go", "run", "helloEcho.go"]
