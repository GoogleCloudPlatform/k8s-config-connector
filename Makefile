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

PROJECT_ID ?= $(shell gcloud config get-value project)
SHORT_SHA := $(shell git rev-parse --short=7 HEAD)
BUILDER_IMG ?= gcr.io/${PROJECT_ID}/builder:${SHORT_SHA}
CONTROLLER_IMG ?= gcr.io/${PROJECT_ID}/cnrm/controller:${SHORT_SHA}
RECORDER_IMG ?= gcr.io/${PROJECT_ID}/cnrm/recorder:${SHORT_SHA}
WEBHOOK_IMG ?= gcr.io/${PROJECT_ID}/cnrm/webhook:${SHORT_SHA}
DELETION_DEFENDER_IMG ?= gcr.io/${PROJECT_ID}/cnrm/deletiondefender:${SHORT_SHA}
UNMANAGED_DETECTOR_IMG ?= gcr.io/${PROJECT_ID}/cnrm/unmanageddetector:${SHORT_SHA}
# Detects the location of the user golangci-lint cache.
GOLANGCI_LINT_CACHE := /tmp/golangci-lint
# When updating this, make sure to update the corresponding action in
# ./github/workflows/lint.yaml
GOLANGCI_LINT_VERSION := v1.64.8

# Use Docker BuildKit when building images to allow usage of 'setcap' in
# multi-stage builds (https://github.com/moby/moby/issues/38132)
DOCKER_BUILD := DOCKER_BUILDKIT=1 docker build

CRD_OUTPUT_TMP := config/crds/tmp
CRD_OUTPUT_STAGING := config/crds/tmp/staging
CRD_OUTPUT_FINAL := config/crds/resources
PLATFORM ?= linux/amd64
OUTPUT_TYPE ?= type=docker

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

# Generate CRDs for direct controllers.
.PHONY: generate-crds
generate-crds:
	./dev/tasks/generate-crds

# Generate manifests e.g. CRD, RBAC etc.
.PHONY: manifests
manifests: generate
	make -C operator manifests
	rm -rf config/crds/resources
	rm -rf config/crds/tmp_resources
	go build -o bin/generate-crds ./scripts/generate-crds && ./bin/generate-crds -output-dir=config/crds/tmp_resources
	# add kustomize patches on all CRDs
	mkdir config/crds/resources
	cp config/crds/kustomization.yaml kustomization.yaml
	kustomize edit add resource config/crds/tmp_resources/*.yaml
	kustomize build -o config/crds/resources
	rm -rf config/crds/tmp_resources
	rm kustomization.yaml

	# for direct controllers
	dev/tasks/generate-crds

	# Generating cnrm cluster roles is dependent on the existence of directory
	# config/crds/resources with all the freshly generated CRDs.
	go run ./scripts/generate-cnrm-cluster-roles/main.go

	# Generating list of all supported GVKs is dependent on the existence of directory
	# config/crds/resources with all the freshly generated CRDs.
	go run ./scripts/generate-gvks/main.go -input-dir=config/crds/resources -output-file=pkg/gvks/supportedgvks/gvks_generated.go

# Format code
.PHONY: fmt
fmt:
	mockgcp/dev/fix-gofmt
	make -C operator fmt
	dev/tasks/fix-gofmt
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
	-ignore "**/testdata/**/*" \
	-ignore "experiments/**/testdata/**" \
	-ignore "pkg/gcpclients/generated/**" \
	./

.PHONY: lint
lint:
	mkdir -p ${GOLANGCI_LINT_CACHE}
	docker run --rm -v $(shell pwd):/app \
		-v ${GOLANGCI_LINT_CACHE}:/root/.cache/golangci-lint \
		-w /app golangci/golangci-lint:${GOLANGCI_LINT_VERSION}-alpine \
		golangci-lint run -v --timeout=10m

# Run go vet against code
.PHONY: vet
vet:
	make -C operator vet
	go vet -tags integration ./pkg/... ./cmd/... ./config/tests/...

# Generate code including the dcl (legacy)
.PHONY: generate-including-dcl
generate-including-dcl:
	go work vendor -o temp-vendor # So we can load DCL resources
	go generate ./pkg/dcl/schema/...
	rm -rf temp-vendor
	go generate ./pkg/apis/...
	make -C operator generate
	make fmt

# Generate code
.PHONY: generate
generate:
	go generate ./pkg/apis/...
	make -C operator generate
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

__tooling-image:
	docker buildx build build/tooling \
		--platform="$(PLATFORM)" \
		--output=$(OUTPUT_TYPE) \
		-t kcc-tooling

__controller-gen: __tooling-image
CONTROLLER_GEN=docker run --rm -v $(shell pwd):/wkdir kcc-tooling controller-gen

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

# Deploy controller only, this will skip CRD install in the configured K8s and usually runs much
# faster than "make deploy". It is useful if you only want to quickly apply code change in controller
.PHONY: deploy-controller-autopilot
deploy-controller-autopilot: docker-build docker-push
	kustomize build config/installbundle/releases/scopes/cluster/autopilot-withworkloadidentity | sed -e 's/$${PROJECT_ID?}/${PROJECT_ID}/g'| kubectl apply -f - ${CONTEXT_FLAG}

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
	go mod tidy -compat=1.23

# Should run all needed commands before any PR is sent out.
.PHONY: ready-pr
ready-pr: lint manifests resource-docs ensure fmt

# Upgrades dcl dependencies
.PHONY: upgrade-dcl
upgrade-dcl:
	go get github.com/GoogleCloudPlatform/declarative-resource-client-library
	make ensure

# Build all binaries
.PHONY: all-binary
all-binary: config-connector-bin deletiondefender manager-bin recorder unmanageddetector webhook gke-addon-poststart operator-manager-bin

# Build config-connector binary from cmd/config-connector
.PHONY: config-connector-bin
config-connector-bin:
	go build -o bin/config-connector github.com/GoogleCloudPlatform/k8s-config-connector/cmd/config-connector

# Build deletiondefender binary from cmd/deletiondefender
.PHONY: deletiondefender
deletiondefender:
	go build -o bin/deletiondefender github.com/GoogleCloudPlatform/k8s-config-connector/cmd/deletiondefender

# Build manager binary from cmd/manager
.PHONY: manager-bin
manager-bin:
	go build -o bin/manager github.com/GoogleCloudPlatform/k8s-config-connector/cmd/manager

# Build recorder binary from cmd/recorder
.PHONY: recorder
recorder:
	go build -o bin/recorder github.com/GoogleCloudPlatform/k8s-config-connector/cmd/recorder

# Build unmanageddetector binary from cmd/unmanageddetector
.PHONY: unmanageddetector
unmanageddetector:
	go build -o bin/unmanageddetector github.com/GoogleCloudPlatform/k8s-config-connector/cmd/unmanageddetector

# Build webhook binary from cmd/webhook
.PHONY: webhook
webhook:
	go build -o bin/webhook github.com/GoogleCloudPlatform/k8s-config-connector/cmd/webhook

# Build gke-addon-poststart binary from operator/cmd/gke_addon_poststart
.PHONY: gke-addon-poststart
gke-addon-poststart:
	go build -o bin/gke-addon-poststart github.com/GoogleCloudPlatform/k8s-config-connector/operator/cmd/gke_addon_poststart

# Build operator-manager binary from operator/cmd/manager
.PHONY: operator-manager-bin
operator-manager-bin:
	go build -o bin/operator-manager github.com/GoogleCloudPlatform/k8s-config-connector/operator/cmd/manager

# Build kcc manifests for both standard and autopilot clusters
.PHONY: all-manifests
all-manifests: crd-manifests rbac-manifests build-operator-manifests
	cp config/installbundle/release-manifests/crds.yaml config/installbundle/release-manifests/standard/crds.yaml
	cp config/installbundle/release-manifests/rbac.yaml config/installbundle/release-manifests/standard/rbac.yaml
	kustomize build config/installbundle/release-manifests/standard -o config/installbundle/release-manifests/standard/manifests.yaml
	cp config/installbundle/release-manifests/crds.yaml config/installbundle/release-manifests/autopilot/crds.yaml
	cp config/installbundle/release-manifests/rbac.yaml config/installbundle/release-manifests/autopilot/rbac.yaml
	kustomize build config/installbundle/release-manifests/autopilot -o config/installbundle/release-manifests/autopilot/manifests.yaml


# Build kcc manifests for standard GKE clusters
.PHONY: config-connector-manifests-standard
config-connector-manifests-standard: build-crd-manifests build-rbac-manifests build-operator-manifests
	cp config/installbundle/release-manifests/crds.yaml config/installbundle/release-manifests/standard/crds.yaml
	cp config/installbundle/release-manifests/rbac.yaml config/installbundle/release-manifests/standard/rbac.yaml
	kustomize build config/installbundle/release-manifests/standard -o config/installbundle/release-manifests/standard/manifests.yaml

# Build kcc manifests for autopilot clusters
.PHONY: config-connector-manifests-autopilot
config-connector-manifests-autopilot: build-crd-manifests build-rbac-manifests build-operator-manifests
	cp config/installbundle/release-manifests/crds.yaml config/installbundle/release-manifests/autopilot/crds.yaml
	cp config/installbundle/release-manifests/rbac.yaml config/installbundle/release-manifests/autopilot/rbac.yaml
	kustomize build config/installbundle/release-manifests/autopilot -o config/installbundle/release-manifests/autopilot/manifests.yaml

.PHONY: build-crd-manifests
build-crd-manifests:
	go run sigs.k8s.io/controller-tools/cmd/controller-gen@v0.14.0 crd paths="./operator/pkg/apis/..." output:crd:artifacts:config=operator/config/crd/bases
	kustomize build operator/config/crd -o config/installbundle/release-manifests/crds.yaml

.PHONY: build-rbac-manifests
build-rbac-manifests:
	kustomize build operator/config/rbac -o config/installbundle/release-manifests/rbac.yaml

.PHONY: build-operator-manifests
build-operator-manifests:
	make -C operator docker-build
	kustomize build operator/config/autopilot-manager -o config/installbundle/release-manifests/autopilot/manager.yaml
	kustomize build operator/config/manager -o config/installbundle/release-manifests/standard/manager.yaml

.PHONY: push-operator-manifest
push-operator-manifest:
	make -C operator docker-push

.PHONY: clean-operator-manifests
clean-release-manifests:
	rm config/installbundle/release-manifests/crds.yaml
	rm config/installbundle/release-manifests/rbac.yaml
	rm config/installbundle/release-manifests/standard/manager.yaml
	rm config/installbundle/release-manifests/autopilot/manager.yaml
	rm config/installbundle/release-manifests/standard/manifests.yaml
	rm config/installbundle/release-manifests/autopilot/manifests.yaml

.PHONY: deploy-kcc-standard
deploy-kcc-standard: docker-build docker-push config-connector-manifests-standard push-operator-manifest 
	kubectl apply -f config/installbundle/release-manifests/standard/manifests.yaml ${CONTEXT_FLAG}
	kustomize build config/installbundle/releases/scopes/cluster/withworkloadidentity | sed -e 's/$${PROJECT_ID?}/${PROJECT_ID}/g'| kubectl apply -f - ${CONTEXT_FLAG}

.PHONY: deploy-kcc-autopilot
deploy-kcc-autopilot: docker-build docker-push config-connector-manifests-autopilot push-operator-manifest
	kubectl apply -f config/installbundle/release-manifests/autopilot/manifests.yaml ${CONTEXT_FLAG}
	kustomize build config/installbundle/releases/scopes/cluster/autopilot-withworkloadidentity | sed -e 's/$${PROJECT_ID?}/${PROJECT_ID}/g'| kubectl apply -f - ${CONTEXT_FLAG}

.PHONY: powertool-tests
powertool-tests:
	cd scripts/github-actions/ && ./powertool-test.sh

.PHONY: e2e-scenario-tests
e2e-scenario-tests:
	dev/ci/presubmits/scenarios-tests

# indicate which samples testcases will be run
SAMPLE_TESTCASE ?= TestAllInSeries/samples
# indicate whether the testcases will be run again real/mock GCP
TEST_TARGET ?= mock

.PHONY: e2e-sample-tests
e2e-sample-tests:
	RUN_E2E=1 E2E_KUBE_TARGET=envtest E2E_GCP_TARGET=${TEST_TARGET} KCC_USE_DIRECT_RECONCILERS="ComputeForwardingRule" \ go test -test.count=1 -timeout 3600s -v ./tests/e2e -run ${SAMPLE_TESTCASE}

# orgnization ID for google.com
ORG_ID ?= 433637338589
# billing account for ACP
BILLING_ACCOUNT ?= 010E8D-490B6B-088E1C

.PHONY: operator-e2e-tests
operator-e2e-tests:
	export TEST_ORG_ID=${ORG_ID}
	export TEST_BILLING_ACCOUNT_ID=${BILLING_ACCOUNT}
	cd operator/tests/e2e/ && go test --project-id=${PROJECT_ID}

# Generate Go types for direct resources specified in the config files located under `dev/tools/controllerbuilder/config`.
.PHONY: generate-types
generate-types:
	cd dev/tools/controllerbuilder && \
	./generate-proto.sh && \
	find config -name "*.yaml" -type f | xargs -I {} go run . generate-types --config {}
	dev/tasks/fix-gofmt 
