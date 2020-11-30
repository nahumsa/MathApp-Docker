FROM golang:1.15

RUN mkdir -p /mathapp
WORKDIR /mathapp

COPY go.mod .
COPY go.sum .
COPY src/ /mathapp
RUN go mod download
RUN go mod vendor

ENV GIN_MODE=release
ENV PORT=8080
RUN go build main.go

EXPOSE $PORT

ENTRYPOINT ["./main"]
