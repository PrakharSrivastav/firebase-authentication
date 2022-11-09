# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.19.1-alpine3.15 as builder

# Copy local code to the container image.
WORKDIR /app
COPY . .

# Build the command inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -o binary main.go

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3.15.0
RUN apk add --no-cache ca-certificates sed
# Copy the binary to the production image from the builder stage.
COPY  --from=builder /app/binary /binary
COPY  --from=builder /app/fb.json /fb.json
COPY  --from=builder /app/public/ /public
# Run the web service on container startup.

ENTRYPOINT ["/binary"]
