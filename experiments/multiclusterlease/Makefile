
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

# Kind cluster name
KIND_CLUSTER_NAME ?= multiclusterlease-e2e

# test-e2e: Run the end-to-end tests.
# This is the primary target for E2E testing. It sets up a kind cluster,
# deploys the controller, runs the tests, and tears down the cluster.
.PHONY: test-e2e
test-e2e: kind-up deploy-e2e
	go test -v -tags=e2e ./controllers/...
	$(MAKE) kind-down

# kind-up: Create a kind cluster.
.PHONY: kind-up
kind-up:
	kind create cluster --name $(KIND_CLUSTER_NAME)

# kind-down: Delete the kind cluster.
.PHONY: kind-down
kind-down:
	kind delete cluster --name $(KIND_CLUSTER_NAME)

# kind-load: Build the controller image and load it into the kind cluster.
.PHONY: kind-load
kind-load: docker-build
	kind load docker-image --name $(KIND_CLUSTER_NAME) ${IMG}

# Redefine deploy for E2E tests to ensure the image is loaded first.
.PHONY: deploy-e2e
deploy-e2e: kind-load install
	@if [ -z "$${GCP_SA_KEY_PATH}" ]; then \
		echo "ERROR: GCP_SA_KEY_PATH environment variable must be set and point to a valid GCP SA key file"; \
		exit 1; \
	fi
	kubectl create namespace multiclusterlease-system --dry-run -o yaml | kubectl apply -f -
	kubectl create secret generic gcp-sa-key --from-file=key.json=${GCP_SA_KEY_PATH} -n multiclusterlease-system --dry-run -o yaml | kubectl apply -f -
	cd config/manager && kustomize edit set image controller=${IMG}
	kustomize build config/default | kubectl apply -f -
