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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeSecurityPolicyIdentity{}
	_ identity.Resource   = &ComputeSecurityPolicy{}
)

var ComputeGlobalSecurityPolicyIdentityFormat = gcpurls.Template[ComputeSecurityPolicyIdentity]("compute.googleapis.com", "projects/{project}/global/securityPolicies/{name}")
var ComputeRegionalSecurityPolicyIdentityFormat = gcpurls.Template[ComputeSecurityPolicyIdentity]("compute.googleapis.com", "projects/{project}/regions/{region}/securityPolicies/{name}")

// +k8s:deepcopy-gen=false
type ComputeSecurityPolicyIdentity struct {
	Project string
	Region  string
	Name    string
}

func (i *ComputeSecurityPolicyIdentity) IsGlobal() bool {
	return i.Region == "" || i.Region == "global"
}

func (i *ComputeSecurityPolicyIdentity) String() string {
	if !i.IsGlobal() {
		return ComputeRegionalSecurityPolicyIdentityFormat.ToString(*i)
	}
	return ComputeGlobalSecurityPolicyIdentityFormat.ToString(*i)
}

func (i *ComputeSecurityPolicyIdentity) FromExternal(ref string) error {
	ref = refs.TrimComputeURIPrefix(ref)

	if parsed, match, _ := ComputeGlobalSecurityPolicyIdentityFormat.Parse(ref); match {
		*i = *parsed
		i.Region = "global"
		return nil
	}
	if parsed, match, _ := ComputeRegionalSecurityPolicyIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of ComputeSecurityPolicy external=%q was not known (use %s or %s)", ref, ComputeGlobalSecurityPolicyIdentityFormat.CanonicalForm(), ComputeRegionalSecurityPolicyIdentityFormat.CanonicalForm())
}

func (i *ComputeSecurityPolicyIdentity) Host() string {
	return ComputeGlobalSecurityPolicyIdentityFormat.Host()
}

func getIdentityFromComputeSecurityPolicySpec(ctx context.Context, reader client.Reader, obj client.Object) (*ComputeSecurityPolicyIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &ComputeSecurityPolicyIdentity{
		Project: projectID,
		Name:    resourceID,
	}

	// Read location from the spec if present to handle regional ComputeSecurityPolicy objects.
	u, ok := obj.(*unstructured.Unstructured)
	if !ok {
		m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err == nil {
			u = &unstructured.Unstructured{Object: m}
			ok = true
		}
	}
	if ok {
		if region, _, _ := unstructured.NestedString(u.Object, "spec", "region"); region != "" {
			identity.Region = region
		}
	}

	return identity, nil
}

func (obj *ComputeSecurityPolicy) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeSecurityPolicySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.SelfLink)
	if externalRef != "" {
		statusIdentity := &ComputeSecurityPolicyIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeSecurityPolicy identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
