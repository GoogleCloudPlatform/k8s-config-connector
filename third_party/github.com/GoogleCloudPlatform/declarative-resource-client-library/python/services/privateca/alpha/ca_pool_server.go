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
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/privateca/alpha/privateca_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/privateca/alpha"
)

// CaPoolServer implements the gRPC interface for CaPool.
type CaPoolServer struct{}

// ProtoToCaPoolTierEnum converts a CaPoolTierEnum enum from its proto representation.
func ProtoToPrivatecaAlphaCaPoolTierEnum(e alphapb.PrivatecaAlphaCaPoolTierEnum) *alpha.CaPoolTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.PrivatecaAlphaCaPoolTierEnum_name[int32(e)]; ok {
		e := alpha.CaPoolTierEnum(n[len("PrivatecaAlphaCaPoolTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum converts a CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum enum from its proto representation.
func ProtoToPrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(e alphapb.PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum) *alpha.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum_name[int32(e)]; ok {
		e := alpha.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(n[len("PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum converts a CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum enum from its proto representation.
func ProtoToPrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(e alphapb.PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum) *alpha.CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum_name[int32(e)]; ok {
		e := alpha.CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(n[len("PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum"):])
		return &e
	}
	return nil
}

// ProtoToCaPoolIssuancePolicy converts a CaPoolIssuancePolicy object from its proto representation.
func ProtoToPrivatecaAlphaCaPoolIssuancePolicy(p *alphapb.PrivatecaAlphaCaPoolIssuancePolicy) *alpha.CaPoolIssuancePolicy {
	if p == nil {
		return nil
	}
	obj := &alpha.CaPoolIssuancePolicy{
		MaximumLifetime:       dcl.StringOrNil(p.GetMaximumLifetime()),
		AllowedIssuanceModes:  ProtoToPrivatecaAlphaCaPoolIssuancePolicyAllowedIssuanceModes(p.GetAllowedIssuanceModes()),
		BaselineValues:        ProtoToPrivatecaAlphaCaPoolIssuancePolicyBaselineValues(p.GetBaselineValues()),
		IdentityConstraints:   ProtoToPrivatecaAlphaCaPoolIssuancePolicyIdentityConstraints(p.GetIdentityConstraints()),
		PassthroughExtensions: ProtoToPrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensions(p.GetPassthroughExtensions()),
	}
	for _, r := range p.GetAllowedKeyTypes() {
		obj.AllowedKeyTypes = append(obj.AllowedKeyTypes, *ProtoToPrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypes(r))
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyAllowedKeyTypes converts a CaPoolIssuancePolicyAllowedKeyTypes object from its proto representation.
func ProtoToPrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypes(p *alphapb.PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypes) *alpha.CaPoolIssuancePolicyAllowedKeyTypes {
	if p == nil {
		return nil
	}
	obj := &alpha.CaPoolIssuancePolicyAllowedKeyTypes{
		Rsa:           ProtoToPrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesRsa(p.GetRsa()),
		EllipticCurve: ProtoToPrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve(p.GetEllipticCurve()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyAllowedKeyTypesRsa converts a CaPoolIssuancePolicyAllowedKeyTypesRsa object from its proto representation.
func ProtoToPrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesRsa(p *alphapb.PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesRsa) *alpha.CaPoolIssuancePolicyAllowedKeyTypesRsa {
	if p == nil {
		return nil
	}
	obj := &alpha.CaPoolIssuancePolicyAllowedKeyTypesRsa{
		MinModulusSize: dcl.Int64OrNil(p.GetMinModulusSize()),
		MaxModulusSize: dcl.Int64OrNil(p.GetMaxModulusSize()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve converts a CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve object from its proto representation.
func ProtoToPrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve(p *alphapb.PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve) *alpha.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve {
	if p == nil {
		return nil
	}
	obj := &alpha.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve{
		SignatureAlgorithm: ProtoToPrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(p.GetSignatureAlgorithm()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyAllowedIssuanceModes converts a CaPoolIssuancePolicyAllowedIssuanceModes object from its proto representation.
func ProtoToPrivatecaAlphaCaPoolIssuancePolicyAllowedIssuanceModes(p *alphapb.PrivatecaAlphaCaPoolIssuancePolicyAllowedIssuanceModes) *alpha.CaPoolIssuancePolicyAllowedIssuanceModes {
	if p == nil {
		return nil
	}
	obj := &alpha.CaPoolIssuancePolicyAllowedIssuanceModes{
		AllowCsrBasedIssuance:    dcl.Bool(p.GetAllowCsrBasedIssuance()),
		AllowConfigBasedIssuance: dcl.Bool(p.GetAllowConfigBasedIssuance()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValues converts a CaPoolIssuancePolicyBaselineValues object from its proto representation.
func ProtoToPrivatecaAlphaCaPoolIssuancePolicyBaselineValues(p *alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValues) *alpha.CaPoolIssuancePolicyBaselineValues {
	if p == nil {
		return nil
	}
	obj := &alpha.CaPoolIssuancePolicyBaselineValues{
		KeyUsage:  ProtoToPrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaAlphaCaPoolIssuancePolicyBaselineValuesCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaAlphaCaPoolIssuancePolicyBaselineValuesPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaAlphaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesKeyUsage converts a CaPoolIssuancePolicyBaselineValuesKeyUsage object from its proto representation.
func ProtoToPrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsage(p *alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsage) *alpha.CaPoolIssuancePolicyBaselineValuesKeyUsage {
	if p == nil {
		return nil
	}
	obj := &alpha.CaPoolIssuancePolicyBaselineValuesKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage converts a CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage object from its proto representation.
func ProtoToPrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage(p *alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage) *alpha.CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &alpha.CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage{
		DigitalSignature:  dcl.Bool(p.GetDigitalSignature()),
		ContentCommitment: dcl.Bool(p.GetContentCommitment()),
		KeyEncipherment:   dcl.Bool(p.GetKeyEncipherment()),
		DataEncipherment:  dcl.Bool(p.GetDataEncipherment()),
		KeyAgreement:      dcl.Bool(p.GetKeyAgreement()),
		CertSign:          dcl.Bool(p.GetCertSign()),
		CrlSign:           dcl.Bool(p.GetCrlSign()),
		EncipherOnly:      dcl.Bool(p.GetEncipherOnly()),
		DecipherOnly:      dcl.Bool(p.GetDecipherOnly()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage converts a CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage object from its proto representation.
func ProtoToPrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage(p *alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage) *alpha.CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &alpha.CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage{
		ServerAuth:      dcl.Bool(p.GetServerAuth()),
		ClientAuth:      dcl.Bool(p.GetClientAuth()),
		CodeSigning:     dcl.Bool(p.GetCodeSigning()),
		EmailProtection: dcl.Bool(p.GetEmailProtection()),
		TimeStamping:    dcl.Bool(p.GetTimeStamping()),
		OcspSigning:     dcl.Bool(p.GetOcspSigning()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages converts a CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages object from its proto representation.
func ProtoToPrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages(p *alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages) *alpha.CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &alpha.CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesCaOptions converts a CaPoolIssuancePolicyBaselineValuesCaOptions object from its proto representation.
func ProtoToPrivatecaAlphaCaPoolIssuancePolicyBaselineValuesCaOptions(p *alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesCaOptions) *alpha.CaPoolIssuancePolicyBaselineValuesCaOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.CaPoolIssuancePolicyBaselineValuesCaOptions{
		IsCa:                    dcl.Bool(p.GetIsCa()),
		MaxIssuerPathLength:     dcl.Int64OrNil(p.GetMaxIssuerPathLength()),
		ZeroMaxIssuerPathLength: dcl.Bool(p.GetZeroMaxIssuerPathLength()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesPolicyIds converts a CaPoolIssuancePolicyBaselineValuesPolicyIds object from its proto representation.
func ProtoToPrivatecaAlphaCaPoolIssuancePolicyBaselineValuesPolicyIds(p *alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesPolicyIds) *alpha.CaPoolIssuancePolicyBaselineValuesPolicyIds {
	if p == nil {
		return nil
	}
	obj := &alpha.CaPoolIssuancePolicyBaselineValuesPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesAdditionalExtensions converts a CaPoolIssuancePolicyBaselineValuesAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaAlphaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions(p *alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions) *alpha.CaPoolIssuancePolicyBaselineValuesAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &alpha.CaPoolIssuancePolicyBaselineValuesAdditionalExtensions{
		ObjectId: ProtoToPrivatecaAlphaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId converts a CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId object from its proto representation.
func ProtoToPrivatecaAlphaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId(p *alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId) *alpha.CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &alpha.CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyIdentityConstraints converts a CaPoolIssuancePolicyIdentityConstraints object from its proto representation.
func ProtoToPrivatecaAlphaCaPoolIssuancePolicyIdentityConstraints(p *alphapb.PrivatecaAlphaCaPoolIssuancePolicyIdentityConstraints) *alpha.CaPoolIssuancePolicyIdentityConstraints {
	if p == nil {
		return nil
	}
	obj := &alpha.CaPoolIssuancePolicyIdentityConstraints{
		CelExpression:                   ProtoToPrivatecaAlphaCaPoolIssuancePolicyIdentityConstraintsCelExpression(p.GetCelExpression()),
		AllowSubjectPassthrough:         dcl.Bool(p.GetAllowSubjectPassthrough()),
		AllowSubjectAltNamesPassthrough: dcl.Bool(p.GetAllowSubjectAltNamesPassthrough()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyIdentityConstraintsCelExpression converts a CaPoolIssuancePolicyIdentityConstraintsCelExpression object from its proto representation.
func ProtoToPrivatecaAlphaCaPoolIssuancePolicyIdentityConstraintsCelExpression(p *alphapb.PrivatecaAlphaCaPoolIssuancePolicyIdentityConstraintsCelExpression) *alpha.CaPoolIssuancePolicyIdentityConstraintsCelExpression {
	if p == nil {
		return nil
	}
	obj := &alpha.CaPoolIssuancePolicyIdentityConstraintsCelExpression{
		Expression:  dcl.StringOrNil(p.GetExpression()),
		Title:       dcl.StringOrNil(p.GetTitle()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyPassthroughExtensions converts a CaPoolIssuancePolicyPassthroughExtensions object from its proto representation.
func ProtoToPrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensions(p *alphapb.PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensions) *alpha.CaPoolIssuancePolicyPassthroughExtensions {
	if p == nil {
		return nil
	}
	obj := &alpha.CaPoolIssuancePolicyPassthroughExtensions{}
	for _, r := range p.GetKnownExtensions() {
		obj.KnownExtensions = append(obj.KnownExtensions, *ProtoToPrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(r))
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions converts a CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions(p *alphapb.PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions) *alpha.CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &alpha.CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCaPoolPublishingOptions converts a CaPoolPublishingOptions object from its proto representation.
func ProtoToPrivatecaAlphaCaPoolPublishingOptions(p *alphapb.PrivatecaAlphaCaPoolPublishingOptions) *alpha.CaPoolPublishingOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.CaPoolPublishingOptions{
		PublishCaCert: dcl.Bool(p.GetPublishCaCert()),
		PublishCrl:    dcl.Bool(p.GetPublishCrl()),
	}
	return obj
}

// ProtoToCaPool converts a CaPool resource from its proto representation.
func ProtoToCaPool(p *alphapb.PrivatecaAlphaCaPool) *alpha.CaPool {
	obj := &alpha.CaPool{
		Name:              dcl.StringOrNil(p.GetName()),
		Tier:              ProtoToPrivatecaAlphaCaPoolTierEnum(p.GetTier()),
		IssuancePolicy:    ProtoToPrivatecaAlphaCaPoolIssuancePolicy(p.GetIssuancePolicy()),
		PublishingOptions: ProtoToPrivatecaAlphaCaPoolPublishingOptions(p.GetPublishingOptions()),
		Project:           dcl.StringOrNil(p.GetProject()),
		Location:          dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// CaPoolTierEnumToProto converts a CaPoolTierEnum enum to its proto representation.
func PrivatecaAlphaCaPoolTierEnumToProto(e *alpha.CaPoolTierEnum) alphapb.PrivatecaAlphaCaPoolTierEnum {
	if e == nil {
		return alphapb.PrivatecaAlphaCaPoolTierEnum(0)
	}
	if v, ok := alphapb.PrivatecaAlphaCaPoolTierEnum_value["CaPoolTierEnum"+string(*e)]; ok {
		return alphapb.PrivatecaAlphaCaPoolTierEnum(v)
	}
	return alphapb.PrivatecaAlphaCaPoolTierEnum(0)
}

// CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnumToProto converts a CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum enum to its proto representation.
func PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnumToProto(e *alpha.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum) alphapb.PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum {
	if e == nil {
		return alphapb.PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(0)
	}
	if v, ok := alphapb.PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum_value["CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum"+string(*e)]; ok {
		return alphapb.PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(v)
	}
	return alphapb.PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(0)
}

// CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnumToProto converts a CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum enum to its proto representation.
func PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnumToProto(e *alpha.CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum) alphapb.PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum {
	if e == nil {
		return alphapb.PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(0)
	}
	if v, ok := alphapb.PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum_value["CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum"+string(*e)]; ok {
		return alphapb.PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(v)
	}
	return alphapb.PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(0)
}

// CaPoolIssuancePolicyToProto converts a CaPoolIssuancePolicy object to its proto representation.
func PrivatecaAlphaCaPoolIssuancePolicyToProto(o *alpha.CaPoolIssuancePolicy) *alphapb.PrivatecaAlphaCaPoolIssuancePolicy {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCaPoolIssuancePolicy{}
	p.SetMaximumLifetime(dcl.ValueOrEmptyString(o.MaximumLifetime))
	p.SetAllowedIssuanceModes(PrivatecaAlphaCaPoolIssuancePolicyAllowedIssuanceModesToProto(o.AllowedIssuanceModes))
	p.SetBaselineValues(PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesToProto(o.BaselineValues))
	p.SetIdentityConstraints(PrivatecaAlphaCaPoolIssuancePolicyIdentityConstraintsToProto(o.IdentityConstraints))
	p.SetPassthroughExtensions(PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsToProto(o.PassthroughExtensions))
	sAllowedKeyTypes := make([]*alphapb.PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypes, len(o.AllowedKeyTypes))
	for i, r := range o.AllowedKeyTypes {
		sAllowedKeyTypes[i] = PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesToProto(&r)
	}
	p.SetAllowedKeyTypes(sAllowedKeyTypes)
	return p
}

// CaPoolIssuancePolicyAllowedKeyTypesToProto converts a CaPoolIssuancePolicyAllowedKeyTypes object to its proto representation.
func PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesToProto(o *alpha.CaPoolIssuancePolicyAllowedKeyTypes) *alphapb.PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypes {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypes{}
	p.SetRsa(PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesRsaToProto(o.Rsa))
	p.SetEllipticCurve(PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveToProto(o.EllipticCurve))
	return p
}

// CaPoolIssuancePolicyAllowedKeyTypesRsaToProto converts a CaPoolIssuancePolicyAllowedKeyTypesRsa object to its proto representation.
func PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesRsaToProto(o *alpha.CaPoolIssuancePolicyAllowedKeyTypesRsa) *alphapb.PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesRsa {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesRsa{}
	p.SetMinModulusSize(dcl.ValueOrEmptyInt64(o.MinModulusSize))
	p.SetMaxModulusSize(dcl.ValueOrEmptyInt64(o.MaxModulusSize))
	return p
}

// CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveToProto converts a CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve object to its proto representation.
func PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveToProto(o *alpha.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve) *alphapb.PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve{}
	p.SetSignatureAlgorithm(PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnumToProto(o.SignatureAlgorithm))
	return p
}

// CaPoolIssuancePolicyAllowedIssuanceModesToProto converts a CaPoolIssuancePolicyAllowedIssuanceModes object to its proto representation.
func PrivatecaAlphaCaPoolIssuancePolicyAllowedIssuanceModesToProto(o *alpha.CaPoolIssuancePolicyAllowedIssuanceModes) *alphapb.PrivatecaAlphaCaPoolIssuancePolicyAllowedIssuanceModes {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCaPoolIssuancePolicyAllowedIssuanceModes{}
	p.SetAllowCsrBasedIssuance(dcl.ValueOrEmptyBool(o.AllowCsrBasedIssuance))
	p.SetAllowConfigBasedIssuance(dcl.ValueOrEmptyBool(o.AllowConfigBasedIssuance))
	return p
}

// CaPoolIssuancePolicyBaselineValuesToProto converts a CaPoolIssuancePolicyBaselineValues object to its proto representation.
func PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesToProto(o *alpha.CaPoolIssuancePolicyBaselineValues) *alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValues {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValues{}
	p.SetKeyUsage(PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageToProto(o.KeyUsage))
	p.SetCaOptions(PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesCaOptionsToProto(o.CaOptions))
	sPolicyIds := make([]*alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesPolicyIds, len(o.PolicyIds))
	for i, r := range o.PolicyIds {
		sPolicyIds[i] = PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesPolicyIdsToProto(&r)
	}
	p.SetPolicyIds(sPolicyIds)
	sAiaOcspServers := make([]string, len(o.AiaOcspServers))
	for i, r := range o.AiaOcspServers {
		sAiaOcspServers[i] = r
	}
	p.SetAiaOcspServers(sAiaOcspServers)
	sAdditionalExtensions := make([]*alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CaPoolIssuancePolicyBaselineValuesKeyUsageToProto converts a CaPoolIssuancePolicyBaselineValuesKeyUsage object to its proto representation.
func PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageToProto(o *alpha.CaPoolIssuancePolicyBaselineValuesKeyUsage) *alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsage {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsage{}
	p.SetBaseKeyUsage(PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage))
	p.SetExtendedKeyUsage(PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage))
	sUnknownExtendedKeyUsages := make([]*alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages, len(o.UnknownExtendedKeyUsages))
	for i, r := range o.UnknownExtendedKeyUsages {
		sUnknownExtendedKeyUsages[i] = PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesToProto(&r)
	}
	p.SetUnknownExtendedKeyUsages(sUnknownExtendedKeyUsages)
	return p
}

// CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsageToProto converts a CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage object to its proto representation.
func PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsageToProto(o *alpha.CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage) *alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage{}
	p.SetDigitalSignature(dcl.ValueOrEmptyBool(o.DigitalSignature))
	p.SetContentCommitment(dcl.ValueOrEmptyBool(o.ContentCommitment))
	p.SetKeyEncipherment(dcl.ValueOrEmptyBool(o.KeyEncipherment))
	p.SetDataEncipherment(dcl.ValueOrEmptyBool(o.DataEncipherment))
	p.SetKeyAgreement(dcl.ValueOrEmptyBool(o.KeyAgreement))
	p.SetCertSign(dcl.ValueOrEmptyBool(o.CertSign))
	p.SetCrlSign(dcl.ValueOrEmptyBool(o.CrlSign))
	p.SetEncipherOnly(dcl.ValueOrEmptyBool(o.EncipherOnly))
	p.SetDecipherOnly(dcl.ValueOrEmptyBool(o.DecipherOnly))
	return p
}

// CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageToProto converts a CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage object to its proto representation.
func PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageToProto(o *alpha.CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage) *alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage{}
	p.SetServerAuth(dcl.ValueOrEmptyBool(o.ServerAuth))
	p.SetClientAuth(dcl.ValueOrEmptyBool(o.ClientAuth))
	p.SetCodeSigning(dcl.ValueOrEmptyBool(o.CodeSigning))
	p.SetEmailProtection(dcl.ValueOrEmptyBool(o.EmailProtection))
	p.SetTimeStamping(dcl.ValueOrEmptyBool(o.TimeStamping))
	p.SetOcspSigning(dcl.ValueOrEmptyBool(o.OcspSigning))
	return p
}

// CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesToProto converts a CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages object to its proto representation.
func PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesToProto(o *alpha.CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages) *alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CaPoolIssuancePolicyBaselineValuesCaOptionsToProto converts a CaPoolIssuancePolicyBaselineValuesCaOptions object to its proto representation.
func PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesCaOptionsToProto(o *alpha.CaPoolIssuancePolicyBaselineValuesCaOptions) *alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesCaOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesCaOptions{}
	p.SetIsCa(dcl.ValueOrEmptyBool(o.IsCa))
	p.SetMaxIssuerPathLength(dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength))
	p.SetZeroMaxIssuerPathLength(dcl.ValueOrEmptyBool(o.ZeroMaxIssuerPathLength))
	return p
}

// CaPoolIssuancePolicyBaselineValuesPolicyIdsToProto converts a CaPoolIssuancePolicyBaselineValuesPolicyIds object to its proto representation.
func PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesPolicyIdsToProto(o *alpha.CaPoolIssuancePolicyBaselineValuesPolicyIds) *alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesPolicyIds {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesPolicyIds{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsToProto converts a CaPoolIssuancePolicyBaselineValuesAdditionalExtensions object to its proto representation.
func PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsToProto(o *alpha.CaPoolIssuancePolicyBaselineValuesAdditionalExtensions) *alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions{}
	p.SetObjectId(PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectIdToProto converts a CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId object to its proto representation.
func PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectIdToProto(o *alpha.CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId) *alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CaPoolIssuancePolicyIdentityConstraintsToProto converts a CaPoolIssuancePolicyIdentityConstraints object to its proto representation.
func PrivatecaAlphaCaPoolIssuancePolicyIdentityConstraintsToProto(o *alpha.CaPoolIssuancePolicyIdentityConstraints) *alphapb.PrivatecaAlphaCaPoolIssuancePolicyIdentityConstraints {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCaPoolIssuancePolicyIdentityConstraints{}
	p.SetCelExpression(PrivatecaAlphaCaPoolIssuancePolicyIdentityConstraintsCelExpressionToProto(o.CelExpression))
	p.SetAllowSubjectPassthrough(dcl.ValueOrEmptyBool(o.AllowSubjectPassthrough))
	p.SetAllowSubjectAltNamesPassthrough(dcl.ValueOrEmptyBool(o.AllowSubjectAltNamesPassthrough))
	return p
}

// CaPoolIssuancePolicyIdentityConstraintsCelExpressionToProto converts a CaPoolIssuancePolicyIdentityConstraintsCelExpression object to its proto representation.
func PrivatecaAlphaCaPoolIssuancePolicyIdentityConstraintsCelExpressionToProto(o *alpha.CaPoolIssuancePolicyIdentityConstraintsCelExpression) *alphapb.PrivatecaAlphaCaPoolIssuancePolicyIdentityConstraintsCelExpression {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCaPoolIssuancePolicyIdentityConstraintsCelExpression{}
	p.SetExpression(dcl.ValueOrEmptyString(o.Expression))
	p.SetTitle(dcl.ValueOrEmptyString(o.Title))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	return p
}

// CaPoolIssuancePolicyPassthroughExtensionsToProto converts a CaPoolIssuancePolicyPassthroughExtensions object to its proto representation.
func PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsToProto(o *alpha.CaPoolIssuancePolicyPassthroughExtensions) *alphapb.PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensions {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensions{}
	sKnownExtensions := make([]alphapb.PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum, len(o.KnownExtensions))
	for i, r := range o.KnownExtensions {
		sKnownExtensions[i] = alphapb.PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(alphapb.PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum_value[string(r)])
	}
	p.SetKnownExtensions(sKnownExtensions)
	sAdditionalExtensions := make([]*alphapb.PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensionsToProto converts a CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions object to its proto representation.
func PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensionsToProto(o *alpha.CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions) *alphapb.PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CaPoolPublishingOptionsToProto converts a CaPoolPublishingOptions object to its proto representation.
func PrivatecaAlphaCaPoolPublishingOptionsToProto(o *alpha.CaPoolPublishingOptions) *alphapb.PrivatecaAlphaCaPoolPublishingOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCaPoolPublishingOptions{}
	p.SetPublishCaCert(dcl.ValueOrEmptyBool(o.PublishCaCert))
	p.SetPublishCrl(dcl.ValueOrEmptyBool(o.PublishCrl))
	return p
}

// CaPoolToProto converts a CaPool resource to its proto representation.
func CaPoolToProto(resource *alpha.CaPool) *alphapb.PrivatecaAlphaCaPool {
	p := &alphapb.PrivatecaAlphaCaPool{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetTier(PrivatecaAlphaCaPoolTierEnumToProto(resource.Tier))
	p.SetIssuancePolicy(PrivatecaAlphaCaPoolIssuancePolicyToProto(resource.IssuancePolicy))
	p.SetPublishingOptions(PrivatecaAlphaCaPoolPublishingOptionsToProto(resource.PublishingOptions))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyCaPool handles the gRPC request by passing it to the underlying CaPool Apply() method.
func (s *CaPoolServer) applyCaPool(ctx context.Context, c *alpha.Client, request *alphapb.ApplyPrivatecaAlphaCaPoolRequest) (*alphapb.PrivatecaAlphaCaPool, error) {
	p := ProtoToCaPool(request.GetResource())
	res, err := c.ApplyCaPool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CaPoolToProto(res)
	return r, nil
}

// applyPrivatecaAlphaCaPool handles the gRPC request by passing it to the underlying CaPool Apply() method.
func (s *CaPoolServer) ApplyPrivatecaAlphaCaPool(ctx context.Context, request *alphapb.ApplyPrivatecaAlphaCaPoolRequest) (*alphapb.PrivatecaAlphaCaPool, error) {
	cl, err := createConfigCaPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCaPool(ctx, cl, request)
}

// DeleteCaPool handles the gRPC request by passing it to the underlying CaPool Delete() method.
func (s *CaPoolServer) DeletePrivatecaAlphaCaPool(ctx context.Context, request *alphapb.DeletePrivatecaAlphaCaPoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCaPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCaPool(ctx, ProtoToCaPool(request.GetResource()))

}

// ListPrivatecaAlphaCaPool handles the gRPC request by passing it to the underlying CaPoolList() method.
func (s *CaPoolServer) ListPrivatecaAlphaCaPool(ctx context.Context, request *alphapb.ListPrivatecaAlphaCaPoolRequest) (*alphapb.ListPrivatecaAlphaCaPoolResponse, error) {
	cl, err := createConfigCaPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCaPool(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.PrivatecaAlphaCaPool
	for _, r := range resources.Items {
		rp := CaPoolToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListPrivatecaAlphaCaPoolResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCaPool(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
