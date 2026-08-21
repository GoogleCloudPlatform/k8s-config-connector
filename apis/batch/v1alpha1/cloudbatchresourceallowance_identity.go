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
	_ identity.IdentityV2 = &ResourceAllowanceIdentity{}
	_ identity.Resource   = &CloudBatchResourceAllowance{}
)

var ResourceAllowanceIdentityFormat = gcpurls.Template[ResourceAllowanceIdentity]("batch.googleapis.com", "projects/{project}/locations/{location}/resourceAllowances/{resourceAllowance}")

// +k8s:deepcopy-gen=false
type ResourceAllowanceIdentity struct {
	Project           string
	Location          string
	ResourceAllowance string
}

func (i *ResourceAllowanceIdentity) String() string {
	return ResourceAllowanceIdentityFormat.ToString(*i)
}

func (i *ResourceAllowanceIdentity) FromExternal(ref string) error {
	parsed, match, err := ResourceAllowanceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CloudBatchResourceAllowance external=%q was not known (use %s): %w", ref, ResourceAllowanceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CloudBatchResourceAllowance external=%q was not known (use %s)", ref, ResourceAllowanceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ResourceAllowanceIdentity) Host() string {
	return ResourceAllowanceIdentityFormat.Host()
}

func getIdentityFromResourceAllowanceSpec(ctx context.Context, reader client.Reader, obj client.Object) (*ResourceAllowanceIdentity, error) {
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

	identity := &ResourceAllowanceIdentity{
		Project:           projectID,
		Location:          location,
		ResourceAllowance: resourceID,
	}

	return identity, nil
}

func (obj *CloudBatchResourceAllowance) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromResourceAllowanceSpec(ctx, reader, obj)
}

func (obj *CloudBatchResourceAllowance) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
