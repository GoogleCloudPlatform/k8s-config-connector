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
	"context"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/kubecli"
	syncerv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/client/interceptor"
)

const (
	fleetLeaseName      = "kcc-multicluster-lease"
	fleetLeaseNamespace = "cnrm-system"

	managerComponentLabel = "cnrm.cloud.google.com/component"
	managerComponent      = "cnrm-controller-manager"
	scopedNamespaceLabel  = "cnrm.cloud.google.com/scoped-namespace"
)

// haFleet stands in for what runs on the clusters of an HA group while failover works on them:
//   - the StatefulSet controller and the kubelet, which stop and start the managers' pods;
//   - the MCL controllers, which give an expired lease to a cluster whose managers run, and
//     report the holder in every cluster's MultiClusterLease;
//   - Config Connector's managers, which renew the lease while they run, and keep their
//     KRMSyncers pointed at the leader.
//
// The fleet moves one step each time failover calls the API of one of its clusters, so the tests
// need no sleeps. It is a model of those components, written from their code.
type haFleet struct {
	t        *testing.T
	clusters []*fleetCluster

	// holder is the cluster that holds the lease.
	holder *fleetCluster
	// holders are the clusters that have held the lease, in order.
	holders []string
	// renewals counts how often the lease was renewed or taken.
	renewals int
	// leaderless counts the steps since the holder's managers stopped.
	leaderless int

	// takeOverSteps is how many steps the lease takes to expire once the holder's managers have
	// stopped. The first cluster that asks for it then takes it.
	takeOverSteps int
	// podStopSteps is how many steps a manager's pod takes to stop.
	podStopSteps int

	steps int
	// onStep, if set, runs at the start of every step.
	onStep func()
	// patches are failover's writes, in order.
	patches []fleetPatch
}

// fleetPatch is one write of failover, and what the fleet looked like when it was made.
type fleetPatch struct {
	cluster    string
	object     string
	replicas   int32
	annotation string
	// holder is the cluster that held the lease.
	holder string
	// running are the clusters that had a running manager.
	running string
}

func (p fleetPatch) String() string {
	if p.replicas == 0 {
		return fmt.Sprintf("stop %s %s (annotation %q) while %s leads and %s run", p.cluster, p.object, p.annotation, p.holder, p.running)
	}
	return fmt.Sprintf("start %s %s (replicas %d, annotation %q) while %s leads", p.cluster, p.object, p.replicas, p.annotation, p.holder)
}

// fleetCluster is one cluster of an haFleet.
type fleetCluster struct {
	fleet *haFleet
	name  string
	gke   operatorv1beta1.GKECluster

	// backend holds the cluster's objects. The fleet uses it directly.
	backend client.WithWatch
	// handle is how failover reaches the cluster. Each call through it moves the fleet one step.
	handle *clusterHandle

	// cannotLead keeps the cluster from taking the lease, as when its MCL controller can't
	// write the lease.
	cannotLead bool
	// stillAsks lets the cluster take an expired lease once while its managers are stopped. Its
	// MCL controller does that for 30 seconds after the managers are gone. Nothing renews a
	// lease that is taken that way.
	stillAsks bool
	// staleLease stops the updates of the cluster's MultiClusterLease, as when its MCL
	// controller has stopped.
	staleLease bool
	// noPods keeps the pods of its managers from starting, as when they can't be scheduled.
	noPods bool
	// stuckPods keeps the pods of its managers from stopping.
	stuckPods bool
	// unreadyPods keeps the pods of its managers from becoming ready.
	unreadyPods bool
	// abandonedSyncers stops the updates of its KRMSyncers, as when its managers can't write
	// them.
	abandonedSyncers bool
	// noSyncer stops the updates of the status of its KRMSyncers, as when its syncer has
	// stopped.
	noSyncer bool
	// syncFailures is what its syncer reports for a KRMSyncer that pulls and doesn't sync, by the
	// name of the KRMSyncer. Without an entry, the syncer reports that the KRMSyncer syncs.
	syncFailures map[string]fleetSyncFailure
	// rejectRead, if set, can fail a read of failover from the cluster. It gets the object or the
	// list that failover asks for.
	rejectRead func(ctx context.Context, obj runtime.Object) error
	// rejectPatch, if set, can fail a write of failover to the cluster. It gets the object and
	// the patch.
	rejectPatch func(obj client.Object, patch string) error

	// stopping counts, for each manager scaled to 0, the steps since its pod was told to stop.
	stopping map[types.NamespacedName]int
}

func (c *fleetCluster) identity() string {
	return c.gke.ResourceName()
}

// newHAFleet returns a fleet of clusters in cluster mode. Each has one manager, and the first
// one is the leader. The clusters' objects are in fake clients.
func newHAFleet(t *testing.T, names ...string) *haFleet {
	t.Helper()
	scheme, err := newHAScheme()
	if err != nil {
		t.Fatalf("newHAScheme() failed: %v", err)
	}
	fleet := newEmptyHAFleet(t)
	for _, name := range names {
		backend := fake.NewClientBuilder().
			WithScheme(scheme).
			WithStatusSubresource(&appsv1.StatefulSet{}, &corev1.Pod{}, &mclv1alpha1.MultiClusterLease{}, &syncerv1alpha1.KRMSyncer{}).
			Build()
		c := fleet.addCluster(name, backend)
		c.create(newConfigConnectorWithMCL(fleetLeaseName, fleetLeaseNamespace, name))
		c.create(&mclv1alpha1.MultiClusterLease{ObjectMeta: metav1.ObjectMeta{Namespace: fleetLeaseNamespace, Name: fleetLeaseName}})
		c.create(newManagerStatefulSet(fleetLeaseNamespace, managerComponent, ""))
	}
	fleet.settle()
	return fleet
}

func newEmptyHAFleet(t *testing.T) *haFleet {
	return &haFleet{t: t, takeOverSteps: 6, podStopSteps: 3}
}

// addCluster adds a cluster whose objects are behind backend. The first cluster is the leader.
func (f *haFleet) addCluster(name string, backend client.WithWatch) *fleetCluster {
	c := &fleetCluster{
		fleet:    f,
		name:     name,
		gke:      operatorv1beta1.GKECluster{Project: "test-project", Location: "us-central1", Name: name},
		backend:  backend,
		stopping: map[types.NamespacedName]int{},
	}
	// A request fails once its context is done, as a request to an API server does.
	stepping := interceptor.NewClient(backend, interceptor.Funcs{
		Get: func(ctx context.Context, backend client.WithWatch, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
			f.step()
			if err := c.failRead(ctx, obj); err != nil {
				return err
			}
			return backend.Get(ctx, key, obj, opts...)
		},
		List: func(ctx context.Context, backend client.WithWatch, list client.ObjectList, opts ...client.ListOption) error {
			f.step()
			if err := c.failRead(ctx, list); err != nil {
				return err
			}
			return backend.List(ctx, list, opts...)
		},
		Patch: func(ctx context.Context, backend client.WithWatch, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
			f.step()
			if err := ctx.Err(); err != nil {
				return err
			}
			if c.rejectPatch != nil {
				data, err := patch.Data(obj)
				if err != nil {
					return err
				}
				if err := c.rejectPatch(obj, string(data)); err != nil {
					return err
				}
			}
			if err := backend.Patch(ctx, obj, patch, opts...); err != nil {
				return err
			}
			f.recordPatch(c, obj)
			return nil
		},
	})
	c.handle = &clusterHandle{contextName: name, client: &kubecli.Client{Client: stepping}}
	f.clusters = append(f.clusters, c)
	if f.holder == nil {
		f.giveLease(c)
	}
	return c
}

func (c *fleetCluster) failRead(ctx context.Context, obj runtime.Object) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if c.rejectRead != nil {
		return c.rejectRead(ctx, obj)
	}
	return nil
}

func (f *haFleet) recordPatch(c *fleetCluster, obj client.Object) {
	p := fleetPatch{
		cluster:    c.name,
		object:     obj.GetNamespace() + "/" + obj.GetName(),
		annotation: obj.GetAnnotations()[failoverReplicasAnnotation],
		holder:     f.holder.name,
		running:    strings.Join(f.runningClusters(), ","),
	}
	if sts, ok := obj.(*appsv1.StatefulSet); ok {
		p.replicas = *sts.Spec.Replicas
	}
	f.patches = append(f.patches, p)
}

func (f *haFleet) cluster(name string) *fleetCluster {
	f.t.Helper()
	for _, c := range f.clusters {
		if c.name == name {
			return c
		}
	}
	f.t.Fatalf("the fleet has no cluster %q", name)
	return nil
}

func (f *haFleet) handles() []*clusterHandle {
	handles := make([]*clusterHandle, 0, len(f.clusters))
	for _, c := range f.clusters {
		handles = append(handles, c.handle)
	}
	return handles
}

func (f *haFleet) runningClusters() []string {
	var running []string
	for _, c := range f.clusters {
		if c.running() {
			running = append(running, c.name)
		}
	}
	return running
}

// settle runs the fleet until its controllers have caught up with the objects, as they would
// have by the time someone runs failover.
func (f *haFleet) settle() {
	for i := 0; i < f.takeOverSteps+f.podStopSteps+2; i++ {
		f.step()
	}
}

// step moves every component of the fleet forward once.
func (f *haFleet) step() {
	f.t.Helper()
	f.steps++
	if f.onStep != nil {
		f.onStep()
	}
	for _, c := range f.clusters {
		c.runPods()
	}
	f.renewOrTakeLease()
	for _, c := range f.clusters {
		c.reportLease()
		c.maintainSyncers()
		c.reportSync()
	}
}

// renewOrTakeLease is the lease itself: the holder keeps it for as long as one of its managers
// runs. After that it expires, and the first cluster that asks for it takes it: one with a
// running manager, or one whose MCL controller still asks for it.
func (f *haFleet) renewOrTakeLease() {
	if f.holder.running() {
		f.renewals++
		f.leaderless = 0
		return
	}
	f.leaderless++
	if f.leaderless < f.takeOverSteps {
		return
	}
	for _, c := range f.clusters {
		if c.cannotLead || !(c.running() || c.stillAsks) {
			continue
		}
		c.stillAsks = false
		f.giveLease(c)
		return
	}
}

// giveLease makes the cluster the holder of the lease.
func (f *haFleet) giveLease(c *fleetCluster) {
	f.holder = c
	f.holders = append(f.holders, c.name)
	f.renewals++
	f.leaderless = 0
}

func (f *haFleet) renewTime() string {
	return time.Date(2026, time.November, 18, 12, 0, 0, 0, time.UTC).Add(time.Duration(f.renewals) * time.Second).Format(time.RFC3339)
}

func (c *fleetCluster) create(obj client.Object) {
	c.fleet.t.Helper()
	if err := c.backend.Create(context.Background(), obj); err != nil {
		c.fleet.t.Fatalf("cluster %q: creating %T %s/%s: %v", c.name, obj, obj.GetNamespace(), obj.GetName(), err)
	}
}

func (c *fleetCluster) managers() []appsv1.StatefulSet {
	c.fleet.t.Helper()
	list := &appsv1.StatefulSetList{}
	if err := c.backend.List(context.Background(), list, client.MatchingLabels{managerComponentLabel: managerComponent}); err != nil {
		c.fleet.t.Fatalf("cluster %q: listing the managers: %v", c.name, err)
	}
	return list.Items
}

func (c *fleetCluster) manager(namespace, name string) *appsv1.StatefulSet {
	c.fleet.t.Helper()
	sts := &appsv1.StatefulSet{}
	if err := c.backend.Get(context.Background(), types.NamespacedName{Namespace: namespace, Name: name}, sts); err != nil {
		c.fleet.t.Fatalf("cluster %q: getting StatefulSet %s/%s: %v", c.name, namespace, name, err)
	}
	return sts
}

func managerPodKey(sts *appsv1.StatefulSet) types.NamespacedName {
	return types.NamespacedName{Namespace: sts.Namespace, Name: sts.Name + "-0"}
}

func (c *fleetCluster) hasPod(sts *appsv1.StatefulSet) bool {
	c.fleet.t.Helper()
	err := c.backend.Get(context.Background(), managerPodKey(sts), &corev1.Pod{})
	if err != nil && !apierrors.IsNotFound(err) {
		c.fleet.t.Fatalf("cluster %q: getting pod %s: %v", c.name, managerPodKey(sts), err)
	}
	return err == nil
}

// running reports whether a manager of the cluster runs. A manager that was scaled to 0 runs
// until its pod is gone.
func (c *fleetCluster) running() bool {
	for _, sts := range c.managers() {
		if c.hasPod(&sts) {
			return true
		}
	}
	return false
}

// runPods is the StatefulSet controller and the kubelet: it starts and stops the pod of each
// manager, and reports it in the status of the StatefulSet.
func (c *fleetCluster) runPods() {
	t := c.fleet.t
	t.Helper()
	ctx := context.Background()
	for _, sts := range c.managers() {
		key := client.ObjectKeyFromObject(&sts)
		running := c.hasPod(&sts)

		if *sts.Spec.Replicas == 0 {
			if running && !c.stuckPods {
				c.stopping[key]++
				if c.stopping[key] >= c.fleet.podStopSteps {
					pod := &corev1.Pod{ObjectMeta: metav1.ObjectMeta{Namespace: sts.Namespace, Name: managerPodKey(&sts).Name}}
					if err := c.backend.Delete(ctx, pod, client.GracePeriodSeconds(0)); err != nil {
						t.Fatalf("cluster %q: deleting pod %s: %v", c.name, managerPodKey(&sts), err)
					}
					running = false
				}
			}
		} else {
			delete(c.stopping, key)
			if !running && !c.noPods {
				c.create(&corev1.Pod{
					ObjectMeta: metav1.ObjectMeta{Namespace: sts.Namespace, Name: managerPodKey(&sts).Name, Labels: sts.Spec.Template.Labels},
					Spec:       sts.Spec.Template.Spec,
				})
				running = true
			}
		}

		status := appsv1.StatefulSetStatus{ObservedGeneration: sts.Generation}
		if running {
			status.Replicas = 1
			if !c.unreadyPods && *sts.Spec.Replicas > 0 {
				status.ReadyReplicas = 1
				status.AvailableReplicas = 1
			}
		}
		if !reflect.DeepEqual(sts.Status, status) {
			sts.Status = status
			if err := c.backend.Status().Update(ctx, &sts); err != nil {
				t.Fatalf("cluster %q: updating the status of StatefulSet %s: %v", c.name, key, err)
			}
		}
	}
}

// reportLease is the cluster's MCL controller: it writes the holder of the lease and its last
// renewal to the status of the cluster's MultiClusterLease.
func (c *fleetCluster) reportLease() {
	if c.staleLease {
		return
	}
	c.setLeaseStatus(c.fleet.holder.identity(), c.fleet.renewTime())
}

func (c *fleetCluster) setLeaseStatus(holder, renewTime string) {
	t := c.fleet.t
	t.Helper()
	ctx := context.Background()
	lease := &mclv1alpha1.MultiClusterLease{}
	if err := c.backend.Get(ctx, types.NamespacedName{Namespace: fleetLeaseNamespace, Name: fleetLeaseName}, lease); err != nil {
		if apierrors.IsNotFound(err) {
			return
		}
		t.Fatalf("cluster %q: getting the MultiClusterLease: %v", c.name, err)
	}
	status := lease.Status.DeepCopy()
	status.GlobalHolderIdentity, status.GlobalRenewTime = nil, nil
	if holder != "" {
		status.GlobalHolderIdentity = &holder
	}
	if renewTime != "" {
		status.GlobalRenewTime = &renewTime
	}
	if reflect.DeepEqual(&lease.Status, status) {
		return
	}
	lease.Status = *status
	if err := c.backend.Status().Update(ctx, lease); err != nil {
		t.Fatalf("cluster %q: updating the status of the MultiClusterLease: %v", c.name, err)
	}
}

// syncerNameOf returns the name of the KRMSyncer that a manager maintains: the lease's in cluster
// mode, with the namespace appended in namespaced mode (SyncerIntegration.getDesiredSyncers in
// pkg/controller/kccmanager).
func syncerNameOf(sts *appsv1.StatefulSet) string {
	if namespace := sts.Labels[scopedNamespaceLabel]; namespace != "" {
		return fleetLeaseName + "-" + namespace
	}
	return fleetLeaseName
}

// maintainSyncers is the cluster's Config Connector managers: each running manager keeps its
// KRMSyncer naming the leader, suspended on the leader itself and pulling everywhere else.
func (c *fleetCluster) maintainSyncers() {
	t := c.fleet.t
	t.Helper()
	if c.abandonedSyncers || !c.syncs() {
		return
	}
	ctx := context.Background()
	leader := c.fleet.holder
	for _, sts := range c.managers() {
		if !c.hasPod(&sts) {
			continue
		}
		name := syncerNameOf(&sts)
		spec := newSyncerSpec(leader.gke, c == leader)

		syncer := &syncerv1alpha1.KRMSyncer{}
		err := c.backend.Get(ctx, types.NamespacedName{Namespace: fleetLeaseNamespace, Name: name}, syncer)
		switch {
		case apierrors.IsNotFound(err):
			c.create(&syncerv1alpha1.KRMSyncer{ObjectMeta: metav1.ObjectMeta{Namespace: fleetLeaseNamespace, Name: name}, Spec: spec})
		case err != nil:
			t.Fatalf("cluster %q: getting KRMSyncer %s: %v", c.name, name, err)
		case !reflect.DeepEqual(syncer.Spec, spec):
			syncer.Spec = spec
			if err := c.backend.Update(ctx, syncer); err != nil {
				t.Fatalf("cluster %q: updating KRMSyncer %s: %v", c.name, name, err)
			}
		}
	}
}

// fleetSyncFailure is what a syncer reports for a KRMSyncer that doesn't sync: the reason and
// the message of its Ready condition, and the number of objects that fail.
type fleetSyncFailure struct {
	reason        string
	message       string
	failedObjects int32
}

// syncerReady returns the Ready condition that a working syncer reports for a KRMSyncer:
// Suspended if spec.suspend is set, as on the leader, and Syncing otherwise.
func syncerReady(syncer *syncerv1alpha1.KRMSyncer) metav1.Condition {
	ready := metav1.Condition{
		Type:               syncerv1alpha1.ConditionReady,
		Status:             metav1.ConditionTrue,
		Reason:             syncerv1alpha1.ReasonSyncing,
		Message:            "Syncing resources",
		ObservedGeneration: syncer.Generation,
	}
	if syncer.Spec.Suspend {
		ready.Reason, ready.Message = syncerv1alpha1.ReasonSuspended, "Sync is suspended by spec.suspend"
	}
	return ready
}

// reportSync is the cluster's syncer: it reports, in the status of each KRMSyncer, the spec
// that it has processed and whether it syncs.
func (c *fleetCluster) reportSync() {
	t := c.fleet.t
	t.Helper()
	if c.noSyncer {
		return
	}
	for _, syncer := range c.syncers() {
		ready := syncerReady(&syncer)
		var failedObjects int32
		if failure, failing := c.syncFailures[syncer.Name]; failing && !syncer.Spec.Suspend {
			ready.Status, ready.Reason, ready.Message = metav1.ConditionFalse, failure.reason, failure.message
			failedObjects = failure.failedObjects
		}

		status := syncer.Status.DeepCopy()
		status.ObservedGeneration = syncer.Generation
		status.FailedObjects = failedObjects
		meta.SetStatusCondition(&status.Conditions, ready)
		if reflect.DeepEqual(&syncer.Status, status) {
			continue
		}
		syncer.Status = *status
		if err := c.backend.Status().Update(context.Background(), &syncer); err != nil {
			t.Fatalf("cluster %q: updating the status of KRMSyncer %s: %v", c.name, syncer.Name, err)
		}
	}
}

// newSyncerSpec returns the spec of a KRMSyncer that names remote as its remote cluster.
func newSyncerSpec(remote operatorv1beta1.GKECluster, suspend bool) syncerv1alpha1.KRMSyncerSpec {
	return syncerv1alpha1.KRMSyncerSpec{
		Suspend: suspend,
		Mode:    syncerv1alpha1.ModePull,
		Remote: &syncerv1alpha1.RemoteConfig{GKECluster: &syncerv1alpha1.GKECluster{
			Project: remote.Project, Location: remote.Location, Name: remote.Name, Endpoint: syncerv1alpha1.GKEEndpointDefault,
		}},
		Rules: []syncerv1alpha1.ResourceRule{{Group: "*.cnrm.cloud.google.com", Version: "*", Kind: "*", SyncFields: []string{"status"}}},
	}
}

// syncs reports whether resource replication is on in the cluster's ConfigConnector.
func (c *fleetCluster) syncs() bool {
	c.fleet.t.Helper()
	cc := &operatorv1beta1.ConfigConnector{}
	if err := c.backend.Get(context.Background(), types.NamespacedName{Name: configConnectorResourceName}, cc); err != nil {
		c.fleet.t.Fatalf("cluster %q: getting the ConfigConnector: %v", c.name, err)
	}
	mcl := multiClusterLeaseSpec(cc)
	return mcl != nil && mcl.ResourceReplicationMode != "Disabled"
}

func (c *fleetCluster) syncers() []syncerv1alpha1.KRMSyncer {
	c.fleet.t.Helper()
	list := &syncerv1alpha1.KRMSyncerList{}
	if err := c.backend.List(context.Background(), list, client.InNamespace(fleetLeaseNamespace)); err != nil {
		c.fleet.t.Fatalf("cluster %q: listing the KRMSyncers: %v", c.name, err)
	}
	return list.Items
}

// newManagerStatefulSet returns a StatefulSet of Config Connector's manager with the labels that
// the operator gives it. scopedNamespace is the namespace that a per-namespace manager is for,
// and empty in cluster mode.
func newManagerStatefulSet(namespace, name, scopedNamespace string) *appsv1.StatefulSet {
	labels := map[string]string{
		managerComponentLabel:          managerComponent,
		"cnrm.cloud.google.com/system": "true",
	}
	if scopedNamespace != "" {
		labels[scopedNamespaceLabel] = scopedNamespace
	}
	replicas := int32(1)
	return &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:   namespace,
			Name:        name,
			Labels:      labels,
			Annotations: map[string]string{"cnrm.cloud.google.com/version": "1.162.0"},
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas:    &replicas,
			ServiceName: "cnrm-manager",
			Selector:    &metav1.LabelSelector{MatchLabels: labels},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{Labels: labels},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{Name: "manager", Image: "gcr.io/gke-release/cnrm/controller:1.162.0"}},
				},
			},
		},
	}
}

// useNamespacedMode replaces the cluster's manager with one manager per namespace, named and
// labeled as the operator does. managerNamespaces says where each one runs: cnrm-system, or a
// namespace of its own.
func (c *fleetCluster) useNamespacedMode(managerNamespaces map[string]string) {
	t := c.fleet.t
	t.Helper()
	ctx := context.Background()
	if err := c.backend.Delete(ctx, c.manager(fleetLeaseNamespace, managerComponent)); err != nil {
		t.Fatalf("cluster %q: deleting the manager: %v", c.name, err)
	}
	pod := &corev1.Pod{ObjectMeta: metav1.ObjectMeta{Namespace: fleetLeaseNamespace, Name: managerComponent + "-0"}}
	if err := c.backend.Delete(ctx, pod); err != nil {
		t.Fatalf("cluster %q: deleting the manager's pod: %v", c.name, err)
	}
	for _, syncer := range c.syncers() {
		if err := c.backend.Delete(ctx, &syncer); err != nil {
			t.Fatalf("cluster %q: deleting KRMSyncer %s: %v", c.name, syncer.Name, err)
		}
	}
	for namespace, managerNamespace := range managerNamespaces {
		c.create(newManagerStatefulSet(managerNamespace, managerComponent+"-"+namespace, namespace))
	}
}

// stopManagerAsFailoverDoes leaves a manager as an interrupted failover does: scaled to 0, with
// the replicas it had in the annotation.
func (c *fleetCluster) stopManagerAsFailoverDoes(namespace, name string) {
	c.fleet.t.Helper()
	sts := c.manager(namespace, name)
	sts.Annotations[failoverReplicasAnnotation] = "1"
	c.scaleManager(sts, 0)
}

func (c *fleetCluster) scaleManager(sts *appsv1.StatefulSet, replicas int32) {
	c.fleet.t.Helper()
	sts.Spec.Replicas = &replicas
	if err := c.backend.Update(context.Background(), sts); err != nil {
		c.fleet.t.Fatalf("cluster %q: scaling StatefulSet %s/%s: %v", c.name, sts.Namespace, sts.Name, err)
	}
}

// requireManager checks the replicas of a manager, and that failover's annotation is gone.
func (c *fleetCluster) requireManager(namespace, name string, wantReplicas int32) {
	t := c.fleet.t
	t.Helper()
	sts := c.manager(namespace, name)
	if got := *sts.Spec.Replicas; got != wantReplicas {
		t.Errorf("cluster %q: StatefulSet %s/%s has %d replicas, want %d", c.name, namespace, name, got, wantReplicas)
	}
	if value, found := sts.Annotations[failoverReplicasAnnotation]; found {
		t.Errorf("cluster %q: StatefulSet %s/%s still has the annotation %s=%q", c.name, namespace, name, failoverReplicasAnnotation, value)
	}
	if got, want := sts.Annotations["cnrm.cloud.google.com/version"], "1.162.0"; got != want {
		t.Errorf("cluster %q: StatefulSet %s/%s lost its other annotations: version is %q, want %q", c.name, namespace, name, got, want)
	}
}

// requireLeader checks that leader holds the lease, and that every running manager has caught up
// with it: the leader's KRMSyncers are suspended, and the others pull from the leader.
func (f *haFleet) requireLeader(leader *fleetCluster) {
	t := f.t
	t.Helper()
	if f.holder != leader {
		t.Fatalf("cluster %q holds the lease, want cluster %q", f.holder.name, leader.name)
	}
	for _, c := range f.clusters {
		if !c.syncs() {
			continue
		}
		syncers := map[string]syncerv1alpha1.KRMSyncer{}
		for _, syncer := range c.syncers() {
			syncers[syncer.Name] = syncer
		}
		for _, sts := range c.managers() {
			if !c.hasPod(&sts) {
				continue
			}
			syncer, found := syncers[syncerNameOf(&sts)]
			if !found {
				t.Errorf("cluster %q has no KRMSyncer %s", c.name, syncerNameOf(&sts))
				continue
			}
			if got := syncer.Spec.Remote.GKECluster.Name; got != leader.gke.Name {
				t.Errorf("cluster %q: KRMSyncer %s names cluster %q as its remote, want the leader %q", c.name, syncer.Name, got, leader.gke.Name)
			}
			if got, want := syncer.Spec.Suspend, c == leader; got != want {
				t.Errorf("cluster %q: KRMSyncer %s has suspend=%v, want %v", c.name, syncer.Name, got, want)
			}
		}
	}
}

// requirePatches checks failover's writes, and what the fleet looked like at each one.
func (f *haFleet) requirePatches(want ...string) {
	f.t.Helper()
	got := make([]string, 0, len(f.patches))
	for _, p := range f.patches {
		got = append(got, p.String())
	}
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		f.t.Errorf("failover's writes:\n  %s\nwant:\n  %s", strings.Join(got, "\n  "), strings.Join(want, "\n  "))
	}
}
