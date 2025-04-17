# Copyright 2024 Google LLC
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


# syntax=docker/dockerfile:1

# ----------------- Build Container ---------------------------
# Build the go app.
# Explicitly set to latest vs golang:1.22
FROM golang:1.24.2 AS build-stage

# Set destination for COPY
WORKDIR /go/src/app

# Download Go modules
# https://docs.docker.com/reference/dockerfile/#copy
COPY ./ ./
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -v -o expander main.go


FROM golang:1.24.2 as bins

WORKDIR /workspace

# Version string to embed in built binary.
ARG HELM_VERSION=v3.16.2

# Install Helm
RUN wget https://get.helm.sh/helm-${HELM_VERSION}-linux-amd64.tar.gz -O /tmp/helm-${HELM_VERSION}-linux-amd64.tar.gz && \
  tar -zxvf /tmp/helm-${HELM_VERSION}-linux-amd64.tar.gz -C /tmp && \
  mv /tmp/linux-amd64/helm /usr/local/bin/helm && \
  rm -rf /tmp/linux-amd64 /tmp/helm-${HELM_VERSION}-linux-amd64.tar.gz

#ARG VERSION
#  -ldflags "-X kpt.dev/configsync/pkg/version.VERSION=${VERSION}" \

FROM gcr.io/distroless/static:latest as helm-sync
# Setting HOME ensures that whatever UID this ultimately runs as can write files.
ENV HOME=/tmp
WORKDIR /
COPY --from=bins /usr/local/bin/helm /usr/local/bin/helm
COPY --from=build-stage /go/src/app/expander .

ENTRYPOINT ["helm"]

# Switch to non-root user
USER 1000
