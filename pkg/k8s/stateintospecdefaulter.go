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

package k8s

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/kccstate"
)

// StateIntoSpecDefaulter contains the required 'defaultValue' field and the
// optional 'userOverride' field.
type StateIntoSpecDefaulter struct {
	client client.Client
}

func NewStateIntoSpecDefaulter(client client.Client) Defaulter {
	return &StateIntoSpecDefaulter{
		client: client,
	}
}

func (v *StateIntoSpecDefaulter) ApplyDefaults(ctx context.Context, resource client.Object) (changed bool, err error) {
	val, found := GetAnnotation(StateIntoSpecAnnotation, resource)
	if found {
		if !isAcceptedStateIntoSpecValue(val, resource.GetObjectKind().GroupVersionKind()) {
			return false, fmt.Errorf("invalid value %q for %q annotation in kind %v", val, StateIntoSpecAnnotation, resource.GetObjectKind().GroupVersionKind().Kind)
		}
		return false, nil
	}

	annotationValue := StateIntoSpecDefaultValueV1Beta1

	namespacedName := types.NamespacedName{Name: resource.GetName(), Namespace: resource.GetNamespace()}
	cc, ccc, err := kccstate.FetchLiveKCCState(ctx, v.client, namespacedName)
	if err != nil {
		return false, fmt.Errorf("error getting ConfigConnector and ConfigConnectorContext objects: %w", err)
	}

	if cc.Spec.StateIntoSpec != nil {
		switch *cc.Spec.StateIntoSpec {
		case operatorv1beta1.StateIntoSpecMerge:
			annotationValue = StateMergeIntoSpec
		case operatorv1beta1.StateIntoSpecAbsent:
			annotationValue = StateAbsentInSpec

		default:
			return false, fmt.Errorf("invalid value %q for spec.stateIntoSpec in ConfigConnector, should be Absent or Merge (Absent recommended)", *cc.Spec.StateIntoSpec)
		}
	}

	if ccc.Spec.StateIntoSpec != nil {
		switch *ccc.Spec.StateIntoSpec {
		case operatorv1beta1.StateIntoSpecMerge:
			annotationValue = StateMergeIntoSpec
		case operatorv1beta1.StateIntoSpecAbsent:
			annotationValue = StateAbsentInSpec

		default:
			return false, fmt.Errorf("invalid value %q for spec.stateIntoSpec in ConfigConnectorContext, should be Absent or Merge (Absent recommended)", *ccc.Spec.StateIntoSpec)
		}
	}

	setStateIntoSpecDefaultValueIfAllowed(annotationValue, resource)
	return true, nil
}

func setStateIntoSpecDefaultValueIfAllowed(defaultValue string, obj client.Object) {
	if defaultValue == StateMergeIntoSpec && !SupportsStateIntoSpecMerge(obj.GetObjectKind().GroupVersionKind()) {
		klog.Infof("%v doesn't support %q so the %q annotation is always defaulted to %q",
			obj.GetObjectKind().GroupVersionKind().Kind, StateMergeIntoSpec,
			StateIntoSpecAnnotation, StateAbsentInSpec)
		SetAnnotation(StateIntoSpecAnnotation, StateAbsentInSpec, obj)
		return

	}
	SetAnnotation(StateIntoSpecAnnotation, defaultValue, obj)
}

func isAcceptedStateIntoSpecValue(value string, gvk schema.GroupVersionKind) bool {
	if SupportsStateIntoSpecMerge(gvk) {
		for _, v := range StateIntoSpecAnnotationValues {
			if value == v {
				return true
			}
		}
		return false
	}
	return value == StateAbsentInSpec
}
