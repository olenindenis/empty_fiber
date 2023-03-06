FROM golang:1.20-buster as builder

WORKDIR /go/src/app

COPY . .

RUN go build -u -o api cmd/api/main.go
RUN go build -u -o cli cmd/migrations/main.go

FROM gcr.io/distroless/base

WORKDIR /bin

COPY --from=builder /go/src/app/api .
COPY --from=builder /go/src/app/cli .
COPY --from=builder /go/src/app/.env .
COPY --from=builder /go/src/app/database/. database/.

ENTRYPOINT ["./api"]