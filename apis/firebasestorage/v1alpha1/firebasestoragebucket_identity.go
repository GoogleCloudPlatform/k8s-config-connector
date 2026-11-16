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

var _ identity.IdentityV2 = &FirebaseStorageBucketIdentity{}
var _ identity.Resource = &FirebaseStorageBucket{}

var FirebaseStorageBucketIdentityFormat = gcpurls.Template[FirebaseStorageBucketIdentity](
	"firebasestorage.googleapis.com",
	"projects/{project}/buckets/{bucket}",
)

// +k8s:deepcopy-gen=false
// FirebaseStorageBucketIdentity is the identity of a GCP FirebaseStorageBucket resource.
type FirebaseStorageBucketIdentity struct {
	Project string
	Bucket  string
}

func (i *FirebaseStorageBucketIdentity) String() string {
	return FirebaseStorageBucketIdentityFormat.ToString(*i)
}

func (i *FirebaseStorageBucketIdentity) FromExternal(ref string) error {
	parsed, match, err := FirebaseStorageBucketIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of FirebaseStorageBucket external=%q was not known (use %s): %w", ref, FirebaseStorageBucketIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of FirebaseStorageBucket external=%q was not known (use %s)", ref, FirebaseStorageBucketIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *FirebaseStorageBucketIdentity) Host() string {
	return FirebaseStorageBucketIdentityFormat.Host()
}

func (i *FirebaseStorageBucketIdentity) ParentString() string {
	return "projects/" + i.Project
}

func getIdentityFromFirebaseStorageBucketSpec(ctx context.Context, reader client.Reader, obj *FirebaseStorageBucket) (*FirebaseStorageBucketIdentity, error) {
	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	bucketID := common.ValueOf(obj.Spec.ResourceID)
	if bucketID == "" {
		bucketID = obj.GetName()
	}
	if bucketID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	return &FirebaseStorageBucketIdentity{
		Project: projectID,
		Bucket:  bucketID,
	}, nil
}

func (obj *FirebaseStorageBucket) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		externalRef = common.ValueOf(obj.Status.Name)
	}
	if externalRef != "" {
		id := &FirebaseStorageBucketIdentity{}
		if err := id.FromExternal(externalRef); err != nil {
			return nil, err
		}
		return id, nil
	}

	return getIdentityFromFirebaseStorageBucketSpec(ctx, reader, obj)
}
