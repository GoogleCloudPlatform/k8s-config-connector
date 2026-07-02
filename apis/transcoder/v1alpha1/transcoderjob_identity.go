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
	_ identity.IdentityV2 = &TranscoderJobIdentity{}
	_ identity.Resource   = &TranscoderJob{}
)

var TranscoderJobIdentityFormat = gcpurls.Template[TranscoderJobIdentity]("transcoder.googleapis.com", "projects/{project}/locations/{location}/jobs/{job}")

// TranscoderJobIdentity is the identity of a GCP TranscoderJob resource.
// +k8s:deepcopy-gen=false
type TranscoderJobIdentity struct {
	Project  string
	Location string
	Job      string
}

func (i *TranscoderJobIdentity) String() string {
	return TranscoderJobIdentityFormat.ToString(*i)
}

func (i *TranscoderJobIdentity) FromExternal(ref string) error {
	parsed, match, err := TranscoderJobIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of TranscoderJob external=%q was not known (use %s): %w", ref, TranscoderJobIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of TranscoderJob external=%q was not known (use %s)", ref, TranscoderJobIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *TranscoderJobIdentity) Host() string {
	return TranscoderJobIdentityFormat.Host()
}

func (i *TranscoderJobIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func getIdentityFromTranscoderJobSpec(ctx context.Context, reader client.Reader, obj *TranscoderJob) (*TranscoderJobIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("spec.location must be specified")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &TranscoderJobIdentity{
		Project:  projectID,
		Location: location,
		Job:      resourceID,
	}
	return identity, nil
}

func (obj *TranscoderJob) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromTranscoderJobSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &TranscoderJobIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change TranscoderJob identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
