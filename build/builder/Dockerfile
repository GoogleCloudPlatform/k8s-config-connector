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
FROM golang:1.26.11 AS builder

# Copy in the Go source code
WORKDIR /go/src/github.com/GoogleCloudPlatform/k8s-config-connector

COPY go.mod go.mod
COPY go.sum go.sum
COPY third_party/ third_party/
COPY mockgcp/ mockgcp/
# go.mod replaces the generator module with this directory, so its go.mod
# is needed to load the module graph.
COPY dev/tools/controllerbuilder/go.mod dev/tools/controllerbuilder/go.mod
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY apis/   apis/
COPY dev/metadata/   dev/metadata/
COPY pkg/    pkg/
COPY ptr/    ptr/
COPY cmd/    cmd/
COPY config/ config/
COPY mockgcp/ mockgcp/
COPY operator/ operator/
COPY third_party/ third_party/
COPY version/ version/
COPY scripts/generate-third-party-licenses scripts/generate-third-party-licenses
COPY scripts/resource-autogen scripts/resource-autogen
# Build the binaries from source in parallel. Note: the "-s -w" flags strip the tables
# needed for debuggers, but not the line numbers for panics, and -trimpath reduces binary size.
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -p 4 -ldflags "-s -w -buildid=" -o linux_amd64/ \
      github.com/GoogleCloudPlatform/k8s-config-connector/cmd/deletiondefender \
      github.com/GoogleCloudPlatform/k8s-config-connector/cmd/webhook \
      github.com/GoogleCloudPlatform/k8s-config-connector/cmd/recorder \
      github.com/GoogleCloudPlatform/k8s-config-connector/cmd/unmanageddetector \
      github.com/GoogleCloudPlatform/k8s-config-connector/cmd/manager \
      github.com/GoogleCloudPlatform/k8s-config-connector/cmd/config-connector

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -trimpath -p 4 -ldflags "-s -w -buildid=" -o linux_arm64/ \
      github.com/GoogleCloudPlatform/k8s-config-connector/cmd/deletiondefender \
      github.com/GoogleCloudPlatform/k8s-config-connector/cmd/webhook \
      github.com/GoogleCloudPlatform/k8s-config-connector/cmd/recorder \
      github.com/GoogleCloudPlatform/k8s-config-connector/cmd/unmanageddetector \
      github.com/GoogleCloudPlatform/k8s-config-connector/cmd/manager \
      github.com/GoogleCloudPlatform/k8s-config-connector/cmd/config-connector

# Build the MultiClusterLease controller and the KRM syncer, used for multi-cluster
# leader election, from the forks in third_party/github.com/gke-labs. Each fork is
# built as its own module, with its own go.mod and go.sum, not with KCC's module graph.
# GOWORK=off ignores a local go.work, and GOTOOLCHAIN=local builds them with this
# image's Go, like the other binaries. Each one ships in its own image
# (build/multiclusterlease and build/krmsyncer). The output names are explicit because
# the MultiClusterLease main package is cmd/manager, which would collide with KCC's
# manager binary.
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    for arch in amd64 arm64; do \
      GOWORK=off GOTOOLCHAIN=local CGO_ENABLED=0 GOOS=linux GOARCH=${arch} \
        go build -C third_party/github.com/gke-labs/multicluster-leader-election \
        -trimpath -ldflags "-s -w -buildid=" -o "${PWD}/linux_${arch}/multiclusterlease" ./cmd/manager && \
      GOWORK=off GOTOOLCHAIN=local CGO_ENABLED=0 GOOS=linux GOARCH=${arch} \
        go build -C third_party/github.com/gke-labs/kube-etl/syncer \
        -trimpath -ldflags "-s -w -buildid=" -o "${PWD}/linux_${arch}/krmsyncer" . || exit 1; \
    done

RUN --mount=type=cache,target=/go/pkg/mod \
    go run scripts/generate-third-party-licenses/main.go

FROM scratch AS artifacts
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/linux_amd64 /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/linux_amd64
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/linux_arm64 /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/linux_arm64
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/THIRD_PARTY_NOTICES /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/THIRD_PARTY_NOTICES
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/MIRRORED_LIBRARY_SOURCE /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/MIRRORED_LIBRARY_SOURCE

