# Copyright 2026 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This Dockefile builds a thin image containing the kompanion binary

# Build the kompanion binary
FROM golang:1.25.3 AS builder

# Copy in the Go source code
WORKDIR /go/src/github.com/GoogleCloudPlatform/k8s-config-connector
COPY experiments/kompanion/ experiments/kompanion/
COPY apis/ apis/
COPY dev/metadata/   dev/metadata/
COPY mockgcp/ mockgcp/
COPY operator/ operator/
COPY pkg/ pkg/
COPY third_party/ third_party/
COPY version/ version/
COPY config/ config/
COPY scripts/resource-autogen/ scripts/resource-autogen/
COPY scripts/generate-third-party-licenses scripts/generate-third-party-licenses
COPY go.mod  go.mod
COPY go.sum  go.sum

# Build the binary from source
WORKDIR /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/experiments/kompanion
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o ../../kompanion .

WORKDIR /go/src/github.com/GoogleCloudPlatform/k8s-config-connector

# Generate licenses
RUN go mod vendor -o temp-vendor # So we can load license files
RUN go run scripts/generate-third-party-licenses/main.go
RUN rm -rf temp-vendor

# Prepare a directory containing the binaries and other artifacts, and
# configure any required permissions
FROM alpine:latest AS packager
WORKDIR /kompanion/
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/kompanion .
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/THIRD_PARTY_NOTICES/ THIRD_PARTY_NOTICES/
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/MIRRORED_LIBRARY_SOURCE/ MIRRORED_LIBRARY_SOURCE/

# Set user with UID 1000 as the owner of the directory
RUN chown 1000 -R /kompanion

# Final stage
# Copy the directory into a thin, distroless image (go/gke-distroless)
FROM gcr.io/gke-release/gke-distroless/static:gke_distroless_20230307.00_p0 AS final

WORKDIR /kompanion/
COPY --from=packager /kompanion /kompanion
ENV PATH="/kompanion/:${PATH}"

# Set the user to user with UID 1000 for subsequent commands
USER 1000
ENTRYPOINT ["./kompanion"]
