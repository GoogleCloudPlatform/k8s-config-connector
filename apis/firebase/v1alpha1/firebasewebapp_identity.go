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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.IdentityV2 = &FirebaseWebAppIdentity{}
var _ identity.Resource = &FirebaseWebApp{}

var FirebaseWebAppIdentityFormat = gcpurls.Template[FirebaseWebAppIdentity](
	"firebase.googleapis.com",
	"projects/{project}/webApps/{app}",
)

// +k8s:deepcopy-gen=false
// FirebaseWebAppIdentity is the identity of a GCP FirebaseWebApp resource.
type FirebaseWebAppIdentity struct {
	Project string
	App     string
}

func (i *FirebaseWebAppIdentity) String() string {
	return FirebaseWebAppIdentityFormat.ToString(*i)
}

func (i *FirebaseWebAppIdentity) FromExternal(ref string) error {
	parsed, match, err := FirebaseWebAppIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of FirebaseWebApp external=%q was not known (use %s): %w", ref, FirebaseWebAppIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of FirebaseWebApp external=%q was not known (use %s)", ref, FirebaseWebAppIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *FirebaseWebAppIdentity) Host() string {
	return FirebaseWebAppIdentityFormat.Host()
}

func (i *FirebaseWebAppIdentity) ParentString() string {
	return "projects/" + i.Project
}

func getIdentityFromFirebaseWebAppSpec(ctx context.Context, reader client.Reader, obj *FirebaseWebApp) (*FirebaseWebAppIdentity, error) {
	projectID := common.ValueOf(obj.Spec.Project)
	if projectID == "" {
		var err error
		projectID, err = refsv1beta1.ResolveProjectID(ctx, reader, obj)
		if err != nil {
			return nil, err
		}
	}

	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		return nil, nil
	}

	return &FirebaseWebAppIdentity{
		Project: projectID,
		App:     resourceID,
	}, nil
}

func (obj *FirebaseWebApp) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		externalRef = common.ValueOf(obj.Status.Name)
	}
	if externalRef != "" {
		id := &FirebaseWebAppIdentity{}
		if err := id.FromExternal(externalRef); err != nil {
			return nil, err
		}
		return id, nil
	}

	specIdentity, err := getIdentityFromFirebaseWebAppSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	if specIdentity != nil {
		return specIdentity, nil
	}

	if obj.Status.AppId != nil && *obj.Status.AppId != "" {
		projectID := common.ValueOf(obj.Spec.Project)
		if projectID == "" {
			var err error
			projectID, err = refsv1beta1.ResolveProjectID(ctx, reader, obj)
			if err != nil {
				return nil, err
			}
		}
		if projectID != "" {
			return &FirebaseWebAppIdentity{
				Project: projectID,
				App:     *obj.Status.AppId,
			}, nil
		}
	}

	return nil, nil
}
