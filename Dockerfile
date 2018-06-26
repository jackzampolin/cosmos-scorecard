FROM golang:alpine AS build-env

# Set up dependencies
ENV PACKAGES make git dep

# Set working directory for the build
WORKDIR /go/src/github.com/jackzampolin/cosmos-scorecard

# Add source files
COPY . .

# Install minimum necessary dependencies, build Cosmos SDK, remove packages
RUN apk add --no-cache $PACKAGES && \
    make dep && \
    make linux

# Final image
FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates
WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /go/src/github.com/jackzampolin/cosmos-scorecard/cosmos-scorecard-linux-amd64 /usr/bin/cosmos-scorecard

# Run gaiad by default, omit entrypoint to ease using container with gaiacli
CMD ["cosmos-scorecard"]
