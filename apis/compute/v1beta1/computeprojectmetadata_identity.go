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
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeProjectMetadataIdentity{}
	_ identity.Resource   = &ComputeProjectMetadata{}
)

var ComputeProjectMetadataIdentityFormat = gcpurls.Template[ComputeProjectMetadataIdentity]("compute.googleapis.com", "projects/{project}")

// ComputeProjectMetadataIdentity is the identity of a GCP ComputeProjectMetadata resource.
// +k8s:deepcopy-gen=false
type ComputeProjectMetadataIdentity struct {
	Project string
}

func (i *ComputeProjectMetadataIdentity) String() string {
	return ComputeProjectMetadataIdentityFormat.ToString(*i)
}

func (i *ComputeProjectMetadataIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	if trimmedRef != "" && !strings.Contains(trimmedRef, "/") {
		trimmedRef = "projects/" + trimmedRef
	}

	parsed, match, err := ComputeProjectMetadataIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeProjectMetadata external=%q was not known (use %s): %w", ref, ComputeProjectMetadataIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeProjectMetadata external=%q was not known (use %s)", ref, ComputeProjectMetadataIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeProjectMetadataIdentity) Host() string {
	return ComputeProjectMetadataIdentityFormat.Host()
}

func (i *ComputeProjectMetadataIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s", i.Project)
}

func ParseComputeProjectMetadataExternal(external string) (*ComputeProjectMetadataIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeProjectMetadata external value")
	}
	id := &ComputeProjectMetadataIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeProjectMetadataSpec(ctx context.Context, reader client.Reader, obj *ComputeProjectMetadata) (*ComputeProjectMetadataIdentity, error) {
	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &ComputeProjectMetadataIdentity{
		Project: projectID,
	}
	return identity, nil
}

func (obj *ComputeProjectMetadata) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromComputeProjectMetadataSpec(ctx, reader, obj)
}
