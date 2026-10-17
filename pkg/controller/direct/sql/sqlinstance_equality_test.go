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

package sql

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	api "google.golang.org/api/sqladmin/v1beta4"
)

func TestDiffInstances_StrictPointersMatch(t *testing.T) {
	t.Parallel()
	// desired has all optional struct pointers as nil (typical for a minimal KRM spec)
	desired := &api.DatabaseInstance{
		Settings: &api.Settings{},
	}

	// actual has these fields populated with "empty" defaults by the GCP API
	actual := &api.DatabaseInstance{
		DiskEncryptionConfiguration: &api.DiskEncryptionConfiguration{
			Kind: "sql#diskEncryptionConfiguration",
		},
		ReplicaConfiguration: &api.ReplicaConfiguration{
			Kind: "sql#replicaConfiguration",
		},
		ReplicationCluster: &api.ReplicationCluster{},
		Settings: &api.Settings{
			BackupConfiguration: &api.BackupConfiguration{
				Kind: "sql#backupConfiguration",
			},
			DataCacheConfig: &api.DataCacheConfig{},
			IpConfiguration: &api.IpConfiguration{
				Ipv4Enabled: true,
				SslMode:     "ALLOW_UNENCRYPTED_AND_ENCRYPTED",
			},
			LocationPreference: &api.LocationPreference{
				Kind: "sql#locationPreference",
			},
		},
	}

	diff := DiffInstances(desired, actual)

	// We want HasDiff() to be false because these are semantically equivalent.
	// Nil values in desired (KRM) should match default/empty values in actual (GCP).
	if diff.HasDiff() {
		t.Errorf("DiffInstances() identified unexpected diffs: %v", diff.Fields)
	}
}

func TestDiffInstances_NilSettings(t *testing.T) {
	t.Parallel()
	// desired has Settings as nil
	desired := &api.DatabaseInstance{
		Settings: nil,
	}

	// actual has Settings as an empty object
	actual := &api.DatabaseInstance{
		Settings: &api.Settings{},
	}

	diff := DiffInstances(desired, actual)

	if diff.HasDiff() {
		t.Errorf("DiffInstances() identified unexpected diffs when Settings is nil in desired: %v", diff.Fields)
	}
}

func TestDiffInstances_NilUserLabels(t *testing.T) {
	t.Parallel()
	// desired has nil UserLabels
	desired := &api.DatabaseInstance{
		Settings: &api.Settings{
			UserLabels: nil,
		},
	}

	// actual has an empty UserLabels map
	actual := &api.DatabaseInstance{
		Settings: &api.Settings{
			UserLabels: map[string]string{},
		},
	}

	diff := DiffInstances(desired, actual)

	if diff.HasDiff() {
		t.Errorf("DiffInstances() identified unexpected diffs when UserLabels is nil in desired: %v", diff.Fields)
	}
}

func TestDiffInstances_AuthorizedNetworksSorting(t *testing.T) {
	t.Parallel()
	desired := &api.DatabaseInstance{
		Settings: &api.Settings{
			IpConfiguration: &api.IpConfiguration{
				AuthorizedNetworks: []*api.AclEntry{
					{Name: "b", Value: "2.2.2.2"},
					{Name: "a", Value: "1.1.1.1"},
					{Name: "", Value: "4.4.4.4"},
					{Name: "", Value: "3.3.3.3"},
				},
			},
		},
	}

	actual := &api.DatabaseInstance{
		Settings: &api.Settings{
			IpConfiguration: &api.IpConfiguration{
				AuthorizedNetworks: []*api.AclEntry{
					{Name: "", Value: "3.3.3.3"},
					{Name: "a", Value: "1.1.1.1"},
					{Name: "b", Value: "2.2.2.2"},
					{Name: "", Value: "4.4.4.4"},
				},
			},
		},
	}

	diff := DiffInstances(desired, actual)

	if diff.HasDiff() {
		t.Errorf("DiffInstances() identified unexpected diffs due to AuthorizedNetworks sorting: %v", diff.Fields)
	}
}

func TestDiffInstances_DatabaseFlagsSorting(t *testing.T) {
	t.Parallel()
	desired := &api.DatabaseInstance{
		Settings: &api.Settings{
			DatabaseFlags: []*api.DatabaseFlags{
				{Name: "b", Value: "v2"},
				{Name: "a", Value: "v1"},
			},
		},
	}

	actual := &api.DatabaseInstance{
		Settings: &api.Settings{
			DatabaseFlags: []*api.DatabaseFlags{
				{Name: "a", Value: "v1"},
				{Name: "b", Value: "v2"},
			},
		},
	}

	diff := DiffInstances(desired, actual)

	if diff.HasDiff() {
		t.Errorf("DiffInstances() identified unexpected diffs due to DatabaseFlags sorting: %v", diff.Fields)
	}
}

func TestDiffInstances_UserLabelsDiff(t *testing.T) {
	t.Parallel()
	desired := &api.DatabaseInstance{
		Settings: &api.Settings{
			UserLabels: map[string]string{
				"key1": "val1",
				"key2": "val2-changed",
				"key3": "val3-new",
			},
		},
	}

	actual := &api.DatabaseInstance{
		Settings: &api.Settings{
			UserLabels: map[string]string{
				"key1": "val1",
				"key2": "val2",
				"key4": "val4-removed",
			},
		},
	}

	diff := DiffInstances(desired, actual)

	if !diff.HasDiff() {
		t.Errorf("DiffInstances() expected diffs, but got none")
	}

	expectedDiffs := map[string]struct {
		Old any
		New any
	}{
		".settings.userLabels[*\"key2\"]": {Old: "val2", New: "val2-changed"},
		".settings.userLabels[+\"key3\"]": {Old: nil, New: "val3-new"},
		".settings.userLabels[-\"key4\"]": {Old: "val4-removed", New: nil},
	}

	if len(diff.Fields) != len(expectedDiffs) {
		t.Fatalf("DiffInstances() expected %d diffs, got %d. Fields: %v", len(expectedDiffs), len(diff.Fields), diff.Fields)
	}

	for _, field := range diff.Fields {
		expected, ok := expectedDiffs[field.ID]
		if !ok {
			t.Errorf("Unexpected diff field ID %q", field.ID)
			continue
		}
		if field.Old != expected.Old {
			t.Errorf("For field %q, expected Old=%v, got %v", field.ID, expected.Old, field.Old)
		}
		if field.New != expected.New {
			t.Errorf("For field %q, expected New=%v, got %v", field.ID, expected.New, field.New)
		}
	}
}

func TestDiffInstances_AvailabilityTypeCasing(t *testing.T) {
	desired := &api.DatabaseInstance{
		Settings: &api.Settings{
			AvailabilityType: "Zonal",
		},
	}

	actual := &api.DatabaseInstance{
		Settings: &api.Settings{
			AvailabilityType: "ZONAL",
		},
	}

	diff := DiffInstances(desired, actual)

	if diff.HasDiff() {
		t.Errorf("DiffInstances() expected no diffs, but got: %v", diff.Fields)
	}
}

func TestSQLInstance_ExhaustiveSpecCoverage(t *testing.T) {
	baseline := &krm.SQLInstance{}
	populateObject(reflect.ValueOf(&baseline.Spec))

	// Get baseline GCP
	gcpBaseline, err := SQLInstanceKRMToGCP(baseline, nil, nil)
	if err != nil {
		t.Fatalf("Failed to map baseline KRM to GCP DatabaseInstance: %v", err)
	}

	// Walk KRM SQLInstanceSpec recursively to discover all fields
	var fields []mutableField
	findMutableFields(reflect.ValueOf(&baseline.Spec), ".spec", nil, &fields)

	if len(fields) == 0 {
		t.Fatalf("No mutable fields discovered in krm.SQLInstanceSpec!")
	}

	// Explicitly ignored paths (with documented rationale)
	ignoredSpecFields := map[string]string{
		".spec.rootPassword":                  "Unreadable/write-only field and is handled separately",
		".spec.replicaConfiguration.password": "Unreadable/write-only field and is handled separately",
		".spec.settings.crashSafeReplication": "Deprecated/ignored field; only applicable to first-generation instances",
	}

	for _, f := range fields {
		// Skip cloneSource fields (immutable cloning-only fields, not part of regular state updates or diffs)
		if strings.HasPrefix(f.path, ".spec.cloneSource") {
			continue
		}
		ignored := false
		for prefix, reason := range ignoredSpecFields {
			if strings.HasPrefix(f.path, prefix) {
				t.Logf("Skipping ignored field %s: %s", f.path, reason)
				ignored = true
				break
			}
		}
		if ignored {
			continue
		}

		// Mutate a clone of the baseline
		krmMutated := baseline.DeepCopy()
		mutateFieldBySteps(reflect.ValueOf(&krmMutated.Spec), f.steps)

		// Stage 1 (Mapper Validation)
		gcpMutated, err := SQLInstanceKRMToGCP(krmMutated, nil, nil)
		if err != nil {
			t.Errorf("Path %s: SQLInstanceKRMToGCP returned error: %v", f.path, err)
			continue
		}

		if reflect.DeepEqual(gcpMutated, gcpBaseline) {
			t.Errorf("Top-level mapper SQLInstanceKRMToGCP failed to map field %s!", f.path)
			continue
		}

		// Stage 2 (Diff Validation)
		diff := DiffInstances(gcpMutated, gcpBaseline)
		if !diff.HasDiff() {
			t.Errorf("DiffInstances failed to detect diff for field %s!", f.path)
		}
	}
}

type PathStep struct {
	Kind      reflect.Kind
	FieldName string
	SliceIdx  int
}

type mutableField struct {
	path  string
	steps []PathStep
}

func deref(val reflect.Value) reflect.Value {
	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			val.Set(reflect.New(val.Type().Elem()))
		}
		val = val.Elem()
	}
	return val
}

func isReferenceType(val reflect.Value) bool {
	if val.Kind() != reflect.Struct {
		return false
	}
	_, found := val.Type().FieldByName("External")
	return found
}

func getJSONFieldName(f reflect.StructField) string {
	tag := f.Tag.Get("json")
	if tag == "" || tag == "-" {
		if len(f.Name) == 0 {
			return ""
		}
		return strings.ToLower(f.Name[:1]) + f.Name[1:]
	}
	parts := strings.Split(tag, ",")
	return parts[0]
}

func populatePrimitive(v reflect.Value, strVal string) {
	switch v.Kind() {
	case reflect.String:
		v.SetString(strVal)
	case reflect.Bool:
		v.SetBool(true)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(42)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.SetUint(42)
	case reflect.Float32, reflect.Float64:
		v.SetFloat(42.42)
	}
}

func populateObject(v reflect.Value) {
	v = deref(v)

	if v.Kind() == reflect.Struct && isReferenceType(v) {
		extField := v.FieldByName("External")
		if extField.IsValid() && extField.CanSet() {
			extField.SetString("test-external-ref")
		}
		return
	}

	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)
			if !f.CanSet() {
				continue
			}
			populateObject(f)
		}
	case reflect.Slice:
		if v.IsNil() || v.Len() == 0 {
			v.Set(reflect.MakeSlice(v.Type(), 1, 1))
		}
		for i := 0; i < v.Len(); i++ {
			populateObject(v.Index(i))
		}
	case reflect.Map:
		if v.IsNil() {
			v.Set(reflect.MakeMap(v.Type()))
		}
		keyType := v.Type().Key()
		valType := v.Type().Elem()
		keyVal := reflect.New(keyType).Elem()
		populatePrimitive(keyVal, "testkey")
		valVal := reflect.New(valType).Elem()
		populateObject(valVal)
		v.SetMapIndex(keyVal, valVal)
	default:
		populatePrimitive(v, "testval")
	}
}

func findMutableFields(val reflect.Value, path string, currentSteps []PathStep, results *[]mutableField) {
	val = deref(val)

	if val.Kind() == reflect.Struct && isReferenceType(val) {
		steps := append([]PathStep(nil), currentSteps...)
		steps = append(steps, PathStep{Kind: reflect.Struct, FieldName: "External"})
		*results = append(*results, mutableField{
			path:  path + ".external",
			steps: steps,
		})
		return
	}

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			f := val.Type().Field(i)
			if !f.IsExported() {
				continue
			}
			steps := append([]PathStep(nil), currentSteps...)
			steps = append(steps, PathStep{Kind: reflect.Struct, FieldName: f.Name})
			jsonName := getJSONFieldName(f)
			findMutableFields(val.Field(i), path+"."+jsonName, steps, results)
		}
	case reflect.Slice:
		if val.Type().Elem().Kind() == reflect.Struct || (val.Type().Elem().Kind() == reflect.Ptr && val.Type().Elem().Elem().Kind() == reflect.Struct) {
			for i := 0; i < val.Len(); i++ {
				steps := append([]PathStep(nil), currentSteps...)
				steps = append(steps, PathStep{Kind: reflect.Slice, SliceIdx: i})
				findMutableFields(val.Index(i), fmt.Sprintf("%s[%d]", path, i), steps, results)
			}
		} else {
			*results = append(*results, mutableField{
				path:  path,
				steps: currentSteps,
			})
		}
	case reflect.Map:
		*results = append(*results, mutableField{
			path:  path,
			steps: currentSteps,
		})
	default:
		*results = append(*results, mutableField{
			path:  path,
			steps: currentSteps,
		})
	}
}

func applyMutation(val reflect.Value) {
	val = deref(val)
	switch val.Kind() {
	case reflect.String:
		val.SetString(val.String() + "-mutated")
	case reflect.Bool:
		val.SetBool(!val.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val.SetInt(val.Int() + 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val.SetUint(val.Uint() + 10)
	case reflect.Float32, reflect.Float64:
		val.SetFloat(val.Float() + 10.0)
	case reflect.Slice:
		if val.Len() > 0 {
			elem := val.Index(0)
			applyMutation(elem)
		} else {
			newElem := reflect.New(val.Type().Elem()).Elem()
			applyMutation(newElem)
			val.Set(reflect.Append(val, newElem))
		}
	case reflect.Map:
		if val.IsNil() {
			val.Set(reflect.MakeMap(val.Type()))
		}
		k := reflect.New(val.Type().Key()).Elem()
		populatePrimitive(k, "mutated-key")
		v := reflect.New(val.Type().Elem()).Elem()
		populatePrimitive(v, "mutated-val")
		val.SetMapIndex(k, v)
	}
}

func mutateFieldBySteps(val reflect.Value, steps []PathStep) {
	val = deref(val)

	if len(steps) == 0 {
		applyMutation(val)
		return
	}

	step := steps[0]
	switch step.Kind {
	case reflect.Struct:
		f := val.FieldByName(step.FieldName)
		mutateFieldBySteps(f, steps[1:])
	case reflect.Slice:
		mutateFieldBySteps(val.Index(step.SliceIdx), steps[1:])
	}
}
