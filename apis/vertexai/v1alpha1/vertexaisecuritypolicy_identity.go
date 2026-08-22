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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &VertexAISecurityPolicyIdentity{}
	_ identity.Resource   = &VertexAISecurityPolicy{}
)

var VertexAISecurityPolicyIdentityFormat = gcpurls.Template[VertexAISecurityPolicyIdentity]("modelarmor.googleapis.com", "projects/{project}/locations/{location}/templates/{template}")

// VertexAISecurityPolicyIdentity is the identity of a GCP VertexAISecurityPolicy resource.
// +k8s:deepcopy-gen=false
type VertexAISecurityPolicyIdentity struct {
	Project  string
	Location string
	Template string
}

func (i *VertexAISecurityPolicyIdentity) String() string {
	return VertexAISecurityPolicyIdentityFormat.ToString(*i)
}

func (i *VertexAISecurityPolicyIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAISecurityPolicyIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAISecurityPolicy external=%q was not known (use %s): %w", ref, VertexAISecurityPolicyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAISecurityPolicy external=%q was not known (use %s)", ref, VertexAISecurityPolicyIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *VertexAISecurityPolicyIdentity) Host() string {
	return VertexAISecurityPolicyIdentityFormat.Host()
}

func (i *VertexAISecurityPolicyIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func getIdentityFromVertexAISecurityPolicySpec(ctx context.Context, reader client.Reader, obj *VertexAISecurityPolicy) (*VertexAISecurityPolicyIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &VertexAISecurityPolicyIdentity{
		Project:  projectID,
		Location: location,
		Template: resourceID,
	}
	return identity, nil
}

func (obj *VertexAISecurityPolicy) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromVertexAISecurityPolicySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &VertexAISecurityPolicyIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change VertexAISecurityPolicy identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
