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

package configdeployment

import (
	"encoding/json"
	"testing"

	pb "cloud.google.com/go/config/apiv1/configpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/configdeployment/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestArtifactsLocationMapping(t *testing.T) {
	mapCtx := &direct.MapContext{}

	// Test FromProto
	uri := "gs://my-bucket/artifacts/logs"
	loc := ArtifactsLocation_FromProto(mapCtx, uri)
	if loc == nil {
		t.Fatalf("expected non-nil ArtifactsLocation")
	}
	if loc.BucketRef == nil || loc.BucketRef.External != "my-bucket" {
		t.Errorf("expected bucket my-bucket, got %v", loc.BucketRef)
	}
	if loc.Folder == nil || *loc.Folder != "artifacts/logs" {
		t.Errorf("expected folder artifacts/logs, got %v", loc.Folder)
	}

	// Test ToProto
	outURI := ArtifactsLocation_ToProto(mapCtx, loc)
	if outURI == nil || *outURI != uri {
		t.Errorf("expected %s, got %v", uri, outURI)
	}

	// Test Spec mapping
	spec := &krm.ConfigDeploymentSpec{
		ArtifactsLocation: loc,
	}
	protoDep := ConfigDeploymentSpec_ToProto(mapCtx, spec)
	if protoDep.GetArtifactsGcsBucket() != uri {
		t.Errorf("expected %s, got %s", uri, protoDep.GetArtifactsGcsBucket())
	}

	specBack := ConfigDeploymentSpec_FromProto(mapCtx, protoDep)
	if specBack.ArtifactsLocation == nil || specBack.ArtifactsLocation.BucketRef.External != "my-bucket" {
		t.Errorf("expected roundtrip of ArtifactsLocation, got %v", specBack.ArtifactsLocation)
	}
}

func TestOptionalFieldsPreserveNil(t *testing.T) {
	mapCtx := &direct.MapContext{}

	// Unset proto fields should map to nil in KRM
	emptyProto := &pb.Deployment{}
	spec := ConfigDeploymentSpec_FromProto(mapCtx, emptyProto)
	if spec.ImportExistingResources != nil {
		t.Errorf("expected nil ImportExistingResources, got %v", *spec.ImportExistingResources)
	}
	if spec.TfVersionConstraint != nil {
		t.Errorf("expected nil TfVersionConstraint, got %v", *spec.TfVersionConstraint)
	}

	// Set proto fields should preserve values in KRM
	setProto := &pb.Deployment{
		ImportExistingResources: direct.LazyPtr(true),
		TfVersionConstraint:     direct.LazyPtr("=1.3.10"),
	}
	specSet := ConfigDeploymentSpec_FromProto(mapCtx, setProto)
	if specSet.ImportExistingResources == nil || !*specSet.ImportExistingResources {
		t.Errorf("expected true ImportExistingResources, got %v", specSet.ImportExistingResources)
	}
	if specSet.TfVersionConstraint == nil || *specSet.TfVersionConstraint != "=1.3.10" {
		t.Errorf("expected =1.3.10 TfVersionConstraint, got %v", specSet.TfVersionConstraint)
	}
}

func TestValueMapping(t *testing.T) {
	mapCtx := &direct.MapContext{}

	// Test NullValue
	nullVal := structpb.NewNullValue()
	krmNull := Value_FromProto(mapCtx, nullVal)
	if krmNull == nil || krmNull.NullValue == nil {
		t.Fatalf("expected non-nil NullValue")
	}
	protoNull := Value_ToProto(mapCtx, krmNull)
	if _, ok := protoNull.Kind.(*structpb.Value_NullValue); !ok {
		t.Errorf("expected Value_NullValue, got %T", protoNull.Kind)
	}

	// Test StringValue
	strVal := structpb.NewStringValue("hello")
	krmStr := Value_FromProto(mapCtx, strVal)
	if krmStr == nil || krmStr.StringValue == nil || *krmStr.StringValue != "hello" {
		t.Fatalf("expected hello string, got %v", krmStr)
	}
	protoStr := Value_ToProto(mapCtx, krmStr)
	if protoStr.GetStringValue() != "hello" {
		t.Errorf("expected hello, got %s", protoStr.GetStringValue())
	}

	// Test NumberValue
	numVal := structpb.NewNumberValue(42.5)
	krmNum := Value_FromProto(mapCtx, numVal)
	if krmNum == nil || krmNum.NumberValue == nil || *krmNum.NumberValue != 42.5 {
		t.Fatalf("expected 42.5, got %v", krmNum)
	}
	protoNum := Value_ToProto(mapCtx, krmNum)
	if protoNum.GetNumberValue() != 42.5 {
		t.Errorf("expected 42.5, got %f", protoNum.GetNumberValue())
	}

	// Test BoolValue
	boolVal := structpb.NewBoolValue(true)
	krmBool := Value_FromProto(mapCtx, boolVal)
	if krmBool == nil || krmBool.BoolValue == nil || !*krmBool.BoolValue {
		t.Fatalf("expected true, got %v", krmBool)
	}
	protoBool := Value_ToProto(mapCtx, krmBool)
	if !protoBool.GetBoolValue() {
		t.Errorf("expected true, got %v", protoBool.GetBoolValue())
	}

	// Test StructValue
	st, err := structpb.NewStruct(map[string]interface{}{
		"key": "value",
		"nested": map[string]interface{}{
			"count": 3.0,
		},
	})
	if err != nil {
		t.Fatalf("failed to create structpb.Struct: %v", err)
	}
	stVal := structpb.NewStructValue(st)
	krmStruct := Value_FromProto(mapCtx, stVal)
	if krmStruct == nil || len(krmStruct.StructValue.Raw) == 0 {
		t.Fatalf("expected non-empty StructValue, got %v", krmStruct)
	}

	var parsedMap map[string]interface{}
	if err := json.Unmarshal(krmStruct.StructValue.Raw, &parsedMap); err != nil {
		t.Fatalf("failed to unmarshal StructValue: %v", err)
	}
	if parsedMap["key"] != "value" {
		t.Errorf("expected key=value, got %v", parsedMap["key"])
	}

	protoStructBack := Value_ToProto(mapCtx, krmStruct)
	if protoStructBack == nil {
		t.Fatalf("expected non-nil protoStructBack")
	}
	stBackKind, ok := protoStructBack.Kind.(*structpb.Value_StructValue)
	if !ok {
		t.Fatalf("expected Value_StructValue, got %T", protoStructBack.Kind)
	}
	if stBackKind.StructValue.Fields["key"].GetStringValue() != "value" {
		t.Errorf("expected key=value in struct, got %v", stBackKind.StructValue.Fields["key"])
	}
}
