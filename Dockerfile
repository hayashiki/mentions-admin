#ARG GO_VERSION=1.15
#ARG NODE_VERSION=14.15.2

# Build Next
#FROM node:${NODE_VERSION}-alpine AS node-builder
FROM node:14-buster-slim as nodebuilder
WORKDIR /app
COPY web/package.json web/yarn.lock ./
RUN yarn install --frozen-lockfile
COPY web/ .
ENV NEXT_TELEMETRY_DISABLED=1
RUN yarn run export

# Build Go
FROM golang:1.15-buster as go-builder

WORKDIR /app
COPY go.* ./
RUN go mod download

COPY main.go main.go
#COPY cmd ./cmd
COPY pkg ./pkg
COPY --from=nodebuilder /app/dist ./web/dist
RUN go generate
RUN go build -mod=readonly -v -o server

#FROM debian:buster-slim
#RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
#    ca-certificates && \
#    rm -rf /var/lib/apt/lists/*
FROM gcr.io/distroless/base

# Copy the binary to the production image from the builder stage.
COPY --from=go-builder /app/server /app/server

# Run the web service on container startup.
CMD ["/app/server"]
