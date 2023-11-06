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

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
)

// func newObjectDependencyTracker
type dependencySet map[schema.GroupKind]map[types.NamespacedName]int64

var _ fmt.Stringer = dependencySet{}

func (x dependencySet) String() string {
	var sb strings.Builder
	for gvk, objects := range x {
		if sb.Len() != 0 {
			fmt.Fprintf(&sb, ",")
		}
		fmt.Fprintf(&sb, "%v:[", gvk.String())
		for nn, rv := range objects {
			fmt.Fprintf(&sb, "%v@%d", nn.String(), rv)
		}
		fmt.Fprintf(&sb, "]")
	}
	return sb.String()
}

func (x dependencySet) Add(gk schema.GroupKind, nn types.NamespacedName, rv string) {
	rvInt, err := strconv.ParseInt(rv, 10, 64)
	if err != nil {
		klog.Fatalf("error parsing resource version %q", rv)
	}
	m := x[gk]
	if m == nil {
		m = make(map[types.NamespacedName]int64)
		x[gk] = m
	}
	m[nn] = rvInt
}
