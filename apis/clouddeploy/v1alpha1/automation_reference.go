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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &CloudDeployAutomationRef{}
var _ refsv1beta1.Ref = &CloudDeployAutomationRef{}
var _ refsv1beta1.ExternalRef = &CloudDeployAutomationRef{}

type CloudDeployAutomationRef struct {
	// A reference to an externally managed CloudDeployAutomation resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/deliveryPipelines/{{deliveryPipelineID}}/automations/{{automationID}}".
	External string `json:"external,omitempty"`

	// The name of a CloudDeployAutomation resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CloudDeployAutomation resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *CloudDeployAutomationRef) GetGVK() schema.GroupVersionKind {
	return CloudDeployAutomationGVK
}

func (r *CloudDeployAutomationRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
}

func (r *CloudDeployAutomationRef) GetExternal() string {
	return r.External
}

func (r *CloudDeployAutomationRef) SetExternal(external string) {
	r.External = external
}

func (r *CloudDeployAutomationRef) ValidateExternal(external string) error {
	actualIdentity := &CloudDeployAutomationIdentity{}
	return actualIdentity.FromExternal(external)
}

func (r *CloudDeployAutomationRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
}

// NormalizedExternal provision the "External" value for other resource that depends on CloudDeployAutomation.
func (r *CloudDeployAutomationRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if err := r.Normalize(ctx, reader, otherNamespace); err != nil {
		return "", err
	}
	return r.External, nil
}

func (r *CloudDeployAutomationRef) Build(ctx context.Context, reader client.Reader, otherNamespace string, identity identity.Identity) error {
	external, err := r.NormalizedExternal(ctx, reader, otherNamespace)
	if err != nil {
		return err
	}
	return identity.FromExternal(external)
}

func (r *CloudDeployAutomationRef) ParseExternalToIdentity() (identity.Identity, error) {
	if r.External == "" {
		return nil, nil
	}
	actualIdentity := &CloudDeployAutomationIdentity{}
	if err := actualIdentity.FromExternal(r.External); err != nil {
		return nil, err
	}
	return actualIdentity, nil
}
