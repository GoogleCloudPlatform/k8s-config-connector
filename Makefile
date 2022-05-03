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

all: test manager operator config-connector

# Run tests
test: generate fmt vet manifests
	make -C operator test
	go test -v ./pkg/... ./cmd/... ./config/tests/... ./scripts/generate-go-crd-clients/... -coverprofile cover.out -count=1

# Build config-connector binary
config-connector: generate fmt vet
	./scripts/config-connector/build.sh

# Build the operator's manager binary
operator:
	make -C operator manager

# Build manager binary
manager: generate fmt vet
	go build -o bin/manager github.com/GoogleCloudPlatform/k8s-config-connector/cmd/manager

# Generate manifests e.g. CRD, RBAC etc.
manifests: generate
	make -C operator manifests
	rm -rf config/crds/resources
	rm -rf config/crds/tmp_resources
	go build -o bin/generate-crds ./scripts/generate-crds && ./bin/generate-crds -output-dir=config/crds/tmp_resources
	go run ./scripts/generate-cnrm-cluster-roles/main.go
	# add kustomize patches on all CRDs
	mkdir config/crds/resources
	cp config/crds/kustomization.yaml kustomization.yaml
	kustomize edit add resource config/crds/tmp_resources/*.yaml
	kustomize build -o config/crds/resources
	rm -rf config/crds/tmp_resources
	rm kustomization.yaml

# Format code
fmt: addlicense-binary
	make -C operator fmt
	goimports -w pkg cmd scripts config/tests
	addlicense -c "Google LLC" -l apache \
	-ignore "vendor/**" -ignore "third_party/**" \
	-ignore "config/crds/**" -ignore "config/cloudcodesnippets/**" \
	-ignore "**/*.html" -ignore "config/installbundle/components/clusterroles/cnrm_admin.yaml" \
	-ignore "config/installbundle/components/clusterroles/cnrm_viewer.yaml" \
	-ignore "operator/channels/**" \
	-ignore "operator/config/crd/bases/**" \
	-ignore "operator/config/gke-addon/image_configmap.yaml" \
	-ignore "operator/config/rbac/cnrm_viewer_role.yaml" \
	-ignore "operator/vendor/**" \
	./

lint:
	for f in `find pkg cmd -name "*.go"`; do golint -set_exit_status $$f || exit $?; done

# Run go vet against code
vet:
	make -C operator vet
	go vet -tags integration ./pkg/... ./cmd/... ./config/tests/...

# Generate code
generate:
	# Don't run go generate on `pkg/clients/generated` in the normal development flow due to high latency.
	# This path will be covered by `generate-go-client` target specifically.
	go generate $$(go list ./pkg/... ./cmd/... | grep -v ./pkg/clients/generated)
	make fmt

# Find or download addlicense binary
addlicense-binary:
ifeq (, $(shell which addlicense))
	GOFLAGS='' go install github.com/google/addlicense@v1.0.0
endif