FROM golang:1.12

RUN mkdir /app

COPY ./src /app

WORKDIR /app

RUN go build

CMD ["/app/weather-api"]