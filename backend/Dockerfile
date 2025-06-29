FROM golang:1.24

WORKDIR /app
COPY . .

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

RUN apt-get update && apt-get install -y gcc

RUN go mod download
RUN go build -o main .

EXPOSE 5000

CMD ["./main"]
