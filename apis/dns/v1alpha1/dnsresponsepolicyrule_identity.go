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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &DNSResponsePolicyRuleIdentity{}
	_ identity.Resource   = &DNSResponsePolicyRule{}
)

var (
	DNSResponsePolicyRuleIdentityFormat         = gcpurls.Template[DNSResponsePolicyRuleIdentity]("dns.googleapis.com", "projects/{project}/locations/{location}/responsePolicies/{responsePolicy}/rules/{rule}")
	DNSResponsePolicyRuleIdentityFallbackFormat = gcpurls.Template[DNSResponsePolicyRuleIdentity]("dns.googleapis.com", "projects/{project}/responsePolicies/{responsePolicy}/rules/{rule}")
)

// DNSResponsePolicyRuleIdentity is the identity of a GCP DNSResponsePolicyRule.
// +k8s:deepcopy-gen=false
type DNSResponsePolicyRuleIdentity struct {
	Project        string
	Location       string
	ResponsePolicy string
	Rule           string
}

func (i *DNSResponsePolicyRuleIdentity) String() string {
	if i.Location != "" {
		return DNSResponsePolicyRuleIdentityFormat.ToString(*i)
	}
	return DNSResponsePolicyRuleIdentityFallbackFormat.ToString(*i)
}

func (i *DNSResponsePolicyRuleIdentity) FromExternal(ref string) error {
	if parsed, match, err := DNSResponsePolicyRuleIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	} else if err != nil {
		return err
	}
	if parsed, match, err := DNSResponsePolicyRuleIdentityFallbackFormat.Parse(ref); match {
		*i = *parsed
		return nil
	} else if err != nil {
		return err
	}
	return fmt.Errorf("format of DNSResponsePolicyRule external=%q was not known (use %s or %s)", ref, DNSResponsePolicyRuleIdentityFormat.CanonicalForm(), DNSResponsePolicyRuleIdentityFallbackFormat.CanonicalForm())
}

func (i *DNSResponsePolicyRuleIdentity) Host() string {
	return DNSResponsePolicyRuleIdentityFormat.Host()
}

func getIdentityFromDNSResponsePolicyRuleSpec(ctx context.Context, reader client.Reader, obj *DNSResponsePolicyRule) (*DNSResponsePolicyRuleIdentity, error) {
	ruleID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	responsePolicyStr := obj.Spec.ResponsePolicy
	if responsePolicyStr == "" {
		return nil, fmt.Errorf("spec.responsePolicy is required")
	}

	var responsePolicyID string
	var responsePolicyLocation string

	policyIdentity := &DNSResponsePolicyIdentity{}
	if err := policyIdentity.FromExternal(responsePolicyStr); err != nil {
		return nil, fmt.Errorf("invalid responsePolicy: %w", err)
	}

	if policyIdentity.Project != projectID {
		return nil, fmt.Errorf("responsePolicy project %q must match project %q", policyIdentity.Project, projectID)
	}
	responsePolicyID = policyIdentity.ResponsePolicy
	responsePolicyLocation = policyIdentity.Location

	identity := &DNSResponsePolicyRuleIdentity{
		Project:        projectID,
		Location:       responsePolicyLocation,
		ResponsePolicy: responsePolicyID,
		Rule:           ruleID,
	}
	return identity, nil
}

func (obj *DNSResponsePolicyRule) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDNSResponsePolicyRuleSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
