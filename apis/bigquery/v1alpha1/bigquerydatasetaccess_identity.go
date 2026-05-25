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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &BigQueryDatasetAccessIdentity{}
	_ identity.Resource   = &BigQueryDatasetAccess{}
)

var BigQueryDatasetAccessIdentityFormat = gcpurls.Template[BigQueryDatasetAccessIdentity]("bigquery.googleapis.com", "projects/{project}/datasets/{dataset}")

// BigQueryDatasetAccessIdentity defines the resource identifier for BigQueryDatasetAccess.
// +k8s:deepcopy-gen=false
type BigQueryDatasetAccessIdentity struct {
	Project string
	Dataset string
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
	bigQueryDatasetAccess, err := TypedCopy[BigQueryDatasetAccess](obj)
	if err != nil {
		return nil, err
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, bigQueryDatasetAccess)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	return &BigQueryDatasetAccessIdentity{
		Project: projectID,
		Dataset: bigQueryDatasetAccess.Spec.DatasetId,
	}, nil
}

// GetIdentity builds a BigQueryDatasetAccessIdentity from the Config Connector BigQueryDatasetAccess object.
func (obj *BigQueryDatasetAccess) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBigQueryDatasetAccessSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	return specIdentity, nil
}

func TypedCopy[T any](obj client.Object) (*T, error) {
	if val, ok := any(obj).(*T); ok {
		if ro, ok := any(val).(runtime.Object); ok {
			return any(ro.DeepCopyObject()).(*T), nil
		}
	}
	u, ok := obj.(*unstructured.Unstructured)
	if !ok {
		return nil, fmt.Errorf("expected %T or *unstructured.Unstructured, got %T", *new(T), obj)
	}
	res := new(T)
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, res); err != nil {
		return nil, fmt.Errorf("error converting unstructured to %T: %w", res, err)
	}
	return res, nil
}
