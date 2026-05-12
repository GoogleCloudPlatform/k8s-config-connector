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
	_ identity.IdentityV2 = &EventarcEnrollmentIdentity{}
	_ identity.Resource   = &EventarcEnrollment{}
)

var EventarcEnrollmentIdentityFormat = gcpurls.Template[EventarcEnrollmentIdentity]("eventarc.googleapis.com", "projects/{project}/locations/{location}/enrollments/{enrollment}")

// +k8s:deepcopy-gen=false
type EventarcEnrollmentIdentity struct {
	Project    string
	Location   string
	Enrollment string
}

func (i *EventarcEnrollmentIdentity) String() string {
	return EventarcEnrollmentIdentityFormat.ToString(*i)
}

func (i *EventarcEnrollmentIdentity) FromExternal(ref string) error {
	parsed, match, err := EventarcEnrollmentIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of EventarcEnrollment external=%q was not known (use %s): %w", ref, EventarcEnrollmentIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of EventarcEnrollment external=%q was not known (use %s)", ref, EventarcEnrollmentIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *EventarcEnrollmentIdentity) Host() string {
	return EventarcEnrollmentIdentityFormat.Host()
}

func NewEventarcEnrollmentIdentity(ctx context.Context, reader client.Reader, obj client.Object) (*EventarcEnrollmentIdentity, error) {
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

	identity := &EventarcEnrollmentIdentity{
		Project:    projectID,
		Location:   location,
		Enrollment: resourceID,
	}
	return identity, nil
}

func (obj *EventarcEnrollment) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := NewEventarcEnrollmentIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &EventarcEnrollmentIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change EventarcEnrollment identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
