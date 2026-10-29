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

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

const (
	cnrmSystemNamespace         = "cnrm-system"
	cnrmSystemSAGroup           = "system:serviceaccounts:cnrm-system"
	workloadIdentityGSAAnnot    = "iam.gke.io/gcp-service-account"
	componentLabelKey           = "cnrm.cloud.google.com/component"
	managerStatefulSetComponent = "cnrm-controller-manager"

	mclServiceAccountName    = "kcc-mcl"
	syncerServiceAccountName = "kcc-syncer"

	mclClusterRoleName    = "cnrm-mcl-controller"
	syncerClusterRoleName = "cnrm-syncer-controller"

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
	objs := make([]client.Object, 0, 8)
	objs = append(objs, buildMCLManifests(params)...)
	objs = append(objs, buildSyncerManifests(params)...)
	return objs
}

func buildMCLManifests(params clusterManifestParams) []client.Object {
	return []client.Object{
		buildWorkloadIdentityServiceAccount(mclServiceAccountName, params.project),
		buildMCLClusterRole(),
		buildClusterRoleBinding(mclClusterRoleName, mclServiceAccountName),
		buildMCLDeployment(params),
	}
}

func buildMCLClusterRole() *rbacv1.ClusterRole {
	return &rbacv1.ClusterRole{
		TypeMeta:   metav1.TypeMeta{APIVersion: "rbac.authorization.k8s.io/v1", Kind: "ClusterRole"},
		ObjectMeta: metav1.ObjectMeta{Name: mclClusterRoleName},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{"multicluster.core.cnrm.cloud.google.com"},
				Resources: []string{"multiclusterleases", "multiclusterleases/status"},
				Verbs:     []string{"get", "list", "watch", "create", "update", "patch", "delete"},
			},
		},
	}
}

func buildMCLDeployment(params clusterManifestParams) *appsv1.Deployment {
	command := []string{mclBinaryPath, fmt.Sprintf("--gcs-bucket=%s", params.bucket)}
	return buildControllerDeployment(mclDeploymentName, "mcl-controller", mclServiceAccountName, params.controllerImage, command)
}

func buildSyncerManifests(params clusterManifestParams) []client.Object {
	return []client.Object{
		buildWorkloadIdentityServiceAccount(syncerServiceAccountName, params.project),
		buildSyncerClusterRole(),
		buildClusterRoleBinding(syncerClusterRoleName, syncerServiceAccountName),
		buildSyncerDeployment(params),
	}
}

func buildSyncerClusterRole() *rbacv1.ClusterRole {
	return &rbacv1.ClusterRole{
		TypeMeta:   metav1.TypeMeta{APIVersion: "rbac.authorization.k8s.io/v1", Kind: "ClusterRole"},
		ObjectMeta: metav1.ObjectMeta{Name: syncerClusterRoleName},
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
}

func buildSyncerDeployment(params clusterManifestParams) *appsv1.Deployment {
	command := []string{syncerBinaryPath}
	return buildControllerDeployment(syncerDeploymentName, "syncer-controller", syncerServiceAccountName, params.controllerImage, command)
}

func renderClusterYAML(contextName string, objs []client.Object, cc *operatorv1beta1.ConfigConnector, params clusterManifestParams) (string, error) {
	allObjects := make([]client.Object, 0, len(objs)+1)
	allObjects = append(allObjects, objs...)
	allObjects = append(allObjects, buildRenderedConfigConnector(cc, params))

	var b strings.Builder
	fmt.Fprintf(&b, "# Cluster context: %s\n", contextName)
	for _, obj := range allObjects {
		raw, err := yaml.Marshal(obj)
		if err != nil {
			return "", fmt.Errorf("marshaling %s %q to YAML: %w", obj.GetObjectKind().GroupVersionKind().Kind, obj.GetName(), err)
		}
		b.WriteString("---\n")
		b.Write(raw)
	}
	return b.String(), nil
}

func buildRenderedConfigConnector(cc *operatorv1beta1.ConfigConnector, params clusterManifestParams) *operatorv1beta1.ConfigConnector {
	var spec operatorv1beta1.ConfigConnectorSpec
	if cc != nil {
		spec = *cc.Spec.DeepCopy()
	}
	setMultiClusterLeaseSpec(&spec, params)
	return &operatorv1beta1.ConfigConnector{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "core.cnrm.cloud.google.com/v1beta1",
			Kind:       "ConfigConnector",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: configConnectorResourceName,
		},
		Spec: spec,
	}
}

func setMultiClusterLeaseSpec(spec *operatorv1beta1.ConfigConnectorSpec, params clusterManifestParams) {
	if spec.Experiments == nil {
		spec.Experiments = &operatorv1beta1.CCExperiments{}
	}
	spec.Experiments.MultiClusterLease = &operatorv1beta1.MultiClusterLeaseSpec{
		LeaseName:                params.leaseName,
		Namespace:                params.leaseNamespace,
		ClusterCandidateIdentity: params.clusterIdentity,
		ResourceReplicationMode:  "Status",
	}
}

func buildWorkloadIdentityServiceAccount(saName, project string) *corev1.ServiceAccount {
	gsaEmail := fmt.Sprintf("%s@%s.iam.gserviceaccount.com", saName, project)
	return &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{APIVersion: "v1", Kind: "ServiceAccount"},
		ObjectMeta: metav1.ObjectMeta{
			Name:      saName,
			Namespace: cnrmSystemNamespace,
			Annotations: map[string]string{
				workloadIdentityGSAAnnot: gsaEmail,
			},
		},
	}
}

func buildClusterRoleBinding(roleName, saName string) *rbacv1.ClusterRoleBinding {
	return &rbacv1.ClusterRoleBinding{
		TypeMeta:   metav1.TypeMeta{APIVersion: "rbac.authorization.k8s.io/v1", Kind: "ClusterRoleBinding"},
		ObjectMeta: metav1.ObjectMeta{Name: roleName},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     roleName,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      saName,
				Namespace: cnrmSystemNamespace,
			},
			{
				Kind:     "Group",
				APIGroup: "rbac.authorization.k8s.io",
				Name:     cnrmSystemSAGroup,
			},
		},
	}
}

func buildControllerDeployment(name, component, saName, image string, command []string) *appsv1.Deployment {
	podLabels := map[string]string{"app": component}
	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{APIVersion: "apps/v1", Kind: "Deployment"},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: cnrmSystemNamespace,
			Labels: map[string]string{
				componentLabelKey: component,
			},
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{MatchLabels: podLabels},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{Labels: podLabels},
				Spec: corev1.PodSpec{
					ServiceAccountName: saName,
					Containers: []corev1.Container{
						{
							Name:    "manager",
							Image:   image,
							Command: command,
						},
					},
				},
			},
		},
	}
}
