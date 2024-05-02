# URL verifier tool

This is a simple tool to verify the URLs. It will check if the URLs are valid and if they are not, it will print the URLs that are invalid.

## Why this tool?

I needed a tool to verify URLs for a bigger project in Go. I decided to "wrap" it in a simple tool that can be also used in a CI/CD pipeline. Shoutout to [davidmytton](https://pkg.go.dev/github.com/davidmytton/url-verifier) for making a great library that I used to build the tool.

## How to use

There is a ready to use docker image available at [Docker Hub](https://hub.docker.com/r/skyheis/url-verifier-tool). You can simply run the following command to use the tool.

```bash
docker run --rm skyheis/url-verifier:latest-arm64 https://example.com https://example.com/invalid
```

The image have been built for `arm64` and `amd64` architectures. You can use the `latest-arm64` of `latest-amd64` tag to get the latest version of the tool.

## How to build

You can build the tool by running the following command.

```bash
go build -o url-verifier
```

In case of cross-compiling, you can use the following command.

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o url-verifier
```

Also, you can build the docker image by running the following command.

```bash
docker build --platform arm64 --build-arg GOARCH=arm64 -t url-verifier .
```
