FROM golang:latest
RUN mkdir /appserver/
ADD . /appserver/
WORKDIR /appserver/
RUN go build -o main .
CMD ["/appserver/main"]