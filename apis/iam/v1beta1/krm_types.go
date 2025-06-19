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

package v1beta1

import (
	"fmt"
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

func (ref *ResourceReference) GroupVersionKind() schema.GroupVersionKind {
	return schema.FromAPIVersionAndKind(ref.APIVersion, ref.Kind)
}

func (ref *ResourceReference) SetGroupVersionKind(gvk schema.GroupVersionKind) {
	ref.APIVersion, ref.Kind = gvk.ToAPIVersionAndKind()
}

func (ms *MemberSource) Validate() error {
	v := reflect.ValueOf(ms).Elem()
	var count int
	for i := 0; i < v.NumField(); i++ {
		if !v.Field(i).IsNil() {
			count++
		}
	}
	if count > 1 {
		return fmt.Errorf("%d memberFrom refs found. Only one subfield of MemberSource can be set", count)
	}
	return nil
}
