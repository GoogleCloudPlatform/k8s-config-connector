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

package watchset

import (
	"context"
	"sync"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

type ControllerManager struct {
	mgr *Manager

	events       chan event.GenericEvent
	mutex        sync.Mutex
	interestSets map[types.NamespacedName]*InterestSet
}

func (w *Manager) NewControllerManager(controller controller.Controller) (*ControllerManager, error) {
	events := make(chan event.GenericEvent)
	if err := controller.Watch(source.TypedChannel(events, &handler.EnqueueRequestForObject{})); err != nil {
		return nil, err
	}

	return &ControllerManager{
		mgr:          w,
		interestSets: make(map[types.NamespacedName]*InterestSet),
		events:       events,
	}, nil
}

var _ client.Object = clientObject{}

// clientObject is a concrete client.Object to pass to watch events.
type clientObject struct {
	runtime.Object
	*metav1.ObjectMeta
}

func (m *ControllerManager) Stop() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for _, s := range m.interestSets {
		s.Close()
	}
	m.interestSets = nil
}

func (m *ControllerManager) ReconcileStart(ctx context.Context, id types.NamespacedName) *ControllerInterestSet {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	dependencySet := newDependencySet()

	interestSet := m.interestSets[id]
	if interestSet == nil {
		obj := &unstructured.Unstructured{}
		obj.SetNamespace(id.Namespace)
		obj.SetName(id.Name)
		interestSet = m.mgr.newInterestSet(func() {
			m.events <- event.GenericEvent{Object: obj}
		})
		m.interestSets[id] = interestSet
	}

	return &ControllerInterestSet{
		manager:       m,
		id:            id,
		DependencySet: dependencySet,
		interestSet:   interestSet,
	}
}

func (s *ControllerInterestSet) ReconcileFailed() {
	s.interestSet.AddDependencies(s.DependencySet)
}

func (s *ControllerInterestSet) ReconcileSuccess() {
	s.interestSet.ReplaceAllDependencies(s.DependencySet)
}

func (s *ControllerInterestSet) Close() {
	s.interestSet.Close()
}

type ControllerInterestSet struct {
	manager *ControllerManager
	id      types.NamespacedName
	*DependencySet
	interestSet *InterestSet
}
