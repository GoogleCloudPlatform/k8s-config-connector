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

ARG GKE_DISTROLESS_IMG=gcr.io/gke-release/gke-distroless/static:gke_distroless_20260207.00_p0

# Prepare a directory containing the precompiled binaries and other artifacts
FROM alpine:latest AS packager
WORKDIR /configconnector-operator/
COPY bin/operator manager
COPY bin/gke_addon_poststart gke_addon_poststart
COPY operator/channels/ channels/
COPY operator/autopilot-channels/ autopilot-channels/
COPY THIRD_PARTY_NOTICES/ THIRD_PARTY_NOTICES/
COPY MIRRORED_LIBRARY_SOURCE/ MIRRORED_LIBRARY_SOURCE/

# Set executable permissions and user with UID 1000 as the owner of the directory
RUN chmod +x manager gke_addon_poststart && chown 1000 -R /configconnector-operator

# Copy the directory into a thin, distroless image (go/gke-distroless)
FROM ${GKE_DISTROLESS_IMG} AS final
WORKDIR /configconnector-operator/
COPY --from=packager /configconnector-operator /configconnector-operator
ENV PATH="/configconnector-operator/:${PATH}"

# Set the user to user with UID 1000 for subsequent commands
USER 1000
ENTRYPOINT ["./manager"]
