FROM golang:1.21-alpine as builder

ARG PACKAGE_PREFIX=""
ARG GIT_TAG="dev"
ARG GIT_COMMIT=""


WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . /app/

RUN go build -o app -ldflags="\
    -X '$PACKAGE_PREFIX/internal/pkg/git.Version=$GIT_TAG' \
    -X '$PACKAGE_PREFIX/internal/pkg/git.Commit=$GIT_COMMIT'" \
    main.go

FROM scratch

COPY --from=builder /app/app .

# import curl from current repository image
COPY --from=ghcr.io/tarampampam/curl:8.0.1 /bin/curl /bin/curl

# Docs: <https://docs.docker.com/engine/reference/builder/#healthcheck>
HEALTHCHECK --interval=5s --timeout=2s --retries=5 --start-period=2s CMD [ \
    "curl", "--fail", "http://127.0.0.1:8080/ready" \
]

ENTRYPOINT ["/app", "server"]



