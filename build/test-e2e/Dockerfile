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

FROM golang:1.24.6

# Install dependencies
RUN apt-get update && apt-get install -y git curl python3 && rm -rf /var/lib/apt/lists/*

# Install gcloud (needed by some tests and for setup-gcloud compatibility)
RUN curl https://sdk.cloud.google.com | bash
ENV PATH=$PATH:/root/google-cloud-sdk/bin

# Install setup-envtest
RUN go install sigs.k8s.io/controller-runtime/tools/setup-envtest@release-0.22

# Pre-download envtest assets
# We run it once to populate the cache
RUN /go/bin/setup-envtest use -p path

WORKDIR /go/src/github.com/GoogleCloudPlatform/k8s-config-connector
COPY . .

# Prebuild the test binary to save time in each job
# We put it in /usr/local/bin so it's not overwritten by actions/checkout
RUN mkdir -p .build && go test -c -o /usr/local/bin/tests-e2e ./tests/e2e
