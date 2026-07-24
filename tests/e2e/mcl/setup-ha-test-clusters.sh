#!/usr/bin/env bash

# Copyright 2026 Google LLC
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
cd "${REPO_ROOT}"

export KUBECONFIG="${KUBECONFIG:-${REPO_ROOT}/.kubeconfig-mcl}"

CLUSTER_A="kcc-cluster-a"
CLUSTER_B="kcc-cluster-b"
NAMESPACE="cnrm-system"
LEASE_NAME="kcc-leader-lease"
TMP_DIR="/tmp/mcl-e2e-setup"

mkdir -p "${TMP_DIR}"

# 1. Build Binaries and Images
echo ">>> Building KCC and Fake MCL images..."
export CGO_ENABLED=0
go build -o "${TMP_DIR}/manager" ./cmd/manager

cat <<EOF > "${TMP_DIR}/fake-mcl-controller.go"
package main
import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"
	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)
func main() {
	var leaseName, namespace string
	flag.StringVar(&leaseName, "lease-name", "kcc-leader-lease", "Name of the lease")
	flag.StringVar(&namespace, "namespace", "cnrm-system", "Namespace of the lease")
	flag.Parse()
	s := runtime.NewScheme()
	scheme.AddToScheme(s)
	mclv1alpha1.AddToScheme(s)
	c, err := client.New(config.GetConfigOrDie(), client.Options{Scheme: s})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Fake MCL controller starting. Watching %s/%s\n", namespace, leaseName)
	for {
		ctx := context.Background()
		lease := &mclv1alpha1.MultiClusterLease{}
		err := c.Get(ctx, client.ObjectKey{Namespace: namespace, Name: leaseName}, lease)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		if lease.Spec.HolderIdentity != nil && *lease.Spec.HolderIdentity != "" {
			candidateIdentity := *lease.Spec.HolderIdentity
			candidateRenewTime := lease.Spec.RenewTime
			updateStatus := false
			if lease.Status.GlobalHolderIdentity == nil || *lease.Status.GlobalHolderIdentity == "" {
				updateStatus = true
			} else if *lease.Status.GlobalHolderIdentity == candidateIdentity {
				updateStatus = true
			} else {
				if lease.Status.GlobalRenewTime != nil {
					lastRenew, err := time.Parse(time.RFC3339Nano, *lease.Status.GlobalRenewTime)
					if err == nil {
						if time.Since(lastRenew) > 10*time.Second {
							fmt.Printf("Global leader %s timed out, allowing %s to take over\n", *lease.Status.GlobalHolderIdentity, candidateIdentity)
							updateStatus = true
						}
					}
				}
			}
			if updateStatus {
				renewTimeStr := candidateRenewTime.Format(time.RFC3339Nano)
				duration := *lease.Spec.LeaseDurationSeconds
				generation := lease.Generation
				patch := client.MergeFrom(lease.DeepCopy())
				lease.Status.GlobalHolderIdentity = &candidateIdentity
				lease.Status.GlobalRenewTime = &renewTimeStr
				lease.Status.GlobalLeaseDurationSeconds = &duration
				lease.Status.ObservedGeneration = &generation
				lease.Status.Conditions = []metav1.Condition{
					{
						Type: "BackendHealthy",
						Status: metav1.ConditionTrue,
						LastTransitionTime: metav1.Now(),
						Reason: "FakeSuccess",
						Message: "Fake MCL controller is running",
					},
				}
				err := c.Status().Patch(ctx, lease, patch)
				if err != nil {
					fmt.Printf("Error patching status: %v\n", err)
				} else {
					fmt.Printf("Global leader is now: %s\n", candidateIdentity)
				}
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
}
EOF
go build -o "${TMP_DIR}/fake-mcl" "${TMP_DIR}/fake-mcl-controller.go"

cat <<EOF > "${TMP_DIR}/Dockerfile.kcc"
FROM gcr.io/distroless/static:latest
COPY manager /manager
ENTRYPOINT ["/manager"]
EOF
docker build -t kcc-manager:latest -f "${TMP_DIR}/Dockerfile.kcc" "${TMP_DIR}"

cat <<EOF > "${TMP_DIR}/Dockerfile.mcl"
FROM gcr.io/distroless/static:latest
COPY fake-mcl /manager
ENTRYPOINT ["/manager"]
EOF
docker build -t fake-mcl-controller:latest -f "${TMP_DIR}/Dockerfile.mcl" "${TMP_DIR}"

setup_cluster() {
  local name=$1
  local context="kind-${name}"

  if ! kind get clusters | grep -q "^${name}$"; then
    echo "Creating kind cluster: ${name}..."
    kind create cluster --name "${name}"
  fi

  echo "Loading images into ${name}..."
  kind load docker-image kcc-manager:latest --name "${name}"
  kind load docker-image fake-mcl-controller:latest --name "${name}"

  echo "Installing essential CRDs in ${name}..."
  kubectl --context "${context}" apply -f operator/config/crd/bases/core.cnrm.cloud.google.com_configconnectors.yaml
  kubectl --context "${context}" apply -f tests/e2e/mcl/mcl_crd.yaml

  echo "Creating namespace ${name}..."
  kubectl --context "${context}" create namespace "${NAMESPACE}" || true

  # RBAC
  cat <<EOF | kubectl --context "${context}" apply -f -
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: kcc-admin-binding
subjects:
- kind: ServiceAccount
  name: default
  namespace: ${NAMESPACE}
roleRef:
  kind: ClusterRole
  name: cluster-admin
  apiGroup: rbac.authorization.k8s.io
EOF

  # Deploy KCC manager
  echo "Deploying KCC manager in ${name}..."
  cat <<EOF | kubectl --context "${context}" apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: cnrm-controller-manager
  namespace: ${NAMESPACE}
  labels:
    control-plane: controller-manager
spec:
  replicas: 1
  selector:
    matchLabels:
      control-plane: controller-manager
  template:
    metadata:
      labels:
        control-plane: controller-manager
    spec:
      containers:
      - name: manager
        image: kcc-manager:latest
        imagePullPolicy: IfNotPresent
        args:
        - --leader-election-type=multicluster
        - --scoped-namespace=${NAMESPACE}
        env:
        - name: HUB_KUBECONFIG
          value: /etc/hub/kubeconfig
        volumeMounts:
        - name: hub-kubeconfig
          mountPath: /etc/hub
      volumes:
      - name: hub-kubeconfig
        secret:
          secretName: kcc-lease-hub-kubeconfig
          optional: true
EOF

  # ConfigConnector object
  cat <<EOF | kubectl --context "${context}" apply -f -
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: configconnector.core.cnrm.cloud.google.com
spec:
  mode: cluster
  experiments:
    multiClusterLease:
      leaseName: ${LEASE_NAME}
      namespace: ${NAMESPACE}
      clusterCandidateIdentity: ${name}
EOF

# Deploy Fake MCL
echo "Deploying fake MCL controller in cluster ${name}..."
cat <<EOF | kubectl --context "kind-${name}" apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mcl-controller
  namespace: ${NAMESPACE}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: mcl-controller
  template:
    metadata:
      labels:
        app: mcl-controller
    spec:
      containers:
      - name: manager
        image: fake-mcl-controller:latest
        imagePullPolicy: IfNotPresent
        args:
        - --lease-name=${LEASE_NAME}
        - --namespace=${NAMESPACE}
EOF
}

# Setup Cluster A
setup_cluster "${CLUSTER_A}"

# Setup Cluster B
setup_cluster "${CLUSTER_B}"

# Cleanup local temp files
rm -rf "${TMP_DIR}"
echo "Setup Complete!"
