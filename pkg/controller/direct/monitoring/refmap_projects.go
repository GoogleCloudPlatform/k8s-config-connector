// Copyright 2024 Google LLC
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

package monitoring

import (
	"fmt"
	"reflect"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/mappings"
)

type refMapProjects struct {
}

var _ mappings.ResourceRefMapper = &refMapProjects{}

func (r *refMapProjects) KRMToCloud(in reflect.Value) (reflect.Value, error) {
	switch in.Kind() {
	// TODO: Push slice conversion "up"?
	case reflect.Slice:
		n := in.Len()
		var out []string
		for i := 0; i < n; i++ {
			v := in.Index(i)
			v2, err := r.KRMToCloud(v)
			if err != nil {
				return reflect.Value{}, err
			}
			out = append(out, v2.String())
		}
		return reflect.ValueOf(out), nil
	case reflect.Struct:
		ref := in.Interface().(v1alpha1.ResourceRef)
		s := ref.External
		return reflect.ValueOf(s), nil
	default:
		return reflect.Value{}, fmt.Errorf("unhandled kind in refMapProjects::KRMToCloud: %v", in.Kind())
	}

}

func (r *refMapProjects) CloudToKRM(in reflect.Value) (reflect.Value, error) {
	s := ""
	switch in.Kind() {
	case reflect.String:
		s = in.String()
	// TODO: Push slice conversion "up"?
	case reflect.Slice:
		n := in.Len()
		var out []v1alpha1.ResourceRef
		for i := 0; i < n; i++ {
			v := in.Index(i)
			v2, err := r.CloudToKRM(v)
			if err != nil {
				return reflect.Value{}, err
			}
			out = append(out, *v2.Interface().(*v1alpha1.ResourceRef))
		}
		return reflect.ValueOf(out), nil
	default:
		return reflect.Value{}, fmt.Errorf("unhandled kind in refMapProjects::CloudToKRM: %v", in.Kind())
	}

	ref := &v1alpha1.ResourceRef{}
	ref.External = s
	return reflect.ValueOf(ref), nil
}
