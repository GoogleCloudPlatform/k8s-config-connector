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

package bigtable

import (
	"testing"

	gcp "cloud.google.com/go/bigtable"
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
)

func strPtr(s string) *string { return &s }
func boolPtr(b bool) *bool    { return &b }

func TestBigtableMaterializedViewSpec_ToMaterializedViewInfo(t *testing.T) {
	t.Run("nil input", func(t *testing.T) {
		if got := BigtableMaterializedViewSpec_ToMaterializedViewInfo(nil, nil); got != nil {
			t.Fatalf("expected nil, got %#v", got)
		}
	})

	t.Run("deletionProtection nil", func(t *testing.T) {
		in := &krmv1alpha1.BigtableMaterializedViewSpec{
			Query:              strPtr("SELECT * FROM tbl"),
			DeletionProtection: nil,
		}
		got := BigtableMaterializedViewSpec_ToMaterializedViewInfo(nil, in)
		if got == nil {
			t.Fatalf("unexpected nil result")
		}
		if got.Query != *in.Query {
			t.Errorf("Query: got %q, want %q", got.Query, *in.Query)
		}
		if got.DeletionProtection != gcp.None {
			t.Errorf("DeletionProtection: got %v, want gcp.None", got.DeletionProtection)
		}
	})

	t.Run("deletionProtection true", func(t *testing.T) {
		in := &krmv1alpha1.BigtableMaterializedViewSpec{
			Query:              strPtr("SELECT 1"),
			DeletionProtection: boolPtr(true),
		}
		got := BigtableMaterializedViewSpec_ToMaterializedViewInfo(nil, in)
		if got == nil {
			t.Fatalf("unexpected nil result")
		}
		if got.DeletionProtection != gcp.Protected {
			t.Errorf("DeletionProtection: got %v, want gcp.Protected", got.DeletionProtection)
		}
	})

	t.Run("deletionProtection false", func(t *testing.T) {
		in := &krmv1alpha1.BigtableMaterializedViewSpec{
			Query:              strPtr("SELECT 2"),
			DeletionProtection: boolPtr(false),
		}
		got := BigtableMaterializedViewSpec_ToMaterializedViewInfo(nil, in)
		if got == nil {
			t.Fatalf("unexpected nil result")
		}
		if got.DeletionProtection != gcp.Unprotected {
			t.Errorf("DeletionProtection: got %v, want gcp.Unprotected", got.DeletionProtection)
		}
	})
}

func TestBigtableMaterializedViewInfo_ToBigtableMaterializedView(t *testing.T) {
	t.Run("nil input", func(t *testing.T) {
		if got := BigtableMaterializedViewInfo_ToBigtableMaterializedView(nil, nil); got != nil {
			t.Fatalf("expected nil, got %#v", got)
		}
	})

	t.Run("protected maps to true", func(t *testing.T) {
		in := &gcp.MaterializedViewInfo{
			Query:              "q",
			DeletionProtection: gcp.Protected,
		}
		got := BigtableMaterializedViewInfo_ToBigtableMaterializedView(nil, in)
		if got == nil {
			t.Fatalf("unexpected nil result")
		}
		if got.Query != in.Query {
			t.Errorf("Query: got %q, want %q", got.Query, in.Query)
		}
		if got.DeletionProtection != true {
			t.Errorf("DeletionProtection: got %v, want true", got.DeletionProtection)
		}
	})

	t.Run("none/unprotected maps to false", func(t *testing.T) {
		cases := []gcp.DeletionProtection{gcp.None, gcp.Unprotected}
		for _, c := range cases {
			in := &gcp.MaterializedViewInfo{
				Query:              "qq",
				DeletionProtection: c,
			}
			got := BigtableMaterializedViewInfo_ToBigtableMaterializedView(nil, in)
			if got == nil {
				t.Fatalf("unexpected nil result for case %v", c)
			}
			if got.DeletionProtection != false {
				t.Errorf("case %v: DeletionProtection: got %v, want false", c, got.DeletionProtection)
			}
		}
	})
}

func TestBigtableMaterializedView_ToBigtableMaterializedViewInfo(t *testing.T) {
	t.Run("nil input", func(t *testing.T) {
		if got := BigtableMaterializedView_ToBigtableMaterializedViewInfo(nil, nil); got != nil {
			t.Fatalf("expected nil, got: %#v", got)
		}
	})

	t.Run("DeletionProtection true -> Protected", func(t *testing.T) {
		in := &pb.MaterializedView{
			Query:              "pv1",
			DeletionProtection: true,
		}
		got := BigtableMaterializedView_ToBigtableMaterializedViewInfo(nil, in)
		if got == nil {
			t.Fatalf("unexpected nil result")
		}
		if got.Query != in.Query {
			t.Errorf("Query mismatch: got %q, want %q", got.Query, in.Query)
		}
		if got.DeletionProtection != gcp.Protected {
			t.Errorf("DeletionProtection mismatch: got %v, want %v", got.DeletionProtection, gcp.Protected)
		}
	})

	t.Run("DeletionProtection false -> Unprotected", func(t *testing.T) {
		in := &pb.MaterializedView{
			Query:              "pv2",
			DeletionProtection: false,
		}
		got := BigtableMaterializedView_ToBigtableMaterializedViewInfo(nil, in)
		if got == nil {
			t.Fatalf("unexpected nil result")
		}
		if got.Query != in.Query {
			t.Errorf("Query mismatch: got %q, want %q", got.Query, in.Query)
		}
		if got.DeletionProtection != gcp.Unprotected {
			t.Errorf("DeletionProtection mismatch: got %v, want %v", got.DeletionProtection, gcp.Unprotected)
		}
	})
}
