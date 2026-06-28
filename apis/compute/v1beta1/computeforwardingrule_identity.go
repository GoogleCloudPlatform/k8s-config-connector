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
	_ identity.IdentityV2 = &ComputeForwardingRuleIdentity{}
	_ identity.Resource   = &ComputeForwardingRule{}
)

var ComputeGlobalForwardingRuleIdentityFormat = gcpurls.Template[ComputeForwardingRuleIdentity]("compute.googleapis.com", "projects/{project}/global/forwardingRules/{forwardingrule}")
var ComputeRegionalForwardingRuleIdentityFormat = gcpurls.Template[ComputeForwardingRuleIdentity]("compute.googleapis.com", "projects/{project}/regions/{region}/forwardingRules/{forwardingrule}")

// ComputeForwardingRuleIdentity is the identity of a GCP ComputeForwardingRule resource.
// +k8s:deepcopy-gen=false
type ComputeForwardingRuleIdentity struct {
	Project        string
	Region         string
	ForwardingRule string
}

func (i *ComputeForwardingRuleIdentity) IsGlobal() bool {
	return i.Region == "" || i.Region == "global"
}

func (i *ComputeForwardingRuleIdentity) String() string {
	if !i.IsGlobal() {
		return ComputeRegionalForwardingRuleIdentityFormat.ToString(*i)
	}
	return ComputeGlobalForwardingRuleIdentityFormat.ToString(*i)
}

func (i *ComputeForwardingRuleIdentity) FromExternal(ref string) error {
	ref = apirefs.TrimComputeURIPrefix(ref)

	if parsed, match, _ := ComputeGlobalForwardingRuleIdentityFormat.Parse(ref); match {
		*i = *parsed
		i.Region = "global"
		return nil
	}
	if parsed, match, _ := ComputeRegionalForwardingRuleIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of ComputeForwardingRule external=%q was not known (use %s or %s)", ref, ComputeGlobalForwardingRuleIdentityFormat.CanonicalForm(), ComputeRegionalForwardingRuleIdentityFormat.CanonicalForm())
}

func (i *ComputeForwardingRuleIdentity) Host() string {
	return ComputeGlobalForwardingRuleIdentityFormat.Host()
}

func (i *ComputeForwardingRuleIdentity) ParentString() string {
	if !i.IsGlobal() {
		return fmt.Sprintf("projects/%s/regions/%s", i.Project, i.Region)
	}
	return fmt.Sprintf("projects/%s/global", i.Project)
}

func getIdentityFromComputeForwardingRuleSpec(ctx context.Context, reader client.Reader, obj *ComputeForwardingRule) (*ComputeForwardingRuleIdentity, error) {
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
		location = *obj.Spec.Location
	}

	identity := &ComputeForwardingRuleIdentity{
		Project:        projectID,
		Region:         location,
		ForwardingRule: resourceID,
	}
	return identity, nil
}

func (obj *ComputeForwardingRule) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeForwardingRuleSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.SelfLink)
	if externalRef == "" {
		externalRef = common.ValueOf(obj.Status.ExternalRef)
	}
	if externalRef != "" {
		statusIdentity := &ComputeForwardingRuleIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeForwardingRule identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
