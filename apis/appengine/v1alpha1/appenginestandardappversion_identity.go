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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &AppEngineStandardAppVersionIdentity{}
	_ identity.Resource   = &AppEngineStandardAppVersion{}
)

var AppEngineStandardAppVersionIdentityFormat = gcpurls.Template[AppEngineStandardAppVersionIdentity]("appengine.googleapis.com", "apps/{app}/services/{service}/versions/{version}")

// AppEngineStandardAppVersionIdentity is the identity of a GCP AppEngineStandardAppVersion resource.
// +k8s:deepcopy-gen=false
type AppEngineStandardAppVersionIdentity struct {
	App     string
	Service string
	Version string
}

func (i *AppEngineStandardAppVersionIdentity) String() string {
	return AppEngineStandardAppVersionIdentityFormat.ToString(*i)
}

func (i *AppEngineStandardAppVersionIdentity) FromExternal(ref string) error {
	parsed, match, err := AppEngineStandardAppVersionIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of AppEngineStandardAppVersion external=%q was not known (use %s): %w", ref, AppEngineStandardAppVersionIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of AppEngineStandardAppVersion external=%q was not known (use %s)", ref, AppEngineStandardAppVersionIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *AppEngineStandardAppVersionIdentity) Host() string {
	return AppEngineStandardAppVersionIdentityFormat.Host()
}

func (i *AppEngineStandardAppVersionIdentity) ParentString() string {
	return "apps/" + i.App + "/services/" + i.Service
}

func getIdentityFromAppEngineStandardAppVersionSpec(ctx context.Context, reader client.Reader, obj *AppEngineStandardAppVersion) (*AppEngineStandardAppVersionIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	if obj.Spec.ServiceRef == nil {
		return nil, fmt.Errorf("spec.serviceRef is required")
	}
	serviceRef := *obj.Spec.ServiceRef
	if err := serviceRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("resolving spec.serviceRef: %w", err)
	}
	parsed := &AppEngineStandardAppVersionIdentity{}
	if err := parsed.FromExternal(strings.TrimSuffix(serviceRef.GetExternal(), "/") + "/versions/" + resourceID); err != nil {
		return nil, fmt.Errorf("parsing spec.serviceRef external=%q: %w", serviceRef.GetExternal(), err)
	}
	parsed.Version = resourceID
	return parsed, nil
}

func (obj *AppEngineStandardAppVersion) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromAppEngineStandardAppVersionSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &AppEngineStandardAppVersionIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change AppEngineStandardAppVersion identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
