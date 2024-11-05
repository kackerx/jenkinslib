FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod .

ENV GOPROXY=https://goproxy.cn,direct

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o myserver ./cmd/server.go

FROM scratch

ENV TZ=Asia/Shanghai

WORKDIR /app

COPY --from=builder /app/myserver .

CMD ["/app/myserver"]
