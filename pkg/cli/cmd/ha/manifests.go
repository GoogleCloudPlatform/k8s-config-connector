// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ha

import (
	"fmt"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

const (
	cnrmSystemNamespace      = "cnrm-system"
	workloadIdentityGSAAnnot = "iam.gke.io/gcp-service-account"
	componentLabelKey        = "cnrm.cloud.google.com/component"
	managerStatefulSetComponent = "cnrm-controller-manager"

	mclServiceAccountName    = "kcc-mcl"
	syncerServiceAccountName = "kcc-syncer"

	mclDeploymentName    = "multiclusterlease-controller-manager"
	syncerDeploymentName = "syncer-controller-manager"

	mclBinaryPath    = "/configconnector/mcl-manager"
	syncerBinaryPath = "/configconnector/krmsyncer"
)

type clusterManifestParams struct {
	clusterIdentity string
	project         string
	bucket          string
	leaseName       string
	leaseNamespace  string
	controllerImage string
}

// buildClusterHAManifests constructs the Kubernetes ServiceAccounts, RBAC, and Deployments
// needed to run MultiClusterLease and KRMSyncer on a cluster.
func buildClusterHAManifests(params clusterManifestParams) []client.Object {
	mclGSA := fmt.Sprintf("%s@%s.iam.gserviceaccount.com", mclServiceAccountName, params.project)
	syncerGSA := fmt.Sprintf("%s@%s.iam.gserviceaccount.com", syncerServiceAccountName, params.project)

	mclSA := &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{APIVersion: "v1", Kind: "ServiceAccount"},
		ObjectMeta: metav1.ObjectMeta{
			Name:      mclServiceAccountName,
			Namespace: cnrmSystemNamespace,
			Annotations: map[string]string{
				workloadIdentityGSAAnnot: mclGSA,
			},
		},
	}

	syncerSA := &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{APIVersion: "v1", Kind: "ServiceAccount"},
		ObjectMeta: metav1.ObjectMeta{
			Name:      syncerServiceAccountName,
			Namespace: cnrmSystemNamespace,
			Annotations: map[string]string{
				workloadIdentityGSAAnnot: syncerGSA,
			},
		},
	}

	mclClusterRole := &rbacv1.ClusterRole{
		TypeMeta: metav1.TypeMeta{APIVersion: "rbac.authorization.k8s.io/v1", Kind: "ClusterRole"},
		ObjectMeta: metav1.ObjectMeta{
			Name: "cnrm-mcl-controller",
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{"multicluster.core.cnrm.cloud.google.com"},
				Resources: []string{"multiclusterleases", "multiclusterleases/status"},
				Verbs:     []string{"get", "list", "watch", "create", "update", "patch", "delete"},
			},
		},
	}

	mclClusterRoleBinding := &rbacv1.ClusterRoleBinding{
		TypeMeta: metav1.TypeMeta{APIVersion: "rbac.authorization.k8s.io/v1", Kind: "ClusterRoleBinding"},
		ObjectMeta: metav1.ObjectMeta{
			Name: "cnrm-mcl-controller",
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     mclClusterRole.Name,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      mclServiceAccountName,
				Namespace: cnrmSystemNamespace,
			},
		},
	}

	syncerClusterRole := &rbacv1.ClusterRole{
		TypeMeta: metav1.TypeMeta{APIVersion: "rbac.authorization.k8s.io/v1", Kind: "ClusterRole"},
		ObjectMeta: metav1.ObjectMeta{
			Name: "cnrm-syncer-controller",
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{"syncer.gkelabs.io"},
				Resources: []string{"krmsyncers", "krmsyncers/status"},
				Verbs:     []string{"get", "list", "watch", "create", "update", "patch", "delete"},
			},
			{
				APIGroups: []string{"*.cnrm.cloud.google.com"},
				Resources: []string{"*"},
				Verbs:     []string{"get", "list", "watch", "create", "update", "patch", "delete"},
			},
		},
	}

	syncerClusterRoleBinding := &rbacv1.ClusterRoleBinding{
		TypeMeta: metav1.TypeMeta{APIVersion: "rbac.authorization.k8s.io/v1", Kind: "ClusterRoleBinding"},
		ObjectMeta: metav1.ObjectMeta{
			Name: "cnrm-syncer-controller",
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     syncerClusterRole.Name,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      syncerServiceAccountName,
				Namespace: cnrmSystemNamespace,
			},
		},
	}

	mclLabels := map[string]string{"app": "mcl-controller"}
	mclDeploy := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{APIVersion: "apps/v1", Kind: "Deployment"},
		ObjectMeta: metav1.ObjectMeta{
			Name:      mclDeploymentName,
			Namespace: cnrmSystemNamespace,
			Labels: map[string]string{
				componentLabelKey: "mcl-controller",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{MatchLabels: mclLabels},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{Labels: mclLabels},
				Spec: corev1.PodSpec{
					ServiceAccountName: mclServiceAccountName,
					Containers: []corev1.Container{
						{
							Name:    "manager",
							Image:   params.controllerImage,
							Command: []string{mclBinaryPath, fmt.Sprintf("--gcs-bucket=%s", params.bucket)},
						},
					},
				},
			},
		},
	}

	syncerLabels := map[string]string{"app": "syncer-controller"}
	syncerDeploy := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{APIVersion: "apps/v1", Kind: "Deployment"},
		ObjectMeta: metav1.ObjectMeta{
			Name:      syncerDeploymentName,
			Namespace: cnrmSystemNamespace,
			Labels: map[string]string{
				componentLabelKey: "syncer-controller",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{MatchLabels: syncerLabels},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{Labels: syncerLabels},
				Spec: corev1.PodSpec{
					ServiceAccountName: syncerServiceAccountName,
					Containers: []corev1.Container{
						{
							Name:    "manager",
							Image:   params.controllerImage,
							Command: []string{syncerBinaryPath},
						},
					},
				},
			},
		},
	}

	return []client.Object{
		mclSA,
		syncerSA,
		mclClusterRole,
		mclClusterRoleBinding,
		syncerClusterRole,
		syncerClusterRoleBinding,
		mclDeploy,
		syncerDeploy,
	}
}

func renderClusterYAML(contextName string, objs []client.Object, params clusterManifestParams) (string, error) {
	var b strings.Builder
	fmt.Fprintf(&b, "# Cluster context: %s\n", contextName)
	for _, obj := range objs {
		raw, err := yaml.Marshal(obj)
		if err != nil {
			return "", fmt.Errorf("marshaling %s/%s to YAML: %w", obj.GetObjectKind().GroupVersionKind().Kind, obj.GetName(), err)
		}
		b.WriteString("---\n")
		b.Write(raw)
	}
	b.WriteString("---\n")
	fmt.Fprintf(&b, `apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: %s
spec:
  experiments:
    multiClusterLease:
      leaseName: %s
      namespace: %s
      clusterCandidateIdentity: %s
      resourceReplicationMode: Status
`, configConnectorResourceName, params.leaseName, params.leaseNamespace, params.clusterIdentity)
	return b.String(), nil
}
