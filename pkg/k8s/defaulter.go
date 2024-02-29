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
	"strings"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Defaulter interface {
	ApplyDefaults(ctx context.Context, obj client.Object) (changed bool, err error)
}

// StateIntoSpecDefaulter contains the required 'defaultValue' field and the
// optional 'userOverride' field.
type StateIntoSpecDefaulter struct {
	defaultValue string
	userOverride *string
}

func NewStateIntoSpecDefaulter(defaultValue string, userOverride *string) (Defaulter, error) {
	if !isAcceptedValue(defaultValue, StateIntoSpecAnnotationValues) {
		return nil, fmt.Errorf("invalid default value '%v' for '%v' annotation, need to be one of {%v}", defaultValue, StateIntoSpecAnnotation, strings.Join(StateIntoSpecAnnotationValues, ", "))
	}
	if userOverride != nil && !isAcceptedValue(*userOverride, StateIntoSpecAnnotationValues) {
		return nil, fmt.Errorf("invalid user override value '%v' for '%v' annotation, need to be one of {%v}", userOverride, StateIntoSpecAnnotation, strings.Join(StateIntoSpecAnnotationValues, ", "))
	}
	return &StateIntoSpecDefaulter{
		defaultValue: defaultValue,
		userOverride: userOverride,
	}, nil
}

func (v *StateIntoSpecDefaulter) ApplyDefaults(_ context.Context, resource client.Object) (changed bool, err error) {
	// Validate or set the default value (cluster-level or namespace-level) for
	// the 'state-into-spec' annotation.
	if err := ValidateOrDefaultStateIntoSpecAnnotation(resource, v.getValue()); err != nil {
		return false, fmt.Errorf("error validating or defaulting the '%v' annotation for resource '%v': %w", StateIntoSpecAnnotation, GetNamespacedName(resource), err)
	}
	return true, nil
}

func (v *StateIntoSpecDefaulter) getValue() string {
	if v.userOverride == nil {
		return v.defaultValue
	}
	return *v.userOverride
}
