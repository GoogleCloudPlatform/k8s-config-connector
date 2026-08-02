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

COPY go.mod go.mod
COPY go.sum go.sum
COPY third_party/ third_party/
COPY mockgcp/ mockgcp/
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY apis/   apis/
COPY dev/metadata/   dev/metadata/
COPY pkg/    pkg/
COPY cmd/    cmd/
COPY config/ config/
COPY mockgcp/ mockgcp/
COPY operator/ operator/
COPY third_party/ third_party/
COPY version/ version/
COPY scripts/ scripts/

# Build the binaries from source in parallel. Note: the "-s -w" flags strip the tables
# needed for debuggers, but not the line numbers for panics, and -trimpath reduces binary size.
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -p 2 -ldflags "-s -w -buildid=" -o . \
      github.com/GoogleCloudPlatform/k8s-config-connector/cmd/deletiondefender \
      github.com/GoogleCloudPlatform/k8s-config-connector/cmd/webhook \
      github.com/GoogleCloudPlatform/k8s-config-connector/cmd/recorder \
      github.com/GoogleCloudPlatform/k8s-config-connector/cmd/unmanageddetector \
      github.com/GoogleCloudPlatform/k8s-config-connector/cmd/manager \
      github.com/GoogleCloudPlatform/k8s-config-connector/cmd/config-connector

RUN go mod vendor -o temp-vendor && \
    go run scripts/generate-third-party-licenses/main.go && \
    rm -rf temp-vendor
