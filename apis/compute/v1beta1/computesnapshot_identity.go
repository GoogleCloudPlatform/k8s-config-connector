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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeSnapshotIdentity{}
	_ identity.Resource   = &ComputeSnapshot{}
)

var GlobalComputeSnapshotIdentityFormat = gcpurls.Template[ComputeSnapshotIdentity]("compute.googleapis.com", "projects/{project}/global/snapshots/{snapshot}")
var RegionalComputeSnapshotIdentityFormat = gcpurls.Template[ComputeSnapshotIdentity]("compute.googleapis.com", "projects/{project}/regions/{region}/snapshots/{snapshot}")

// ComputeSnapshotIdentity is the identity of a GCP ComputeSnapshot resource.
// +k8s:deepcopy-gen=false
type ComputeSnapshotIdentity struct {
	Project  string
	Region   string
	Snapshot string
}

func (i *ComputeSnapshotIdentity) IsRegional() bool {
	return i.Region != ""
}

func (i *ComputeSnapshotIdentity) String() string {
	if i.IsRegional() {
		return RegionalComputeSnapshotIdentityFormat.ToString(*i)
	}
	return GlobalComputeSnapshotIdentityFormat.ToString(*i)
}

func (i *ComputeSnapshotIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	if parsed, match, err := GlobalComputeSnapshotIdentityFormat.Parse(trimmedRef); err == nil && match {
		*i = *parsed
		return nil
	}
	if parsed, match, err := RegionalComputeSnapshotIdentityFormat.Parse(trimmedRef); err == nil && match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of ComputeSnapshot external=%q was not known (use %s or %s)", ref, GlobalComputeSnapshotIdentityFormat.CanonicalForm(), RegionalComputeSnapshotIdentityFormat.CanonicalForm())
}

func (i *ComputeSnapshotIdentity) Host() string {
	return GlobalComputeSnapshotIdentityFormat.Host()
}

func (i *ComputeSnapshotIdentity) ParentString() string {
	if i.IsRegional() {
		return fmt.Sprintf("projects/%s/regions/%s", i.Project, i.Region)
	}
	return fmt.Sprintf("projects/%s/global", i.Project)
}

func ParseComputeSnapshotExternal(external string) (*ComputeSnapshotIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeSnapshot external value")
	}
	id := &ComputeSnapshotIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeSnapshotSpec(ctx context.Context, reader client.Reader, obj *ComputeSnapshot) (*ComputeSnapshotIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ComputeSnapshotIdentity{
		Project:  projectID,
		Snapshot: resourceID,
	}
	return identity, nil
}

func (obj *ComputeSnapshot) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeSnapshotSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.selfLink, if present.
	selfLink := common.ValueOf(obj.Status.SelfLink)
	if selfLink != "" {
		statusIdentity := &ComputeSnapshotIdentity{}
		if err := statusIdentity.FromExternal(selfLink); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeSnapshot identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
