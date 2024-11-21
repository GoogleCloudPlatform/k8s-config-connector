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
	"strings"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/testing/protocmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

func lastComponent(s string) string {
	i := strings.LastIndex(s, "/")
	return s[i+1:]
}

func ComputeChangedFields(actual proto.Message, desired proto.Message) sets.Set[string] {
	changes := sets.New[string]()
	actualReflect := actual.ProtoReflect()
	desiredReflect := desired.ProtoReflect()
	actualReflect.Range(func(field protoreflect.FieldDescriptor, actualValue protoreflect.Value) bool {
		desiredValue := desiredReflect.Get(field)
		if !actualValue.Equal(desiredValue) {
			changes.Insert(field.JSONName())
		}
		return true
	})
	desiredReflect.Range(func(field protoreflect.FieldDescriptor, desiredValue protoreflect.Value) bool {
		actualValue := actualReflect.Get(field)
		if !actualValue.Equal(desiredValue) {
			changes.Insert(field.JSONName())
		}
		return true
	})
	if changes.Len() != 0 {
		klog.V(2).Infof("ComputeChangedFields found diff fields=%v, diff=%v", sets.List(changes), cmp.Diff(actual, desired, protocmp.Transform()))
	}
	return changes
}

func setStatus(u *unstructured.Unstructured, typedStatus any) error {
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}

	// Use existing values for conditions/observedGeneration; they are managed in k8s not the GCP API
	old, _, _ := unstructured.NestedMap(u.Object, "status")
	if old != nil {
		status["conditions"] = old["conditions"]
		status["observedGeneration"] = old["observedGeneration"]
	}

	u.Object["status"] = status

	return nil
}
