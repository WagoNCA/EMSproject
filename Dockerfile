FROM golang:1.26.4

WORKDIR /app

COPY ./EMS-API .

RUN go build -mod=vendor -o ems-app .

CMD ["go", "run", "-mod=vendor", "main.go"]