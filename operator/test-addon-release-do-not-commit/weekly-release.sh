#!/bin/bash
# Copyright 2022 Google LLC
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


set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"

ADDON_BUCKET=kcc-addon-internal
OPERATOR_INTERNAL_BUCKET=kcc-operator-internal
OPERATOR_PUBLIC_BUCKET=configconnector-operator
GKE_OPERATOR_IMAGE=gcr.io/gke-release/cnrm/operator
CNRM_EAP_CONTAINER_REGISTRY=gcr.io/cnrm-eap/cnrm/operator

## GKE addon release ###

# Copy contents from most recent private bucket
#ADDON_DIR=$(mktemp -d /tmp/addon-release.XXXXXXXX)
#mkdir $ADDON_DIR/cluster-bundle
#ADDON_DIR=$ADDON_DIR/cluster-bundle

# Build images and manifests
KUSTOMIZE="go run -mod=mod sigs.k8s.io/kustomize/kustomize/v5@v5.0.1"
cd $REPO_ROOT/operator
mkdir -p 1.126.2/
ADDON_DIR=$(pwd)/1.126.2
${KUSTOMIZE} build config/gke-addon > $ADDON_DIR/configconnector-operator.yaml
cp config/gke-addon/image_configmap.yaml  $ADDON_DIR/image_configmap.yaml

echo $ADDON_DIR
#cat $ADDON_DIR/configconnector-operator.yaml

# Extract the version
VERSION=$(grep -oP -m 1 'cnrm.cloud.google.com/operator-version: \K.+' $ADDON_DIR/configconnector-operator.yaml)
echo $VERSION

# Extract the short SHA
SHORT_SHA=$(grep -oP -m 1 "gcr.io/maqiuyu-kcc-test-2/cnrm/operator:\K.+" $ADDON_DIR/configconnector-operator.yaml)
echo $SHORT_SHA

# Extract the digest
DIGEST=$(gcloud container images describe gcr.io/maqiuyu-kcc-test-2/cnrm/operator:${SHORT_SHA} --format="get(image_summary.digest)")
echo $DIGEST

# Prepare google3 addon clusterbundle submission
CLIENT_NAME=configconnector_addon_20251029
p4 g4d -f $CLIENT_NAME
GOOGLE3_TARGET_DIR=/google/src/cloud/$USER/$CLIENT_NAME/google3/cloud/kubernetes/distro/components/configconnector/

cp -a $ADDON_DIR/ $GOOGLE3_TARGET_DIR/
cd $GOOGLE3_TARGET_DIR/1.126.2
#sed -i "s/version: [0-9]\+\.[0-9]\+\.[0-9]\+.*/version: $VERSION/g" kcc-component-builder.yaml
sed -i "/${GKE_OPERATOR_IMAGE//\//\\/}/ s/$/@${DIGEST}/" configconnector-operator.yaml
p4 reopen
p4 change --desc "Updating ConfigConnector clusterbundle to $VERSION

NO_BUG=Upgrading the version of ConfigConnector clusterbundle."
