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

ENTRYPOINT ["/app", "version"]



