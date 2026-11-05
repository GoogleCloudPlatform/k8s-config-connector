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
	_ identity.IdentityV2 = &ComputeManagedSSLCertificateIdentity{}
	_ identity.Resource   = &ComputeManagedSSLCertificate{}
)

var ComputeManagedSSLCertificateIdentityFormat = gcpurls.Template[ComputeManagedSSLCertificateIdentity]("compute.googleapis.com", "projects/{project}/global/sslCertificates/{sslCertificate}")

// ComputeManagedSSLCertificateIdentity is the identity of a GCP ComputeManagedSSLCertificate resource.
// +k8s:deepcopy-gen=false
type ComputeManagedSSLCertificateIdentity struct {
	Project        string
	SslCertificate string
}

func (i *ComputeManagedSSLCertificateIdentity) String() string {
	return ComputeManagedSSLCertificateIdentityFormat.ToString(*i)
}

func (i *ComputeManagedSSLCertificateIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeManagedSSLCertificateIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeManagedSSLCertificate external=%q was not known (use %s): %w", ref, ComputeManagedSSLCertificateIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeManagedSSLCertificate external=%q was not known (use %s)", ref, ComputeManagedSSLCertificateIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeManagedSSLCertificateIdentity) Host() string {
	return ComputeManagedSSLCertificateIdentityFormat.Host()
}

func (i *ComputeManagedSSLCertificateIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/global", i.Project)
}

func ParseComputeManagedSSLCertificateExternal(external string) (*ComputeManagedSSLCertificateIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeManagedSSLCertificate external value")
	}
	id := &ComputeManagedSSLCertificateIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeManagedSSLCertificateSpec(ctx context.Context, reader client.Reader, obj *ComputeManagedSSLCertificate) (*ComputeManagedSSLCertificateIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &ComputeManagedSSLCertificateIdentity{
		Project:        projectID,
		SslCertificate: resourceID,
	}
	return identity, nil
}

func (obj *ComputeManagedSSLCertificate) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeManagedSSLCertificateSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.externalRef or status.observedState.selfLink, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" && obj.Status.ObservedState != nil {
		externalRef = common.ValueOf(obj.Status.ObservedState.SelfLink)
	}
	if externalRef != "" {
		statusIdentity := &ComputeManagedSSLCertificateIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeManagedSSLCertificate identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
