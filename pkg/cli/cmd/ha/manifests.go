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
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

// manifestBuilder constructs the Kubernetes objects and ConfigConnector mutation for a single HA cluster.
type manifestBuilder struct {
	contextName       string
	candidateIdentity string
	projectID         string
	controllerImage   string
	mclImage          string
	syncerImage       string
	leaseName         string
	leaseNamespace    string
	bucket            string
}

func newManifestBuilder(contextName, candidateIdentity, projectID, controllerImage, leaseName, leaseNamespace, bucket string) *manifestBuilder {
	return &manifestBuilder{
		contextName:       contextName,
		candidateIdentity: candidateIdentity,
		projectID:         projectID,
		controllerImage:   controllerImage,
		leaseName:         leaseName,
		leaseNamespace:    leaseNamespace,
		bucket:            bucket,
	}
}

type controllerComponent struct {
	crd                *apiextensionsv1.CustomResourceDefinition
	deploymentName     string
	componentLabel     string
	serviceAccountName string
	gsaName            string
	clusterRoleName    string
	roleBindingName    string
	rules              []rbacv1.PolicyRule
	container          corev1.Container
}

const (
	mclCRDName               = "multiclusterleases.multicluster.core.cnrm.cloud.google.com"
	mclDeploymentName        = "multiclusterlease-controller-manager"
	mclServiceAccountName    = "kcc-mcl"
	syncerCRDName            = "krmsyncers.syncer.gkelabs.io"
	syncerDeploymentName     = "syncer-controller-manager"
	syncerServiceAccountName = "kcc-syncer"
)

func (b *manifestBuilder) buildClusterObjects() []client.Object {
	mclComponent := controllerComponent{
		crd: buildPermissiveV1Alpha1CRD(
			mclCRDName,
			"multicluster.core.cnrm.cloud.google.com",
			"MultiClusterLease",
			"MultiClusterLeaseList",
			"multiclusterleases",
			"multiclusterlease",
		),
		deploymentName:     mclDeploymentName,
		componentLabel:     "multiclusterlease-controller",
		serviceAccountName: mclServiceAccountName,
		gsaName:            mclServiceAccountName,
		clusterRoleName:    "cnrm-multiclusterlease-controller",
		roleBindingName:    "cnrm-multiclusterlease-controller",
		rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{"multicluster.core.cnrm.cloud.google.com"},
				Resources: []string{"multiclusterleases", "multiclusterleases/status"},
				Verbs:     []string{"get", "list", "watch", "create", "update", "patch", "delete"},
			},
			{
				APIGroups: []string{""},
				Resources: []string{"events"},
				Verbs:     []string{"create", "patch"},
			},
		},
		container: b.buildMCLContainer(),
	}

	syncerComponent := controllerComponent{
		crd: buildPermissiveV1Alpha1CRD(
			syncerCRDName,
			"syncer.gkelabs.io",
			"KRMSyncer",
			"KRMSyncerList",
			"krmsyncers",
			"krmsyncer",
		),
		deploymentName:     syncerDeploymentName,
		componentLabel:     "syncer-controller",
		serviceAccountName: syncerServiceAccountName,
		gsaName:            syncerServiceAccountName,
		clusterRoleName:    "cnrm-syncer-controller",
		roleBindingName:    "cnrm-syncer-controller",
		rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{"syncer.gkelabs.io"},
				Resources: []string{"krmsyncers", "krmsyncers/status"},
				Verbs:     []string{"get", "list", "watch", "create", "update", "patch", "delete"},
			},
			{
				APIGroups: []string{""},
				Resources: []string{"secrets", "namespaces", "events"},
				Verbs:     []string{"get", "list", "watch", "create", "update", "patch"},
			},
			{
				APIGroups: []string{"apiextensions.k8s.io"},
				Resources: []string{"customresourcedefinitions"},
				Verbs:     []string{"get", "list", "watch"},
			},
			{
				APIGroups: []string{"*.cnrm.cloud.google.com"},
				Resources: []string{"*"},
				Verbs:     []string{"get", "list", "watch", "create", "update", "patch", "delete"},
			},
		},
		container: b.buildSyncerContainer(),
	}

	objs := make([]client.Object, 0, 10)
	objs = append(objs, b.buildComponentObjects(mclComponent)...)
	objs = append(objs, b.buildComponentObjects(syncerComponent)...)
	return objs
}

func buildPermissiveV1Alpha1CRD(name, group, kind, listKind, plural, singular string) *apiextensionsv1.CustomResourceDefinition {
	return &apiextensionsv1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apiextensions.k8s.io/v1",
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: apiextensionsv1.CustomResourceDefinitionSpec{
			Group: group,
			Names: apiextensionsv1.CustomResourceDefinitionNames{
				Kind:     kind,
				ListKind: listKind,
				Plural:   plural,
				Singular: singular,
			},
			Scope: apiextensionsv1.NamespaceScoped,
			Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
				{
					Name:    "v1alpha1",
					Served:  true,
					Storage: true,
					Schema: &apiextensionsv1.CustomResourceValidation{
						OpenAPIV3Schema: &apiextensionsv1.JSONSchemaProps{
							Type:                   "object",
							XPreserveUnknownFields: ptr.To(true),
						},
					},
					Subresources: &apiextensionsv1.CustomResourceSubresources{
						Status: &apiextensionsv1.CustomResourceSubresourceStatus{},
					},
				},
			},
		},
	}
}

const mclBinaryPath = "/configconnector/mcl-manager"

func (b *manifestBuilder) buildMCLContainer() corev1.Container {
	if b.mclImage != "" {
		return corev1.Container{
			Name:  "manager",
			Image: b.mclImage,
			Args:  []string{"--gcs-bucket=" + b.bucket},
		}
	}

	return corev1.Container{
		Name:    "manager",
		Image:   b.controllerImage,
		Command: []string{mclBinaryPath, "--gcs-bucket=" + b.bucket},
	}
}

const syncerBinaryPath = "/configconnector/syncer-manager"

func (b *manifestBuilder) buildSyncerContainer() corev1.Container {
	if b.syncerImage != "" {
		return corev1.Container{
			Name:  "manager",
			Image: b.syncerImage,
			Args:  []string{"--mode=incluster"},
		}
	}

	return corev1.Container{
		Name:    "manager",
		Image:   b.controllerImage,
		Command: []string{syncerBinaryPath},
	}
}

const (
	manifestSystemNamespace       = "cnrm-system"
	manifestSystemSAGroup         = "system:serviceaccounts:cnrm-system"
	workloadIdentityGKEAnnotation = "iam.gke.io/gcp-service-account"
)

func (b *manifestBuilder) buildComponentObjects(comp controllerComponent) []client.Object {
	gsa := fmt.Sprintf("%s@%s.iam.gserviceaccount.com", comp.gsaName, b.projectID)
	labels := map[string]string{
		"cnrm.cloud.google.com/component": comp.componentLabel,
		"cnrm.cloud.google.com/system":    "true",
	}

	sa := &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ServiceAccount",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      comp.serviceAccountName,
			Namespace: manifestSystemNamespace,
			Annotations: map[string]string{
				workloadIdentityGKEAnnotation: gsa,
			},
			Labels: labels,
		},
	}

	role := &rbacv1.ClusterRole{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "ClusterRole",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   comp.clusterRoleName,
			Labels: labels,
		},
		Rules: comp.rules,
	}

	binding := &rbacv1.ClusterRoleBinding{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "ClusterRoleBinding",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   comp.roleBindingName,
			Labels: labels,
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     comp.clusterRoleName,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      comp.serviceAccountName,
				Namespace: manifestSystemNamespace,
			},
			{
				Kind:     "Group",
				Name:     manifestSystemSAGroup,
				APIGroup: "rbac.authorization.k8s.io",
			},
		},
	}

	replicas := int32(1)
	deployment := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      comp.deploymentName,
			Namespace: manifestSystemNamespace,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: comp.serviceAccountName,
					Containers:         []corev1.Container{comp.container},
				},
			},
		},
	}

	return []client.Object{comp.crd, sa, role, binding, deployment}
}

func (b *manifestBuilder) renderClusterYAML(cc *operatorv1beta1.ConfigConnector) (string, error) {
	objs := b.buildClusterObjects()

	ccCopy := cc.DeepCopy()
	ccCopy.TypeMeta = metav1.TypeMeta{
		APIVersion: operatorv1beta1.GroupVersion.String(),
		Kind:       "ConfigConnector",
	}
	ccCopy.ManagedFields = nil
	b.configureMultiClusterLease(&ccCopy.Spec)
	objs = append(objs, ccCopy)

	var sb strings.Builder
	fmt.Fprintf(&sb, "# Cluster: %s (candidate identity: %s)\n", b.contextName, b.candidateIdentity)
	for _, obj := range objs {
		rendered, err := serializeObjectYAML(obj)
		if err != nil {
			return "", fmt.Errorf("marshaling %T to YAML for cluster %q: %w", obj, b.contextName, err)
		}

		sb.WriteString("---\n")
		sb.WriteString(rendered)
	}

	return sb.String(), nil
}

func serializeObjectYAML(obj client.Object) (string, error) {
	raw, err := yaml.Marshal(obj)
	if err != nil {
		return "", err
	}

	var doc map[string]any
	if err := yaml.Unmarshal(raw, &doc); err != nil {
		return "", err
	}
	delete(doc, "status")
	if metaMap, ok := doc["metadata"].(map[string]any); ok {
		delete(metaMap, "creationTimestamp")
	}

	cleaned, err := yaml.Marshal(doc)
	if err != nil {
		return "", err
	}
	return string(cleaned), nil
}

const defaultReplicationMode = "Status"

func (b *manifestBuilder) configureMultiClusterLease(spec *operatorv1beta1.ConfigConnectorSpec) {
	if spec.Experiments == nil {
		spec.Experiments = &operatorv1beta1.CCExperiments{}
	}
	if spec.Experiments.MultiClusterLease == nil {
		spec.Experiments.MultiClusterLease = &operatorv1beta1.MultiClusterLeaseSpec{}
	}

	mcl := spec.Experiments.MultiClusterLease
	mcl.LeaseName = b.leaseName
	mcl.Namespace = b.leaseNamespace
	if mcl.ClusterCandidateIdentity == "" {
		mcl.ClusterCandidateIdentity = b.candidateIdentity
	}
	if mcl.ResourceReplicationMode == "" {
		mcl.ResourceReplicationMode = defaultReplicationMode
	}
}

func (b *manifestBuilder) setImageOverrides(mclImage, syncerImage string) {
	b.mclImage = mclImage
	b.syncerImage = syncerImage
}

func (b *manifestBuilder) leaseKey() string {
	return b.leaseNamespace + "/" + b.leaseName
}
