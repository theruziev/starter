FROM golang:1.24-alpine as builder

ARG PACKAGE_PREFIX=""
ARG GIT_TAG="dev"
ARG GIT_COMMIT=""

RUN apk update && apk upgrade && apk add --no-cache ca-certificates tzdata shadow
RUN update-ca-certificates
RUN useradd -u 10001 app


WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . /app/

RUN go build -o app -ldflags="\
    -X '$PACKAGE_PREFIX/internal/pkg/info.Version=$GIT_TAG' \
    -X '$PACKAGE_PREFIX/internal/pkg/info.Commit=$GIT_COMMIT'" \
    main.go

FROM scratch

COPY --from=builder /app/app /srv/app
COPY --from=builder /etc/ssl /etc/ssl/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/passwd

# import curl from current repository image
COPY --from=ghcr.io/tarampampam/curl:8.0.1 /bin/curl /bin/curl

# Docs: <https://docs.docker.com/engine/reference/builder/#healthcheck>
HEALTHCHECK --interval=5s --timeout=2s --retries=5 --start-period=2s CMD [ \
    "curl", "--fail", "http://127.0.0.1:8080/ready" \
]


USER app

ENTRYPOINT ["/srv/app", "server"]



