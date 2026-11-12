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
	operatork8s "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	syncercrd "github.com/gke-labs/kube-etl/syncer/config/crd"
	mclcrd "github.com/gke-labs/multicluster-leader-election/config/crd/bases"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

// manifestBuilder constructs and renders the Kubernetes manifests required for multi-cluster HA on a single cluster.
type manifestBuilder struct {
	contextName string
	// gkeCluster is the GKE cluster this Config Connector runs in; its resource name is the
	// cluster's identity in the multi-cluster election.
	gkeCluster operatorv1beta1.GKECluster
	project    string
	// peerProjects are the projects of the other clusters in the HA group.
	peerProjects []string
	// mode is the mode that Config Connector runs in on this cluster (ConfigConnector.GetMode).
	mode           string
	mclImage       string
	syncerImage    string
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
	// extraBindings bind the component's ServiceAccount, or its peers, to ClusterRoles that
	// Config Connector installs.
	extraBindings      []*rbacv1.ClusterRoleBinding
	podSecurityContext *corev1.PodSecurityContext
	container          corev1.Container
}

func newManifestBuilder(contextName string, gkeCluster operatorv1beta1.GKECluster, project, mode, mclImage, syncerImage, leaseName, leaseNamespace, bucket string) *manifestBuilder {
	return &manifestBuilder{
		contextName:    contextName,
		gkeCluster:     gkeCluster,
		project:        project,
		mode:           mode,
		mclImage:       mclImage,
		syncerImage:    syncerImage,
		leaseName:      leaseName,
		leaseNamespace: leaseNamespace,
		bucket:         bucket,
	}
}

// setPeerProjects sets the projects of the other clusters in the HA group. Their syncers read
// Config Connector resources from this cluster while it leads.
func (b *manifestBuilder) setPeerProjects(peerProjects []string) {
	b.peerProjects = peerProjects
}

const (
	// manifestManagedByLabelKey and manifestManagedByLabelValue label every object that setup
	// applies, so that they can be listed together.
	manifestManagedByLabelKey   = "app.kubernetes.io/managed-by"
	manifestManagedByLabelValue = "config-connector-ha-setup"
)

// buildClusterObjects returns the objects that setup applies to a cluster, in the order it
// applies them.
func (b *manifestBuilder) buildClusterObjects() []client.Object {
	var objs []client.Object
	objs = append(objs, buildCRDObjects()...)
	// The manager's access comes before the components. On a cluster that an earlier version of
	// setup configured, applying the component bindings removes a group subject that used to
	// give the manager this access.
	objs = append(objs, b.buildManagerObjects()...)
	objs = append(objs, b.buildComponentObjects(b.mclComponent())...)
	objs = append(objs, b.buildComponentObjects(b.syncerComponent())...)

	for _, obj := range objs {
		labels := obj.GetLabels()
		if labels == nil {
			labels = map[string]string{}
		}
		labels[manifestManagedByLabelKey] = manifestManagedByLabelValue
		obj.SetLabels(labels)
	}
	return objs
}

// buildCRDObjects returns the MultiClusterLease and KRMSyncer CRDs, as the forks in
// third_party/github.com/gke-labs define them.
func buildCRDObjects() []client.Object {
	return []client.Object{
		mustUnmarshalCRD(mclcrd.MultiClusterLeaseCRDYAML),
		mustUnmarshalCRD(syncercrd.KRMSyncerCRDYAML),
	}
}

func mustUnmarshalCRD(raw []byte) *apiextensionsv1.CustomResourceDefinition {
	crd := &apiextensionsv1.CustomResourceDefinition{}
	if err := yaml.Unmarshal(raw, crd); err != nil {
		panic(fmt.Sprintf("failed to unmarshal embedded CRD YAML: %v", err))
	}
	return crd
}

const (
	mclAPIGroup    = "multicluster.core.cnrm.cloud.google.com"
	syncerAPIGroup = "syncer.gkelabs.io"
	rbacAPIGroup   = "rbac.authorization.k8s.io"

	manifestSystemNamespace = "cnrm-system"
	manifestSystemSAGroup   = "system:serviceaccounts:cnrm-system"

	managerHARoleName                   = "cnrm-manager-ha"
	controllerManagerServiceAccountName = "cnrm-controller-manager"
)

// buildManagerObjects grants Config Connector's manager what multi-cluster leader election needs
// in the lease namespace:
//   - multiclusterleases: the leader election lock gets the lease, creates it, and patches it to
//     renew.
//   - krmsyncers: the manager gets each KRMSyncer, then creates or updates it so that it pulls
//     from the leader.
//
// Both use uncached clients, so neither needs list or watch.
func (b *manifestBuilder) buildManagerObjects() []client.Object {
	role := &rbacv1.Role{
		TypeMeta: metav1.TypeMeta{APIVersion: "rbac.authorization.k8s.io/v1", Kind: "Role"},
		ObjectMeta: metav1.ObjectMeta{
			Name:      managerHARoleName,
			Namespace: b.leaseNamespace,
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{mclAPIGroup},
				Resources: []string{"multiclusterleases"},
				Verbs:     []string{"get", "create", "patch"},
			},
			{
				APIGroups: []string{syncerAPIGroup},
				Resources: []string{"krmsyncers"},
				Verbs:     []string{"get", "create", "update"},
			},
		},
	}

	binding := &rbacv1.RoleBinding{
		TypeMeta: metav1.TypeMeta{APIVersion: "rbac.authorization.k8s.io/v1", Kind: "RoleBinding"},
		ObjectMeta: metav1.ObjectMeta{
			Name:      managerHARoleName,
			Namespace: b.leaseNamespace,
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: rbacAPIGroup,
			Kind:     "Role",
			Name:     managerHARoleName,
		},
		Subjects: b.managerSubjects(),
	}

	return []client.Object{role, binding}
}

// managerSubjects returns the identities that Config Connector's manager runs as on this cluster.
func (b *manifestBuilder) managerSubjects() []rbacv1.Subject {
	if b.mode == operatork8s.NamespacedMode {
		// Namespaced mode runs one manager per namespace, each as its own ServiceAccount in
		// cnrm-system (cnrm-controller-manager-<namespace>). Namespaces can be enabled after
		// setup runs, so the subject is every ServiceAccount in cnrm-system rather than a list
		// of names. That includes the ServiceAccounts in cnrm-system that aren't managers, such
		// as the MCL controller's and the syncer's.
		//
		// TODO: Managers that run in their own namespace (ConfigConnectorContext
		// spec.managerNamespace) are not in this group, so they get no access.
		return []rbacv1.Subject{
			{
				APIGroup: rbacAPIGroup,
				Kind:     "Group",
				Name:     manifestSystemSAGroup,
			},
		}
	}

	return []rbacv1.Subject{
		{
			Kind:      "ServiceAccount",
			Name:      controllerManagerServiceAccountName,
			Namespace: manifestSystemNamespace,
		},
	}
}

const (
	mclServiceAccountName = "kcc-mcl"
	mclClusterRoleName    = "cnrm-mcl-controller"
	mclDeploymentName     = "multiclusterlease-controller-manager"
)

// mclComponent describes the MultiClusterLease controller, which holds this cluster's lease in
// the GCS bucket.
func (b *manifestBuilder) mclComponent() controllerComponent {
	return controllerComponent{
		deploymentName:     mclDeploymentName,
		componentLabel:     "mcl-controller",
		serviceAccountName: mclServiceAccountName,
		clusterRoleName:    mclClusterRoleName,
		// These are the rules in the fork's config/rbac/role.yaml. The controller watches leases
		// in every namespace, so they are in a ClusterRole.
		policyRules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{mclAPIGroup},
				Resources: []string{"multiclusterleases"},
				Verbs:     []string{"get", "list", "watch", "create", "update", "patch", "delete"},
			},
			{
				APIGroups: []string{mclAPIGroup},
				Resources: []string{"multiclusterleases/status"},
				Verbs:     []string{"get", "update", "patch"},
			},
			{
				APIGroups: []string{mclAPIGroup},
				Resources: []string{"multiclusterleases/finalizers"},
				Verbs:     []string{"update"},
			},
			{
				APIGroups: []string{""},
				Resources: []string{"events"},
				Verbs:     []string{"create", "patch"},
			},
		},
		// There is no runAsNonRoot here, unlike for the syncer: the fork's own image names its
		// user (nonroot), and the kubelet can only verify a numeric user.
		container: corev1.Container{
			Name:  "manager",
			Image: b.mclImage,
			Args:  []string{fmt.Sprintf("--gcs-bucket=%s", b.bucket)},
			// The requests and limits are the fork's (config/manager/manager.yaml).
			Resources:       resourceRequirements("300m", "256Mi", "500m", "512Mi"),
			SecurityContext: unprivilegedSecurityContext(),
		},
	}
}

const (
	syncerServiceAccountName = "kcc-syncer"
	syncerClusterRoleName    = "cnrm-syncer-controller"
	syncerDeploymentName     = "syncer-controller-manager"

	syncerAdminBindingName = "cnrm-syncer-admin"
	syncerPeersBindingName = "cnrm-syncer-peers"
	cnrmAdminClusterRole   = "cnrm-admin"
	cnrmViewerClusterRole  = "cnrm-viewer"

	syncerDefaultSyncPeriodArg = "--sync-period=1h"
)

// syncerComponent describes the KRM syncer, which copies Config Connector resources from the
// leader to this cluster while this cluster is a standby.
func (b *manifestBuilder) syncerComponent() controllerComponent {
	runAsNonRoot := true
	return controllerComponent{
		deploymentName:     syncerDeploymentName,
		componentLabel:     "syncer-controller",
		serviceAccountName: syncerServiceAccountName,
		clusterRoleName:    syncerClusterRoleName,
		policyRules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{syncerAPIGroup},
				Resources: []string{"krmsyncers"},
				Verbs:     []string{"get", "list", "watch"},
			},
			{
				APIGroups: []string{syncerAPIGroup},
				Resources: []string{"krmsyncers/status"},
				Verbs:     []string{"get", "update", "patch"},
			},
			{
				// The syncer creates the namespace of a resource that it copies when the
				// namespace is missing.
				APIGroups: []string{""},
				Resources: []string{"namespaces"},
				Verbs:     []string{"get", "list", "watch", "create"},
			},
		},
		extraBindings: []*rbacv1.ClusterRoleBinding{
			b.buildSyncerAdminBinding(),
			b.buildSyncerPeersBinding(),
		},
		// The security context, requests and limits are the fork's (config/manager/manager.yaml).
		podSecurityContext: &corev1.PodSecurityContext{RunAsNonRoot: &runAsNonRoot},
		container: corev1.Container{
			Name:            "manager",
			Image:           b.syncerImage,
			Args:            []string{syncerDefaultSyncPeriodArg},
			Resources:       resourceRequirements("10m", "64Mi", "500m", "128Mi"),
			SecurityContext: unprivilegedSecurityContext(),
		},
	}
}

// buildSyncerAdminBinding lets this cluster's syncer write the Config Connector resources that it
// copies from the leader.
//
// It binds Config Connector's own cnrm-admin ClusterRole, because RBAC can't match API groups by
// suffix. That role lists every Config Connector resource group and nothing else: it covers
// neither Secrets nor the objects that configure Config Connector itself (ConfigConnector,
// ConfigConnectorContext, MultiClusterLease).
func (b *manifestBuilder) buildSyncerAdminBinding() *rbacv1.ClusterRoleBinding {
	return &rbacv1.ClusterRoleBinding{
		TypeMeta:   metav1.TypeMeta{APIVersion: "rbac.authorization.k8s.io/v1", Kind: "ClusterRoleBinding"},
		ObjectMeta: metav1.ObjectMeta{Name: syncerAdminBindingName},
		RoleRef: rbacv1.RoleRef{
			APIGroup: rbacAPIGroup,
			Kind:     "ClusterRole",
			Name:     cnrmAdminClusterRole,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      syncerServiceAccountName,
				Namespace: manifestSystemNamespace,
			},
		},
	}
}

// buildSyncerPeersBinding lets the other clusters' syncers read Config Connector resources on
// this cluster, which they do while this cluster leads.
//
// A syncer authenticates to another cluster as its Google service account, so the subjects are
// users named by email. This binding is their only access to objects in this cluster: the IAM
// role that setup grants them (syncerClusterAccessRole) only lets them find and connect to it.
//
// The grant is cluster-wide in namespaced mode too, because the syncer watches each kind across
// all namespaces.
func (b *manifestBuilder) buildSyncerPeersBinding() *rbacv1.ClusterRoleBinding {
	subjects := make([]rbacv1.Subject, 0, len(b.peerProjects))
	for _, peerProject := range b.peerProjects {
		subjects = append(subjects, rbacv1.Subject{
			APIGroup: rbacAPIGroup,
			Kind:     "User",
			Name:     gsaEmail(syncerServiceAccountName, peerProject),
		})
	}

	return &rbacv1.ClusterRoleBinding{
		TypeMeta:   metav1.TypeMeta{APIVersion: "rbac.authorization.k8s.io/v1", Kind: "ClusterRoleBinding"},
		ObjectMeta: metav1.ObjectMeta{Name: syncerPeersBindingName},
		RoleRef: rbacv1.RoleRef{
			APIGroup: rbacAPIGroup,
			Kind:     "ClusterRole",
			Name:     cnrmViewerClusterRole,
		},
		Subjects: subjects,
	}
}

func resourceRequirements(cpuRequest, memoryRequest, cpuLimit, memoryLimit string) corev1.ResourceRequirements {
	return corev1.ResourceRequirements{
		Requests: corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse(cpuRequest),
			corev1.ResourceMemory: resource.MustParse(memoryRequest),
		},
		Limits: corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse(cpuLimit),
			corev1.ResourceMemory: resource.MustParse(memoryLimit),
		},
	}
}

func unprivilegedSecurityContext() *corev1.SecurityContext {
	allowPrivilegeEscalation := false
	return &corev1.SecurityContext{
		AllowPrivilegeEscalation: &allowPrivilegeEscalation,
		Capabilities:             &corev1.Capabilities{Drop: []corev1.Capability{"ALL"}},
	}
}

const (
	workloadIdentityGSAAnnot  = "iam.gke.io/gcp-service-account"
	manifestComponentLabelKey = "cnrm.cloud.google.com/component"

	// componentMinReadySeconds is how long a component's pod must stay ready before its
	// Deployment counts it as available. The MCL controller exits right after it starts when it
	// can't read the lease bucket, and setup must not take such a pod for a working controller
	// (see clusterConfigurer.waitForMCLController).
	componentMinReadySeconds = 10
	// componentTerminationGracePeriodSeconds is the forks' value. It bounds how long a rollout
	// waits for the old pod to stop.
	componentTerminationGracePeriodSeconds = 10
)

func (b *manifestBuilder) buildComponentObjects(comp controllerComponent) []client.Object {
	podLabels := map[string]string{"app": comp.componentLabel}
	replicas := int32(1)
	terminationGracePeriodSeconds := int64(componentTerminationGracePeriodSeconds)

	sa := &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{APIVersion: "v1", Kind: "ServiceAccount"},
		ObjectMeta: metav1.ObjectMeta{
			Name:      comp.serviceAccountName,
			Namespace: manifestSystemNamespace,
			Annotations: map[string]string{
				workloadIdentityGSAAnnot: gsaEmail(comp.serviceAccountName, b.project),
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
			APIGroup: rbacAPIGroup,
			Kind:     "ClusterRole",
			Name:     comp.clusterRoleName,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      comp.serviceAccountName,
				Namespace: manifestSystemNamespace,
			},
		},
	}

	deployment := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{APIVersion: "apps/v1", Kind: "Deployment"},
		ObjectMeta: metav1.ObjectMeta{
			Name:      comp.deploymentName,
			Namespace: manifestSystemNamespace,
			Labels: map[string]string{
				manifestComponentLabelKey: comp.componentLabel,
			},
		},
		Spec: appsv1.DeploymentSpec{
			// Neither component runs its own leader election, so two of its pods must never run
			// at once: there is one replica, and Recreate stops the old pod before it starts the
			// new one.
			//
			// While the MCL controller's pod is down, nothing renews this cluster's lease in the
			// bucket. A rollout that takes longer than the lease duration (15 seconds), for
			// example to pull a new image, can therefore move leadership to another cluster.
			Replicas: &replicas,
			Strategy: appsv1.DeploymentStrategy{
				Type: appsv1.RecreateDeploymentStrategyType,
			},
			MinReadySeconds: componentMinReadySeconds,
			Selector:        &metav1.LabelSelector{MatchLabels: podLabels},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{Labels: podLabels},
				Spec: corev1.PodSpec{
					ServiceAccountName:            comp.serviceAccountName,
					SecurityContext:               comp.podSecurityContext,
					TerminationGracePeriodSeconds: &terminationGracePeriodSeconds,
					Containers:                    []corev1.Container{comp.container},
				},
			},
		},
	}

	objs := make([]client.Object, 0, 4+len(comp.extraBindings))
	objs = append(objs, sa, role, binding)
	for _, extra := range comp.extraBindings {
		objs = append(objs, extra)
	}
	objs = append(objs, deployment)
	return objs
}

func (b *manifestBuilder) renderClusterYAML(cc *operatorv1beta1.ConfigConnector) (string, error) {
	objs := b.buildClusterObjects()
	objs = append(objs, b.buildRenderedConfigConnector(cc))

	var out strings.Builder
	fmt.Fprintf(&out, "# Cluster context: %s\n", b.contextName)
	for _, obj := range objs {
		raw, err := marshalManifest(obj)
		if err != nil {
			return "", fmt.Errorf("marshaling %s %q to YAML: %w", obj.GetObjectKind().GroupVersionKind().Kind, obj.GetName(), err)
		}

		out.WriteString("---\n")
		out.Write(raw)
	}

	return out.String(), nil
}

// marshalManifest renders obj as YAML without its status, which only the cluster sets.
func marshalManifest(obj client.Object) ([]byte, error) {
	fields, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	delete(fields, "status")
	return yaml.Marshal(fields)
}

const manifestConfigConnectorName = "configconnector.core.cnrm.cloud.google.com"

func (b *manifestBuilder) buildRenderedConfigConnector(cc *operatorv1beta1.ConfigConnector) *operatorv1beta1.ConfigConnector {
	var spec operatorv1beta1.ConfigConnectorSpec
	if cc != nil {
		spec = *cc.Spec.DeepCopy()
	}

	b.configureMultiClusterLease(&spec)
	return &operatorv1beta1.ConfigConnector{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "core.cnrm.cloud.google.com/v1beta1",
			Kind:       "ConfigConnector",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: manifestConfigConnectorName,
		},
		Spec: spec,
	}
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
	// clusterCandidateIdentity is deprecated and rejected by the ConfigConnector CRD; the
	// cluster's identity is now derived from gkeCluster.
	mcl.ClusterCandidateIdentity = ""
	if mcl.GKECluster == nil {
		gkeCluster := b.gkeCluster
		mcl.GKECluster = &gkeCluster
	}
	if mcl.ResourceReplicationMode == "" {
		mcl.ResourceReplicationMode = defaultReplicationMode
	}
}

func (b *manifestBuilder) leaseKey() string {
	return b.leaseNamespace + "/" + b.leaseName
}
