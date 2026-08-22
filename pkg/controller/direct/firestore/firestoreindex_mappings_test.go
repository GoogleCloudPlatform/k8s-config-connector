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
	"testing"

	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

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
