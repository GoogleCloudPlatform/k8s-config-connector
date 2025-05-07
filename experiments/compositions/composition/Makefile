# GCP project to use for pushing manifests
GCP_PROJECT_ID ?= $(shell gcloud config get-value project)

# Image URLs to use for building/pushing image targets
GIT_IMG_VERSION ?= $(shell git rev-parse --short HEAD)
IMG_VERSION ?= v0.0.1
IMG_REGISTRY ?= gcr.io/$(GCP_PROJECT_ID)
IMG ?= $(IMG_REGISTRY)/composition:$(IMG_VERSION)
JINJA_IMG ?= $(IMG_REGISTRY)/expander-jinja2:$(IMG_VERSION)
GETTER_IMG ?= $(IMG_REGISTRY)/expander-getter:$(IMG_VERSION)
GOPREFIX ?= GOWORK=off

# ENVTEST_K8S_VERSION refers to the version of kubebuilder assets to be downloaded by envtest binary.
ENVTEST_K8S_VERSION = 1.28.0
KIND_CLUSTER ?= kind



# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# CONTAINER_TOOL defines the container tool to be used for building images.
# Be aware that the target commands are only tested with Docker which is
# scaffolded by default. However, you might want to replace it to use other
# tools. (i.e. podman)
CONTAINER_TOOL ?= docker

# Setting SHELL to bash allows bash commands to be executed by recipes.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: all
all: build

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk command is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: manifests
manifests: ## Generate WebhookConfiguration, ClusterRole and CustomResourceDefinition objects.
	$(GOPREFIX) go run sigs.k8s.io/controller-tools/cmd/controller-gen@$(CONTROLLER_TOOLS_VERSION) rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases

.PHONY: generate
generate: ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
	$(GOPREFIX) go run sigs.k8s.io/controller-tools/cmd/controller-gen@$(CONTROLLER_TOOLS_VERSION) object:headerFile="hack/boilerplate.go.txt" paths="./..."
	$(GOPREFIX) go run -mod=readonly golang.org/x/tools/cmd/goimports@latest -w api cmd internal tests proto

.PHONY: fmt
fmt: ## Run go fmt against code.
	$(GOPREFIX) go fmt ./...
	GOFLAGS= $(GOPREFIX) go run github.com/google/addlicense@04bfe4ee9ca5764577b029acc6a1957fd1997153 -c "Google LLC" -l apache ./
	$(GOPREFIX) go run -mod=readonly golang.org/x/tools/cmd/goimports@latest -w api cmd internal tests proto

.PHONY: vet
vet: ## Run go vet against code.
	$(GOPREFIX) go vet ./...

.PHONY: test
test: manifests generate fmt vet envtest ## Run tests.
	KUBEBUILDER_ASSETS="$(shell $(ENVTEST) use $(ENVTEST_K8S_VERSION) --bin-dir $(LOCALBIN) -p path)" $(GOPREFIX) go test ./... -coverprofile cover.out

GOLANGCI_LINT = $(shell pwd)/bin/golangci-lint
GOLANGCI_LINT_VERSION ?= v1.64.8
golangci-lint:
	@[ -f $(GOLANGCI_LINT) ] || { \
	set -e ;\
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell dirname $(GOLANGCI_LINT)) $(GOLANGCI_LINT_VERSION) ;\
	}

.PHONY: lint
lint: golangci-lint ## Run golangci-lint linter & yamllint
	$(GOPREFIX) $(GOLANGCI_LINT) run

.PHONY: lint-fix
lint-fix: golangci-lint ## Run golangci-lint linter and perform fixes
	$(GOPREFIX) $(GOLANGCI_LINT) run --fix

##@ Build

.PHONY: build
build: manifests generate fmt vet ## Build manager binary.
	$(GOPREFIX) go build -o bin/manager cmd/main.go

.PHONY: run
run: manifests generate fmt vet ## Run a controller from your host.
	$(GOPREFIX) go run ./cmd/main.go

.PHONY: debug
debug: generate fmt vet manifests
	GO111MODULE=on dlv debug ./cmd/main.go

# If you wish to build the manager image targeting other platforms you can use the --platform flag.
# (i.e. docker build --platform linux/arm64). However, you must enable docker buildKit for it.
# More info: https://docs.docker.com/develop/develop-images/build_enhancements/
.PHONY: docker-build
docker-build: ## Build docker image with the manager.
	$(CONTAINER_TOOL) build -t ${IMG} .

.PHONY: docker-push
docker-push: ## Push docker image with the manager.
	$(CONTAINER_TOOL) push ${IMG}

# PLATFORMS defines the target platforms for the manager image be built to provide support to multiple
# architectures. (i.e. make docker-buildx IMG=myregistry/mypoperator:0.0.1). To use this option you need to:
# - be able to use docker buildx. More info: https://docs.docker.com/build/buildx/
# - have enabled BuildKit. More info: https://docs.docker.com/develop/develop-images/build_enhancements/
# - be able to push the image to your registry (i.e. if you do not set a valid value via IMG=<myregistry/image:<tag>> then the export will fail)
# To adequately provide solutions that are compatible with multiple platforms, you should consider using this option.
PLATFORMS ?= linux/arm64,linux/amd64,linux/s390x,linux/ppc64le
.PHONY: docker-buildx
docker-buildx: ## Build and push docker image for the manager for cross-platform support
	# copy existing Dockerfile and insert --platform=${BUILDPLATFORM} into Dockerfile.cross, and preserve the original Dockerfile
	sed -e '1 s/\(^FROM\)/FROM --platform=\$$\{BUILDPLATFORM\}/; t' -e ' 1,// s//FROM --platform=\$$\{BUILDPLATFORM\}/' Dockerfile > Dockerfile.cross
	- $(CONTAINER_TOOL) buildx create --name project-v3-builder
	$(CONTAINER_TOOL) buildx use project-v3-builder
	- $(CONTAINER_TOOL) buildx build --push --platform=$(PLATFORMS) --tag ${IMG} -f Dockerfile.cross .
	- $(CONTAINER_TOOL) buildx rm project-v3-builder
	rm Dockerfile.cross

##@ Testing

.PHONY: common-test-manifests
common-test-manifests: manifests kustomize
	mkdir -p release/test
	cd ../facade && $(KUSTOMIZE) build config/crd -o ../composition/release/test/facade_crds.yaml
	$(KUSTOMIZE) build config/crd -o release/test/crds.yaml

.PHONY: release-manifests
release-manifests: manifests kustomize
	$(KUSTOMIZE) build config/crd -o release/crds.yaml
	$(KUSTOMIZE) build config/release -o release/manifest.yaml

.PHONY: release-test-cc-manifests
release-test-cc-manifests: common-test-manifests
	cd config/manager && $(KUSTOMIZE) edit set image controller=${IMG}
	cd config/default && $(KUSTOMIZE) edit add patch --namespace system --name controller-manager --kind Deployment --patch "[{\"op\": \"add\", \"path\": \"/spec/template/spec/containers/1/imagePullPolicy\", \"value\": \"Always\"}]"
	cd config/expanders && $(KUSTOMIZE) edit set image expander-jinja2=${JINJA_IMG}
	cd config/expanders && $(KUSTOMIZE) edit set image expander-getter=${GETTER_IMG}
	$(KUSTOMIZE) build config/default -o release/test/cc-operator.yaml
	cd config/manager && $(KUSTOMIZE) edit set image controller=composition:latest
	cd config/default && $(KUSTOMIZE) edit remove patch --namespace system --name controller-manager --kind Deployment --patch "[{\"op\": \"add\", \"path\": \"/spec/template/spec/containers/1/imagePullPolicy\", \"value\": \"Always\"}]"
	cd config/expanders && $(KUSTOMIZE) edit set image expander-jinja2=expander-jinja2:latest
	cd config/expanders && $(KUSTOMIZE) edit set image expander-getter=expander-getter:latest

.PHONY: release-kind-manifests
release-test-kind-manifests: common-test-manifests
	cd config/manager && $(KUSTOMIZE) edit set image controller=${IMG}
	cd config/expanders && $(KUSTOMIZE) edit set image expander-jinja2=${JINJA_IMG}
	cd config/expanders && $(KUSTOMIZE) edit set image expander-getter=${GETTER_IMG}
	$(KUSTOMIZE) build config/default -o release/test/kind-operator.yaml
	cd config/manager && $(KUSTOMIZE) edit set image controller=composition:latest
	cd config/expanders && $(KUSTOMIZE) edit set image expander-jinja2=expander-jinja2:latest
	cd config/expanders && $(KUSTOMIZE) edit set image expander-getter=expander-getter:latest

.PHONY: deploy-kind
deploy-kind: release-test-kind-manifests docker-build docker-build-inline docker-build-expander-jinja2 docker-build-expander-getter 
	kind delete clusters ${KIND_CLUSTER} || true
	kind create cluster --name ${KIND_CLUSTER}
	kind load docker-image ${IMG} --name ${KIND_CLUSTER}
	# for testing samples/FirstComposition
	docker image pull nginx:1.16.0
	kind load docker-image nginx:1.16.0 --name ${KIND_CLUSTER}
	kind load docker-image ${INLINE_IMG} --name ${KIND_CLUSTER}
	kind load docker-image ${JINJA_IMG} --name ${KIND_CLUSTER}
	kind load docker-image ${GETTER_IMG} --name ${KIND_CLUSTER}
	kubectl --context kind-${KIND_CLUSTER} apply -f release/test/crds.yaml
	kubectl --context kind-${KIND_CLUSTER} apply -f release/test/facade_crds.yaml
	sleep 5 # for CRDs to be registered
	kubectl --context kind-${KIND_CLUSTER} apply -f release/test/kind-operator.yaml
	sleep 5
	kubectl --context kind-${KIND_CLUSTER} get pods -A


.PHONY: run-test
run-test: release-test-kind-manifests docker-build docker-build-inline docker-build-expander-jinja2 docker-build-expander-getter fmt #docker-push docker-push-inline docker-push-expander-jinja2 docker-push-expander-getter
	cd tests/testcases/ && $(GOPREFIX) go test -v -timeout 3600s -run "^${TESTCASE}" --images=${IMG},${INLINE_IMG},${JINJA_IMG},${GETTER_IMG}

.PHONY: e2e-test
e2e-test: release-test-kind-manifests docker-build docker-build-inline docker-build-expander-jinja2 docker-build-expander-getter fmt #docker-push docker-push-inline docker-push-expander-jinja2 docker-push-expander-getter
	cd tests/testcases/ && $(GOPREFIX) go test -v -timeout 3600s -run ./... --images=${IMG},${INLINE_IMG},${JINJA_IMG},${GETTER_IMG}

.PHONY: kcc-test
kcc-test: release-test-cc-manifests docker-build docker-build-inline docker-build-expander-jinja2 docker-build-expander-getter fmt docker-push docker-push-inline docker-push-expander-jinja2 docker-push-expander-getter
	cd tests/testcases/ && $(GOPREFIX) go test -v -timeout 3600s -run "^TestKCCSampleCloudSQL" --use-cc=true --images=${IMG},${INLINE_IMG},${JINJA_IMG},${GETTER_IMG}

##@ Deployment

ifndef ignore-not-found
  ignore-not-found = false
endif

.PHONY: install
install: manifests kustomize ## Install CRDs into the K8s cluster specified in ~/.kube/config.
	$(KUSTOMIZE) build config/crd | $(KUBECTL) apply -f -

.PHONY: uninstall
uninstall: manifests kustomize ## Uninstall CRDs from the K8s cluster specified in ~/.kube/config. Call with ignore-not-found=true to ignore resource not found errors during deletion.
	$(KUSTOMIZE) build config/crd | $(KUBECTL) delete --ignore-not-found=$(ignore-not-found) -f -

.PHONY: deploy
deploy: manifests kustomize ## Deploy controller to the K8s cluster specified in ~/.kube/config.
	cd config/manager && $(KUSTOMIZE) edit set image controller=${IMG}
	$(KUSTOMIZE) build config/default | $(KUBECTL) apply -f -
	cd config/manager && $(KUSTOMIZE) edit set image controller=composition:latest


.PHONY: undeploy
undeploy: ## Undeploy controller from the K8s cluster specified in ~/.kube/config. Call with ignore-not-found=true to ignore resource not found errors during deletion.
	$(KUSTOMIZE) build config/default | $(KUBECTL) delete --ignore-not-found=$(ignore-not-found) -f -

##@ Build Dependencies

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Tool Binaries
KUBECTL ?= kubectl
KUSTOMIZE ?= $(LOCALBIN)/kustomize
ENVTEST ?= $(LOCALBIN)/setup-envtest

## Tool Versions
KUSTOMIZE_VERSION ?= v5.2.1
CONTROLLER_TOOLS_VERSION ?= v0.14.0

.PHONY: kustomize
kustomize: $(KUSTOMIZE) ## Download kustomize locally if necessary. If wrong version is installed, it will be removed before downloading.
$(KUSTOMIZE): $(LOCALBIN)
	@if test -x $(LOCALBIN)/kustomize && ! $(LOCALBIN)/kustomize version | grep -q $(KUSTOMIZE_VERSION); then \
		echo "$(LOCALBIN)/kustomize version is not expected $(KUSTOMIZE_VERSION). Removing it before installing."; \
		rm -rf $(LOCALBIN)/kustomize; \
	fi
	test -s $(LOCALBIN)/kustomize || GOBIN=$(LOCALBIN) GO111MODULE=on $(GOPREFIX) go install sigs.k8s.io/kustomize/kustomize/v5@$(KUSTOMIZE_VERSION)

.PHONY: envtest
envtest: $(ENVTEST) ## Download envtest-setup locally if necessary.
$(ENVTEST): $(LOCALBIN)
	test -s $(LOCALBIN)/setup-envtest || GOBIN=$(LOCALBIN) $(GOPREFIX) go install sigs.k8s.io/controller-runtime/tools/setup-envtest@latest


###### ----------- include other make targets
include inline.mk
include expanders.mk