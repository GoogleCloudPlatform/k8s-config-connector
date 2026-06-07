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

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ refs.Ref = &CertificateDNSAuthorizationRef{}
	_ refs.Ref = &CertificateIssuanceConfigRef{}
)

var CertificateManagerCertificateIssuanceConfigGVK = schema.GroupVersionKind{
	Group:   "certificatemanager.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "CertificateManagerCertificateIssuanceConfig",
}

type certificateManagerDNSAuthorizationIdentity struct {
	Project          string
	Location         string
	DNSAuthorization string
}

var certificateManagerDNSAuthorizationIdentityFormat = gcpurls.Template[certificateManagerDNSAuthorizationIdentity]("certificatemanager.googleapis.com", "projects/{project}/locations/{location}/dnsAuthorizations/{dnsauthorization}")

type certificateManagerCertificateIssuanceConfigIdentity struct {
	Project                   string
	Location                  string
	CertificateIssuanceConfig string
}

var certificateManagerCertificateIssuanceConfigIdentityFormat = gcpurls.Template[certificateManagerCertificateIssuanceConfigIdentity]("certificatemanager.googleapis.com", "projects/{project}/locations/{location}/certificateIssuanceConfigs/{certificateissuanceconfig}")

// CertificateDNSAuthorizationRef implementation

func (r *CertificateDNSAuthorizationRef) GetGVK() schema.GroupVersionKind {
	return CertificateManagerDNSAuthorizationGVK
}

func (r *CertificateDNSAuthorizationRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *CertificateDNSAuthorizationRef) GetExternal() string {
	return r.External
}

func (r *CertificateDNSAuthorizationRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *CertificateDNSAuthorizationRef) ValidateExternal(ref string) error {
	_, match, err := certificateManagerDNSAuthorizationIdentityFormat.Parse(ref)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("format of external=%q was not known (use %s)", ref, certificateManagerDNSAuthorizationIdentityFormat.CanonicalForm())
	}
	return nil
}

func (r *CertificateDNSAuthorizationRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		resourceID, err := refs.GetResourceID(u)
		if err != nil {
			return ""
		}
		location, err := refs.GetLocation(u)
		if err != nil {
			return ""
		}
		if location == "" {
			location = "global"
		}
		projectID, err := refs.ResolveProjectID(ctx, reader, u)
		if err != nil {
			return ""
		}
		return fmt.Sprintf("//certificatemanager.googleapis.com/projects/%s/locations/%s/dnsAuthorizations/%s", projectID, location, resourceID)
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

// CertificateIssuanceConfigRef implementation

func (r *CertificateIssuanceConfigRef) GetGVK() schema.GroupVersionKind {
	return CertificateManagerCertificateIssuanceConfigGVK
}

func (r *CertificateIssuanceConfigRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *CertificateIssuanceConfigRef) GetExternal() string {
	return r.External
}

func (r *CertificateIssuanceConfigRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *CertificateIssuanceConfigRef) ValidateExternal(ref string) error {
	_, match, err := certificateManagerCertificateIssuanceConfigIdentityFormat.Parse(ref)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("format of external=%q was not known (use %s)", ref, certificateManagerCertificateIssuanceConfigIdentityFormat.CanonicalForm())
	}
	return nil
}

func (r *CertificateIssuanceConfigRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		resourceID, err := refs.GetResourceID(u)
		if err != nil {
			return ""
		}
		location, err := refs.GetLocation(u)
		if err != nil {
			return ""
		}
		if location == "" {
			location = "global"
		}
		projectID, err := refs.ResolveProjectID(ctx, reader, u)
		if err != nil {
			return ""
		}
		return fmt.Sprintf("//certificatemanager.googleapis.com/projects/%s/locations/%s/certificateIssuanceConfigs/%s", projectID, location, resourceID)
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

func init() {
	refs.Register(&CertificateDNSAuthorizationRef{})
	refs.Register(&CertificateIssuanceConfigRef{})
}
