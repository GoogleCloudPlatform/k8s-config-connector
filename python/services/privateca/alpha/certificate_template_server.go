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

// CertificateTemplateServer implements the gRPC interface for CertificateTemplate.
type CertificateTemplateServer struct{}

// ProtoToCertificateTemplatePassthroughExtensionsKnownExtensionsEnum converts a CertificateTemplatePassthroughExtensionsKnownExtensionsEnum enum from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(e alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum) *alpha.CertificateTemplatePassthroughExtensionsKnownExtensionsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum_name[int32(e)]; ok {
		e := alpha.CertificateTemplatePassthroughExtensionsKnownExtensionsEnum(n[len("PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateTemplatePredefinedValues converts a CertificateTemplatePredefinedValues object from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplatePredefinedValues(p *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValues) *alpha.CertificateTemplatePredefinedValues {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplatePredefinedValues{
		KeyUsage:  ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesKeyUsage converts a CertificateTemplatePredefinedValuesKeyUsage object from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsage(p *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsage) *alpha.CertificateTemplatePredefinedValuesKeyUsage {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplatePredefinedValuesKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage converts a CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage object from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage(p *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage) *alpha.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage{
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

// ProtoToCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage converts a CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage object from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage(p *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage) *alpha.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage{
		ServerAuth:      dcl.Bool(p.GetServerAuth()),
		ClientAuth:      dcl.Bool(p.GetClientAuth()),
		CodeSigning:     dcl.Bool(p.GetCodeSigning()),
		EmailProtection: dcl.Bool(p.GetEmailProtection()),
		TimeStamping:    dcl.Bool(p.GetTimeStamping()),
		OcspSigning:     dcl.Bool(p.GetOcspSigning()),
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages converts a CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages object from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages(p *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages) *alpha.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesCaOptions converts a CertificateTemplatePredefinedValuesCaOptions object from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesCaOptions(p *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesCaOptions) *alpha.CertificateTemplatePredefinedValuesCaOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplatePredefinedValuesCaOptions{
		IsCa:                dcl.Bool(p.GetIsCa()),
		MaxIssuerPathLength: dcl.Int64OrNil(p.GetMaxIssuerPathLength()),
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesPolicyIds converts a CertificateTemplatePredefinedValuesPolicyIds object from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesPolicyIds(p *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesPolicyIds) *alpha.CertificateTemplatePredefinedValuesPolicyIds {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplatePredefinedValuesPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesAdditionalExtensions converts a CertificateTemplatePredefinedValuesAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensions(p *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensions) *alpha.CertificateTemplatePredefinedValuesAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplatePredefinedValuesAdditionalExtensions{
		ObjectId: ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId converts a CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId object from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId(p *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId) *alpha.CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplateIdentityConstraints converts a CertificateTemplateIdentityConstraints object from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplateIdentityConstraints(p *alphapb.PrivatecaAlphaCertificateTemplateIdentityConstraints) *alpha.CertificateTemplateIdentityConstraints {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplateIdentityConstraints{
		CelExpression:                   ProtoToPrivatecaAlphaCertificateTemplateIdentityConstraintsCelExpression(p.GetCelExpression()),
		AllowSubjectPassthrough:         dcl.Bool(p.GetAllowSubjectPassthrough()),
		AllowSubjectAltNamesPassthrough: dcl.Bool(p.GetAllowSubjectAltNamesPassthrough()),
	}
	return obj
}

// ProtoToCertificateTemplateIdentityConstraintsCelExpression converts a CertificateTemplateIdentityConstraintsCelExpression object from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplateIdentityConstraintsCelExpression(p *alphapb.PrivatecaAlphaCertificateTemplateIdentityConstraintsCelExpression) *alpha.CertificateTemplateIdentityConstraintsCelExpression {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplateIdentityConstraintsCelExpression{
		Expression:  dcl.StringOrNil(p.GetExpression()),
		Title:       dcl.StringOrNil(p.GetTitle()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ProtoToCertificateTemplatePassthroughExtensions converts a CertificateTemplatePassthroughExtensions object from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplatePassthroughExtensions(p *alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensions) *alpha.CertificateTemplatePassthroughExtensions {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplatePassthroughExtensions{}
	for _, r := range p.GetKnownExtensions() {
		obj.KnownExtensions = append(obj.KnownExtensions, *ProtoToPrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(r))
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaAlphaCertificateTemplatePassthroughExtensionsAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateTemplatePassthroughExtensionsAdditionalExtensions converts a CertificateTemplatePassthroughExtensionsAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplatePassthroughExtensionsAdditionalExtensions(p *alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsAdditionalExtensions) *alpha.CertificateTemplatePassthroughExtensionsAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplatePassthroughExtensionsAdditionalExtensions{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplate converts a CertificateTemplate resource from its proto representation.
func ProtoToCertificateTemplate(p *alphapb.PrivatecaAlphaCertificateTemplate) *alpha.CertificateTemplate {
	obj := &alpha.CertificateTemplate{
		Name:                  dcl.StringOrNil(p.GetName()),
		PredefinedValues:      ProtoToPrivatecaAlphaCertificateTemplatePredefinedValues(p.GetPredefinedValues()),
		IdentityConstraints:   ProtoToPrivatecaAlphaCertificateTemplateIdentityConstraints(p.GetIdentityConstraints()),
		PassthroughExtensions: ProtoToPrivatecaAlphaCertificateTemplatePassthroughExtensions(p.GetPassthroughExtensions()),
		Description:           dcl.StringOrNil(p.GetDescription()),
		CreateTime:            dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:            dcl.StringOrNil(p.GetUpdateTime()),
		Project:               dcl.StringOrNil(p.GetProject()),
		Location:              dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// CertificateTemplatePassthroughExtensionsKnownExtensionsEnumToProto converts a CertificateTemplatePassthroughExtensionsKnownExtensionsEnum enum to its proto representation.
func PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnumToProto(e *alpha.CertificateTemplatePassthroughExtensionsKnownExtensionsEnum) alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum {
	if e == nil {
		return alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(0)
	}
	if v, ok := alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum_value["CertificateTemplatePassthroughExtensionsKnownExtensionsEnum"+string(*e)]; ok {
		return alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(v)
	}
	return alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(0)
}

// CertificateTemplatePredefinedValuesToProto converts a CertificateTemplatePredefinedValues object to its proto representation.
func PrivatecaAlphaCertificateTemplatePredefinedValuesToProto(o *alpha.CertificateTemplatePredefinedValues) *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValues {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplatePredefinedValues{}
	p.SetKeyUsage(PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageToProto(o.KeyUsage))
	p.SetCaOptions(PrivatecaAlphaCertificateTemplatePredefinedValuesCaOptionsToProto(o.CaOptions))
	sPolicyIds := make([]*alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesPolicyIds, len(o.PolicyIds))
	for i, r := range o.PolicyIds {
		sPolicyIds[i] = PrivatecaAlphaCertificateTemplatePredefinedValuesPolicyIdsToProto(&r)
	}
	p.SetPolicyIds(sPolicyIds)
	sAiaOcspServers := make([]string, len(o.AiaOcspServers))
	for i, r := range o.AiaOcspServers {
		sAiaOcspServers[i] = r
	}
	p.SetAiaOcspServers(sAiaOcspServers)
	sAdditionalExtensions := make([]*alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CertificateTemplatePredefinedValuesKeyUsageToProto converts a CertificateTemplatePredefinedValuesKeyUsage object to its proto representation.
func PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageToProto(o *alpha.CertificateTemplatePredefinedValuesKeyUsage) *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsage {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsage{}
	p.SetBaseKeyUsage(PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage))
	p.SetExtendedKeyUsage(PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage))
	sUnknownExtendedKeyUsages := make([]*alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages, len(o.UnknownExtendedKeyUsages))
	for i, r := range o.UnknownExtendedKeyUsages {
		sUnknownExtendedKeyUsages[i] = PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesToProto(&r)
	}
	p.SetUnknownExtendedKeyUsages(sUnknownExtendedKeyUsages)
	return p
}

// CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageToProto converts a CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage object to its proto representation.
func PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageToProto(o *alpha.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage) *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage{}
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

// CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageToProto converts a CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage object to its proto representation.
func PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageToProto(o *alpha.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage) *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage{}
	p.SetServerAuth(dcl.ValueOrEmptyBool(o.ServerAuth))
	p.SetClientAuth(dcl.ValueOrEmptyBool(o.ClientAuth))
	p.SetCodeSigning(dcl.ValueOrEmptyBool(o.CodeSigning))
	p.SetEmailProtection(dcl.ValueOrEmptyBool(o.EmailProtection))
	p.SetTimeStamping(dcl.ValueOrEmptyBool(o.TimeStamping))
	p.SetOcspSigning(dcl.ValueOrEmptyBool(o.OcspSigning))
	return p
}

// CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesToProto converts a CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages object to its proto representation.
func PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesToProto(o *alpha.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages) *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateTemplatePredefinedValuesCaOptionsToProto converts a CertificateTemplatePredefinedValuesCaOptions object to its proto representation.
func PrivatecaAlphaCertificateTemplatePredefinedValuesCaOptionsToProto(o *alpha.CertificateTemplatePredefinedValuesCaOptions) *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesCaOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesCaOptions{}
	p.SetIsCa(dcl.ValueOrEmptyBool(o.IsCa))
	p.SetMaxIssuerPathLength(dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength))
	return p
}

// CertificateTemplatePredefinedValuesPolicyIdsToProto converts a CertificateTemplatePredefinedValuesPolicyIds object to its proto representation.
func PrivatecaAlphaCertificateTemplatePredefinedValuesPolicyIdsToProto(o *alpha.CertificateTemplatePredefinedValuesPolicyIds) *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesPolicyIds {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesPolicyIds{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateTemplatePredefinedValuesAdditionalExtensionsToProto converts a CertificateTemplatePredefinedValuesAdditionalExtensions object to its proto representation.
func PrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensionsToProto(o *alpha.CertificateTemplatePredefinedValuesAdditionalExtensions) *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensions{}
	p.SetObjectId(PrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdToProto converts a CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId object to its proto representation.
func PrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdToProto(o *alpha.CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId) *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateTemplateIdentityConstraintsToProto converts a CertificateTemplateIdentityConstraints object to its proto representation.
func PrivatecaAlphaCertificateTemplateIdentityConstraintsToProto(o *alpha.CertificateTemplateIdentityConstraints) *alphapb.PrivatecaAlphaCertificateTemplateIdentityConstraints {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplateIdentityConstraints{}
	p.SetCelExpression(PrivatecaAlphaCertificateTemplateIdentityConstraintsCelExpressionToProto(o.CelExpression))
	p.SetAllowSubjectPassthrough(dcl.ValueOrEmptyBool(o.AllowSubjectPassthrough))
	p.SetAllowSubjectAltNamesPassthrough(dcl.ValueOrEmptyBool(o.AllowSubjectAltNamesPassthrough))
	return p
}

// CertificateTemplateIdentityConstraintsCelExpressionToProto converts a CertificateTemplateIdentityConstraintsCelExpression object to its proto representation.
func PrivatecaAlphaCertificateTemplateIdentityConstraintsCelExpressionToProto(o *alpha.CertificateTemplateIdentityConstraintsCelExpression) *alphapb.PrivatecaAlphaCertificateTemplateIdentityConstraintsCelExpression {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplateIdentityConstraintsCelExpression{}
	p.SetExpression(dcl.ValueOrEmptyString(o.Expression))
	p.SetTitle(dcl.ValueOrEmptyString(o.Title))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	return p
}

// CertificateTemplatePassthroughExtensionsToProto converts a CertificateTemplatePassthroughExtensions object to its proto representation.
func PrivatecaAlphaCertificateTemplatePassthroughExtensionsToProto(o *alpha.CertificateTemplatePassthroughExtensions) *alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensions {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensions{}
	sKnownExtensions := make([]alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum, len(o.KnownExtensions))
	for i, r := range o.KnownExtensions {
		sKnownExtensions[i] = alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum_value[string(r)])
	}
	p.SetKnownExtensions(sKnownExtensions)
	sAdditionalExtensions := make([]*alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaAlphaCertificateTemplatePassthroughExtensionsAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CertificateTemplatePassthroughExtensionsAdditionalExtensionsToProto converts a CertificateTemplatePassthroughExtensionsAdditionalExtensions object to its proto representation.
func PrivatecaAlphaCertificateTemplatePassthroughExtensionsAdditionalExtensionsToProto(o *alpha.CertificateTemplatePassthroughExtensionsAdditionalExtensions) *alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsAdditionalExtensions{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateTemplateToProto converts a CertificateTemplate resource to its proto representation.
func CertificateTemplateToProto(resource *alpha.CertificateTemplate) *alphapb.PrivatecaAlphaCertificateTemplate {
	p := &alphapb.PrivatecaAlphaCertificateTemplate{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetPredefinedValues(PrivatecaAlphaCertificateTemplatePredefinedValuesToProto(resource.PredefinedValues))
	p.SetIdentityConstraints(PrivatecaAlphaCertificateTemplateIdentityConstraintsToProto(resource.IdentityConstraints))
	p.SetPassthroughExtensions(PrivatecaAlphaCertificateTemplatePassthroughExtensionsToProto(resource.PassthroughExtensions))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyCertificateTemplate handles the gRPC request by passing it to the underlying CertificateTemplate Apply() method.
func (s *CertificateTemplateServer) applyCertificateTemplate(ctx context.Context, c *alpha.Client, request *alphapb.ApplyPrivatecaAlphaCertificateTemplateRequest) (*alphapb.PrivatecaAlphaCertificateTemplate, error) {
	p := ProtoToCertificateTemplate(request.GetResource())
	res, err := c.ApplyCertificateTemplate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CertificateTemplateToProto(res)
	return r, nil
}

// applyPrivatecaAlphaCertificateTemplate handles the gRPC request by passing it to the underlying CertificateTemplate Apply() method.
func (s *CertificateTemplateServer) ApplyPrivatecaAlphaCertificateTemplate(ctx context.Context, request *alphapb.ApplyPrivatecaAlphaCertificateTemplateRequest) (*alphapb.PrivatecaAlphaCertificateTemplate, error) {
	cl, err := createConfigCertificateTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCertificateTemplate(ctx, cl, request)
}

// DeleteCertificateTemplate handles the gRPC request by passing it to the underlying CertificateTemplate Delete() method.
func (s *CertificateTemplateServer) DeletePrivatecaAlphaCertificateTemplate(ctx context.Context, request *alphapb.DeletePrivatecaAlphaCertificateTemplateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCertificateTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCertificateTemplate(ctx, ProtoToCertificateTemplate(request.GetResource()))

}

// ListPrivatecaAlphaCertificateTemplate handles the gRPC request by passing it to the underlying CertificateTemplateList() method.
func (s *CertificateTemplateServer) ListPrivatecaAlphaCertificateTemplate(ctx context.Context, request *alphapb.ListPrivatecaAlphaCertificateTemplateRequest) (*alphapb.ListPrivatecaAlphaCertificateTemplateResponse, error) {
	cl, err := createConfigCertificateTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCertificateTemplate(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.PrivatecaAlphaCertificateTemplate
	for _, r := range resources.Items {
		rp := CertificateTemplateToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListPrivatecaAlphaCertificateTemplateResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCertificateTemplate(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
