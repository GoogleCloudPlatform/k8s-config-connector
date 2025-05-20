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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// AnywhereCacheIdentity defines the resource reference to StorageAnywhereCache, which "External" field
// holds the GCP identifier for the KRM object.
type AnywhereCacheIdentity struct {
	parent *AnywhereCacheParent
	id     string
}

func (i *AnywhereCacheIdentity) String() string {
	return i.parent.String() + "/anywhereCaches/" + i.id
}

func (i *AnywhereCacheIdentity) ID() string {
	return i.id
}

func (i *AnywhereCacheIdentity) HasKnownId() bool {
	return i.id != ""
}

func (i *AnywhereCacheIdentity) Parent() *AnywhereCacheParent {
	return i.parent
}

type AnywhereCacheParent struct {
	BucketName string
}

func (p *AnywhereCacheParent) String() string {
	return "projects/_/buckets/" + p.BucketName
}

// New builds a AnywhereCacheIdentity from parent, and resourceID
func GetAnywhereCacheIdentity(parent *AnywhereCacheParent, id string) *AnywhereCacheIdentity {
	return &AnywhereCacheIdentity{
		parent: parent,
		id:     id,
	}
}

// New builds a AnywhereCacheIdentity from the Config Connector AnywhereCache object.
func NewAnywhereCacheIdentity(ctx context.Context, reader client.Reader, obj *StorageAnywhereCache) (*AnywhereCacheIdentity, error) {

	// Get Parent
	storageBucketRef, err := refsv1beta1.ResolveStorageBucketRef(ctx, reader, obj, obj.Spec.BucketRef)
	if err != nil {
		return nil, err
	}
	bucketName, err := ParseBucketExternal(storageBucketRef.External)
	if err != nil {
		return nil, err
	}

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseAnywhereCacheExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.BucketName != bucketName {
			return nil, fmt.Errorf("spec.BucketName changed, expect %s, got %s", actualParent.BucketName, bucketName)
		}
		if resourceID != "" && actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
		resourceID = actualResourceID
	}
	return &AnywhereCacheIdentity{
		parent: &AnywhereCacheParent{
			BucketName: bucketName,
		},
		id: resourceID,
	}, nil
}

func ParseAnywhereCacheExternal(external string) (parent *AnywhereCacheParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "buckets" || tokens[4] != "anywhereCaches" {
		return nil, "", fmt.Errorf("format of StorageAnywhereCache external=%q was not known (use projects/_/buckets/{{bucket}}/anywhereCaches/{{anywhereCacheID}})", external)
	}
	parent = &AnywhereCacheParent{
		BucketName: tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}

func ParseBucketExternal(external string) (string, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "buckets" {
		return "", fmt.Errorf("format of BucketRef external should be projects/{project_id}/buckets/{bucket_name}, but received external=%q", external)
	}
	return tokens[3], nil
}
