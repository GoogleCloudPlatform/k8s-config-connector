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

package v1beta1

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ Ref = &RedisClusterRef{}

// RedisClusterRef represents a reference to a RedisCluster resource.
type RedisClusterRef struct {
	// A reference to an externally managed target.
	// Should be in the format:
	// projects/<project>/locations/<region>/clusters/<cluster-id>
	// +optional
	External string `json:"external,omitempty"`

	// The `name` field of a `RedisCluster` resource.
	// +optional
	Name string `json:"name,omitempty"`

	// The `namespace` field of a `RedisCluster` resource.
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

func (r *RedisClusterRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "redis.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "RedisCluster",
	}
}

func (r *RedisClusterRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *RedisClusterRef) GetExternal() string {
	return r.External
}

func (r *RedisClusterRef) SetExternal(ref string) {
	r.External = ref
}

func (r *RedisClusterRef) ValidateExternal(ref string) error {
	if ref == "" {
		return fmt.Errorf("external cannot be empty")
	}
	return nil
}

func (r *RedisClusterRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return Normalize(ctx, reader, r, defaultNamespace)
}
