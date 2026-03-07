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

package bigtable

import (
	"reflect"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TestBigtableSchemaBundleMapper(t *testing.T) {
	mapCtx := &direct.MapContext{}

	protoDescriptors := []byte("test-descriptors")
	krmSpec := &krm.BigtableSchemaBundleSpec{
		ProtoSchema: &krm.ProtoSchema{
			ProtoDescriptors: protoDescriptors,
		},
	}

	protoObj := BigtableSchemaBundleSpec_v1alpha1_ToProto(mapCtx, krmSpec)
	if mapCtx.Err() != nil {
		t.Fatalf("Spec to Proto failed: %v", mapCtx.Err())
	}

	if !reflect.DeepEqual(protoObj.GetProtoSchema().ProtoDescriptors, protoDescriptors) {
		t.Errorf("Proto descriptors mismatch, got %v, want %v", protoObj.GetProtoSchema().ProtoDescriptors, protoDescriptors)
	}

	krmSpec2 := BigtableSchemaBundleSpec_v1alpha1_FromProto(mapCtx, protoObj)
	if mapCtx.Err() != nil {
		t.Fatalf("Proto to Spec failed: %v", mapCtx.Err())
	}

	if !reflect.DeepEqual(krmSpec2.ProtoSchema.ProtoDescriptors, protoDescriptors) {
		t.Errorf("KRM descriptors mismatch, got %v, want %v", krmSpec2.ProtoSchema.ProtoDescriptors, protoDescriptors)
	}
}
