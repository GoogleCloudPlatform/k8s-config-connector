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
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

var (
	_ identity.IdentityV2 = &ComputeSnapshotIdentity{}
)

var (
	ComputeSnapshotGlobalIdentityFormat   = gcpurls.Template[ComputeSnapshotIdentity]("compute.googleapis.com", "projects/{project}/global/snapshots/{snapshot}")
	ComputeSnapshotRegionalIdentityFormat = gcpurls.Template[ComputeSnapshotIdentity]("compute.googleapis.com", "projects/{project}/regions/{location}/snapshots/{snapshot}")
)

// ComputeSnapshotIdentity is the identity of a GCP ComputeSnapshot resource.
// +k8s:deepcopy-gen=false
type ComputeSnapshotIdentity struct {
	Project  string
	Location string
	Snapshot string
}

func (i *ComputeSnapshotIdentity) IsRegional() bool {
	return i.Location != "" && i.Location != "global"
}

func (i *ComputeSnapshotIdentity) String() string {
	if i.IsRegional() {
		return ComputeSnapshotRegionalIdentityFormat.ToString(*i)
	}
	return ComputeSnapshotGlobalIdentityFormat.ToString(*i)
}

func (i *ComputeSnapshotIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeSnapshotGlobalIdentityFormat.Parse(trimmedRef)
	if err == nil && match {
		*i = *parsed
		i.Location = "global"
		return nil
	}

	parsed, match, err = ComputeSnapshotRegionalIdentityFormat.Parse(trimmedRef)
	if err == nil && match {
		*i = *parsed
		return nil
	}

	return fmt.Errorf("format of ComputeSnapshot external=%q was not known (expected global %q or regional %q)", ref, ComputeSnapshotGlobalIdentityFormat.CanonicalForm(), ComputeSnapshotRegionalIdentityFormat.CanonicalForm())
}

func (i *ComputeSnapshotIdentity) Host() string {
	return ComputeSnapshotGlobalIdentityFormat.Host()
}

func (i *ComputeSnapshotIdentity) ParentString() string {
	if i.IsRegional() {
		return fmt.Sprintf("projects/%s/regions/%s", i.Project, i.Location)
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
