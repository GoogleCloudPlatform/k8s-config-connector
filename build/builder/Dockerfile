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

# This Dockerfile builds an image containing builds of all the binaries built
# from source (manager, webhook, etc.)
FROM golang:1.26.5 AS builder

# Copy in the Go source code
WORKDIR /go/src/github.com/GoogleCloudPlatform/k8s-config-connector
COPY apis/   apis/
COPY dev/metadata/   dev/metadata/
COPY pkg/    pkg/
COPY cmd/    cmd/
COPY config/ config/
COPY mockgcp/ mockgcp/
COPY operator/ operator/
COPY third_party/ third_party/
COPY version/ version/
COPY scripts/generate-third-party-licenses scripts/generate-third-party-licenses
COPY scripts/resource-autogen scripts/resource-autogen
COPY go.mod go.mod
COPY go.sum go.sum
# Build the binaries from source. Note: the "-s -w" flags strip the tables
# needed for debuggers, but not the line numbers for panics
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o linux_amd64/deletiondefender github.com/GoogleCloudPlatform/k8s-config-connector/cmd/deletiondefender && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o linux_amd64/webhook github.com/GoogleCloudPlatform/k8s-config-connector/cmd/webhook && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o linux_amd64/recorder github.com/GoogleCloudPlatform/k8s-config-connector/cmd/recorder && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o linux_amd64/unmanageddetector github.com/GoogleCloudPlatform/k8s-config-connector/cmd/unmanageddetector && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o linux_amd64/manager github.com/GoogleCloudPlatform/k8s-config-connector/cmd/manager && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o linux_amd64/config-connector github.com/GoogleCloudPlatform/k8s-config-connector/cmd/config-connector

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o linux_arm64/deletiondefender github.com/GoogleCloudPlatform/k8s-config-connector/cmd/deletiondefender && \
    CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o linux_arm64/webhook github.com/GoogleCloudPlatform/k8s-config-connector/cmd/webhook && \
    CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o linux_arm64/recorder github.com/GoogleCloudPlatform/k8s-config-connector/cmd/recorder && \
    CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o linux_arm64/unmanageddetector github.com/GoogleCloudPlatform/k8s-config-connector/cmd/unmanageddetector && \
    CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o linux_arm64/manager github.com/GoogleCloudPlatform/k8s-config-connector/cmd/manager && \
    CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o linux_arm64/config-connector github.com/GoogleCloudPlatform/k8s-config-connector/cmd/config-connector

RUN go mod vendor -o temp-vendor && \
    go run scripts/generate-third-party-licenses/main.go && \
    rm -rf temp-vendor

FROM scratch AS artifacts
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/linux_amd64 /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/linux_amd64
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/linux_arm64 /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/linux_arm64
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/THIRD_PARTY_NOTICES /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/THIRD_PARTY_NOTICES
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/MIRRORED_LIBRARY_SOURCE /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/MIRRORED_LIBRARY_SOURCE

