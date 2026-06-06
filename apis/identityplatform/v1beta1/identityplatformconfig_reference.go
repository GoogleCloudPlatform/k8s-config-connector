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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &ConfigDefaultTenantLocationRef{}
var _ refsv1beta1.Ref = &ConfigFunctionUriRef{}

func init() {
	refsv1beta1.Register(&ConfigDefaultTenantLocationRef{}, nil)
	refsv1beta1.Register(&ConfigFunctionUriRef{}, nil)
}

func (r *ConfigDefaultTenantLocationRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "resourcemanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "Folder",
	}
}

func (r *ConfigDefaultTenantLocationRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ConfigDefaultTenantLocationRef) GetExternal() string {
	return r.External
}

func (r *ConfigDefaultTenantLocationRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ConfigDefaultTenantLocationRef) ValidateExternal(ref string) error {
	if ref == "" {
		return fmt.Errorf("external reference cannot be empty")
	}
	if !strings.HasPrefix(ref, "folders/") && !strings.HasPrefix(ref, "organizations/") {
		return fmt.Errorf("format of external default tenant location %q was not known (must be folders/{{folderID}} or organizations/{{orgID}})", ref)
	}
	return nil
}

func (r *ConfigDefaultTenantLocationRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
}

func (r *ConfigFunctionUriRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "cloudfunctions.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "CloudFunctionsFunction",
	}
}

func (r *ConfigFunctionUriRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ConfigFunctionUriRef) GetExternal() string {
	return r.External
}

func (r *ConfigFunctionUriRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ConfigFunctionUriRef) ValidateExternal(ref string) error {
	if ref == "" {
		return fmt.Errorf("external reference cannot be empty")
	}
	if !strings.HasPrefix(ref, "http://") && !strings.HasPrefix(ref, "https://") {
		return fmt.Errorf("format of external function URI %q was not known (must be a valid http or https URL)", ref)
	}
	return nil
}

func (r *ConfigFunctionUriRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External == "" {
		key := types.NamespacedName{
			Name:      r.Name,
			Namespace: r.Namespace,
		}
		if key.Namespace == "" {
			key.Namespace = defaultNamespace
		}
		u := &unstructured.Unstructured{}
		u.SetGroupVersionKind(r.GetGVK())
		if err := reader.Get(ctx, key, u); err != nil {
			if apierrors.IsNotFound(err) {
				return k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
			}
			return fmt.Errorf("reading referenced %s %s: %w", r.GetGVK(), key, err)
		}
		// Get external from status.httpsTrigger.url.
		url, _, err := unstructured.NestedString(u.Object, "status", "httpsTrigger", "url")
		if err != nil {
			return fmt.Errorf("reading status.httpsTrigger.url: %w", err)
		}
		if url == "" {
			return k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
		}
		r.SetExternal(url)
	}

	return r.ValidateExternal(r.GetExternal())
}
