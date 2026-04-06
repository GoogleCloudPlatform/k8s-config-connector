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

package bigquery

import (
	"reflect"
	"testing"

	bigquery "google.golang.org/api/bigquery/v2"
)

func TestPolicyTagsEqual(t *testing.T) {
	tests := []struct {
		name      string
		a         *bigquery.TableFieldSchemaPolicyTags
		b         *bigquery.TableFieldSchemaPolicyTags
		wantEqual bool
	}{
		{
			name:      "both nil",
			a:         nil,
			b:         nil,
			wantEqual: true,
		},
		{
			name:      "a nil, b empty",
			a:         nil,
			b:         &bigquery.TableFieldSchemaPolicyTags{Names: []string{}},
			wantEqual: true,
		},
		{
			name:      "a empty, b nil",
			a:         &bigquery.TableFieldSchemaPolicyTags{Names: []string{}},
			b:         nil,
			wantEqual: true,
		},
		{
			name:      "both empty",
			a:         &bigquery.TableFieldSchemaPolicyTags{Names: []string{}},
			b:         &bigquery.TableFieldSchemaPolicyTags{Names: []string{}},
			wantEqual: true,
		},
		{
			name:      "a nil, b not empty",
			a:         nil,
			b:         &bigquery.TableFieldSchemaPolicyTags{Names: []string{"tag1"}},
			wantEqual: false,
		},
		{
			name:      "a not empty, b nil",
			a:         &bigquery.TableFieldSchemaPolicyTags{Names: []string{"tag1"}},
			b:         nil,
			wantEqual: false,
		},
		{
			name:      "equal tags",
			a:         &bigquery.TableFieldSchemaPolicyTags{Names: []string{"tag1", "tag2"}},
			b:         &bigquery.TableFieldSchemaPolicyTags{Names: []string{"tag2", "tag1"}},
			wantEqual: true,
		},
		{
			name:      "different tags",
			a:         &bigquery.TableFieldSchemaPolicyTags{Names: []string{"tag1"}},
			b:         &bigquery.TableFieldSchemaPolicyTags{Names: []string{"tag2"}},
			wantEqual: false,
		},
		{
			name:      "non-nil objects, a.Names nil, b.Names empty",
			a:         &bigquery.TableFieldSchemaPolicyTags{Names: nil},
			b:         &bigquery.TableFieldSchemaPolicyTags{Names: []string{}},
			wantEqual: true,
		},
		{
			name:      "non-nil objects, a.Names empty, b.Names nil",
			a:         &bigquery.TableFieldSchemaPolicyTags{Names: []string{}},
			b:         &bigquery.TableFieldSchemaPolicyTags{Names: nil},
			wantEqual: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := policyTagsEqual(tt.a, tt.b); got != tt.wantEqual {
				t.Errorf("policyTagsEqual() = %v, want %v", got, tt.wantEqual)
			}
		})
	}
}

func TestPolicyTagsEqualNoInPlaceSort(t *testing.T) {
	a := &bigquery.TableFieldSchemaPolicyTags{Names: []string{"b", "a"}}
	b := &bigquery.TableFieldSchemaPolicyTags{Names: []string{"a", "b"}}

	aOrig := []string{"b", "a"}
	bOrig := []string{"a", "b"}

	if !policyTagsEqual(a, b) {
		t.Errorf("policyTagsEqual() = false, want true")
	}

	if !reflect.DeepEqual(a.Names, aOrig) {
		t.Errorf("policyTagsEqual modified input a: got %v, want %v", a.Names, aOrig)
	}
	if !reflect.DeepEqual(b.Names, bOrig) {
		t.Errorf("policyTagsEqual modified input b: got %v, want %v", b.Names, bOrig)
	}
}
