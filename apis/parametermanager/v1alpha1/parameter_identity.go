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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &ParameterIdentity{}

// ParameterIdentity defines the resource reference to ParameterManagerParameter, which "External" field
// holds the GCP identifier for the KRM object.
type ParameterIdentity struct {
	parent *parent.ProjectAndLocationParent
	id     string
}

func (i *ParameterIdentity) String() string {
	return i.parent.String() + "/parameters/" + i.id
}

func (i *ParameterIdentity) ID() string {
	return i.id
}

func (i *ParameterIdentity) Parent() *parent.ProjectAndLocationParent {
	return i.parent
}

func (i *ParameterIdentity) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/parameters/")
	if len(tokens) != 2 {
		return fmt.Errorf("format of parameters external=%q was not known (use projects/{{projectID}}/locations/{{location}}/parameters/{{parameterID}})", ref)
	}
	i.parent = &parent.ProjectAndLocationParent{}
	if err := i.parent.FromExternal(tokens[0]); err != nil {
		return err
	}
	i.id = tokens[1]
	if i.id == "" {
		return fmt.Errorf("parameterID was empty in external=%q", ref)
	}
	return nil
}

var _ identity.Resource = &ParameterManagerParameter{}

func (obj *ParameterManagerParameter) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	parameter := &ParameterIdentity{}
	parameter.parent = &parent.ProjectAndLocationParent{}

	// Resolve user-configured Parent
	if err := obj.Spec.ProjectAndLocationRef.Build(ctx, reader, obj.GetNamespace(), parameter.parent); err != nil {
		return nil, err
	}

	// Get user-configured ID
	parameter.id = common.ValueOf(obj.Spec.ResourceID)
	if parameter.id == "" {
		parameter.id = obj.GetName()
	}
	if parameter.id == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Validate against the ID stored in status.externalRef, if any
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &ParameterIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != parameter.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, parameter.String())
		}
	}
	return parameter, nil
}
