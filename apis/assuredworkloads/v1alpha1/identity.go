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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type WorkloadIdentity struct {
	OrganizationID string
	Location       string
	WorkloadID     string
}

var _ identity.Identity = &WorkloadIdentity{}

var WorkloadFormat = gcpurls.Template[WorkloadIdentity]("assuredworkloads.googleapis.com", "organizations/{organizationID}/locations/{location}/workloads/{workloadID}")

func (i *WorkloadIdentity) Host() string {
	return WorkloadFormat.Host()
}

func (i *WorkloadIdentity) String() string {
	return WorkloadFormat.ToString(*i)
}

func (i *WorkloadIdentity) FromExternal(ref string) error {
	parsed, match, err := WorkloadFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of AssuredWorkloadsWorkload external=%q was not known (use %s): %w", ref, WorkloadFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of AssuredWorkloadsWorkload external=%q was not known (use %s)", ref, WorkloadFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func NewWorkloadIdentity(ctx context.Context, reader client.Reader, obj *AssuredWorkloadsWorkload) (*WorkloadIdentity, error) {
	// Assumes references have been normalized
	organizationRef, err := refs.ResolveOrganization(ctx, reader, obj, obj.Spec.OrganizationRef)
	if err != nil {
		return nil, err
	}
	organizationID := organizationRef.OrganizationID

	location := obj.Spec.Location
	workloadID := ""
	if obj.Spec.ResourceID != nil {
		workloadID = *obj.Spec.ResourceID
	} else {
		workloadID = obj.Name
	}

	return &WorkloadIdentity{
		OrganizationID: organizationID,
		Location:       location,
		WorkloadID:     workloadID,
	}, nil
}

func (i *WorkloadIdentity) Parent() *parent.OrganizationLocationParent {
	return &parent.OrganizationLocationParent{
		OrganizationID: i.OrganizationID,
		Location:       i.Location,
	}
}
