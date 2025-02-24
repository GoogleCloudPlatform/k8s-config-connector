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

// CertificateTemplateServer implements the gRPC interface for CertificateTemplate.
type CertificateTemplateServer struct{}

// ProtoToCertificateTemplatePassthroughExtensionsKnownExtensionsEnum converts a CertificateTemplatePassthroughExtensionsKnownExtensionsEnum enum from its proto representation.
func ProtoToPrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(e privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum) *privateca.CertificateTemplatePassthroughExtensionsKnownExtensionsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum_name[int32(e)]; ok {
		e := privateca.CertificateTemplatePassthroughExtensionsKnownExtensionsEnum(n[len("PrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateTemplatePredefinedValues converts a CertificateTemplatePredefinedValues object from its proto representation.
func ProtoToPrivatecaCertificateTemplatePredefinedValues(p *privatecapb.PrivatecaCertificateTemplatePredefinedValues) *privateca.CertificateTemplatePredefinedValues {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplatePredefinedValues{
		KeyUsage:  ProtoToPrivatecaCertificateTemplatePredefinedValuesKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaCertificateTemplatePredefinedValuesCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaCertificateTemplatePredefinedValuesPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaCertificateTemplatePredefinedValuesAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesKeyUsage converts a CertificateTemplatePredefinedValuesKeyUsage object from its proto representation.
func ProtoToPrivatecaCertificateTemplatePredefinedValuesKeyUsage(p *privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsage) *privateca.CertificateTemplatePredefinedValuesKeyUsage {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplatePredefinedValuesKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage converts a CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage object from its proto representation.
func ProtoToPrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage(p *privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage) *privateca.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage{
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
func ProtoToPrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage(p *privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage) *privateca.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage{
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
func ProtoToPrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages(p *privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages) *privateca.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesCaOptions converts a CertificateTemplatePredefinedValuesCaOptions object from its proto representation.
func ProtoToPrivatecaCertificateTemplatePredefinedValuesCaOptions(p *privatecapb.PrivatecaCertificateTemplatePredefinedValuesCaOptions) *privateca.CertificateTemplatePredefinedValuesCaOptions {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplatePredefinedValuesCaOptions{
		IsCa:                dcl.Bool(p.GetIsCa()),
		MaxIssuerPathLength: dcl.Int64OrNil(p.GetMaxIssuerPathLength()),
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesPolicyIds converts a CertificateTemplatePredefinedValuesPolicyIds object from its proto representation.
func ProtoToPrivatecaCertificateTemplatePredefinedValuesPolicyIds(p *privatecapb.PrivatecaCertificateTemplatePredefinedValuesPolicyIds) *privateca.CertificateTemplatePredefinedValuesPolicyIds {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplatePredefinedValuesPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesAdditionalExtensions converts a CertificateTemplatePredefinedValuesAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaCertificateTemplatePredefinedValuesAdditionalExtensions(p *privatecapb.PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensions) *privateca.CertificateTemplatePredefinedValuesAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplatePredefinedValuesAdditionalExtensions{
		ObjectId: ProtoToPrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId converts a CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId object from its proto representation.
func ProtoToPrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId(p *privatecapb.PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId) *privateca.CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplateIdentityConstraints converts a CertificateTemplateIdentityConstraints object from its proto representation.
func ProtoToPrivatecaCertificateTemplateIdentityConstraints(p *privatecapb.PrivatecaCertificateTemplateIdentityConstraints) *privateca.CertificateTemplateIdentityConstraints {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplateIdentityConstraints{
		CelExpression:                   ProtoToPrivatecaCertificateTemplateIdentityConstraintsCelExpression(p.GetCelExpression()),
		AllowSubjectPassthrough:         dcl.Bool(p.GetAllowSubjectPassthrough()),
		AllowSubjectAltNamesPassthrough: dcl.Bool(p.GetAllowSubjectAltNamesPassthrough()),
	}
	return obj
}

// ProtoToCertificateTemplateIdentityConstraintsCelExpression converts a CertificateTemplateIdentityConstraintsCelExpression object from its proto representation.
func ProtoToPrivatecaCertificateTemplateIdentityConstraintsCelExpression(p *privatecapb.PrivatecaCertificateTemplateIdentityConstraintsCelExpression) *privateca.CertificateTemplateIdentityConstraintsCelExpression {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplateIdentityConstraintsCelExpression{
		Expression:  dcl.StringOrNil(p.GetExpression()),
		Title:       dcl.StringOrNil(p.GetTitle()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ProtoToCertificateTemplatePassthroughExtensions converts a CertificateTemplatePassthroughExtensions object from its proto representation.
func ProtoToPrivatecaCertificateTemplatePassthroughExtensions(p *privatecapb.PrivatecaCertificateTemplatePassthroughExtensions) *privateca.CertificateTemplatePassthroughExtensions {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplatePassthroughExtensions{}
	for _, r := range p.GetKnownExtensions() {
		obj.KnownExtensions = append(obj.KnownExtensions, *ProtoToPrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(r))
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateTemplatePassthroughExtensionsAdditionalExtensions converts a CertificateTemplatePassthroughExtensionsAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensions(p *privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensions) *privateca.CertificateTemplatePassthroughExtensionsAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplatePassthroughExtensionsAdditionalExtensions{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplate converts a CertificateTemplate resource from its proto representation.
func ProtoToCertificateTemplate(p *privatecapb.PrivatecaCertificateTemplate) *privateca.CertificateTemplate {
	obj := &privateca.CertificateTemplate{
		Name:                  dcl.StringOrNil(p.GetName()),
		PredefinedValues:      ProtoToPrivatecaCertificateTemplatePredefinedValues(p.GetPredefinedValues()),
		IdentityConstraints:   ProtoToPrivatecaCertificateTemplateIdentityConstraints(p.GetIdentityConstraints()),
		PassthroughExtensions: ProtoToPrivatecaCertificateTemplatePassthroughExtensions(p.GetPassthroughExtensions()),
		Description:           dcl.StringOrNil(p.GetDescription()),
		CreateTime:            dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:            dcl.StringOrNil(p.GetUpdateTime()),
		Project:               dcl.StringOrNil(p.GetProject()),
		Location:              dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// CertificateTemplatePassthroughExtensionsKnownExtensionsEnumToProto converts a CertificateTemplatePassthroughExtensionsKnownExtensionsEnum enum to its proto representation.
func PrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnumToProto(e *privateca.CertificateTemplatePassthroughExtensionsKnownExtensionsEnum) privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum {
	if e == nil {
		return privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(0)
	}
	if v, ok := privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum_value["CertificateTemplatePassthroughExtensionsKnownExtensionsEnum"+string(*e)]; ok {
		return privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(v)
	}
	return privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(0)
}

// CertificateTemplatePredefinedValuesToProto converts a CertificateTemplatePredefinedValues object to its proto representation.
func PrivatecaCertificateTemplatePredefinedValuesToProto(o *privateca.CertificateTemplatePredefinedValues) *privatecapb.PrivatecaCertificateTemplatePredefinedValues {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplatePredefinedValues{}
	p.SetKeyUsage(PrivatecaCertificateTemplatePredefinedValuesKeyUsageToProto(o.KeyUsage))
	p.SetCaOptions(PrivatecaCertificateTemplatePredefinedValuesCaOptionsToProto(o.CaOptions))
	sPolicyIds := make([]*privatecapb.PrivatecaCertificateTemplatePredefinedValuesPolicyIds, len(o.PolicyIds))
	for i, r := range o.PolicyIds {
		sPolicyIds[i] = PrivatecaCertificateTemplatePredefinedValuesPolicyIdsToProto(&r)
	}
	p.SetPolicyIds(sPolicyIds)
	sAiaOcspServers := make([]string, len(o.AiaOcspServers))
	for i, r := range o.AiaOcspServers {
		sAiaOcspServers[i] = r
	}
	p.SetAiaOcspServers(sAiaOcspServers)
	sAdditionalExtensions := make([]*privatecapb.PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CertificateTemplatePredefinedValuesKeyUsageToProto converts a CertificateTemplatePredefinedValuesKeyUsage object to its proto representation.
func PrivatecaCertificateTemplatePredefinedValuesKeyUsageToProto(o *privateca.CertificateTemplatePredefinedValuesKeyUsage) *privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsage {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsage{}
	p.SetBaseKeyUsage(PrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage))
	p.SetExtendedKeyUsage(PrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage))
	sUnknownExtendedKeyUsages := make([]*privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages, len(o.UnknownExtendedKeyUsages))
	for i, r := range o.UnknownExtendedKeyUsages {
		sUnknownExtendedKeyUsages[i] = PrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesToProto(&r)
	}
	p.SetUnknownExtendedKeyUsages(sUnknownExtendedKeyUsages)
	return p
}

// CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageToProto converts a CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage object to its proto representation.
func PrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageToProto(o *privateca.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage) *privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage{}
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
func PrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageToProto(o *privateca.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage) *privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage{}
	p.SetServerAuth(dcl.ValueOrEmptyBool(o.ServerAuth))
	p.SetClientAuth(dcl.ValueOrEmptyBool(o.ClientAuth))
	p.SetCodeSigning(dcl.ValueOrEmptyBool(o.CodeSigning))
	p.SetEmailProtection(dcl.ValueOrEmptyBool(o.EmailProtection))
	p.SetTimeStamping(dcl.ValueOrEmptyBool(o.TimeStamping))
	p.SetOcspSigning(dcl.ValueOrEmptyBool(o.OcspSigning))
	return p
}

// CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesToProto converts a CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages object to its proto representation.
func PrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesToProto(o *privateca.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages) *privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateTemplatePredefinedValuesCaOptionsToProto converts a CertificateTemplatePredefinedValuesCaOptions object to its proto representation.
func PrivatecaCertificateTemplatePredefinedValuesCaOptionsToProto(o *privateca.CertificateTemplatePredefinedValuesCaOptions) *privatecapb.PrivatecaCertificateTemplatePredefinedValuesCaOptions {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplatePredefinedValuesCaOptions{}
	p.SetIsCa(dcl.ValueOrEmptyBool(o.IsCa))
	p.SetMaxIssuerPathLength(dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength))
	return p
}

// CertificateTemplatePredefinedValuesPolicyIdsToProto converts a CertificateTemplatePredefinedValuesPolicyIds object to its proto representation.
func PrivatecaCertificateTemplatePredefinedValuesPolicyIdsToProto(o *privateca.CertificateTemplatePredefinedValuesPolicyIds) *privatecapb.PrivatecaCertificateTemplatePredefinedValuesPolicyIds {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplatePredefinedValuesPolicyIds{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateTemplatePredefinedValuesAdditionalExtensionsToProto converts a CertificateTemplatePredefinedValuesAdditionalExtensions object to its proto representation.
func PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsToProto(o *privateca.CertificateTemplatePredefinedValuesAdditionalExtensions) *privatecapb.PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensions{}
	p.SetObjectId(PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdToProto converts a CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId object to its proto representation.
func PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdToProto(o *privateca.CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId) *privatecapb.PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateTemplateIdentityConstraintsToProto converts a CertificateTemplateIdentityConstraints object to its proto representation.
func PrivatecaCertificateTemplateIdentityConstraintsToProto(o *privateca.CertificateTemplateIdentityConstraints) *privatecapb.PrivatecaCertificateTemplateIdentityConstraints {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplateIdentityConstraints{}
	p.SetCelExpression(PrivatecaCertificateTemplateIdentityConstraintsCelExpressionToProto(o.CelExpression))
	p.SetAllowSubjectPassthrough(dcl.ValueOrEmptyBool(o.AllowSubjectPassthrough))
	p.SetAllowSubjectAltNamesPassthrough(dcl.ValueOrEmptyBool(o.AllowSubjectAltNamesPassthrough))
	return p
}

// CertificateTemplateIdentityConstraintsCelExpressionToProto converts a CertificateTemplateIdentityConstraintsCelExpression object to its proto representation.
func PrivatecaCertificateTemplateIdentityConstraintsCelExpressionToProto(o *privateca.CertificateTemplateIdentityConstraintsCelExpression) *privatecapb.PrivatecaCertificateTemplateIdentityConstraintsCelExpression {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplateIdentityConstraintsCelExpression{}
	p.SetExpression(dcl.ValueOrEmptyString(o.Expression))
	p.SetTitle(dcl.ValueOrEmptyString(o.Title))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	return p
}

// CertificateTemplatePassthroughExtensionsToProto converts a CertificateTemplatePassthroughExtensions object to its proto representation.
func PrivatecaCertificateTemplatePassthroughExtensionsToProto(o *privateca.CertificateTemplatePassthroughExtensions) *privatecapb.PrivatecaCertificateTemplatePassthroughExtensions {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplatePassthroughExtensions{}
	sKnownExtensions := make([]privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum, len(o.KnownExtensions))
	for i, r := range o.KnownExtensions {
		sKnownExtensions[i] = privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum_value[string(r)])
	}
	p.SetKnownExtensions(sKnownExtensions)
	sAdditionalExtensions := make([]*privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CertificateTemplatePassthroughExtensionsAdditionalExtensionsToProto converts a CertificateTemplatePassthroughExtensionsAdditionalExtensions object to its proto representation.
func PrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensionsToProto(o *privateca.CertificateTemplatePassthroughExtensionsAdditionalExtensions) *privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensions{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateTemplateToProto converts a CertificateTemplate resource to its proto representation.
func CertificateTemplateToProto(resource *privateca.CertificateTemplate) *privatecapb.PrivatecaCertificateTemplate {
	p := &privatecapb.PrivatecaCertificateTemplate{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetPredefinedValues(PrivatecaCertificateTemplatePredefinedValuesToProto(resource.PredefinedValues))
	p.SetIdentityConstraints(PrivatecaCertificateTemplateIdentityConstraintsToProto(resource.IdentityConstraints))
	p.SetPassthroughExtensions(PrivatecaCertificateTemplatePassthroughExtensionsToProto(resource.PassthroughExtensions))
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
func (s *CertificateTemplateServer) applyCertificateTemplate(ctx context.Context, c *privateca.Client, request *privatecapb.ApplyPrivatecaCertificateTemplateRequest) (*privatecapb.PrivatecaCertificateTemplate, error) {
	p := ProtoToCertificateTemplate(request.GetResource())
	res, err := c.ApplyCertificateTemplate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CertificateTemplateToProto(res)
	return r, nil
}

// applyPrivatecaCertificateTemplate handles the gRPC request by passing it to the underlying CertificateTemplate Apply() method.
func (s *CertificateTemplateServer) ApplyPrivatecaCertificateTemplate(ctx context.Context, request *privatecapb.ApplyPrivatecaCertificateTemplateRequest) (*privatecapb.PrivatecaCertificateTemplate, error) {
	cl, err := createConfigCertificateTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCertificateTemplate(ctx, cl, request)
}

// DeleteCertificateTemplate handles the gRPC request by passing it to the underlying CertificateTemplate Delete() method.
func (s *CertificateTemplateServer) DeletePrivatecaCertificateTemplate(ctx context.Context, request *privatecapb.DeletePrivatecaCertificateTemplateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCertificateTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCertificateTemplate(ctx, ProtoToCertificateTemplate(request.GetResource()))

}

// ListPrivatecaCertificateTemplate handles the gRPC request by passing it to the underlying CertificateTemplateList() method.
func (s *CertificateTemplateServer) ListPrivatecaCertificateTemplate(ctx context.Context, request *privatecapb.ListPrivatecaCertificateTemplateRequest) (*privatecapb.ListPrivatecaCertificateTemplateResponse, error) {
	cl, err := createConfigCertificateTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCertificateTemplate(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*privatecapb.PrivatecaCertificateTemplate
	for _, r := range resources.Items {
		rp := CertificateTemplateToProto(r)
		protos = append(protos, rp)
	}
	p := &privatecapb.ListPrivatecaCertificateTemplateResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCertificateTemplate(ctx context.Context, service_account_file string) (*privateca.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return privateca.NewClient(conf), nil
}
