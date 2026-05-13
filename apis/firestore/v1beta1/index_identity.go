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

var (
	_ identity.IdentityV2 = &FirestoreIndexIdentity{}
	_ identity.Resource   = &FirestoreIndex{}
)

var FirestoreIndexIdentityFormat = gcpurls.Template[FirestoreIndexIdentity]("firestore.googleapis.com", "projects/{project}/databases/{database}/collectionGroups/{collectionGroup}/indexes/{index}")

// +k8s:deepcopy-gen=false
type FirestoreIndexIdentity struct {
	Project         string
	Database        string
	CollectionGroup string
	Index           string
}

func (i *FirestoreIndexIdentity) String() string {
	return FirestoreIndexIdentityFormat.ToString(*i)
}

func (i *FirestoreIndexIdentity) FromExternal(ref string) error {
	parsed, match, err := FirestoreIndexIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of FirestoreIndex external=%q was not known (use %s): %w", ref, FirestoreIndexIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of FirestoreIndex external=%q was not known (use %s)", ref, FirestoreIndexIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *FirestoreIndexIdentity) Host() string {
	return FirestoreIndexIdentityFormat.Host()
}

func (i *FirestoreIndexIdentity) Parent() *FirestoreCollectionGroupIdentity {
	return &FirestoreCollectionGroupIdentity{
		Project:         i.Project,
		Database:        i.Database,
		CollectionGroup: i.CollectionGroup,
	}
}

// +k8s:deepcopy-gen=false
type FirestoreCollectionGroupIdentity struct {
	Project         string
	Database        string
	CollectionGroup string
}

var FirestoreCollectionGroupIdentityFormat = gcpurls.Template[FirestoreCollectionGroupIdentity]("firestore.googleapis.com", "projects/{project}/databases/{database}/collectionGroups/{collectionGroup}")

func (i *FirestoreCollectionGroupIdentity) String() string {
	return FirestoreCollectionGroupIdentityFormat.ToString(*i)
}

func getIdentityFromFirestoreIndexSpec(ctx context.Context, reader client.Reader, obj client.Object) (*FirestoreIndexIdentity, error) {
	spec := obj.(*FirestoreIndex).Spec

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	database := "(default)"
	if spec.Database != nil {
		database = *spec.Database
	}

	collection := spec.Collection
	if collection == "" {
		return nil, fmt.Errorf("collection is required")
	}

	identity := &FirestoreIndexIdentity{
		Project:         projectID,
		Database:        database,
		CollectionGroup: collection,
		Index:           "", // Will be populated from status or creation
	}

	return identity, nil
}
func (obj *FirestoreIndex) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromFirestoreIndexSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.Name)
	if externalRef != "" {
		statusIdentity := &FirestoreIndexIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.Project != specIdentity.Project || statusIdentity.Database != specIdentity.Database || statusIdentity.CollectionGroup != specIdentity.CollectionGroup {
			return nil, fmt.Errorf("cannot change FirestoreIndex identity (old=%q, new parent=%q)", externalRef, specIdentity.Parent().String())
		}

		specIdentity.Index = statusIdentity.Index
	}

	return specIdentity, nil
}
