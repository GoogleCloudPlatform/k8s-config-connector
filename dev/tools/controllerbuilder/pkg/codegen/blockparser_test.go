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

package codegen

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestParseGeneratedGoBlocks(t *testing.T) {
	src := []byte(`// Copyright 2026 Google LLC

package v1alpha1

import (
	"fmt"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/foo/v1alpha1"
)

// +kcc:proto=google.cloud.foo.v1.Alpha
type Alpha struct {
	Name *string
}

/* unreachable type Beta
// +kcc:proto=google.cloud.foo.v1.Beta
type Beta struct {
}
*/

// +kcc:observedstate:proto=google.cloud.foo.v1.Alpha
type AlphaObservedState struct {
	State *string
}

func Alpha_FromProto(in *string) *krm.Alpha {
	_ = fmt.Sprintf
	return nil
}

/*
func Beta_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Beta) *string {
	return nil
}
*/
`)

	blocks, imports, err := parseGeneratedGoBlocks(src, "v1alpha1", false)
	if err != nil {
		t.Fatalf("parseGeneratedGoBlocks failed: %v", err)
	}

	wantImports := map[string]string{
		"fmt": "",
		"github.com/GoogleCloudPlatform/k8s-config-connector/apis/foo/v1alpha1": "krm",
	}

	if diff := cmp.Diff(wantImports, imports); diff != "" {
		t.Errorf("imports mismatch (-want +got):\n%s", diff)
	}

	wantBlocks := []generatedBlock{
		{
			name:     "Alpha",
			kind:     "type",
			protoFQN: "google.cloud.foo.v1.Alpha",
		},
		{
			name:        "Beta",
			kind:        "type",
			isCommented: true,
			protoFQN:    "google.cloud.foo.v1.Beta",
		},
		{
			name:            "AlphaObservedState",
			kind:            "type",
			protoFQN:        "google.cloud.foo.v1.Alpha",
			isObservedState: true,
		},
		{
			name:        "Alpha_FromProto",
			kind:        "func",
			krmTypeName: "Alpha",
		},
		{
			name:        "Beta_v1beta1_ToProto",
			kind:        "func",
			isCommented: true,
			krmTypeName: "Beta",
		},
	}

	ignoreFields := cmpopts.IgnoreFields(generatedBlock{}, "rawText", "offset")
	if diff := cmp.Diff(wantBlocks, blocks, cmp.AllowUnexported(generatedBlock{}), ignoreFields); diff != "" {
		t.Errorf("blocks mismatch (-want +got):\n%s", diff)
	}
}
