# Copyright 2025 Google LLC
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

REPO_ROOT := $(shell git rev-parse --show-toplevel)
MCL_MOD_PATH := $(shell go list -m -f '{{.Dir}}' github.com/gke-labs/multicluster-leader-election)
DOCKER_NET := kcc-net
GCS_MOCK_NAME := gcs-mock
GCS_BUCKET := kcc-mcl-test
MCL_IMG := multicluster-leader-election:latest
IMAGE_TAG := latest
KCC_CONTROLLER_IMG := controller:$(IMAGE_TAG)
KCC_OPERATOR_IMG := operator:$(IMAGE_TAG)
KUSTOMIZE := go run sigs.k8s.io/kustomize/kustomize/v5@v5.3.0

.PHONY: setup
setup: create-network start-gcs-mock build-images create-clusters install-mcl install-kcc patch-kcc

.PHONY: build-images
build-images:
	# Build MCL manager
	docker build -t $(MCL_IMG) $(MCL_MOD_PATH)
	# Build KCC Operator
	docker build -t $(KCC_OPERATOR_IMG) $(REPO_ROOT) -f $(REPO_ROOT)/operator/Dockerfile
	# Build KCC Controller (manager)
	printf "FROM golang:1.24.6 AS builder\nWORKDIR /go/src/github.com/GoogleCloudPlatform/k8s-config-connector\nCOPY . .\nRUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager github.com/GoogleCloudPlatform/k8s-config-connector/cmd/manager\nFROM gcr.io/distroless/static:nonroot\nWORKDIR /configconnector/\nCOPY --from=builder /go/src/github.com/GoogleCloudPlatform/k8s-config-connector/manager .\nUSER 1000\nENTRYPOINT [\"./manager\"]" | docker build -t $(KCC_CONTROLLER_IMG) -

.PHONY: create-network
create-network:
	docker network create $(DOCKER_NET) || true

.PHONY: start-gcs-mock
start-gcs-mock:
	docker stop $(GCS_MOCK_NAME) || true
	docker rm $(GCS_MOCK_NAME) || true
	docker run -d --name $(GCS_MOCK_NAME) --network $(DOCKER_NET) -p 4443:4443 fsouza/fake-gcs-server -scheme http
	# Wait for GCS mock to be ready
	for i in {1..30}; do curl -s http://localhost:4443/storage/v1/b && break || sleep 2; done
	# Create bucket
	curl -s -X POST "http://localhost:4443/storage/v1/b?project=test-project" -H "Content-Type: application/json" -d '{"name": "$(GCS_BUCKET)"}'

.PHONY: create-clusters
create-clusters:
	# Create Cluster A
	kind create cluster --name cluster-a --image kindest/node:v1.29.1 --wait 5m || true
	docker network connect $(DOCKER_NET) cluster-a-control-plane || true
	kind load docker-image $(MCL_IMG) --name cluster-a
	kind load docker-image $(KCC_CONTROLLER_IMG) --name cluster-a
	kind load docker-image $(KCC_OPERATOR_IMG) --name cluster-a

	# Create Cluster B
	kind create cluster --name cluster-b --image kindest/node:v1.29.1 --wait 5m || true
	docker network connect $(DOCKER_NET) cluster-b-control-plane || true
	kind load docker-image $(MCL_IMG) --name cluster-b
	kind load docker-image $(KCC_CONTROLLER_IMG) --name cluster-b
	kind load docker-image $(KCC_OPERATOR_IMG) --name cluster-b

.PHONY: install-mcl
install-mcl:
	for cluster in cluster-a cluster-b; do \
		kubectl --context kind-$$cluster apply -f $(MCL_MOD_PATH)/config/crd/bases; \
		$(KUSTOMIZE) build $(MCL_MOD_PATH)/config/default | \
			sed "s@image: .*@image: $(MCL_IMG)@g" | \
			kubectl --context kind-$$cluster apply -f -; \
		kubectl --context kind-$$cluster -n multiclusterlease-system wait --for=condition=Available deployment/multiclusterlease-controller-manager --timeout=60s; \
		kubectl --context kind-$$cluster -n multiclusterlease-system patch deployment multiclusterlease-controller-manager --patch \
			'{"spec":{"template":{"spec":{"containers":[{"name":"manager","args":["--gcs-bucket=$(GCS_BUCKET)","--verbose"],"env":[{"name":"STORAGE_EMULATOR_HOST","value":"http://$(GCS_MOCK_NAME):4443"}]}]}}}}'; \
	done

.PHONY: install-kcc
install-kcc:
	for cluster in cluster-a cluster-b; do \
		kubectl --context kind-$$cluster apply -f $(REPO_ROOT)/config/crds/resources; \
		$(KUSTOMIZE) build $(REPO_ROOT)/operator/config/default | \
			sed "s@image: .*@image: $(KCC_OPERATOR_IMG)@g" | \
			kubectl --context kind-$$cluster apply -f -; \
		kubectl --context kind-$$cluster wait -n configconnector-operator-system --for=jsonpath='{.status.readyReplicas}'=1 statefulset/configconnector-operator --timeout=60s; \
		sed "s/CLUSTER_ID/$$cluster/" cc.yaml.tmpl | kubectl --context kind-$$cluster apply -f -; \
	done

.PHONY: patch-kcc
patch-kcc:
	for cluster in cluster-a cluster-b; do \
		kubectl --context kind-$$cluster wait -n cnrm-system --for=condition=Ready pod -l cnrm.cloud.google.com/component=cnrm-controller-manager --timeout=120s; \
		kubectl --context kind-$$cluster -n cnrm-system patch statefulset cnrm-controller-manager --patch \
			'{"spec":{"template":{"spec":{"containers":[{"name":"manager","image":"$(KCC_CONTROLLER_IMG)","env":[{"name":"STORAGE_EMULATOR_HOST","value":"http://$(GCS_MOCK_NAME):4443"}]}]}}}}'; \
	done

.PHONY: test
test:
	bats mcl.bats

.PHONY: teardown
teardown:
	kind delete cluster --name cluster-a || true
	kind delete cluster --name cluster-b || true
	docker stop $(GCS_MOCK_NAME) || true
	docker rm $(GCS_MOCK_NAME) || true
	docker network rm $(DOCKER_NET) || true
