FROM golang:alpine

RUN apk add git

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN go get -d

RUN go mod tidy

RUN go build -o main .

CMD ["/app/main"]

EXPOSE 8080
