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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/kccscheme"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Register registers a Ref implementation and optionally its corresponding runtime.Object.
// It is thread-safe.
func Register(ref Ref, objs ...runtime.Object) {
	kccscheme.RegisterRef(ref, ref.GetGVK())

	if len(objs) > 0 && objs[0] != nil {
		kccscheme.RegisterType(ref.GetGVK(), objs[0])
	}
}

// NewRef returns a new instance of Ref for the given GroupKind.
// It is thread-safe.
func NewRef(gk schema.GroupKind) (Ref, error) {
	ref, err := kccscheme.NewRef(gk)
	if err != nil {
		return nil, err
	}
	return ref.(Ref), nil
}

// NewRefByKind returns a new instance of Ref for the given Kind.
// It is thread-safe.
func NewRefByKind(kind string) (Ref, error) {
	ref, err := kccscheme.NewRefByKind(kind)
	if err != nil {
		return nil, err
	}
	return ref.(Ref), nil
}

// NewObject returns a new instance of runtime.Object for the given GroupKind.
// It is thread-safe.
func NewObject(gk schema.GroupKind) (runtime.Object, error) {
	return kccscheme.NewObject(gk)
}

// PreferredGVK returns the preferred GroupVersionKind for the given GroupKind.
// It is thread-safe.
func PreferredGVK(gk schema.GroupKind) (schema.GroupVersionKind, bool) {
	return kccscheme.PreferredGVK(gk)
}
