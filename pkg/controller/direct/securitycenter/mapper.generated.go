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

package securitycenter

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/securitycenter/apiv2/securitycenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ExternalSystem_FromProto(mapCtx *direct.MapContext, in *pb.ExternalSystem) *krm.ExternalSystem {
	if in == nil {
		return nil
	}
	out := &krm.ExternalSystem{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Assignees = in.Assignees
	out.ExternalUid = direct.LazyPtr(in.GetExternalUid())
	out.Status = direct.LazyPtr(in.GetStatus())
	out.ExternalSystemUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExternalSystemUpdateTime())
	out.CaseURI = direct.LazyPtr(in.GetCaseUri())
	out.CasePriority = direct.LazyPtr(in.GetCasePriority())
	out.CaseSla = direct.StringTimestamp_FromProto(mapCtx, in.GetCaseSla())
	out.CaseCreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCaseCreateTime())
	out.CaseCloseTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCaseCloseTime())
	out.TicketInfo = ExternalSystem_TicketInfo_FromProto(mapCtx, in.GetTicketInfo())
	return out
}
func ExternalSystem_ToProto(mapCtx *direct.MapContext, in *krm.ExternalSystem) *pb.ExternalSystem {
	if in == nil {
		return nil
	}
	out := &pb.ExternalSystem{}
	out.Name = direct.ValueOf(in.Name)
	out.Assignees = in.Assignees
	out.ExternalUid = direct.ValueOf(in.ExternalUid)
	out.Status = direct.ValueOf(in.Status)
	out.ExternalSystemUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.ExternalSystemUpdateTime)
	out.CaseUri = direct.ValueOf(in.CaseURI)
	out.CasePriority = direct.ValueOf(in.CasePriority)
	out.CaseSla = direct.StringTimestamp_ToProto(mapCtx, in.CaseSla)
	out.CaseCreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CaseCreateTime)
	out.CaseCloseTime = direct.StringTimestamp_ToProto(mapCtx, in.CaseCloseTime)
	out.TicketInfo = ExternalSystem_TicketInfo_ToProto(mapCtx, in.TicketInfo)
	return out
}
func ExternalSystem_TicketInfo_FromProto(mapCtx *direct.MapContext, in *pb.ExternalSystem_TicketInfo) *krm.ExternalSystem_TicketInfo {
	if in == nil {
		return nil
	}
	out := &krm.ExternalSystem_TicketInfo{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Assignee = direct.LazyPtr(in.GetAssignee())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.URI = direct.LazyPtr(in.GetUri())
	out.Status = direct.LazyPtr(in.GetStatus())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func ExternalSystem_TicketInfo_ToProto(mapCtx *direct.MapContext, in *krm.ExternalSystem_TicketInfo) *pb.ExternalSystem_TicketInfo {
	if in == nil {
		return nil
	}
	out := &pb.ExternalSystem_TicketInfo{}
	out.Id = direct.ValueOf(in.ID)
	out.Assignee = direct.ValueOf(in.Assignee)
	out.Description = direct.ValueOf(in.Description)
	out.Uri = direct.ValueOf(in.URI)
	out.Status = direct.ValueOf(in.Status)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
