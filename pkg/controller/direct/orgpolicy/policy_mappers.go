// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.orgpolicy.v2.Policy
// api.group: orgpolicy.cnrm.cloud.google.com

package orgpolicy

import (
	"encoding/json"

	pb "cloud.google.com/go/orgpolicy/apiv2/orgpolicypb"
	"google.golang.org/genproto/googleapis/type/expr"
	"google.golang.org/protobuf/types/known/structpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/orgpolicy/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func PolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.PolicySpec) *krm.PolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.PolicySpec{}
	// MISSING: Etag
	// MISSING: UpdateTime
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, PolicySpec_PolicyRule_FromProto)
	out.InheritFromParent = direct.LazyPtr(in.GetInheritFromParent())
	// Expected code `in.GetReset()`, actual code `in.GetReset_()`
	out.Reset = direct.LazyPtr(in.GetReset_())
	return out
}
func PolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.PolicySpec) *pb.PolicySpec {
	if in == nil {
		return nil
	}
	out := &pb.PolicySpec{}
	// MISSING: Etag
	// MISSING: UpdateTime
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, PolicySpec_PolicyRule_ToProto)
	out.InheritFromParent = direct.ValueOf(in.InheritFromParent)
	// Expected code `in.GetReset()`, actual code `in.GetReset_()`
	out.Reset_ = direct.ValueOf(in.Reset)
	return out
}

func Expr_FromProto(mapCtx *direct.MapContext, in *expr.Expr) *krm.Expr {
	if in == nil {
		return nil
	}
	out := &krm.Expr{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Expression = direct.LazyPtr(in.GetExpression())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Title = direct.LazyPtr(in.GetTitle())
	return out
}

func Expr_ToProto(mapCtx *direct.MapContext, in *krm.Expr) *expr.Expr {
	if in == nil {
		return nil
	}
	out := &expr.Expr{}
	out.Description = direct.ValueOf(in.Description)
	out.Expression = direct.ValueOf(in.Expression)
	out.Location = direct.ValueOf(in.Location)
	out.Title = direct.ValueOf(in.Title)
	return out
}

func PolicySpec_PolicyRule_FromProto(mapCtx *direct.MapContext, in *pb.PolicySpec_PolicyRule) *krm.PolicySpec_PolicyRule {
	if in == nil {
		return nil
	}
	out := &krm.PolicySpec_PolicyRule{}
	if in.GetKind() != nil {
		switch kind := in.GetKind().(type) {
		case *pb.PolicySpec_PolicyRule_Values:
			out.Values = PolicySpec_PolicyRule_StringValues_FromProto(mapCtx, kind.Values)
		case *pb.PolicySpec_PolicyRule_AllowAll:
			out.AllowAll = &kind.AllowAll
		case *pb.PolicySpec_PolicyRule_DenyAll:
			out.DenyAll = &kind.DenyAll
		case *pb.PolicySpec_PolicyRule_Enforce:
			out.Enforce = &kind.Enforce
		default:
			mapCtx.Errorf("unknown oneof kind %T", kind)
			return nil
		}
	}
	out.Condition = Expr_FromProto(mapCtx, in.GetCondition())
	out.Parameters = PolicySpec_PolicyRule_Parameters_FromProto(mapCtx, in.GetParameters())
	return out
}

func PolicySpec_PolicyRule_ToProto(mapCtx *direct.MapContext, in *krm.PolicySpec_PolicyRule) *pb.PolicySpec_PolicyRule {
	if in == nil {
		return nil
	}
	out := &pb.PolicySpec_PolicyRule{}
	if in.Values != nil {
		out.Kind = &pb.PolicySpec_PolicyRule_Values{Values: PolicySpec_PolicyRule_StringValues_ToProto(mapCtx, in.Values)}
	} else if in.AllowAll != nil {
		out.Kind = &pb.PolicySpec_PolicyRule_AllowAll{AllowAll: direct.ValueOf(in.AllowAll)}
	} else if in.DenyAll != nil {
		out.Kind = &pb.PolicySpec_PolicyRule_DenyAll{DenyAll: direct.ValueOf(in.DenyAll)}
	} else if in.Enforce != nil {
		out.Kind = &pb.PolicySpec_PolicyRule_Enforce{Enforce: direct.ValueOf(in.Enforce)}
	} else {
		out.Kind = nil
	}
	out.Condition = Expr_ToProto(mapCtx, in.Condition)
	out.Parameters = PolicySpec_PolicyRule_Parameters_ToProto(mapCtx, in.Parameters)
	return out
}

// PolicySpec_PolicyRule_Parameters_FromProto converts the dynamic
// google.protobuf.Struct returned by the Org Policy API into the typed
// KRM struct. We round-trip via JSON: structpb.Struct -> map[string]any
// -> json bytes -> typed struct. Unknown keys are dropped (rejected at
// the CRD layer on user input, but if the server returns a key we don't
// model yet we want the controller to keep working).
func PolicySpec_PolicyRule_Parameters_FromProto(mapCtx *direct.MapContext, in *structpb.Struct) *krm.PolicySpec_PolicyRule_Parameters {
	if in == nil {
		return nil
	}
	b, err := json.Marshal(in.AsMap())
	if err != nil {
		mapCtx.Errorf("marshalling structpb.Struct to json: %v", err)
		return nil
	}
	out := &krm.PolicySpec_PolicyRule_Parameters{}
	if err := json.Unmarshal(b, out); err != nil {
		mapCtx.Errorf("unmarshalling parameters: %v", err)
		return nil
	}
	return out
}

// PolicySpec_PolicyRule_Parameters_ToProto converts the typed KRM struct
// into a google.protobuf.Struct. We round-trip via JSON to flatten the
// struct into a map with omitempty fields skipped, then build the
// structpb.Struct from that map.
func PolicySpec_PolicyRule_Parameters_ToProto(mapCtx *direct.MapContext, in *krm.PolicySpec_PolicyRule_Parameters) *structpb.Struct {
	if in == nil {
		return nil
	}
	b, err := json.Marshal(in)
	if err != nil {
		mapCtx.Errorf("marshalling parameters to json: %v", err)
		return nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		mapCtx.Errorf("unmarshalling json to map: %v", err)
		return nil
	}
	s, err := structpb.NewStruct(m)
	if err != nil {
		mapCtx.Errorf("error converting map to structpb.Struct: %v", err)
		return nil
	}
	return s
}
