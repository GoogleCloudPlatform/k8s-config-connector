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
FROM golang:1.22.2 AS builder

WORKDIR /go/src/github.com/GoogleCloudPlatform/k8s-config-connector

# Copy in the go.mod files and pre-download to enable caching
COPY go.mod go.mod
COPY go.sum go.sum
COPY mockgcp/go.mod mockgcp/go.mod
COPY third_party/github.com/hashicorp/terraform-provider-google-beta/go.mod third_party/github.com/hashicorp/terraform-provider-google-beta/go.mod
RUN go mod download

# Copy in the Go source code
COPY pkg/    pkg/
COPY cmd/    cmd/
COPY config/ config/
COPY mockgcp/ mockgcp/
COPY operator/ operator/
COPY third_party/ third_party/
COPY scripts/generate-third-party-licenses scripts/generate-third-party-licenses
COPY scripts/resource-autogen scripts/resource-autogen

# Build the binaries from source. Note: the "-s -w" flags strip the tables
# needed for debuggers, but not the line numbers for panics
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o deletiondefender github.com/GoogleCloudPlatform/k8s-config-connector/cmd/deletiondefender
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o webhook github.com/GoogleCloudPlatform/k8s-config-connector/cmd/webhook
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o recorder github.com/GoogleCloudPlatform/k8s-config-connector/cmd/recorder
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o unmanageddetector github.com/GoogleCloudPlatform/k8s-config-connector/cmd/unmanageddetector
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o manager github.com/GoogleCloudPlatform/k8s-config-connector/cmd/manager

RUN go mod vendor -o temp-vendor # So we can load license files
RUN go run scripts/generate-third-party-licenses/main.go
RUN rm -rf temp-vendor
