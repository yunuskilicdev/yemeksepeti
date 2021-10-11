FROM golang:1.17-alpine

RUN apk add --no-cache git
WORKDIR src/github.com/yunuskilicdev/yemeksepeti
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/yemeksepeti .

# This container exposes port 8080 to the ouside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./out/yemeksepeti"]