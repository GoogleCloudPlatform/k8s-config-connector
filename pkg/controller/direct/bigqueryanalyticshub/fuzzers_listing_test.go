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
	"math/rand"
	"reflect"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryanalyticshub/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/fuzz"
)

var ValidCategories = []string{
	"CATEGORY_UNSPECIFIED",
	"CATEGORY_OTHERS",
	"CATEGORY_ADVERTISING_AND_MARKETING",
	"CATEGORY_COMMERCE",
	"CATEGORY_CLIMATE_AND_ENVIRONMENT",
	"CATEGORY_DEMOGRAPHICS",
	"CATEGORY_ECONOMICS",
	"CATEGORY_EDUCATION",
	"CATEGORY_ENERGY",
	"CATEGORY_FINANCIAL",
	"CATEGORY_GAMING",
	"CATEGORY_GEOSPATIAL",
	"CATEGORY_HEALTHCARE_AND_LIFE_SCIENCE",
	"CATEGORY_MEDIA",
	"CATEGORY_PUBLIC_SECTOR",
	"CATEGORY_RETAIL",
	"CATEGORY_SPORTS",
	"CATEGORY_SCIENCE_AND_RESEARCH",
	"CATEGORY_TRANSPORTATION_AND_LOGISTICS",
	"CATEGORY_TRAVEL_AND_TOURISM",
}

// KRM ->ToProto-> proto ->FromProto-> KRM
func FuzzListingSpec(f *testing.F) {
	// spec
	f.Fuzz(func(t *testing.T, seed int64) {
		randStream := rand.New(rand.NewSource(seed))

		overrides := &fuzz.FillerConfig{
			FieldOverrides: map[string]fuzz.OverrideFiller{
				".Categories": func(t *testing.T, fieldName string, field reflect.Value) {
					// up to two fields can be set for this slice
					slice := reflect.MakeSlice(field.Type(), 2, 2)

					perms := randStream.Perm(len(ValidCategories))
					// invariants: len(perms) >> 2
					slice.Index(0).SetString(ValidCategories[perms[0]])
					slice.Index(1).SetString(ValidCategories[perms[1]])

					field.Set(slice)
				},
				".DiscoveryType": func(t *testing.T, fieldName string, field reflect.Value) {

					values := []string{"DISCOVERY_TYPE_PRIVATE", "DISCOVERY_TYPE_PUBLIC"}
					selectedValue := values[randStream.Intn(len(values))]
					field.Elem().SetString(selectedValue)
				},
			},
		}
		filler := fuzz.NewRandomFiller(&fuzz.FillerConfig{
			Stream:         randStream,
			FieldOverrides: overrides.FieldOverrides,
		})

		// krm1
		k1 := &krm.BigQueryAnalyticsHubListingSpec{}

		// fill
		filler.Fill(t, k1)

		// ToProto
		mapCtx := &direct.MapContext{}
		p := BigQueryAnalyticsHubListingSpec_ToProto(mapCtx, k1)
		if mapCtx.Err() != nil {
			t.Fatalf("error converting KRM to proto: %v \n KRM: %s", mapCtx.Err(), test.PrettyPrintJSON(t, k1))
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
			t.Logf("k1 =\n%v", test.PrettyPrintYAML(t, k1))
			t.Logf("k2 =\n%v", test.PrettyPrintYAML(t, k2))
			t.Errorf("roundtrip failed; diff:\n%s", diff)
		}
	})
}
