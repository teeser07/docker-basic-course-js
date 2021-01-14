FROM golang:latest
WORKDIR /app/backend
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main .
RUN chmod +x ./main
ENTRYPOINT ["./main"]

EXPOSE 3030