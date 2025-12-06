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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &MigrationJobIdentity{}

const (
	MigrationJobIDURL = parent.ProjectAndLocationURL + "/migrationJobs/{{migrationJobID}}"
)

// MigrationJobIdentity defines the resource reference to CloudDMSMigrationJob, which "External" field
// holds the GCP identifier for the KRM object.
type MigrationJobIdentity struct {
	parent *parent.ProjectAndLocationParent
	id     string
}

func (i *MigrationJobIdentity) String() string {
	return i.parent.String() + "/migrationJobs/" + i.id
}

func (i *MigrationJobIdentity) ID() string {
	return i.id
}

func (i *MigrationJobIdentity) Parent() *parent.ProjectAndLocationParent {
	return i.parent
}

func (i *MigrationJobIdentity) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/migrationJobs/")
	if len(tokens) != 2 {
		return fmt.Errorf("format of CloudDMSMigrationJob external=%q was not known (use projects/{{projectID}}/locations/{{location}}/migrationJobs/{{migrationJobID}})", ref)
	}
	i.parent = &parent.ProjectAndLocationParent{}
	if err := i.parent.FromExternal(tokens[0]); err != nil {
		return err
	}
	i.id = tokens[1]
	if i.id == "" {
		return fmt.Errorf("migrationJobID was empty in external=%q", ref)
	}
	return nil
}

var _ identity.Resource = &CloudDMSMigrationJob{}

func (obj *CloudDMSMigrationJob) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	migrationJob := &MigrationJobIdentity{}
	migrationJob.parent = &parent.ProjectAndLocationParent{} // Initialize the parent field

	// Resolve user-configured Parent
	if err := obj.Spec.ProjectAndLocationRef.Build(ctx, reader, obj.GetNamespace(), migrationJob.parent); err != nil {
		return nil, err
	}

	// Get user-configured ID
	migrationJob.id = common.ValueOf(obj.Spec.ResourceID)
	if migrationJob.id == "" {
		migrationJob.id = obj.GetName()
	}
	if migrationJob.id == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Validate against the ID stored in status.externalRef, if any
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &MigrationJobIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != migrationJob.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, migrationJob.String())
		}
	}
	return migrationJob, nil
}
