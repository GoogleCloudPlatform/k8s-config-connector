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

ARG BUILDER_IMG
ARG GKE_DISTROLESS_IMG

FROM ${BUILDER_IMG} AS builder

# Prepare a directory containing the binaries and other artifacts, and
# configure any required permissions
FROM alpine:latest AS packager
ARG BUILDER_IMG
WORKDIR /configconnector-operator/
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/operator-manager ./manager
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/gke_addon_poststart .
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/operator/channels/ channels/
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/operator/autopilot-channels/ autopilot-channels/
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/THIRD_PARTY_NOTICES/ THIRD_PARTY_NOTICES/
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/MIRRORED_LIBRARY_SOURCE/ MIRRORED_LIBRARY_SOURCE/

# Set user with UID 1000 as the owner of the directory
RUN chown 1000 -R /configconnector-operator

# Copy the directory into a thin, distroless image (go/gke-distroless)
FROM ${GKE_DISTROLESS_IMG} AS final
WORKDIR /configconnector-operator/
COPY --from=packager /configconnector-operator /configconnector-operator
ENV PATH="/configconnector-operator/:${PATH}"

# Set the user to user with UID 1000 for subsequent commands
USER 1000
ENTRYPOINT ["./manager"]
