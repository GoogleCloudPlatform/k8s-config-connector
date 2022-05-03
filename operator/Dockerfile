# Copyright 2022 Google LLC
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

# This Dockefile builds a thin image containing the manager binary

# Build the manager binary
FROM golang:1.17 AS builder

ENV GOFLAGS "-mod=vendor"

# Copy in the Go source code
WORKDIR /go/src/github.com/GoogleCloudPlatform/k8s-config-connector
COPY operator/pkg/      operator/pkg/
COPY operator/cmd/      operator/cmd/
COPY operator/channels/ operator/channels/
COPY pkg/ pkg/
COPY vendor/ vendor/
COPY scripts/generate-third-party-licenses scripts/generate-third-party-licenses
COPY go.mod  go.mod
COPY go.sum  go.sum

# Build the binary from source
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager github.com/GoogleCloudPlatform/k8s-config-connector/operator/cmd/manager
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o gke_addon_poststart github.com/GoogleCloudPlatform/k8s-config-connector/operator/cmd/gke_addon_poststart
RUN go run scripts/generate-third-party-licenses/main.go

# Build a specific version of kubectl to be used by the
# kubebuilder-declarative-pattern library.
RUN curl -fsSL https://dl.k8s.io/v1.19.14/bin/linux/amd64/kubectl > kubectl
RUN chmod a+rx kubectl

# Prepare a directory containing the binaries and other artifacts, and
# configure any required permissions
FROM alpine:latest AS packager
WORKDIR /configconnector-operator/
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/manager .
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/gke_addon_poststart .
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/operator/channels/ channels/
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/kubectl kubectl
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/THIRD_PARTY_NOTICES/ THIRD_PARTY_NOTICES/
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/MIRRORED_LIBRARY_SOURCE/ MIRRORED_LIBRARY_SOURCE/

# Set user with UID 1000 as the owner of the directory
RUN chown 1000 -R /configconnector-operator

# Copy the directory into a thin, distroless image (go/gke-distroless)
FROM gke.gcr.io/gke-distroless/static:latest AS final
WORKDIR /configconnector-operator/
COPY --from=packager /configconnector-operator /configconnector-operator
ENV PATH="/configconnector-operator/:${PATH}"

# Set the user to user with UID 1000 for subsequent commands
USER 1000
ENTRYPOINT ["./manager"]