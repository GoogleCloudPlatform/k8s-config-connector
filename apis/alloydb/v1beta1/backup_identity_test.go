// Copyright 2025 Google LLC
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

package v1beta1

import (
	"maps"
	"testing"

	util "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/identity"
)

func TestBackupParse(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		parsedMap map[string]string
		hasError  bool
	}{
		{
			name:      "Normal parse",
			input:     "projects/myProject/locations/mylocation/backups/mybackup",
			parsedMap: map[string]string{"projects": "myProject", "locations": "mylocation", "backups": "mybackup"},
			hasError:  false,
		},
		{
			name:      "Normal parse with leading slash",
			input:     "/projects/p1/locations/l1/backups/b1",
			parsedMap: map[string]string{"projects": "p1", "locations": "l1", "backups": "b1"},
			hasError:  false,
		},
		{
			name:      "Normal parse with domain",
			input:     "alloydb.googleapis.com/projects/first/locations/second/backups/third",
			parsedMap: map[string]string{"projects": "first", "locations": "second", "backups": "third"},
			hasError:  false,
		},
		{
			name:      "Normal parse with slashed domain",
			input:     "//alloydb.googleapis.com/projects/athos/locations/porthos/backups/aramis",
			parsedMap: map[string]string{"projects": "athos", "locations": "porthos", "backups": "aramis"},
			hasError:  false,
		},
		{
			name:      "Normal parse with wrong domain",
			input:     "//anthos.googleapis.com/projects/myProject/locations/mylocation/backups/mybackup",
			parsedMap: nil,
			hasError:  true,
		},
		{
			name:      "Normal parse with wrong project key",
			input:     "org/myProject/locations/mylocation/backups/mybackup",
			parsedMap: nil,
			hasError:  true,
		},
		{
			name:      "Normal parse with wrong location key",
			input:     "projects/myProject/regions/mylocation/backups/mybackup",
			parsedMap: nil,
			hasError:  true,
		},
		{
			name:      "Normal parse with wrong backup key",
			input:     "projects/myProject/locations/mylocation/cluster/mybackup",
			parsedMap: nil,
			hasError:  true,
		},
	}

	for _, tc := range tests {
		err, result := util.ParseIdentityMap(tc.input, parser, 2)
		if tc.hasError {
			if err == nil {
				t.Fatalf("Test %s expected error but did not get one", tc.name)
			}
			continue
		}
		// Error no expected at this point
		if err != nil {
			t.Fatalf("Test %s did not expect error but got %v", tc.name, err)
		}
		if result == nil {
			t.Fatalf("Test %s expected a result but did not get one", tc.name)
		}
		if !maps.Equal(result, tc.parsedMap) {
			t.Fatalf("Test %s bad result %v != %v", tc.name, result, tc.parsedMap)
		}
	}
}
