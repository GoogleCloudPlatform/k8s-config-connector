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

package lint

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"

	"k8s.io/apimachinery/pkg/util/sets"
)

// identityCollectionRe extracts the collection segment from an Identity's String() method:
//
//	return i.parent.String() + "/managedFolders/" + i.id
//
// Anchoring on parent.String() ensures only the resource's direct collection segment is matched.
var identityCollectionRe = regexp.MustCompile(`parent\.String\(\)\s*\+\s*"/([a-zA-Z0-9]+)/"`)

// pathSegmentRe finds path segments in recorded traffic to extract observed GCP casings.
var pathSegmentRe = regexp.MustCompile(`/([a-zA-Z0-9]+)/`)

// TestIdentityCollectionCasing validates that the collection segment emitted by an Identity
// uses the canonical casing observed in GCP HTTP traffic logs.
//
// Unlike KRM field names (validated by TestCRDsAcronyms), external reference collection
// segments form part of the GCP resource identifier (status.externalRef) and must match
// the GCP API path casing.
func TestIdentityCollectionCasing(t *testing.T) {
	t.Parallel()

	identityCollections, err := loadIdentityCollections("../../apis")
	if err != nil {
		t.Fatalf("scanning identity files: %v", err)
	}
	if len(identityCollections) == 0 {
		t.Fatal("found no identity collections; the regex or the layout changed")
	}

	interest := sets.NewString()
	for c := range identityCollections {
		interest.Insert(strings.ToLower(c))
	}

	observed, err := observedCollectionCasings("../../pkg/test/resourcefixture/testdata", interest)
	if err != nil {
		t.Fatalf("scanning recorded traffic: %v", err)
	}

	var errs []string
	for collection, files := range identityCollections {
		seen, ok := observed[strings.ToLower(collection)]
		if !ok || seen.Len() == 0 {
			// No traffic mentions this collection at all: nothing to compare.
			continue
		}
		if seen.Has(collection) {
			continue
		}
		sort.Strings(files)
		errs = append(errs, fmt.Sprintf(
			"[identity-casing] collection %q is never used in recorded traffic (traffic uses %s): %s",
			collection, strings.Join(seen.List(), ", "), strings.Join(files, " ")))
	}
	sort.Strings(errs)

	// Ratchet: these are shipped resources, so fixing one changes a published
	// status.externalRef. The list may shrink, never grow.
	test.CompareRatchetFile(t, "testdata/exceptions/identity_collection_casing.txt", strings.Join(errs, "\n"))
}

// TestIdentityCollectionRegex verifies that identityCollectionRe matches the
// resource collection segment in Identity.String() rather than earlier segments
// (such as "/locations/") in Parent.String().
func TestIdentityCollectionRegex(t *testing.T) {
	// Shape of a real identity file: the Parent's String() comes first in the
	// file and mentions /locations/, then the Identity's String().
	const body = `
func (p *FooParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

func (i *FooIdentity) String() string {
	return i.parent.String() + "/fooBars/" + i.id
}
`
	m := identityCollectionRe.FindStringSubmatch(body)
	if m == nil {
		t.Fatal("regex matched nothing; identity file layout may have changed")
	}
	if got := m[1]; got != "fooBars" {
		t.Errorf("extracted %q, want %q; the regex is matching the parent's segment", got, "fooBars")
	}
}

// loadIdentityCollections maps each collection segment to the identity files
// that build it.
func loadIdentityCollections(apisDir string) (map[string][]string, error) {
	out := map[string][]string{}
	err := filepath.Walk(apisDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(path, "_identity.go") {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		m := identityCollectionRe.FindSubmatch(data)
		if m == nil {
			return nil
		}
		collection := string(m[1])
		rel := strings.TrimPrefix(path, apisDir+"/")
		out[collection] = append(out[collection], rel)
		return nil
	})
	return out, err
}

// observedCollectionCasings records, for each collection we care about, every
// exact casing seen in recorded traffic. Keyed by the lowercased name so
// differing casings collapse into one entry.
func observedCollectionCasings(testdataDir string, interest sets.String) (map[string]sets.String, error) {
	out := map[string]sets.String{}
	err := filepath.Walk(testdataDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || filepath.Base(path) != "_http.log" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		for _, m := range pathSegmentRe.FindAllSubmatch(data, -1) {
			seg := string(m[1])
			key := strings.ToLower(seg)
			if !interest.Has(key) {
				continue
			}
			if out[key] == nil {
				out[key] = sets.NewString()
			}
			out[key].Insert(seg)
		}
		return nil
	})
	return out, err
}
