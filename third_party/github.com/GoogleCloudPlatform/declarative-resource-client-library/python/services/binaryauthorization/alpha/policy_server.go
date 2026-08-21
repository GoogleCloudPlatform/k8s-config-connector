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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/binaryauthorization/alpha/binaryauthorization_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization/alpha"
)

// PolicyServer implements the gRPC interface for Policy.
type PolicyServer struct{}

// ProtoToPolicyClusterAdmissionRulesEvaluationModeEnum converts a PolicyClusterAdmissionRulesEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationAlphaPolicyClusterAdmissionRulesEvaluationModeEnum(e alphapb.BinaryauthorizationAlphaPolicyClusterAdmissionRulesEvaluationModeEnum) *alpha.PolicyClusterAdmissionRulesEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BinaryauthorizationAlphaPolicyClusterAdmissionRulesEvaluationModeEnum_name[int32(e)]; ok {
		e := alpha.PolicyClusterAdmissionRulesEvaluationModeEnum(n[len("BinaryauthorizationAlphaPolicyClusterAdmissionRulesEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyClusterAdmissionRulesEnforcementModeEnum converts a PolicyClusterAdmissionRulesEnforcementModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationAlphaPolicyClusterAdmissionRulesEnforcementModeEnum(e alphapb.BinaryauthorizationAlphaPolicyClusterAdmissionRulesEnforcementModeEnum) *alpha.PolicyClusterAdmissionRulesEnforcementModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BinaryauthorizationAlphaPolicyClusterAdmissionRulesEnforcementModeEnum_name[int32(e)]; ok {
		e := alpha.PolicyClusterAdmissionRulesEnforcementModeEnum(n[len("BinaryauthorizationAlphaPolicyClusterAdmissionRulesEnforcementModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum converts a PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(e alphapb.BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum) *alpha.PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum_name[int32(e)]; ok {
		e := alpha.PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(n[len("BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum converts a PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(e alphapb.BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum) *alpha.PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum_name[int32(e)]; ok {
		e := alpha.PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(n[len("BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum converts a PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(e alphapb.BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum) *alpha.PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum_name[int32(e)]; ok {
		e := alpha.PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(n[len("BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum converts a PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(e alphapb.BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum) *alpha.PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum_name[int32(e)]; ok {
		e := alpha.PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(n[len("BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum converts a PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(e alphapb.BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum) *alpha.PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum_name[int32(e)]; ok {
		e := alpha.PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(n[len("BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum converts a PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(e alphapb.BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum) *alpha.PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum_name[int32(e)]; ok {
		e := alpha.PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(n[len("BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyDefaultAdmissionRuleEvaluationModeEnum converts a PolicyDefaultAdmissionRuleEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationAlphaPolicyDefaultAdmissionRuleEvaluationModeEnum(e alphapb.BinaryauthorizationAlphaPolicyDefaultAdmissionRuleEvaluationModeEnum) *alpha.PolicyDefaultAdmissionRuleEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BinaryauthorizationAlphaPolicyDefaultAdmissionRuleEvaluationModeEnum_name[int32(e)]; ok {
		e := alpha.PolicyDefaultAdmissionRuleEvaluationModeEnum(n[len("BinaryauthorizationAlphaPolicyDefaultAdmissionRuleEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyDefaultAdmissionRuleEnforcementModeEnum converts a PolicyDefaultAdmissionRuleEnforcementModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationAlphaPolicyDefaultAdmissionRuleEnforcementModeEnum(e alphapb.BinaryauthorizationAlphaPolicyDefaultAdmissionRuleEnforcementModeEnum) *alpha.PolicyDefaultAdmissionRuleEnforcementModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BinaryauthorizationAlphaPolicyDefaultAdmissionRuleEnforcementModeEnum_name[int32(e)]; ok {
		e := alpha.PolicyDefaultAdmissionRuleEnforcementModeEnum(n[len("BinaryauthorizationAlphaPolicyDefaultAdmissionRuleEnforcementModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyGlobalPolicyEvaluationModeEnum converts a PolicyGlobalPolicyEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnum(e alphapb.BinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnum) *alpha.PolicyGlobalPolicyEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnum_name[int32(e)]; ok {
		e := alpha.PolicyGlobalPolicyEvaluationModeEnum(n[len("BinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyAdmissionWhitelistPatterns converts a PolicyAdmissionWhitelistPatterns object from its proto representation.
func ProtoToBinaryauthorizationAlphaPolicyAdmissionWhitelistPatterns(p *alphapb.BinaryauthorizationAlphaPolicyAdmissionWhitelistPatterns) *alpha.PolicyAdmissionWhitelistPatterns {
	if p == nil {
		return nil
	}
	obj := &alpha.PolicyAdmissionWhitelistPatterns{
		NamePattern: dcl.StringOrNil(p.GetNamePattern()),
	}
	return obj
}

// ProtoToPolicyClusterAdmissionRules converts a PolicyClusterAdmissionRules object from its proto representation.
func ProtoToBinaryauthorizationAlphaPolicyClusterAdmissionRules(p *alphapb.BinaryauthorizationAlphaPolicyClusterAdmissionRules) *alpha.PolicyClusterAdmissionRules {
	if p == nil {
		return nil
	}
	obj := &alpha.PolicyClusterAdmissionRules{
		EvaluationMode:  ProtoToBinaryauthorizationAlphaPolicyClusterAdmissionRulesEvaluationModeEnum(p.GetEvaluationMode()),
		EnforcementMode: ProtoToBinaryauthorizationAlphaPolicyClusterAdmissionRulesEnforcementModeEnum(p.GetEnforcementMode()),
	}
	for _, r := range p.GetRequireAttestationsBy() {
		obj.RequireAttestationsBy = append(obj.RequireAttestationsBy, r)
	}
	return obj
}

// ProtoToPolicyKubernetesNamespaceAdmissionRules converts a PolicyKubernetesNamespaceAdmissionRules object from its proto representation.
func ProtoToBinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRules(p *alphapb.BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRules) *alpha.PolicyKubernetesNamespaceAdmissionRules {
	if p == nil {
		return nil
	}
	obj := &alpha.PolicyKubernetesNamespaceAdmissionRules{
		EvaluationMode:  ProtoToBinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(p.GetEvaluationMode()),
		EnforcementMode: ProtoToBinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(p.GetEnforcementMode()),
	}
	for _, r := range p.GetRequireAttestationsBy() {
		obj.RequireAttestationsBy = append(obj.RequireAttestationsBy, r)
	}
	return obj
}

// ProtoToPolicyKubernetesServiceAccountAdmissionRules converts a PolicyKubernetesServiceAccountAdmissionRules object from its proto representation.
func ProtoToBinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRules(p *alphapb.BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRules) *alpha.PolicyKubernetesServiceAccountAdmissionRules {
	if p == nil {
		return nil
	}
	obj := &alpha.PolicyKubernetesServiceAccountAdmissionRules{
		EvaluationMode:  ProtoToBinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(p.GetEvaluationMode()),
		EnforcementMode: ProtoToBinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(p.GetEnforcementMode()),
	}
	for _, r := range p.GetRequireAttestationsBy() {
		obj.RequireAttestationsBy = append(obj.RequireAttestationsBy, r)
	}
	return obj
}

// ProtoToPolicyIstioServiceIdentityAdmissionRules converts a PolicyIstioServiceIdentityAdmissionRules object from its proto representation.
func ProtoToBinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRules(p *alphapb.BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRules) *alpha.PolicyIstioServiceIdentityAdmissionRules {
	if p == nil {
		return nil
	}
	obj := &alpha.PolicyIstioServiceIdentityAdmissionRules{
		EvaluationMode:  ProtoToBinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(p.GetEvaluationMode()),
		EnforcementMode: ProtoToBinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(p.GetEnforcementMode()),
	}
	for _, r := range p.GetRequireAttestationsBy() {
		obj.RequireAttestationsBy = append(obj.RequireAttestationsBy, r)
	}
	return obj
}

// ProtoToPolicyDefaultAdmissionRule converts a PolicyDefaultAdmissionRule object from its proto representation.
func ProtoToBinaryauthorizationAlphaPolicyDefaultAdmissionRule(p *alphapb.BinaryauthorizationAlphaPolicyDefaultAdmissionRule) *alpha.PolicyDefaultAdmissionRule {
	if p == nil {
		return nil
	}
	obj := &alpha.PolicyDefaultAdmissionRule{
		EvaluationMode:  ProtoToBinaryauthorizationAlphaPolicyDefaultAdmissionRuleEvaluationModeEnum(p.GetEvaluationMode()),
		EnforcementMode: ProtoToBinaryauthorizationAlphaPolicyDefaultAdmissionRuleEnforcementModeEnum(p.GetEnforcementMode()),
	}
	for _, r := range p.GetRequireAttestationsBy() {
		obj.RequireAttestationsBy = append(obj.RequireAttestationsBy, r)
	}
	return obj
}

// ProtoToPolicy converts a Policy resource from its proto representation.
func ProtoToPolicy(p *alphapb.BinaryauthorizationAlphaPolicy) *alpha.Policy {
	obj := &alpha.Policy{
		DefaultAdmissionRule:       ProtoToBinaryauthorizationAlphaPolicyDefaultAdmissionRule(p.GetDefaultAdmissionRule()),
		Description:                dcl.StringOrNil(p.GetDescription()),
		GlobalPolicyEvaluationMode: ProtoToBinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnum(p.GetGlobalPolicyEvaluationMode()),
		SelfLink:                   dcl.StringOrNil(p.GetSelfLink()),
		Project:                    dcl.StringOrNil(p.GetProject()),
		UpdateTime:                 dcl.StringOrNil(p.GetUpdateTime()),
	}
	for _, r := range p.GetAdmissionWhitelistPatterns() {
		obj.AdmissionWhitelistPatterns = append(obj.AdmissionWhitelistPatterns, *ProtoToBinaryauthorizationAlphaPolicyAdmissionWhitelistPatterns(r))
	}
	return obj
}

// PolicyClusterAdmissionRulesEvaluationModeEnumToProto converts a PolicyClusterAdmissionRulesEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationAlphaPolicyClusterAdmissionRulesEvaluationModeEnumToProto(e *alpha.PolicyClusterAdmissionRulesEvaluationModeEnum) alphapb.BinaryauthorizationAlphaPolicyClusterAdmissionRulesEvaluationModeEnum {
	if e == nil {
		return alphapb.BinaryauthorizationAlphaPolicyClusterAdmissionRulesEvaluationModeEnum(0)
	}
	if v, ok := alphapb.BinaryauthorizationAlphaPolicyClusterAdmissionRulesEvaluationModeEnum_value["PolicyClusterAdmissionRulesEvaluationModeEnum"+string(*e)]; ok {
		return alphapb.BinaryauthorizationAlphaPolicyClusterAdmissionRulesEvaluationModeEnum(v)
	}
	return alphapb.BinaryauthorizationAlphaPolicyClusterAdmissionRulesEvaluationModeEnum(0)
}

// PolicyClusterAdmissionRulesEnforcementModeEnumToProto converts a PolicyClusterAdmissionRulesEnforcementModeEnum enum to its proto representation.
func BinaryauthorizationAlphaPolicyClusterAdmissionRulesEnforcementModeEnumToProto(e *alpha.PolicyClusterAdmissionRulesEnforcementModeEnum) alphapb.BinaryauthorizationAlphaPolicyClusterAdmissionRulesEnforcementModeEnum {
	if e == nil {
		return alphapb.BinaryauthorizationAlphaPolicyClusterAdmissionRulesEnforcementModeEnum(0)
	}
	if v, ok := alphapb.BinaryauthorizationAlphaPolicyClusterAdmissionRulesEnforcementModeEnum_value["PolicyClusterAdmissionRulesEnforcementModeEnum"+string(*e)]; ok {
		return alphapb.BinaryauthorizationAlphaPolicyClusterAdmissionRulesEnforcementModeEnum(v)
	}
	return alphapb.BinaryauthorizationAlphaPolicyClusterAdmissionRulesEnforcementModeEnum(0)
}

// PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnumToProto converts a PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnumToProto(e *alpha.PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum) alphapb.BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum {
	if e == nil {
		return alphapb.BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(0)
	}
	if v, ok := alphapb.BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum_value["PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum"+string(*e)]; ok {
		return alphapb.BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(v)
	}
	return alphapb.BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(0)
}

// PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnumToProto converts a PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum enum to its proto representation.
func BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnumToProto(e *alpha.PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum) alphapb.BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum {
	if e == nil {
		return alphapb.BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(0)
	}
	if v, ok := alphapb.BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum_value["PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum"+string(*e)]; ok {
		return alphapb.BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(v)
	}
	return alphapb.BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(0)
}

// PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnumToProto converts a PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnumToProto(e *alpha.PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum) alphapb.BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum {
	if e == nil {
		return alphapb.BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(0)
	}
	if v, ok := alphapb.BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum_value["PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum"+string(*e)]; ok {
		return alphapb.BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(v)
	}
	return alphapb.BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(0)
}

// PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnumToProto converts a PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum enum to its proto representation.
func BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnumToProto(e *alpha.PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum) alphapb.BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum {
	if e == nil {
		return alphapb.BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(0)
	}
	if v, ok := alphapb.BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum_value["PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum"+string(*e)]; ok {
		return alphapb.BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(v)
	}
	return alphapb.BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(0)
}

// PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnumToProto converts a PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnumToProto(e *alpha.PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum) alphapb.BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum {
	if e == nil {
		return alphapb.BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(0)
	}
	if v, ok := alphapb.BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum_value["PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum"+string(*e)]; ok {
		return alphapb.BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(v)
	}
	return alphapb.BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(0)
}

// PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnumToProto converts a PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum enum to its proto representation.
func BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnumToProto(e *alpha.PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum) alphapb.BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum {
	if e == nil {
		return alphapb.BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(0)
	}
	if v, ok := alphapb.BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum_value["PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum"+string(*e)]; ok {
		return alphapb.BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(v)
	}
	return alphapb.BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(0)
}

// PolicyDefaultAdmissionRuleEvaluationModeEnumToProto converts a PolicyDefaultAdmissionRuleEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationAlphaPolicyDefaultAdmissionRuleEvaluationModeEnumToProto(e *alpha.PolicyDefaultAdmissionRuleEvaluationModeEnum) alphapb.BinaryauthorizationAlphaPolicyDefaultAdmissionRuleEvaluationModeEnum {
	if e == nil {
		return alphapb.BinaryauthorizationAlphaPolicyDefaultAdmissionRuleEvaluationModeEnum(0)
	}
	if v, ok := alphapb.BinaryauthorizationAlphaPolicyDefaultAdmissionRuleEvaluationModeEnum_value["PolicyDefaultAdmissionRuleEvaluationModeEnum"+string(*e)]; ok {
		return alphapb.BinaryauthorizationAlphaPolicyDefaultAdmissionRuleEvaluationModeEnum(v)
	}
	return alphapb.BinaryauthorizationAlphaPolicyDefaultAdmissionRuleEvaluationModeEnum(0)
}

// PolicyDefaultAdmissionRuleEnforcementModeEnumToProto converts a PolicyDefaultAdmissionRuleEnforcementModeEnum enum to its proto representation.
func BinaryauthorizationAlphaPolicyDefaultAdmissionRuleEnforcementModeEnumToProto(e *alpha.PolicyDefaultAdmissionRuleEnforcementModeEnum) alphapb.BinaryauthorizationAlphaPolicyDefaultAdmissionRuleEnforcementModeEnum {
	if e == nil {
		return alphapb.BinaryauthorizationAlphaPolicyDefaultAdmissionRuleEnforcementModeEnum(0)
	}
	if v, ok := alphapb.BinaryauthorizationAlphaPolicyDefaultAdmissionRuleEnforcementModeEnum_value["PolicyDefaultAdmissionRuleEnforcementModeEnum"+string(*e)]; ok {
		return alphapb.BinaryauthorizationAlphaPolicyDefaultAdmissionRuleEnforcementModeEnum(v)
	}
	return alphapb.BinaryauthorizationAlphaPolicyDefaultAdmissionRuleEnforcementModeEnum(0)
}

// PolicyGlobalPolicyEvaluationModeEnumToProto converts a PolicyGlobalPolicyEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnumToProto(e *alpha.PolicyGlobalPolicyEvaluationModeEnum) alphapb.BinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnum {
	if e == nil {
		return alphapb.BinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnum(0)
	}
	if v, ok := alphapb.BinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnum_value["PolicyGlobalPolicyEvaluationModeEnum"+string(*e)]; ok {
		return alphapb.BinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnum(v)
	}
	return alphapb.BinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnum(0)
}

// PolicyAdmissionWhitelistPatternsToProto converts a PolicyAdmissionWhitelistPatterns object to its proto representation.
func BinaryauthorizationAlphaPolicyAdmissionWhitelistPatternsToProto(o *alpha.PolicyAdmissionWhitelistPatterns) *alphapb.BinaryauthorizationAlphaPolicyAdmissionWhitelistPatterns {
	if o == nil {
		return nil
	}
	p := &alphapb.BinaryauthorizationAlphaPolicyAdmissionWhitelistPatterns{}
	p.SetNamePattern(dcl.ValueOrEmptyString(o.NamePattern))
	return p
}

// PolicyClusterAdmissionRulesToProto converts a PolicyClusterAdmissionRules object to its proto representation.
func BinaryauthorizationAlphaPolicyClusterAdmissionRulesToProto(o *alpha.PolicyClusterAdmissionRules) *alphapb.BinaryauthorizationAlphaPolicyClusterAdmissionRules {
	if o == nil {
		return nil
	}
	p := &alphapb.BinaryauthorizationAlphaPolicyClusterAdmissionRules{}
	p.SetEvaluationMode(BinaryauthorizationAlphaPolicyClusterAdmissionRulesEvaluationModeEnumToProto(o.EvaluationMode))
	p.SetEnforcementMode(BinaryauthorizationAlphaPolicyClusterAdmissionRulesEnforcementModeEnumToProto(o.EnforcementMode))
	sRequireAttestationsBy := make([]string, len(o.RequireAttestationsBy))
	for i, r := range o.RequireAttestationsBy {
		sRequireAttestationsBy[i] = r
	}
	p.SetRequireAttestationsBy(sRequireAttestationsBy)
	return p
}

// PolicyKubernetesNamespaceAdmissionRulesToProto converts a PolicyKubernetesNamespaceAdmissionRules object to its proto representation.
func BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesToProto(o *alpha.PolicyKubernetesNamespaceAdmissionRules) *alphapb.BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRules {
	if o == nil {
		return nil
	}
	p := &alphapb.BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRules{}
	p.SetEvaluationMode(BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnumToProto(o.EvaluationMode))
	p.SetEnforcementMode(BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnumToProto(o.EnforcementMode))
	sRequireAttestationsBy := make([]string, len(o.RequireAttestationsBy))
	for i, r := range o.RequireAttestationsBy {
		sRequireAttestationsBy[i] = r
	}
	p.SetRequireAttestationsBy(sRequireAttestationsBy)
	return p
}

// PolicyKubernetesServiceAccountAdmissionRulesToProto converts a PolicyKubernetesServiceAccountAdmissionRules object to its proto representation.
func BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesToProto(o *alpha.PolicyKubernetesServiceAccountAdmissionRules) *alphapb.BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRules {
	if o == nil {
		return nil
	}
	p := &alphapb.BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRules{}
	p.SetEvaluationMode(BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnumToProto(o.EvaluationMode))
	p.SetEnforcementMode(BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnumToProto(o.EnforcementMode))
	sRequireAttestationsBy := make([]string, len(o.RequireAttestationsBy))
	for i, r := range o.RequireAttestationsBy {
		sRequireAttestationsBy[i] = r
	}
	p.SetRequireAttestationsBy(sRequireAttestationsBy)
	return p
}

// PolicyIstioServiceIdentityAdmissionRulesToProto converts a PolicyIstioServiceIdentityAdmissionRules object to its proto representation.
func BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesToProto(o *alpha.PolicyIstioServiceIdentityAdmissionRules) *alphapb.BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRules {
	if o == nil {
		return nil
	}
	p := &alphapb.BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRules{}
	p.SetEvaluationMode(BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnumToProto(o.EvaluationMode))
	p.SetEnforcementMode(BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnumToProto(o.EnforcementMode))
	sRequireAttestationsBy := make([]string, len(o.RequireAttestationsBy))
	for i, r := range o.RequireAttestationsBy {
		sRequireAttestationsBy[i] = r
	}
	p.SetRequireAttestationsBy(sRequireAttestationsBy)
	return p
}

// PolicyDefaultAdmissionRuleToProto converts a PolicyDefaultAdmissionRule object to its proto representation.
func BinaryauthorizationAlphaPolicyDefaultAdmissionRuleToProto(o *alpha.PolicyDefaultAdmissionRule) *alphapb.BinaryauthorizationAlphaPolicyDefaultAdmissionRule {
	if o == nil {
		return nil
	}
	p := &alphapb.BinaryauthorizationAlphaPolicyDefaultAdmissionRule{}
	p.SetEvaluationMode(BinaryauthorizationAlphaPolicyDefaultAdmissionRuleEvaluationModeEnumToProto(o.EvaluationMode))
	p.SetEnforcementMode(BinaryauthorizationAlphaPolicyDefaultAdmissionRuleEnforcementModeEnumToProto(o.EnforcementMode))
	sRequireAttestationsBy := make([]string, len(o.RequireAttestationsBy))
	for i, r := range o.RequireAttestationsBy {
		sRequireAttestationsBy[i] = r
	}
	p.SetRequireAttestationsBy(sRequireAttestationsBy)
	return p
}

// PolicyToProto converts a Policy resource to its proto representation.
func PolicyToProto(resource *alpha.Policy) *alphapb.BinaryauthorizationAlphaPolicy {
	p := &alphapb.BinaryauthorizationAlphaPolicy{}
	p.SetDefaultAdmissionRule(BinaryauthorizationAlphaPolicyDefaultAdmissionRuleToProto(resource.DefaultAdmissionRule))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetGlobalPolicyEvaluationMode(BinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnumToProto(resource.GlobalPolicyEvaluationMode))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	sAdmissionWhitelistPatterns := make([]*alphapb.BinaryauthorizationAlphaPolicyAdmissionWhitelistPatterns, len(resource.AdmissionWhitelistPatterns))
	for i, r := range resource.AdmissionWhitelistPatterns {
		sAdmissionWhitelistPatterns[i] = BinaryauthorizationAlphaPolicyAdmissionWhitelistPatternsToProto(&r)
	}
	p.SetAdmissionWhitelistPatterns(sAdmissionWhitelistPatterns)
	mClusterAdmissionRules := make(map[string]*alphapb.BinaryauthorizationAlphaPolicyClusterAdmissionRules, len(resource.ClusterAdmissionRules))
	for k, r := range resource.ClusterAdmissionRules {
		mClusterAdmissionRules[k] = BinaryauthorizationAlphaPolicyClusterAdmissionRulesToProto(&r)
	}
	p.SetClusterAdmissionRules(mClusterAdmissionRules)
	mKubernetesNamespaceAdmissionRules := make(map[string]*alphapb.BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRules, len(resource.KubernetesNamespaceAdmissionRules))
	for k, r := range resource.KubernetesNamespaceAdmissionRules {
		mKubernetesNamespaceAdmissionRules[k] = BinaryauthorizationAlphaPolicyKubernetesNamespaceAdmissionRulesToProto(&r)
	}
	p.SetKubernetesNamespaceAdmissionRules(mKubernetesNamespaceAdmissionRules)
	mKubernetesServiceAccountAdmissionRules := make(map[string]*alphapb.BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRules, len(resource.KubernetesServiceAccountAdmissionRules))
	for k, r := range resource.KubernetesServiceAccountAdmissionRules {
		mKubernetesServiceAccountAdmissionRules[k] = BinaryauthorizationAlphaPolicyKubernetesServiceAccountAdmissionRulesToProto(&r)
	}
	p.SetKubernetesServiceAccountAdmissionRules(mKubernetesServiceAccountAdmissionRules)
	mIstioServiceIdentityAdmissionRules := make(map[string]*alphapb.BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRules, len(resource.IstioServiceIdentityAdmissionRules))
	for k, r := range resource.IstioServiceIdentityAdmissionRules {
		mIstioServiceIdentityAdmissionRules[k] = BinaryauthorizationAlphaPolicyIstioServiceIdentityAdmissionRulesToProto(&r)
	}
	p.SetIstioServiceIdentityAdmissionRules(mIstioServiceIdentityAdmissionRules)

	return p
}

// applyPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) applyPolicy(ctx context.Context, c *alpha.Client, request *alphapb.ApplyBinaryauthorizationAlphaPolicyRequest) (*alphapb.BinaryauthorizationAlphaPolicy, error) {
	p := ProtoToPolicy(request.GetResource())
	res, err := c.ApplyPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := PolicyToProto(res)
	return r, nil
}

// applyBinaryauthorizationAlphaPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) ApplyBinaryauthorizationAlphaPolicy(ctx context.Context, request *alphapb.ApplyBinaryauthorizationAlphaPolicyRequest) (*alphapb.BinaryauthorizationAlphaPolicy, error) {
	cl, err := createConfigPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyPolicy(ctx, cl, request)
}

// DeletePolicy handles the gRPC request by passing it to the underlying Policy Delete() method.
func (s *PolicyServer) DeleteBinaryauthorizationAlphaPolicy(ctx context.Context, request *alphapb.DeleteBinaryauthorizationAlphaPolicyRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for Policy")

}

// ListBinaryauthorizationAlphaPolicy is a no-op method because Policy has no list method.
func (s *PolicyServer) ListBinaryauthorizationAlphaPolicy(_ context.Context, _ *alphapb.ListBinaryauthorizationAlphaPolicyRequest) (*alphapb.ListBinaryauthorizationAlphaPolicyResponse, error) {
	return nil, nil
}

func createConfigPolicy(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
