
# Image URL to use all building/pushing image targets
PROJECT_ID ?= $(shell gcloud config get-value project)
IMG ?= gcr.io/${PROJECT_ID}/multiclusterlease/controller:latest

CRD_OPTIONS ?= "crd"

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

all: manager

# Run tests
test: generate fmt vet manifests
	go test ./... -coverprofile cover.out

# Build manager binary
manager: generate fmt vet
	go build -o bin/manager cmd/manager/main.go

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet manifests
	go run cmd/manager/main.go

# Install CRDs into a cluster
install: manifests
	kustomize build config/crd | kubectl apply -f -

# Uninstall CRDs from a cluster
uninstall: manifests
	kustomize build config/crd | kubectl delete -f -

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy: manifests
	cd config/manager && kustomize edit set image controller=${IMG}
	kustomize build config/default | kubectl apply -f -

# Generate manifests e.g. CRD, RBAC etc.
manifests: controller-gen
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Generate code
generate: controller-gen
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."

# Build the docker image
docker-build: test
	docker build . -t ${IMG}

# Push the docker image
docker-push:
	docker push ${IMG}

# find or download controller-gen
# download controller-gen if necessary
controller-gen:
ifeq (, $(shell which controller-gen))
	@{ \
	set -e ;
	CONTROLLER_GEN_TMP_DIR=$(mktemp -d) ;
	cd $CONTROLLER_GEN_TMP_DIR ;
	go mod init tmp ;
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.15.0 ;
	rm -rf $CONTROLLER_GEN_TMP_DIR ;
	}
CONTROLLER_GEN=$(GOBIN)/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif

##@ E2E Testing

# Kind cluster names
KIND_CLUSTER_1_NAME ?= multiclusterlease-e2e-1
KIND_CLUSTER_2_NAME ?= multiclusterlease-e2e-2

# test-e2e: Run the single-cluster end-to-end tests.
.PHONY: test-e2e
test-e2e:
	$(MAKE) kind-up CLUSTER_NAME=$(KIND_CLUSTER_1_NAME)
	$(MAKE) deploy-to-cluster CLUSTER_NAME=$(KIND_CLUSTER_1_NAME)
	go test -v -tags=e2e ./controllers/...
	$(MAKE) kind-down CLUSTER_NAME=$(KIND_CLUSTER_1_NAME)

# test-e2e-multi: Run the multi-cluster end-to-end tests.
.PHONY: test-e2e-multi
test-e2e-multi: kind-up-all deploy-e2e-all
	go test -v -tags=e2e_multi ./controllers/...
	$(MAKE) kind-down-all

# kind-up-all: Create two kind clusters.
.PHONY: kind-up-all
kind-up-all:
	$(MAKE) kind-up CLUSTER_NAME=$(KIND_CLUSTER_1_NAME)
	$(MAKE) kind-up CLUSTER_NAME=$(KIND_CLUSTER_2_NAME)

# kind-down-all: Delete both kind clusters.
.PHONY: kind-down-all
kind-down-all:
	$(MAKE) kind-down CLUSTER_NAME=$(KIND_CLUSTER_1_NAME)
	$(MAKE) kind-down CLUSTER_NAME=$(KIND_CLUSTER_2_NAME)

# deploy-e2e-all: Deploy the controller to both clusters.
.PHONY: deploy-e2e-all
deploy-e2e-all:
	$(MAKE) deploy-to-cluster CLUSTER_NAME=$(KIND_CLUSTER_1_NAME)
	$(MAKE) deploy-to-cluster CLUSTER_NAME=$(KIND_CLUSTER_2_NAME)

# kind-up: Create a single kind cluster.
.PHONY: kind-up
kind-up:
	kind create cluster --name $(CLUSTER_NAME)

# kind-down: Delete a single kind cluster.
.PHONY: kind-down
kind-down:
	kind delete cluster --name $(CLUSTER_NAME)

# kind-load: Build the controller image and load it into a single kind cluster.
.PHONY: kind-load
kind-load: docker-build
	kind load docker-image --name $(CLUSTER_NAME) ${IMG}

# deploy-to-cluster: Deploy manifests to a specific cluster.
.PHONY: deploy-to-cluster
deploy-to-cluster: kind-load
	@if [ -z "${GCP_SA_KEY_PATH}" ]; then \
		echo "ERROR: GCP_SA_KEY_PATH environment variable must be set and point to a valid GCP SA key file"; \
		exit 1; \
	fi
	kubectl config use-context kind-$(CLUSTER_NAME)
	$(MAKE) install
	kubectl create namespace multiclusterlease-system --dry-run -o yaml | kubectl apply -f -
	kubectl create secret generic gcp-sa-key --from-file=key.json=${GCP_SA_KEY_PATH} -n multiclusterlease-system --dry-run -o yaml | kubectl apply -f -
	cd config/manager && kustomize edit set image controller=${IMG}
	kustomize build config/default | kubectl apply -f -

