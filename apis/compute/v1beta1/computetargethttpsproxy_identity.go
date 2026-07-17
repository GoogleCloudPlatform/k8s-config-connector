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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeTargetHTTPSProxyIdentity{}
	_ identity.Resource   = &ComputeTargetHTTPSProxy{}
)

var ComputeGlobalTargetHTTPSProxyIdentityFormat = gcpurls.Template[ComputeTargetHTTPSProxyIdentity]("compute.googleapis.com", "projects/{project}/global/targetHttpsProxies/{targetHttpsProxy}")
var ComputeRegionalTargetHTTPSProxyIdentityFormat = gcpurls.Template[ComputeTargetHTTPSProxyIdentity]("compute.googleapis.com", "projects/{project}/regions/{region}/targetHttpsProxies/{targetHttpsProxy}")

// ComputeTargetHTTPSProxyIdentity is the identity of a GCP ComputeTargetHTTPSProxy resource.
// +k8s:deepcopy-gen=false
type ComputeTargetHTTPSProxyIdentity struct {
	Project          string
	Region           string
	TargetHttpsProxy string
}

func (i *ComputeTargetHTTPSProxyIdentity) IsGlobal() bool {
	return i.Region == "" || i.Region == "global"
}

func (i *ComputeTargetHTTPSProxyIdentity) String() string {
	if !i.IsGlobal() {
		return ComputeRegionalTargetHTTPSProxyIdentityFormat.ToString(*i)
	}
	return ComputeGlobalTargetHTTPSProxyIdentityFormat.ToString(*i)
}

func (i *ComputeTargetHTTPSProxyIdentity) FromExternal(ref string) error {
	ref = apirefs.TrimComputeURIPrefix(ref)

	if parsed, match, _ := ComputeGlobalTargetHTTPSProxyIdentityFormat.Parse(ref); match {
		*i = *parsed
		i.Region = "global"
		return nil
	}
	if parsed, match, _ := ComputeRegionalTargetHTTPSProxyIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of ComputeTargetHTTPSProxy external=%q was not known (use %s or %s)", ref, ComputeGlobalTargetHTTPSProxyIdentityFormat.CanonicalForm(), ComputeRegionalTargetHTTPSProxyIdentityFormat.CanonicalForm())
}

func (i *ComputeTargetHTTPSProxyIdentity) Host() string {
	return ComputeGlobalTargetHTTPSProxyIdentityFormat.Host()
}

func (i *ComputeTargetHTTPSProxyIdentity) ParentString() string {
	if !i.IsGlobal() {
		return fmt.Sprintf("projects/%s/regions/%s", i.Project, i.Region)
	}
	return fmt.Sprintf("projects/%s/global", i.Project)
}

func ParseComputeTargetHTTPSProxyExternal(external string) (*ComputeTargetHTTPSProxyIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeTargetHTTPSProxy external value")
	}
	id := &ComputeTargetHTTPSProxyIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeTargetHTTPSProxySpec(ctx context.Context, reader client.Reader, obj *ComputeTargetHTTPSProxy) (*ComputeTargetHTTPSProxyIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	location := "global"
	if obj.Spec.Location != nil {
		location = common.ValueOf(obj.Spec.Location)
	}

	identity := &ComputeTargetHTTPSProxyIdentity{
		Project:          projectID,
		Region:           location,
		TargetHttpsProxy: resourceID,
	}
	return identity, nil
}

func (obj *ComputeTargetHTTPSProxy) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeTargetHTTPSProxySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.externalRef or status.selfLink, if present.
	// We'll check externalRef first, then fallback to selfLink.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		externalRef = common.ValueOf(obj.Status.SelfLink)
	}
	if externalRef != "" {
		statusIdentity := &ComputeTargetHTTPSProxyIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeTargetHTTPSProxy identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
