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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &VertexAICustomJobRef{}

// VertexAICustomJobRef is a reference to a VertexAICustomJob.
type VertexAICustomJobRef struct {
	// A reference to an externally managed VertexAICustomJob resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/customJobs/{{customJobID}}".
	External string `json:"external,omitempty"`

	// The name of a VertexAICustomJob resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAICustomJob resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&VertexAICustomJobRef{}, &VertexAICustomJob{})
}

func (r *VertexAICustomJobRef) GetGVK() schema.GroupVersionKind {
	return VertexAICustomJobGVK
}

func (r *VertexAICustomJobRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *VertexAICustomJobRef) GetExternal() string {
	return r.External
}

func (r *VertexAICustomJobRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *VertexAICustomJobRef) ValidateExternal(ref string) error {
	id := &VertexAICustomJobIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *VertexAICustomJobRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &VertexAICustomJobIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *VertexAICustomJobRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		structuredObj, err := common.ToStructuredType[*VertexAICustomJob](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromVertexAICustomJobSpec(ctx, reader, structuredObj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

var VertexAIExperimentGVK = schema.GroupVersionKind{
	Group:   "vertexai.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "VertexAIExperiment",
}

var _ refs.Ref = &VertexAIExperimentRef{}

// VertexAIExperimentRef is a reference to a VertexAIExperiment.
type VertexAIExperimentRef struct {
	// A reference to an externally managed VertexAIExperiment resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/metadataStores/{{metadataStore}}/contexts/{{experimentID}}".
	External string `json:"external,omitempty"`

	// The name of a VertexAIExperiment resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAIExperiment resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&VertexAIExperimentRef{}, nil)
}

func (r *VertexAIExperimentRef) GetGVK() schema.GroupVersionKind {
	return VertexAIExperimentGVK
}

func (r *VertexAIExperimentRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *VertexAIExperimentRef) GetExternal() string {
	return r.External
}

func (r *VertexAIExperimentRef) SetExternal(external string) {
	r.External = external
	r.Name = ""
	r.Namespace = ""
}

var VertexAIExperimentIdentityFormat = gcpurls.Template[VertexAIExperimentIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/metadataStores/{metadataStore}/contexts/{experiment}")

type VertexAIExperimentIdentity struct {
	Project       string
	Location      string
	MetadataStore string
	Experiment    string
}

func (i *VertexAIExperimentIdentity) String() string {
	return VertexAIExperimentIdentityFormat.ToString(*i)
}

func (i *VertexAIExperimentIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAIExperimentIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAIExperiment external=%q was not known (use %s): %w", ref, VertexAIExperimentIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAIExperiment external=%q was not known (use %s)", ref, VertexAIExperimentIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (r *VertexAIExperimentRef) ValidateExternal(ref string) error {
	id := &VertexAIExperimentIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *VertexAIExperimentRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &VertexAIExperimentIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *VertexAIExperimentRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		// No fallback since we do not have a KRM type/controller for Experiment yet.
		return ""
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

var VertexAIExperimentRunGVK = schema.GroupVersionKind{
	Group:   "vertexai.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "VertexAIExperimentRun",
}

var _ refs.Ref = &VertexAIExperimentRunRef{}

// VertexAIExperimentRunRef is a reference to a VertexAIExperimentRun.
type VertexAIExperimentRunRef struct {
	// A reference to an externally managed VertexAIExperimentRun resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/metadataStores/{{metadataStore}}/contexts/{{experimentRunID}}".
	External string `json:"external,omitempty"`

	// The name of a VertexAIExperimentRun resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAIExperimentRun resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&VertexAIExperimentRunRef{}, nil)
}

func (r *VertexAIExperimentRunRef) GetGVK() schema.GroupVersionKind {
	return VertexAIExperimentRunGVK
}

func (r *VertexAIExperimentRunRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *VertexAIExperimentRunRef) GetExternal() string {
	return r.External
}

func (r *VertexAIExperimentRunRef) SetExternal(external string) {
	r.External = external
	r.Name = ""
	r.Namespace = ""
}

var VertexAIExperimentRunIdentityFormat = gcpurls.Template[VertexAIExperimentRunIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/metadataStores/{metadataStore}/contexts/{experimentRun}")

type VertexAIExperimentRunIdentity struct {
	Project       string
	Location      string
	MetadataStore string
	ExperimentRun string
}

func (i *VertexAIExperimentRunIdentity) String() string {
	return VertexAIExperimentRunIdentityFormat.ToString(*i)
}

func (i *VertexAIExperimentRunIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAIExperimentRunIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAIExperimentRun external=%q was not known (use %s): %w", ref, VertexAIExperimentRunIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAIExperimentRun external=%q was not known (use %s)", ref, VertexAIExperimentRunIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (r *VertexAIExperimentRunRef) ValidateExternal(ref string) error {
	id := &VertexAIExperimentRunIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *VertexAIExperimentRunRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &VertexAIExperimentRunIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *VertexAIExperimentRunRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		// No fallback since we do not have a KRM type/controller for ExperimentRun yet.
		return ""
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
