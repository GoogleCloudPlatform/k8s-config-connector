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

package fuzztesting

import (
	"math/rand"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"

	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/register"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func FuzzAllMappers(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		randStream := rand.New(rand.NewSource(seed))

		fuzzer := fuzztesting.ChooseFuzzer(randStream.Int63())
		nextSeed := randStream.Int63()
		fuzzer(t, nextSeed)
	})
}

func TestSomeMappers(t *testing.T) {
	t.Parallel()
	seed := time.Now().UnixNano()
	randStream := rand.New(rand.NewSource(seed))

	for i := 0; i < 100000; i++ {
		fuzzer := fuzztesting.ChooseFuzzer(randStream.Int63())
		nextSeed := randStream.Int63()
		fuzzer(t, nextSeed)
	}
}

func TestListTypesWithoutFuzzers(t *testing.T) {
	crds, err := crdloader.LoadAllCRDs()
	if err != nil {
		t.Fatalf("failed to load all CRDs: %v", err)
	}

	// GVKs that have registered fuzzers
	fuzzedGroupKinds := make(map[schema.GroupKind]bool)
	fuzzedKinds := make(map[string]bool)

	fuzzers := fuzztesting.GetRegisteredFuzzers()
	for _, f := range fuzzers {
		specType := getFuzzerSpecType(f)
		if specType == nil {
			continue
		}
		kind := strings.TrimSuffix(specType.Name(), "Spec")
		fuzzedKinds[kind] = true

		// Try to parse the service/group name from package path
		pkgPath := specType.PkgPath()
		parts := strings.Split(pkgPath, "/")
		var service string
		// Search for "apis" segment
		for i, part := range parts {
			if part == "apis" && i+1 < len(parts) {
				service = parts[i+1]
				break
			}
		}
		if service == "" && len(parts) >= 2 {
			service = parts[len(parts)-2]
		}
		if service != "" {
			group := service + ".cnrm.cloud.google.com"
			fuzzedGroupKinds[schema.GroupKind{Group: group, Kind: kind}] = true
		}
	}

	// Map CRDs to unique GroupKinds
	var missing []schema.GroupKind
	var allGroupKinds []schema.GroupKind
	seenGK := make(map[schema.GroupKind]bool)

	for _, crd := range crds {
		gk := schema.GroupKind{
			Group: crd.Spec.Group,
			Kind:  crd.Spec.Names.Kind,
		}
		if seenGK[gk] {
			continue
		}
		seenGK[gk] = true
		allGroupKinds = append(allGroupKinds, gk)

		// Check both fuzzedGroupKinds (full match) and fuzzedKinds (kind-only match as fallback)
		if !fuzzedGroupKinds[gk] && !fuzzedKinds[gk.Kind] {
			missing = append(missing, gk)
		}
	}

	// Sort missing for nice stable output
	sort.Slice(missing, func(i, j int) bool {
		if missing[i].Group != missing[j].Group {
			return missing[i].Group < missing[j].Group
		}
		return missing[i].Kind < missing[j].Kind
	})

	t.Logf("=== Informational: KCC Kinds Without Fuzzers ===")
	t.Logf("Total Kinds in CRDs: %d", len(allGroupKinds))
	t.Logf("Kinds with Fuzzers: %d", len(allGroupKinds)-len(missing))
	t.Logf("Kinds without Fuzzers: %d", len(missing))
	t.Logf("------------------------------------------------")
	for _, gk := range missing {
		t.Logf("  - %s (%s)", gk.Kind, gk.Group)
	}
	t.Logf("================================================")
}

func getFuzzerSpecType(fuzzer any) reflect.Type {
	val := reflect.ValueOf(fuzzer)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil
	}
	field := val.FieldByName("SpecFromProto")
	if !field.IsValid() {
		field = val.FieldByName("SpecFromAPI")
	}
	if !field.IsValid() {
		return nil
	}
	ft := field.Type()
	if ft.Kind() != reflect.Func || ft.NumOut() < 1 {
		return nil
	}
	outType := ft.Out(0)
	if outType.Kind() == reflect.Ptr {
		return outType.Elem()
	}
	return outType
}
