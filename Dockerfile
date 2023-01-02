# Build the manager binary
FROM golang:1.18 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY controllers/ controllers/
COPY commons/ commons/
COPY services/ services/
COPY utilities/ utilities/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o bundle-rest-app main.go

# Use distroless as minimal base image to package the bundle-rest-app binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/bundle-rest-app .
USER 65532:65532

ENTRYPOINT ["/bundle-rest-app"]
