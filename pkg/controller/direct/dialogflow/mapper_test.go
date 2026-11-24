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

package dialogflow

import (
	"encoding/json"
	"reflect"
	"testing"

	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/structpb"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/utils/ptr"
)

func TestResponseMessage_ToolCall_ToProto(t *testing.T) {
	tool := "projects/test-project/locations/global/agents/agent-1/tools/tool-1"
	action := "queryAction"
	inputParamsJSON := `{"city":"Seattle"}`

	krmMsg := &krm.ResponseMessage{
		ToolCall: &krm.ToolCall{
			Tool:            ptr.To(tool),
			Action:          ptr.To(action),
			InputParameters: apiextensionsv1.JSON{Raw: []byte(inputParamsJSON)},
		},
	}

	mapCtx := &direct.MapContext{}
	protoMsg := ResponseMessage_ToProto(mapCtx, krmMsg)
	if err := mapCtx.Err(); err != nil {
		t.Fatalf("unexpected error in ResponseMessage_ToProto: %v", err)
	}

	if protoMsg == nil {
		t.Fatalf("expected non-nil proto message")
	}

	toolCall := protoMsg.GetToolCall()
	if toolCall == nil {
		t.Fatalf("expected non-nil ToolCall in ResponseMessage proto, got nil (Message: %v)", protoMsg.GetMessage())
	}

	if got := toolCall.GetTool(); got != tool {
		t.Errorf("Tool mismatch: got %q, want %q", got, tool)
	}
	if got := toolCall.GetAction(); got != action {
		t.Errorf("Action mismatch: got %q, want %q", got, action)
	}

	expectedStruct, err := structpb.NewStruct(map[string]interface{}{"city": "Seattle"})
	if err != nil {
		t.Fatalf("failed to create expected structpb: %v", err)
	}
	if got := toolCall.GetInputParameters().AsMap(); !reflect.DeepEqual(got, expectedStruct.AsMap()) {
		t.Errorf("InputParameters mismatch: got %v, want %v", got, expectedStruct.AsMap())
	}
}

func TestResponseMessage_ToolCall_FromProto(t *testing.T) {
	tool := "projects/test-project/locations/global/agents/agent-1/tools/tool-1"
	action := "queryAction"
	inputParams, err := structpb.NewStruct(map[string]interface{}{"city": "Seattle"})
	if err != nil {
		t.Fatalf("failed to create structpb: %v", err)
	}

	protoMsg := &pb.ResponseMessage{
		Message: &pb.ResponseMessage_ToolCall{
			ToolCall: &pb.ToolCall{
				Tool:            tool,
				Action:          action,
				InputParameters: inputParams,
			},
		},
	}

	mapCtx := &direct.MapContext{}
	krmMsg := ResponseMessage_FromProto(mapCtx, protoMsg)
	if err := mapCtx.Err(); err != nil {
		t.Fatalf("unexpected error in ResponseMessage_FromProto: %v", err)
	}

	if krmMsg == nil {
		t.Fatalf("expected non-nil KRM ResponseMessage")
	}

	if krmMsg.ToolCall == nil {
		t.Fatalf("expected non-nil ToolCall in KRM ResponseMessage")
	}

	if got := ptr.Deref(krmMsg.ToolCall.Tool, ""); got != tool {
		t.Errorf("Tool mismatch: got %q, want %q", got, tool)
	}
	if got := ptr.Deref(krmMsg.ToolCall.Action, ""); got != action {
		t.Errorf("Action mismatch: got %q, want %q", got, action)
	}

	var gotMap map[string]interface{}
	if err := json.Unmarshal(krmMsg.ToolCall.InputParameters.Raw, &gotMap); err != nil {
		t.Fatalf("failed to unmarshal KRM ToolCall InputParameters: %v", err)
	}
	if !reflect.DeepEqual(gotMap, map[string]interface{}{"city": "Seattle"}) {
		t.Errorf("InputParameters JSON mismatch: got %v, want %v", gotMap, map[string]interface{}{"city": "Seattle"})
	}
}

func TestResponseMessage_ToolCall_NilAndEmpty(t *testing.T) {
	mapCtx := &direct.MapContext{}

	// Nil ResponseMessage
	if got := ResponseMessage_ToProto(mapCtx, nil); got != nil {
		t.Errorf("expected nil proto from nil KRM ResponseMessage, got %v", got)
	}
	if got := ResponseMessage_FromProto(mapCtx, nil); got != nil {
		t.Errorf("expected nil KRM from nil proto ResponseMessage, got %v", got)
	}

	// ResponseMessage with nil ToolCall
	emptyKrm := &krm.ResponseMessage{}
	emptyProto := ResponseMessage_ToProto(mapCtx, emptyKrm)
	if emptyProto == nil {
		t.Fatalf("expected non-nil proto for empty KRM ResponseMessage")
	}
	if emptyProto.GetToolCall() != nil {
		t.Errorf("expected nil ToolCall in proto, got %v", emptyProto.GetToolCall())
	}

	// ToolCall nil
	if got := ToolCall_ToProto(mapCtx, nil); got != nil {
		t.Errorf("expected nil proto from nil ToolCall, got %v", got)
	}
	if got := ToolCall_FromProto(mapCtx, nil); got != nil {
		t.Errorf("expected nil KRM from nil ToolCall, got %v", got)
	}
}
