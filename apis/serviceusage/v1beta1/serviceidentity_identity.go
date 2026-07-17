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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ServiceIdentityIdentity{}
	_ identity.Resource   = &ServiceIdentity{}
)

var ServiceIdentityIdentityFormat = gcpurls.Template[ServiceIdentityIdentity]("serviceusage.googleapis.com", "projects/{project}/services/{service}/identity")

// ServiceIdentityIdentity is the identity of a GCP ServiceIdentity resource.
// +k8s:deepcopy-gen=false
type ServiceIdentityIdentity struct {
	Project string
	Service string
}

func (i *ServiceIdentityIdentity) String() string {
	return ServiceIdentityIdentityFormat.ToString(*i)
}

func (i *ServiceIdentityIdentity) FromExternal(ref string) error {
	// Strip optional scheme and host for manual parsing of legacy format
	s := strings.TrimPrefix(ref, "https:")
	s = strings.TrimPrefix(s, "http:")
	s = strings.TrimPrefix(s, "//")
	s = strings.TrimPrefix(s, "serviceusage.googleapis.com/")
	s = strings.Trim(s, "/")

	// Support legacy format for backward compatibility:
	// projects/{{projectID}}/locations/{{location}}/serviceidentitys/{{serviceidentityID}}
	tokens := strings.Split(s, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "serviceidentitys" {
		i.Project = tokens[1]
		i.Service = tokens[5]
		return nil
	}

	parsed, match, err := ServiceIdentityIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ServiceIdentity external=%q was not known (use %s): %w", ref, ServiceIdentityIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ServiceIdentity external=%q was not known (use %s)", ref, ServiceIdentityIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ServiceIdentityIdentity) Host() string {
	return ServiceIdentityIdentityFormat.Host()
}

func (i *ServiceIdentityIdentity) ParentString() string {
	return "projects/" + i.Project
}

func getIdentityFromServiceIdentitySpec(ctx context.Context, reader client.Reader, obj *ServiceIdentity) (*ServiceIdentityIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ServiceIdentityIdentity{
		Project: projectID,
		Service: resourceID,
	}
	return identity, nil
}

func (obj *ServiceIdentity) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromServiceIdentitySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	return specIdentity, nil
}
