FROM golang:1.24.4

WORKDIR /app
ENV GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o email-service ./cmd && ls -l

EXPOSE 4003

CMD ["./email-service"]
