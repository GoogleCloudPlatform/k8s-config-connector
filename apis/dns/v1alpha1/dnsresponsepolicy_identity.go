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
	_ identity.IdentityV2 = &DNSResponsePolicyIdentity{}
	_ identity.Resource   = &DNSResponsePolicy{}
)

var (
	DNSResponsePolicyIdentityFormat         = gcpurls.Template[DNSResponsePolicyIdentity]("dns.googleapis.com", "projects/{project}/locations/{location}/responsePolicies/{responsePolicy}")
	DNSResponsePolicyIdentityFallbackFormat = gcpurls.Template[DNSResponsePolicyIdentity]("dns.googleapis.com", "projects/{project}/responsePolicies/{responsePolicy}")
)

// DNSResponsePolicyIdentity is the identity of a GCP DNSResponsePolicy.
// +k8s:deepcopy-gen=false
type DNSResponsePolicyIdentity struct {
	Project        string
	Location       string
	ResponsePolicy string
}

func (i *DNSResponsePolicyIdentity) String() string {
	if i.Location != "" {
		return DNSResponsePolicyIdentityFormat.ToString(*i)
	}
	return DNSResponsePolicyIdentityFallbackFormat.ToString(*i)
}

func (i *DNSResponsePolicyIdentity) FromExternal(ref string) error {
	if parsed, match, err := DNSResponsePolicyIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	} else if err != nil {
		return err
	}
	if parsed, match, err := DNSResponsePolicyIdentityFallbackFormat.Parse(ref); match {
		*i = *parsed
		return nil
	} else if err != nil {
		return err
	}
	return fmt.Errorf("format of DNSResponsePolicy external=%q was not known (use %s or %s)", ref, DNSResponsePolicyIdentityFormat.CanonicalForm(), DNSResponsePolicyIdentityFallbackFormat.CanonicalForm())
}

func (i *DNSResponsePolicyIdentity) Host() string {
	return DNSResponsePolicyIdentityFormat.Host()
}

func getIdentityFromDNSResponsePolicySpec(ctx context.Context, reader client.Reader, obj *DNSResponsePolicy) (*DNSResponsePolicyIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &DNSResponsePolicyIdentity{
		Project:        projectID,
		Location:       "global",
		ResponsePolicy: resourceID,
	}
	return identity, nil
}

func (obj *DNSResponsePolicy) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDNSResponsePolicySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
