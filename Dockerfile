FROM golang:1.6

RUN go get github.com/websysd/websysd

EXPOSE 7050

ENTRYPOINT ["/go/bin/websysd"]
