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

package v1alpha1

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestWasmPluginIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *WasmPluginIdentity
		wantErr bool
	}{
		{
			name: "valid regional wasm plugin reference",
			ref:  "projects/my-project/locations/us-central1/wasmPlugins/my-wasmplugin",
			want: &WasmPluginIdentity{
				parent: &parent.ProjectAndLocationParent{
					ProjectID: "my-project",
					Location:  "us-central1",
				},
				id: "my-wasmplugin",
			},
			wantErr: false,
		},
		{
			name: "valid global wasm plugin reference",
			ref:  "projects/my-project/locations/global/wasmPlugins/my-wasmplugin",
			want: &WasmPluginIdentity{
				parent: &parent.ProjectAndLocationParent{
					ProjectID: "my-project",
					Location:  "global",
				},
				id: "my-wasmplugin",
			},
			wantErr: false,
		},
		{
			name:    "invalid format (missing wasmPlugins)",
			ref:     "projects/my-project/locations/global/my-wasmplugin",
			wantErr: true,
		},
		{
			name:    "invalid format (wrong prefix)",
			ref:     "locations/global/wasmPlugins/my-wasmplugin",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &WasmPluginIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("WasmPluginIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.id != tt.want.id || i.parent.ProjectID != tt.want.parent.ProjectID || i.parent.Location != tt.want.parent.Location {
					t.Errorf("WasmPluginIdentity.FromExternal() = %+v, want %+v", i, tt.want)
				}
			}
		})
	}
}

func TestWasmPluginIdentity_GetIdentity(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	if err := AddToScheme(scheme); err != nil {
		t.Fatalf("failed to add to scheme: %v", err)
	}
	cl := fake.NewClientBuilder().WithScheme(scheme).Build()

	obj := &NetworkServicesWasmPlugin{}
	obj.SetName("my-wasmplugin")
	obj.SetNamespace("test-ns")
	obj.Spec.ProjectAndLocationRef = &parent.ProjectAndLocationRef{
		ProjectRef: &refsv1beta1.ProjectRef{
			External: "my-project",
		},
		Location: "global",
	}

	identity, err := obj.GetIdentity(ctx, cl)
	if err != nil {
		t.Fatalf("unexpected error getting identity: %v", err)
	}

	wantStr := "projects/my-project/locations/global/wasmPlugins/my-wasmplugin"
	if identity.String() != wantStr {
		t.Errorf("identity.String() = %q, want %q", identity.String(), wantStr)
	}

	idVal, ok := identity.(*WasmPluginIdentity)
	if !ok {
		t.Fatalf("identity is not of type *WasmPluginIdentity")
	}

	if idVal.ID() != "my-wasmplugin" {
		t.Errorf("identity.ID() = %q, want %q", idVal.ID(), "my-wasmplugin")
	}
}
