FROM --platform=amd64 golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go build ./src/main.go

EXPOSE 8000

CMD ["./main" ]