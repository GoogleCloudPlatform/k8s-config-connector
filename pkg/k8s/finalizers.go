// Copyright 2022 Google LLC
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

package k8s

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func EnsureFinalizer(o metav1.Object, finalizer string) (found bool) {
	for _, f := range o.GetFinalizers() {
		if f == finalizer {
			return true
		}
	}
	o.SetFinalizers(append(o.GetFinalizers(), finalizer))
	return false
}

// EnsureFinalizers adds the specified finalizers, returning true if the finalizers were already present (i.e. no changes)
func EnsureFinalizers(o metav1.Object, finalizers ...string) (found bool) {
	found = true
	for _, f := range finalizers {
		partialFound := EnsureFinalizer(o, f)
		if !partialFound {
			found = false
		}
	}
	return found
}

func RemoveFinalizer(o metav1.Object, finalizer string) {
	found := false
	var finalizers []string
	for _, f := range o.GetFinalizers() {
		if f != finalizer {
			finalizers = append(finalizers, f)
		} else {
			found = true
		}
	}
	if found {
		o.SetFinalizers(finalizers)
	}
}

func HasFinalizer(o metav1.Object, finalizer string) bool {
	for _, f := range o.GetFinalizers() {
		if f == finalizer {
			return true
		}
	}
	return false
}
