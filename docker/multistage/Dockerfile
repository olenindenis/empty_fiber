FROM golang:1.17-buster as builder

WORKDIR /go/src/app

COPY . .

RUN go build -v -o api cmd/api/main.go
RUN go build -v -o cli cmd/migrations/main.go

FROM gcr.io/distroless/base

WORKDIR /bin

COPY --from=builder /go/src/app/api .
COPY --from=builder /go/src/app/cli .
COPY --from=builder /go/src/app/.env .
COPY --from=builder /go/src/app/database/. database/.

ENTRYPOINT ["./api"]