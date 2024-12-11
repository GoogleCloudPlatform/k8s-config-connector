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

package bigqueryanalyticshub

import (
	"encoding/json"
	"reflect"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryanalyticshub/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"golang.org/x/exp/rand"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/fuzz"
)

// KRM ->ToProto-> proto ->FromProto-> KRM
func FuzzListingSpec(f *testing.F) {
	// spec
	f.Fuzz(func(t *testing.T, seed int64) {
		randStream := rand.New(rand.NewSource(uint64(seed)))
		stringEnumAllowableValues := map[string][]interface{}{
			"MyStringEnumType": {"CATEGORY_DEMOGRAPHICS"},
		}
		overrides := &fuzz.FillerConfig{
			StringEnumAllowableValues: stringEnumAllowableValues,
			FieldOverrides: map[string]fuzz.OverrideFiller{
				".Categories": func(t *testing.T, fieldName string, field reflect.Value) {
					// up to two fields can be set for this slice

					slice := reflect.MakeSlice(field.Type(), 2, 2)
					slice.Index(0).SetString("CATEGORY_DEMOGRAPHICS")
					slice.Index(1).SetString("CATEGORY_MEDIA") // todo acpana make random
					field.Set(slice)
				},
				".DiscoveryType": func(t *testing.T, fieldName string, field reflect.Value) {
					if field.Kind() == reflect.Ptr {
						if field.IsNil() {
							field.Set(reflect.New(field.Type().Elem()))
						}
					}
					values := []string{"DISCOVERY_TYPE_PRIVATE", "DISCOVERY_TYPE_PUBLIC"}
					selectedValue := values[randStream.Intn(len(values))]
					field.Elem().SetString(selectedValue)
				},
			},
		}
		filler := fuzz.NewRandomFillerWithConfig(seed, overrides)

		// krm1
		k1 := &krm.BigQueryAnalyticsHubListingSpec{}

		// fill
		filler.Fill(t, k1)

		// ToProto
		mapCtx := &direct.MapContext{}
		p := BigQueryAnalyticsHubListingSpec_ToProto(mapCtx, k1)
		if mapCtx.Err() != nil {
			t.Fatalf("error converting KRM to proto: %v \n KRM: %s", mapCtx.Err(), prettyPrint(t, k1))
		}

		// krm2 : FromProto
		mapCtx = &direct.MapContext{}
		k2 := BigQueryAnalyticsHubListingSpec_FromProto(mapCtx, p)
		if mapCtx.Err() != nil {
			t.Fatalf("error mapping from proto to KRM: %v", mapCtx.Err())
		}

		// Using cmpopts.IgnoreFields to ignore specific fields in ProjectRef
		opts := cmp.Options{
			cmpopts.IgnoreFields(krm.BigQueryAnalyticsHubListingSpec{}, "Location"),
			cmpopts.IgnoreFields(krm.BigQueryAnalyticsHubListingSpec{}, "ResourceID"),
			cmpopts.IgnoreFields(krm.BigQueryAnalyticsHubListingSpec{}, "ProjectRef"),
			cmpopts.IgnoreFields(krm.BigQueryAnalyticsHubListingSpec{}, "DataExchangeRef"),
			cmpopts.IgnoreFields(refs.BigQueryTableRef{}, "Name", "Namespace"),
			cmpopts.IgnoreFields(krm.BigQueryDatasetSource{}, "DatasetRef"),
			cmpopts.IgnoreFields(refs.BigQueryDatasetRef{}, "External", "Name", "Namespace"),
		}

		// compare
		if diff := cmp.Diff(k1, k2, opts...); diff != "" {
			t.Logf("k1 = %v", k1)
			t.Logf("k2 = %v", k2)
			t.Errorf("roundtrip failed; diff:\n%s", diff)
		}
	})
}

func prettyPrint(t *testing.T, k *krm.BigQueryAnalyticsHubListingSpec) string {
	encoded, err := json.MarshalIndent(k, "", " ")
	if err != nil {
		t.Fatalf("error encoding to json: %v", err)
	}

	return string(encoded)
}
