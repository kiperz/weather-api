FROM golang:1.12

RUN mkdir /app

WORKDIR /app

COPY ./src .

RUN go build -a -o /app/weather-api

EXPOSE 8080

ENTRYPOINT ["/app/weather-api"]