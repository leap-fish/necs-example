FROM golang:1.22-bookworm

RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD .. /app
WORKDIR /app
RUN go build -o app server/cmd/server/server.go
CMD ["/app/app"]