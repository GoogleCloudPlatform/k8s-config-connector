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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// CustomClassIdentity is the identity of a SpeechCustomClass.
var (
	_ identity.IdentityV2 = &CustomClassIdentity{}
	_ identity.Resource   = &SpeechCustomClass{}
)

var CustomClassIdentityFormat = gcpurls.Template[CustomClassIdentity]("speech.googleapis.com", "projects/{project}/locations/{location}/customClasses/{customclass}")

// +k8s:deepcopy-gen=false
type CustomClassIdentity struct {
	Project     string
	Location    string
	CustomClass string
}

func (i *CustomClassIdentity) String() string {
	return CustomClassIdentityFormat.ToString(*i)
}

func (i *CustomClassIdentity) FromExternal(ref string) error {
	parsed, match, err := CustomClassIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of SpeechCustomClass external=%q was not known (use %s): %w", ref, CustomClassIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of SpeechCustomClass external=%q was not known (use %s)", ref, CustomClassIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *CustomClassIdentity) Host() string {
	return CustomClassIdentityFormat.Host()
}

func (i *CustomClassIdentity) ExternalIdentifier() *string {
	return &i.CustomClass
}

func getIdentityFromSpeechCustomClassSpec(ctx context.Context, reader client.Reader, obj client.Object) (*CustomClassIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &CustomClassIdentity{
		Project:     projectID,
		Location:    location,
		CustomClass: resourceID,
	}
	return identity, nil
}

func (obj *SpeechCustomClass) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromSpeechCustomClassSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &CustomClassIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change SpeechCustomClass identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
