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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ParameterManagerParameterVersionIdentity{}
	_ identity.Resource   = &ParameterManagerParameterVersion{}
)

var ParameterManagerParameterVersionIdentityFormat = gcpurls.Template[ParameterManagerParameterVersionIdentity]("parametermanager.googleapis.com", "projects/{project}/locations/{location}/parameters/{parameter}/versions/{version}")

// +k8s:deepcopy-gen=false
type ParameterManagerParameterVersionIdentity struct {
	Project   string
	Location  string
	Parameter string
	Version   string
}

func (i *ParameterManagerParameterVersionIdentity) String() string {
	return ParameterManagerParameterVersionIdentityFormat.ToString(*i)
}

func (i *ParameterManagerParameterVersionIdentity) FromExternal(ref string) error {
	parsed, match, err := ParameterManagerParameterVersionIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ParameterManagerParameterVersion external=%q was not known (use %s): %w", ref, ParameterManagerParameterVersionIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ParameterManagerParameterVersion external=%q was not known (use %s)", ref, ParameterManagerParameterVersionIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ParameterManagerParameterVersionIdentity) Host() string {
	return ParameterManagerParameterVersionIdentityFormat.Host()
}

func getIdentityFromParameterManagerParameterVersionSpec(ctx context.Context, reader client.Reader, obj client.Object) (*ParameterManagerParameterVersionIdentity, error) {
	var parameterRef *ParameterRef
	var resourceID *string
	var namespace string

	if typed, ok := obj.(*ParameterManagerParameterVersion); ok {
		parameterRef = typed.Spec.ParameterRef
		resourceID = typed.Spec.ResourceID
		namespace = typed.GetNamespace()
	} else if u, ok := obj.(*unstructured.Unstructured); ok {
		namespace = u.GetNamespace()

		rawParameterRef, found, err := unstructured.NestedMap(u.Object, "spec", "parameterRef")
		if err != nil || !found {
			return nil, fmt.Errorf("cannot find spec.parameterRef")
		}

		parameterRef = &ParameterRef{}
		if val, ok := rawParameterRef["external"].(string); ok {
			parameterRef.External = val
		}
		if val, ok := rawParameterRef["name"].(string); ok {
			parameterRef.Name = val
		}
		if val, ok := rawParameterRef["namespace"].(string); ok {
			parameterRef.Namespace = val
		}

		if val, found, err := unstructured.NestedString(u.Object, "spec", "resourceID"); err == nil && found {
			resourceID = &val
		}
	} else {
		return nil, fmt.Errorf("unsupported object type: %T", obj)
	}

	if parameterRef == nil {
		return nil, fmt.Errorf("parameterRef is required")
	}

	// Normalize parent reference
	if err := parameterRef.Normalize(ctx, reader, namespace); err != nil {
		return nil, fmt.Errorf("failed to normalize parameterRef: %w", err)
	}

	// Resolve parent identity
	parentID := &ParameterIdentity{}
	if err := parentID.FromExternal(parameterRef.External); err != nil {
		return nil, fmt.Errorf("failed to parse parent externalRef %q: %w", parameterRef.External, err)
	}

	idVal := common.ValueOf(resourceID)
	if idVal == "" {
		idVal = obj.GetName()
	}
	if idVal == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	identity := &ParameterManagerParameterVersionIdentity{
		Project:   parentID.Parent().ProjectID,
		Location:  parentID.Parent().Location,
		Parameter: parentID.ID(),
		Version:   idVal,
	}
	return identity, nil
}

func (obj *ParameterManagerParameterVersion) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromParameterManagerParameterVersionSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Validate against the ID stored in status.externalRef, if any
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &ParameterManagerParameterVersionIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, specIdentity.String())
		}
	}

	return specIdentity, nil
}
