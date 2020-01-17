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

all: docker-build docker-push deploy

docker-build-app:
	docker build ./app -t gcr.io/${PROJECT_ID}/bookstore/app:${SHORT_SHA}

docker-build-inventory:
	docker build ./inventory -t gcr.io/${PROJECT_ID}/bookstore/inventory:${SHORT_SHA}

docker-build-purchases:
	docker build ./purchases -t gcr.io/${PROJECT_ID}/bookstore/purchases:${SHORT_SHA}

docker-build-users:
	docker build ./users -t gcr.io/${PROJECT_ID}/bookstore/users:${SHORT_SHA}

docker-build: docker-build-app docker-build-inventory docker-build-purchases docker-build-users
	cp config/image_patch_template.yaml config/image_patch.yaml
	sed -i'' -e 's/\[PROJECT_ID\]/${PROJECT_ID}/g' config/image_patch.yaml
	sed -i'' -e 's/\[TAG\]/${SHORT_SHA}/g' config/image_patch.yaml

docker-push-app:
	docker push gcr.io/${PROJECT_ID}/bookstore/app:${SHORT_SHA}

docker-push-inventory:
	docker push gcr.io/${PROJECT_ID}/bookstore/inventory:${SHORT_SHA}

docker-push-purchases:
	docker push gcr.io/${PROJECT_ID}/bookstore/purchases:${SHORT_SHA}

docker-push-users:
	docker push gcr.io/${PROJECT_ID}/bookstore/users:${SHORT_SHA}

docker-push: docker-push-app docker-push-inventory docker-push-purchases docker-push-users

deploy:
	kustomize build config/ | kubectl apply -f -
