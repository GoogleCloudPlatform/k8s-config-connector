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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.IdentityV2 = &FirebaseHostingSiteIdentity{}
var _ identity.Resource = &FirebaseHostingSite{}

var FirebaseHostingSiteIdentityFormat = gcpurls.Template[FirebaseHostingSiteIdentity](
	"firebasehosting.googleapis.com",
	"projects/{project}/sites/{site}",
)

// +k8s:deepcopy-gen=false
// FirebaseHostingSiteIdentity is the identity of a GCP FirebaseHostingSite resource.
type FirebaseHostingSiteIdentity struct {
	Project string
	Site    string
}

func (i *FirebaseHostingSiteIdentity) String() string {
	return FirebaseHostingSiteIdentityFormat.ToString(*i)
}

func (i *FirebaseHostingSiteIdentity) FromExternal(ref string) error {
	parsed, match, err := FirebaseHostingSiteIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of FirebaseHostingSite external=%q was not known (use %s): %w", ref, FirebaseHostingSiteIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of FirebaseHostingSite external=%q was not known (use %s)", ref, FirebaseHostingSiteIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *FirebaseHostingSiteIdentity) Host() string {
	return FirebaseHostingSiteIdentityFormat.Host()
}

func (i *FirebaseHostingSiteIdentity) ParentString() string {
	return "projects/" + i.Project
}

func getIdentityFromFirebaseHostingSiteSpec(ctx context.Context, reader client.Reader, obj *FirebaseHostingSite) (*FirebaseHostingSiteIdentity, error) {
	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	siteID := common.ValueOf(obj.Spec.ResourceID)
	if siteID == "" {
		siteID = obj.GetName()
	}

	return &FirebaseHostingSiteIdentity{
		Project: projectID,
		Site:    siteID,
	}, nil
}

func (obj *FirebaseHostingSite) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromFirebaseHostingSiteSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.Name)
	if externalRef != "" {
		if !strings.Contains(externalRef, "/") {
			if externalRef != specIdentity.Site {
				return nil, fmt.Errorf("cannot change FirebaseHostingSite identity (old=%q, new=%q)", externalRef, specIdentity.Site)
			}
		} else {
			statusIdentity := &FirebaseHostingSiteIdentity{}
			if err := statusIdentity.FromExternal(externalRef); err != nil {
				return nil, err
			}

			if statusIdentity.Project == specIdentity.Project {
				if statusIdentity.String() != specIdentity.String() {
					return nil, fmt.Errorf("cannot change FirebaseHostingSite identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
				}
			} else {
				if statusIdentity.Site != specIdentity.Site {
					return nil, fmt.Errorf("cannot change FirebaseHostingSite site (old=%q, new=%q)", statusIdentity.Site, specIdentity.Site)
				}
			}
		}
	}

	return specIdentity, nil
}
