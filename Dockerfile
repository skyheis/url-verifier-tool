# syntax=docker/dockerfile:1
FROM --platform=arm64 golang:1.22

# Set destination for COPY
WORKDIR /app

ARG GOARCH

COPY . ./

RUN go mod download
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o url-verifier main.go

FROM alpine:3.19
COPY --from=0 /app/url-verifier /url-verifier

ENTRYPOINT [ "/url-verifier" ]
