FROM golang
WORKDIR /go/src/github.com/CGTYK/auth
COPY main.go .
RUN go build main.go
EXPOSE 8080
ENTRYPOINT ["./main"]