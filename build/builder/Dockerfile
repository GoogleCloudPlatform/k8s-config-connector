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

# Lightweight builder stage that uses host-compiled Go binaries
FROM alpine:latest AS builder

WORKDIR /go/src/github.com/GoogleCloudPlatform/k8s-config-connector

COPY bin/deletiondefender deletiondefender
COPY bin/webhook webhook
COPY bin/recorder recorder
COPY bin/unmanageddetector unmanageddetector
COPY bin/manager manager
COPY bin/config-connector config-connector
COPY THIRD_PARTY_NOTICES/ THIRD_PARTY_NOTICES/
COPY MIRRORED_LIBRARY_SOURCE/ MIRRORED_LIBRARY_SOURCE/
