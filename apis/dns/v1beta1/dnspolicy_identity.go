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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &DNSPolicyIdentity{}
	_ identity.Resource   = &DNSPolicy{}
)

var DNSPolicyIdentityFormat = gcpurls.Template[DNSPolicyIdentity]("dns.googleapis.com", "projects/{project}/policies/{policy}")

// DNSPolicyIdentity is the identity of a GCP DNSPolicy.
// +k8s:deepcopy-gen=false
type DNSPolicyIdentity struct {
	Project string
	Policy  string
}

func (i *DNSPolicyIdentity) String() string {
	return DNSPolicyIdentityFormat.ToString(*i)
}

func (i *DNSPolicyIdentity) FromExternal(ref string) error {
	parsed, match, err := DNSPolicyIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DNSPolicy external=%q was not known (use %s): %w", ref, DNSPolicyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DNSPolicy external=%q was not known (use %s)", ref, DNSPolicyIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DNSPolicyIdentity) Host() string {
	return DNSPolicyIdentityFormat.Host()
}

func getIdentityFromDNSPolicySpec(ctx context.Context, reader client.Reader, obj *DNSPolicy) (*DNSPolicyIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &DNSPolicyIdentity{
		Project: projectID,
		Policy:  resourceID,
	}
	return identity, nil
}

func (obj *DNSPolicy) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDNSPolicySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
