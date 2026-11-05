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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.ServerGeneratedIdentity = &SQLSSLCertIdentity{}
	_ identity.Resource                = &SQLSSLCert{}
)

var SQLSSLCertIdentityFormat = gcpurls.Template[SQLSSLCertIdentity]("sqladmin.googleapis.com", "projects/{project}/instances/{instance}/sslCerts/{sha1Fingerprint}")

// SQLSSLCertIdentity is the identity of a GCP SQLSSLCert resource.
// +k8s:deepcopy-gen=false
type SQLSSLCertIdentity struct {
	Project         string
	Instance        string
	Sha1Fingerprint string
}

func (i *SQLSSLCertIdentity) String() string {
	return SQLSSLCertIdentityFormat.ToString(*i)
}

func (i *SQLSSLCertIdentity) FromExternal(ref string) error {
	ref = strings.TrimPrefix(ref, "https://")
	ref = strings.TrimPrefix(ref, "http://")
	ref = strings.TrimPrefix(ref, "//")
	ref = strings.TrimPrefix(ref, "sqladmin.googleapis.com/")
	ref = strings.TrimPrefix(ref, "sql/v1beta4/")

	parsed, match, err := SQLSSLCertIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of SQLSSLCert external=%q was not known (use %s): %w", ref, SQLSSLCertIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of SQLSSLCert external=%q was not known (use %s)", ref, SQLSSLCertIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *SQLSSLCertIdentity) Host() string {
	return SQLSSLCertIdentityFormat.Host()
}

func (i *SQLSSLCertIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/instances/%s", i.Project, i.Instance)
}

func (i *SQLSSLCertIdentity) HasIdentitySpecified() bool {
	return i.Sha1Fingerprint != ""
}

func getIdentityFromSQLSSLCertSpec(ctx context.Context, reader client.Reader, obj *SQLSSLCert) (*SQLSSLCertIdentity, error) {
	// For SQLSSLCert, sha1Fingerprint is a server-generated ID and is optional for creation.
	// We retrieve it directly from Spec.ResourceID without falling back to GetName().
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	instanceRef := obj.Spec.InstanceRef
	instance, err := refsv1beta1.ResolveSQLInstanceRef(ctx, reader, obj, &instanceRef)
	if err != nil {
		return nil, fmt.Errorf("resolving spec.instanceRef: %w", err)
	}
	if instance == nil {
		return nil, fmt.Errorf("spec.instanceRef is required")
	}

	if instance.SQLInstanceName == "" {
		return nil, fmt.Errorf("cannot resolve instance name from spec.instanceRef")
	}

	projectID := instance.ProjectID
	if projectID == "" {
		projectID, err = refsv1beta1.ResolveProjectID(ctx, reader, obj)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve project: %w", err)
		}
	}

	return &SQLSSLCertIdentity{
		Project:         projectID,
		Instance:        instance.SQLInstanceName,
		Sha1Fingerprint: resourceID,
	}, nil
}

func (obj *SQLSSLCert) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromSQLSSLCertSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	statusSha1 := common.ValueOf(obj.Status.Sha1Fingerprint)
	if statusSha1 != "" {
		if specIdentity.Sha1Fingerprint == "" {
			specIdentity.Sha1Fingerprint = statusSha1
		} else if specIdentity.Sha1Fingerprint != statusSha1 {
			return nil, fmt.Errorf("cannot change SQLSSLCert identity (old=%q, new=%q)", statusSha1, specIdentity.Sha1Fingerprint)
		}
	}

	return specIdentity, nil
}
