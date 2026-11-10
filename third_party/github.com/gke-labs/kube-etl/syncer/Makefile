# Image URL to use all building/pushing image targets
PROJECT_ID ?= $(shell gcloud config get-value project)
# Image goes to Artifact Registry in IMAGE_LOCATION, a region (e.g.
# us-central1) or multi-region (us, europe, asia), matching deploy.sh's
# --image-location. It defaults to the region of DEST_LOCATION (the
# destination cluster's region or zone), derived with location_region from
# lib.sh. Set IMAGE_LOCATION or IMG directly to override. The krmsyncer Docker
# repository must exist in that location (deploy.sh creates it).
DEST_LOCATION ?=
IMAGE_LOCATION ?= $(if $(DEST_LOCATION),$(shell bash -c 'source ./lib.sh && location_region "$$1"' _ '$(DEST_LOCATION)'))
IMG ?= $(if $(IMAGE_LOCATION),$(IMAGE_LOCATION)-docker.pkg.dev/$(PROJECT_ID)/krmsyncer/controller:latest)

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

.PHONY: all
all: build

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: test
test: fmt vet ## Run tests.
	export KUBEBUILDER_ASSETS=$$(setup-envtest use -p path); \
	go test ./... -coverprofile cover.out

.PHONY: build
build: fmt vet ## Build manager binary.
	go build -o bin/manager main.go

.PHONY: run
run: fmt vet ## Run a controller from your host.
	go run ./main.go

.PHONY: test manifests generate
# Generate manifests (CRDs, RBAC, etc.)
manifests:
	controller-gen crd rbac:roleName=krmsyncer-manager-role paths="./..." output:crd:artifacts:config=config/crd/

# Generate code (deepcopy, etc.)
generate:
	controller-gen object paths="./..."

# Fail early if no image is set (DEST_LOCATION, IMAGE_LOCATION or IMG).
.PHONY: check-img
check-img:
	@test -n "$(IMG)" || { echo "Set DEST_LOCATION (e.g. make $(MAKECMDGOALS) DEST_LOCATION=us-central1-a), IMAGE_LOCATION or IMG" >&2; exit 1; }

# Build the docker image
docker-build: check-img test
	docker build . -t ${IMG}

# Push the docker image
docker-push: check-img
	docker push ${IMG}

# Install CRDs into a cluster
install: manifests
	kustomize build config/crd | kubectl apply -f -

# Uninstall CRDs from a cluster
uninstall: manifests
	kustomize build config/crd | kubectl delete -f -

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config.
# The controller:latest placeholder in config/manager/manager.yaml is replaced with ${IMG}.
deploy: check-img manifests
	kustomize build config/default | sed 's#image: controller:latest#image: ${IMG}#' | kubectl apply -f -

## Undeploy controller from the K8s cluster specified in ~/.kube/config.
undeploy:
	kustomize build config/default | kubectl delete --ignore-not-found=true -f -

##@ Integration Testing

.PHONY: test-integration
test-integration: manifests generate
	go test -v ./controllers/...

##@ Optional Integration Tests via Kind (Requires Kubeconfig)
#.PHONY: kind-up-all
#kind-up-all:
#	$(MAKE) kind-up CLUSTER_NAME=syncer-source KUBECONFIG_PATH=$(shell pwd)/test/integration/source.kubeconfig
#	$(MAKE) kind-up CLUSTER_NAME=syncer-dest KUBECONFIG_PATH=$(shell pwd)/test/integration/dest.kubeconfig
#	# Install CRDs to source cluster
#	kubectl --kubeconfig=$(shell pwd)/test/integration/source.kubeconfig apply -f config/crd/syncer.core.cnrm.cloud.google.com_syncers.yaml
#
#.PHONY: kind-down-all
#kind-down-all:
#	$(MAKE) kind-down CLUSTER_NAME=$(KIND_CLUSTER_SOURCE)
#	$(MAKE) kind-down CLUSTER_NAME=$(KIND_CLUSTER_DEST)
#
#.PHONY: kind-up
#kind-up:
#	@if ! kind get clusters | grep -q "^$(CLUSTER_NAME)$$"; then \
#		kind create cluster --name $(CLUSTER_NAME); \
#	else \
#		echo "Cluster $(CLUSTER_NAME) already exists"; \
#	fi
#	kind export kubeconfig --name $(CLUSTER_NAME) --kubeconfig $(KUBECONFIG_PATH)
#
#.PHONY: kind-down
#kind-down:
#	kind delete cluster --name $(CLUSTER_NAME)
