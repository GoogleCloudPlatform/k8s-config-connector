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
	_ identity.IdentityV2 = &ParallelstoreInstanceIdentity{}
	_ identity.Resource   = &ParallelstoreInstance{}
)

var ParallelstoreInstanceIdentityFormat = gcpurls.Template[ParallelstoreInstanceIdentity]("parallelstore.googleapis.com", "projects/{project}/locations/{location}/instances/{instance}")

// ParallelstoreInstanceIdentity is the identity of a GCP ParallelstoreInstance resource.
// +k8s:deepcopy-gen=false
type ParallelstoreInstanceIdentity struct {
	Project  string
	Location string
	Instance string
}

func (i *ParallelstoreInstanceIdentity) String() string {
	return ParallelstoreInstanceIdentityFormat.ToString(*i)
}

func (i *ParallelstoreInstanceIdentity) FromExternal(ref string) error {
	parsed, match, err := ParallelstoreInstanceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ParallelstoreInstance external=%q was not known (use %s): %w", ref, ParallelstoreInstanceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ParallelstoreInstance external=%q was not known (use %s)", ref, ParallelstoreInstanceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ParallelstoreInstanceIdentity) Host() string {
	return ParallelstoreInstanceIdentityFormat.Host()
}

func (i *ParallelstoreInstanceIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func getIdentityFromParallelstoreInstanceSpec(ctx context.Context, reader client.Reader, obj *ParallelstoreInstance) (*ParallelstoreInstanceIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &ParallelstoreInstanceIdentity{
		Project:  projectID,
		Location: location,
		Instance: resourceID,
	}
	return identity, nil
}

func (obj *ParallelstoreInstance) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromParallelstoreInstanceSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Validate against the ID stored in status.externalRef, if any
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &ParallelstoreInstanceIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, specIdentity.String())
		}
	}
	return specIdentity, nil
}

func (obj *ParallelstoreInstance) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
