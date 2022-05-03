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

export GOFLAGS=-mod=vendor

# enable multi-versions feature for CRDs
CRD_OPTIONS ?= "crd"

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

CONTROLLER_GEN := $(GOBIN)/controller-gen
CONTROLLER_GEN_VERSION := 0.6.2

all: manager

# Run tests
test: generate fmt vet manifests
	go test -v ./pkg/... -coverprofile cover.out

# Build manager binary
manager: generate fmt vet
	go build -o bin/manager cmd/manager/main.go

# Generate manifests e.g. CRD, RBAC etc.
manifests: controller-gen
	$(CONTROLLER_GEN) $(CRD_OPTIONS) paths="./pkg/apis/..." output:crd:artifacts:config=config/crd/bases
	go run scripts/generate-image-configmap/main.go

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Generate code
generate: controller-gen
	$(CONTROLLER_GEN) object paths="./..."

# Download controller-gen binary if not found at the path specified by
# CONTROLLER_GEN or if the binary found is not the version specified by
# CONTROLLER_GEN_VERSION.
controller-gen:
ifneq (Version: v$(CONTROLLER_GEN_VERSION), $(shell $(CONTROLLER_GEN) --version))
	GOFLAGS='' go install sigs.k8s.io/controller-tools/cmd/controller-gen@v$(CONTROLLER_GEN_VERSION)
endif
