# GO ENV
FROM golang:latest
# Maintainer
LABEL maintainer = "Cherry Yang <qiruiy5@uci.edu>"
# Set GIN_MODE & PORT
ENV GIN_MODE=release
ENV PORT=8080
# Create project dir in Docker Image
WORKDIR /go/src/cherryang.com
# Copy src code into Docker Image
COPY . /go/src/cherryang.com
# Install gin framework
RUN go get github.com/gin-gonic/gin
# Build application
RUN go build main.go
# Expose PORT
EXPOSE $PORT
# Set application entrypoint
ENTRYPOINT ["./main"]