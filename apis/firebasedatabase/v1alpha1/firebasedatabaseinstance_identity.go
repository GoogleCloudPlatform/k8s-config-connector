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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.IdentityV2 = &FirebaseDatabaseInstanceIdentity{}
var _ identity.Resource = &FirebaseDatabaseInstance{}

var FirebaseDatabaseInstanceIdentityFormat = gcpurls.Template[FirebaseDatabaseInstanceIdentity](
	"firebasedatabase.googleapis.com",
	"projects/{project}/locations/{location}/instances/{instance}",
)

// +k8s:deepcopy-gen=false
// FirebaseDatabaseInstanceIdentity is the identity of a GCP FirebaseDatabaseInstance resource.
type FirebaseDatabaseInstanceIdentity struct {
	Project  string
	Location string
	Instance string
}

func (i *FirebaseDatabaseInstanceIdentity) String() string {
	return FirebaseDatabaseInstanceIdentityFormat.ToString(*i)
}

func (i *FirebaseDatabaseInstanceIdentity) FromExternal(ref string) error {
	parsed, match, err := FirebaseDatabaseInstanceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of FirebaseDatabaseInstance external=%q was not known (use %s): %w", ref, FirebaseDatabaseInstanceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of FirebaseDatabaseInstance external=%q was not known (use %s)", ref, FirebaseDatabaseInstanceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *FirebaseDatabaseInstanceIdentity) Host() string {
	return FirebaseDatabaseInstanceIdentityFormat.Host()
}

func (i *FirebaseDatabaseInstanceIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func getIdentityFromFirebaseDatabaseInstanceSpec(ctx context.Context, reader client.Reader, obj *FirebaseDatabaseInstance) (*FirebaseDatabaseInstanceIdentity, error) {
	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	location := obj.Spec.Region
	if location == "" {
		return nil, fmt.Errorf("spec.region is required")
	}

	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	return &FirebaseDatabaseInstanceIdentity{
		Project:  projectID,
		Location: location,
		Instance: resourceID,
	}, nil
}

func (obj *FirebaseDatabaseInstance) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		externalRef = common.ValueOf(obj.Status.Name)
	}
	if externalRef != "" {
		id := &FirebaseDatabaseInstanceIdentity{}
		if err := id.FromExternal(externalRef); err != nil {
			return nil, err
		}
		return id, nil
	}

	return getIdentityFromFirebaseDatabaseInstanceSpec(ctx, reader, obj)
}
