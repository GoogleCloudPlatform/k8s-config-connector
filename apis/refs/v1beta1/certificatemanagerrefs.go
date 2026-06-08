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
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ Ref = &CertificateManagerCertificateRef{}
var _ Ref = &CertificateManagerTrustConfigRef{}

var CertificateManagerCertificateGVK = schema.GroupVersionKind{
	Group:   "certificatemanager.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "CertificateManagerCertificate",
}

var CertificateManagerTrustConfigGVK = schema.GroupVersionKind{
	Group:   "certificatemanager.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "CertificateManagerTrustConfig",
}

func init() {
	Register(&CertificateManagerCertificateRef{})
	Register(&CertificateManagerTrustConfigRef{})
}

// CertificateManagerCertificateRef is a reference to a CertificateManagerCertificate.
type CertificateManagerCertificateRef struct {
	// A reference to an externally managed CertificateManagerCertificate resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/certificates/{{certificate}}".
	External string `json:"external,omitempty"`

	// The name of a CertificateManagerCertificate resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CertificateManagerCertificate resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *CertificateManagerCertificateRef) GetGVK() schema.GroupVersionKind {
	return CertificateManagerCertificateGVK
}

func (r *CertificateManagerCertificateRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *CertificateManagerCertificateRef) GetExternal() string {
	return r.External
}

func (r *CertificateManagerCertificateRef) SetExternal(ref string) {
	r.External = ref
}

func (r *CertificateManagerCertificateRef) ValidateExternal(ref string) error {
	if !strings.HasPrefix(ref, "projects/") && !strings.HasPrefix(ref, "//certificatemanager.googleapis.com/projects/") {
		return fmt.Errorf("external reference format %q is not known; expected projects/<project>/locations/<location>/certificates/<name> or //certificatemanager.googleapis.com/projects/<project>/locations/<location>/certificates/<name>", ref)
	}
	return nil
}

func (r *CertificateManagerCertificateRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return Normalize(ctx, reader, r, defaultNamespace)
}

// CertificateManagerTrustConfigRef is a reference to a CertificateManagerTrustConfig.
type CertificateManagerTrustConfigRef struct {
	// A reference to an externally managed CertificateManagerTrustConfig resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/trustConfigs/{{trustConfig}}".
	External string `json:"external,omitempty"`

	// The name of a CertificateManagerTrustConfig resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CertificateManagerTrustConfig resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *CertificateManagerTrustConfigRef) GetGVK() schema.GroupVersionKind {
	return CertificateManagerTrustConfigGVK
}

func (r *CertificateManagerTrustConfigRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *CertificateManagerTrustConfigRef) GetExternal() string {
	return r.External
}

func (r *CertificateManagerTrustConfigRef) SetExternal(ref string) {
	r.External = ref
}

func (r *CertificateManagerTrustConfigRef) ValidateExternal(ref string) error {
	if !strings.HasPrefix(ref, "projects/") && !strings.HasPrefix(ref, "//certificatemanager.googleapis.com/projects/") {
		return fmt.Errorf("external reference format %q is not known; expected projects/<project>/locations/<location>/trustConfigs/<name> or //certificatemanager.googleapis.com/projects/<project>/locations/<location>/trustConfigs/<name>", ref)
	}
	return nil
}

func (r *CertificateManagerTrustConfigRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return Normalize(ctx, reader, r, defaultNamespace)
}
