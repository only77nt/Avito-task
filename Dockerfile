FROM golang:latest

RUN apt-get update
RUN apt-get install vim -y
RUN go get "github.com/gorilla/mux"