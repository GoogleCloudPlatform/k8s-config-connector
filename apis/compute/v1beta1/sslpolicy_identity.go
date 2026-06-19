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
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeSSLPolicyIdentity{}
	_ identity.Resource   = &ComputeSSLPolicy{}
)

var ComputeGlobalSSLPolicyIdentityFormat = gcpurls.Template[ComputeSSLPolicyIdentity]("compute.googleapis.com", "projects/{project}/global/sslPolicies/{sslpolicy}")
var ComputeRegionalSSLPolicyIdentityFormat = gcpurls.Template[ComputeSSLPolicyIdentity]("compute.googleapis.com", "projects/{project}/regions/{region}/sslPolicies/{sslpolicy}")

// ComputeSSLPolicyIdentity is the identity of a GCP ComputeSSLPolicy resource.
// +k8s:deepcopy-gen=false
type ComputeSSLPolicyIdentity struct {
	Project   string
	Region    string
	SslPolicy string
}

func (i *ComputeSSLPolicyIdentity) IsGlobal() bool {
	return i.Region == "" || i.Region == "global"
}

func (i *ComputeSSLPolicyIdentity) String() string {
	if !i.IsGlobal() {
		return ComputeRegionalSSLPolicyIdentityFormat.ToString(*i)
	}
	return ComputeGlobalSSLPolicyIdentityFormat.ToString(*i)
}

func (i *ComputeSSLPolicyIdentity) FromExternal(ref string) error {
	ref = apirefs.TrimComputeURIPrefix(ref)

	if parsed, match, _ := ComputeGlobalSSLPolicyIdentityFormat.Parse(ref); match {
		*i = *parsed
		i.Region = "global"
		return nil
	}
	if parsed, match, _ := ComputeRegionalSSLPolicyIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of ComputeSSLPolicy external=%q was not known (use %s or %s)", ref, ComputeGlobalSSLPolicyIdentityFormat.CanonicalForm(), ComputeRegionalSSLPolicyIdentityFormat.CanonicalForm())
}

func (i *ComputeSSLPolicyIdentity) Host() string {
	return ComputeGlobalSSLPolicyIdentityFormat.Host()
}

func (i *ComputeSSLPolicyIdentity) ParentString() string {
	if !i.IsGlobal() {
		return fmt.Sprintf("projects/%s/regions/%s", i.Project, i.Region)
	}
	return fmt.Sprintf("projects/%s/global", i.Project)
}

func ParseComputeSSLPolicyExternal(external string) (*ComputeSSLPolicyIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeSSLPolicy external value")
	}
	id := &ComputeSSLPolicyIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeSSLPolicySpec(ctx context.Context, reader client.Reader, obj *ComputeSSLPolicy) (*ComputeSSLPolicyIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ComputeSSLPolicyIdentity{
		Project:   projectID,
		Region:    "global", // ComputeSSLPolicy is global in KCC
		SslPolicy: resourceID,
	}
	return identity, nil
}

func (obj *ComputeSSLPolicy) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeSSLPolicySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.selfLink, if present.
	selfLink := common.ValueOf(obj.Status.SelfLink)
	if selfLink != "" {
		statusIdentity := &ComputeSSLPolicyIdentity{}
		if err := statusIdentity.FromExternal(selfLink); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeSSLPolicy identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
