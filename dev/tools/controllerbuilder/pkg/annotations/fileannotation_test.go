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

package annotations

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSimpleRoundTrip(t *testing.T) {
	in := `
// This is a file

// +tool:mockgcp-service
// http.host: gkehub.googleapis.com
// proto.service: google.cloud.gkehub.v1beta.GkeHub
// proto.service: google.cloud.gkehub.v1beta1.GkeHubMembershipService

// After the new line we can continue with our logic
func main() {
}`

	want := `// +tool:mockgcp-service
// http.host: gkehub.googleapis.com
// proto.service: google.cloud.gkehub.v1beta.GkeHub
// proto.service: google.cloud.gkehub.v1beta1.GkeHubMembershipService
`

	annotations, err := FindFileAnnotations([]byte(in), []string{"+tool:"})
	if err != nil {
		t.Fatalf("unexpected error from FindFileAnnotations: %v", err)
	}

	if len(annotations) != 1 {
		t.Fatalf("expected %d annotations, got %d", 1, len(annotations))
	}

	annotation := annotations[0]

	got := annotation.FormatGo()

	if got != want {
		diff := cmp.Diff(got, want)
		t.Errorf("unexpected output; got %q, want %q\ndiff=%v", got, want, diff)
	}
}
