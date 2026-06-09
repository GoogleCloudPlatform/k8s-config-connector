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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &CloudFunctionsFunctionIdentity{}
	_ identity.Resource   = &CloudFunctionsFunction{}
)

var CloudFunctionsFunctionIdentityFormat = gcpurls.Template[CloudFunctionsFunctionIdentity]("cloudfunctions.googleapis.com", "projects/{project}/locations/{location}/functions/{function}")

// +k8s:deepcopy-gen=false
type CloudFunctionsFunctionIdentity struct {
	Project  string
	Location string
	Function string
}

func (i *CloudFunctionsFunctionIdentity) String() string {
	return CloudFunctionsFunctionIdentityFormat.ToString(*i)
}

func (i *CloudFunctionsFunctionIdentity) FromExternal(ref string) error {
	parsed, match, err := CloudFunctionsFunctionIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CloudFunctionsFunction external=%q was not known (use %s): %w", ref, CloudFunctionsFunctionIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CloudFunctionsFunction external=%q was not known (use %s)", ref, CloudFunctionsFunctionIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *CloudFunctionsFunctionIdentity) Host() string {
	return CloudFunctionsFunctionIdentityFormat.Host()
}

func getIdentityFromCloudFunctionsFunctionSpec(ctx context.Context, reader client.Reader, obj client.Object) (*CloudFunctionsFunctionIdentity, error) {
	u, ok := obj.(*unstructured.Unstructured)
	if !ok {
		m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			return nil, fmt.Errorf("failed to convert to unstructured: %w", err)
		}
		u = &unstructured.Unstructured{Object: m}
	}

	region, _, err := unstructured.NestedString(u.Object, "spec", "region")
	if err != nil || region == "" {
		return nil, fmt.Errorf("cannot resolve spec.region: %w", err)
	}

	resourceID, _, err := unstructured.NestedString(u.Object, "spec", "resourceID")
	if err != nil {
		return nil, fmt.Errorf("cannot resolve spec.resourceID: %w", err)
	}
	if resourceID == "" {
		resourceID = u.GetName()
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &CloudFunctionsFunctionIdentity{
		Project:  projectID,
		Location: region,
		Function: resourceID,
	}
	return identity, nil
}

func (obj *CloudFunctionsFunction) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromCloudFunctionsFunctionSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status values if present.
	// CloudFunctions does not yet have status.externalRef, so we will use specIdentity.
	return specIdentity, nil
}
