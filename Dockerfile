#FROM ubuntu:latest
#LABEL authors="dale"
#
#ENTRYPOINT ["top", "-b"]

FROM golang:alpine AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY go.mod go.sum cmd/main.go ./
COPY . .
RUN go mod download
RUN go build -o main .
WORKDIR /dist
RUN cp /build/main .

FROM scratch
# CA 인증서 복사
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /dist/main .
COPY .env .
ENTRYPOINT ["/main"]