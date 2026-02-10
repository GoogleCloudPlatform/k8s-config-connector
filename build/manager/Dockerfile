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

# This Dockerfile builds a thin image containing the manager binary
ARG BUILDER_IMG
ARG GKE_DISTROLESS_IMG

# Build the manager binary
FROM ${BUILDER_IMG} AS builder

# Prepare a directory containing the binary and other artifacts, and configure
# any required permissions
FROM alpine:latest AS packager
WORKDIR /configconnector/
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/manager .
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/THIRD_PARTY_NOTICES/ THIRD_PARTY_NOTICES/
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/MIRRORED_LIBRARY_SOURCE/ MIRRORED_LIBRARY_SOURCE/

# Allow the binary to bind to privileged ports (ports below 1024)
RUN apk --update add --no-cache libcap
RUN setcap cap_net_bind_service+eip manager

# Set user with UID 1000 as the owner of the directory
RUN chown 1000 /configconnector

# Copy the directory into a thin, distroless image (go/gke-distroless)
FROM ${GKE_DISTROLESS_IMG} AS final
WORKDIR /configconnector/
COPY --from=packager /configconnector /configconnector
ENV PATH="/configconnector/:${PATH}"

# Set the user to user with UID 1000 for subsequent commands
USER 1000
ENTRYPOINT ["./manager"]
