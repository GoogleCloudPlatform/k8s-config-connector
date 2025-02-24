#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

export SERVICE=iam
export HTTP_HOST=iam.googleapis.com
export PROTO_SERVICE=google.iam.admin.v1.IAM
export PROTO_MESSAGE=google.iam.admin.v1.Role

RESOURCE=serviceaccount GCLOUD_COMMAND="gcloud iam service-accounts" ./_generate-resource.sh || true

RESOURCE=role GCLOUD_COMMAND="gcloud iam roles" ./_generate-resource.sh || true

RESOURCE=workloadidentitypool GCLOUD_COMMAND="gcloud iam workload-identity-pools" ./_generate-resource.sh || true

RESOURCE=oauthclient GCLOUD_COMMAND="gcloud iam oauth-clients" ./_generate-resource.sh || true
