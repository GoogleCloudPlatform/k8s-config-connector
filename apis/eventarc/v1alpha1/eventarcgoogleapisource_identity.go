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
	_ identity.IdentityV2 = &EventarcGoogleApiSourceIdentity{}
	_ identity.Resource   = &EventarcGoogleApiSource{}
)

var EventarcGoogleApiSourceIdentityFormat = gcpurls.Template[EventarcGoogleApiSourceIdentity]("eventarc.googleapis.com", "projects/{project}/locations/{location}/googleApiSources/{google_api_source}")

// +k8s:deepcopy-gen=false
type EventarcGoogleApiSourceIdentity struct {
	Project           string
	Location          string
	Google_api_source string
}

func (i *EventarcGoogleApiSourceIdentity) String() string {
	return EventarcGoogleApiSourceIdentityFormat.ToString(*i)
}

func (i *EventarcGoogleApiSourceIdentity) FromExternal(ref string) error {
	parsed, match, err := EventarcGoogleApiSourceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of EventarcGoogleApiSource external=%q was not known (use %s): %w", ref, EventarcGoogleApiSourceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of EventarcGoogleApiSource external=%q was not known (use %s)", ref, EventarcGoogleApiSourceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *EventarcGoogleApiSourceIdentity) Host() string {
	return EventarcGoogleApiSourceIdentityFormat.Host()
}

func (i *EventarcGoogleApiSourceIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func NewEventarcGoogleApiSourceIdentity(ctx context.Context, reader client.Reader, obj client.Object) (*EventarcGoogleApiSourceIdentity, error) {
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

	identity := &EventarcGoogleApiSourceIdentity{
		Project:           projectID,
		Location:          location,
		Google_api_source: resourceID,
	}
	return identity, nil
}

func (obj *EventarcGoogleApiSource) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := NewEventarcGoogleApiSourceIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &EventarcGoogleApiSourceIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change EventarcGoogleApiSource identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

func (obj *EventarcGoogleApiSource) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
