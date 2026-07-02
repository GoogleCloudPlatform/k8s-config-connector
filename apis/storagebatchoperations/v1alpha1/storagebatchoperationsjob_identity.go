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
	_ identity.IdentityV2 = &StorageBatchOperationsJobIdentity{}
	_ identity.Resource   = &StorageBatchOperationsJob{}
)

var StorageBatchOperationsJobIdentityFormat = gcpurls.Template[StorageBatchOperationsJobIdentity]("storagebatchoperations.googleapis.com", "projects/{project}/locations/{location}/jobs/{job}")

// +k8s:deepcopy-gen=false
type StorageBatchOperationsJobIdentity struct {
	Project  string
	Location string
	Job      string
}

func (i *StorageBatchOperationsJobIdentity) String() string {
	return StorageBatchOperationsJobIdentityFormat.ToString(*i)
}

func (i *StorageBatchOperationsJobIdentity) FromExternal(ref string) error {
	parsed, match, err := StorageBatchOperationsJobIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of StorageBatchOperationsJob external=%q was not known (use %s): %w", ref, StorageBatchOperationsJobIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of StorageBatchOperationsJob external=%q was not known (use %s)", ref, StorageBatchOperationsJobIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *StorageBatchOperationsJobIdentity) Host() string {
	return StorageBatchOperationsJobIdentityFormat.Host()
}

func getIdentityFromStorageBatchOperationsJobSpec(ctx context.Context, reader client.Reader, obj client.Object) (*StorageBatchOperationsJobIdentity, error) {
	jobObj, ok := obj.(*StorageBatchOperationsJob)
	if !ok {
		return nil, fmt.Errorf("object is not a StorageBatchOperationsJob")
	}
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location := jobObj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &StorageBatchOperationsJobIdentity{
		Project:  projectID,
		Location: location,
		Job:      resourceID,
	}
	return identity, nil
}

func (obj *StorageBatchOperationsJob) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromStorageBatchOperationsJobSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &StorageBatchOperationsJobIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change StorageBatchOperationsJob identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
