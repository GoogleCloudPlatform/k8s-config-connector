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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &AgentSearchEngineRef{}

var AgentSearchEngineGVK = GroupVersion.WithKind("AgentSearchEngine")

type AgentSearchEngineRef struct {
	// A reference to an externally managed AgentSearchEngine resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/collections/{{collectionID}}/engines/{{engineID}}".
	External string `json:"external,omitempty"`

	// Name of the referent.
	// +optional
	// Name string `json:"name,omitempty"`

	// Namespace of the referent.
	// +optional
	// Namespace string `json:"namespace,omitempty"`
}

func (r *AgentSearchEngineRef) GetGVK() schema.GroupVersionKind {
	return AgentSearchEngineGVK
}

func (r *AgentSearchEngineRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{}
}

func (r *AgentSearchEngineRef) GetExternal() string {
	return r.External
}

func (r *AgentSearchEngineRef) SetExternal(ref string) {
	r.External = ref
}

func (r *AgentSearchEngineRef) ValidateExternal(ref string) error {
	if _, err := ParseAgentSearchEngineExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *AgentSearchEngineRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External == "" {
		return fmt.Errorf("external reference must be specified for %s", AgentSearchEngineGVK.Kind)
	}
	return r.ValidateExternal(r.External)
}

func ParseAgentSearchEngineExternal(external string) (string, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "collections" || tokens[6] != "engines" {
		return "", fmt.Errorf("format of AgentSearchEngine external=%q was not known (use projects/{{projectID}}/locations/{{location}}/collections/{{collectionID}}/engines/{{engineID}})", external)
	}
	return external, nil
}
