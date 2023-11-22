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
	"fmt"
	"strconv"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
)

type resourceVersion int64

type DependencySet struct {
	// Note: GroupVersionResource (instead of GroupResource) because we typically will only care about one version anyway,
	// and because sometimes we do care - when we want to read the object.
	objects map[schema.GroupVersionResource]map[types.NamespacedName]resourceVersion

	lists map[schema.GroupVersionResource]map[listOpKey]listOp
}

func newDependencySet() *DependencySet {
	return &DependencySet{
		objects: make(map[schema.GroupVersionResource]map[types.NamespacedName]resourceVersion),
		lists:   make(map[schema.GroupVersionResource]map[listOpKey]listOp),
	}
}

var _ fmt.Stringer = dependencySet{}

func (x *DependencySet) String() string {
	var sb strings.Builder
	for gvk, objects := range x.objects {
		if sb.Len() != 0 {
			fmt.Fprintf(&sb, ",")
		}
		fmt.Fprintf(&sb, "%v:[", gvk.String())
		for nn, rv := range objects {
			fmt.Fprintf(&sb, "%v@%d", nn.String(), rv)
		}
		fmt.Fprintf(&sb, "]")
	}
	for gvk, ops := range x.lists {
		if sb.Len() != 0 {
			fmt.Fprintf(&sb, ",")
		}
		fmt.Fprintf(&sb, "list:%v:[", gvk.String())
		for _, op := range ops {
			fmt.Fprintf(&sb, "%v", op)
		}
		fmt.Fprintf(&sb, "]")

	}
	return sb.String()
}

func (x *DependencySet) WatchObject(gvr schema.GroupVersionResource, nn types.NamespacedName, rv string) {
	rvInt := int64(0)
	if rv != "" {
		n, err := strconv.ParseInt(rv, 10, 64)
		if err != nil {
			klog.Fatalf("error parsing resource version %q", rv)
		}
		rvInt = n
	}

	m := x.objects[gvr]
	if m == nil {
		m = make(map[types.NamespacedName]resourceVersion)
		x.objects[gvr] = m
	}
	m[nn] = resourceVersion(rvInt)
}

type listOp struct {
}

type listOpKey struct {
	Namespace     string
	LabelSelector string
	FieldSelector string
}

func (x *DependencySet) WatchList(gvr schema.GroupVersionResource, ns string, opts metav1.ListOptions) {
	key := listOpKey{
		Namespace:     ns,
		LabelSelector: opts.LabelSelector,
		FieldSelector: opts.FieldSelector,
	}
	op := listOp{}

	m := x.lists[gvr]
	if m == nil {
		m = make(map[listOpKey]listOp)
		x.lists[gvr] = m
	}
	m[key] = op
}
