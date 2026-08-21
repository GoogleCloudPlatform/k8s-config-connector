#!/bin/bash
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

set -o errexit
set -o nounset
set -o pipefail

# Install build-essential and basic tools
sudo apt-get update
sudo apt-get install -y build-essential curl jq apt-transport-https ca-certificates gnupg

# Install gcloud
if ! command -v gcloud &> /dev/null; then
    echo "Installing gcloud..."
    curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo gpg --dearmor -o /usr/share/keyrings/cloud.google.com.gpg
    echo "deb [signed-by=/usr/share/keyrings/cloud.google.com.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | sudo tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
    sudo apt-get update && sudo apt-get install -y google-cloud-cli
fi

# Install kubectl
if ! command -v kubectl &> /dev/null; then
    echo "Installing kubectl..."
    sudo apt-get install -y kubectl
fi

# Install kustomize (using the version from Makefile)
KUSTOMIZE_VERSION=v5.3.0
if ! command -v kustomize &> /dev/null; then
    echo "Installing kustomize ${KUSTOMIZE_VERSION}..."
    curl -sSfL "https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2F${KUSTOMIZE_VERSION}/kustomize_${KUSTOMIZE_VERSION}_$(go env GOOS)_$(go env GOARCH).tar.gz" | tar xz
    sudo mv kustomize /usr/local/bin/
fi

# Install yq
if ! command -v yq &> /dev/null; then
    echo "Installing yq..."
    go install github.com/mikefarah/yq/v4@latest
    sudo ln -s "${HOME}/go/bin/yq" /usr/local/bin/yq || true
fi

# Setup envtest assets
echo "Setting up envtest assets..."
go run sigs.k8s.io/controller-runtime/tools/setup-envtest@release-0.22 use -p path

echo "Setup complete!"
echo "--------------------------------------------------------"
echo "To start testing, please perform the following steps:"
echo "1. Authenticate with Google Cloud:"
echo "   gcloud auth login"
echo "   gcloud auth application-default login"
echo ""
echo "2. Set up your GCP project:"
echo "   gcloud config set project YOUR_PROJECT_ID"
echo ""
echo "3. Export KUBEBUILDER_ASSETS in your shell:"
echo "   export KUBEBUILDER_ASSETS=\$(go run sigs.k8s.io/controller-runtime/tools/setup-envtest@release-0.22 use -p path)"
echo ""
echo "4. Run the e2e tests:"
echo "   E2E_KUBE_TARGET=envtest RUN_E2E=1 E2E_GCP_TARGET=real GOLDEN_OBJECT_CHECKS=1 GOLDEN_REQUEST_CHECKS=1 WRITE_GOLDEN_OUTPUT=1 go test -test.count=1 -timeout 3600s -v ./tests/e2e -run TestAllInSeries/fixtures/pubsubtopic"
echo "--------------------------------------------------------"
