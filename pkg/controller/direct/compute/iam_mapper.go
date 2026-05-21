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

package compute

import (
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	iampb "cloud.google.com/go/iam/apiv1/iampb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/genproto/googleapis/type/expr"
)

func IAMPolicy_FromProto(mapCtx *direct.MapContext, in *computepb.Policy) *iampb.Policy {
	if in == nil {
		return nil
	}
	out := &iampb.Policy{}
	for _, b := range in.Bindings {
		out.Bindings = append(out.Bindings, IAMBinding_FromProto(mapCtx, b))
	}
	for _, a := range in.AuditConfigs {
		out.AuditConfigs = append(out.AuditConfigs, IAMAuditConfig_FromProto(mapCtx, a))
	}
	if in.Etag != nil {
		out.Etag = []byte(*in.Etag)
	}
	if in.Version != nil {
		out.Version = *in.Version
	}
	return out
}

func IAMPolicy_ToProto(mapCtx *direct.MapContext, in *iampb.Policy) *computepb.Policy {
	if in == nil {
		return nil
	}
	out := &computepb.Policy{}
	for _, b := range in.Bindings {
		out.Bindings = append(out.Bindings, IAMBinding_ToProto(mapCtx, b))
	}
	for _, a := range in.AuditConfigs {
		out.AuditConfigs = append(out.AuditConfigs, IAMAuditConfig_ToProto(mapCtx, a))
	}
	if len(in.Etag) > 0 {
		out.Etag = direct.LazyPtr(string(in.Etag))
	}
	out.Version = direct.LazyPtr(in.Version)
	return out
}

func IAMBinding_FromProto(mapCtx *direct.MapContext, in *computepb.Binding) *iampb.Binding {
	if in == nil {
		return nil
	}
	out := &iampb.Binding{}
	out.Role = direct.ValueOf(in.Role)
	out.Members = in.Members
	out.Condition = IAMCondition_FromProto(mapCtx, in.Condition)
	return out
}

func IAMBinding_ToProto(mapCtx *direct.MapContext, in *iampb.Binding) *computepb.Binding {
	if in == nil {
		return nil
	}
	out := &computepb.Binding{}
	out.Role = direct.LazyPtr(in.Role)
	out.Members = in.Members
	out.Condition = IAMCondition_ToProto(mapCtx, in.Condition)
	return out
}

func IAMAuditConfig_FromProto(mapCtx *direct.MapContext, in *computepb.AuditConfig) *iampb.AuditConfig {
	if in == nil {
		return nil
	}
	out := &iampb.AuditConfig{}
	out.Service = direct.ValueOf(in.Service)
	for _, c := range in.AuditLogConfigs {
		out.AuditLogConfigs = append(out.AuditLogConfigs, IAMAuditLogConfig_FromProto(mapCtx, c))
	}
	return out
}

func IAMAuditConfig_ToProto(mapCtx *direct.MapContext, in *iampb.AuditConfig) *computepb.AuditConfig {
	if in == nil {
		return nil
	}
	out := &computepb.AuditConfig{}
	out.Service = direct.LazyPtr(in.Service)
	for _, c := range in.AuditLogConfigs {
		out.AuditLogConfigs = append(out.AuditLogConfigs, IAMAuditLogConfig_ToProto(mapCtx, c))
	}
	return out
}

func IAMAuditLogConfig_FromProto(mapCtx *direct.MapContext, in *computepb.AuditLogConfig) *iampb.AuditLogConfig {
	if in == nil {
		return nil
	}
	out := &iampb.AuditLogConfig{}
	if in.LogType != nil {
		out.LogType = iampb.AuditLogConfig_LogType(iampb.AuditLogConfig_LogType_value[*in.LogType])
	}
	out.ExemptedMembers = in.ExemptedMembers
	return out
}

func IAMAuditLogConfig_ToProto(mapCtx *direct.MapContext, in *iampb.AuditLogConfig) *computepb.AuditLogConfig {
	if in == nil {
		return nil
	}
	out := &computepb.AuditLogConfig{}
	out.LogType = direct.LazyPtr(in.LogType.String())
	out.ExemptedMembers = in.ExemptedMembers
	return out
}

func IAMCondition_FromProto(mapCtx *direct.MapContext, in *computepb.Expr) *expr.Expr {
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

func IAMCondition_ToProto(mapCtx *direct.MapContext, in *expr.Expr) *computepb.Expr {
	if in == nil {
		return nil
	}
	out := &computepb.Expr{}
	out.Description = direct.LazyPtr(in.Description)
	out.Expression = direct.LazyPtr(in.Expression)
	out.Location = direct.LazyPtr(in.Location)
	out.Title = direct.LazyPtr(in.Title)
	return out
}
