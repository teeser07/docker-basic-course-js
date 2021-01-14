FROM golang:latest AS build
WORKDIR /app/backend
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main .

FROM alpine
RUN apk --no-cache add tzdata
ENV TZ Asia/Bangkok
WORKDIR /app
COPY --from=build /app/backend/main .
RUN chmod +x ./main
ENTRYPOINT ./main

EXPOSE 3030