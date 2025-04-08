// Copyright 2024 Google LLC
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
	"strings"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type StorageBucketRef struct {
	/* The StorageBucket selfLink, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `StorageBucket` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `StorageBucket` resource. */
	Namespace string `json:"namespace,omitempty"`
}

// ResolveStorageBucketRef will resolve a StorageBucketRef.
func ResolveStorageBucketRef(ctx context.Context, reader client.Reader, src client.Object, ref *StorageBucketRef) (*StorageBucketRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on StorageBucketRef")
	}
	if ref.Name != "" && ref.External != "" {
		return nil, fmt.Errorf("cannot specify both name and external on StorageBucketRef")
	}

	// External should be in the `projects/[project_id]/buckets/[bucket]` format
	if ref.External != "" {
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "buckets" {
			ref = &StorageBucketRef{
				External: fmt.Sprintf("projects/%s/buckets/%s", tokens[1], tokens[3]),
			}
			return ref, nil
		}
		return nil, fmt.Errorf("format of StorageBucketRef external=%q was not known (use projects/[project_id]/buckets/[bucket])", ref.External)
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	// Fetch object from k8s cluster to construct the external form
	bucket := &unstructured.Unstructured{}
	bucket.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "storage.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "StorageBucket",
	})
	if err := reader.Get(ctx, key, bucket); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced StorageBucketRef %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced StorageBucketRef %v: %w", key, err)
	}

	storagebucketID, err := GetResourceID(bucket)
	if err != nil {
		return nil, err
	}

	projectID, err := ResolveProjectID(ctx, reader, bucket)
	if err != nil {
		return nil, err
	}
	ref = &StorageBucketRef{
		External: fmt.Sprintf("projects/%s/buckets/%s", projectID, storagebucketID),
	}

	return ref, nil
}
