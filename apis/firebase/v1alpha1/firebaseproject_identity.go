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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.IdentityV2 = &FirebaseProjectIdentity{}
var _ identity.Resource = &FirebaseProject{}

var FirebaseProjectIdentityFormat = gcpurls.Template[FirebaseProjectIdentity](
	"firebase.googleapis.com",
	"projects/{project}",
)

// +k8s:deepcopy-gen=false
// FirebaseProjectIdentity is the identity of a GCP FirebaseProject resource.
type FirebaseProjectIdentity struct {
	Project string
}

func (i *FirebaseProjectIdentity) String() string {
	return FirebaseProjectIdentityFormat.ToString(*i)
}

func (i *FirebaseProjectIdentity) FromExternal(ref string) error {
	parsed, match, err := FirebaseProjectIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of FirebaseProject external=%q was not known (use %s): %w", ref, FirebaseProjectIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of FirebaseProject external=%q was not known (use %s)", ref, FirebaseProjectIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *FirebaseProjectIdentity) Host() string {
	return FirebaseProjectIdentityFormat.Host()
}

func (i *FirebaseProjectIdentity) ParentString() string {
	return ""
}

func getIdentityFromFirebaseProjectSpec(ctx context.Context, reader client.Reader, obj *FirebaseProject) (*FirebaseProjectIdentity, error) {
	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return &FirebaseProjectIdentity{
		Project: projectID,
	}, nil
}

func (obj *FirebaseProject) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromFirebaseProjectSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.projectNumber if present and formatted.
	// FirebaseProject does not have status.externalRef or status.name in the baseline CRD.
	// If status.projectNumber is populated, check if it contradicts (if matched as project).
	if obj.Status.ProjectNumber != nil {
		pn := *obj.Status.ProjectNumber
		if pn != "" && strings.Contains(pn, "/") {
			statusIdentity := &FirebaseProjectIdentity{}
			if err := statusIdentity.FromExternal(pn); err == nil {
				if statusIdentity.Project != specIdentity.Project {
					return nil, fmt.Errorf("cannot change FirebaseProject identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
				}
			}
		}
	}

	return specIdentity, nil
}
