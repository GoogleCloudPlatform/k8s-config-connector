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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &AccessContextManagerServicePerimeterResourceIdentity{}
	_ identity.Resource   = &AccessContextManagerServicePerimeterResource{}
)

var AccessContextManagerServicePerimeterResourceIdentityFormat = gcpurls.Template[AccessContextManagerServicePerimeterResourceIdentity]("accesscontextmanager.googleapis.com", "accessPolicies/{accessPolicy}/servicePerimeters/{servicePerimeter}/projects/{project}")

// +k8s:deepcopy-gen=false
type AccessContextManagerServicePerimeterResourceIdentity struct {
	AccessPolicy     string
	ServicePerimeter string
	Project          string
}

func (i *AccessContextManagerServicePerimeterResourceIdentity) String() string {
	return AccessContextManagerServicePerimeterResourceIdentityFormat.ToString(*i)
}

func (i *AccessContextManagerServicePerimeterResourceIdentity) FromExternal(ref string) error {
	parsed, match, err := AccessContextManagerServicePerimeterResourceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of AccessContextManagerServicePerimeterResource external=%q was not known (use %s): %w", ref, AccessContextManagerServicePerimeterResourceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of AccessContextManagerServicePerimeterResource external=%q was not known (use %s)", ref, AccessContextManagerServicePerimeterResourceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *AccessContextManagerServicePerimeterResourceIdentity) Host() string {
	return AccessContextManagerServicePerimeterResourceIdentityFormat.Host()
}

func getIdentityFromAccessContextManagerServicePerimeterResourceSpec(ctx context.Context, reader client.Reader, obj *AccessContextManagerServicePerimeterResource) (*AccessContextManagerServicePerimeterResourceIdentity, error) {
	if obj.Spec.PerimeterNameRef == nil {
		return nil, fmt.Errorf("spec.perimeterNameRef is required")
	}

	if err := obj.Spec.PerimeterNameRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("cannot resolve perimeterNameRef: %w", err)
	}

	perimeterExternal := obj.Spec.PerimeterNameRef.GetExternal()
	perimeterIdentity := &ServicePerimeterIdentity{}
	if err := perimeterIdentity.FromExternal(perimeterExternal); err != nil {
		return nil, fmt.Errorf("cannot parse perimeterNameRef external: %w", err)
	}

	if obj.Spec.ResourceRef == nil {
		return nil, fmt.Errorf("spec.resourceRef is required")
	}

	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ResourceRef)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resourceRef: %w", err)
	}

	identity := &AccessContextManagerServicePerimeterResourceIdentity{
		AccessPolicy:     perimeterIdentity.AccessPolicy,
		ServicePerimeter: perimeterIdentity.ServicePerimeter,
		Project:          projectRef.ProjectID,
	}
	return identity, nil
}

func (obj *AccessContextManagerServicePerimeterResource) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromAccessContextManagerServicePerimeterResourceSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &AccessContextManagerServicePerimeterResourceIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change AccessContextManagerServicePerimeterResource identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
