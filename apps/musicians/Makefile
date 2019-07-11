# Copyright 2019 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

PROJECT_ID := $(shell gcloud config get-value project)
SHORT_SHA := $(shell git rev-parse --short HEAD)
IMG ?= gcr.io/${PROJECT_ID}/user:${SHORT_SHA}

all: docker-build docker-push deploy

.user:
	go get ./user/...
	go build -o bin/user ./user/...

generate-manifests:
	sed -e 's@image: .*@image: '"${IMG}"'@' manifests/base/musicians-deployment-image-patch.tmpl.yaml > manifests/base/musicians-deployment-image-patch.yaml

destroy: generate-manifests
	kustomize build manifests/overlays/prod/ | kubectl delete --ignore-not-found -f -

deploy: generate-manifests
	kustomize build manifests/overlays/prod/ | kubectl apply -f -

docker-build:
	docker build . -t ${IMG}

docker-push:
	docker push ${IMG}
