// Copyright 2024 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"context"
	"errors"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	binaryauthorizationpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/binaryauthorization/binaryauthorization_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization"
)

// PolicyServer implements the gRPC interface for Policy.
type PolicyServer struct{}

// ProtoToPolicyClusterAdmissionRulesEvaluationModeEnum converts a PolicyClusterAdmissionRulesEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationPolicyClusterAdmissionRulesEvaluationModeEnum(e binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEvaluationModeEnum) *binaryauthorization.PolicyClusterAdmissionRulesEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEvaluationModeEnum_name[int32(e)]; ok {
		e := binaryauthorization.PolicyClusterAdmissionRulesEvaluationModeEnum(n[len("BinaryauthorizationPolicyClusterAdmissionRulesEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyClusterAdmissionRulesEnforcementModeEnum converts a PolicyClusterAdmissionRulesEnforcementModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationPolicyClusterAdmissionRulesEnforcementModeEnum(e binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEnforcementModeEnum) *binaryauthorization.PolicyClusterAdmissionRulesEnforcementModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEnforcementModeEnum_name[int32(e)]; ok {
		e := binaryauthorization.PolicyClusterAdmissionRulesEnforcementModeEnum(n[len("BinaryauthorizationPolicyClusterAdmissionRulesEnforcementModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum converts a PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(e binaryauthorizationpb.BinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum) *binaryauthorization.PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := binaryauthorizationpb.BinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum_name[int32(e)]; ok {
		e := binaryauthorization.PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(n[len("BinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum converts a PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(e binaryauthorizationpb.BinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum) *binaryauthorization.PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := binaryauthorizationpb.BinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum_name[int32(e)]; ok {
		e := binaryauthorization.PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(n[len("BinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum converts a PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(e binaryauthorizationpb.BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum) *binaryauthorization.PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := binaryauthorizationpb.BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum_name[int32(e)]; ok {
		e := binaryauthorization.PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(n[len("BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum converts a PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(e binaryauthorizationpb.BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum) *binaryauthorization.PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := binaryauthorizationpb.BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum_name[int32(e)]; ok {
		e := binaryauthorization.PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(n[len("BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum converts a PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(e binaryauthorizationpb.BinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum) *binaryauthorization.PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := binaryauthorizationpb.BinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum_name[int32(e)]; ok {
		e := binaryauthorization.PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(n[len("BinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum converts a PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(e binaryauthorizationpb.BinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum) *binaryauthorization.PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := binaryauthorizationpb.BinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum_name[int32(e)]; ok {
		e := binaryauthorization.PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(n[len("BinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyDefaultAdmissionRuleEvaluationModeEnum converts a PolicyDefaultAdmissionRuleEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationPolicyDefaultAdmissionRuleEvaluationModeEnum(e binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEvaluationModeEnum) *binaryauthorization.PolicyDefaultAdmissionRuleEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEvaluationModeEnum_name[int32(e)]; ok {
		e := binaryauthorization.PolicyDefaultAdmissionRuleEvaluationModeEnum(n[len("BinaryauthorizationPolicyDefaultAdmissionRuleEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyDefaultAdmissionRuleEnforcementModeEnum converts a PolicyDefaultAdmissionRuleEnforcementModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationPolicyDefaultAdmissionRuleEnforcementModeEnum(e binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEnforcementModeEnum) *binaryauthorization.PolicyDefaultAdmissionRuleEnforcementModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEnforcementModeEnum_name[int32(e)]; ok {
		e := binaryauthorization.PolicyDefaultAdmissionRuleEnforcementModeEnum(n[len("BinaryauthorizationPolicyDefaultAdmissionRuleEnforcementModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyGlobalPolicyEvaluationModeEnum converts a PolicyGlobalPolicyEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationPolicyGlobalPolicyEvaluationModeEnum(e binaryauthorizationpb.BinaryauthorizationPolicyGlobalPolicyEvaluationModeEnum) *binaryauthorization.PolicyGlobalPolicyEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := binaryauthorizationpb.BinaryauthorizationPolicyGlobalPolicyEvaluationModeEnum_name[int32(e)]; ok {
		e := binaryauthorization.PolicyGlobalPolicyEvaluationModeEnum(n[len("BinaryauthorizationPolicyGlobalPolicyEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyAdmissionWhitelistPatterns converts a PolicyAdmissionWhitelistPatterns object from its proto representation.
func ProtoToBinaryauthorizationPolicyAdmissionWhitelistPatterns(p *binaryauthorizationpb.BinaryauthorizationPolicyAdmissionWhitelistPatterns) *binaryauthorization.PolicyAdmissionWhitelistPatterns {
	if p == nil {
		return nil
	}
	obj := &binaryauthorization.PolicyAdmissionWhitelistPatterns{
		NamePattern: dcl.StringOrNil(p.GetNamePattern()),
	}
	return obj
}

// ProtoToPolicyClusterAdmissionRules converts a PolicyClusterAdmissionRules object from its proto representation.
func ProtoToBinaryauthorizationPolicyClusterAdmissionRules(p *binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRules) *binaryauthorization.PolicyClusterAdmissionRules {
	if p == nil {
		return nil
	}
	obj := &binaryauthorization.PolicyClusterAdmissionRules{
		EvaluationMode:  ProtoToBinaryauthorizationPolicyClusterAdmissionRulesEvaluationModeEnum(p.GetEvaluationMode()),
		EnforcementMode: ProtoToBinaryauthorizationPolicyClusterAdmissionRulesEnforcementModeEnum(p.GetEnforcementMode()),
	}
	for _, r := range p.GetRequireAttestationsBy() {
		obj.RequireAttestationsBy = append(obj.RequireAttestationsBy, r)
	}
	return obj
}

// ProtoToPolicyKubernetesNamespaceAdmissionRules converts a PolicyKubernetesNamespaceAdmissionRules object from its proto representation.
func ProtoToBinaryauthorizationPolicyKubernetesNamespaceAdmissionRules(p *binaryauthorizationpb.BinaryauthorizationPolicyKubernetesNamespaceAdmissionRules) *binaryauthorization.PolicyKubernetesNamespaceAdmissionRules {
	if p == nil {
		return nil
	}
	obj := &binaryauthorization.PolicyKubernetesNamespaceAdmissionRules{
		EvaluationMode:  ProtoToBinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(p.GetEvaluationMode()),
		EnforcementMode: ProtoToBinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(p.GetEnforcementMode()),
	}
	for _, r := range p.GetRequireAttestationsBy() {
		obj.RequireAttestationsBy = append(obj.RequireAttestationsBy, r)
	}
	return obj
}

// ProtoToPolicyKubernetesServiceAccountAdmissionRules converts a PolicyKubernetesServiceAccountAdmissionRules object from its proto representation.
func ProtoToBinaryauthorizationPolicyKubernetesServiceAccountAdmissionRules(p *binaryauthorizationpb.BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRules) *binaryauthorization.PolicyKubernetesServiceAccountAdmissionRules {
	if p == nil {
		return nil
	}
	obj := &binaryauthorization.PolicyKubernetesServiceAccountAdmissionRules{
		EvaluationMode:  ProtoToBinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(p.GetEvaluationMode()),
		EnforcementMode: ProtoToBinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(p.GetEnforcementMode()),
	}
	for _, r := range p.GetRequireAttestationsBy() {
		obj.RequireAttestationsBy = append(obj.RequireAttestationsBy, r)
	}
	return obj
}

// ProtoToPolicyIstioServiceIdentityAdmissionRules converts a PolicyIstioServiceIdentityAdmissionRules object from its proto representation.
func ProtoToBinaryauthorizationPolicyIstioServiceIdentityAdmissionRules(p *binaryauthorizationpb.BinaryauthorizationPolicyIstioServiceIdentityAdmissionRules) *binaryauthorization.PolicyIstioServiceIdentityAdmissionRules {
	if p == nil {
		return nil
	}
	obj := &binaryauthorization.PolicyIstioServiceIdentityAdmissionRules{
		EvaluationMode:  ProtoToBinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(p.GetEvaluationMode()),
		EnforcementMode: ProtoToBinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(p.GetEnforcementMode()),
	}
	for _, r := range p.GetRequireAttestationsBy() {
		obj.RequireAttestationsBy = append(obj.RequireAttestationsBy, r)
	}
	return obj
}

// ProtoToPolicyDefaultAdmissionRule converts a PolicyDefaultAdmissionRule object from its proto representation.
func ProtoToBinaryauthorizationPolicyDefaultAdmissionRule(p *binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRule) *binaryauthorization.PolicyDefaultAdmissionRule {
	if p == nil {
		return nil
	}
	obj := &binaryauthorization.PolicyDefaultAdmissionRule{
		EvaluationMode:  ProtoToBinaryauthorizationPolicyDefaultAdmissionRuleEvaluationModeEnum(p.GetEvaluationMode()),
		EnforcementMode: ProtoToBinaryauthorizationPolicyDefaultAdmissionRuleEnforcementModeEnum(p.GetEnforcementMode()),
	}
	for _, r := range p.GetRequireAttestationsBy() {
		obj.RequireAttestationsBy = append(obj.RequireAttestationsBy, r)
	}
	return obj
}

// ProtoToPolicy converts a Policy resource from its proto representation.
func ProtoToPolicy(p *binaryauthorizationpb.BinaryauthorizationPolicy) *binaryauthorization.Policy {
	obj := &binaryauthorization.Policy{
		DefaultAdmissionRule:       ProtoToBinaryauthorizationPolicyDefaultAdmissionRule(p.GetDefaultAdmissionRule()),
		Description:                dcl.StringOrNil(p.GetDescription()),
		GlobalPolicyEvaluationMode: ProtoToBinaryauthorizationPolicyGlobalPolicyEvaluationModeEnum(p.GetGlobalPolicyEvaluationMode()),
		SelfLink:                   dcl.StringOrNil(p.GetSelfLink()),
		Project:                    dcl.StringOrNil(p.GetProject()),
		UpdateTime:                 dcl.StringOrNil(p.GetUpdateTime()),
	}
	for _, r := range p.GetAdmissionWhitelistPatterns() {
		obj.AdmissionWhitelistPatterns = append(obj.AdmissionWhitelistPatterns, *ProtoToBinaryauthorizationPolicyAdmissionWhitelistPatterns(r))
	}
	return obj
}

// PolicyClusterAdmissionRulesEvaluationModeEnumToProto converts a PolicyClusterAdmissionRulesEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationPolicyClusterAdmissionRulesEvaluationModeEnumToProto(e *binaryauthorization.PolicyClusterAdmissionRulesEvaluationModeEnum) binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEvaluationModeEnum {
	if e == nil {
		return binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEvaluationModeEnum(0)
	}
	if v, ok := binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEvaluationModeEnum_value["PolicyClusterAdmissionRulesEvaluationModeEnum"+string(*e)]; ok {
		return binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEvaluationModeEnum(v)
	}
	return binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEvaluationModeEnum(0)
}

// PolicyClusterAdmissionRulesEnforcementModeEnumToProto converts a PolicyClusterAdmissionRulesEnforcementModeEnum enum to its proto representation.
func BinaryauthorizationPolicyClusterAdmissionRulesEnforcementModeEnumToProto(e *binaryauthorization.PolicyClusterAdmissionRulesEnforcementModeEnum) binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEnforcementModeEnum {
	if e == nil {
		return binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEnforcementModeEnum(0)
	}
	if v, ok := binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEnforcementModeEnum_value["PolicyClusterAdmissionRulesEnforcementModeEnum"+string(*e)]; ok {
		return binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEnforcementModeEnum(v)
	}
	return binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEnforcementModeEnum(0)
}

// PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnumToProto converts a PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnumToProto(e *binaryauthorization.PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum) binaryauthorizationpb.BinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum {
	if e == nil {
		return binaryauthorizationpb.BinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(0)
	}
	if v, ok := binaryauthorizationpb.BinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum_value["PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum"+string(*e)]; ok {
		return binaryauthorizationpb.BinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(v)
	}
	return binaryauthorizationpb.BinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(0)
}

// PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnumToProto converts a PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum enum to its proto representation.
func BinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnumToProto(e *binaryauthorization.PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum) binaryauthorizationpb.BinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum {
	if e == nil {
		return binaryauthorizationpb.BinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(0)
	}
	if v, ok := binaryauthorizationpb.BinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum_value["PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum"+string(*e)]; ok {
		return binaryauthorizationpb.BinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(v)
	}
	return binaryauthorizationpb.BinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(0)
}

// PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnumToProto converts a PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnumToProto(e *binaryauthorization.PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum) binaryauthorizationpb.BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum {
	if e == nil {
		return binaryauthorizationpb.BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(0)
	}
	if v, ok := binaryauthorizationpb.BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum_value["PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum"+string(*e)]; ok {
		return binaryauthorizationpb.BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(v)
	}
	return binaryauthorizationpb.BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(0)
}

// PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnumToProto converts a PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum enum to its proto representation.
func BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnumToProto(e *binaryauthorization.PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum) binaryauthorizationpb.BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum {
	if e == nil {
		return binaryauthorizationpb.BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(0)
	}
	if v, ok := binaryauthorizationpb.BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum_value["PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum"+string(*e)]; ok {
		return binaryauthorizationpb.BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(v)
	}
	return binaryauthorizationpb.BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(0)
}

// PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnumToProto converts a PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnumToProto(e *binaryauthorization.PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum) binaryauthorizationpb.BinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum {
	if e == nil {
		return binaryauthorizationpb.BinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(0)
	}
	if v, ok := binaryauthorizationpb.BinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum_value["PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum"+string(*e)]; ok {
		return binaryauthorizationpb.BinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(v)
	}
	return binaryauthorizationpb.BinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(0)
}

// PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnumToProto converts a PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum enum to its proto representation.
func BinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnumToProto(e *binaryauthorization.PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum) binaryauthorizationpb.BinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum {
	if e == nil {
		return binaryauthorizationpb.BinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(0)
	}
	if v, ok := binaryauthorizationpb.BinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum_value["PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum"+string(*e)]; ok {
		return binaryauthorizationpb.BinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(v)
	}
	return binaryauthorizationpb.BinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(0)
}

// PolicyDefaultAdmissionRuleEvaluationModeEnumToProto converts a PolicyDefaultAdmissionRuleEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationPolicyDefaultAdmissionRuleEvaluationModeEnumToProto(e *binaryauthorization.PolicyDefaultAdmissionRuleEvaluationModeEnum) binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEvaluationModeEnum {
	if e == nil {
		return binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEvaluationModeEnum(0)
	}
	if v, ok := binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEvaluationModeEnum_value["PolicyDefaultAdmissionRuleEvaluationModeEnum"+string(*e)]; ok {
		return binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEvaluationModeEnum(v)
	}
	return binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEvaluationModeEnum(0)
}

// PolicyDefaultAdmissionRuleEnforcementModeEnumToProto converts a PolicyDefaultAdmissionRuleEnforcementModeEnum enum to its proto representation.
func BinaryauthorizationPolicyDefaultAdmissionRuleEnforcementModeEnumToProto(e *binaryauthorization.PolicyDefaultAdmissionRuleEnforcementModeEnum) binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEnforcementModeEnum {
	if e == nil {
		return binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEnforcementModeEnum(0)
	}
	if v, ok := binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEnforcementModeEnum_value["PolicyDefaultAdmissionRuleEnforcementModeEnum"+string(*e)]; ok {
		return binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEnforcementModeEnum(v)
	}
	return binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEnforcementModeEnum(0)
}

// PolicyGlobalPolicyEvaluationModeEnumToProto converts a PolicyGlobalPolicyEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationPolicyGlobalPolicyEvaluationModeEnumToProto(e *binaryauthorization.PolicyGlobalPolicyEvaluationModeEnum) binaryauthorizationpb.BinaryauthorizationPolicyGlobalPolicyEvaluationModeEnum {
	if e == nil {
		return binaryauthorizationpb.BinaryauthorizationPolicyGlobalPolicyEvaluationModeEnum(0)
	}
	if v, ok := binaryauthorizationpb.BinaryauthorizationPolicyGlobalPolicyEvaluationModeEnum_value["PolicyGlobalPolicyEvaluationModeEnum"+string(*e)]; ok {
		return binaryauthorizationpb.BinaryauthorizationPolicyGlobalPolicyEvaluationModeEnum(v)
	}
	return binaryauthorizationpb.BinaryauthorizationPolicyGlobalPolicyEvaluationModeEnum(0)
}

// PolicyAdmissionWhitelistPatternsToProto converts a PolicyAdmissionWhitelistPatterns object to its proto representation.
func BinaryauthorizationPolicyAdmissionWhitelistPatternsToProto(o *binaryauthorization.PolicyAdmissionWhitelistPatterns) *binaryauthorizationpb.BinaryauthorizationPolicyAdmissionWhitelistPatterns {
	if o == nil {
		return nil
	}
	p := &binaryauthorizationpb.BinaryauthorizationPolicyAdmissionWhitelistPatterns{}
	p.SetNamePattern(dcl.ValueOrEmptyString(o.NamePattern))
	return p
}

// PolicyClusterAdmissionRulesToProto converts a PolicyClusterAdmissionRules object to its proto representation.
func BinaryauthorizationPolicyClusterAdmissionRulesToProto(o *binaryauthorization.PolicyClusterAdmissionRules) *binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRules {
	if o == nil {
		return nil
	}
	p := &binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRules{}
	p.SetEvaluationMode(BinaryauthorizationPolicyClusterAdmissionRulesEvaluationModeEnumToProto(o.EvaluationMode))
	p.SetEnforcementMode(BinaryauthorizationPolicyClusterAdmissionRulesEnforcementModeEnumToProto(o.EnforcementMode))
	sRequireAttestationsBy := make([]string, len(o.RequireAttestationsBy))
	for i, r := range o.RequireAttestationsBy {
		sRequireAttestationsBy[i] = r
	}
	p.SetRequireAttestationsBy(sRequireAttestationsBy)
	return p
}

// PolicyKubernetesNamespaceAdmissionRulesToProto converts a PolicyKubernetesNamespaceAdmissionRules object to its proto representation.
func BinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesToProto(o *binaryauthorization.PolicyKubernetesNamespaceAdmissionRules) *binaryauthorizationpb.BinaryauthorizationPolicyKubernetesNamespaceAdmissionRules {
	if o == nil {
		return nil
	}
	p := &binaryauthorizationpb.BinaryauthorizationPolicyKubernetesNamespaceAdmissionRules{}
	p.SetEvaluationMode(BinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnumToProto(o.EvaluationMode))
	p.SetEnforcementMode(BinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnumToProto(o.EnforcementMode))
	sRequireAttestationsBy := make([]string, len(o.RequireAttestationsBy))
	for i, r := range o.RequireAttestationsBy {
		sRequireAttestationsBy[i] = r
	}
	p.SetRequireAttestationsBy(sRequireAttestationsBy)
	return p
}

// PolicyKubernetesServiceAccountAdmissionRulesToProto converts a PolicyKubernetesServiceAccountAdmissionRules object to its proto representation.
func BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesToProto(o *binaryauthorization.PolicyKubernetesServiceAccountAdmissionRules) *binaryauthorizationpb.BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRules {
	if o == nil {
		return nil
	}
	p := &binaryauthorizationpb.BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRules{}
	p.SetEvaluationMode(BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnumToProto(o.EvaluationMode))
	p.SetEnforcementMode(BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnumToProto(o.EnforcementMode))
	sRequireAttestationsBy := make([]string, len(o.RequireAttestationsBy))
	for i, r := range o.RequireAttestationsBy {
		sRequireAttestationsBy[i] = r
	}
	p.SetRequireAttestationsBy(sRequireAttestationsBy)
	return p
}

// PolicyIstioServiceIdentityAdmissionRulesToProto converts a PolicyIstioServiceIdentityAdmissionRules object to its proto representation.
func BinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesToProto(o *binaryauthorization.PolicyIstioServiceIdentityAdmissionRules) *binaryauthorizationpb.BinaryauthorizationPolicyIstioServiceIdentityAdmissionRules {
	if o == nil {
		return nil
	}
	p := &binaryauthorizationpb.BinaryauthorizationPolicyIstioServiceIdentityAdmissionRules{}
	p.SetEvaluationMode(BinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnumToProto(o.EvaluationMode))
	p.SetEnforcementMode(BinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnumToProto(o.EnforcementMode))
	sRequireAttestationsBy := make([]string, len(o.RequireAttestationsBy))
	for i, r := range o.RequireAttestationsBy {
		sRequireAttestationsBy[i] = r
	}
	p.SetRequireAttestationsBy(sRequireAttestationsBy)
	return p
}

// PolicyDefaultAdmissionRuleToProto converts a PolicyDefaultAdmissionRule object to its proto representation.
func BinaryauthorizationPolicyDefaultAdmissionRuleToProto(o *binaryauthorization.PolicyDefaultAdmissionRule) *binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRule {
	if o == nil {
		return nil
	}
	p := &binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRule{}
	p.SetEvaluationMode(BinaryauthorizationPolicyDefaultAdmissionRuleEvaluationModeEnumToProto(o.EvaluationMode))
	p.SetEnforcementMode(BinaryauthorizationPolicyDefaultAdmissionRuleEnforcementModeEnumToProto(o.EnforcementMode))
	sRequireAttestationsBy := make([]string, len(o.RequireAttestationsBy))
	for i, r := range o.RequireAttestationsBy {
		sRequireAttestationsBy[i] = r
	}
	p.SetRequireAttestationsBy(sRequireAttestationsBy)
	return p
}

// PolicyToProto converts a Policy resource to its proto representation.
func PolicyToProto(resource *binaryauthorization.Policy) *binaryauthorizationpb.BinaryauthorizationPolicy {
	p := &binaryauthorizationpb.BinaryauthorizationPolicy{}
	p.SetDefaultAdmissionRule(BinaryauthorizationPolicyDefaultAdmissionRuleToProto(resource.DefaultAdmissionRule))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetGlobalPolicyEvaluationMode(BinaryauthorizationPolicyGlobalPolicyEvaluationModeEnumToProto(resource.GlobalPolicyEvaluationMode))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	sAdmissionWhitelistPatterns := make([]*binaryauthorizationpb.BinaryauthorizationPolicyAdmissionWhitelistPatterns, len(resource.AdmissionWhitelistPatterns))
	for i, r := range resource.AdmissionWhitelistPatterns {
		sAdmissionWhitelistPatterns[i] = BinaryauthorizationPolicyAdmissionWhitelistPatternsToProto(&r)
	}
	p.SetAdmissionWhitelistPatterns(sAdmissionWhitelistPatterns)
	mClusterAdmissionRules := make(map[string]*binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRules, len(resource.ClusterAdmissionRules))
	for k, r := range resource.ClusterAdmissionRules {
		mClusterAdmissionRules[k] = BinaryauthorizationPolicyClusterAdmissionRulesToProto(&r)
	}
	p.SetClusterAdmissionRules(mClusterAdmissionRules)
	mKubernetesNamespaceAdmissionRules := make(map[string]*binaryauthorizationpb.BinaryauthorizationPolicyKubernetesNamespaceAdmissionRules, len(resource.KubernetesNamespaceAdmissionRules))
	for k, r := range resource.KubernetesNamespaceAdmissionRules {
		mKubernetesNamespaceAdmissionRules[k] = BinaryauthorizationPolicyKubernetesNamespaceAdmissionRulesToProto(&r)
	}
	p.SetKubernetesNamespaceAdmissionRules(mKubernetesNamespaceAdmissionRules)
	mKubernetesServiceAccountAdmissionRules := make(map[string]*binaryauthorizationpb.BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRules, len(resource.KubernetesServiceAccountAdmissionRules))
	for k, r := range resource.KubernetesServiceAccountAdmissionRules {
		mKubernetesServiceAccountAdmissionRules[k] = BinaryauthorizationPolicyKubernetesServiceAccountAdmissionRulesToProto(&r)
	}
	p.SetKubernetesServiceAccountAdmissionRules(mKubernetesServiceAccountAdmissionRules)
	mIstioServiceIdentityAdmissionRules := make(map[string]*binaryauthorizationpb.BinaryauthorizationPolicyIstioServiceIdentityAdmissionRules, len(resource.IstioServiceIdentityAdmissionRules))
	for k, r := range resource.IstioServiceIdentityAdmissionRules {
		mIstioServiceIdentityAdmissionRules[k] = BinaryauthorizationPolicyIstioServiceIdentityAdmissionRulesToProto(&r)
	}
	p.SetIstioServiceIdentityAdmissionRules(mIstioServiceIdentityAdmissionRules)

	return p
}

// applyPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) applyPolicy(ctx context.Context, c *binaryauthorization.Client, request *binaryauthorizationpb.ApplyBinaryauthorizationPolicyRequest) (*binaryauthorizationpb.BinaryauthorizationPolicy, error) {
	p := ProtoToPolicy(request.GetResource())
	res, err := c.ApplyPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := PolicyToProto(res)
	return r, nil
}

// applyBinaryauthorizationPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) ApplyBinaryauthorizationPolicy(ctx context.Context, request *binaryauthorizationpb.ApplyBinaryauthorizationPolicyRequest) (*binaryauthorizationpb.BinaryauthorizationPolicy, error) {
	cl, err := createConfigPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyPolicy(ctx, cl, request)
}

// DeletePolicy handles the gRPC request by passing it to the underlying Policy Delete() method.
func (s *PolicyServer) DeleteBinaryauthorizationPolicy(ctx context.Context, request *binaryauthorizationpb.DeleteBinaryauthorizationPolicyRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for Policy")

}

// ListBinaryauthorizationPolicy is a no-op method because Policy has no list method.
func (s *PolicyServer) ListBinaryauthorizationPolicy(_ context.Context, _ *binaryauthorizationpb.ListBinaryauthorizationPolicyRequest) (*binaryauthorizationpb.ListBinaryauthorizationPolicyResponse, error) {
	return nil, nil
}

func createConfigPolicy(ctx context.Context, service_account_file string) (*binaryauthorization.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return binaryauthorization.NewClient(conf), nil
}
