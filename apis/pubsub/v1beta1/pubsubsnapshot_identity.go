// Copyright 2025 Google LLC
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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &PubSubSnapshotIdentity{}
	_ identity.Resource   = &PubSubSnapshot{}
)

var PubSubSnapshotIdentityFormat = gcpurls.Template[PubSubSnapshotIdentity]("pubsub.googleapis.com", "projects/{project}/snapshots/{snapshot}")

// +k8s:deepcopy-gen=false

// PubSubSnapshotIdentity is the identity of a GCP PubSubSnapshot resource.
type PubSubSnapshotIdentity struct {
	Project  string
	Snapshot string
}

func (i *PubSubSnapshotIdentity) String() string {
	return PubSubSnapshotIdentityFormat.ToString(*i)
}

func (i *PubSubSnapshotIdentity) FromExternal(ref string) error {
	parsed, match, err := PubSubSnapshotIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of PubSubSnapshot external=%q was not known (use %s): %w", ref, PubSubSnapshotIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of PubSubSnapshot external=%q was not known (use %s)", ref, PubSubSnapshotIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *PubSubSnapshotIdentity) Host() string {
	return PubSubSnapshotIdentityFormat.Host()
}

func (i *PubSubSnapshotIdentity) ParentString() string {
	return "projects/" + i.Project
}

func getIdentityFromPubSubSnapshotSpec(ctx context.Context, reader client.Reader, obj *PubSubSnapshot) (*PubSubSnapshotIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &PubSubSnapshotIdentity{
		Project:  projectID,
		Snapshot: resourceID,
	}
	return identity, nil
}

func (obj *PubSubSnapshot) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromPubSubSnapshotSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	if obj.Status.ExternalRef != nil && *obj.Status.ExternalRef != "" {
		actualId := &PubSubSnapshotIdentity{}
		if err := actualId.FromExternal(*obj.Status.ExternalRef); err != nil {
			return nil, err
		}
		if actualId.Project != specIdentity.Project {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualId.Project, specIdentity.Project)
		}
		if actualId.Snapshot != specIdentity.Snapshot {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				specIdentity.Snapshot, actualId.Snapshot)
		}
	}

	return specIdentity, nil
}
