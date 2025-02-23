#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

export SERVICE=dataproc
export PROTO_SERVICE=google.cloud.dataproc.v1.ClusterController
export CRD_VERSION=v1alpha1
export CRD_GROUP=${SERVICE}.cnrm.cloud.google.com


CRD_VERSION=v1beta1 CRD_KIND=DataprocCluster RESOURCE=Cluster PROTO_MESSAGE=google.cloud.dataproc.v1.Cluster ./_generate-resource.sh || true
