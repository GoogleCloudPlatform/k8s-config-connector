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
	_ identity.IdentityV2 = &DNSRecordSetIdentity{}
	_ identity.Resource   = &DNSRecordSet{}
)

var (
	DNSRecordSetIdentityFormat         = gcpurls.Template[DNSRecordSetIdentity]("dns.googleapis.com", "projects/{project}/locations/{location}/managedZones/{managedZone}/rrsets/{name}")
	DNSRecordSetIdentityFallbackFormat = gcpurls.Template[DNSRecordSetIdentity]("dns.googleapis.com", "projects/{project}/managedZones/{managedZone}/rrsets/{name}")
)

// DNSRecordSetIdentity is the identity of a GCP DNSRecordSet resource.
// +k8s:deepcopy-gen=false
type DNSRecordSetIdentity struct {
	Project     string
	Location    string
	ManagedZone string
	Name        string
}

func (i *DNSRecordSetIdentity) String() string {
	if i.Location != "" {
		return DNSRecordSetIdentityFormat.ToString(*i)
	}
	return DNSRecordSetIdentityFallbackFormat.ToString(*i)
}

func (i *DNSRecordSetIdentity) FromExternal(ref string) error {
	if parsed, match, err := DNSRecordSetIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	} else if err != nil {
		return err
	}
	if parsed, match, err := DNSRecordSetIdentityFallbackFormat.Parse(ref); match {
		*i = *parsed
		return nil
	} else if err != nil {
		return err
	}
	return fmt.Errorf("format of DNSRecordSet external=%q was not known (use %s or %s)", ref, DNSRecordSetIdentityFormat.CanonicalForm(), DNSRecordSetIdentityFallbackFormat.CanonicalForm())
}

func (i *DNSRecordSetIdentity) Host() string {
	return DNSRecordSetIdentityFormat.Host()
}

func getIdentityFromDNSRecordSetSpec(ctx context.Context, reader client.Reader, obj *DNSRecordSet) (*DNSRecordSetIdentity, error) {
	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	managedZoneRef := obj.Spec.ManagedZoneRef.DeepCopy()
	if err := managedZoneRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("cannot resolve managedZoneRef: %w", err)
	}
	managedZoneIdentityRaw, err := managedZoneRef.ParseExternalToIdentity()
	if err != nil {
		return nil, fmt.Errorf("cannot parse managedZoneRef: %w", err)
	}
	managedZoneIdentity := managedZoneIdentityRaw.(*DNSManagedZoneIdentity)

	if managedZoneIdentity.Project != projectID {
		return nil, fmt.Errorf("managedZoneRef project %q must match project %q", managedZoneIdentity.Project, projectID)
	}

	if obj.Spec.Name == "" {
		return nil, fmt.Errorf("spec.name is required")
	}

	if obj.Spec.Type == "" {
		return nil, fmt.Errorf("spec.type is required")
	}

	identity := &DNSRecordSetIdentity{
		Project:     projectID,
		ManagedZone: managedZoneIdentity.ManagedZone,
		Name:        obj.Spec.Name,
	}
	return identity, nil
}

func (obj *DNSRecordSet) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDNSRecordSetSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
