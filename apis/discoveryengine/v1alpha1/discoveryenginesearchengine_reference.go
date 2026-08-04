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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var DiscoveryEngineSearchEngineGVK = schema.GroupVersionKind{
	Group:   "discoveryengine.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "DiscoveryEngineSearchEngine",
}

var _ refs.Ref = &DiscoveryEngineSearchEngineRef{}

// DiscoveryEngineSearchEngineRef is a reference to a DiscoveryEngineSearchEngine.
type DiscoveryEngineSearchEngineRef struct {
	// A reference to an externally managed DiscoveryEngineSearchEngine resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/collections/{{collectionID}}/dataStores/{{dataStoreID}}/siteSearchEngine".
	External string `json:"external,omitempty"`

	// The name of a DiscoveryEngineSearchEngine resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DiscoveryEngineSearchEngine resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&DiscoveryEngineSearchEngineRef{}, &DiscoveryEngineSearchEngine{})
}

func (r *DiscoveryEngineSearchEngineRef) GetGVK() schema.GroupVersionKind {
	return DiscoveryEngineSearchEngineGVK
}

func (r *DiscoveryEngineSearchEngineRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *DiscoveryEngineSearchEngineRef) GetExternal() string {
	return r.External
}

func (r *DiscoveryEngineSearchEngineRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *DiscoveryEngineSearchEngineRef) ValidateExternal(ref string) error {
	id := &DiscoveryEngineSearchEngineIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *DiscoveryEngineSearchEngineRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &DiscoveryEngineSearchEngineIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *DiscoveryEngineSearchEngineRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	// Since DiscoveryEngineSearchEngine is a modern resource supporting status.externalRef,
	// we delegate Normalize strictly to refs.Normalize as per Step 4 critical warning in SKILL.md.
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}
