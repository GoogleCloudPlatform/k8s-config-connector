// Copyright 2024 Google LLC
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

// TODO: move to apis/appengine when we implement the AppEngine resource

var (
	_ Ref = &AppEngineApplicationRef{}
	_ Ref = &AppEngineServiceRef{}
	_ Ref = &AppEngineVersionRef{}
)

type AppEngineApplicationRef struct {
	// Format: projects/{projects_id}/iap_web/appengine-{app_id}
	External string `json:"external,omitempty"`
}

func (r *AppEngineApplicationRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "appengine.cnrm.cloud.google.com", Version: "v1beta1", Kind: "AppEngineApplication"}
}

func (r *AppEngineApplicationRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{}
}

func (r *AppEngineApplicationRef) GetExternal() string {
	return r.External
}

func (r *AppEngineApplicationRef) SetExternal(ref string) {
	r.External = ref
}

func (r *AppEngineApplicationRef) ValidateExternal(ref string) error {
	if ref == "" {
		return fmt.Errorf("invalid AppEngineApplicationRef: must specify external")
	}
	return nil
}

func (r *AppEngineApplicationRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return r.ValidateExternal(r.External)
}

type AppEngineServiceRef struct {
	// Format: projects/{projects_id}/iap_web/appengine-{app_id}/service/{service_id}
	External string `json:"external,omitempty"`
}

func (r *AppEngineServiceRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "appengine.cnrm.cloud.google.com", Version: "v1beta1", Kind: "AppEngineService"}
}

func (r *AppEngineServiceRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{}
}

func (r *AppEngineServiceRef) GetExternal() string {
	return r.External
}

func (r *AppEngineServiceRef) SetExternal(ref string) {
	r.External = ref
}

func (r *AppEngineServiceRef) ValidateExternal(ref string) error {
	if ref == "" {
		return fmt.Errorf("invalid AppEngineServiceRef: must specify external")
	}
	return nil
}

func (r *AppEngineServiceRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return r.ValidateExternal(r.External)
}

type AppEngineVersionRef struct {
	// Format: projects/{projects_id}/iap_web/appengine-{app_id}/service/{service_id}/version/{version_id}
	External string `json:"external,omitempty"`
}

func (r *AppEngineVersionRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "appengine.cnrm.cloud.google.com", Version: "v1beta1", Kind: "AppEngineVersion"}
}

func (r *AppEngineVersionRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{}
}

func (r *AppEngineVersionRef) GetExternal() string {
	return r.External
}

func (r *AppEngineVersionRef) SetExternal(ref string) {
	r.External = ref
}

func (r *AppEngineVersionRef) ValidateExternal(ref string) error {
	if ref == "" {
		return fmt.Errorf("invalid AppEngineVersionRef: must specify external")
	}
	return nil
}

func (r *AppEngineVersionRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return r.ValidateExternal(r.External)
}

func ResolveAppEngineApplicationID(ctx context.Context, reader client.Reader, otherNamespace string, ref *AppEngineApplicationRef) (string, error) {
	if ref == nil {
		return "", nil
	}

	if ref.External != "" {
		return ref.External, nil
	}

	return "", fmt.Errorf("invalid AppEngineApplicationRef: must specify external")
}

func ResolveAppEngineServiceID(ctx context.Context, reader client.Reader, otherNamespace string, ref *AppEngineServiceRef) (string, error) {
	if ref == nil {
		return "", nil
	}

	if ref.External != "" {
		return ref.External, nil
	}

	return "", fmt.Errorf("invalid AppEngineServiceRef: must specify external")
}

func ResolveAppEngineVersionID(ctx context.Context, reader client.Reader, otherNamespace string, ref *AppEngineVersionRef) (string, error) {
	if ref == nil {
		return "", nil
	}

	if ref.External != "" {
		return ref.External, nil
	}

	return "", fmt.Errorf("invalid AppEngineVersionRef: must specify external")
}
