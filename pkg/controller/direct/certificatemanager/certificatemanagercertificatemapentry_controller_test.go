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

package certificatemanager

import (
	"testing"
)

func TestEqualCertificateRef(t *testing.T) {
	testCases := []struct {
		name string
		a    string
		b    string
		want bool
	}{
		{
			name: "identical project ID",
			a:    "projects/my-proj/locations/global/certificates/my-cert",
			b:    "projects/my-proj/locations/global/certificates/my-cert",
			want: true,
		},
		{
			name: "project number vs project ID with matching cert",
			a:    "projects/123456789/locations/global/certificates/my-cert",
			b:    "projects/my-proj/locations/global/certificates/my-cert",
			want: true,
		},
		{
			name: "project ID vs project number with matching cert",
			a:    "projects/my-proj/locations/global/certificates/my-cert",
			b:    "projects/123456789/locations/global/certificates/my-cert",
			want: true,
		},
		{
			name: "different certificate names",
			a:    "projects/123456789/locations/global/certificates/cert-a",
			b:    "projects/my-proj/locations/global/certificates/cert-b",
			want: false,
		},
		{
			name: "different locations",
			a:    "projects/123456789/locations/us-central1/certificates/my-cert",
			b:    "projects/my-proj/locations/global/certificates/my-cert",
			want: false,
		},
		{
			name: "different project IDs",
			a:    "projects/proj-a/locations/global/certificates/my-cert",
			b:    "projects/proj-b/locations/global/certificates/my-cert",
			want: false,
		},
		{
			name: "different project numbers",
			a:    "projects/11111/locations/global/certificates/my-cert",
			b:    "projects/22222/locations/global/certificates/my-cert",
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := equalCertificateRef(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("equalCertificateRef(%q, %q) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestEqualCertificateSlices(t *testing.T) {
	testCases := []struct {
		name    string
		desired []string
		actual  []string
		want    bool
	}{
		{
			name:    "empty slices",
			desired: []string{},
			actual:  []string{},
			want:    true,
		},
		{
			name:    "different lengths",
			desired: []string{"projects/my-proj/locations/global/certificates/cert-1"},
			actual:  []string{},
			want:    false,
		},
		{
			name:    "single cert with project number suppression",
			desired: []string{"projects/my-proj/locations/global/certificates/cert-1"},
			actual:  []string{"projects/123456789/locations/global/certificates/cert-1"},
			want:    true,
		},
		{
			name: "multiple certs in order",
			desired: []string{
				"projects/my-proj/locations/global/certificates/cert-1",
				"projects/my-proj/locations/global/certificates/cert-2",
			},
			actual: []string{
				"projects/123456789/locations/global/certificates/cert-1",
				"projects/123456789/locations/global/certificates/cert-2",
			},
			want: true,
		},
		{
			name: "multiple certs out of order",
			desired: []string{
				"projects/my-proj/locations/global/certificates/cert-2",
				"projects/my-proj/locations/global/certificates/cert-1",
			},
			actual: []string{
				"projects/123456789/locations/global/certificates/cert-1",
				"projects/123456789/locations/global/certificates/cert-2",
			},
			want: true,
		},
		{
			name: "multiple certs mismatch",
			desired: []string{
				"projects/my-proj/locations/global/certificates/cert-1",
				"projects/my-proj/locations/global/certificates/cert-3",
			},
			actual: []string{
				"projects/123456789/locations/global/certificates/cert-1",
				"projects/123456789/locations/global/certificates/cert-2",
			},
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := equalCertificateSlices(tc.desired, tc.actual)
			if got != tc.want {
				t.Errorf("equalCertificateSlices(%v, %v) = %v, want %v", tc.desired, tc.actual, got, tc.want)
			}
		})
	}
}
