FROM golang:1.18 as base

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

FROM base AS development
    RUN go install github.com/cortesi/modd/cmd/modd@latest && \
        go install github.com/go-delve/delve/cmd/dlv@latest
    CMD modd -f ./cmd/server/modd.conf

FROM base as compiler
    COPY . .
    RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o jobs-manager ./cmd/server/main.go

FROM gcr.io/distroless/static:nonroot as release
WORKDIR /usr/app
COPY --from=compiler /workspace/jobs-manager .
COPY ./swagger ./swagger
USER 65532:65532

ENTRYPOINT ["./jobs-manager"]