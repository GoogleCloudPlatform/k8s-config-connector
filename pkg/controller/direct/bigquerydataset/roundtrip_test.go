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
	"math/rand"
	"reflect"
	"strconv"
	"testing"

	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/fuzz"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var accessFieldNames = []string{".UserByEmail", ",GroupByEmail", ".Domain", ".IamMember", ".SpecialGroup", ".View", ".Routine", ".Dataset"}

func FuzzBigQueryDataSetSpec(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		stream := rand.New(rand.NewSource((seed)))

		// this is a pointer string field that needs an int.
		// the value is between 48-168 hours.
		funcF := func(t *testing.T, fieldName string, field reflect.Value) {
			maxT := rand.Intn(121) + 48
			field.SetString(strconv.Itoa(maxT))
		}
		funcMs := func(t *testing.T, fieldName string, field reflect.Value) {
			field.SetInt((stream.Int63()) / 1e6)
		}
		// only one field in the access struct should be set.
		funcAccess := func(t *testing.T, fieldName string, field reflect.Value) {
			count := rand.Intn(10) + 1
			slice := reflect.MakeSlice(field.Type(), count, count)
			for j := 0; j < count; j++ {
				accessRand := rand.Intn(8)
				element := reflect.New(field.Type().Elem()).Elem()
				f := element.Type().Field(accessRand).Name
				fillAccess(t, accessFieldNames[accessRand], reflect.ValueOf(&f))
				slice.Index(j).Set(element)
			}
			field.Set(slice)
		}

		overrides := map[string]fuzz.OverrideFiller{
			".MaxTimeTravelHours":           funcF,
			".DefaultPartitionExpirationMs": funcMs,
			".DefaultTableExpirationMs":     funcMs,
			".Access":                       funcAccess,
		}
		filler := fuzz.NewRandomFiller(&fuzz.FillerConfig{Stream: stream, FieldOverrides: overrides})

		// krm1
		k1 := &krm.BigQueryDatasetSpec{}

		// fill
		filler.Fill(t, k1)

		// ToProto
		ctx := &direct.MapContext{}
		p := BigQueryDatasetSpec_ToProto(ctx, k1)
		if ctx.Err() != nil {
			t.Fatalf("error converting KRM to proto: %v \n KRM: %s", ctx.Err(), prettyPrint(t, k1))
		}
		t.Logf("proto= %+v", p)
		k2 := BigQueryDatasetSpec_FromProto(ctx, p)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to KRM: %v", ctx.Err())
		}
		// Using cmpopts.IgnoreFields to ignore specific fields in ProjectRef
		opts := cmp.Options{
			// project ref
			cmpopts.IgnoreFields(refs.ProjectRef{}, "External", "Name", "Namespace"),

			// other ref
			cmpopts.IgnoreFields(kmsv1beta1.KMSKeyRef_OneOf{}, "KMSCryptoKeyRef", "External", "Name", "Namespace"),
			cmpopts.IgnoreFields(kmsv1beta1.KMSKeyRef_OneOf{}, "AutoKeyRef", "External", "Name", "Namespace"),

			// unroundtrippable fields (for now)
			cmpopts.IgnoreFields(krm.BigQueryDatasetSpec{}, "ResourceID"),
			cmpopts.IgnoreFields(krm.BigQueryDatasetSpec{}, "ProjectRef"),
			cmpopts.IgnoreFields(krm.BigQueryDatasetSpec{}, "IsCaseInsensitive"),
		}
		// compare
		if diff := cmp.Diff(k1, k2, opts...); diff != "" {
			t.Logf("k1 = %+v", k1)
			t.Logf("k2 = %+v", k2)
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

func fillAccess(t *testing.T, fieldName string, field reflect.Value) {
	switch field.Kind() {
	case reflect.Ptr:
		if field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))
			fillAccess(t, fieldName, field.Elem())
		}
	case reflect.String:
		field.SetString(RandomString(10))

	case reflect.Struct:
		for i := 0; i < field.NumField(); i++ {
			structFieldName := field.Type().Field(i).Name
			nestedStructFieldname := fieldName + "." + structFieldName
			fillAccess(t, nestedStructFieldname, field.Field(i))
		}
	default:
		t.Fatalf("Unhandled access field kind: %v", field.Kind())
	}
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func RandomString(length int) string {
	return StringWithCharset(length, charset)
}
