FROM golang:1.20

WORKDIR /app

COPY go.mod .
COPY main.go .
COPY receipt.go .

RUN go get
RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]