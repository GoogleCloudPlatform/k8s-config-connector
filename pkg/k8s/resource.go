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
	"encoding/json"
	"fmt"
	"reflect"

	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/structured-merge-diff/v4/fieldpath"
)

// Resource represents a resource in KRM
type Resource struct {
	// Fundamental fields
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              map[string]interface{} `json:"spec,omitempty"`
	Status            map[string]interface{} `json:"status,omitempty"`

	// Fields related to KRM processing

	// ManagedFields is the set of spec fields whose desired state is managed
	// by Kubernetes. Fields that are not part of this set are considered
	// unmanaged, and their values in etcd will be updated to match the
	// underlying API.
	//
	// If this object is nil, all fields in the spec in etcd are considered
	// managed and their values will be constantly enforced.
	ManagedFields *fieldpath.Set `json:"-"`
}

func (r *Resource) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: r.GetNamespace(),
		Name:      r.GetName(),
	}
}

// NewResource creates a Resource based on the given unstructured. NewResource
// can be used to create deep copies of a Resource by calling NewResource
// multiple times on the same unstructured since the Resource objects created
// are separate copies.
func NewResource(u *unstructured.Unstructured) (*Resource, error) {
	resource := &Resource{}
	if err := util.Marshal(u, resource); err != nil {
		return nil, err
	}
	managedFields, err := GetK8sManagedFields(u)
	if err != nil {
		return nil, err
	}
	resource.ManagedFields = managedFields
	return resource, nil
}

func (r *Resource) MarshalAsUnstructured() (*unstructured.Unstructured, error) {
	u := &unstructured.Unstructured{}
	if err := util.Marshal(r, u); err != nil {
		return nil, fmt.Errorf("error marshing resource to Unstructured %w", err)
	}
	removeNilCreationTimestamp(u.Object)
	return u, nil
}

func (r *Resource) IsResourceIDConfigured() (bool, error) {
	val, exists, err := unstructured.NestedString(r.Spec, ResourceIDFieldName)
	if err != nil {
		return false, fmt.Errorf("error getting the value of "+
			"\"spec.%s\": %w", ResourceIDFieldName, err)
	}

	if !exists {
		return false, nil
	}

	if val == "" {
		return false, fmt.Errorf("the value of '%s' is invalid: '' (empty "+
			"string)", ResourceIDFieldPath)
	}
	return true, nil
}

func (r *Resource) HasResourceIDField() bool {
	_, ok := r.Spec[ResourceIDFieldName]
	return ok
}

// The creation timestamp value is 'nil' after marshalling a new ObjectMeta
func removeNilCreationTimestamp(object map[string]interface{}) {
	metadata, ok := object["metadata"].(map[string]interface{})
	if !ok {
		panic("expected object to have a metadata field of type 'map[string]interface{}'")
	}
	creationTimestampKey := "creationTimestamp"
	if _, ok := metadata[creationTimestampKey]; ok {
		if creationTimeStamp, ok := metadata[creationTimestampKey]; ok {
			if creationTimeStamp == nil {
				delete(metadata, creationTimestampKey)
			}
		}
	}
}

func IsResourceReady(r *Resource) bool {
	cond, found := GetReadyCondition(r)
	return found && cond.Status == corev1.ConditionTrue
}

func GetReadyCondition(r *Resource) (condition k8sv1alpha1.Condition, found bool) {
	if currConditionsRaw, ok := r.Status["conditions"].([]interface{}); ok {
		if currConditions, err := MarshalAsConditionsSlice(currConditionsRaw); err == nil {
			for _, condition := range currConditions {
				if condition.Type == k8sv1alpha1.ReadyConditionType {
					return condition, true
				}
			}
		}
	}
	return k8sv1alpha1.Condition{}, false
}

func ReadyConditionMatches(resource *Resource, status corev1.ConditionStatus, rs, msg string) bool {
	cond, found := GetReadyCondition(resource)
	if !found {
		return false
	}
	return ConditionsEqualIgnoreTransitionTime(cond, NewCustomReadyCondition(status, rs, msg))
}

func IsSpecOrStatusUpdateRequired(resource *Resource, original *Resource) bool {
	if !reflect.DeepEqual(resource.Spec, original.Spec) {
		return true
	}
	if !reflect.DeepEqual(resource.Status, original.Status) {
		return true
	}
	// JSON marshall will turn all numbers to float64 type, we convert generation to float64 for comparison
	if len(resource.Status) == 0 || resource.Status["observedGeneration"] != float64(original.GetGeneration()) {
		return true
	}
	return false
}

func IsAnnotationsUpdateRequired(resource *Resource, original *Resource) bool {
	return !reflect.DeepEqual(resource.GetAnnotations(), original.GetAnnotations())
}

func MarshalObjectAsUnstructured(o metav1.Object) (*unstructured.Unstructured, error) {
	b, err := json.Marshal(o)
	if err != nil {
		return nil, fmt.Errorf("error marshalling object %v: %w", o.GetName(), err)
	}
	u := &unstructured.Unstructured{}
	if err := json.Unmarshal(b, u); err != nil {
		return nil, fmt.Errorf("error unmarshalling object %v to unstructured.Unstructured: %w", o.GetName(), err)
	}
	removeNilCreationTimestamp(u.Object)
	return u, nil
}

func IsDirectByAnnotation(r *Resource) bool {
	return r.Annotations[AlphaReconcilerAnnotation] == "direct"
}
