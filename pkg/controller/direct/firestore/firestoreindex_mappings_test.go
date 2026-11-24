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

package firestore

import (
	"context"
	"testing"

	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestFirestoreIndexExportIncludesResourceID(t *testing.T) {
	adapter := &firestoreIndexAdapter{
		id: &krm.FirestoreIndexIdentity{
			Project:         "my-project",
			Database:        "my-database",
			CollectionGroup: "Customer",
			Index:           "CICAgOjXh4EK",
		},
		actual: &pb.Index{},
	}

	obj, err := adapter.Export(context.Background())
	if err != nil {
		t.Fatalf("Export() returned error: %v", err)
	}

	got, found, err := unstructured.NestedString(obj.Object, "spec", "resourceID")
	if err != nil {
		t.Fatalf("reading spec.resourceID: %v", err)
	}
	if !found || got != "CICAgOjXh4EK" {
		t.Fatalf("spec.resourceID = %q, found=%t; want %q", got, found, "CICAgOjXh4EK")
	}
}

func TestFirestoreIndexDefinitionsEqual(t *testing.T) {
	tests := []struct {
		name    string
		desired *pb.Index
		actual  *pb.Index
		want    bool
	}{
		{
			name: "native API adds name field using last order",
			desired: &pb.Index{
				Fields: []*pb.Index_IndexField{indexOrderField("email", pb.Index_IndexField_DESCENDING)},
			},
			actual: &pb.Index{
				QueryScope: pb.Index_COLLECTION,
				Fields: []*pb.Index_IndexField{
					indexOrderField("email", pb.Index_IndexField_DESCENDING),
					indexOrderField("__name__", pb.Index_IndexField_DESCENDING),
				},
				State:   pb.Index_READY,
				Density: pb.Index_SPARSE_ALL,
			},
			want: true,
		},
		{
			name: "datastore API adds ascending name field",
			desired: &pb.Index{
				ApiScope: pb.Index_DATASTORE_MODE_API,
				Fields:   []*pb.Index_IndexField{indexOrderField("createdAt", pb.Index_IndexField_DESCENDING)},
			},
			actual: &pb.Index{
				ApiScope:   pb.Index_DATASTORE_MODE_API,
				QueryScope: pb.Index_COLLECTION,
				Fields: []*pb.Index_IndexField{
					indexOrderField("createdAt", pb.Index_IndexField_DESCENDING),
					indexOrderField("__name__", pb.Index_IndexField_ASCENDING),
				},
			},
			want: true,
		},
		{
			name: "different definition",
			desired: &pb.Index{
				Fields: []*pb.Index_IndexField{indexOrderField("email", pb.Index_IndexField_ASCENDING)},
			},
			actual: &pb.Index{
				Fields: []*pb.Index_IndexField{
					indexOrderField("email", pb.Index_IndexField_DESCENDING),
					indexOrderField("__name__", pb.Index_IndexField_DESCENDING),
				},
			},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := firestoreIndexDefinitionsEqual(tc.desired, tc.actual); got != tc.want {
				t.Fatalf("firestoreIndexDefinitionsEqual() = %t, want %t", got, tc.want)
			}
		})
	}
}

func indexOrderField(path string, order pb.Index_IndexField_Order) *pb.Index_IndexField {
	return &pb.Index_IndexField{
		FieldPath: path,
		ValueMode: &pb.Index_IndexField_Order_{Order: order},
	}
}

func TestFirestoreIndexDatastoreModeMapping(t *testing.T) {
	apiScope := "DATASTORE_MODE_API"
	queryScope := "COLLECTION_RECURSIVE"
	mapCtx := &direct.MapContext{}

	proto := FirestoreIndexSpec_v1beta1_ToProto(mapCtx, &krm.FirestoreIndexSpec{APIScope: &apiScope, QueryScope: &queryScope})
	if err := mapCtx.Err(); err != nil {
		t.Fatalf("mapping FirestoreIndex spec to proto: %v", err)
	}
	if got, want := proto.GetApiScope(), pb.Index_DATASTORE_MODE_API; got != want {
		t.Errorf("api scope mapped to proto as %v, want %v", got, want)
	}
	if got, want := proto.GetQueryScope(), pb.Index_COLLECTION_RECURSIVE; got != want {
		t.Errorf("query scope mapped to proto as %v, want %v", got, want)
	}

	mapCtx = &direct.MapContext{}
	spec := FirestoreIndexSpec_v1beta1_FromProto(mapCtx, &pb.Index{ApiScope: pb.Index_DATASTORE_MODE_API, QueryScope: pb.Index_COLLECTION_RECURSIVE})
	if err := mapCtx.Err(); err != nil {
		t.Fatalf("mapping FirestoreIndex proto to spec: %v", err)
	}
	if spec.APIScope == nil || *spec.APIScope != apiScope {
		t.Errorf("api scope mapped from proto as %v, want %q", spec.APIScope, apiScope)
	}
	if spec.QueryScope == nil || *spec.QueryScope != queryScope {
		t.Errorf("query scope mapped from proto as %v, want %q", spec.QueryScope, queryScope)
	}
}
