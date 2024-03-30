FROM golang:latest

WORKDIR /go/src/first-go-api

COPY . . 

RUN go build -o main cmd/main.go

CMD [ "./main" ]



