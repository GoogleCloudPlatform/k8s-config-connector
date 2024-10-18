#!/bin/bash -x
# Copyright 2024 Google LLC
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

export PROJECT_ID=compositions-${USER}
export REGION=us-central1
export CONFIG_CONTROLLER_NAME=compositions

create_project() {
    if gcloud projects describe ${PROJECT_ID} ; then
        echo "Reusing existing project: ${PROJECT_ID}"
        gcloud auth application-default set-quota-project ${PROJECT_ID}
    else
        echo "Creating project: ${PROJECT_ID}"
        gcloud projects create ${PROJECT_ID} --folder=78821025809
        gcloud auth application-default set-quota-project ${PROJECT_ID}
        gcloud beta billing projects link ${PROJECT_ID} --billing-account 010E8D-490B6B-088E1C
        gcloud services enable krmapihosting.googleapis.com container.googleapis.com cloudresourcemanager.googleapis.com
        gcloud services enable anthos.googleapis.com   serviceusage.googleapis.com
        gcloud compute networks create default --subnet-mode=auto || true
    fi
    gcloud config set project ${PROJECT_ID}
}

grant_cc_permissions() {
    export SA_EMAIL="$(kubectl get ConfigConnectorContext -n config-control \
        -o jsonpath='{.items[0].spec.googleServiceAccount}' 2> /dev/null)"
    gcloud projects add-iam-policy-binding "${PROJECT_ID}" \
        --member "serviceAccount:${SA_EMAIL}" \
        --role "roles/owner" \
        --project "${PROJECT_ID}"
}

create_config_controller() {
    if gcloud anthos config controller describe ${CONFIG_CONTROLLER_NAME} --location=${REGION} ; then
        echo "Reusing existing Config Controller: ${CONFIG_CONTROLLER_NAME}"
    else
        echo "Creating Config Controller: ${CONFIG_CONTROLLER_NAME}"
        gcloud anthos config controller create ${CONFIG_CONTROLLER_NAME} --location=${REGION}
    fi
}

point_kubeconfig_to_cc() {
    gcloud anthos config controller get-credentials ${CONFIG_CONTROLLER_NAME} --location ${REGION}
}

allow_custom_workloads() {
    kubectl patch k8sallowedresources.constraints.gatekeeper.sh block-workloads --patch '{"spec":{"enforcementAction":"dryrun"}}' --type merge
    kubectl label validatingwebhookconfigurations.admissionregistration.k8s.io gatekeeper-validating-webhook-configuration policycontroller.configmanagement.gke.io/managed-by-operator-
    kubectl patch validatingwebhookconfigurations.admissionregistration.k8s.io gatekeeper-validating-webhook-configuration --type=json -p '[ {"op":"remove","path":"/webhooks"} ]'
}

build_and_push() {
    (cd facade; make)
    (cd composition; make)
    (cd composition; make docker-build-inline)
    (cd composition; make docker-push-inline)
    (cd composition; make docker-build)
    (cd composition; make docker-push)
    (cd expanders/jinja-expander; make docker-build)
    (cd expanders/jinja-expander; make docker-push)
}

allow_image_pulls() {
    export defaultGCESA="$(gcloud iam service-accounts list --format=json | jq '.[] | select(.displayName == "Compute Engine default service account") | .email' | xargs echo)"
    export bucket="$(gsutil ls | grep artifacts.)"
    gsutil iam ch serviceAccount:${defaultGCESA}:roles/storage.objectViewer ${bucket}
}

create_project
create_config_controller
point_kubeconfig_to_cc
grant_cc_permissions
allow_custom_workloads
build_and_push
allow_image_pulls