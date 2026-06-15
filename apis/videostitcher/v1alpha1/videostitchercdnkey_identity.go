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
	_ identity.IdentityV2 = &VideoStitcherCdnKeyIdentity{}
	_ identity.Resource   = &VideoStitcherCdnKey{}
)

// VideoStitcherCdnKeyIdentity represents the GCP identity of a VideoStitcherCdnKey.
var VideoStitcherCdnKeyIdentityFormat = gcpurls.Template[VideoStitcherCdnKeyIdentity]("videostitcher.googleapis.com", "projects/{project}/locations/{location}/cdnKeys/{cdnKey}")

// VideoStitcherCdnKeyIdentity is the identity of a GCP VideoStitcherCdnKey resource.
// +k8s:deepcopy-gen=false
type VideoStitcherCdnKeyIdentity struct {
	Project  string
	Location string
	CdnKey   string
}

func (i *VideoStitcherCdnKeyIdentity) String() string {
	return VideoStitcherCdnKeyIdentityFormat.ToString(*i)
}

func (i *VideoStitcherCdnKeyIdentity) FromExternal(ref string) error {
	parsed, match, err := VideoStitcherCdnKeyIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VideoStitcherCdnKey external=%q was not known (use %s): %w", ref, VideoStitcherCdnKeyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VideoStitcherCdnKey external=%q was not known (use %s)", ref, VideoStitcherCdnKeyIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *VideoStitcherCdnKeyIdentity) Host() string {
	return VideoStitcherCdnKeyIdentityFormat.Host()
}

// ParentString returns the parent GCP resource path.
func (i *VideoStitcherCdnKeyIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func getIdentityFromVideoStitcherCdnKeySpec(ctx context.Context, reader client.Reader, obj *VideoStitcherCdnKey) (*VideoStitcherCdnKeyIdentity, error) {
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

	identity := &VideoStitcherCdnKeyIdentity{
		Project:  projectID,
		Location: location,
		CdnKey:   resourceID,
	}
	return identity, nil
}

func (obj *VideoStitcherCdnKey) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromVideoStitcherCdnKeySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &VideoStitcherCdnKeyIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change VideoStitcherCdnKey identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
