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
	_ identity.IdentityV2 = &CertificateManagerCertificateMapEntryIdentity{}
	_ identity.Resource   = &CertificateManagerCertificateMapEntry{}
)

var CertificateManagerCertificateMapEntryIdentityFormat = gcpurls.Template[CertificateManagerCertificateMapEntryIdentity]("certificatemanager.googleapis.com", "projects/{project}/locations/global/certificateMaps/{certificatemap}/certificateMapEntries/{certificatemapentry}")

// +k8s:deepcopy-gen=false

// CertificateManagerCertificateMapEntryIdentity is the identity of a GCP CertificateManagerCertificateMapEntry resource.
type CertificateManagerCertificateMapEntryIdentity struct {
	Project             string
	CertificateMap      string
	CertificateMapEntry string
}

func (i *CertificateManagerCertificateMapEntryIdentity) String() string {
	return CertificateManagerCertificateMapEntryIdentityFormat.ToString(*i)
}

func (i *CertificateManagerCertificateMapEntryIdentity) FromExternal(ref string) error {
	parsed, match, err := CertificateManagerCertificateMapEntryIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CertificateManagerCertificateMapEntry external=%q was not known (use %s): %w", ref, CertificateManagerCertificateMapEntryIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CertificateManagerCertificateMapEntry external=%q was not known (use %s)", ref, CertificateManagerCertificateMapEntryIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *CertificateManagerCertificateMapEntryIdentity) Host() string {
	return CertificateManagerCertificateMapEntryIdentityFormat.Host()
}

func (i *CertificateManagerCertificateMapEntryIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/global/certificateMaps/%s", i.Project, i.CertificateMap)
}

func getIdentityFromCertificateManagerCertificateMapEntrySpec(ctx context.Context, reader client.Reader, obj *CertificateManagerCertificateMapEntry) (*CertificateManagerCertificateMapEntryIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	mapRef := obj.Spec.MapRef
	if err := mapRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("resolving spec.mapRef: %w", err)
	}

	mapIDRaw, err := mapRef.ParseExternalToIdentity()
	if err != nil {
		return nil, fmt.Errorf("parsing mapRef: %w", err)
	}

	mapID, ok := mapIDRaw.(*CertificateManagerCertificateMapIdentity)
	if !ok {
		return nil, fmt.Errorf("expected CertificateManagerCertificateMapIdentity from mapRef")
	}

	identity := &CertificateManagerCertificateMapEntryIdentity{
		Project:             projectID,
		CertificateMap:      mapID.CertificateMap,
		CertificateMapEntry: resourceID,
	}
	return identity, nil
}

func (obj *CertificateManagerCertificateMapEntry) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromCertificateManagerCertificateMapEntrySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
