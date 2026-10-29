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

// manifestBuilder constructs and renders the Kubernetes manifests required for multi-cluster HA.
type manifestBuilder struct {
	leaseName      string
	leaseNamespace string
	bucket         string
}

type controllerComponent struct {
	deploymentName     string
	componentLabel     string
	serviceAccountName string
	clusterRoleName    string
	policyRules        []rbacv1.PolicyRule
	container          corev1.Container
}

func newManifestBuilder(leaseName, leaseNamespace, bucket string) *manifestBuilder {
	return &manifestBuilder{
		leaseName:      leaseName,
		leaseNamespace: leaseNamespace,
		bucket:         bucket,
	}
}

func (b *manifestBuilder) buildClusterObjects(project, controllerImage string) []client.Object {
	mclComponent := controllerComponent{
		deploymentName:     mclDeploymentName,
		componentLabel:     "mcl-controller",
		serviceAccountName: mclServiceAccountName,
		clusterRoleName:    mclClusterRoleName,
		policyRules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{"multicluster.core.cnrm.cloud.google.com"},
				Resources: []string{"multiclusterleases", "multiclusterleases/status"},
				Verbs:     []string{"get", "list", "watch", "create", "update", "patch", "delete"},
			},
		},
		container: corev1.Container{
			Name:    "manager",
			Image:   controllerImage,
			Command: []string{mclBinaryPath, fmt.Sprintf("--gcs-bucket=%s", b.bucket)},
		},
	}

	syncerComponent := controllerComponent{
		deploymentName:     syncerDeploymentName,
		componentLabel:     "syncer-controller",
		serviceAccountName: syncerServiceAccountName,
		clusterRoleName:    syncerClusterRoleName,
		policyRules: []rbacv1.PolicyRule{
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
		container: corev1.Container{
			Name:    "manager",
			Image:   controllerImage,
			Command: []string{syncerBinaryPath},
		},
	}

	objs := make([]client.Object, 0, 8)
	objs = append(objs, buildComponentObjects(project, mclComponent)...)
	objs = append(objs, buildComponentObjects(project, syncerComponent)...)
	return objs
}

func (b *manifestBuilder) renderClusterYAML(gkeInfo gkeClusterInfo, cc *operatorv1beta1.ConfigConnector, controllerImage string) (string, error) {
	objs := b.buildClusterObjects(gkeInfo.project, controllerImage)
	objs = append(objs, b.buildRenderedConfigConnector(cc, gkeInfo.identity))

	var out strings.Builder
	fmt.Fprintf(&out, "# Cluster context: %s\n", gkeInfo.contextName)
	for _, obj := range objs {
		raw, err := yaml.Marshal(obj)
		if err != nil {
			return "", fmt.Errorf("marshaling %s %q to YAML: %w", obj.GetObjectKind().GroupVersionKind().Kind, obj.GetName(), err)
		}

		out.WriteString("---\n")
		out.Write(raw)
	}

	return out.String(), nil
}

func (b *manifestBuilder) configureMultiClusterLease(spec *operatorv1beta1.ConfigConnectorSpec, clusterIdentity string) {
	if spec.Experiments == nil {
		spec.Experiments = &operatorv1beta1.CCExperiments{}
	}

	spec.Experiments.MultiClusterLease = &operatorv1beta1.MultiClusterLeaseSpec{
		LeaseName:                b.leaseName,
		Namespace:                b.leaseNamespace,
		ClusterCandidateIdentity: clusterIdentity,
		ResourceReplicationMode:  "Status",
	}
}

func (b *manifestBuilder) buildRenderedConfigConnector(cc *operatorv1beta1.ConfigConnector, clusterIdentity string) *operatorv1beta1.ConfigConnector {
	var spec operatorv1beta1.ConfigConnectorSpec
	if cc != nil {
		spec = *cc.Spec.DeepCopy()
	}

	b.configureMultiClusterLease(&spec, clusterIdentity)
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

func buildComponentObjects(project string, comp controllerComponent) []client.Object {
	gsaEmail := fmt.Sprintf("%s@%s.iam.gserviceaccount.com", comp.serviceAccountName, project)
	podLabels := map[string]string{"app": comp.componentLabel}

	sa := &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{APIVersion: "v1", Kind: "ServiceAccount"},
		ObjectMeta: metav1.ObjectMeta{
			Name:      comp.serviceAccountName,
			Namespace: cnrmSystemNamespace,
			Annotations: map[string]string{
				workloadIdentityGSAAnnot: gsaEmail,
			},
		},
	}

	role := &rbacv1.ClusterRole{
		TypeMeta:   metav1.TypeMeta{APIVersion: "rbac.authorization.k8s.io/v1", Kind: "ClusterRole"},
		ObjectMeta: metav1.ObjectMeta{Name: comp.clusterRoleName},
		Rules:      comp.policyRules,
	}

	binding := &rbacv1.ClusterRoleBinding{
		TypeMeta:   metav1.TypeMeta{APIVersion: "rbac.authorization.k8s.io/v1", Kind: "ClusterRoleBinding"},
		ObjectMeta: metav1.ObjectMeta{Name: comp.clusterRoleName},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     comp.clusterRoleName,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      comp.serviceAccountName,
				Namespace: cnrmSystemNamespace,
			},
			{
				Kind:     "Group",
				APIGroup: "rbac.authorization.k8s.io",
				Name:     cnrmSystemSAGroup,
			},
		},
	}

	deployment := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{APIVersion: "apps/v1", Kind: "Deployment"},
		ObjectMeta: metav1.ObjectMeta{
			Name:      comp.deploymentName,
			Namespace: cnrmSystemNamespace,
			Labels: map[string]string{
				componentLabelKey: comp.componentLabel,
			},
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{MatchLabels: podLabels},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{Labels: podLabels},
				Spec: corev1.PodSpec{
					ServiceAccountName: comp.serviceAccountName,
					Containers:         []corev1.Container{comp.container},
				},
			},
		},
	}

	return []client.Object{sa, role, binding, deployment}
}
