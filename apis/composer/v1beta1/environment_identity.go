// Copyright 2025 Google LLC
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
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &EnvironmentIdentity{}

// EnvironmentIdentity defines the resource reference to ComposerEnvironment, which "External" field
// holds the GCP identifier for the KRM object.
type EnvironmentIdentity struct {
	parent *parent.ProjectAndLocationParent
	id     string
}

func (i *EnvironmentIdentity) String() string {
	return i.parent.String() + "/environments/" + i.id
}

func (i *EnvironmentIdentity) ID() string {
	return i.id
}

func (i *EnvironmentIdentity) Parent() *parent.ProjectAndLocationParent {
	return i.parent
}

func (i *EnvironmentIdentity) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/environments/")
	if len(tokens) != 2 {
		return fmt.Errorf("format of ComposerEnvironment external=%q was not known (use projects/{{projectID}}/locations/{{location}}/environments/{{environmentID}})", ref)
	}
	i.parent = &parent.ProjectAndLocationParent{}
	if err := i.parent.FromExternal(tokens[0]); err != nil {
		return err
	}
	i.id = tokens[1]
	if i.id == "" {
		return fmt.Errorf("environmentID was empty in external=%q", ref)
	}
	return nil
}

var _ identity.Resource = &ComposerEnvironment{}

func (obj *ComposerEnvironment) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	environment := &EnvironmentIdentity{
		parent: &parent.ProjectAndLocationParent{},
	}

	// Resolve user-configured Parent
	if err := obj.Spec.ParentRef.Build(ctx, reader, obj.GetNamespace(), environment.parent); err != nil {
		return nil, err
	}

	// Get user-configured ID
	environment.id = common.ValueOf(obj.Spec.ResourceID)
	if environment.id == "" {
		environment.id = obj.GetName()
	}
	if environment.id == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Validate against the ID stored in status.externalRef, if any
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &EnvironmentIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != environment.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, environment.String())
		}
	}
	return environment, nil
}
