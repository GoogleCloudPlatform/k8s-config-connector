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

func TestFirestoreDatabaseTypeMapping(t *testing.T) {
	databaseType := "DATASTORE_MODE"
	mapCtx := &direct.MapContext{}

	proto := FirestoreDatabaseSpec_v1beta1_ToProto(mapCtx, &krm.FirestoreDatabaseSpec{Type: &databaseType})
	if err := mapCtx.Err(); err != nil {
		t.Fatalf("mapping FirestoreDatabase spec to proto: %v", err)
	}
	if got, want := proto.GetType(), pb.Database_DATASTORE_MODE; got != want {
		t.Errorf("database type mapped to proto as %v, want %v", got, want)
	}

	mapCtx = &direct.MapContext{}
	spec := FirestoreDatabaseSpec_v1beta1_FromProto(mapCtx, &pb.Database{Type: pb.Database_DATASTORE_MODE})
	if err := mapCtx.Err(); err != nil {
		t.Fatalf("mapping FirestoreDatabase proto to spec: %v", err)
	}
	if spec.Type == nil || *spec.Type != databaseType {
		t.Errorf("database type mapped from proto as %v, want %q", spec.Type, databaseType)
	}
}
