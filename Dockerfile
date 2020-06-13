# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.14

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/alexanderkarlis/npi

# Build the golang-docker command inside the container.
RUN go install github.com/alexanderkarlis/npi

# Run the golang-docker command when the container starts.
ENTRYPOINT /go/bin/npi

# http server listens on port 8080.
EXPOSE 8080