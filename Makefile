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
BUILDER_IMG ?= gcr.io/${PROJECT_ID}/builder:${SHORT_SHA}
CONTROLLER_IMG ?= gcr.io/${PROJECT_ID}/controller:${SHORT_SHA}
RECORDER_IMG ?= gcr.io/${PROJECT_ID}/recorder:${SHORT_SHA}
WEBHOOK_IMG ?= gcr.io/${PROJECT_ID}/webhook:${SHORT_SHA}
DELETION_DEFENDER_IMG ?= gcr.io/${PROJECT_ID}/deletiondefender:${SHORT_SHA}
UNMANAGED_DETECTOR_IMG ?= gcr.io/${PROJECT_ID}/unmanageddetector:${SHORT_SHA}

# Use Docker BuildKit when building images to allow usage of 'setcap' in
# multi-stage builds (https://github.com/moby/moby/issues/38132)
DOCKER_BUILD := DOCKER_BUILDKIT=1 docker build

ifneq ($(origin KUBECONTEXT), undefined)
CONTEXT_FLAG := --context ${KUBECONTEXT}
endif

.PHONY: all
all: test manager operator config-connector

# Run tests
.PHONY: test
test: generate fmt vet manifests
	./scripts/unit-test.sh

# Build config-connector binary
.PHONY: config-connector
config-connector: generate fmt vet
	./scripts/config-connector/build.sh

# Build the operator's manager binary
.PHONY: operator
operator:
	make -C operator manager

# Build manager binary
.PHONY: manager
manager: generate fmt vet
	go build -o bin/manager github.com/GoogleCloudPlatform/k8s-config-connector/cmd/manager

# Generate manifests e.g. CRD, RBAC etc.
.PHONY: manifests
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
.PHONY: fmt
fmt:
	mockgcp/dev/fix-gofmt
	make -C operator fmt
	go run -mod=readonly golang.org/x/tools/cmd/goimports@latest -w pkg cmd scripts config/tests
	# 04bfe4ee9ca5764577b029acc6a1957fd1997153 includes fix to not log "Skipped" for each skipped file
	GOFLAGS= go run github.com/google/addlicense@04bfe4ee9ca5764577b029acc6a1957fd1997153 -c "Google LLC" -l apache \
	-ignore ".build/**" -ignore "vendor/**" -ignore "third_party/**" \
	-ignore "config/crds/**" -ignore "config/cloudcodesnippets/**" \
	-ignore "**/*.html" -ignore "config/installbundle/components/clusterroles/cnrm_admin.yaml" \
	-ignore "config/installbundle/components/clusterroles/cnrm_viewer.yaml" \
	-ignore "operator/channels/**" \
	-ignore "operator/autopilot-channels/**" \
	-ignore "operator/config/crd/bases/**" \
	-ignore "operator/config/gke-addon/image_configmap.yaml" \
	-ignore "operator/config/rbac/cnrm_viewer_role.yaml" \
	-ignore "operator/vendor/**" \
	./

.PHONY: lint
lint:
	for f in `find pkg cmd -name "*.go"`; do golint -set_exit_status $$f || exit $?; done

# Run go vet against code
.PHONY: vet
vet:
	make -C operator vet
	go vet -tags integration ./pkg/... ./cmd/... ./config/tests/...

# Generate code
.PHONY: generate
generate:
	# Don't run go generate on `pkg/clients/generated` in the normal development flow due to high latency.
	# This path will be covered by `generate-go-client` target specifically.
	go mod vendor -o temp-vendor # So we can load DCL resources
	go generate $$(go list ./pkg/... ./cmd/... ./scripts/resource-autogen/... | grep -v ./pkg/clients/generated)
	rm -rf temp-vendor
	make fmt

# Build the docker images
.PHONY: docker-build
docker-build: docker-build-manager docker-build-recorder docker-build-webhook docker-build-deletiondefender docker-build-unmanageddetector

# build all the binaries into the builder docker image
.PHONY: docker-build-builder
docker-build-builder:
	$(DOCKER_BUILD) . -f build/builder/Dockerfile -t ${BUILDER_IMG}

# Build the manager docker image
.PHONY: docker-build-manager
docker-build-manager: docker-build-builder
	$(DOCKER_BUILD) -t ${CONTROLLER_IMG} --build-arg BUILDER_IMG=${BUILDER_IMG} - < build/manager/Dockerfile
	@echo "updating kustomize image patch file for manager resource"
	cp config/installbundle/components/manager/base/manager_image_patch_template.yaml config/installbundle/components/manager/base/manager_image_patch.yaml
	sed -i'' -e 's@image: .*@image: '"${CONTROLLER_IMG}"'@' ./config/installbundle/components/manager/base/manager_image_patch.yaml

# Build the recorder docker image
.PHONY: docker-build-recorder
docker-build-recorder: docker-build-builder
	$(DOCKER_BUILD) -t ${RECORDER_IMG} --build-arg BUILDER_IMG=${BUILDER_IMG} - < build/recorder/Dockerfile
	@echo "updating kustomize image patch file for recorder resource"
	cp config/installbundle/components/recorder/recorder_image_patch_template.yaml config/installbundle/components/recorder/recorder_image_patch.yaml
	sed -i'' -e 's@image: .*@image: '"${RECORDER_IMG}"'@' ./config/installbundle/components/recorder/recorder_image_patch.yaml

# Build the webhook docker image
.PHONY: docker-build-webhook
docker-build-webhook: docker-build-builder
	$(DOCKER_BUILD) -t ${WEBHOOK_IMG} --build-arg BUILDER_IMG=${BUILDER_IMG} - < build/webhook/Dockerfile
	@echo "updating kustomize image patch file for webhook resource"
	cp config/installbundle/components/webhook/webhook_image_patch_template.yaml config/installbundle/components/webhook/webhook_image_patch.yaml
	sed -i'' -e 's@image: .*@image: '"${WEBHOOK_IMG}"'@' ./config/installbundle/components/webhook/webhook_image_patch.yaml

.PHONY: docker-build-deletiondefender
docker-build-deletiondefender: docker-build-builder
	$(DOCKER_BUILD) -t ${DELETION_DEFENDER_IMG} --build-arg BUILDER_IMG=${BUILDER_IMG} - < build/deletiondefender/Dockerfile
	@echo "updating kustomize image patch file for deletion defender resource"
	cp config/installbundle/components/deletiondefender/deletiondefender_image_patch_template.yaml config/installbundle/components/deletiondefender/deletiondefender_image_patch.yaml
	sed -i'' -e 's@image: .*@image: '"${DELETION_DEFENDER_IMG}"'@' ./config/installbundle/components/deletiondefender/deletiondefender_image_patch.yaml

.PHONY: docker-build-unmanageddetector
docker-build-unmanageddetector: docker-build-builder
	$(DOCKER_BUILD) -t ${UNMANAGED_DETECTOR_IMG} --build-arg BUILDER_IMG=${BUILDER_IMG} - < build/unmanageddetector/Dockerfile
	@echo "updating kustomize image patch file for unmanaged detector resource"
	cp config/installbundle/components/unmanageddetector/unmanageddetector_image_patch_template.yaml config/installbundle/components/unmanageddetector/unmanageddetector_image_patch.yaml
	sed -i'' -e 's@image: .*@image: '"${UNMANAGED_DETECTOR_IMG}"'@' ./config/installbundle/components/unmanageddetector/unmanageddetector_image_patch.yaml

# Push the docker image
.PHONY: docker-push
docker-push:
	docker push ${CONTROLLER_IMG}
	docker push ${RECORDER_IMG}
	docker push ${WEBHOOK_IMG}
	docker push ${DELETION_DEFENDER_IMG}
	docker push ${UNMANAGED_DETECTOR_IMG}

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
.PHONY: deploy
deploy: manifests install
	kustomize build config/installbundle/releases/scopes/cluster/withworkloadidentity | sed -e 's/$${PROJECT_ID?}/${PROJECT_ID}/g'| kubectl apply -f - ${CONTEXT_FLAG}

# Install CRDs into a cluster
.PHONY: install
install: manifests
	kubectl apply -f config/crds/resources/ ${CONTEXT_FLAG}

# Deploy controller only, this will skip CRD install in the configured K8s and usually runs much
# faster than "make deploy". It is useful if you only want to quickly apply code change in controller
.PHONY: deploy-controller
deploy-controller: docker-build docker-push
	kustomize build config/installbundle/releases/scopes/cluster/withworkloadidentity | sed -e 's/$${PROJECT_ID?}/${PROJECT_ID}/g'| kubectl apply -f - ${CONTEXT_FLAG}

# Generate CRD go clients
.PHONY: generate-go-client
generate-go-client:
	./scripts/generate-go-crd-clients/generate-clients.sh

# Generate google3 docs
.PHONY: resource-docs
resource-docs:
	@go run ./scripts/generate-google3-docs/resource-reference/main.go
	@go run ./scripts/generate-google3-docs/resource-lists/main.go

# Run against the configured Kubernetes cluster in ~/.kube/config
.PHONY: run
run: generate fmt vet
	SERVICE_MAPPING_DIR=config/servicemappings/ go run ./cmd/manager/main.go

# Ensures dependencies are up-to-date
.PHONY: ensure
ensure:
	go mod tidy -compat=1.19

# Should run all needed commands before any PR is sent out.
.PHONY: ready-pr
ready-pr: manifests resource-docs generate-go-client

# Upgrades dcl dependencies
.PHONY: upgrade-dcl
upgrade-dcl:
	go get github.com/GoogleCloudPlatform/declarative-resource-client-library
	make ensure
