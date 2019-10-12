# Start from golang base image
FROM golang:alpine AS build
# Install Git in Alpine
RUN apk add git
# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/localleon/octoprint-exporter
# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .
# Download all the dependencies and Install the package in the Builder
RUN go get -d -v ./... && go install -v ./...

# Build the actual Container Image
FROM alpine:latest
# Copy binarys from builder
COPY --from=build /go/bin/octoprint-exporter /bin/octoprint-exporter
# Set Entrypoint, Flags and Port
EXPOSE 9112
CMD ["/bin/octoprint-exporter","--config=/bin/config.yaml"]
