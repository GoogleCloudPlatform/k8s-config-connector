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
	_ identity.IdentityV2 = &AIPlatformScheduleIdentity{}
	_ identity.Resource   = &AIPlatformSchedule{}
)

var AIPlatformScheduleIdentityFormat = gcpurls.Template[AIPlatformScheduleIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/schedules/{schedule}")

// AIPlatformScheduleIdentity is the identity of a GCP AIPlatformSchedule resource.
// +k8s:deepcopy-gen=false
type AIPlatformScheduleIdentity struct {
	Project  string
	Location string
	Schedule string
}

func (i *AIPlatformScheduleIdentity) String() string {
	return AIPlatformScheduleIdentityFormat.ToString(*i)
}

func (i *AIPlatformScheduleIdentity) FromExternal(ref string) error {
	parsed, match, err := AIPlatformScheduleIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of AIPlatformSchedule external=%q was not known (use %s): %w", ref, AIPlatformScheduleIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of AIPlatformSchedule external=%q was not known (use %s)", ref, AIPlatformScheduleIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *AIPlatformScheduleIdentity) Host() string {
	return AIPlatformScheduleIdentityFormat.Host()
}

func (i *AIPlatformScheduleIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func getIdentityFromAIPlatformScheduleSpec(ctx context.Context, reader client.Reader, obj *AIPlatformSchedule) (*AIPlatformScheduleIdentity, error) {
	// AIPlatformSchedule only supports service-generated IDs.
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &AIPlatformScheduleIdentity{
		Project:  projectID,
		Location: location,
		Schedule: resourceID,
	}
	return identity, nil
}

func (obj *AIPlatformSchedule) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromAIPlatformScheduleSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &AIPlatformScheduleIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		isNumeric := func(s string) bool {
			for _, r := range s {
				if r < '0' || r > '9' {
					return false
				}
			}
			return len(s) > 0
		}

		projectMatch := statusIdentity.Project == specIdentity.Project || isNumeric(statusIdentity.Project)
		if specIdentity.Schedule == "" {
			if !projectMatch || statusIdentity.Location != specIdentity.Location {
				return nil, fmt.Errorf("cannot change AIPlatformSchedule parent (old project/location=%s/%s, new project/location=%s/%s)",
					statusIdentity.Project, statusIdentity.Location, specIdentity.Project, specIdentity.Location)
			}
			return statusIdentity, nil
		}

		if statusIdentity.Schedule != specIdentity.Schedule || !projectMatch || statusIdentity.Location != specIdentity.Location {
			return nil, fmt.Errorf("cannot change AIPlatformSchedule identity (old=%q, new=%q)",
				statusIdentity.String(), specIdentity.String())
		}
		return statusIdentity, nil
	}

	return specIdentity, nil
}
