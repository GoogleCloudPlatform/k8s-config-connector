// Copyright 2024 Google LLC
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

package bigquerydataset

import (
	"encoding/json"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/fuzz"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func FuzzBigQueryDataSetSpec(f *testing.F) {

	f.Fuzz(func(t *testing.T, seed int64) {
		filler := fuzz.NewRandomFiller(seed, nil, nil)

		// krm1
		k1 := &krm.BigQueryDatasetSpec{}

		// fill
		filler.Fill(t, k1)

		// ToProto
		ctx := &direct.MapContext{}
		p := BigQueryDatasetSpec_ToAPI(ctx, k1, "test")
		if ctx.Err() != nil {
			t.Fatalf("error converting KRM to proto: %v \n KRM: %s", ctx.Err(), prettyPrint(t, k1))
		}
		k2 := BigQueryDatasetSpec_FromAPI(ctx, p)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to KRM: %v", ctx.Err())
		}
		// Using cmpopts.IgnoreFields to ignore specific fields in ProjectRef
		opts := cmp.Options{
			// project ref
			cmpopts.IgnoreFields(refs.ProjectRef{}, "External", "Name", "Namespace"),

			// other ref
			cmpopts.IgnoreFields(refs.KMSCryptoKeyRef{}, "External", "Name", "Namespace"),

			// unroundtrippable fields (for now)
			cmpopts.IgnoreFields(krm.BigQueryDatasetSpec{}, "ResourceID"),
			cmpopts.IgnoreFields(krm.BigQueryDatasetSpec{}, "ProjectRef"),
			cmpopts.IgnoreFields(krm.BigQueryDatasetSpec{}, "IsCaseInsensitive"),
		}
		// compare
		if diff := cmp.Diff(k1, k2, opts...); diff != "" {
			t.Logf("k1 = %v", k1)
			t.Logf("k2 = %v", k2)
			t.Errorf("roundtrip failed; diff:\n%s", diff)
		}
	})
}

func prettyPrint(t *testing.T, k *krm.BigQueryDatasetSpec) string {
	encoded, err := json.MarshalIndent(k, "", " ")
	if err != nil {
		t.Fatalf("error encoding to json: %v", err)
	}

	return string(encoded)
}
