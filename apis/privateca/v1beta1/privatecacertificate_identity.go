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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/privateca/privatecarefs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &PrivateCACertificateIdentity{}
	_ identity.Resource   = &PrivateCACertificate{}
)

var PrivateCACertificateIdentityFormat = gcpurls.Template[PrivateCACertificateIdentity]("privateca.googleapis.com", "projects/{project}/locations/{location}/caPools/{caPool}/certificates/{certificate}")

// PrivateCACertificateIdentity is the identity of a GCP PrivateCACertificate resource.
// +k8s:deepcopy-gen=false
type PrivateCACertificateIdentity struct {
	Project     string
	Location    string
	CAPool      string
	Certificate string
}

func (i *PrivateCACertificateIdentity) String() string {
	return PrivateCACertificateIdentityFormat.ToString(*i)
}

func (i *PrivateCACertificateIdentity) FromExternal(ref string) error {
	parsed, match, err := PrivateCACertificateIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of PrivateCACertificate external=%q was not known (use %s): %w", ref, PrivateCACertificateIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of PrivateCACertificate external=%q was not known (use %s)", ref, PrivateCACertificateIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *PrivateCACertificateIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/caPools/%s", i.Project, i.Location, i.CAPool)
}

func (i *PrivateCACertificateIdentity) Host() string {
	return PrivateCACertificateIdentityFormat.Host()
}

func getIdentityFromPrivateCACertificateSpec(ctx context.Context, reader client.Reader, obj *PrivateCACertificate) (*PrivateCACertificateIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	if err := obj.Spec.CaPoolRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("cannot resolve caPoolRef: %w", err)
	}
	caPoolExternal := obj.Spec.CaPoolRef.External
	caPoolIdentity := &privatecarefs.PrivateCACAPoolIdentity{}
	if err := caPoolIdentity.FromExternal(caPoolExternal); err != nil {
		return nil, fmt.Errorf("cannot parse caPoolRef external string %q: %w", caPoolExternal, err)
	}

	identity := &PrivateCACertificateIdentity{
		Project:     projectID,
		Location:    location,
		CAPool:      caPoolIdentity.CAPool,
		Certificate: resourceID,
	}
	return identity, nil
}

func (obj *PrivateCACertificate) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromPrivateCACertificateSpec(ctx, reader, obj)
}
