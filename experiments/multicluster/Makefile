
# Image URL to use all building/pushing image targets
IMG ?= controller:latest

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

CONTROLLER_GEN := $(GOBIN)/controller-gen
CONTROLLER_GEN_VERSION := 0.10.0

all: manager

# Run tests
test: generate fmt vet manifests
	go test ./... -coverprofile cover.out

# Build manager binary
manager: generate fmt vet
	go build -o bin/manager main.go

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet manifests
	go run ./main.go

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
	$(CONTROLLER_GEN) rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases

# Run go fmt against code
fmt:
	go fmt ./...
	go run github.com/google/addlicense@04bfe4ee9ca5764577b029acc6a1957fd1997153 -c "Google LLC" -l apache \
	-ignore "config/crd/**" \
	-ignore "config/rbac/**" \
	./

# Run go vet against code
vet:
	go vet ./...

# Download controller-gen binary if not found at the path specified by
# CONTROLLER_GEN or if the binary found is not the version specified by
# CONTROLLER_GEN_VERSION.
.PHONY: controller-gen
controller-gen:
ifneq (Version: v$(CONTROLLER_GEN_VERSION), $(shell $(CONTROLLER_GEN) --version))
	GOFLAGS='' go install sigs.k8s.io/controller-tools/cmd/controller-gen@v$(CONTROLLER_GEN_VERSION)
endif

# Generate code
generate: controller-gen
	$(CONTROLLER_GEN) object paths="./..."

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
	set -e ;\
	CONTROLLER_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$CONTROLLER_GEN_TMP_DIR ;\
	go mod init tmp ;\
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.5 ;\
	rm -rf $$CONTROLLER_GEN_TMP_DIR ;\
	}
CONTROLLER_GEN=$(GOBIN)/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif
