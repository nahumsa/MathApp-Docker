# Choosing a build
FROM golang:1.15
# Creating folder
RUN mkdir -p /mathapp
# Defining working repository
WORKDIR /mathapp
# Adding files for the container
COPY go.mod .
COPY go.sum .
COPY src/ /mathapp
# Downloading dependencies
RUN go mod download
RUN go mod vendor
# Setting up enviroment variables
ENV GIN_MODE=release
ENV PORT=8080
# Building the binary
RUN go build -o mathappbin
# Exposing the port
EXPOSE $PORT
# Running app
ENTRYPOINT ["./mathappbin"]
