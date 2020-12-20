
FROM golang:latest

LABEL maintainer = "Cherry Yang <qiruiy5@uci.edu>"

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /go/src/cherryang.com

COPY . /go/src/cherryang.com

RUN go get github.com/gin-gonic/gin

RUN go build main.go

EXPOSE $PORT

ENTRYPOINT ["./main"]