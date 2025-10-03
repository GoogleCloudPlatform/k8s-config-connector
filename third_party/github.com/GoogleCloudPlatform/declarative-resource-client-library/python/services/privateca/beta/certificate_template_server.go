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

// CertificateTemplateServer implements the gRPC interface for CertificateTemplate.
type CertificateTemplateServer struct{}

// ProtoToCertificateTemplatePassthroughExtensionsKnownExtensionsEnum converts a CertificateTemplatePassthroughExtensionsKnownExtensionsEnum enum from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(e betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum) *beta.CertificateTemplatePassthroughExtensionsKnownExtensionsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum_name[int32(e)]; ok {
		e := beta.CertificateTemplatePassthroughExtensionsKnownExtensionsEnum(n[len("PrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateTemplatePredefinedValues converts a CertificateTemplatePredefinedValues object from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplatePredefinedValues(p *betapb.PrivatecaBetaCertificateTemplatePredefinedValues) *beta.CertificateTemplatePredefinedValues {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplatePredefinedValues{
		KeyUsage:  ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesKeyUsage converts a CertificateTemplatePredefinedValuesKeyUsage object from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesKeyUsage(p *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsage) *beta.CertificateTemplatePredefinedValuesKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplatePredefinedValuesKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage converts a CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage object from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage(p *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage) *beta.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage{
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
func ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage(p *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage) *beta.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage{
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
func ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages(p *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages) *beta.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesCaOptions converts a CertificateTemplatePredefinedValuesCaOptions object from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesCaOptions(p *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesCaOptions) *beta.CertificateTemplatePredefinedValuesCaOptions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplatePredefinedValuesCaOptions{
		IsCa:                dcl.Bool(p.GetIsCa()),
		MaxIssuerPathLength: dcl.Int64OrNil(p.GetMaxIssuerPathLength()),
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesPolicyIds converts a CertificateTemplatePredefinedValuesPolicyIds object from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesPolicyIds(p *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesPolicyIds) *beta.CertificateTemplatePredefinedValuesPolicyIds {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplatePredefinedValuesPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesAdditionalExtensions converts a CertificateTemplatePredefinedValuesAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensions(p *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensions) *beta.CertificateTemplatePredefinedValuesAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplatePredefinedValuesAdditionalExtensions{
		ObjectId: ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId converts a CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId object from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId(p *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId) *beta.CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplateIdentityConstraints converts a CertificateTemplateIdentityConstraints object from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplateIdentityConstraints(p *betapb.PrivatecaBetaCertificateTemplateIdentityConstraints) *beta.CertificateTemplateIdentityConstraints {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplateIdentityConstraints{
		CelExpression:                   ProtoToPrivatecaBetaCertificateTemplateIdentityConstraintsCelExpression(p.GetCelExpression()),
		AllowSubjectPassthrough:         dcl.Bool(p.GetAllowSubjectPassthrough()),
		AllowSubjectAltNamesPassthrough: dcl.Bool(p.GetAllowSubjectAltNamesPassthrough()),
	}
	return obj
}

// ProtoToCertificateTemplateIdentityConstraintsCelExpression converts a CertificateTemplateIdentityConstraintsCelExpression object from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplateIdentityConstraintsCelExpression(p *betapb.PrivatecaBetaCertificateTemplateIdentityConstraintsCelExpression) *beta.CertificateTemplateIdentityConstraintsCelExpression {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplateIdentityConstraintsCelExpression{
		Expression:  dcl.StringOrNil(p.GetExpression()),
		Title:       dcl.StringOrNil(p.GetTitle()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ProtoToCertificateTemplatePassthroughExtensions converts a CertificateTemplatePassthroughExtensions object from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplatePassthroughExtensions(p *betapb.PrivatecaBetaCertificateTemplatePassthroughExtensions) *beta.CertificateTemplatePassthroughExtensions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplatePassthroughExtensions{}
	for _, r := range p.GetKnownExtensions() {
		obj.KnownExtensions = append(obj.KnownExtensions, *ProtoToPrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(r))
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaBetaCertificateTemplatePassthroughExtensionsAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateTemplatePassthroughExtensionsAdditionalExtensions converts a CertificateTemplatePassthroughExtensionsAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplatePassthroughExtensionsAdditionalExtensions(p *betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsAdditionalExtensions) *beta.CertificateTemplatePassthroughExtensionsAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplatePassthroughExtensionsAdditionalExtensions{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplate converts a CertificateTemplate resource from its proto representation.
func ProtoToCertificateTemplate(p *betapb.PrivatecaBetaCertificateTemplate) *beta.CertificateTemplate {
	obj := &beta.CertificateTemplate{
		Name:                  dcl.StringOrNil(p.GetName()),
		PredefinedValues:      ProtoToPrivatecaBetaCertificateTemplatePredefinedValues(p.GetPredefinedValues()),
		IdentityConstraints:   ProtoToPrivatecaBetaCertificateTemplateIdentityConstraints(p.GetIdentityConstraints()),
		PassthroughExtensions: ProtoToPrivatecaBetaCertificateTemplatePassthroughExtensions(p.GetPassthroughExtensions()),
		Description:           dcl.StringOrNil(p.GetDescription()),
		CreateTime:            dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:            dcl.StringOrNil(p.GetUpdateTime()),
		Project:               dcl.StringOrNil(p.GetProject()),
		Location:              dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// CertificateTemplatePassthroughExtensionsKnownExtensionsEnumToProto converts a CertificateTemplatePassthroughExtensionsKnownExtensionsEnum enum to its proto representation.
func PrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnumToProto(e *beta.CertificateTemplatePassthroughExtensionsKnownExtensionsEnum) betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum {
	if e == nil {
		return betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum_value["CertificateTemplatePassthroughExtensionsKnownExtensionsEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(v)
	}
	return betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(0)
}

// CertificateTemplatePredefinedValuesToProto converts a CertificateTemplatePredefinedValues object to its proto representation.
func PrivatecaBetaCertificateTemplatePredefinedValuesToProto(o *beta.CertificateTemplatePredefinedValues) *betapb.PrivatecaBetaCertificateTemplatePredefinedValues {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplatePredefinedValues{}
	p.SetKeyUsage(PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageToProto(o.KeyUsage))
	p.SetCaOptions(PrivatecaBetaCertificateTemplatePredefinedValuesCaOptionsToProto(o.CaOptions))
	sPolicyIds := make([]*betapb.PrivatecaBetaCertificateTemplatePredefinedValuesPolicyIds, len(o.PolicyIds))
	for i, r := range o.PolicyIds {
		sPolicyIds[i] = PrivatecaBetaCertificateTemplatePredefinedValuesPolicyIdsToProto(&r)
	}
	p.SetPolicyIds(sPolicyIds)
	sAiaOcspServers := make([]string, len(o.AiaOcspServers))
	for i, r := range o.AiaOcspServers {
		sAiaOcspServers[i] = r
	}
	p.SetAiaOcspServers(sAiaOcspServers)
	sAdditionalExtensions := make([]*betapb.PrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CertificateTemplatePredefinedValuesKeyUsageToProto converts a CertificateTemplatePredefinedValuesKeyUsage object to its proto representation.
func PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageToProto(o *beta.CertificateTemplatePredefinedValuesKeyUsage) *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsage{}
	p.SetBaseKeyUsage(PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage))
	p.SetExtendedKeyUsage(PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage))
	sUnknownExtendedKeyUsages := make([]*betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages, len(o.UnknownExtendedKeyUsages))
	for i, r := range o.UnknownExtendedKeyUsages {
		sUnknownExtendedKeyUsages[i] = PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesToProto(&r)
	}
	p.SetUnknownExtendedKeyUsages(sUnknownExtendedKeyUsages)
	return p
}

// CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageToProto converts a CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage object to its proto representation.
func PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageToProto(o *beta.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage) *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage{}
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
func PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageToProto(o *beta.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage) *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage{}
	p.SetServerAuth(dcl.ValueOrEmptyBool(o.ServerAuth))
	p.SetClientAuth(dcl.ValueOrEmptyBool(o.ClientAuth))
	p.SetCodeSigning(dcl.ValueOrEmptyBool(o.CodeSigning))
	p.SetEmailProtection(dcl.ValueOrEmptyBool(o.EmailProtection))
	p.SetTimeStamping(dcl.ValueOrEmptyBool(o.TimeStamping))
	p.SetOcspSigning(dcl.ValueOrEmptyBool(o.OcspSigning))
	return p
}

// CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesToProto converts a CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages object to its proto representation.
func PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesToProto(o *beta.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages) *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateTemplatePredefinedValuesCaOptionsToProto converts a CertificateTemplatePredefinedValuesCaOptions object to its proto representation.
func PrivatecaBetaCertificateTemplatePredefinedValuesCaOptionsToProto(o *beta.CertificateTemplatePredefinedValuesCaOptions) *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesCaOptions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplatePredefinedValuesCaOptions{}
	p.SetIsCa(dcl.ValueOrEmptyBool(o.IsCa))
	p.SetMaxIssuerPathLength(dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength))
	return p
}

// CertificateTemplatePredefinedValuesPolicyIdsToProto converts a CertificateTemplatePredefinedValuesPolicyIds object to its proto representation.
func PrivatecaBetaCertificateTemplatePredefinedValuesPolicyIdsToProto(o *beta.CertificateTemplatePredefinedValuesPolicyIds) *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesPolicyIds {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplatePredefinedValuesPolicyIds{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateTemplatePredefinedValuesAdditionalExtensionsToProto converts a CertificateTemplatePredefinedValuesAdditionalExtensions object to its proto representation.
func PrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensionsToProto(o *beta.CertificateTemplatePredefinedValuesAdditionalExtensions) *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensions{}
	p.SetObjectId(PrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdToProto converts a CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId object to its proto representation.
func PrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdToProto(o *beta.CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId) *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateTemplateIdentityConstraintsToProto converts a CertificateTemplateIdentityConstraints object to its proto representation.
func PrivatecaBetaCertificateTemplateIdentityConstraintsToProto(o *beta.CertificateTemplateIdentityConstraints) *betapb.PrivatecaBetaCertificateTemplateIdentityConstraints {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplateIdentityConstraints{}
	p.SetCelExpression(PrivatecaBetaCertificateTemplateIdentityConstraintsCelExpressionToProto(o.CelExpression))
	p.SetAllowSubjectPassthrough(dcl.ValueOrEmptyBool(o.AllowSubjectPassthrough))
	p.SetAllowSubjectAltNamesPassthrough(dcl.ValueOrEmptyBool(o.AllowSubjectAltNamesPassthrough))
	return p
}

// CertificateTemplateIdentityConstraintsCelExpressionToProto converts a CertificateTemplateIdentityConstraintsCelExpression object to its proto representation.
func PrivatecaBetaCertificateTemplateIdentityConstraintsCelExpressionToProto(o *beta.CertificateTemplateIdentityConstraintsCelExpression) *betapb.PrivatecaBetaCertificateTemplateIdentityConstraintsCelExpression {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplateIdentityConstraintsCelExpression{}
	p.SetExpression(dcl.ValueOrEmptyString(o.Expression))
	p.SetTitle(dcl.ValueOrEmptyString(o.Title))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	return p
}

// CertificateTemplatePassthroughExtensionsToProto converts a CertificateTemplatePassthroughExtensions object to its proto representation.
func PrivatecaBetaCertificateTemplatePassthroughExtensionsToProto(o *beta.CertificateTemplatePassthroughExtensions) *betapb.PrivatecaBetaCertificateTemplatePassthroughExtensions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplatePassthroughExtensions{}
	sKnownExtensions := make([]betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum, len(o.KnownExtensions))
	for i, r := range o.KnownExtensions {
		sKnownExtensions[i] = betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum_value[string(r)])
	}
	p.SetKnownExtensions(sKnownExtensions)
	sAdditionalExtensions := make([]*betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaBetaCertificateTemplatePassthroughExtensionsAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CertificateTemplatePassthroughExtensionsAdditionalExtensionsToProto converts a CertificateTemplatePassthroughExtensionsAdditionalExtensions object to its proto representation.
func PrivatecaBetaCertificateTemplatePassthroughExtensionsAdditionalExtensionsToProto(o *beta.CertificateTemplatePassthroughExtensionsAdditionalExtensions) *betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsAdditionalExtensions{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateTemplateToProto converts a CertificateTemplate resource to its proto representation.
func CertificateTemplateToProto(resource *beta.CertificateTemplate) *betapb.PrivatecaBetaCertificateTemplate {
	p := &betapb.PrivatecaBetaCertificateTemplate{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetPredefinedValues(PrivatecaBetaCertificateTemplatePredefinedValuesToProto(resource.PredefinedValues))
	p.SetIdentityConstraints(PrivatecaBetaCertificateTemplateIdentityConstraintsToProto(resource.IdentityConstraints))
	p.SetPassthroughExtensions(PrivatecaBetaCertificateTemplatePassthroughExtensionsToProto(resource.PassthroughExtensions))
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
func (s *CertificateTemplateServer) applyCertificateTemplate(ctx context.Context, c *beta.Client, request *betapb.ApplyPrivatecaBetaCertificateTemplateRequest) (*betapb.PrivatecaBetaCertificateTemplate, error) {
	p := ProtoToCertificateTemplate(request.GetResource())
	res, err := c.ApplyCertificateTemplate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CertificateTemplateToProto(res)
	return r, nil
}

// applyPrivatecaBetaCertificateTemplate handles the gRPC request by passing it to the underlying CertificateTemplate Apply() method.
func (s *CertificateTemplateServer) ApplyPrivatecaBetaCertificateTemplate(ctx context.Context, request *betapb.ApplyPrivatecaBetaCertificateTemplateRequest) (*betapb.PrivatecaBetaCertificateTemplate, error) {
	cl, err := createConfigCertificateTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCertificateTemplate(ctx, cl, request)
}

// DeleteCertificateTemplate handles the gRPC request by passing it to the underlying CertificateTemplate Delete() method.
func (s *CertificateTemplateServer) DeletePrivatecaBetaCertificateTemplate(ctx context.Context, request *betapb.DeletePrivatecaBetaCertificateTemplateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCertificateTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCertificateTemplate(ctx, ProtoToCertificateTemplate(request.GetResource()))

}

// ListPrivatecaBetaCertificateTemplate handles the gRPC request by passing it to the underlying CertificateTemplateList() method.
func (s *CertificateTemplateServer) ListPrivatecaBetaCertificateTemplate(ctx context.Context, request *betapb.ListPrivatecaBetaCertificateTemplateRequest) (*betapb.ListPrivatecaBetaCertificateTemplateResponse, error) {
	cl, err := createConfigCertificateTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCertificateTemplate(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.PrivatecaBetaCertificateTemplate
	for _, r := range resources.Items {
		rp := CertificateTemplateToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListPrivatecaBetaCertificateTemplateResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCertificateTemplate(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
