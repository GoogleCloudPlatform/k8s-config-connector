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

package securesourcemanager

import (
	"fmt"
	"reflect"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/mappings"
)

type refMapKMSKey struct {
}

var _ mappings.ResourceRefMapper = &refMapKMSKey{}

func (r *refMapKMSKey) KRMToCloud(in reflect.Value) (reflect.Value, error) {
	switch in.Kind() {
	case reflect.Struct:
		ref := in.Interface().(v1alpha1.ResourceRef)
		s := ref.External
		return reflect.ValueOf(s), nil
	case reflect.Ptr:
		if in.IsNil() {
			var p *string
			return reflect.ValueOf(p), nil
		}
		return r.KRMToCloud(in.Elem())
	default:
		return reflect.Value{}, fmt.Errorf("unhandled kind in refMapKMSKey::KRMToCloud: %v", in.Kind())
	}
}

func (r *refMapKMSKey) CloudToKRM(in reflect.Value) (reflect.Value, error) {
	s := ""
	switch in.Kind() {
	case reflect.String:
		s = in.String()
	default:
		return reflect.Value{}, fmt.Errorf("unhandled kind in refMapKMSKey::CloudToKRM: %v", in.Kind())
	}

	ref := &v1alpha1.ResourceRef{}
	ref.External = s
	return reflect.ValueOf(ref), nil
}
