FROM golang:1.16

WORKDIR /app/tree
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build

# Run the executable
CMD ["./tree"]
