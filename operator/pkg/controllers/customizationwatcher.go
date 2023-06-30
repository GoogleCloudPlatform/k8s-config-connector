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

	customizev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/customize/v1alpha1"
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
	// CustomizationCRDsToWatch contains all the customization CRDs to watch
	CustomizationCRDsToWatch = []schema.GroupVersionResource{
		corekcck8s.ToGVR(customizev1alpha1.ControllerResourceGroupVersionKind),
	}
)

// CustomizationWatcher setup watches on 'triggerGVRs', raises events on 'target'.
type CustomizationWatcher struct {
	// triggerGVRs contains all the GVRs to watch. An event on triggerGVRs will raise an event on the target.
	triggerGVRs []schema.GroupVersionResource
	targetNN    types.NamespacedName
	// watchRegistered tracks the GVRs we are currently watching, avoid duplicate watches.
	watchRegistered map[string]struct{}
	// events is the channel that an event is raised and send to when CustomizationWatcher
	// receives a watch event on the triggerGVRs it is watching.
	events        chan event.GenericEvent
	dynamicClient dynamic.Interface
	log           logr.Logger
}

func NewWithDynamicClient(dc dynamic.Interface, gvrs []schema.GroupVersionResource, logger logr.Logger) *CustomizationWatcher {
	return &CustomizationWatcher{
		dynamicClient:   dc,
		triggerGVRs:     gvrs,
		watchRegistered: make(map[string]struct{}),
		events:          make(chan event.GenericEvent),
		log:             logger.WithName("customization-watcher"),
	}
}

// Events returns a channel with events raised on target.
func (w *CustomizationWatcher) Events() chan event.GenericEvent {
	return w.events
}

// EnsureWatchStarted starts watches on triggerGVRs if not already done so.
func (w *CustomizationWatcher) EnsureWatchStarted(ctx context.Context, nn types.NamespacedName) error {
	w.targetNN = nn
	for _, gvr := range w.triggerGVRs {
		if _, started := w.watchRegistered[gvr.String()]; started {
			continue
		}
		w.log.Info("starting watch", "trigger GVR", gvr.String(), "target NamespacedName", nn)
		go w.startWatch(ctx, gvr)
	}
	return nil
}

// startWatch starts a watch for changes to "triggerGVR" and raises events on "targetNN".
func (w *CustomizationWatcher) startWatch(ctx context.Context, triggerGVR schema.GroupVersionResource) {
	log := w.log.WithValues("trigger GVR", triggerGVR, "target NamespacedName", w.targetNN)
	opts := metav1.ListOptions{AllowWatchBookmarks: true}
	triggerEvents, err := w.dynamicClient.Resource(triggerGVR).Namespace(w.targetNN.Namespace).Watch(ctx, opts)
	if err != nil {
		log.Error(err, "failed to start watch")
		return
	}
	log.Info("watch started")
	w.watchRegistered[triggerGVR.String()] = struct{}{}
	defer func() {
		triggerEvents.Stop()
		delete(w.watchRegistered, triggerGVR.String())
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
				log.Error(err, "unexpected error from watch")
			}
			genEvent := event.GenericEvent{}
			genEvent.Object = &unstructured.Unstructured{}
			genEvent.Object.SetNamespace(w.targetNN.Namespace)
			genEvent.Object.SetName(w.targetNN.Name)
			w.events <- genEvent
		}
	}
}
