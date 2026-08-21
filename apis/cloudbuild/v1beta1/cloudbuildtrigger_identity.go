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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &CloudBuildTriggerIdentity{}
	_ identity.Resource   = &CloudBuildTrigger{}
)

var CloudBuildTriggerIdentityFormat = gcpurls.Template[CloudBuildTriggerIdentity]("cloudbuild.googleapis.com", "projects/{project}/locations/{location}/triggers/{trigger}")

// +k8s:deepcopy-gen=false
type CloudBuildTriggerIdentity struct {
	Project  string
	Location string
	Trigger  string
}

func (i *CloudBuildTriggerIdentity) String() string {
	return CloudBuildTriggerIdentityFormat.ToString(*i)
}

func (i *CloudBuildTriggerIdentity) FromExternal(ref string) error {
	parsed, match, err := CloudBuildTriggerIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CloudBuildTrigger external=%q was not known (use %s): %w", ref, CloudBuildTriggerIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CloudBuildTrigger external=%q was not known (use %s)", ref, CloudBuildTriggerIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *CloudBuildTriggerIdentity) Host() string {
	return CloudBuildTriggerIdentityFormat.Host()
}

func getIdentityFromCloudBuildTriggerSpec(ctx context.Context, reader client.Reader, obj *CloudBuildTrigger) (*CloudBuildTriggerIdentity, error) {
	location := common.ValueOf(obj.Spec.Location)
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	// Trigger ID is server-generated. Use TriggerId from status if it exists.
	triggerID := common.ValueOf(obj.Status.TriggerId)
	if triggerID == "" {
		return nil, nil
	}

	return &CloudBuildTriggerIdentity{
		Project:  projectID,
		Location: location,
		Trigger:  triggerID,
	}, nil
}

func (obj *CloudBuildTrigger) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromCloudBuildTriggerSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Use Status.ExternalRef if it exists
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		actualIdentity := &CloudBuildTriggerIdentity{}
		if err := actualIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if specIdentity != nil {
			if actualIdentity.Project != specIdentity.Project {
				return nil, fmt.Errorf("project changed, expect %s, got %s", actualIdentity.Project, specIdentity.Project)
			}
			if actualIdentity.Location != specIdentity.Location {
				return nil, fmt.Errorf("location changed, expect %s, got %s", actualIdentity.Location, specIdentity.Location)
			}
		}

		return actualIdentity, nil
	}

	if specIdentity == nil {
		return nil, nil
	}

	return specIdentity, nil
}
