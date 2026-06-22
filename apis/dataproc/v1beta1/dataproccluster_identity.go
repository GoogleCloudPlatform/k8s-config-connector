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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &DataprocClusterIdentity{}
)

var (
	DataprocClusterIdentityFormat         = gcpurls.Template[DataprocClusterIdentity]("dataproc.googleapis.com", "v1/projects/{project}/regions/{region}/clusters/{cluster}")
	DataprocClusterIdentityFormatRelative = gcpurls.Template[DataprocClusterIdentity]("", "projects/{project}/regions/{region}/clusters/{cluster}")
)

// +k8s:deepcopy-gen=false
type DataprocClusterIdentity struct {
	Project string
	Region  string
	Cluster string
}

func (i *DataprocClusterIdentity) String() string {
	return DataprocClusterIdentityFormat.ToString(*i)
}

func (i *DataprocClusterIdentity) FromExternal(ref string) error {
	parsed, match, err := DataprocClusterIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DataprocCluster external=%q was not known (use %s): %w", ref, DataprocClusterIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		parsed, match, err = DataprocClusterIdentityFormatRelative.Parse(ref)
		if err != nil {
			return fmt.Errorf("format of DataprocCluster external=%q was not known (use %s): %w", ref, DataprocClusterIdentityFormat.CanonicalForm(), err)
		}
		if !match {
			return fmt.Errorf("format of DataprocCluster external=%q was not known (use %s or %s)", ref, DataprocClusterIdentityFormat.CanonicalForm(), DataprocClusterIdentityFormatRelative.CanonicalForm())
		}
	}

	*i = *parsed
	return nil
}

func (i *DataprocClusterIdentity) Host() string {
	return DataprocClusterIdentityFormat.Host()
}

func getIdentityFromDataprocClusterSpec(ctx context.Context, reader client.Reader, obj client.Object) (*DataprocClusterIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// DataprocCluster uses "location" in KRM but "regions" in GCP URL.
	region, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve region: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &DataprocClusterIdentity{
		Project: projectID,
		Region:  region,
		Cluster: resourceID,
	}
	return identity, nil
}

func GetIdentity(ctx context.Context, reader client.Reader, obj client.Object) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDataprocClusterSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	u, ok := obj.(*unstructured.Unstructured)
	if ok {
		externalRef, _, _ := unstructured.NestedString(u.Object, "status", "externalRef")
		if externalRef != "" {
			// Validate desired with actual
			statusIdentity := &DataprocClusterIdentity{}
			if err := statusIdentity.FromExternal(externalRef); err != nil {
				return nil, err
			}

			if statusIdentity.String() != specIdentity.String() {
				return nil, fmt.Errorf("cannot change DataprocCluster identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
			}
		}
	}

	return specIdentity, nil
}
