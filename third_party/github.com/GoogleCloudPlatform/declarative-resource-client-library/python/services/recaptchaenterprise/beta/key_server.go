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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/recaptchaenterprise/beta/recaptchaenterprise_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/recaptchaenterprise/beta"
)

// KeyServer implements the gRPC interface for Key.
type KeyServer struct{}

// ProtoToKeyWebSettingsIntegrationTypeEnum converts a KeyWebSettingsIntegrationTypeEnum enum from its proto representation.
func ProtoToRecaptchaenterpriseBetaKeyWebSettingsIntegrationTypeEnum(e betapb.RecaptchaenterpriseBetaKeyWebSettingsIntegrationTypeEnum) *beta.KeyWebSettingsIntegrationTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.RecaptchaenterpriseBetaKeyWebSettingsIntegrationTypeEnum_name[int32(e)]; ok {
		e := beta.KeyWebSettingsIntegrationTypeEnum(n[len("RecaptchaenterpriseBetaKeyWebSettingsIntegrationTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToKeyWebSettingsChallengeSecurityPreferenceEnum converts a KeyWebSettingsChallengeSecurityPreferenceEnum enum from its proto representation.
func ProtoToRecaptchaenterpriseBetaKeyWebSettingsChallengeSecurityPreferenceEnum(e betapb.RecaptchaenterpriseBetaKeyWebSettingsChallengeSecurityPreferenceEnum) *beta.KeyWebSettingsChallengeSecurityPreferenceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.RecaptchaenterpriseBetaKeyWebSettingsChallengeSecurityPreferenceEnum_name[int32(e)]; ok {
		e := beta.KeyWebSettingsChallengeSecurityPreferenceEnum(n[len("RecaptchaenterpriseBetaKeyWebSettingsChallengeSecurityPreferenceEnum"):])
		return &e
	}
	return nil
}

// ProtoToKeyTestingOptionsTestingChallengeEnum converts a KeyTestingOptionsTestingChallengeEnum enum from its proto representation.
func ProtoToRecaptchaenterpriseBetaKeyTestingOptionsTestingChallengeEnum(e betapb.RecaptchaenterpriseBetaKeyTestingOptionsTestingChallengeEnum) *beta.KeyTestingOptionsTestingChallengeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.RecaptchaenterpriseBetaKeyTestingOptionsTestingChallengeEnum_name[int32(e)]; ok {
		e := beta.KeyTestingOptionsTestingChallengeEnum(n[len("RecaptchaenterpriseBetaKeyTestingOptionsTestingChallengeEnum"):])
		return &e
	}
	return nil
}

// ProtoToKeyWafSettingsWafServiceEnum converts a KeyWafSettingsWafServiceEnum enum from its proto representation.
func ProtoToRecaptchaenterpriseBetaKeyWafSettingsWafServiceEnum(e betapb.RecaptchaenterpriseBetaKeyWafSettingsWafServiceEnum) *beta.KeyWafSettingsWafServiceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.RecaptchaenterpriseBetaKeyWafSettingsWafServiceEnum_name[int32(e)]; ok {
		e := beta.KeyWafSettingsWafServiceEnum(n[len("RecaptchaenterpriseBetaKeyWafSettingsWafServiceEnum"):])
		return &e
	}
	return nil
}

// ProtoToKeyWafSettingsWafFeatureEnum converts a KeyWafSettingsWafFeatureEnum enum from its proto representation.
func ProtoToRecaptchaenterpriseBetaKeyWafSettingsWafFeatureEnum(e betapb.RecaptchaenterpriseBetaKeyWafSettingsWafFeatureEnum) *beta.KeyWafSettingsWafFeatureEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.RecaptchaenterpriseBetaKeyWafSettingsWafFeatureEnum_name[int32(e)]; ok {
		e := beta.KeyWafSettingsWafFeatureEnum(n[len("RecaptchaenterpriseBetaKeyWafSettingsWafFeatureEnum"):])
		return &e
	}
	return nil
}

// ProtoToKeyWebSettings converts a KeyWebSettings object from its proto representation.
func ProtoToRecaptchaenterpriseBetaKeyWebSettings(p *betapb.RecaptchaenterpriseBetaKeyWebSettings) *beta.KeyWebSettings {
	if p == nil {
		return nil
	}
	obj := &beta.KeyWebSettings{
		AllowAllDomains:             dcl.Bool(p.GetAllowAllDomains()),
		AllowAmpTraffic:             dcl.Bool(p.GetAllowAmpTraffic()),
		IntegrationType:             ProtoToRecaptchaenterpriseBetaKeyWebSettingsIntegrationTypeEnum(p.GetIntegrationType()),
		ChallengeSecurityPreference: ProtoToRecaptchaenterpriseBetaKeyWebSettingsChallengeSecurityPreferenceEnum(p.GetChallengeSecurityPreference()),
	}
	for _, r := range p.GetAllowedDomains() {
		obj.AllowedDomains = append(obj.AllowedDomains, r)
	}
	return obj
}

// ProtoToKeyAndroidSettings converts a KeyAndroidSettings object from its proto representation.
func ProtoToRecaptchaenterpriseBetaKeyAndroidSettings(p *betapb.RecaptchaenterpriseBetaKeyAndroidSettings) *beta.KeyAndroidSettings {
	if p == nil {
		return nil
	}
	obj := &beta.KeyAndroidSettings{
		AllowAllPackageNames: dcl.Bool(p.GetAllowAllPackageNames()),
	}
	for _, r := range p.GetAllowedPackageNames() {
		obj.AllowedPackageNames = append(obj.AllowedPackageNames, r)
	}
	return obj
}

// ProtoToKeyIosSettings converts a KeyIosSettings object from its proto representation.
func ProtoToRecaptchaenterpriseBetaKeyIosSettings(p *betapb.RecaptchaenterpriseBetaKeyIosSettings) *beta.KeyIosSettings {
	if p == nil {
		return nil
	}
	obj := &beta.KeyIosSettings{
		AllowAllBundleIds: dcl.Bool(p.GetAllowAllBundleIds()),
	}
	for _, r := range p.GetAllowedBundleIds() {
		obj.AllowedBundleIds = append(obj.AllowedBundleIds, r)
	}
	return obj
}

// ProtoToKeyTestingOptions converts a KeyTestingOptions object from its proto representation.
func ProtoToRecaptchaenterpriseBetaKeyTestingOptions(p *betapb.RecaptchaenterpriseBetaKeyTestingOptions) *beta.KeyTestingOptions {
	if p == nil {
		return nil
	}
	obj := &beta.KeyTestingOptions{
		TestingScore:     dcl.Float64OrNil(p.GetTestingScore()),
		TestingChallenge: ProtoToRecaptchaenterpriseBetaKeyTestingOptionsTestingChallengeEnum(p.GetTestingChallenge()),
	}
	return obj
}

// ProtoToKeyWafSettings converts a KeyWafSettings object from its proto representation.
func ProtoToRecaptchaenterpriseBetaKeyWafSettings(p *betapb.RecaptchaenterpriseBetaKeyWafSettings) *beta.KeyWafSettings {
	if p == nil {
		return nil
	}
	obj := &beta.KeyWafSettings{
		WafService: ProtoToRecaptchaenterpriseBetaKeyWafSettingsWafServiceEnum(p.GetWafService()),
		WafFeature: ProtoToRecaptchaenterpriseBetaKeyWafSettingsWafFeatureEnum(p.GetWafFeature()),
	}
	return obj
}

// ProtoToKey converts a Key resource from its proto representation.
func ProtoToKey(p *betapb.RecaptchaenterpriseBetaKey) *beta.Key {
	obj := &beta.Key{
		Name:            dcl.StringOrNil(p.GetName()),
		DisplayName:     dcl.StringOrNil(p.GetDisplayName()),
		WebSettings:     ProtoToRecaptchaenterpriseBetaKeyWebSettings(p.GetWebSettings()),
		AndroidSettings: ProtoToRecaptchaenterpriseBetaKeyAndroidSettings(p.GetAndroidSettings()),
		IosSettings:     ProtoToRecaptchaenterpriseBetaKeyIosSettings(p.GetIosSettings()),
		CreateTime:      dcl.StringOrNil(p.GetCreateTime()),
		TestingOptions:  ProtoToRecaptchaenterpriseBetaKeyTestingOptions(p.GetTestingOptions()),
		WafSettings:     ProtoToRecaptchaenterpriseBetaKeyWafSettings(p.GetWafSettings()),
		Project:         dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// KeyWebSettingsIntegrationTypeEnumToProto converts a KeyWebSettingsIntegrationTypeEnum enum to its proto representation.
func RecaptchaenterpriseBetaKeyWebSettingsIntegrationTypeEnumToProto(e *beta.KeyWebSettingsIntegrationTypeEnum) betapb.RecaptchaenterpriseBetaKeyWebSettingsIntegrationTypeEnum {
	if e == nil {
		return betapb.RecaptchaenterpriseBetaKeyWebSettingsIntegrationTypeEnum(0)
	}
	if v, ok := betapb.RecaptchaenterpriseBetaKeyWebSettingsIntegrationTypeEnum_value["KeyWebSettingsIntegrationTypeEnum"+string(*e)]; ok {
		return betapb.RecaptchaenterpriseBetaKeyWebSettingsIntegrationTypeEnum(v)
	}
	return betapb.RecaptchaenterpriseBetaKeyWebSettingsIntegrationTypeEnum(0)
}

// KeyWebSettingsChallengeSecurityPreferenceEnumToProto converts a KeyWebSettingsChallengeSecurityPreferenceEnum enum to its proto representation.
func RecaptchaenterpriseBetaKeyWebSettingsChallengeSecurityPreferenceEnumToProto(e *beta.KeyWebSettingsChallengeSecurityPreferenceEnum) betapb.RecaptchaenterpriseBetaKeyWebSettingsChallengeSecurityPreferenceEnum {
	if e == nil {
		return betapb.RecaptchaenterpriseBetaKeyWebSettingsChallengeSecurityPreferenceEnum(0)
	}
	if v, ok := betapb.RecaptchaenterpriseBetaKeyWebSettingsChallengeSecurityPreferenceEnum_value["KeyWebSettingsChallengeSecurityPreferenceEnum"+string(*e)]; ok {
		return betapb.RecaptchaenterpriseBetaKeyWebSettingsChallengeSecurityPreferenceEnum(v)
	}
	return betapb.RecaptchaenterpriseBetaKeyWebSettingsChallengeSecurityPreferenceEnum(0)
}

// KeyTestingOptionsTestingChallengeEnumToProto converts a KeyTestingOptionsTestingChallengeEnum enum to its proto representation.
func RecaptchaenterpriseBetaKeyTestingOptionsTestingChallengeEnumToProto(e *beta.KeyTestingOptionsTestingChallengeEnum) betapb.RecaptchaenterpriseBetaKeyTestingOptionsTestingChallengeEnum {
	if e == nil {
		return betapb.RecaptchaenterpriseBetaKeyTestingOptionsTestingChallengeEnum(0)
	}
	if v, ok := betapb.RecaptchaenterpriseBetaKeyTestingOptionsTestingChallengeEnum_value["KeyTestingOptionsTestingChallengeEnum"+string(*e)]; ok {
		return betapb.RecaptchaenterpriseBetaKeyTestingOptionsTestingChallengeEnum(v)
	}
	return betapb.RecaptchaenterpriseBetaKeyTestingOptionsTestingChallengeEnum(0)
}

// KeyWafSettingsWafServiceEnumToProto converts a KeyWafSettingsWafServiceEnum enum to its proto representation.
func RecaptchaenterpriseBetaKeyWafSettingsWafServiceEnumToProto(e *beta.KeyWafSettingsWafServiceEnum) betapb.RecaptchaenterpriseBetaKeyWafSettingsWafServiceEnum {
	if e == nil {
		return betapb.RecaptchaenterpriseBetaKeyWafSettingsWafServiceEnum(0)
	}
	if v, ok := betapb.RecaptchaenterpriseBetaKeyWafSettingsWafServiceEnum_value["KeyWafSettingsWafServiceEnum"+string(*e)]; ok {
		return betapb.RecaptchaenterpriseBetaKeyWafSettingsWafServiceEnum(v)
	}
	return betapb.RecaptchaenterpriseBetaKeyWafSettingsWafServiceEnum(0)
}

// KeyWafSettingsWafFeatureEnumToProto converts a KeyWafSettingsWafFeatureEnum enum to its proto representation.
func RecaptchaenterpriseBetaKeyWafSettingsWafFeatureEnumToProto(e *beta.KeyWafSettingsWafFeatureEnum) betapb.RecaptchaenterpriseBetaKeyWafSettingsWafFeatureEnum {
	if e == nil {
		return betapb.RecaptchaenterpriseBetaKeyWafSettingsWafFeatureEnum(0)
	}
	if v, ok := betapb.RecaptchaenterpriseBetaKeyWafSettingsWafFeatureEnum_value["KeyWafSettingsWafFeatureEnum"+string(*e)]; ok {
		return betapb.RecaptchaenterpriseBetaKeyWafSettingsWafFeatureEnum(v)
	}
	return betapb.RecaptchaenterpriseBetaKeyWafSettingsWafFeatureEnum(0)
}

// KeyWebSettingsToProto converts a KeyWebSettings object to its proto representation.
func RecaptchaenterpriseBetaKeyWebSettingsToProto(o *beta.KeyWebSettings) *betapb.RecaptchaenterpriseBetaKeyWebSettings {
	if o == nil {
		return nil
	}
	p := &betapb.RecaptchaenterpriseBetaKeyWebSettings{}
	p.SetAllowAllDomains(dcl.ValueOrEmptyBool(o.AllowAllDomains))
	p.SetAllowAmpTraffic(dcl.ValueOrEmptyBool(o.AllowAmpTraffic))
	p.SetIntegrationType(RecaptchaenterpriseBetaKeyWebSettingsIntegrationTypeEnumToProto(o.IntegrationType))
	p.SetChallengeSecurityPreference(RecaptchaenterpriseBetaKeyWebSettingsChallengeSecurityPreferenceEnumToProto(o.ChallengeSecurityPreference))
	sAllowedDomains := make([]string, len(o.AllowedDomains))
	for i, r := range o.AllowedDomains {
		sAllowedDomains[i] = r
	}
	p.SetAllowedDomains(sAllowedDomains)
	return p
}

// KeyAndroidSettingsToProto converts a KeyAndroidSettings object to its proto representation.
func RecaptchaenterpriseBetaKeyAndroidSettingsToProto(o *beta.KeyAndroidSettings) *betapb.RecaptchaenterpriseBetaKeyAndroidSettings {
	if o == nil {
		return nil
	}
	p := &betapb.RecaptchaenterpriseBetaKeyAndroidSettings{}
	p.SetAllowAllPackageNames(dcl.ValueOrEmptyBool(o.AllowAllPackageNames))
	sAllowedPackageNames := make([]string, len(o.AllowedPackageNames))
	for i, r := range o.AllowedPackageNames {
		sAllowedPackageNames[i] = r
	}
	p.SetAllowedPackageNames(sAllowedPackageNames)
	return p
}

// KeyIosSettingsToProto converts a KeyIosSettings object to its proto representation.
func RecaptchaenterpriseBetaKeyIosSettingsToProto(o *beta.KeyIosSettings) *betapb.RecaptchaenterpriseBetaKeyIosSettings {
	if o == nil {
		return nil
	}
	p := &betapb.RecaptchaenterpriseBetaKeyIosSettings{}
	p.SetAllowAllBundleIds(dcl.ValueOrEmptyBool(o.AllowAllBundleIds))
	sAllowedBundleIds := make([]string, len(o.AllowedBundleIds))
	for i, r := range o.AllowedBundleIds {
		sAllowedBundleIds[i] = r
	}
	p.SetAllowedBundleIds(sAllowedBundleIds)
	return p
}

// KeyTestingOptionsToProto converts a KeyTestingOptions object to its proto representation.
func RecaptchaenterpriseBetaKeyTestingOptionsToProto(o *beta.KeyTestingOptions) *betapb.RecaptchaenterpriseBetaKeyTestingOptions {
	if o == nil {
		return nil
	}
	p := &betapb.RecaptchaenterpriseBetaKeyTestingOptions{}
	p.SetTestingScore(dcl.ValueOrEmptyDouble(o.TestingScore))
	p.SetTestingChallenge(RecaptchaenterpriseBetaKeyTestingOptionsTestingChallengeEnumToProto(o.TestingChallenge))
	return p
}

// KeyWafSettingsToProto converts a KeyWafSettings object to its proto representation.
func RecaptchaenterpriseBetaKeyWafSettingsToProto(o *beta.KeyWafSettings) *betapb.RecaptchaenterpriseBetaKeyWafSettings {
	if o == nil {
		return nil
	}
	p := &betapb.RecaptchaenterpriseBetaKeyWafSettings{}
	p.SetWafService(RecaptchaenterpriseBetaKeyWafSettingsWafServiceEnumToProto(o.WafService))
	p.SetWafFeature(RecaptchaenterpriseBetaKeyWafSettingsWafFeatureEnumToProto(o.WafFeature))
	return p
}

// KeyToProto converts a Key resource to its proto representation.
func KeyToProto(resource *beta.Key) *betapb.RecaptchaenterpriseBetaKey {
	p := &betapb.RecaptchaenterpriseBetaKey{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetWebSettings(RecaptchaenterpriseBetaKeyWebSettingsToProto(resource.WebSettings))
	p.SetAndroidSettings(RecaptchaenterpriseBetaKeyAndroidSettingsToProto(resource.AndroidSettings))
	p.SetIosSettings(RecaptchaenterpriseBetaKeyIosSettingsToProto(resource.IosSettings))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetTestingOptions(RecaptchaenterpriseBetaKeyTestingOptionsToProto(resource.TestingOptions))
	p.SetWafSettings(RecaptchaenterpriseBetaKeyWafSettingsToProto(resource.WafSettings))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyKey handles the gRPC request by passing it to the underlying Key Apply() method.
func (s *KeyServer) applyKey(ctx context.Context, c *beta.Client, request *betapb.ApplyRecaptchaenterpriseBetaKeyRequest) (*betapb.RecaptchaenterpriseBetaKey, error) {
	p := ProtoToKey(request.GetResource())
	res, err := c.ApplyKey(ctx, p)
	if err != nil {
		return nil, err
	}
	r := KeyToProto(res)
	return r, nil
}

// applyRecaptchaenterpriseBetaKey handles the gRPC request by passing it to the underlying Key Apply() method.
func (s *KeyServer) ApplyRecaptchaenterpriseBetaKey(ctx context.Context, request *betapb.ApplyRecaptchaenterpriseBetaKeyRequest) (*betapb.RecaptchaenterpriseBetaKey, error) {
	cl, err := createConfigKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyKey(ctx, cl, request)
}

// DeleteKey handles the gRPC request by passing it to the underlying Key Delete() method.
func (s *KeyServer) DeleteRecaptchaenterpriseBetaKey(ctx context.Context, request *betapb.DeleteRecaptchaenterpriseBetaKeyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteKey(ctx, ProtoToKey(request.GetResource()))

}

// ListRecaptchaenterpriseBetaKey handles the gRPC request by passing it to the underlying KeyList() method.
func (s *KeyServer) ListRecaptchaenterpriseBetaKey(ctx context.Context, request *betapb.ListRecaptchaenterpriseBetaKeyRequest) (*betapb.ListRecaptchaenterpriseBetaKeyResponse, error) {
	cl, err := createConfigKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListKey(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.RecaptchaenterpriseBetaKey
	for _, r := range resources.Items {
		rp := KeyToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListRecaptchaenterpriseBetaKeyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigKey(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
