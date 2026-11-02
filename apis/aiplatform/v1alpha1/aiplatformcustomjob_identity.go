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
	"strconv"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &AIPlatformCustomJobIdentity{}
	_ identity.Resource   = &AIPlatformCustomJob{}
)

var AIPlatformCustomJobIdentityFormat = gcpurls.Template[AIPlatformCustomJobIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/customJobs/{customJob}")

// AIPlatformCustomJobIdentity is the identity of a GCP AIPlatformCustomJob resource.
// +k8s:deepcopy-gen=false
type AIPlatformCustomJobIdentity struct {
	Project   string
	Location  string
	CustomJob string
}

func (i *AIPlatformCustomJobIdentity) String() string {
	return AIPlatformCustomJobIdentityFormat.ToString(*i)
}

func (i *AIPlatformCustomJobIdentity) FromExternal(ref string) error {
	parsed, match, err := AIPlatformCustomJobIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of AIPlatformCustomJob external=%q was not known (use %s): %w", ref, AIPlatformCustomJobIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of AIPlatformCustomJob external=%q was not known (use %s)", ref, AIPlatformCustomJobIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *AIPlatformCustomJobIdentity) Host() string {
	return AIPlatformCustomJobIdentityFormat.Host()
}

func (i *AIPlatformCustomJobIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func getIdentityFromAIPlatformCustomJobSpec(ctx context.Context, reader client.Reader, obj *AIPlatformCustomJob) (*AIPlatformCustomJobIdentity, error) {
	// AIPlatformCustomJob only supports service-generated IDs (GCP assigns an ID upon creation).
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &AIPlatformCustomJobIdentity{
		Project:   projectID,
		Location:  location,
		CustomJob: resourceID,
	}
	return identity, nil
}

func (obj *AIPlatformCustomJob) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromAIPlatformCustomJobSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &AIPlatformCustomJobIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if !IsProjectIDMatch(statusIdentity.Project, specIdentity.Project) {
			return nil, fmt.Errorf("cannot change AIPlatformCustomJob project (old=%q, new=%q)", statusIdentity.Project, specIdentity.Project)
		}
		if statusIdentity.Location != specIdentity.Location {
			return nil, fmt.Errorf("cannot change AIPlatformCustomJob location (old=%q, new=%q)", statusIdentity.Location, specIdentity.Location)
		}

		if specIdentity.CustomJob != "" {
			if statusIdentity.CustomJob != specIdentity.CustomJob {
				return nil, fmt.Errorf("cannot change AIPlatformCustomJob identity (old=%q, new=%q)", statusIdentity.CustomJob, specIdentity.CustomJob)
			}
		}

		return statusIdentity, nil
	}

	return specIdentity, nil
}

// IsProjectIDMatch returns true if two project identifiers are considered matching,
// or if we can't reliably compare them because one is an alphanumeric project ID
// and the other is a numeric project number.
func IsProjectIDMatch(p1, p2 string) bool {
	if p1 == p2 {
		return true
	}
	if p1 == "" || p2 == "" {
		return false
	}
	p1IsNumeric := isNumeric(p1)
	p2IsNumeric := isNumeric(p2)
	if p1IsNumeric != p2IsNumeric {
		// Skip strict matching if one is project ID (alphanumeric) and the other is project number (numeric).
		return true
	}
	return p1 == p2
}

func isNumeric(s string) bool {
	_, err := strconv.ParseUint(s, 10, 64)
	return err == nil
}
