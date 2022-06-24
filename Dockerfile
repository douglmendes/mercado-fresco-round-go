FROM golang:1.18.3

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
COPY . /go/src
RUN go build cmd/server/main.go
EXPOSE 8080 8080

CMD ["/main"]