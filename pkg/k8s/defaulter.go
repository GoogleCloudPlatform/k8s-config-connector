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

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	operatork8s "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
)

type Defaulter interface {
	ApplyDefaults(ctx context.Context, obj client.Object) (changed bool, err error)
}

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
	annotationValue := StateIntoSpecDefaultValueV1Beta1

	cccNamespacedName := types.NamespacedName{
		Namespace: resource.GetNamespace(),
		Name:      operatork8s.ConfigConnectorContextAllowedName,
	}
	ccc := &operatorv1beta1.ConfigConnectorContext{}
	if err := v.client.Get(ctx, cccNamespacedName, ccc); err != nil {
		if apierrors.IsNotFound(err) {
			ccc = nil
		} else {
			return false, fmt.Errorf("error getting ConfigConnectorContext object %v/%v: %w", cccNamespacedName.Namespace, cccNamespacedName.Name, err)
		}
	}

	if ccc != nil && ccc.Spec.StateIntoSpec != nil {
		switch *ccc.Spec.StateIntoSpec {
		case operatorv1beta1.StateIntoSpecMerge:
			annotationValue = StateMergeIntoSpec
		case operatorv1beta1.StateIntoSpecAbsent:
			annotationValue = StateAbsentInSpec

		default:
			return false, fmt.Errorf("invalid value %q for spec.stateIntoSpec, should be Absent or Merge (Absent recommended)", *ccc.Spec.StateIntoSpec)
		}
	}

	// Validate or set the default value (cluster-level or namespace-level) for
	// the 'state-into-spec' annotation.
	if err := ValidateOrDefaultStateIntoSpecAnnotation(resource, annotationValue); err != nil {
		return false, fmt.Errorf("error validating or defaulting the '%v' annotation for resource '%v': %w", StateIntoSpecAnnotation, GetNamespacedName(resource), err)
	}
	return true, nil
}
