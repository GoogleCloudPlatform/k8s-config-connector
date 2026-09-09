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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &DialogflowCXAgentRef{}
var DialogflowCXAgentGVK = GroupVersion.WithKind("DialogflowCXAgent")

// DialogflowCXAgentRef is a reference to a DialogflowCXAgent.
type DialogflowCXAgentRef struct {
	/* A reference to an externally managed DialogflowCXAgent resource.
	   Should be in the format "projects/{{projectID}}/locations/{{location}}/agents/{{agentID}}". */
	External string `json:"external,omitempty"`

	/* NOTYET
	// The name of a DialogflowCXAgent resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DialogflowCXAgent resource.
	Namespace string `json:"namespace,omitempty"`
	*/
}

func (r *DialogflowCXAgentRef) GetGVK() schema.GroupVersionKind {
	return DialogflowCXAgentGVK
}

func (r *DialogflowCXAgentRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{}
}

func (r *DialogflowCXAgentRef) GetExternal() string {
	return r.External
}

func (r *DialogflowCXAgentRef) SetExternal(ref string) {
	r.External = ref
}

func (r *DialogflowCXAgentRef) ValidateExternal(ref string) error {
	id := &DialogflowCXAgentIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *DialogflowCXAgentRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External == "" {
		return fmt.Errorf("external reference must be specified for %s", DialogflowCXAgentGVK.Kind)
	}
	return r.ValidateExternal(r.External)
}
