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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/privateca/beta/privateca_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/privateca/beta"
)

// CaPoolServer implements the gRPC interface for CaPool.
type CaPoolServer struct{}

// ProtoToCaPoolTierEnum converts a CaPoolTierEnum enum from its proto representation.
func ProtoToPrivatecaBetaCaPoolTierEnum(e betapb.PrivatecaBetaCaPoolTierEnum) *beta.CaPoolTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCaPoolTierEnum_name[int32(e)]; ok {
		e := beta.CaPoolTierEnum(n[len("PrivatecaBetaCaPoolTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum converts a CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum enum from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(e betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum) *beta.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum_name[int32(e)]; ok {
		e := beta.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(n[len("PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum converts a CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum enum from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(e betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum) *beta.CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum_name[int32(e)]; ok {
		e := beta.CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(n[len("PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum"):])
		return &e
	}
	return nil
}

// ProtoToCaPoolIssuancePolicy converts a CaPoolIssuancePolicy object from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicy(p *betapb.PrivatecaBetaCaPoolIssuancePolicy) *beta.CaPoolIssuancePolicy {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicy{
		MaximumLifetime:       dcl.StringOrNil(p.GetMaximumLifetime()),
		AllowedIssuanceModes:  ProtoToPrivatecaBetaCaPoolIssuancePolicyAllowedIssuanceModes(p.GetAllowedIssuanceModes()),
		BaselineValues:        ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValues(p.GetBaselineValues()),
		IdentityConstraints:   ProtoToPrivatecaBetaCaPoolIssuancePolicyIdentityConstraints(p.GetIdentityConstraints()),
		PassthroughExtensions: ProtoToPrivatecaBetaCaPoolIssuancePolicyPassthroughExtensions(p.GetPassthroughExtensions()),
	}
	for _, r := range p.GetAllowedKeyTypes() {
		obj.AllowedKeyTypes = append(obj.AllowedKeyTypes, *ProtoToPrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypes(r))
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyAllowedKeyTypes converts a CaPoolIssuancePolicyAllowedKeyTypes object from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypes(p *betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypes) *beta.CaPoolIssuancePolicyAllowedKeyTypes {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyAllowedKeyTypes{
		Rsa:           ProtoToPrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesRsa(p.GetRsa()),
		EllipticCurve: ProtoToPrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve(p.GetEllipticCurve()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyAllowedKeyTypesRsa converts a CaPoolIssuancePolicyAllowedKeyTypesRsa object from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesRsa(p *betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesRsa) *beta.CaPoolIssuancePolicyAllowedKeyTypesRsa {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyAllowedKeyTypesRsa{
		MinModulusSize: dcl.Int64OrNil(p.GetMinModulusSize()),
		MaxModulusSize: dcl.Int64OrNil(p.GetMaxModulusSize()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve converts a CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve object from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve(p *betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve) *beta.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve{
		SignatureAlgorithm: ProtoToPrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(p.GetSignatureAlgorithm()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyAllowedIssuanceModes converts a CaPoolIssuancePolicyAllowedIssuanceModes object from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyAllowedIssuanceModes(p *betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedIssuanceModes) *beta.CaPoolIssuancePolicyAllowedIssuanceModes {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyAllowedIssuanceModes{
		AllowCsrBasedIssuance:    dcl.Bool(p.GetAllowCsrBasedIssuance()),
		AllowConfigBasedIssuance: dcl.Bool(p.GetAllowConfigBasedIssuance()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValues converts a CaPoolIssuancePolicyBaselineValues object from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValues(p *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValues) *beta.CaPoolIssuancePolicyBaselineValues {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyBaselineValues{
		KeyUsage:  ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesKeyUsage converts a CaPoolIssuancePolicyBaselineValuesKeyUsage object from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsage(p *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsage) *beta.CaPoolIssuancePolicyBaselineValuesKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyBaselineValuesKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage converts a CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage object from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage(p *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage) *beta.CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage{
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
func ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage(p *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage) *beta.CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage{
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
func ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages(p *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages) *beta.CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesCaOptions converts a CaPoolIssuancePolicyBaselineValuesCaOptions object from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesCaOptions(p *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesCaOptions) *beta.CaPoolIssuancePolicyBaselineValuesCaOptions {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyBaselineValuesCaOptions{
		IsCa:                    dcl.Bool(p.GetIsCa()),
		MaxIssuerPathLength:     dcl.Int64OrNil(p.GetMaxIssuerPathLength()),
		ZeroMaxIssuerPathLength: dcl.Bool(p.GetZeroMaxIssuerPathLength()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesPolicyIds converts a CaPoolIssuancePolicyBaselineValuesPolicyIds object from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesPolicyIds(p *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesPolicyIds) *beta.CaPoolIssuancePolicyBaselineValuesPolicyIds {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyBaselineValuesPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesAdditionalExtensions converts a CaPoolIssuancePolicyBaselineValuesAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions(p *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions) *beta.CaPoolIssuancePolicyBaselineValuesAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyBaselineValuesAdditionalExtensions{
		ObjectId: ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId converts a CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId object from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId(p *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId) *beta.CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyIdentityConstraints converts a CaPoolIssuancePolicyIdentityConstraints object from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyIdentityConstraints(p *betapb.PrivatecaBetaCaPoolIssuancePolicyIdentityConstraints) *beta.CaPoolIssuancePolicyIdentityConstraints {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyIdentityConstraints{
		CelExpression:                   ProtoToPrivatecaBetaCaPoolIssuancePolicyIdentityConstraintsCelExpression(p.GetCelExpression()),
		AllowSubjectPassthrough:         dcl.Bool(p.GetAllowSubjectPassthrough()),
		AllowSubjectAltNamesPassthrough: dcl.Bool(p.GetAllowSubjectAltNamesPassthrough()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyIdentityConstraintsCelExpression converts a CaPoolIssuancePolicyIdentityConstraintsCelExpression object from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyIdentityConstraintsCelExpression(p *betapb.PrivatecaBetaCaPoolIssuancePolicyIdentityConstraintsCelExpression) *beta.CaPoolIssuancePolicyIdentityConstraintsCelExpression {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyIdentityConstraintsCelExpression{
		Expression:  dcl.StringOrNil(p.GetExpression()),
		Title:       dcl.StringOrNil(p.GetTitle()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyPassthroughExtensions converts a CaPoolIssuancePolicyPassthroughExtensions object from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyPassthroughExtensions(p *betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensions) *beta.CaPoolIssuancePolicyPassthroughExtensions {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyPassthroughExtensions{}
	for _, r := range p.GetKnownExtensions() {
		obj.KnownExtensions = append(obj.KnownExtensions, *ProtoToPrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(r))
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions converts a CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions(p *betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions) *beta.CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCaPoolPublishingOptions converts a CaPoolPublishingOptions object from its proto representation.
func ProtoToPrivatecaBetaCaPoolPublishingOptions(p *betapb.PrivatecaBetaCaPoolPublishingOptions) *beta.CaPoolPublishingOptions {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolPublishingOptions{
		PublishCaCert: dcl.Bool(p.GetPublishCaCert()),
		PublishCrl:    dcl.Bool(p.GetPublishCrl()),
	}
	return obj
}

// ProtoToCaPool converts a CaPool resource from its proto representation.
func ProtoToCaPool(p *betapb.PrivatecaBetaCaPool) *beta.CaPool {
	obj := &beta.CaPool{
		Name:              dcl.StringOrNil(p.GetName()),
		Tier:              ProtoToPrivatecaBetaCaPoolTierEnum(p.GetTier()),
		IssuancePolicy:    ProtoToPrivatecaBetaCaPoolIssuancePolicy(p.GetIssuancePolicy()),
		PublishingOptions: ProtoToPrivatecaBetaCaPoolPublishingOptions(p.GetPublishingOptions()),
		Project:           dcl.StringOrNil(p.GetProject()),
		Location:          dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// CaPoolTierEnumToProto converts a CaPoolTierEnum enum to its proto representation.
func PrivatecaBetaCaPoolTierEnumToProto(e *beta.CaPoolTierEnum) betapb.PrivatecaBetaCaPoolTierEnum {
	if e == nil {
		return betapb.PrivatecaBetaCaPoolTierEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCaPoolTierEnum_value["CaPoolTierEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCaPoolTierEnum(v)
	}
	return betapb.PrivatecaBetaCaPoolTierEnum(0)
}

// CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnumToProto converts a CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum enum to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnumToProto(e *beta.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum) betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum {
	if e == nil {
		return betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum_value["CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(v)
	}
	return betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(0)
}

// CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnumToProto converts a CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum enum to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnumToProto(e *beta.CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum) betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum {
	if e == nil {
		return betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum_value["CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(v)
	}
	return betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(0)
}

// CaPoolIssuancePolicyToProto converts a CaPoolIssuancePolicy object to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyToProto(o *beta.CaPoolIssuancePolicy) *betapb.PrivatecaBetaCaPoolIssuancePolicy {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicy{}
	p.SetMaximumLifetime(dcl.ValueOrEmptyString(o.MaximumLifetime))
	p.SetAllowedIssuanceModes(PrivatecaBetaCaPoolIssuancePolicyAllowedIssuanceModesToProto(o.AllowedIssuanceModes))
	p.SetBaselineValues(PrivatecaBetaCaPoolIssuancePolicyBaselineValuesToProto(o.BaselineValues))
	p.SetIdentityConstraints(PrivatecaBetaCaPoolIssuancePolicyIdentityConstraintsToProto(o.IdentityConstraints))
	p.SetPassthroughExtensions(PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsToProto(o.PassthroughExtensions))
	sAllowedKeyTypes := make([]*betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypes, len(o.AllowedKeyTypes))
	for i, r := range o.AllowedKeyTypes {
		sAllowedKeyTypes[i] = PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesToProto(&r)
	}
	p.SetAllowedKeyTypes(sAllowedKeyTypes)
	return p
}

// CaPoolIssuancePolicyAllowedKeyTypesToProto converts a CaPoolIssuancePolicyAllowedKeyTypes object to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesToProto(o *beta.CaPoolIssuancePolicyAllowedKeyTypes) *betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypes {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypes{}
	p.SetRsa(PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesRsaToProto(o.Rsa))
	p.SetEllipticCurve(PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveToProto(o.EllipticCurve))
	return p
}

// CaPoolIssuancePolicyAllowedKeyTypesRsaToProto converts a CaPoolIssuancePolicyAllowedKeyTypesRsa object to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesRsaToProto(o *beta.CaPoolIssuancePolicyAllowedKeyTypesRsa) *betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesRsa {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesRsa{}
	p.SetMinModulusSize(dcl.ValueOrEmptyInt64(o.MinModulusSize))
	p.SetMaxModulusSize(dcl.ValueOrEmptyInt64(o.MaxModulusSize))
	return p
}

// CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveToProto converts a CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve object to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveToProto(o *beta.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve) *betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve{}
	p.SetSignatureAlgorithm(PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnumToProto(o.SignatureAlgorithm))
	return p
}

// CaPoolIssuancePolicyAllowedIssuanceModesToProto converts a CaPoolIssuancePolicyAllowedIssuanceModes object to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyAllowedIssuanceModesToProto(o *beta.CaPoolIssuancePolicyAllowedIssuanceModes) *betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedIssuanceModes {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedIssuanceModes{}
	p.SetAllowCsrBasedIssuance(dcl.ValueOrEmptyBool(o.AllowCsrBasedIssuance))
	p.SetAllowConfigBasedIssuance(dcl.ValueOrEmptyBool(o.AllowConfigBasedIssuance))
	return p
}

// CaPoolIssuancePolicyBaselineValuesToProto converts a CaPoolIssuancePolicyBaselineValues object to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyBaselineValuesToProto(o *beta.CaPoolIssuancePolicyBaselineValues) *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValues {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValues{}
	p.SetKeyUsage(PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageToProto(o.KeyUsage))
	p.SetCaOptions(PrivatecaBetaCaPoolIssuancePolicyBaselineValuesCaOptionsToProto(o.CaOptions))
	sPolicyIds := make([]*betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesPolicyIds, len(o.PolicyIds))
	for i, r := range o.PolicyIds {
		sPolicyIds[i] = PrivatecaBetaCaPoolIssuancePolicyBaselineValuesPolicyIdsToProto(&r)
	}
	p.SetPolicyIds(sPolicyIds)
	sAiaOcspServers := make([]string, len(o.AiaOcspServers))
	for i, r := range o.AiaOcspServers {
		sAiaOcspServers[i] = r
	}
	p.SetAiaOcspServers(sAiaOcspServers)
	sAdditionalExtensions := make([]*betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CaPoolIssuancePolicyBaselineValuesKeyUsageToProto converts a CaPoolIssuancePolicyBaselineValuesKeyUsage object to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageToProto(o *beta.CaPoolIssuancePolicyBaselineValuesKeyUsage) *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsage{}
	p.SetBaseKeyUsage(PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage))
	p.SetExtendedKeyUsage(PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage))
	sUnknownExtendedKeyUsages := make([]*betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages, len(o.UnknownExtendedKeyUsages))
	for i, r := range o.UnknownExtendedKeyUsages {
		sUnknownExtendedKeyUsages[i] = PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesToProto(&r)
	}
	p.SetUnknownExtendedKeyUsages(sUnknownExtendedKeyUsages)
	return p
}

// CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsageToProto converts a CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage object to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsageToProto(o *beta.CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage) *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage{}
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
func PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageToProto(o *beta.CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage) *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage{}
	p.SetServerAuth(dcl.ValueOrEmptyBool(o.ServerAuth))
	p.SetClientAuth(dcl.ValueOrEmptyBool(o.ClientAuth))
	p.SetCodeSigning(dcl.ValueOrEmptyBool(o.CodeSigning))
	p.SetEmailProtection(dcl.ValueOrEmptyBool(o.EmailProtection))
	p.SetTimeStamping(dcl.ValueOrEmptyBool(o.TimeStamping))
	p.SetOcspSigning(dcl.ValueOrEmptyBool(o.OcspSigning))
	return p
}

// CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesToProto converts a CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages object to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesToProto(o *beta.CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages) *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CaPoolIssuancePolicyBaselineValuesCaOptionsToProto converts a CaPoolIssuancePolicyBaselineValuesCaOptions object to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyBaselineValuesCaOptionsToProto(o *beta.CaPoolIssuancePolicyBaselineValuesCaOptions) *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesCaOptions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesCaOptions{}
	p.SetIsCa(dcl.ValueOrEmptyBool(o.IsCa))
	p.SetMaxIssuerPathLength(dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength))
	p.SetZeroMaxIssuerPathLength(dcl.ValueOrEmptyBool(o.ZeroMaxIssuerPathLength))
	return p
}

// CaPoolIssuancePolicyBaselineValuesPolicyIdsToProto converts a CaPoolIssuancePolicyBaselineValuesPolicyIds object to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyBaselineValuesPolicyIdsToProto(o *beta.CaPoolIssuancePolicyBaselineValuesPolicyIds) *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesPolicyIds {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesPolicyIds{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsToProto converts a CaPoolIssuancePolicyBaselineValuesAdditionalExtensions object to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsToProto(o *beta.CaPoolIssuancePolicyBaselineValuesAdditionalExtensions) *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions{}
	p.SetObjectId(PrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectIdToProto converts a CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId object to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectIdToProto(o *beta.CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId) *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CaPoolIssuancePolicyIdentityConstraintsToProto converts a CaPoolIssuancePolicyIdentityConstraints object to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyIdentityConstraintsToProto(o *beta.CaPoolIssuancePolicyIdentityConstraints) *betapb.PrivatecaBetaCaPoolIssuancePolicyIdentityConstraints {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyIdentityConstraints{}
	p.SetCelExpression(PrivatecaBetaCaPoolIssuancePolicyIdentityConstraintsCelExpressionToProto(o.CelExpression))
	p.SetAllowSubjectPassthrough(dcl.ValueOrEmptyBool(o.AllowSubjectPassthrough))
	p.SetAllowSubjectAltNamesPassthrough(dcl.ValueOrEmptyBool(o.AllowSubjectAltNamesPassthrough))
	return p
}

// CaPoolIssuancePolicyIdentityConstraintsCelExpressionToProto converts a CaPoolIssuancePolicyIdentityConstraintsCelExpression object to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyIdentityConstraintsCelExpressionToProto(o *beta.CaPoolIssuancePolicyIdentityConstraintsCelExpression) *betapb.PrivatecaBetaCaPoolIssuancePolicyIdentityConstraintsCelExpression {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyIdentityConstraintsCelExpression{}
	p.SetExpression(dcl.ValueOrEmptyString(o.Expression))
	p.SetTitle(dcl.ValueOrEmptyString(o.Title))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	return p
}

// CaPoolIssuancePolicyPassthroughExtensionsToProto converts a CaPoolIssuancePolicyPassthroughExtensions object to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsToProto(o *beta.CaPoolIssuancePolicyPassthroughExtensions) *betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensions{}
	sKnownExtensions := make([]betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum, len(o.KnownExtensions))
	for i, r := range o.KnownExtensions {
		sKnownExtensions[i] = betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum_value[string(r)])
	}
	p.SetKnownExtensions(sKnownExtensions)
	sAdditionalExtensions := make([]*betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensionsToProto converts a CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions object to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensionsToProto(o *beta.CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions) *betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CaPoolPublishingOptionsToProto converts a CaPoolPublishingOptions object to its proto representation.
func PrivatecaBetaCaPoolPublishingOptionsToProto(o *beta.CaPoolPublishingOptions) *betapb.PrivatecaBetaCaPoolPublishingOptions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolPublishingOptions{}
	p.SetPublishCaCert(dcl.ValueOrEmptyBool(o.PublishCaCert))
	p.SetPublishCrl(dcl.ValueOrEmptyBool(o.PublishCrl))
	return p
}

// CaPoolToProto converts a CaPool resource to its proto representation.
func CaPoolToProto(resource *beta.CaPool) *betapb.PrivatecaBetaCaPool {
	p := &betapb.PrivatecaBetaCaPool{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetTier(PrivatecaBetaCaPoolTierEnumToProto(resource.Tier))
	p.SetIssuancePolicy(PrivatecaBetaCaPoolIssuancePolicyToProto(resource.IssuancePolicy))
	p.SetPublishingOptions(PrivatecaBetaCaPoolPublishingOptionsToProto(resource.PublishingOptions))
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
func (s *CaPoolServer) applyCaPool(ctx context.Context, c *beta.Client, request *betapb.ApplyPrivatecaBetaCaPoolRequest) (*betapb.PrivatecaBetaCaPool, error) {
	p := ProtoToCaPool(request.GetResource())
	res, err := c.ApplyCaPool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CaPoolToProto(res)
	return r, nil
}

// applyPrivatecaBetaCaPool handles the gRPC request by passing it to the underlying CaPool Apply() method.
func (s *CaPoolServer) ApplyPrivatecaBetaCaPool(ctx context.Context, request *betapb.ApplyPrivatecaBetaCaPoolRequest) (*betapb.PrivatecaBetaCaPool, error) {
	cl, err := createConfigCaPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCaPool(ctx, cl, request)
}

// DeleteCaPool handles the gRPC request by passing it to the underlying CaPool Delete() method.
func (s *CaPoolServer) DeletePrivatecaBetaCaPool(ctx context.Context, request *betapb.DeletePrivatecaBetaCaPoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCaPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCaPool(ctx, ProtoToCaPool(request.GetResource()))

}

// ListPrivatecaBetaCaPool handles the gRPC request by passing it to the underlying CaPoolList() method.
func (s *CaPoolServer) ListPrivatecaBetaCaPool(ctx context.Context, request *betapb.ListPrivatecaBetaCaPoolRequest) (*betapb.ListPrivatecaBetaCaPoolResponse, error) {
	cl, err := createConfigCaPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCaPool(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.PrivatecaBetaCaPool
	for _, r := range resources.Items {
		rp := CaPoolToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListPrivatecaBetaCaPoolResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCaPool(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
