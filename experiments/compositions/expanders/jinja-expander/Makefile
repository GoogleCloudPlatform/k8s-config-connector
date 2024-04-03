# GCP project to use for pushing manifests
GCP_PROJECT_ID ?= $(shell gcloud config get-value project)
# Image URL to use all building/pushing image targets
IMG_VERSION ?= v0.0.1.alpha
IMG ?= gcr.io/$(GCP_PROJECT_ID)/expander-jinja2:$(IMG_VERSION)

# Setting SHELL to bash allows bash commands to be executed by recipes.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: all
all: test

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


.PHONY: test
test: ## Run tests.
	rm -fr expanded/expanded
	docker run --rm -v ${PWD}/tests/test1/inputs:/inputs -v ${PWD}/tests/test1/expanded:/expanded gcr.io/cdcs-test/expander-jinja2:latest /inputs/template /inputs/values --format=yaml -o expanded/expanded
	diff tests/test1/expanded/expanded tests/test1/expanded/test-expected || echo "FAILED !!!!!!!!!"


##@ Build


.PHONY: clean
clean: ## clean binary.
	docker rmi ${IMG} .

# If you wish to build the manager image targeting other platforms you can use the --platform flag.
# (i.e. docker build --platform linux/arm64). However, you must enable docker buildKit for it.
# More info: https://docs.docker.com/develop/develop-images/build_enhancements/
.PHONY: docker-build
docker-build: ## Build docker image with the manager.
	docker build -t ${IMG} .

.PHONY: docker-push
docker-push: ## Push docker image with the manager.
	docker push ${IMG}