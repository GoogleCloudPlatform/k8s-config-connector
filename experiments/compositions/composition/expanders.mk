## Variables

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
	docker rmi ${JINJA_IMG} .

# If you wish to build the manager image targeting other platforms you can use the --platform flag.
# (i.e. docker build --platform linux/arm64). However, you must enable docker buildKit for it.
# More info: https://docs.docker.com/develop/develop-images/build_enhancements/
.PHONY: docker-build-expander-jinja2
docker-build-expander-jinja2: build-expander-jinja2 ## Build docker image with the manager.
	docker build -t ${JINJA_IMG} -f Dockerfile.jinja2.expander .

.PHONY: docker-push-expander-jinja2
docker-push-expander-jinja2: ## Push docker image with the manager.
	docker push ${JINJA_IMG}

.PHONY: docker-run-expander-jinja2
docker-run-expander-jinja2: docker-build-expander-jinja2
	docker run -p 8443:8443 ${JINJA_IMG}

.PHONY: unit-test-expander-jinja2
unit-test-expander-jinja2: deploy-kind
	kubectl patch service -n composition-system composition-jinja2-v0-0-1 -p '{"spec":{"type":"LoadBalancer"}}'
	nodeip=$$(kubectl get nodes -o json  | jq '.items[0].status.addresses[0].address' | xargs echo );\
	nodeport=$$(kubectl get service -n composition-system composition-jinja2-v0-0-1 -o json | jq ".spec.ports[0].nodePort");\
	echo $$nodeip:$$nodeport; \
	cd expanders/jinja2 && go test -v --addr=$$nodeip:$$nodeport


###### ----------- Getter Expander ----------------------------

##@ Getter expander pod

.PHONY: build-expander-getter
build-expander-getter: build-protos fmt vet ## Build binary.
	go build -v -o expanders/bin/getter ./expanders/getter

.PHONY: clean-expander-getter
clean-expander-getter: ## clean binary.
	rm -fr expanders/bin/getter
	docker rmi ${GETTER_IMG} .

# If you wish to build the manager image targeting other platforms you can use the --platform flag.
# (i.e. docker build --platform linux/arm64). However, you must enable docker buildKit for it.
# More info: https://docs.docker.com/develop/develop-images/build_enhancements/
.PHONY: docker-build-expander-getter
docker-build-expander-getter: build-expander-getter ## Build docker image with the manager.
	docker build -t ${GETTER_IMG} -f Dockerfile.getter.expander .

.PHONY: docker-push-expander-getter
docker-push-expander-getter: ## Push docker image with the manager.
	docker push ${GETTER_IMG}

.PHONY: docker-run-expander-getter
docker-run-expander-getter: docker-build-expander-getter
	docker run -p 8443:8443 ${GETTER_IMG}

.PHONY: unit-test-expander-getter
unit-test-expander-getter: deploy-kind
	kubectl patch service -n composition-system composition-getter-v0-0-1 -p '{"spec":{"type":"LoadBalancer"}}'
	nodeip=$$(kubectl get nodes -o json  | jq '.items[0].status.addresses[0].address' | xargs echo );\
	nodeport=$$(kubectl get service -n composition-system composition-getter-v0-0-1 -o json | jq ".spec.ports[0].nodePort");\
	echo $$nodeip:$$nodeport; \
	cd expanders/getter && go test -v --addr=$$nodeip:$$nodeport