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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &CertificateManagerCertificateRef{}

var CertificateManagerCertificateGVK = schema.GroupVersionKind{
	Group:   "certificatemanager.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "CertificateManagerCertificate",
}

// A reference to a CertificateManagerCertificate resource.
type CertificateManagerCertificateRef struct {
	// Allowed value: The format `//certificatemanager.googleapis.com/projects/{{project}}/locations/{{location}}/certificates/{{name}}`.
	External string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
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
	if !strings.HasPrefix(ref, "//certificatemanager.googleapis.com/projects/") {
		return fmt.Errorf("external reference format %q is not known; expected //certificatemanager.googleapis.com/projects/<project>/locations/<location>/certificates/<name>", ref)
	}
	return nil
}

func (r *CertificateManagerCertificateRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if err := refsv1beta1.ValidateNameAndExternal(r.Name, r.External); err != nil {
		return "", fmt.Errorf("in CertificateManagerCertificate reference: %w", err)
	}
	if r.External != "" {
		return r.External, nil
	}

	namespace := r.Namespace
	if namespace == "" {
		namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(CertificateManagerCertificateGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced CertificateManagerCertificate %s: %w", key, err)
	}

	// CertificateManagerCertificate is an unmigrated Terraform-based CRD and does not expose status.externalRef.
	// We must manually construct the external reference.
	resourceID, err := refsv1beta1.GetResourceID(u)
	if err != nil {
		return "", fmt.Errorf("getting resourceID for %s: %w", key, err)
	}

	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
	if err != nil {
		return "", fmt.Errorf("resolving projectID for %s: %w", key, err)
	}

	location, err := refsv1beta1.GetLocation(u)
	if err != nil {
		return "", fmt.Errorf("getting location for %s: %w", key, err)
	}

	return fmt.Sprintf("//certificatemanager.googleapis.com/projects/%s/locations/%s/certificates/%s", projectID, location, resourceID), nil
}

func (r *CertificateManagerCertificateRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) string {
		resourceID, err := refsv1beta1.GetResourceID(u)
		if err != nil {
			return ""
		}
		projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
		if err != nil {
			return ""
		}
		location, err := refsv1beta1.GetLocation(u)
		if err != nil {
			return ""
		}

		if resourceID != "" && projectID != "" && location != "" {
			return fmt.Sprintf("//certificatemanager.googleapis.com/projects/%s/locations/%s/certificates/%s", projectID, location, resourceID)
		}
		return ""
	})
}
