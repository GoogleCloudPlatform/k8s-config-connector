// Copyright 2025 Google LLC
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

package preview

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// GKNN is the canonical identity for a kube object; it is short for Group-Kind-Namespaced-Name
// (Version is an encoding artifact, and does not change the identity of the object)
type GKNN struct {
	Group     string
	Kind      string
	Namespace string
	Name      string
}

// Recorder holds the information from reconciling the objects
type Recorder struct {
	mutex   sync.Mutex
	objects map[GKNN]*objectInfo

	reconcileTrackerMutex sync.Mutex
	// Track if a resource has been reconciled.
	ReconciledResources map[GKNN]map[k8s.ReconcilerType]bool
	gknCount            map[GKN]int
	// Number of resources has not been reconciled.
	RemainResourcesCount int
}

// NewRecorder creates a new Recorder.
func NewRecorder() *Recorder {
	return &Recorder{
		objects:              make(map[GKNN]*objectInfo),
		ReconciledResources:  make(map[GKNN]map[k8s.ReconcilerType]bool),
		gknCount:             make(map[GKN]int),
		RemainResourcesCount: 0,
	}
}

// objectInfo holds the activity from reconciling the objects
type objectInfo struct {
	events []event
}

type event struct {
	// eventType is the type of event
	eventType EventType
	// diff is the diff that was recorded
	diff *structuredreporting.Diff
	// kubeAction is the kube action that was recorded
	kubeAction *kubeAction
	// gcpAction is the gcp action that was recorded
	gcpAction *gcpAction
	// object is the object that was reconciled
	object *unstructured.Unstructured
	// reconcilerType is the type of reconciler that start/ends
	reconcilerType k8s.ReconcilerType
}

type EventType string

const (
	EventTypeReconcileStart EventType = "reconcileStart"
	EventTypeReconcileEnd   EventType = "reconcileEnd"
	EventTypeDiff           EventType = "diff"
	EventTypeKubeAction     EventType = "kubeAction"
	EventTypeGCPAction      EventType = "gcpAction"
)

// kubeAction holds a kubernetes action that was recorded
type kubeAction struct {
	method string
	action Action
}

// Action indicates whether we blocked or ignored the requested action.
type Action string

const (
	// ActionIgnored indicates that the action was ignored.
	ActionIgnored Action = "ignored"
	// ActionBlocked indicates that the action was blocked.
	ActionBlocked Action = "blocked"
)

// gcpAction holds a GCP action that was recorded
type gcpAction struct {
	method string
	url    string
	body   string
	action Action
}

// NewStructuredReportingListener creates a new StructuredReportingListener.
func (r *Recorder) NewStructuredReportingListener() structuredreporting.Listener {
	return &structuredReportingListener{recorder: r}
}

// structuredReportingListener is a listener for structured reporting events.
type structuredReportingListener struct {
	recorder *Recorder
}

// OnError is called by the structured reporting subsystem when an error occurs.
func (l *structuredReportingListener) OnError(ctx context.Context, err error, args ...any) {
	blockedGCPError, ok := ExtractBlockedGCPError(err)
	if !ok {
		return
	}
	l.recorder.recordGCPAction(ctx, blockedGCPError, args, ActionBlocked)
}

// OnReconcileStart is called by the structured reporting subsystem when a reconcile starts.
func (l *structuredReportingListener) OnReconcileStart(ctx context.Context, u *unstructured.Unstructured, t k8s.ReconcilerType) {
	l.recorder.recordReconcileStart(ctx, u, t)
}

// OnReconcileEnd is called by the structured reporting subsystem when a reconcile ends.
func (l *structuredReportingListener) OnReconcileEnd(ctx context.Context, u *unstructured.Unstructured, result reconcile.Result, err error, t k8s.ReconcilerType) {
	l.recorder.recordReconcileEnd(ctx, u, result, err, t)
}

// OnDiff is called by the structured reporting subsystem when a diff occurs.
func (l *structuredReportingListener) OnDiff(ctx context.Context, diff *structuredreporting.Diff) {
	l.recorder.recordDiff(ctx, diff)
}

// RecordBlockedKubeMethod is called by the interceptingKubeClient when a write operation is blocked.
func (r *Recorder) RecordBlockedKubeMethod(ctx context.Context, method string, args ...any) {
	r.recordKubeAction(ctx, method, args, ActionBlocked)
}

// RecordIgnoredKubeMethod is called by the interceptingKubeClient when a read operation is ignored.
func (r *Recorder) RecordIgnoredKubeMethod(ctx context.Context, method string, args ...any) {
	r.recordKubeAction(ctx, method, args, ActionIgnored)
}

// recordDiff captures the diff into our recorder.
func (r *Recorder) recordDiff(ctx context.Context, diff *structuredreporting.Diff) {

	gknn := GKNN{}

	if u := diff.Object; u != nil {
		gknn = gknnFromUnstructured(u)
	}

	if done := r.GKNNDoneReconcile(gknn); done {
		return
	}

	info := r.getObjectInfo(gknn)
	info.events = append(info.events, event{
		eventType: EventTypeDiff,
		diff:      diff,
	})
}

// recordReconcileStart captures the reconcile start into our recorder.
func (r *Recorder) recordReconcileStart(ctx context.Context, u *unstructured.Unstructured, t k8s.ReconcilerType) {
	gknn := gknnFromUnstructured(u)
	if done := r.GKNNDoneReconcile(gknn); done {
		return
	}

	info := r.getObjectInfo(gknn)
	info.events = append(info.events, event{
		eventType:      EventTypeReconcileStart,
		object:         u.DeepCopy(),
		reconcilerType: t,
	})
}

// recordReconcileEnd captures the reconcile end into our recorder.
func (r *Recorder) recordReconcileEnd(ctx context.Context, u *unstructured.Unstructured, result reconcile.Result, err error, t k8s.ReconcilerType) {
	gknn := gknnFromUnstructured(u)

	if done := r.GKNNDoneReconcile(gknn); done {
		return
	}

	info := r.getObjectInfo(gknn)
	info.events = append(info.events, event{
		eventType:      EventTypeReconcileEnd,
		object:         u.DeepCopy(),
		reconcilerType: t,
	})
	r.reconcileTrackerMutex.Lock()
	r.ReconciledResources[gknn][t] = true
	r.reconcileTrackerMutex.Unlock()

	if r.GKNNDoneReconcile(gknn) {
		r.reconcileTrackerMutex.Lock()
		r.RemainResourcesCount--
		r.reconcileTrackerMutex.Unlock()
	}
}

func (r *Recorder) GKNNDoneReconcile(gknn GKNN) bool {
	r.reconcileTrackerMutex.Lock()
	defer r.reconcileTrackerMutex.Unlock()
	for _, done := range r.ReconciledResources[gknn] {
		if !done {
			return false
		}
	}
	return true
}

func gknnFromUnstructured(u *unstructured.Unstructured) GKNN {
	return GKNN{
		Group:     u.GroupVersionKind().Group,
		Kind:      u.GroupVersionKind().Kind,
		Namespace: u.GetNamespace(),
		Name:      u.GetName(),
	}
}

// recordKubeAction captures the kube action into our recorder.
func (r *Recorder) recordKubeAction(ctx context.Context, method string, args []any, action Action) {
	var gknn GKNN

	kubeAction := &kubeAction{
		method: method,
		action: action,
	}

	for _, arg := range args {
		switch arg := arg.(type) {
		case *unstructured.Unstructured:
			gvk := arg.GroupVersionKind()
			gknn = GKNN{
				Group:     gvk.Group,
				Kind:      gvk.Kind,
				Namespace: arg.GetNamespace(),
				Name:      arg.GetName(),
			}
			// We could capture the object here: kubeAction.object = arg.DeepCopy()

		case []client.SubResourceUpdateOption:
			// ignore

		case []client.UpdateOption:
			// ignore

		default:
			klog.Fatalf("unhandled arg type %T", arg)
		}
	}

	if done := r.GKNNDoneReconcile(gknn); done {
		return
	}

	info := r.getObjectInfo(gknn)
	info.events = append(info.events, event{
		eventType:  EventTypeKubeAction,
		kubeAction: kubeAction,
	})

}

// recordGCPAction captures the GCP action into our recorder.
func (r *Recorder) recordGCPAction(ctx context.Context, err *BlockedGCPError, args []any, action Action) {
	var gknn GKNN

	gcpAction := &gcpAction{
		method: err.Method,
		body:   err.Body,
		url:    err.URL,
		action: action,
	}

	for _, arg := range args {
		switch arg := arg.(type) {
		case *k8s.Resource:
			group := strings.Split(arg.APIVersion, "/")[0]
			gknn = GKNN{
				Group:     group,
				Kind:      arg.Kind,
				Namespace: arg.Namespace,
				Name:      arg.Name,
			}
		default:
			klog.Fatalf("unhandled arg type %T", arg)
		}
	}

	if done := r.GKNNDoneReconcile(gknn); done {
		return
	}

	info := r.getObjectInfo(gknn)
	info.events = append(info.events, event{
		eventType: EventTypeGCPAction,
		gcpAction: gcpAction,
	})
}

// getObjectInfo returns the objectInfo for the given GKNN.
// If there is no objectInfo, it creates a new one and returns it.
func (r *Recorder) getObjectInfo(gknn GKNN) *objectInfo {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	info := r.objects[gknn]
	if info == nil {
		info = &objectInfo{}
		r.objects[gknn] = info
	}
	return info
}

func (r *Recorder) DoneReconciling() bool {
	r.reconcileTrackerMutex.Lock()
	defer r.reconcileTrackerMutex.Unlock()
	return r.RemainResourcesCount == 0
}

// TODO: Implement concurrent worker by GVRs.
func (r *Recorder) PreloadGKNN(ctx context.Context, config *rest.Config) error {
	klog.Infof("Preloading the list of resources to reconcile")
	// Make a copy of config to increase QPS and burst.
	// This would not effect the config for the Manager.
	config = rest.CopyConfig(config)

	if config.QPS == 0 {
		config.QPS = 100
		config.Burst = 20
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("error creating Kubernetes clientset: %w", err)
	}

	discoveryClient := clientset.Discovery()
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("error creating dynamic client: %w", err)
	}
	apiResourceLists, err := discoveryClient.ServerPreferredResources()
	if err != nil {
		return fmt.Errorf("failed to get preferred resources: %w", err)
	}
	resourceControllerConfig := resourceconfig.LoadConfig()
	for _, apiResourceList := range apiResourceLists {
		if !strings.Contains(apiResourceList.GroupVersion, ".cnrm.cloud.google.com/") {
			continue
		}

		apiResourceListGroupVersion, err := schema.ParseGroupVersion(apiResourceList.GroupVersion)
		if err != nil {
			klog.Warningf("skipping unparseable groupVersion %q", apiResourceList.GroupVersion)
			continue
		}
		for _, apiResource := range apiResourceList.APIResources {
			if !apiResource.Namespaced {
				continue
			}
			if !contains(apiResource.Verbs, "list") {
				continue
			}
			gvr := schema.GroupVersionResource{
				Group:    apiResource.Group,
				Version:  apiResource.Version,
				Resource: apiResource.Name,
			}
			if gvr.Group == "" {
				gvr.Group = apiResourceListGroupVersion.Group
			}
			if gvr.Version == "" {
				gvr.Version = apiResourceListGroupVersion.Version
			}
			// Not tracking CC and CCC objects.
			if strings.HasSuffix(gvr.Group, "core.cnrm.cloud.google.com") {
				continue
			}
			resources, err := dynamicClient.Resource(gvr).List(ctx, metav1.ListOptions{})
			if err != nil {
				return fmt.Errorf("fetching gvr %s resources: %w", gvr, err)
			}
			gvk := schema.GroupVersionKind{
				Group:   gvr.Group,
				Version: gvr.Version,
				Kind:    apiResource.Kind,
			}
			config, err := resourceControllerConfig.GetControllersForGVK(gvk)
			if err != nil {
				klog.Warningf("error getting controller config found for GroupKind %v", gvk.GroupKind())
				continue
			}
			for _, resource := range resources.Items {
				r.ReconciledResources[GKNN{
					Group:     gvr.Group,
					Kind:      resource.GetKind(),
					Namespace: resource.GetNamespace(),
					Name:      resource.GetName(),
				}] = make(map[k8s.ReconcilerType]bool)
				gkn := GKN{
					Group:     gvr.Group,
					Kind:      resource.GetKind(),
					Namespace: resource.GetNamespace(),
				}
				r.gknCount[gkn] += 1
				for _, controllerType := range config.SupportedControllers {
					r.ReconciledResources[GKNN{
						Group:     gvr.Group,
						Kind:      resource.GetKind(),
						Namespace: resource.GetNamespace(),
						Name:      resource.GetName(),
					}][controllerType] = false
				}
			}
			r.RemainResourcesCount += len(resources.Items)
		}
	}
	klog.Infof("Got %d objects to reconcile", r.RemainResourcesCount)
	return nil
}

// contains checks if a slice contains a specific string.
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if strings.EqualFold(s, str) {
			return true
		}
	}
	return false
}
