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

# This Dockerfile builds a thin image containing the multiclusterlease binary, which
# multi-cluster leader election needs on every cluster. The builder builds it from the
# fork in third_party/github.com/gke-labs/multicluster-leader-election, with the fork's
# own go.mod.
ARG BUILDER_IMG
ARG GKE_DISTROLESS_IMG

# Build the multiclusterlease binary
FROM --platform=$BUILDPLATFORM ${BUILDER_IMG} AS builder

# Copy the binary and other artifacts into a thin, distroless image (go/gke-distroless).
# The files stay owned by root: UID 1000 only needs to read and execute them.
FROM ${GKE_DISTROLESS_IMG} AS final
ARG TARGETARCH=amd64
# The upstream version that the fork was copied from (go.mod's require line)
ARG UPSTREAM_VERSION
LABEL org.opencontainers.image.title="multiclusterlease" \
      org.opencontainers.image.source="https://github.com/GoogleCloudPlatform/k8s-config-connector" \
      org.opencontainers.image.url="https://github.com/gke-labs/multicluster-leader-election" \
      org.opencontainers.image.version="${UPSTREAM_VERSION}"
WORKDIR /configconnector/
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/linux_${TARGETARCH}/multiclusterlease .
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/THIRD_PARTY_NOTICES/ THIRD_PARTY_NOTICES/
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/MIRRORED_LIBRARY_SOURCE/ MIRRORED_LIBRARY_SOURCE/
ENV PATH="/configconnector/:${PATH}"

# Set the user to user with UID 1000 for subsequent commands
USER 1000
ENTRYPOINT ["./multiclusterlease"]
