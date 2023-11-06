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
	"sync"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
)

// InterestSet represents a client that wants to know the ongoing status of a set of objects.
// The guarantee is that once we call ReplaceAllObjects, we will try to watch objects from that point onwards.
// TODO: Should we have a status on the InterestSet, to indicate that all our watches are healthy?
// TODO: Can we / should we fold this into applyset?  Similar data structure...
type InterestSet struct {
	parent *Manager
	mutex  sync.Mutex
	closed bool

	// subscribers []InterestSetSubscriber
	byGVR    map[schema.GroupVersionResource]*interestSetGVR
	callback func()
}

// Close indicates that this InterestSet is no longer in use.
func (s *InterestSet) Close() {
	s.mutex.Lock()
	s.closed = true
	s.byGVR = nil
	s.callback = nil
	s.mutex.Unlock()

	s.parent.updateInterests()
}

type GRNN struct {
	GR schema.GroupResource
	NN types.NamespacedName
}

// mergeDependencies adds all the dependencies to the interest set.
// If prune is true, other dependencies will be removed.
func (s *InterestSet) mergeDependencies(dependencies *DependencySet, prune bool) {
	s.mutex.Lock()
	for gvr, objects := range dependencies.objects {
		gvrInterest := s.byGVR[gvr]
		if gvrInterest == nil {
			gvrInterest = &interestSetGVR{
				objects:     make(map[types.NamespacedName]objectState),
				listObjects: make(map[types.NamespacedName]objectState),
				listOps:     make(map[listOpKey]listOpState),
			}
			s.byGVR[gvr] = gvrInterest
		}
		gvrInterest.mergeDependencies(objects, prune)
	}
	for gvr, listOps := range dependencies.lists {
		gvrInterest := s.byGVR[gvr]
		if gvrInterest == nil {
			gvrInterest = &interestSetGVR{
				objects:     make(map[types.NamespacedName]objectState),
				listObjects: make(map[types.NamespacedName]objectState),
				listOps:     make(map[listOpKey]listOpState),
			}
			s.byGVR[gvr] = gvrInterest
		}
		gvrInterest.mergeListDependencies(listOps, prune)
	}

	if prune {
		for gvk := range s.byGVR {
			if len(dependencies.objects[gvk]) == 0 {
				// TODO: clean up interest?
				delete(s.byGVR, gvk)
			}
		}
	}

	s.mutex.Unlock()

	s.parent.updateInterests()
}

// ReplaceAllDependencies completely replaces all the objects of interest.
func (s *InterestSet) ReplaceAllDependencies(dependencies *DependencySet) {
	s.mergeDependencies(dependencies, true)
}

// ReplaceAllDependencies completely replaces all the objects of interest.
func (s *InterestSet) AddDependencies(dependencies *DependencySet) {
	s.mergeDependencies(dependencies, false)
}

func (s *InterestSet) onEvent(gvr schema.GroupVersionResource, ev *objectEvent) {

	s.mutex.Lock()
	status := s.byGVR[gvr]
	s.mutex.Unlock()

	if status != nil {
		status.mutex.Lock()
		found := status.matches(ev)
		status.mutex.Unlock()
		if found {
			if s.callback != nil {
				s.callback()
			}
		}
	}
}

type objectState struct {
}
type listOpState struct {
}

type interestSetGVR struct {
	mutex       sync.Mutex
	objects     map[types.NamespacedName]objectState
	listObjects map[types.NamespacedName]objectState
	listOps     map[listOpKey]listOpState
}

func (i *interestSetGVR) mergeDependencies(objects map[types.NamespacedName]resourceVersion, prune bool) {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	for nn := range objects {
		_, found := i.objects[nn]
		if !found {
			i.objects[nn] = objectState{}
		}
	}

	if prune {
		for nn := range i.objects {
			if _, found := objects[nn]; !found {
				delete(i.objects, nn)
			}
		}
	}
}

func (i *interestSetGVR) mergeListDependencies(listOps map[listOpKey]listOp, prune bool) {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	for listOpKey := range listOps {
		_, found := i.listOps[listOpKey]
		if !found {
			i.listOps[listOpKey] = listOpState{}
		}
	}

	if prune {
		for listOpKey := range i.listOps {
			if _, found := listOps[listOpKey]; !found {
				delete(i.listOps, listOpKey)
			}
		}
	}
}

func (i *interestSetGVR) matches(ev *objectEvent) bool {
	nn := ev.ID

	_, found := i.objects[nn]
	if found {
		// A direct object match
		return true
	}

	_, wasInList := i.listObjects[nn]

	// If it was in a list previously, it may no longer match
	// If it was not in a list previously, it may now be in a list

	nowInList := false

	for k := range i.listOps {
		if k.Namespace != "" && k.Namespace != nn.Namespace {
			continue
		}
		if k.FieldSelector != "" {
			klog.Warningf("ignoring FieldSelector %q", k.FieldSelector)
		}
		if k.LabelSelector != "" {
			// TODO: Cache this
			labelSelector, err := labels.Parse(k.LabelSelector)
			if err != nil {
				klog.Warningf("error parsing labelSelector %q", k.LabelSelector)
				continue
			}

			matchesLabels := labelSelector.Matches(labels.Set(ev.Labels))
			if !matchesLabels {
				continue
			}
		}

		nowInList = true
		break
	}

	if !wasInList && nowInList {
		i.listObjects[nn] = objectState{}
	}

	if wasInList && !nowInList {
		delete(i.listObjects, nn)
	}

	// If it matches one of the list operations, we want to send it.
	// If it matched previously, we want to send if it is now _not_ in the list
	return nowInList || wasInList // technically nowInList || (wasInList && !nowInList)
}
