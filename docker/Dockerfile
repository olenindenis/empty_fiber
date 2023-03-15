#FROM golangci/golangci-lint:latest as linter
#ENV GO111MODULE=off
#WORKDIR /go/src/app
#
#COPY . .
#
#RUN golangci-lint run -u --skip-dirs cmd,docs,pkg,utils

FROM golang:1.20-buster as builder

WORKDIR /go/src/app

COPY . .

RUN go build -v -o api cmd/api/main.go
RUN go build -v -o migrations cmd/migrations/main.go

FROM builder as prod

WORKDIR /bin

COPY --from=builder /go/src/app/api .
COPY --from=builder /go/src/app/migrations .
COPY --from=builder /go/src/app/.env .
COPY --from=builder /go/src/app/database/. database/.

CMD api

FROM builder as dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /go/src/app

CMD air