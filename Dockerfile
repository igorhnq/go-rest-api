FROM golang:1.24.5-alpine

WORKDIR /go/src/app

COPY . .

RUN go build -o main cmd/main.go

EXPOSE 8080

CMD ["./main"]