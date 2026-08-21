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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ServiceDirectoryEndpointIdentity{}
	_ identity.Resource   = &ServiceDirectoryEndpoint{}
)

var ServiceDirectoryEndpointIdentityFormat = gcpurls.Template[ServiceDirectoryEndpointIdentity]("servicedirectory.googleapis.com", "projects/{project}/locations/{location}/namespaces/{namespace}/services/{service}/endpoints/{endpoint}")

// +k8s:deepcopy-gen=false

// ServiceDirectoryEndpointIdentity is the identity of a GCP ServiceDirectoryEndpoint resource.
type ServiceDirectoryEndpointIdentity struct {
	Project   string
	Location  string
	Namespace string
	Service   string
	Endpoint  string
}

func (i *ServiceDirectoryEndpointIdentity) String() string {
	return ServiceDirectoryEndpointIdentityFormat.ToString(*i)
}

func (i *ServiceDirectoryEndpointIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s", i.Project, i.Location, i.Namespace, i.Service)
}

func (i *ServiceDirectoryEndpointIdentity) FromExternal(ref string) error {
	ref = strings.TrimPrefix(ref, "https://servicedirectory.googleapis.com/v1/")
	ref = strings.TrimPrefix(ref, "https://servicedirectory.googleapis.com/v1beta1/")
	ref = strings.TrimPrefix(ref, "https://servicedirectory.googleapis.com/")

	parsed, match, err := ServiceDirectoryEndpointIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ServiceDirectoryEndpoint external=%q was not known (use %s): %w", ref, ServiceDirectoryEndpointIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ServiceDirectoryEndpoint external=%q was not known (use %s)", ref, ServiceDirectoryEndpointIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ServiceDirectoryEndpointIdentity) Host() string {
	return ServiceDirectoryEndpointIdentityFormat.Host()
}

func getIdentityFromServiceDirectoryEndpointSpec(ctx context.Context, reader client.Reader, obj *ServiceDirectoryEndpoint) (*ServiceDirectoryEndpointIdentity, error) {
	serviceRef := obj.Spec.ServiceRef
	if err := (&serviceRef).Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("resolving spec.serviceRef: %w", err)
	}

	serviceIdentity := &ServiceDirectoryServiceIdentity{}
	if err := serviceIdentity.FromExternal(serviceRef.GetExternal()); err != nil {
		return nil, fmt.Errorf("parsing serviceRef.external=%q: %w", serviceRef.GetExternal(), err)
	}

	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	identity := &ServiceDirectoryEndpointIdentity{
		Project:   serviceIdentity.Project,
		Location:  serviceIdentity.Location,
		Namespace: serviceIdentity.Namespace,
		Service:   serviceIdentity.Service,
		Endpoint:  resourceID,
	}
	return identity, nil
}

func (obj *ServiceDirectoryEndpoint) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromServiceDirectoryEndpointSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.Name)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &ServiceDirectoryEndpointIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ServiceDirectoryEndpoint identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
