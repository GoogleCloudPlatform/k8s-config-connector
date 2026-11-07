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

package dialogflowcx

import (
	"testing"

	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krmdialogflowcxv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflowcx/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/structpb"
	"k8s.io/utils/ptr"
)

func TestParameterDefaultValueMapping(t *testing.T) {
	mapCtx := &direct.MapContext{}

	// Test nil
	if got := Parameter_DefaultValue_FromProto(mapCtx, nil); got != nil {
		t.Errorf("expected nil, got %v", got)
	}
	if got := Parameter_DefaultValue_ToProto(mapCtx, nil); got != nil {
		t.Errorf("expected nil, got %v", got)
	}

	// Test string value
	val := structpb.NewStringValue("test-val")
	s := Parameter_DefaultValue_FromProto(mapCtx, val)
	if s == nil || *s != `"test-val"` {
		t.Errorf("unexpected fromProto result: %v", s)
	}
	toVal := Parameter_DefaultValue_ToProto(mapCtx, ptr.To("test-val"))
	if toVal.GetStringValue() != "test-val" {
		t.Errorf("unexpected toProto result: %v", toVal)
	}
}

func TestResponseMessageMapping(t *testing.T) {
	mapCtx := &direct.MapContext{}

	// Test nil
	if got := ResponseMessage_v1alpha1_FromProto(mapCtx, nil); got != nil {
		t.Errorf("expected nil, got %v", got)
	}
	if got := ResponseMessage_v1alpha1_ToProto(mapCtx, nil); got != nil {
		t.Errorf("expected nil, got %v", got)
	}

	// Test with text and channel
	protoMsg := &pb.ResponseMessage{
		Message: &pb.ResponseMessage_Text_{
			Text: &pb.ResponseMessage_Text{
				Text: []string{"hello world"},
			},
		},
		Channel: "test-channel",
	}

	krmMsg := ResponseMessage_v1alpha1_FromProto(mapCtx, protoMsg)
	if krmMsg == nil {
		t.Fatalf("expected non-nil KRM response message")
	}
	if krmMsg.Channel == nil || *krmMsg.Channel != "test-channel" {
		t.Errorf("expected channel 'test-channel', got %v", krmMsg.Channel)
	}

	backProto := ResponseMessage_v1alpha1_ToProto(mapCtx, krmMsg)
	if backProto == nil || backProto.GetChannel() != "test-channel" {
		t.Errorf("expected roundtrip channel 'test-channel', got %v", backProto)
	}
}

func TestDialogflowCXPageSpecMapping(t *testing.T) {
	mapCtx := &direct.MapContext{}

	protoPage := &pb.Page{
		DisplayName: "test-page",
	}

	var krmSpec *krmdialogflowcxv1alpha1.DialogflowCXPageSpec = DialogflowCXPageSpec_v1alpha1_FromProto(mapCtx, protoPage)
	if krmSpec == nil {
		t.Fatalf("expected non-nil KRM spec")
	}
	if krmSpec.DisplayName != "test-page" {
		t.Errorf("expected displayName 'test-page', got %v", krmSpec.DisplayName)
	}

	backProto := DialogflowCXPageSpec_v1alpha1_ToProto(mapCtx, krmSpec)
	if backProto == nil || backProto.GetDisplayName() != "test-page" {
		t.Errorf("expected roundtrip displayName 'test-page', got %v", backProto)
	}
}
