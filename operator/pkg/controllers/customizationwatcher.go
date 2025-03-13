// Copyright 2023 Google LLC
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

package controllers

import (
	"context"
	"fmt"
	"sync"

	customizev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/customize/v1beta1"
	corekcck8s "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/dynamic"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

var (
	// CustomizationCRsToWatch contains all the customization CRs to watch
	CustomizationCRsToWatch = []schema.GroupVersionResource{
		corekcck8s.ToGVR(customizev1beta1.ControllerResourceGroupVersionKind),
		corekcck8s.ToGVR(customizev1beta1.ValidatingWebhookConfigurationCustomizationGroupVersionKind),
		corekcck8s.ToGVR(customizev1beta1.MutatingWebhookConfigurationCustomizationGroupVersionKind),
		corekcck8s.ToGVR(customizev1beta1.ControllerReconcilerGroupVersionKind),
	}
	// NamespacedCustomizationCRsToWatch contains all the namespaced customization CRs to watch
	NamespacedCustomizationCRsToWatch = []schema.GroupVersionResource{
		corekcck8s.ToGVR(customizev1beta1.NamespacedControllerResourceGroupVersionKind),
		corekcck8s.ToGVR(customizev1beta1.NamespacedControllerReconcilerGroupVersionKind),
	}
)

// CustomizationWatcher setup watches on 'triggerGVRs'. It is used by the CC / CCC operator.
//
// For ConfigConnector operator, CustomizationWatcher setups a cluster-scoped watch on the customization CRs. Any
// changes to the customization CRs raises a watch event on the ConfigConnector object, which then triggers a reconciliation.
//
// For ConfigConnectorContext operator, CustomizationWatcher setups a cluster-scoped watch, despite the fact that the
// CRs being watched are namespaced. This is to keep the number of watches low when there are large number of namespaces
// managed by Config Connector. Any changes to the namespaced customization CRs raises a watch event on the ConfigConnectorContext
// object in the same namespace.
//
// The raised watch events are sent to "events" channel, which is watched by CC / CCC operator.
type CustomizationWatcher struct {
	// triggerGVRs contains all the GVRs to watch. An event on triggerGVRs will raise an event on the target.
	triggerGVRs []schema.GroupVersionResource
	// watchRegistered tracks the GVRs we are currently watching, avoid duplicate watches.
	watchRegistered map[string]struct{}
	// watchRegisteredMu protects access to watchRegistered map.
	watchRegisteredMu sync.RWMutex
	// lastRV caches the last reported resource version to filter out duplicated watch events.
	lastRV map[string]string
	// events is the channel that an event is raised and send to when CustomizationWatcher
	// receives a watch event on the triggerGVRs it is watching.
	events        chan event.GenericEvent
	dynamicClient dynamic.Interface
	log           logr.Logger
}

type CustomizationWatcherOptions struct {
	TriggerGVRs []schema.GroupVersionResource
	Log         logr.Logger
}

func NewWithDynamicClient(dc dynamic.Interface, opts CustomizationWatcherOptions) *CustomizationWatcher {
	return &CustomizationWatcher{
		dynamicClient:   dc,
		triggerGVRs:     opts.TriggerGVRs,
		log:             opts.Log.WithName("customization-watcher"),
		watchRegistered: make(map[string]struct{}),
		lastRV:          make(map[string]string),
		events:          make(chan event.GenericEvent),
	}
}

// Events returns a channel with events raised on target.
func (w *CustomizationWatcher) Events() chan event.GenericEvent {
	return w.events
}

// EnsureWatchStarted starts watches on triggerGVRs if not already done so.
func (w *CustomizationWatcher) EnsureWatchStarted(ctx context.Context, targetNN types.NamespacedName) error {
	for _, gvr := range w.triggerGVRs {
		go w.startWatch(ctx, gvr, targetNN)
	}
	return nil
}

// startWatch starts a watch for changes to "triggerGVR", raises events on target object which is in
// the same namespace as the watch events received on "triggerGVR".
func (w *CustomizationWatcher) startWatch(ctx context.Context, triggerGVR schema.GroupVersionResource, targetNN types.NamespacedName) {
	// if watch for the triggerGVR is already started / registered, skip; otherwise register the watch for the triggerGVR.
	w.watchRegisteredMu.Lock()
	if _, found := w.watchRegistered[triggerGVR.String()]; found {
		w.watchRegisteredMu.Unlock()
		return
	}
	w.watchRegistered[triggerGVR.String()] = struct{}{}
	w.watchRegisteredMu.Unlock()

	// make sure we de-register the watch for the triggerGVR when the watch is stopped.
	defer func() {
		w.watchRegisteredMu.Lock()
		delete(w.watchRegistered, triggerGVR.String())
		w.watchRegisteredMu.Unlock()
	}()

	log := w.log.WithValues("trigger GVR", triggerGVR.String(), "target NamespacedName", targetNN)
	opts := metav1.ListOptions{AllowWatchBookmarks: true}
	triggerEvents, err := w.dynamicClient.Resource(triggerGVR).Watch(ctx, opts)
	if err != nil {
		log.Error(err, "failed to start watch")
		return
	}
	log.Info("watch started")
	defer func() {
		triggerEvents.Stop()
		log.Info("watch stopped")
	}()

	for {
		select {
		case <-ctx.Done():
			log.Info("watch context cancelled")
			return
		case triggerEvent, ok := <-triggerEvents.ResultChan():
			if !ok {
				log.Info("watch channel closed")
				return
			}
			switch triggerEvent.Type {
			case watch.Bookmark:
				continue
			case watch.Error:
				log.Error(err, "unexpected event type from watch")
				return
			}

			u, ok := triggerEvent.Object.(*unstructured.Unstructured)
			if !ok {
				log.Error(fmt.Errorf("expect watch event object to be *unstructure.Unstructured"), "unexpected event object from watch", "event object", triggerEvent.Object)
				continue
			}
			lastRVKey := fmt.Sprintf("gvr=%s;namespace=%s;name=%s", triggerGVR.String(), u.GetNamespace(), u.GetName()) // a unique key given GVR, namespace and name
			switch triggerEvent.Type {
			case watch.Deleted:
				delete(w.lastRV, lastRVKey)
			case watch.Added, watch.Modified:
				rv := u.GetResourceVersion()
				if previousRV, found := w.lastRV[lastRVKey]; found && previousRV == rv {
					// Filter out duplicated watch events
					continue
				}
				w.lastRV[lastRVKey] = rv
			}

			genEvent := event.GenericEvent{}
			genEvent.Object = &unstructured.Unstructured{}
			genEvent.Object.SetNamespace(u.GetNamespace()) // raise event in the same namespace as the trigger watch event
			genEvent.Object.SetName(targetNN.Name)         // raise event on the target name
			w.events <- genEvent
		}
	}
}
