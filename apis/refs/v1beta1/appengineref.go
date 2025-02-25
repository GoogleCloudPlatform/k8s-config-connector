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

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TODO: move to apis/appengine when we implement the AppEngine resource

type AppEngineApplicationRef struct {
	// Format: projects/{projects_id}/iap_web/appengine-{app_id}
	External string `json:"external,omitempty"`
}

type AppEngineServiceRef struct {
	// Format: projects/{projects_id}/iap_web/appengine-{app_id}/service/{service_id}
	External string `json:"external,omitempty"`
}

type AppEngineVersionRef struct {
	// Format: projects/{projects_id}/iap_web/appengine-{app_id}/service/{service_id}/version/{version_id}
	External string `json:"external,omitempty"`
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
