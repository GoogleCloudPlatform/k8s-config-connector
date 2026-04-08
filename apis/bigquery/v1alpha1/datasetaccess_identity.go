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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// BigQueryDatasetAccessIdentityURL is the format for the externalRef of a BigQueryDatasetAccess.
	BigQueryDatasetAccessIdentityURL = "projects/{project}/datasets/{dataset}/access/{accessID}"
)

var (
	_ identity.IdentityV2 = &BigQueryDatasetAccessIdentity{}
	_ identity.Resource   = &BigQueryDatasetAccess{}
)

var BigQueryDatasetAccessIdentityFormat = gcpurls.Template[BigQueryDatasetAccessIdentity](
	"bigquery.googleapis.com",
	BigQueryDatasetAccessIdentityURL,
)

// BigQueryDatasetAccessIdentity represents the identity of a BigQueryDatasetAccess.
// +k8s:deepcopy-gen=false
type BigQueryDatasetAccessIdentity struct {
	Project  string
	Dataset  string
	AccessID string
}

func (i *BigQueryDatasetAccessIdentity) String() string {
	return BigQueryDatasetAccessIdentityFormat.ToString(*i)
}

func (i *BigQueryDatasetAccessIdentity) FromExternal(ref string) error {
	parsed, match, err := BigQueryDatasetAccessIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BigQueryDatasetAccess external=%q was not known (use %s): %w", ref, BigQueryDatasetAccessIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BigQueryDatasetAccess external=%q was not known (use %s)", ref, BigQueryDatasetAccessIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BigQueryDatasetAccessIdentity) Host() string {
	return BigQueryDatasetAccessIdentityFormat.Host()
}

func getIdentityFromBigQueryDatasetAccessSpec(ctx context.Context, reader client.Reader, obj client.Object) (*BigQueryDatasetAccessIdentity, error) {
	u, ok := obj.(*unstructured.Unstructured)
	if !ok {
		m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			return nil, fmt.Errorf("cannot convert to unstructured: %w", err)
		}
		u = &unstructured.Unstructured{Object: m}
	}

	resourceID, err := v1beta1.GetResourceID(u)
	if err != nil {
		return nil, err
	}

	datasetID, _, _ := unstructured.NestedString(u.Object, "spec", "datasetId")
	if datasetID == "" {
		return nil, fmt.Errorf("spec.datasetId not set")
	}

	projectID, err := v1beta1.ResolveProjectID(ctx, reader, u)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	return &BigQueryDatasetAccessIdentity{
		Project:  projectID,
		Dataset:  datasetID,
		AccessID: resourceID,
	}, nil
}

func (obj *BigQueryDatasetAccess) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBigQueryDatasetAccessSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
