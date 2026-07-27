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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &DialogflowCXAgentRef{}

var DialogflowCXAgentGVK = schema.GroupVersionKind{
	Group:   "dialogflowcx.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "DialogflowCXAgent",
}

// DialogflowCXAgentRef is a reference to a GCP DialogflowCXAgent.
type DialogflowCXAgentRef struct {
	// A reference to an externally managed DialogflowCXAgent resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/agents/{{agentID}}".
	External string `json:"external,omitempty"`

	// The name of a DialogflowCXAgent resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DialogflowCXAgent resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&DialogflowCXAgentRef{})
}

func (r *DialogflowCXAgentRef) GetGVK() schema.GroupVersionKind {
	return DialogflowCXAgentGVK
}

func (r *DialogflowCXAgentRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *DialogflowCXAgentRef) GetExternal() string {
	return r.External
}

func (r *DialogflowCXAgentRef) SetExternal(ref string) {
	r.External = ref
}

func (r *DialogflowCXAgentRef) ValidateExternal(ref string) error {
	if !strings.HasPrefix(ref, "projects/") {
		return fmt.Errorf("DialogflowCXAgent external %q must start with 'projects/'", ref)
	}
	parts := strings.Split(ref, "/")
	if len(parts) != 6 || parts[2] != "locations" || parts[4] != "agents" {
		return fmt.Errorf("DialogflowCXAgent external %q must be in format projects/{project}/locations/{location}/agents/{agent}", ref)
	}
	return nil
}

func (r *DialogflowCXAgentRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &DialogflowCXAgentIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *DialogflowCXAgentRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		name, _, _ := unstructured.NestedString(u.Object, "status", "name")
		return name
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

// +k8s:deepcopy-gen=false
type DialogflowCXAgentIdentity struct {
	Project  string
	Location string
	Agent    string
}

var _ identity.Identity = &DialogflowCXAgentIdentity{}

func (i *DialogflowCXAgentIdentity) Host() string {
	return "dialogflow.googleapis.com"
}

func (i *DialogflowCXAgentIdentity) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/agents/%s", i.Project, i.Location, i.Agent)
}

func (i *DialogflowCXAgentIdentity) FromExternal(ref string) error {
	if !strings.HasPrefix(ref, "projects/") {
		return fmt.Errorf("DialogflowCXAgent external %q must start with 'projects/'", ref)
	}
	parts := strings.Split(ref, "/")
	if len(parts) != 6 || parts[2] != "locations" || parts[4] != "agents" {
		return fmt.Errorf("DialogflowCXAgent external %q must be in format projects/{project}/locations/{location}/agents/{agent}", ref)
	}
	i.Project = parts[1]
	i.Location = parts[3]
	i.Agent = parts[5]
	return nil
}
