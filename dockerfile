FROM golang:1.20

WORKDIR /app

COPY go.mod .
COPY main.go .
COPY receipt.go .
COPY process_receipt.go .
COPY get_points.go .
COPY get_points_test.go . 

RUN go get

RUN go test -v

RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]