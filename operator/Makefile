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

PROJECT_ID := $(shell gcloud config get-value project)
SHORT_SHA := $(shell git rev-parse --short=7 HEAD)
OPERATOR_IMG ?= gcr.io/${PROJECT_ID}/cnrm/operator:${SHORT_SHA}

# enable multi-versions feature for CRDs
CRD_OPTIONS ?= "crd"

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

CONTROLLER_GEN := $(GOBIN)/controller-gen
CONTROLLER_GEN_VERSION := 0.10.0

.PHONY: all
all: manager

# Run tests
.PHONY: test
test: generate fmt vet manifests
	go test -v ./pkg/... -coverprofile cover.out

# Build manager binary
.PHONY: manager
manager: generate fmt vet
	go build -o bin/manager cmd/manager/main.go

# Generate manifests e.g. CRD, RBAC etc.
.PHONY: manifests
manifests: controller-gen
	$(CONTROLLER_GEN) $(CRD_OPTIONS) paths="./pkg/apis/..." output:crd:artifacts:config=config/crd/bases
	go run scripts/generate-image-configmap/main.go

# Run go fmt against code
.PHONY: fmt
fmt:
	go fmt ./...

# Run go vet against code
.PHONY: vet
vet:
	go vet ./...

# Generate code
.PHONY: generate
generate: controller-gen
	$(CONTROLLER_GEN) object paths="./..."

# Download controller-gen binary if not found at the path specified by
# CONTROLLER_GEN or if the binary found is not the version specified by
# CONTROLLER_GEN_VERSION.
.PHONY: controller-gen
controller-gen:
ifneq (Version: v$(CONTROLLER_GEN_VERSION), $(shell $(CONTROLLER_GEN) --version))
	GOFLAGS='' go install sigs.k8s.io/controller-tools/cmd/controller-gen@v$(CONTROLLER_GEN_VERSION)
endif

# Install CRDs into a cluster
.PHONY: install
install: manifests
	kustomize build config/crd | kubectl apply -f -

# Uninstall CRDs from a cluster
.PHONY: uninstall
uninstall: manifests
	kustomize build config/crd | kubectl delete -f -

# Build the docker image
.PHONY: docker-build
docker-build:
	docker build .. -f Dockerfile -t ${OPERATOR_IMG}
	@echo "updating kustomize image patch file for manager"
	cp config/manager/manager_image_patch_template.yaml config/manager/manager_image_patch.yaml
	sed -i'' -e 's@image: .*@image: '"${OPERATOR_IMG}"'@' config/manager/manager_image_patch.yaml
	cp -f config/autopilot-manager/manager_image_patch_template.yaml config/autopilot-manager/manager_image_patch.yaml
	sed -i'' -e 's@image: .*@image: '"${OPERATOR_IMG}"'@' config/autopilot-manager/manager_image_patch.yaml

# Push the docker image
.PHONY: docker-push
docker-push:
	docker push ${OPERATOR_IMG}

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
# Note: run docker-build first to generate config/manager/manager_image_patch.yaml
.PHONY: deploy
deploy: manifests  # docker-build, docker-push also required, but not explicit to speed up deploy
	kustomize build config/default | kubectl apply -f -

# Destroy controller in the configured Kubernetes cluster in ~/.kube/config
.PHONY: destroy
destroy: manifests
	kustomize build config/default | kubectl delete -f -
