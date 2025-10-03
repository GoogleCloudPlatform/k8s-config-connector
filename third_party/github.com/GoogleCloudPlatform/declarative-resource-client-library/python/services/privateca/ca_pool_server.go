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
	privatecapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/privateca/privateca_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/privateca"
)

// CaPoolServer implements the gRPC interface for CaPool.
type CaPoolServer struct{}

// ProtoToCaPoolTierEnum converts a CaPoolTierEnum enum from its proto representation.
func ProtoToPrivatecaCaPoolTierEnum(e privatecapb.PrivatecaCaPoolTierEnum) *privateca.CaPoolTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := privatecapb.PrivatecaCaPoolTierEnum_name[int32(e)]; ok {
		e := privateca.CaPoolTierEnum(n[len("PrivatecaCaPoolTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum converts a CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum enum from its proto representation.
func ProtoToPrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(e privatecapb.PrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum) *privateca.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := privatecapb.PrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum_name[int32(e)]; ok {
		e := privateca.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(n[len("PrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum converts a CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum enum from its proto representation.
func ProtoToPrivatecaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(e privatecapb.PrivatecaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum) *privateca.CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := privatecapb.PrivatecaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum_name[int32(e)]; ok {
		e := privateca.CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(n[len("PrivatecaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum"):])
		return &e
	}
	return nil
}

// ProtoToCaPoolIssuancePolicy converts a CaPoolIssuancePolicy object from its proto representation.
func ProtoToPrivatecaCaPoolIssuancePolicy(p *privatecapb.PrivatecaCaPoolIssuancePolicy) *privateca.CaPoolIssuancePolicy {
	if p == nil {
		return nil
	}
	obj := &privateca.CaPoolIssuancePolicy{
		MaximumLifetime:       dcl.StringOrNil(p.GetMaximumLifetime()),
		AllowedIssuanceModes:  ProtoToPrivatecaCaPoolIssuancePolicyAllowedIssuanceModes(p.GetAllowedIssuanceModes()),
		BaselineValues:        ProtoToPrivatecaCaPoolIssuancePolicyBaselineValues(p.GetBaselineValues()),
		IdentityConstraints:   ProtoToPrivatecaCaPoolIssuancePolicyIdentityConstraints(p.GetIdentityConstraints()),
		PassthroughExtensions: ProtoToPrivatecaCaPoolIssuancePolicyPassthroughExtensions(p.GetPassthroughExtensions()),
	}
	for _, r := range p.GetAllowedKeyTypes() {
		obj.AllowedKeyTypes = append(obj.AllowedKeyTypes, *ProtoToPrivatecaCaPoolIssuancePolicyAllowedKeyTypes(r))
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyAllowedKeyTypes converts a CaPoolIssuancePolicyAllowedKeyTypes object from its proto representation.
func ProtoToPrivatecaCaPoolIssuancePolicyAllowedKeyTypes(p *privatecapb.PrivatecaCaPoolIssuancePolicyAllowedKeyTypes) *privateca.CaPoolIssuancePolicyAllowedKeyTypes {
	if p == nil {
		return nil
	}
	obj := &privateca.CaPoolIssuancePolicyAllowedKeyTypes{
		Rsa:           ProtoToPrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsa(p.GetRsa()),
		EllipticCurve: ProtoToPrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve(p.GetEllipticCurve()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyAllowedKeyTypesRsa converts a CaPoolIssuancePolicyAllowedKeyTypesRsa object from its proto representation.
func ProtoToPrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsa(p *privatecapb.PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsa) *privateca.CaPoolIssuancePolicyAllowedKeyTypesRsa {
	if p == nil {
		return nil
	}
	obj := &privateca.CaPoolIssuancePolicyAllowedKeyTypesRsa{
		MinModulusSize: dcl.Int64OrNil(p.GetMinModulusSize()),
		MaxModulusSize: dcl.Int64OrNil(p.GetMaxModulusSize()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve converts a CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve object from its proto representation.
func ProtoToPrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve(p *privatecapb.PrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve) *privateca.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve {
	if p == nil {
		return nil
	}
	obj := &privateca.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve{
		SignatureAlgorithm: ProtoToPrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(p.GetSignatureAlgorithm()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyAllowedIssuanceModes converts a CaPoolIssuancePolicyAllowedIssuanceModes object from its proto representation.
func ProtoToPrivatecaCaPoolIssuancePolicyAllowedIssuanceModes(p *privatecapb.PrivatecaCaPoolIssuancePolicyAllowedIssuanceModes) *privateca.CaPoolIssuancePolicyAllowedIssuanceModes {
	if p == nil {
		return nil
	}
	obj := &privateca.CaPoolIssuancePolicyAllowedIssuanceModes{
		AllowCsrBasedIssuance:    dcl.Bool(p.GetAllowCsrBasedIssuance()),
		AllowConfigBasedIssuance: dcl.Bool(p.GetAllowConfigBasedIssuance()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValues converts a CaPoolIssuancePolicyBaselineValues object from its proto representation.
func ProtoToPrivatecaCaPoolIssuancePolicyBaselineValues(p *privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValues) *privateca.CaPoolIssuancePolicyBaselineValues {
	if p == nil {
		return nil
	}
	obj := &privateca.CaPoolIssuancePolicyBaselineValues{
		KeyUsage:  ProtoToPrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaCaPoolIssuancePolicyBaselineValuesCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaCaPoolIssuancePolicyBaselineValuesPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesKeyUsage converts a CaPoolIssuancePolicyBaselineValuesKeyUsage object from its proto representation.
func ProtoToPrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsage(p *privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsage) *privateca.CaPoolIssuancePolicyBaselineValuesKeyUsage {
	if p == nil {
		return nil
	}
	obj := &privateca.CaPoolIssuancePolicyBaselineValuesKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage converts a CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage object from its proto representation.
func ProtoToPrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage(p *privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage) *privateca.CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &privateca.CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage{
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
func ProtoToPrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage(p *privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage) *privateca.CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &privateca.CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage{
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
func ProtoToPrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages(p *privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages) *privateca.CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &privateca.CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesCaOptions converts a CaPoolIssuancePolicyBaselineValuesCaOptions object from its proto representation.
func ProtoToPrivatecaCaPoolIssuancePolicyBaselineValuesCaOptions(p *privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesCaOptions) *privateca.CaPoolIssuancePolicyBaselineValuesCaOptions {
	if p == nil {
		return nil
	}
	obj := &privateca.CaPoolIssuancePolicyBaselineValuesCaOptions{
		IsCa:                    dcl.Bool(p.GetIsCa()),
		MaxIssuerPathLength:     dcl.Int64OrNil(p.GetMaxIssuerPathLength()),
		ZeroMaxIssuerPathLength: dcl.Bool(p.GetZeroMaxIssuerPathLength()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesPolicyIds converts a CaPoolIssuancePolicyBaselineValuesPolicyIds object from its proto representation.
func ProtoToPrivatecaCaPoolIssuancePolicyBaselineValuesPolicyIds(p *privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesPolicyIds) *privateca.CaPoolIssuancePolicyBaselineValuesPolicyIds {
	if p == nil {
		return nil
	}
	obj := &privateca.CaPoolIssuancePolicyBaselineValuesPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesAdditionalExtensions converts a CaPoolIssuancePolicyBaselineValuesAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions(p *privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions) *privateca.CaPoolIssuancePolicyBaselineValuesAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &privateca.CaPoolIssuancePolicyBaselineValuesAdditionalExtensions{
		ObjectId: ProtoToPrivatecaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId converts a CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId object from its proto representation.
func ProtoToPrivatecaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId(p *privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId) *privateca.CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &privateca.CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyIdentityConstraints converts a CaPoolIssuancePolicyIdentityConstraints object from its proto representation.
func ProtoToPrivatecaCaPoolIssuancePolicyIdentityConstraints(p *privatecapb.PrivatecaCaPoolIssuancePolicyIdentityConstraints) *privateca.CaPoolIssuancePolicyIdentityConstraints {
	if p == nil {
		return nil
	}
	obj := &privateca.CaPoolIssuancePolicyIdentityConstraints{
		CelExpression:                   ProtoToPrivatecaCaPoolIssuancePolicyIdentityConstraintsCelExpression(p.GetCelExpression()),
		AllowSubjectPassthrough:         dcl.Bool(p.GetAllowSubjectPassthrough()),
		AllowSubjectAltNamesPassthrough: dcl.Bool(p.GetAllowSubjectAltNamesPassthrough()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyIdentityConstraintsCelExpression converts a CaPoolIssuancePolicyIdentityConstraintsCelExpression object from its proto representation.
func ProtoToPrivatecaCaPoolIssuancePolicyIdentityConstraintsCelExpression(p *privatecapb.PrivatecaCaPoolIssuancePolicyIdentityConstraintsCelExpression) *privateca.CaPoolIssuancePolicyIdentityConstraintsCelExpression {
	if p == nil {
		return nil
	}
	obj := &privateca.CaPoolIssuancePolicyIdentityConstraintsCelExpression{
		Expression:  dcl.StringOrNil(p.GetExpression()),
		Title:       dcl.StringOrNil(p.GetTitle()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyPassthroughExtensions converts a CaPoolIssuancePolicyPassthroughExtensions object from its proto representation.
func ProtoToPrivatecaCaPoolIssuancePolicyPassthroughExtensions(p *privatecapb.PrivatecaCaPoolIssuancePolicyPassthroughExtensions) *privateca.CaPoolIssuancePolicyPassthroughExtensions {
	if p == nil {
		return nil
	}
	obj := &privateca.CaPoolIssuancePolicyPassthroughExtensions{}
	for _, r := range p.GetKnownExtensions() {
		obj.KnownExtensions = append(obj.KnownExtensions, *ProtoToPrivatecaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(r))
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions converts a CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions(p *privatecapb.PrivatecaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions) *privateca.CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &privateca.CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCaPoolPublishingOptions converts a CaPoolPublishingOptions object from its proto representation.
func ProtoToPrivatecaCaPoolPublishingOptions(p *privatecapb.PrivatecaCaPoolPublishingOptions) *privateca.CaPoolPublishingOptions {
	if p == nil {
		return nil
	}
	obj := &privateca.CaPoolPublishingOptions{
		PublishCaCert: dcl.Bool(p.GetPublishCaCert()),
		PublishCrl:    dcl.Bool(p.GetPublishCrl()),
	}
	return obj
}

// ProtoToCaPool converts a CaPool resource from its proto representation.
func ProtoToCaPool(p *privatecapb.PrivatecaCaPool) *privateca.CaPool {
	obj := &privateca.CaPool{
		Name:              dcl.StringOrNil(p.GetName()),
		Tier:              ProtoToPrivatecaCaPoolTierEnum(p.GetTier()),
		IssuancePolicy:    ProtoToPrivatecaCaPoolIssuancePolicy(p.GetIssuancePolicy()),
		PublishingOptions: ProtoToPrivatecaCaPoolPublishingOptions(p.GetPublishingOptions()),
		Project:           dcl.StringOrNil(p.GetProject()),
		Location:          dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// CaPoolTierEnumToProto converts a CaPoolTierEnum enum to its proto representation.
func PrivatecaCaPoolTierEnumToProto(e *privateca.CaPoolTierEnum) privatecapb.PrivatecaCaPoolTierEnum {
	if e == nil {
		return privatecapb.PrivatecaCaPoolTierEnum(0)
	}
	if v, ok := privatecapb.PrivatecaCaPoolTierEnum_value["CaPoolTierEnum"+string(*e)]; ok {
		return privatecapb.PrivatecaCaPoolTierEnum(v)
	}
	return privatecapb.PrivatecaCaPoolTierEnum(0)
}

// CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnumToProto converts a CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum enum to its proto representation.
func PrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnumToProto(e *privateca.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum) privatecapb.PrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum {
	if e == nil {
		return privatecapb.PrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(0)
	}
	if v, ok := privatecapb.PrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum_value["CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum"+string(*e)]; ok {
		return privatecapb.PrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(v)
	}
	return privatecapb.PrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(0)
}

// CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnumToProto converts a CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum enum to its proto representation.
func PrivatecaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnumToProto(e *privateca.CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum) privatecapb.PrivatecaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum {
	if e == nil {
		return privatecapb.PrivatecaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(0)
	}
	if v, ok := privatecapb.PrivatecaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum_value["CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum"+string(*e)]; ok {
		return privatecapb.PrivatecaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(v)
	}
	return privatecapb.PrivatecaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(0)
}

// CaPoolIssuancePolicyToProto converts a CaPoolIssuancePolicy object to its proto representation.
func PrivatecaCaPoolIssuancePolicyToProto(o *privateca.CaPoolIssuancePolicy) *privatecapb.PrivatecaCaPoolIssuancePolicy {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCaPoolIssuancePolicy{}
	p.SetMaximumLifetime(dcl.ValueOrEmptyString(o.MaximumLifetime))
	p.SetAllowedIssuanceModes(PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesToProto(o.AllowedIssuanceModes))
	p.SetBaselineValues(PrivatecaCaPoolIssuancePolicyBaselineValuesToProto(o.BaselineValues))
	p.SetIdentityConstraints(PrivatecaCaPoolIssuancePolicyIdentityConstraintsToProto(o.IdentityConstraints))
	p.SetPassthroughExtensions(PrivatecaCaPoolIssuancePolicyPassthroughExtensionsToProto(o.PassthroughExtensions))
	sAllowedKeyTypes := make([]*privatecapb.PrivatecaCaPoolIssuancePolicyAllowedKeyTypes, len(o.AllowedKeyTypes))
	for i, r := range o.AllowedKeyTypes {
		sAllowedKeyTypes[i] = PrivatecaCaPoolIssuancePolicyAllowedKeyTypesToProto(&r)
	}
	p.SetAllowedKeyTypes(sAllowedKeyTypes)
	return p
}

// CaPoolIssuancePolicyAllowedKeyTypesToProto converts a CaPoolIssuancePolicyAllowedKeyTypes object to its proto representation.
func PrivatecaCaPoolIssuancePolicyAllowedKeyTypesToProto(o *privateca.CaPoolIssuancePolicyAllowedKeyTypes) *privatecapb.PrivatecaCaPoolIssuancePolicyAllowedKeyTypes {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCaPoolIssuancePolicyAllowedKeyTypes{}
	p.SetRsa(PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaToProto(o.Rsa))
	p.SetEllipticCurve(PrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveToProto(o.EllipticCurve))
	return p
}

// CaPoolIssuancePolicyAllowedKeyTypesRsaToProto converts a CaPoolIssuancePolicyAllowedKeyTypesRsa object to its proto representation.
func PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaToProto(o *privateca.CaPoolIssuancePolicyAllowedKeyTypesRsa) *privatecapb.PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsa {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsa{}
	p.SetMinModulusSize(dcl.ValueOrEmptyInt64(o.MinModulusSize))
	p.SetMaxModulusSize(dcl.ValueOrEmptyInt64(o.MaxModulusSize))
	return p
}

// CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveToProto converts a CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve object to its proto representation.
func PrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveToProto(o *privateca.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve) *privatecapb.PrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve{}
	p.SetSignatureAlgorithm(PrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnumToProto(o.SignatureAlgorithm))
	return p
}

// CaPoolIssuancePolicyAllowedIssuanceModesToProto converts a CaPoolIssuancePolicyAllowedIssuanceModes object to its proto representation.
func PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesToProto(o *privateca.CaPoolIssuancePolicyAllowedIssuanceModes) *privatecapb.PrivatecaCaPoolIssuancePolicyAllowedIssuanceModes {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCaPoolIssuancePolicyAllowedIssuanceModes{}
	p.SetAllowCsrBasedIssuance(dcl.ValueOrEmptyBool(o.AllowCsrBasedIssuance))
	p.SetAllowConfigBasedIssuance(dcl.ValueOrEmptyBool(o.AllowConfigBasedIssuance))
	return p
}

// CaPoolIssuancePolicyBaselineValuesToProto converts a CaPoolIssuancePolicyBaselineValues object to its proto representation.
func PrivatecaCaPoolIssuancePolicyBaselineValuesToProto(o *privateca.CaPoolIssuancePolicyBaselineValues) *privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValues {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValues{}
	p.SetKeyUsage(PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageToProto(o.KeyUsage))
	p.SetCaOptions(PrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsToProto(o.CaOptions))
	sPolicyIds := make([]*privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesPolicyIds, len(o.PolicyIds))
	for i, r := range o.PolicyIds {
		sPolicyIds[i] = PrivatecaCaPoolIssuancePolicyBaselineValuesPolicyIdsToProto(&r)
	}
	p.SetPolicyIds(sPolicyIds)
	sAiaOcspServers := make([]string, len(o.AiaOcspServers))
	for i, r := range o.AiaOcspServers {
		sAiaOcspServers[i] = r
	}
	p.SetAiaOcspServers(sAiaOcspServers)
	sAdditionalExtensions := make([]*privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CaPoolIssuancePolicyBaselineValuesKeyUsageToProto converts a CaPoolIssuancePolicyBaselineValuesKeyUsage object to its proto representation.
func PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageToProto(o *privateca.CaPoolIssuancePolicyBaselineValuesKeyUsage) *privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsage {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsage{}
	p.SetBaseKeyUsage(PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage))
	p.SetExtendedKeyUsage(PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage))
	sUnknownExtendedKeyUsages := make([]*privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages, len(o.UnknownExtendedKeyUsages))
	for i, r := range o.UnknownExtendedKeyUsages {
		sUnknownExtendedKeyUsages[i] = PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesToProto(&r)
	}
	p.SetUnknownExtendedKeyUsages(sUnknownExtendedKeyUsages)
	return p
}

// CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsageToProto converts a CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage object to its proto representation.
func PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsageToProto(o *privateca.CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage) *privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage{}
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
func PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageToProto(o *privateca.CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage) *privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage{}
	p.SetServerAuth(dcl.ValueOrEmptyBool(o.ServerAuth))
	p.SetClientAuth(dcl.ValueOrEmptyBool(o.ClientAuth))
	p.SetCodeSigning(dcl.ValueOrEmptyBool(o.CodeSigning))
	p.SetEmailProtection(dcl.ValueOrEmptyBool(o.EmailProtection))
	p.SetTimeStamping(dcl.ValueOrEmptyBool(o.TimeStamping))
	p.SetOcspSigning(dcl.ValueOrEmptyBool(o.OcspSigning))
	return p
}

// CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesToProto converts a CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages object to its proto representation.
func PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesToProto(o *privateca.CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages) *privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CaPoolIssuancePolicyBaselineValuesCaOptionsToProto converts a CaPoolIssuancePolicyBaselineValuesCaOptions object to its proto representation.
func PrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsToProto(o *privateca.CaPoolIssuancePolicyBaselineValuesCaOptions) *privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesCaOptions {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesCaOptions{}
	p.SetIsCa(dcl.ValueOrEmptyBool(o.IsCa))
	p.SetMaxIssuerPathLength(dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength))
	p.SetZeroMaxIssuerPathLength(dcl.ValueOrEmptyBool(o.ZeroMaxIssuerPathLength))
	return p
}

// CaPoolIssuancePolicyBaselineValuesPolicyIdsToProto converts a CaPoolIssuancePolicyBaselineValuesPolicyIds object to its proto representation.
func PrivatecaCaPoolIssuancePolicyBaselineValuesPolicyIdsToProto(o *privateca.CaPoolIssuancePolicyBaselineValuesPolicyIds) *privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesPolicyIds {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesPolicyIds{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsToProto converts a CaPoolIssuancePolicyBaselineValuesAdditionalExtensions object to its proto representation.
func PrivatecaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsToProto(o *privateca.CaPoolIssuancePolicyBaselineValuesAdditionalExtensions) *privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions{}
	p.SetObjectId(PrivatecaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectIdToProto converts a CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId object to its proto representation.
func PrivatecaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectIdToProto(o *privateca.CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId) *privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CaPoolIssuancePolicyIdentityConstraintsToProto converts a CaPoolIssuancePolicyIdentityConstraints object to its proto representation.
func PrivatecaCaPoolIssuancePolicyIdentityConstraintsToProto(o *privateca.CaPoolIssuancePolicyIdentityConstraints) *privatecapb.PrivatecaCaPoolIssuancePolicyIdentityConstraints {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCaPoolIssuancePolicyIdentityConstraints{}
	p.SetCelExpression(PrivatecaCaPoolIssuancePolicyIdentityConstraintsCelExpressionToProto(o.CelExpression))
	p.SetAllowSubjectPassthrough(dcl.ValueOrEmptyBool(o.AllowSubjectPassthrough))
	p.SetAllowSubjectAltNamesPassthrough(dcl.ValueOrEmptyBool(o.AllowSubjectAltNamesPassthrough))
	return p
}

// CaPoolIssuancePolicyIdentityConstraintsCelExpressionToProto converts a CaPoolIssuancePolicyIdentityConstraintsCelExpression object to its proto representation.
func PrivatecaCaPoolIssuancePolicyIdentityConstraintsCelExpressionToProto(o *privateca.CaPoolIssuancePolicyIdentityConstraintsCelExpression) *privatecapb.PrivatecaCaPoolIssuancePolicyIdentityConstraintsCelExpression {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCaPoolIssuancePolicyIdentityConstraintsCelExpression{}
	p.SetExpression(dcl.ValueOrEmptyString(o.Expression))
	p.SetTitle(dcl.ValueOrEmptyString(o.Title))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	return p
}

// CaPoolIssuancePolicyPassthroughExtensionsToProto converts a CaPoolIssuancePolicyPassthroughExtensions object to its proto representation.
func PrivatecaCaPoolIssuancePolicyPassthroughExtensionsToProto(o *privateca.CaPoolIssuancePolicyPassthroughExtensions) *privatecapb.PrivatecaCaPoolIssuancePolicyPassthroughExtensions {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCaPoolIssuancePolicyPassthroughExtensions{}
	sKnownExtensions := make([]privatecapb.PrivatecaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum, len(o.KnownExtensions))
	for i, r := range o.KnownExtensions {
		sKnownExtensions[i] = privatecapb.PrivatecaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(privatecapb.PrivatecaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum_value[string(r)])
	}
	p.SetKnownExtensions(sKnownExtensions)
	sAdditionalExtensions := make([]*privatecapb.PrivatecaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensionsToProto converts a CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions object to its proto representation.
func PrivatecaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensionsToProto(o *privateca.CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions) *privatecapb.PrivatecaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CaPoolPublishingOptionsToProto converts a CaPoolPublishingOptions object to its proto representation.
func PrivatecaCaPoolPublishingOptionsToProto(o *privateca.CaPoolPublishingOptions) *privatecapb.PrivatecaCaPoolPublishingOptions {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCaPoolPublishingOptions{}
	p.SetPublishCaCert(dcl.ValueOrEmptyBool(o.PublishCaCert))
	p.SetPublishCrl(dcl.ValueOrEmptyBool(o.PublishCrl))
	return p
}

// CaPoolToProto converts a CaPool resource to its proto representation.
func CaPoolToProto(resource *privateca.CaPool) *privatecapb.PrivatecaCaPool {
	p := &privatecapb.PrivatecaCaPool{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetTier(PrivatecaCaPoolTierEnumToProto(resource.Tier))
	p.SetIssuancePolicy(PrivatecaCaPoolIssuancePolicyToProto(resource.IssuancePolicy))
	p.SetPublishingOptions(PrivatecaCaPoolPublishingOptionsToProto(resource.PublishingOptions))
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
func (s *CaPoolServer) applyCaPool(ctx context.Context, c *privateca.Client, request *privatecapb.ApplyPrivatecaCaPoolRequest) (*privatecapb.PrivatecaCaPool, error) {
	p := ProtoToCaPool(request.GetResource())
	res, err := c.ApplyCaPool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CaPoolToProto(res)
	return r, nil
}

// applyPrivatecaCaPool handles the gRPC request by passing it to the underlying CaPool Apply() method.
func (s *CaPoolServer) ApplyPrivatecaCaPool(ctx context.Context, request *privatecapb.ApplyPrivatecaCaPoolRequest) (*privatecapb.PrivatecaCaPool, error) {
	cl, err := createConfigCaPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCaPool(ctx, cl, request)
}

// DeleteCaPool handles the gRPC request by passing it to the underlying CaPool Delete() method.
func (s *CaPoolServer) DeletePrivatecaCaPool(ctx context.Context, request *privatecapb.DeletePrivatecaCaPoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCaPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCaPool(ctx, ProtoToCaPool(request.GetResource()))

}

// ListPrivatecaCaPool handles the gRPC request by passing it to the underlying CaPoolList() method.
func (s *CaPoolServer) ListPrivatecaCaPool(ctx context.Context, request *privatecapb.ListPrivatecaCaPoolRequest) (*privatecapb.ListPrivatecaCaPoolResponse, error) {
	cl, err := createConfigCaPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCaPool(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*privatecapb.PrivatecaCaPool
	for _, r := range resources.Items {
		rp := CaPoolToProto(r)
		protos = append(protos, rp)
	}
	p := &privatecapb.ListPrivatecaCaPoolResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCaPool(ctx context.Context, service_account_file string) (*privateca.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return privateca.NewClient(conf), nil
}
