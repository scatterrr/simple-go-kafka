FROM golang:latest

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o app main.go    # <-- specify only main.go here

CMD ["./app"]
