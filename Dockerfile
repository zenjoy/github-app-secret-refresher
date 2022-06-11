FROM --platform=${BUILDPLATFORM} golang:1.16-alpine as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY cmd ./cmd
COPY internal ./internal
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o app cmd/main.go

FROM alpine:3.9.3
RUN apk add --no-cache bash
WORKDIR /app
COPY --from=builder /app/app app
ENTRYPOINT ["/app/app"]
