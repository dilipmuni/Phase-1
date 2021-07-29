FROM golang:alpine

RUN mkdir /user

WORKDIR /user

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

RUN go get github.com/githubnemo/CompileDaemon


EXPOSE 9013


ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main

