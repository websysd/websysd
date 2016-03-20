FROM golang:1.6

RUN go get github.com/ian-kent/websysd

EXPOSE 7050

ENTRYPOINT ["/go/bin/websysd"]
