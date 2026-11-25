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

package kccmanager

import (
	"context"
	"fmt"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
	syncerv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	syncercontrollers "github.com/gke-labs/kube-etl/syncer/controllers"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	crconfig "sigs.k8s.io/controller-runtime/pkg/config"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
)

// A standby gets the leader's spec.resourceID, and keeps it.
//
// The leader's controller writes the ID that a service generated into spec.resourceID. The
// KRMSyncer that the manager configures on a standby has to copy it there: without it, a
// DCL-based controller that starts leading does not find the resource and creates another one.
// And the field has to stay. The syncer applies every rule that matches an object, in order,
// with one field manager, so a status-only rule that also matches the object removes
// spec.resourceID again, until the next rule puts it back.
//
// This runs the syncer's own controllers, with the KRMSyncer exactly as SyncerIntegration writes
// it, between two API servers. There is no admission webhook here, so the syncer's writes are
// admitted; pkg/webhook has the test for that half.
func TestStandbyKeepsLeaderResourceID(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const (
		namespace  = "team-a"
		name       = "a-group"
		resourceID = "1234567890123456789"
		syncs      = 5
	)
	gvk := schema.GroupVersionKind{Group: "monitoring.cnrm.cloud.google.com", Version: "v1beta1", Kind: "MonitoringGroup"}
	key := types.NamespacedName{Namespace: namespace, Name: name}

	kubeEtlModPath, err := getModulePath("github.com/gke-labs/kube-etl/syncer")
	if err != nil {
		t.Fatalf("error getting module path for kube-etl: %v", err)
	}
	groupCRD := filepath.Join(repo.GetRootOrTestFatal(t), "config/crds/resources",
		"apiextensions.k8s.io_v1_customresourcedefinition_monitoringgroups.monitoring.cnrm.cloud.google.com.yaml")

	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = syncerv1alpha1.AddToScheme(scheme)
	startCluster := func(crds ...string) (*rest.Config, client.WithWatch) {
		t.Helper()
		env := &envtest.Environment{CRDDirectoryPaths: crds, ErrorIfCRDPathMissing: true}
		cfg, err := env.Start()
		if err != nil {
			t.Fatalf("error starting envtest: %v", err)
		}
		t.Cleanup(func() { _ = env.Stop() })
		c, err := client.NewWithWatch(cfg, client.Options{Scheme: scheme})
		if err != nil {
			t.Fatalf("error creating client: %v", err)
		}
		return cfg, c
	}
	leaderCfg, leader := startCluster(groupCRD)
	standbyCfg, standby := startCluster(groupCRD, filepath.Join(kubeEtlModPath, "config/crd"))

	for _, c := range []client.Client{leader, standby} {
		for _, ns := range []string{namespace, "cnrm-system"} {
			if err := c.Create(ctx, &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: ns}}); err != nil {
				t.Fatalf("error creating namespace %s: %v", ns, err)
			}
		}
	}

	// The syncer, as it runs on the standby.
	skipNameValidation := true
	mgr, err := ctrl.NewManager(standbyCfg, ctrl.Options{
		Scheme:     scheme,
		Metrics:    metricsserver.Options{BindAddress: "0"},
		Controller: crconfig.Controller{SkipNameValidation: &skipNameValidation},
	})
	if err != nil {
		t.Fatalf("error creating the syncer's manager: %v", err)
	}
	syncer := &syncercontrollers.KRMSyncerReconciler{
		Client:               mgr.GetClient(),
		Scheme:               mgr.GetScheme(),
		Manager:              mgr,
		RemoteConfigProvider: fixedRemote{cfg: leaderCfg},
	}
	if err := syncer.SetupWithManager(mgr); err != nil {
		t.Fatalf("error setting up the syncer: %v", err)
	}
	go func() {
		if err := mgr.Start(ctx); err != nil && ctx.Err() == nil {
			t.Errorf("the syncer's manager stopped: %v", err)
		}
	}()

	// The same manifest on both clusters, as GitOps applies it: no spec.resourceID.
	newGroup := func() *unstructured.Unstructured {
		u := &unstructured.Unstructured{Object: map[string]interface{}{
			"metadata": map[string]interface{}{"name": name, "namespace": namespace},
			"spec": map[string]interface{}{
				"displayName": "a group",
				"filter":      "resource.metadata.name=has_substring(\"nothing\")",
				"projectRef":  map[string]interface{}{"external": "projects/a-project"},
			},
		}}
		u.SetGroupVersionKind(gvk)
		return u
	}
	if err := standby.Create(ctx, newGroup(), client.FieldOwner("a-user")); err != nil {
		t.Fatalf("error creating the standby's object: %v", err)
	}
	// The leader's controller created the group: it wrote the generated ID and the status.
	leaderObj := newGroup()
	if err := leader.Create(ctx, leaderObj, client.FieldOwner("a-user")); err != nil {
		t.Fatalf("error creating the leader's object: %v", err)
	}
	if err := unstructured.SetNestedField(leaderObj.Object, resourceID, "spec", "resourceID"); err != nil {
		t.Fatal(err)
	}
	if err := leader.Update(ctx, leaderObj, client.FieldOwner("cnrm-controller-manager")); err != nil {
		t.Fatalf("error setting spec.resourceID on the leader's object: %v", err)
	}
	setLeaderStatus := func(message string) {
		t.Helper()
		obj := &unstructured.Unstructured{}
		obj.SetGroupVersionKind(gvk)
		if err := leader.Get(ctx, key, obj); err != nil {
			t.Fatalf("error reading the leader's object: %v", err)
		}
		status := map[string]interface{}{
			"observedGeneration": obj.GetGeneration(),
			"conditions": []interface{}{map[string]interface{}{
				"type": "Ready", "status": "True", "reason": "UpToDate", "message": message,
				"lastTransitionTime": "2026-01-01T00:00:00Z",
			}},
		}
		if err := unstructured.SetNestedMap(obj.Object, status, "status"); err != nil {
			t.Fatal(err)
		}
		if err := leader.Status().Update(ctx, obj, client.FieldOwner("cnrm-controller-manager")); err != nil {
			t.Fatalf("error writing the leader's status: %v", err)
		}
	}
	setLeaderStatus("sync 0")

	// The KRMSyncer, as the manager of a standby writes it.
	si := &SyncerIntegration{
		client:          standby,
		apiReader:       standby,
		leaseNN:         types.NamespacedName{Namespace: "cnrm-system", Name: "kcc-multicluster-lease"},
		replicationMode: "Status",
		cluster:         testGKECluster("standby"),
	}
	if err := si.EnsurePullingFromLeader(ctx, testGKECluster("leader")); err != nil {
		t.Fatalf("error configuring the KRMSyncer: %v", err)
	}

	type state struct {
		generation int64
		resourceID string
		message    string
	}
	stateOf := func(u *unstructured.Unstructured) state {
		id, _, _ := unstructured.NestedString(u.Object, "spec", "resourceID")
		s := state{generation: u.GetGeneration(), resourceID: id}
		conditions, _, _ := unstructured.NestedSlice(u.Object, "status", "conditions")
		if len(conditions) > 0 {
			if c, ok := conditions[0].(map[string]interface{}); ok {
				s.message, _ = c["message"].(string)
			}
		}
		return s
	}
	standbyState := func() state {
		t.Helper()
		obj := &unstructured.Unstructured{}
		obj.SetGroupVersionKind(gvk)
		if err := standby.Get(ctx, key, obj); err != nil {
			t.Fatalf("error reading the standby's object: %v", err)
		}
		return stateOf(obj)
	}
	waitForStandby := func(what string, ok func(state) bool) {
		t.Helper()
		var last state
		if err := wait.PollUntilContextTimeout(ctx, 50*time.Millisecond, 30*time.Second, true, func(context.Context) (bool, error) {
			last = standbyState()
			return ok(last), nil
		}); err != nil {
			t.Fatalf("the standby's object did not get %s: %v; it has %+v", what, err, last)
		}
	}

	// 1. The standby gets the leader's spec.resourceID and status.
	waitForStandby("the leader's spec.resourceID and status", func(s state) bool {
		return s.resourceID == resourceID && s.message == "sync 0"
	})
	time.Sleep(time.Second) // let the first sync finish
	before := standbyState()

	// 2. It keeps it: while the leader's status changes, the standby's spec does not.
	list := &unstructured.UnstructuredList{}
	list.SetGroupVersionKind(gvk.GroupVersion().WithKind(gvk.Kind + "List"))
	w, err := standby.Watch(ctx, list, client.InNamespace(namespace), client.MatchingFields{"metadata.name": name})
	if err != nil {
		t.Fatalf("error watching the standby's object: %v", err)
	}
	var mu sync.Mutex
	var seen []state
	done := make(chan struct{})
	go func() {
		defer close(done)
		for ev := range w.ResultChan() {
			if u, ok := ev.Object.(*unstructured.Unstructured); ok {
				mu.Lock()
				seen = append(seen, stateOf(u))
				mu.Unlock()
			}
		}
	}()
	for i := 1; i <= syncs; i++ {
		message := fmt.Sprintf("sync %d", i)
		setLeaderStatus(message)
		waitForStandby("status "+message, func(s state) bool { return s.message == message })
		time.Sleep(300 * time.Millisecond) // let this sync finish
	}
	w.Stop()
	<-done

	after := standbyState()
	mu.Lock()
	defer mu.Unlock()
	withoutID := 0
	for _, s := range seen {
		if s.resourceID == "" {
			withoutID++
		}
	}
	if withoutID > 0 {
		t.Errorf("during %d syncs of the leader's status, %d of the %d states of the standby's object had no spec.resourceID: a rule that does not sync the field removed it, and another rule put it back", syncs, withoutID, len(seen))
	}
	if after.generation != before.generation {
		t.Errorf("during %d syncs of the leader's status, the standby's object went from generation %d to %d: its spec changed, although the leader's spec did not", syncs, before.generation, after.generation)
	}
	if after.resourceID != resourceID {
		t.Errorf("spec.resourceID on the standby is %q at the end, want %q", after.resourceID, resourceID)
	}
}

// fixedRemote makes the syncer reach one API server for every remote cluster.
type fixedRemote struct {
	cfg *rest.Config
}

func (p fixedRemote) Key(namespace string, remote *syncerv1alpha1.RemoteConfig) (string, error) {
	return (&syncercontrollers.GKEConfigProvider{}).Key(namespace, remote)
}

func (p fixedRemote) RESTConfig(context.Context, string, *syncerv1alpha1.RemoteConfig) (*rest.Config, error) {
	return rest.CopyConfig(p.cfg), nil
}
