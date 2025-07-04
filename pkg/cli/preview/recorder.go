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
	"strings"
	"sync"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
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
}

// NewRecorder creates a new Recorder.
func NewRecorder() *Recorder {
	return &Recorder{
		objects: make(map[GKNN]*objectInfo),
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
func (l *structuredReportingListener) OnReconcileStart(ctx context.Context, u *unstructured.Unstructured) {
	l.recorder.recordReconcileStart(ctx, u)
}

// OnReconcileEnd is called by the structured reporting subsystem when a reconcile ends.
func (l *structuredReportingListener) OnReconcileEnd(ctx context.Context, u *unstructured.Unstructured, result reconcile.Result, err error) {
	l.recorder.recordReconcileEnd(ctx, u, result, err)
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
	log := klog.FromContext(ctx)

	gknn := GKNN{}

	if u := diff.Object; u != nil {
		gknn = gknnFromUnstructured(u)
	}

	log.Info("recordDiffs", "gknn", gknn)

	info := r.getObjectInfo(gknn)
	info.events = append(info.events, event{
		eventType: EventTypeDiff,
		diff:      diff,
	})
}

// recordReconcileStart captures the reconcile start into our recorder.
func (r *Recorder) recordReconcileStart(ctx context.Context, u *unstructured.Unstructured) {
	gknn := gknnFromUnstructured(u)

	info := r.getObjectInfo(gknn)
	info.events = append(info.events, event{
		eventType: EventTypeReconcileStart,
		object:    u.DeepCopy(),
	})
}

// recordReconcileEnd captures the reconcile end into our recorder.
func (r *Recorder) recordReconcileEnd(ctx context.Context, u *unstructured.Unstructured, result reconcile.Result, err error) {
	gknn := gknnFromUnstructured(u)

	info := r.getObjectInfo(gknn)
	info.events = append(info.events, event{
		eventType: EventTypeReconcileEnd,
		object:    u.DeepCopy(),
	})
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
	klog.Infof("recordKubeAction %v %v %v", method, args, action)
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
