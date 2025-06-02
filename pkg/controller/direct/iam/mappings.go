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

package iam

import (
	"strings"

	"cloud.google.com/go/iam/apiv1/iampb"
	expr "google.golang.org/genproto/googleapis/type/expr"

	newiamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"
	oldiamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

const (
	LogTypeUnspecified = "LOG_TYPE_UNSPECIFIED"
	AdminRead          = "ADMIN_READ"
	DataWrite          = "DATA_WRITE"
	DataRead           = "DATA_READ"
)

func IAMPolicySpec_ToProto(_ *direct.MapContext, in *oldiamv1beta1.IAMPolicySpec) *iampb.Policy {
	if in == nil {
		return nil
	}

	protoPolicy := &iampb.Policy{
		// Default to version 3, which supports conditions.
		// IAM Server should downgrade to supported versions
		Version: 3,
		//Etag: spec.Etag, NOT YET
	}

	// Map Bindings
	if len(in.Bindings) > 0 {
		protoPolicy.Bindings = make([]*iampb.Binding, 0, len(in.Bindings))
		for _, b := range in.Bindings {
			pbBinding := &iampb.Binding{
				Role:    b.Role,
				Members: make([]string, len(b.Members)),
			}
			for i, member := range b.Members {
				pbBinding.Members[i] = string(member)
			}

			if b.Condition != nil {
				pbBinding.Condition = &expr.Expr{
					Expression:  b.Condition.Expression,
					Title:       b.Condition.Title,
					Description: b.Condition.Description,
					// Location: NOT yet, would beed adapter/ id for this; current types don't hold this information
				}
			}
			protoPolicy.Bindings = append(protoPolicy.Bindings, pbBinding)
		}
	}

	// Map AuditConfigs
	if len(in.AuditConfigs) > 0 {
		protoPolicy.AuditConfigs = make([]*iampb.AuditConfig, 0, len(in.AuditConfigs))
		for _, ac := range in.AuditConfigs {
			pbAuditConfig := &iampb.AuditConfig{
				Service: ac.Service,
			}

			if len(ac.AuditLogConfigs) > 0 {
				pbAuditConfig.AuditLogConfigs = make([]*iampb.AuditLogConfig, 0, len(ac.AuditLogConfigs))
				for _, alc := range ac.AuditLogConfigs {
					pbAlc := &iampb.AuditLogConfig{
						LogType: mapV1Beta1LogTypeToProto(alc.LogType),
					}
					if len(alc.ExemptedMembers) > 0 {
						pbAlc.ExemptedMembers = make([]string, len(alc.ExemptedMembers))
						for i, em := range alc.ExemptedMembers {
							pbAlc.ExemptedMembers[i] = string(em)
						}
					}
					pbAuditConfig.AuditLogConfigs = append(pbAuditConfig.AuditLogConfigs, pbAlc)
				}
			}
			protoPolicy.AuditConfigs = append(protoPolicy.AuditConfigs, pbAuditConfig)
		}
	}

	return protoPolicy
}

func IAMPolicySpec_FromProto(_ *direct.MapContext, in *iampb.Policy) *newiamv1beta1.IAMPolicySpec {
	if in == nil {
		return nil
	}

	out := &newiamv1beta1.IAMPolicySpec{
		Etag: string(in.Etag),
	}

	// Map Bindings from Proto to KRM
	if len(in.Bindings) > 0 {
		out.Bindings = make([]newiamv1beta1.IAMPolicyBinding, 0, len(in.Bindings))
		for _, pbBinding := range in.Bindings {
			binding := newiamv1beta1.IAMPolicyBinding{
				Role:    pbBinding.Role,
				Members: make([]newiamv1beta1.Member, len(pbBinding.Members)),
			}
			for i, member := range pbBinding.Members {
				binding.Members[i] = newiamv1beta1.Member(member)
			}

			if pbBinding.Condition != nil {
				binding.Condition = &newiamv1beta1.IAMCondition{
					Expression:  pbBinding.Condition.Expression,
					Title:       pbBinding.Condition.Title,
					Description: pbBinding.Condition.Description,
				}
			}
			out.Bindings = append(out.Bindings, binding)
		}
	}

	// Map AuditConfigs from Proto to KRM
	if len(in.AuditConfigs) > 0 {
		out.AuditConfigs = make([]newiamv1beta1.IAMPolicyAuditConfig, 0, len(in.AuditConfigs))
		for _, pbAuditConfig := range in.AuditConfigs {
			ac := newiamv1beta1.IAMPolicyAuditConfig{
				Service: pbAuditConfig.Service,
			}

			if len(pbAuditConfig.AuditLogConfigs) > 0 {
				ac.AuditLogConfigs = make([]newiamv1beta1.AuditLogConfig, 0, len(pbAuditConfig.AuditLogConfigs))
				for _, pbAlc := range pbAuditConfig.AuditLogConfigs {
					logTypeString := mapProtoLogTypeToKRM(pbAlc.LogType)

					alc := newiamv1beta1.AuditLogConfig{
						LogType: logTypeString,
					}
					if len(pbAlc.ExemptedMembers) > 0 {
						alc.ExemptedMembers = make([]newiamv1beta1.Member, len(pbAlc.ExemptedMembers))
						for i, em := range pbAlc.ExemptedMembers {
							alc.ExemptedMembers[i] = newiamv1beta1.Member(em)
						}
					}
					ac.AuditLogConfigs = append(ac.AuditLogConfigs, alc)
				}
			}
			out.AuditConfigs = append(out.AuditConfigs, ac)
		}
	}

	return out
}

// mapV1Beta1LogTypeToProto converts a string log type from v1beta1
// to the pb.AuditLogConfig_LogType enum (which is a wrapped int).
func mapV1Beta1LogTypeToProto(logTypeString string) iampb.AuditLogConfig_LogType {
	switch strings.ToUpper(logTypeString) {
	case LogTypeUnspecified:
		return iampb.AuditLogConfig_LOG_TYPE_UNSPECIFIED
	case AdminRead:
		return iampb.AuditLogConfig_ADMIN_READ
	case DataWrite:
		return iampb.AuditLogConfig_DATA_WRITE
	case DataRead:
		return iampb.AuditLogConfig_DATA_READ
	default:
		return iampb.AuditLogConfig_LOG_TYPE_UNSPECIFIED
	}
}

func mapProtoLogTypeToKRM(logTypeEnum iampb.AuditLogConfig_LogType) string {
	switch logTypeEnum {
	case iampb.AuditLogConfig_LOG_TYPE_UNSPECIFIED:
		return LogTypeUnspecified
	case iampb.AuditLogConfig_ADMIN_READ:
		return AdminRead
	case iampb.AuditLogConfig_DATA_WRITE:
		return DataWrite
	case iampb.AuditLogConfig_DATA_READ:
		return DataRead
	default:
		return LogTypeUnspecified
	}
}
