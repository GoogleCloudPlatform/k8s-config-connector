###### ----------- Inline manifest tool section ----------------------------

##@ Inline sidecar pod

INLINE_IMG ?= $(IMG_REGISTRY)/manifests-inline:$(IMG_VERSION)

.PHONY: build-inline
build-inline: fmt vet ## Build binary.
	$(GOPREFIX) go build -v -o bin/inline ./cmd/inline

.PHONY: clean-inline
clean-inline: ## clean binary.
	rm -fr bin/inline
	docker rmi ${IMG} .

# If you wish to build the manager image targeting other platforms you can use the --platform flag.
# (i.e. docker build --platform linux/arm64). However, you must enable docker buildKit for it.
# More info: https://docs.docker.com/develop/develop-images/build_enhancements/
.PHONY: docker-build-inline
docker-build-inline: ## Build docker image with the manager.
	docker build -t ${INLINE_IMG} -f Dockerfile.inline .

.PHONY: docker-push-inline
docker-push-inline: ## Push docker image with the manager.
	docker push ${INLINE_IMG}
