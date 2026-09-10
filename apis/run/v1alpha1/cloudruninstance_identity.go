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
	_ identity.IdentityV2 = &CloudRunInstanceIdentity{}
	_ identity.Resource   = &CloudRunInstance{}
)

var CloudRunInstanceIdentityFormat = gcpurls.Template[CloudRunInstanceIdentity]("run.googleapis.com", "projects/{project}/locations/{location}/instances/{instance}")

// CloudRunInstanceIdentity is the identity of a GCP CloudRunInstance resource.
// +k8s:deepcopy-gen=false
type CloudRunInstanceIdentity struct {
	Project  string
	Location string
	Instance string
}

func (i *CloudRunInstanceIdentity) String() string {
	return CloudRunInstanceIdentityFormat.ToString(*i)
}

func (i *CloudRunInstanceIdentity) FromExternal(ref string) error {
	parsed, match, err := CloudRunInstanceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CloudRunInstance external=%q was not known (use %s): %w", ref, CloudRunInstanceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CloudRunInstance external=%q was not known (use %s)", ref, CloudRunInstanceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *CloudRunInstanceIdentity) Host() string {
	return CloudRunInstanceIdentityFormat.Host()
}

func (i *CloudRunInstanceIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func getIdentityFromCloudRunInstanceSpec(ctx context.Context, reader client.Reader, obj *CloudRunInstance) (*CloudRunInstanceIdentity, error) {
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

	identity := &CloudRunInstanceIdentity{
		Project:  projectID,
		Location: location,
		Instance: resourceID,
	}
	return identity, nil
}

func (obj *CloudRunInstance) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromCloudRunInstanceSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.externalRef, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &CloudRunInstanceIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change CloudRunInstance identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
