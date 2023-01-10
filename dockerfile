FROM golang:latest
WORKDIR /app
COPY . .
WORKDIR /app/cmd/main
RUN go mod tidy
RUN go build -o main .
EXPOSE 9000
CMD ["./main"]