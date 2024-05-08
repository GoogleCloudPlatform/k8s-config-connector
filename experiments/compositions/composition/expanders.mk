## Variables
# TODO(barney-s): rename this to expander-jinja2 and the job/pod version to jinja2-cli
JINJA2_IMG ?= $(IMG_REGISTRY)/expander-gjinja2:$(IMG_VERSION)

###### ----------- Build protos --------------------------------------------

##@ Protos

PROTOC ?= /usr/bin/protoc
.PHONY: protoc
protoc: $(PROTOC) ## Download protoc if necessary
$(PROTOC):
	sudo apt install -y protobuf-compiler
	protoc --version
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: build-protos
build-protos: protoc
	$(PROTOC) --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/expander.proto

###### ----------- Expander client command---------------------------

##@ Expander expand client command

.PHONY: build-expand
build-expand: build-protos fmt vet ## Build binary.
	go build -v -o expanders/bin/expand ./expanders/expand

.PHONY: clean-expand
clean-expand: ## clean binary.
	rm -fr expanders/bin/expand

###### ----------- Jinja2 Expander ----------------------------

##@ Jinja2 expander pod

.PHONY: build-expander-jinja2
build-expander-jinja2: build-protos fmt vet ## Build binary.
	go build -v -o expanders/bin/jinja2 ./expanders/jinja2

.PHONY: clean-expander-jinja2
clean-expander-jinja2: ## clean binary.
	rm -fr expanders/bin/jinja2
	docker rmi ${JINJA2_IMG} .

# If you wish to build the manager image targeting other platforms you can use the --platform flag.
# (i.e. docker build --platform linux/arm64). However, you must enable docker buildKit for it.
# More info: https://docs.docker.com/develop/develop-images/build_enhancements/
.PHONY: docker-build-expander-jinja2
docker-build-expander-jinja2: build-expander-jinja2 ## Build docker image with the manager.
	docker build -t ${JINJA2_IMG} -f Dockerfile.jinja2.expander .

.PHONY: docker-push-expander-jinja2
docker-push-expander-jinja2: ## Push docker image with the manager.
	docker push ${JINJA2_IMG}

.PHONY: docker-run-expander-jinja2
docker-run-expander-jinja2: docker-build-expander-jinja2
	docker run -p 50051:50051 ${JINJA2_IMG}