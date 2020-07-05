FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download

# CompileDaemon to automatically recompiled.
RUN go get github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon --build="go build commands/runserver.go" --command=./runserver