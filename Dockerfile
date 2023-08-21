FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on

RUN cd /build && git clone https://github.com/veluvignesh027/todo.git

RUN cd /build/todo/src/ && go build

EXPOSE 7000

ENTRYPOINT [ "/build/todo/src/src" ]