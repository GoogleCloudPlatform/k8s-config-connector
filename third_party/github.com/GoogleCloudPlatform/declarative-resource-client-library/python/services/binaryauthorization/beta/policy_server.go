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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/binaryauthorization/beta/binaryauthorization_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization/beta"
)

// PolicyServer implements the gRPC interface for Policy.
type PolicyServer struct{}

// ProtoToPolicyClusterAdmissionRulesEvaluationModeEnum converts a PolicyClusterAdmissionRulesEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum(e betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum) *beta.PolicyClusterAdmissionRulesEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum_name[int32(e)]; ok {
		e := beta.PolicyClusterAdmissionRulesEvaluationModeEnum(n[len("BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyClusterAdmissionRulesEnforcementModeEnum converts a PolicyClusterAdmissionRulesEnforcementModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum(e betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum) *beta.PolicyClusterAdmissionRulesEnforcementModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum_name[int32(e)]; ok {
		e := beta.PolicyClusterAdmissionRulesEnforcementModeEnum(n[len("BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum converts a PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(e betapb.BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum) *beta.PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum_name[int32(e)]; ok {
		e := beta.PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(n[len("BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum converts a PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(e betapb.BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum) *beta.PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum_name[int32(e)]; ok {
		e := beta.PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(n[len("BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum converts a PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(e betapb.BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum) *beta.PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum_name[int32(e)]; ok {
		e := beta.PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(n[len("BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum converts a PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(e betapb.BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum) *beta.PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum_name[int32(e)]; ok {
		e := beta.PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(n[len("BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum converts a PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(e betapb.BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum) *beta.PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum_name[int32(e)]; ok {
		e := beta.PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(n[len("BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum converts a PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(e betapb.BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum) *beta.PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum_name[int32(e)]; ok {
		e := beta.PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(n[len("BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyDefaultAdmissionRuleEvaluationModeEnum converts a PolicyDefaultAdmissionRuleEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum(e betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum) *beta.PolicyDefaultAdmissionRuleEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum_name[int32(e)]; ok {
		e := beta.PolicyDefaultAdmissionRuleEvaluationModeEnum(n[len("BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyDefaultAdmissionRuleEnforcementModeEnum converts a PolicyDefaultAdmissionRuleEnforcementModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum(e betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum) *beta.PolicyDefaultAdmissionRuleEnforcementModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum_name[int32(e)]; ok {
		e := beta.PolicyDefaultAdmissionRuleEnforcementModeEnum(n[len("BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyGlobalPolicyEvaluationModeEnum converts a PolicyGlobalPolicyEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum(e betapb.BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum) *beta.PolicyGlobalPolicyEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum_name[int32(e)]; ok {
		e := beta.PolicyGlobalPolicyEvaluationModeEnum(n[len("BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyAdmissionWhitelistPatterns converts a PolicyAdmissionWhitelistPatterns object from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyAdmissionWhitelistPatterns(p *betapb.BinaryauthorizationBetaPolicyAdmissionWhitelistPatterns) *beta.PolicyAdmissionWhitelistPatterns {
	if p == nil {
		return nil
	}
	obj := &beta.PolicyAdmissionWhitelistPatterns{
		NamePattern: dcl.StringOrNil(p.GetNamePattern()),
	}
	return obj
}

// ProtoToPolicyClusterAdmissionRules converts a PolicyClusterAdmissionRules object from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyClusterAdmissionRules(p *betapb.BinaryauthorizationBetaPolicyClusterAdmissionRules) *beta.PolicyClusterAdmissionRules {
	if p == nil {
		return nil
	}
	obj := &beta.PolicyClusterAdmissionRules{
		EvaluationMode:  ProtoToBinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum(p.GetEvaluationMode()),
		EnforcementMode: ProtoToBinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum(p.GetEnforcementMode()),
	}
	for _, r := range p.GetRequireAttestationsBy() {
		obj.RequireAttestationsBy = append(obj.RequireAttestationsBy, r)
	}
	return obj
}

// ProtoToPolicyKubernetesNamespaceAdmissionRules converts a PolicyKubernetesNamespaceAdmissionRules object from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRules(p *betapb.BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRules) *beta.PolicyKubernetesNamespaceAdmissionRules {
	if p == nil {
		return nil
	}
	obj := &beta.PolicyKubernetesNamespaceAdmissionRules{
		EvaluationMode:  ProtoToBinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(p.GetEvaluationMode()),
		EnforcementMode: ProtoToBinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(p.GetEnforcementMode()),
	}
	for _, r := range p.GetRequireAttestationsBy() {
		obj.RequireAttestationsBy = append(obj.RequireAttestationsBy, r)
	}
	return obj
}

// ProtoToPolicyKubernetesServiceAccountAdmissionRules converts a PolicyKubernetesServiceAccountAdmissionRules object from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRules(p *betapb.BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRules) *beta.PolicyKubernetesServiceAccountAdmissionRules {
	if p == nil {
		return nil
	}
	obj := &beta.PolicyKubernetesServiceAccountAdmissionRules{
		EvaluationMode:  ProtoToBinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(p.GetEvaluationMode()),
		EnforcementMode: ProtoToBinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(p.GetEnforcementMode()),
	}
	for _, r := range p.GetRequireAttestationsBy() {
		obj.RequireAttestationsBy = append(obj.RequireAttestationsBy, r)
	}
	return obj
}

// ProtoToPolicyIstioServiceIdentityAdmissionRules converts a PolicyIstioServiceIdentityAdmissionRules object from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRules(p *betapb.BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRules) *beta.PolicyIstioServiceIdentityAdmissionRules {
	if p == nil {
		return nil
	}
	obj := &beta.PolicyIstioServiceIdentityAdmissionRules{
		EvaluationMode:  ProtoToBinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(p.GetEvaluationMode()),
		EnforcementMode: ProtoToBinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(p.GetEnforcementMode()),
	}
	for _, r := range p.GetRequireAttestationsBy() {
		obj.RequireAttestationsBy = append(obj.RequireAttestationsBy, r)
	}
	return obj
}

// ProtoToPolicyDefaultAdmissionRule converts a PolicyDefaultAdmissionRule object from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyDefaultAdmissionRule(p *betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRule) *beta.PolicyDefaultAdmissionRule {
	if p == nil {
		return nil
	}
	obj := &beta.PolicyDefaultAdmissionRule{
		EvaluationMode:  ProtoToBinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum(p.GetEvaluationMode()),
		EnforcementMode: ProtoToBinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum(p.GetEnforcementMode()),
	}
	for _, r := range p.GetRequireAttestationsBy() {
		obj.RequireAttestationsBy = append(obj.RequireAttestationsBy, r)
	}
	return obj
}

// ProtoToPolicy converts a Policy resource from its proto representation.
func ProtoToPolicy(p *betapb.BinaryauthorizationBetaPolicy) *beta.Policy {
	obj := &beta.Policy{
		DefaultAdmissionRule:       ProtoToBinaryauthorizationBetaPolicyDefaultAdmissionRule(p.GetDefaultAdmissionRule()),
		Description:                dcl.StringOrNil(p.GetDescription()),
		GlobalPolicyEvaluationMode: ProtoToBinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum(p.GetGlobalPolicyEvaluationMode()),
		SelfLink:                   dcl.StringOrNil(p.GetSelfLink()),
		Project:                    dcl.StringOrNil(p.GetProject()),
		UpdateTime:                 dcl.StringOrNil(p.GetUpdateTime()),
	}
	for _, r := range p.GetAdmissionWhitelistPatterns() {
		obj.AdmissionWhitelistPatterns = append(obj.AdmissionWhitelistPatterns, *ProtoToBinaryauthorizationBetaPolicyAdmissionWhitelistPatterns(r))
	}
	return obj
}

// PolicyClusterAdmissionRulesEvaluationModeEnumToProto converts a PolicyClusterAdmissionRulesEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnumToProto(e *beta.PolicyClusterAdmissionRulesEvaluationModeEnum) betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum {
	if e == nil {
		return betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum(0)
	}
	if v, ok := betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum_value["PolicyClusterAdmissionRulesEvaluationModeEnum"+string(*e)]; ok {
		return betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum(v)
	}
	return betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum(0)
}

// PolicyClusterAdmissionRulesEnforcementModeEnumToProto converts a PolicyClusterAdmissionRulesEnforcementModeEnum enum to its proto representation.
func BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnumToProto(e *beta.PolicyClusterAdmissionRulesEnforcementModeEnum) betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum {
	if e == nil {
		return betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum(0)
	}
	if v, ok := betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum_value["PolicyClusterAdmissionRulesEnforcementModeEnum"+string(*e)]; ok {
		return betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum(v)
	}
	return betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum(0)
}

// PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnumToProto converts a PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnumToProto(e *beta.PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum) betapb.BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum {
	if e == nil {
		return betapb.BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(0)
	}
	if v, ok := betapb.BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum_value["PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum"+string(*e)]; ok {
		return betapb.BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(v)
	}
	return betapb.BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(0)
}

// PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnumToProto converts a PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum enum to its proto representation.
func BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnumToProto(e *beta.PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum) betapb.BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum {
	if e == nil {
		return betapb.BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(0)
	}
	if v, ok := betapb.BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum_value["PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum"+string(*e)]; ok {
		return betapb.BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(v)
	}
	return betapb.BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(0)
}

// PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnumToProto converts a PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnumToProto(e *beta.PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum) betapb.BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum {
	if e == nil {
		return betapb.BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(0)
	}
	if v, ok := betapb.BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum_value["PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum"+string(*e)]; ok {
		return betapb.BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(v)
	}
	return betapb.BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(0)
}

// PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnumToProto converts a PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum enum to its proto representation.
func BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnumToProto(e *beta.PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum) betapb.BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum {
	if e == nil {
		return betapb.BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(0)
	}
	if v, ok := betapb.BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum_value["PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum"+string(*e)]; ok {
		return betapb.BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(v)
	}
	return betapb.BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(0)
}

// PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnumToProto converts a PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnumToProto(e *beta.PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum) betapb.BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum {
	if e == nil {
		return betapb.BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(0)
	}
	if v, ok := betapb.BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum_value["PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum"+string(*e)]; ok {
		return betapb.BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(v)
	}
	return betapb.BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(0)
}

// PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnumToProto converts a PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum enum to its proto representation.
func BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnumToProto(e *beta.PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum) betapb.BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum {
	if e == nil {
		return betapb.BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(0)
	}
	if v, ok := betapb.BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum_value["PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum"+string(*e)]; ok {
		return betapb.BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(v)
	}
	return betapb.BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(0)
}

// PolicyDefaultAdmissionRuleEvaluationModeEnumToProto converts a PolicyDefaultAdmissionRuleEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnumToProto(e *beta.PolicyDefaultAdmissionRuleEvaluationModeEnum) betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum {
	if e == nil {
		return betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum(0)
	}
	if v, ok := betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum_value["PolicyDefaultAdmissionRuleEvaluationModeEnum"+string(*e)]; ok {
		return betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum(v)
	}
	return betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum(0)
}

// PolicyDefaultAdmissionRuleEnforcementModeEnumToProto converts a PolicyDefaultAdmissionRuleEnforcementModeEnum enum to its proto representation.
func BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnumToProto(e *beta.PolicyDefaultAdmissionRuleEnforcementModeEnum) betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum {
	if e == nil {
		return betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum(0)
	}
	if v, ok := betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum_value["PolicyDefaultAdmissionRuleEnforcementModeEnum"+string(*e)]; ok {
		return betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum(v)
	}
	return betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum(0)
}

// PolicyGlobalPolicyEvaluationModeEnumToProto converts a PolicyGlobalPolicyEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnumToProto(e *beta.PolicyGlobalPolicyEvaluationModeEnum) betapb.BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum {
	if e == nil {
		return betapb.BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum(0)
	}
	if v, ok := betapb.BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum_value["PolicyGlobalPolicyEvaluationModeEnum"+string(*e)]; ok {
		return betapb.BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum(v)
	}
	return betapb.BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum(0)
}

// PolicyAdmissionWhitelistPatternsToProto converts a PolicyAdmissionWhitelistPatterns object to its proto representation.
func BinaryauthorizationBetaPolicyAdmissionWhitelistPatternsToProto(o *beta.PolicyAdmissionWhitelistPatterns) *betapb.BinaryauthorizationBetaPolicyAdmissionWhitelistPatterns {
	if o == nil {
		return nil
	}
	p := &betapb.BinaryauthorizationBetaPolicyAdmissionWhitelistPatterns{}
	p.SetNamePattern(dcl.ValueOrEmptyString(o.NamePattern))
	return p
}

// PolicyClusterAdmissionRulesToProto converts a PolicyClusterAdmissionRules object to its proto representation.
func BinaryauthorizationBetaPolicyClusterAdmissionRulesToProto(o *beta.PolicyClusterAdmissionRules) *betapb.BinaryauthorizationBetaPolicyClusterAdmissionRules {
	if o == nil {
		return nil
	}
	p := &betapb.BinaryauthorizationBetaPolicyClusterAdmissionRules{}
	p.SetEvaluationMode(BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnumToProto(o.EvaluationMode))
	p.SetEnforcementMode(BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnumToProto(o.EnforcementMode))
	sRequireAttestationsBy := make([]string, len(o.RequireAttestationsBy))
	for i, r := range o.RequireAttestationsBy {
		sRequireAttestationsBy[i] = r
	}
	p.SetRequireAttestationsBy(sRequireAttestationsBy)
	return p
}

// PolicyKubernetesNamespaceAdmissionRulesToProto converts a PolicyKubernetesNamespaceAdmissionRules object to its proto representation.
func BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesToProto(o *beta.PolicyKubernetesNamespaceAdmissionRules) *betapb.BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRules {
	if o == nil {
		return nil
	}
	p := &betapb.BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRules{}
	p.SetEvaluationMode(BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnumToProto(o.EvaluationMode))
	p.SetEnforcementMode(BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnumToProto(o.EnforcementMode))
	sRequireAttestationsBy := make([]string, len(o.RequireAttestationsBy))
	for i, r := range o.RequireAttestationsBy {
		sRequireAttestationsBy[i] = r
	}
	p.SetRequireAttestationsBy(sRequireAttestationsBy)
	return p
}

// PolicyKubernetesServiceAccountAdmissionRulesToProto converts a PolicyKubernetesServiceAccountAdmissionRules object to its proto representation.
func BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesToProto(o *beta.PolicyKubernetesServiceAccountAdmissionRules) *betapb.BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRules {
	if o == nil {
		return nil
	}
	p := &betapb.BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRules{}
	p.SetEvaluationMode(BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnumToProto(o.EvaluationMode))
	p.SetEnforcementMode(BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnumToProto(o.EnforcementMode))
	sRequireAttestationsBy := make([]string, len(o.RequireAttestationsBy))
	for i, r := range o.RequireAttestationsBy {
		sRequireAttestationsBy[i] = r
	}
	p.SetRequireAttestationsBy(sRequireAttestationsBy)
	return p
}

// PolicyIstioServiceIdentityAdmissionRulesToProto converts a PolicyIstioServiceIdentityAdmissionRules object to its proto representation.
func BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesToProto(o *beta.PolicyIstioServiceIdentityAdmissionRules) *betapb.BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRules {
	if o == nil {
		return nil
	}
	p := &betapb.BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRules{}
	p.SetEvaluationMode(BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnumToProto(o.EvaluationMode))
	p.SetEnforcementMode(BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnumToProto(o.EnforcementMode))
	sRequireAttestationsBy := make([]string, len(o.RequireAttestationsBy))
	for i, r := range o.RequireAttestationsBy {
		sRequireAttestationsBy[i] = r
	}
	p.SetRequireAttestationsBy(sRequireAttestationsBy)
	return p
}

// PolicyDefaultAdmissionRuleToProto converts a PolicyDefaultAdmissionRule object to its proto representation.
func BinaryauthorizationBetaPolicyDefaultAdmissionRuleToProto(o *beta.PolicyDefaultAdmissionRule) *betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRule {
	if o == nil {
		return nil
	}
	p := &betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRule{}
	p.SetEvaluationMode(BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnumToProto(o.EvaluationMode))
	p.SetEnforcementMode(BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnumToProto(o.EnforcementMode))
	sRequireAttestationsBy := make([]string, len(o.RequireAttestationsBy))
	for i, r := range o.RequireAttestationsBy {
		sRequireAttestationsBy[i] = r
	}
	p.SetRequireAttestationsBy(sRequireAttestationsBy)
	return p
}

// PolicyToProto converts a Policy resource to its proto representation.
func PolicyToProto(resource *beta.Policy) *betapb.BinaryauthorizationBetaPolicy {
	p := &betapb.BinaryauthorizationBetaPolicy{}
	p.SetDefaultAdmissionRule(BinaryauthorizationBetaPolicyDefaultAdmissionRuleToProto(resource.DefaultAdmissionRule))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetGlobalPolicyEvaluationMode(BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnumToProto(resource.GlobalPolicyEvaluationMode))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	sAdmissionWhitelistPatterns := make([]*betapb.BinaryauthorizationBetaPolicyAdmissionWhitelistPatterns, len(resource.AdmissionWhitelistPatterns))
	for i, r := range resource.AdmissionWhitelistPatterns {
		sAdmissionWhitelistPatterns[i] = BinaryauthorizationBetaPolicyAdmissionWhitelistPatternsToProto(&r)
	}
	p.SetAdmissionWhitelistPatterns(sAdmissionWhitelistPatterns)
	mClusterAdmissionRules := make(map[string]*betapb.BinaryauthorizationBetaPolicyClusterAdmissionRules, len(resource.ClusterAdmissionRules))
	for k, r := range resource.ClusterAdmissionRules {
		mClusterAdmissionRules[k] = BinaryauthorizationBetaPolicyClusterAdmissionRulesToProto(&r)
	}
	p.SetClusterAdmissionRules(mClusterAdmissionRules)
	mKubernetesNamespaceAdmissionRules := make(map[string]*betapb.BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRules, len(resource.KubernetesNamespaceAdmissionRules))
	for k, r := range resource.KubernetesNamespaceAdmissionRules {
		mKubernetesNamespaceAdmissionRules[k] = BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesToProto(&r)
	}
	p.SetKubernetesNamespaceAdmissionRules(mKubernetesNamespaceAdmissionRules)
	mKubernetesServiceAccountAdmissionRules := make(map[string]*betapb.BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRules, len(resource.KubernetesServiceAccountAdmissionRules))
	for k, r := range resource.KubernetesServiceAccountAdmissionRules {
		mKubernetesServiceAccountAdmissionRules[k] = BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesToProto(&r)
	}
	p.SetKubernetesServiceAccountAdmissionRules(mKubernetesServiceAccountAdmissionRules)
	mIstioServiceIdentityAdmissionRules := make(map[string]*betapb.BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRules, len(resource.IstioServiceIdentityAdmissionRules))
	for k, r := range resource.IstioServiceIdentityAdmissionRules {
		mIstioServiceIdentityAdmissionRules[k] = BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesToProto(&r)
	}
	p.SetIstioServiceIdentityAdmissionRules(mIstioServiceIdentityAdmissionRules)

	return p
}

// applyPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) applyPolicy(ctx context.Context, c *beta.Client, request *betapb.ApplyBinaryauthorizationBetaPolicyRequest) (*betapb.BinaryauthorizationBetaPolicy, error) {
	p := ProtoToPolicy(request.GetResource())
	res, err := c.ApplyPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := PolicyToProto(res)
	return r, nil
}

// applyBinaryauthorizationBetaPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) ApplyBinaryauthorizationBetaPolicy(ctx context.Context, request *betapb.ApplyBinaryauthorizationBetaPolicyRequest) (*betapb.BinaryauthorizationBetaPolicy, error) {
	cl, err := createConfigPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyPolicy(ctx, cl, request)
}

// DeletePolicy handles the gRPC request by passing it to the underlying Policy Delete() method.
func (s *PolicyServer) DeleteBinaryauthorizationBetaPolicy(ctx context.Context, request *betapb.DeleteBinaryauthorizationBetaPolicyRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for Policy")

}

// ListBinaryauthorizationBetaPolicy is a no-op method because Policy has no list method.
func (s *PolicyServer) ListBinaryauthorizationBetaPolicy(_ context.Context, _ *betapb.ListBinaryauthorizationBetaPolicyRequest) (*betapb.ListBinaryauthorizationBetaPolicyResponse, error) {
	return nil, nil
}

func createConfigPolicy(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
