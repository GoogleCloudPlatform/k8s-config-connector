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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &VertexAIIndexEndpointIdentity{}
	_ identity.Resource   = &VertexAIIndexEndpoint{}
)

var VertexAIIndexEndpointIdentityFormat = gcpurls.Template[VertexAIIndexEndpointIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/indexEndpoints/{indexendpoint}")

// +k8s:deepcopy-gen=false
type VertexAIIndexEndpointIdentity struct {
	Project       string
	Location      string
	Indexendpoint string
}

func (i *VertexAIIndexEndpointIdentity) String() string {
	return VertexAIIndexEndpointIdentityFormat.ToString(*i)
}

func (i *VertexAIIndexEndpointIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAIIndexEndpointIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAIIndexEndpoint external=%q was not known (use %s): %w", ref, VertexAIIndexEndpointIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAIIndexEndpoint external=%q was not known (use %s)", ref, VertexAIIndexEndpointIdentityFormat.CanonicalForm())
	}
	*i = *parsed
	return nil
}

func (i *VertexAIIndexEndpointIdentity) Host() string {
	return VertexAIIndexEndpointIdentityFormat.Host()
}

func (obj *VertexAIIndexEndpoint) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	id, err := getIdentityFromVertexAIIndexEndpointSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	if obj.Status.ExternalRef != nil {
		actual := &VertexAIIndexEndpointIdentity{}
		if err := actual.FromExternal(*obj.Status.ExternalRef); err != nil {
			return nil, err
		}
		if id.String() != actual.String() {
			return nil, fmt.Errorf("cannot change VertexAIIndexEndpoint identity (old=%q, new=%q)", actual.String(), id.String())
		}
	}
	return id, nil
}

func getIdentityFromVertexAIIndexEndpointSpec(ctx context.Context, reader client.Reader, obj client.Object) (*VertexAIIndexEndpointIdentity, error) {
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

	identity := &VertexAIIndexEndpointIdentity{
		Project:       projectID,
		Location:      location,
		Indexendpoint: resourceID,
	}

	return identity, nil
}
