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

var _ identity.IdentityV2 = &FirebaseAndroidAppIdentity{}
var _ identity.Resource = &FirebaseAndroidApp{}

var FirebaseAndroidAppIdentityFormat = gcpurls.Template[FirebaseAndroidAppIdentity](
	"firebase.googleapis.com",
	"projects/{project}/androidApps/{app}",
)

// +k8s:deepcopy-gen=false
// FirebaseAndroidAppIdentity is the identity of a GCP FirebaseAndroidApp resource.
type FirebaseAndroidAppIdentity struct {
	Project string
	App     string
}

func (i *FirebaseAndroidAppIdentity) String() string {
	return FirebaseAndroidAppIdentityFormat.ToString(*i)
}

func (i *FirebaseAndroidAppIdentity) FromExternal(ref string) error {
	parsed, match, err := FirebaseAndroidAppIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of FirebaseAndroidApp external=%q was not known (use %s): %w", ref, FirebaseAndroidAppIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of FirebaseAndroidApp external=%q was not known (use %s)", ref, FirebaseAndroidAppIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *FirebaseAndroidAppIdentity) Host() string {
	return FirebaseAndroidAppIdentityFormat.Host()
}

func (i *FirebaseAndroidAppIdentity) ParentString() string {
	return "projects/" + i.Project
}

func getIdentityFromFirebaseAndroidAppSpec(ctx context.Context, reader client.Reader, obj *FirebaseAndroidApp) (*FirebaseAndroidAppIdentity, error) {
	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		// AndroidApp has service-generated appId, but user might set it for acquisition
		return nil, nil
	}

	return &FirebaseAndroidAppIdentity{
		Project: projectID,
		App:     resourceID,
	}, nil
}

func (obj *FirebaseAndroidApp) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// First check status.externalRef or status.name
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		externalRef = common.ValueOf(obj.Status.Name)
	}
	if externalRef != "" {
		id := &FirebaseAndroidAppIdentity{}
		if err := id.FromExternal(externalRef); err != nil {
			return nil, err
		}
		return id, nil
	}

	// Next check spec
	specIdentity, err := getIdentityFromFirebaseAndroidAppSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	if specIdentity != nil {
		return specIdentity, nil
	}

	// Also check status.appId if present
	if obj.Status.AppId != nil && *obj.Status.AppId != "" {
		projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, obj)
		if err == nil && projectID != "" {
			return &FirebaseAndroidAppIdentity{
				Project: projectID,
				App:     *obj.Status.AppId,
			}, nil
		}
	}

	return nil, nil
}
