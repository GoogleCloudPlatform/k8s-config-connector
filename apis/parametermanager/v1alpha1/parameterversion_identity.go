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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &ParameterVersionIdentity{}

// ParameterIdentity defines the resource reference to BigLakeCatalog, which "External" field
// holds the GCP identifier for the KRM object.
type ParameterVersionIdentity struct {
	parent *ParameterIdentity
	id     string
}

func (i *ParameterVersionIdentity) String() string {
	return i.parent.String() + "/versions/" + i.id
}

func (i *ParameterVersionIdentity) ID() string {
	return i.id
}

func (i *ParameterVersionIdentity) Parent() *ParameterIdentity {
	return i.parent
}

func (i *ParameterVersionIdentity) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/versions/")
	if len(tokens) != 2 {
		return fmt.Errorf("format of parameters external=%q was not known (use projects/{{projectID}}/locations/{{location}}/parameters/{{parameterID}}/versions/{{versionID}})", ref)
	}
	i.parent = &ParameterIdentity{}
	if err := i.parent.FromExternal(tokens[0]); err != nil {
		return err
	}
	i.id = tokens[1]
	if i.id == "" {
		return fmt.Errorf("versionID was empty in external=%q", ref)
	}
	return nil
}

var _ identity.Resource = &ParameterManagerParameterVersion{}

func (obj *ParameterManagerParameterVersion) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Get parent ID
	parentID, err := obj.GetParentIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	// Get resource ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	parameterVersion := &ParameterVersionIdentity{
		parent: parentID.(*ParameterIdentity),
		id:     resourceID,
	}

	// Validate against the ID stored in status.externalRef, if any
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &ParameterVersionIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != parameterVersion.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, parameterVersion.String())
		}
	}

	return parameterVersion, nil
}

func (obj *ParameterManagerParameterVersion) GetParentIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Normalize parent reference
	if err := obj.Spec.ParameterRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}
	// Get parent identity
	parentID := &ParameterIdentity{}
	if err := parentID.FromExternal(obj.Spec.ParameterRef.External); err != nil {
		return nil, err
	}
	return parentID, nil
}
