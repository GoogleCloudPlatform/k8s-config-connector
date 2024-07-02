
GCP_PROJECT_ID ?= $(shell gcloud config get-value project)

# Image URLs to use for building/pushing image targets
GIT_IMG_VERSION ?= $(shell git rev-parse --short HEAD)
IMG_VERSION ?= v0.0.1
IMG_REGISTRY ?= gcr.io/$(GCP_PROJECT_ID)
EXPANDER_IMG ?= $(IMG_REGISTRY)/expander-helm:$(IMG_VERSION)
EXPANDER_BINARY ?= helm
EXPANDER_SERVICE ?= composition-helm-v0-0-1

.PHONY: all
all: build

###### ----------- Helm Expander ----------------------------

##@ Helm expander pod

.PHONY: build
build:
	go build -v -o bin/${EXPANDER_BINARY} ./

.PHONY: clean
clean: ## clean binary.
	rm -fr bin/${EXPANDER_BINARY}
	docker rmi ${EXPANDER_IMG} .

# If you wish to build the manager image targeting other platforms you can use the --platform flag.
# (i.e. docker build --platform linux/arm64). However, you must enable docker buildKit for it.
# More info: https://docs.docker.com/develop/develop-images/build_enhancements/
.PHONY: docker-build
docker-build: #build ## Build docker image with the manager.
	docker build -t ${EXPANDER_IMG} -f Dockerfile .

.PHONY: docker-push
docker-push: ## Push docker image with the manager.
	docker push ${EXPANDER_IMG}

.PHONY: docker-run
docker-run: docker-build
	docker run -p 8443:8443 ${EXPANDER_IMG}

.PHONY: unit-test
unit-test: deploy-kind
	kubectl patch service -n composition-system ${EXPANDER_SERVICE} -p '{"spec":{"type":"LoadBalancer"}}'
	nodeip=$$(kubectl get nodes -o json  | jq '.items[0].status.addresses[0].address' | xargs echo );\
	nodeport=$$(kubectl get service -n composition-system ${EXPANDER_SERVICE} -o json | jq ".spec.ports[0].nodePort");\
	echo $$nodeip:$$nodeport; \
	go test -v --addr=$$nodeip:$$nodeport