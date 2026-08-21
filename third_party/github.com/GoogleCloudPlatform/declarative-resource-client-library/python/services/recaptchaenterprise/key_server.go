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
	recaptchaenterprisepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/recaptchaenterprise/recaptchaenterprise_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/recaptchaenterprise"
)

// KeyServer implements the gRPC interface for Key.
type KeyServer struct{}

// ProtoToKeyWebSettingsIntegrationTypeEnum converts a KeyWebSettingsIntegrationTypeEnum enum from its proto representation.
func ProtoToRecaptchaenterpriseKeyWebSettingsIntegrationTypeEnum(e recaptchaenterprisepb.RecaptchaenterpriseKeyWebSettingsIntegrationTypeEnum) *recaptchaenterprise.KeyWebSettingsIntegrationTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := recaptchaenterprisepb.RecaptchaenterpriseKeyWebSettingsIntegrationTypeEnum_name[int32(e)]; ok {
		e := recaptchaenterprise.KeyWebSettingsIntegrationTypeEnum(n[len("RecaptchaenterpriseKeyWebSettingsIntegrationTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToKeyWebSettingsChallengeSecurityPreferenceEnum converts a KeyWebSettingsChallengeSecurityPreferenceEnum enum from its proto representation.
func ProtoToRecaptchaenterpriseKeyWebSettingsChallengeSecurityPreferenceEnum(e recaptchaenterprisepb.RecaptchaenterpriseKeyWebSettingsChallengeSecurityPreferenceEnum) *recaptchaenterprise.KeyWebSettingsChallengeSecurityPreferenceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := recaptchaenterprisepb.RecaptchaenterpriseKeyWebSettingsChallengeSecurityPreferenceEnum_name[int32(e)]; ok {
		e := recaptchaenterprise.KeyWebSettingsChallengeSecurityPreferenceEnum(n[len("RecaptchaenterpriseKeyWebSettingsChallengeSecurityPreferenceEnum"):])
		return &e
	}
	return nil
}

// ProtoToKeyTestingOptionsTestingChallengeEnum converts a KeyTestingOptionsTestingChallengeEnum enum from its proto representation.
func ProtoToRecaptchaenterpriseKeyTestingOptionsTestingChallengeEnum(e recaptchaenterprisepb.RecaptchaenterpriseKeyTestingOptionsTestingChallengeEnum) *recaptchaenterprise.KeyTestingOptionsTestingChallengeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := recaptchaenterprisepb.RecaptchaenterpriseKeyTestingOptionsTestingChallengeEnum_name[int32(e)]; ok {
		e := recaptchaenterprise.KeyTestingOptionsTestingChallengeEnum(n[len("RecaptchaenterpriseKeyTestingOptionsTestingChallengeEnum"):])
		return &e
	}
	return nil
}

// ProtoToKeyWafSettingsWafServiceEnum converts a KeyWafSettingsWafServiceEnum enum from its proto representation.
func ProtoToRecaptchaenterpriseKeyWafSettingsWafServiceEnum(e recaptchaenterprisepb.RecaptchaenterpriseKeyWafSettingsWafServiceEnum) *recaptchaenterprise.KeyWafSettingsWafServiceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := recaptchaenterprisepb.RecaptchaenterpriseKeyWafSettingsWafServiceEnum_name[int32(e)]; ok {
		e := recaptchaenterprise.KeyWafSettingsWafServiceEnum(n[len("RecaptchaenterpriseKeyWafSettingsWafServiceEnum"):])
		return &e
	}
	return nil
}

// ProtoToKeyWafSettingsWafFeatureEnum converts a KeyWafSettingsWafFeatureEnum enum from its proto representation.
func ProtoToRecaptchaenterpriseKeyWafSettingsWafFeatureEnum(e recaptchaenterprisepb.RecaptchaenterpriseKeyWafSettingsWafFeatureEnum) *recaptchaenterprise.KeyWafSettingsWafFeatureEnum {
	if e == 0 {
		return nil
	}
	if n, ok := recaptchaenterprisepb.RecaptchaenterpriseKeyWafSettingsWafFeatureEnum_name[int32(e)]; ok {
		e := recaptchaenterprise.KeyWafSettingsWafFeatureEnum(n[len("RecaptchaenterpriseKeyWafSettingsWafFeatureEnum"):])
		return &e
	}
	return nil
}

// ProtoToKeyWebSettings converts a KeyWebSettings object from its proto representation.
func ProtoToRecaptchaenterpriseKeyWebSettings(p *recaptchaenterprisepb.RecaptchaenterpriseKeyWebSettings) *recaptchaenterprise.KeyWebSettings {
	if p == nil {
		return nil
	}
	obj := &recaptchaenterprise.KeyWebSettings{
		AllowAllDomains:             dcl.Bool(p.GetAllowAllDomains()),
		AllowAmpTraffic:             dcl.Bool(p.GetAllowAmpTraffic()),
		IntegrationType:             ProtoToRecaptchaenterpriseKeyWebSettingsIntegrationTypeEnum(p.GetIntegrationType()),
		ChallengeSecurityPreference: ProtoToRecaptchaenterpriseKeyWebSettingsChallengeSecurityPreferenceEnum(p.GetChallengeSecurityPreference()),
	}
	for _, r := range p.GetAllowedDomains() {
		obj.AllowedDomains = append(obj.AllowedDomains, r)
	}
	return obj
}

// ProtoToKeyAndroidSettings converts a KeyAndroidSettings object from its proto representation.
func ProtoToRecaptchaenterpriseKeyAndroidSettings(p *recaptchaenterprisepb.RecaptchaenterpriseKeyAndroidSettings) *recaptchaenterprise.KeyAndroidSettings {
	if p == nil {
		return nil
	}
	obj := &recaptchaenterprise.KeyAndroidSettings{
		AllowAllPackageNames: dcl.Bool(p.GetAllowAllPackageNames()),
	}
	for _, r := range p.GetAllowedPackageNames() {
		obj.AllowedPackageNames = append(obj.AllowedPackageNames, r)
	}
	return obj
}

// ProtoToKeyIosSettings converts a KeyIosSettings object from its proto representation.
func ProtoToRecaptchaenterpriseKeyIosSettings(p *recaptchaenterprisepb.RecaptchaenterpriseKeyIosSettings) *recaptchaenterprise.KeyIosSettings {
	if p == nil {
		return nil
	}
	obj := &recaptchaenterprise.KeyIosSettings{
		AllowAllBundleIds: dcl.Bool(p.GetAllowAllBundleIds()),
	}
	for _, r := range p.GetAllowedBundleIds() {
		obj.AllowedBundleIds = append(obj.AllowedBundleIds, r)
	}
	return obj
}

// ProtoToKeyTestingOptions converts a KeyTestingOptions object from its proto representation.
func ProtoToRecaptchaenterpriseKeyTestingOptions(p *recaptchaenterprisepb.RecaptchaenterpriseKeyTestingOptions) *recaptchaenterprise.KeyTestingOptions {
	if p == nil {
		return nil
	}
	obj := &recaptchaenterprise.KeyTestingOptions{
		TestingScore:     dcl.Float64OrNil(p.GetTestingScore()),
		TestingChallenge: ProtoToRecaptchaenterpriseKeyTestingOptionsTestingChallengeEnum(p.GetTestingChallenge()),
	}
	return obj
}

// ProtoToKeyWafSettings converts a KeyWafSettings object from its proto representation.
func ProtoToRecaptchaenterpriseKeyWafSettings(p *recaptchaenterprisepb.RecaptchaenterpriseKeyWafSettings) *recaptchaenterprise.KeyWafSettings {
	if p == nil {
		return nil
	}
	obj := &recaptchaenterprise.KeyWafSettings{
		WafService: ProtoToRecaptchaenterpriseKeyWafSettingsWafServiceEnum(p.GetWafService()),
		WafFeature: ProtoToRecaptchaenterpriseKeyWafSettingsWafFeatureEnum(p.GetWafFeature()),
	}
	return obj
}

// ProtoToKey converts a Key resource from its proto representation.
func ProtoToKey(p *recaptchaenterprisepb.RecaptchaenterpriseKey) *recaptchaenterprise.Key {
	obj := &recaptchaenterprise.Key{
		Name:            dcl.StringOrNil(p.GetName()),
		DisplayName:     dcl.StringOrNil(p.GetDisplayName()),
		WebSettings:     ProtoToRecaptchaenterpriseKeyWebSettings(p.GetWebSettings()),
		AndroidSettings: ProtoToRecaptchaenterpriseKeyAndroidSettings(p.GetAndroidSettings()),
		IosSettings:     ProtoToRecaptchaenterpriseKeyIosSettings(p.GetIosSettings()),
		CreateTime:      dcl.StringOrNil(p.GetCreateTime()),
		TestingOptions:  ProtoToRecaptchaenterpriseKeyTestingOptions(p.GetTestingOptions()),
		WafSettings:     ProtoToRecaptchaenterpriseKeyWafSettings(p.GetWafSettings()),
		Project:         dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// KeyWebSettingsIntegrationTypeEnumToProto converts a KeyWebSettingsIntegrationTypeEnum enum to its proto representation.
func RecaptchaenterpriseKeyWebSettingsIntegrationTypeEnumToProto(e *recaptchaenterprise.KeyWebSettingsIntegrationTypeEnum) recaptchaenterprisepb.RecaptchaenterpriseKeyWebSettingsIntegrationTypeEnum {
	if e == nil {
		return recaptchaenterprisepb.RecaptchaenterpriseKeyWebSettingsIntegrationTypeEnum(0)
	}
	if v, ok := recaptchaenterprisepb.RecaptchaenterpriseKeyWebSettingsIntegrationTypeEnum_value["KeyWebSettingsIntegrationTypeEnum"+string(*e)]; ok {
		return recaptchaenterprisepb.RecaptchaenterpriseKeyWebSettingsIntegrationTypeEnum(v)
	}
	return recaptchaenterprisepb.RecaptchaenterpriseKeyWebSettingsIntegrationTypeEnum(0)
}

// KeyWebSettingsChallengeSecurityPreferenceEnumToProto converts a KeyWebSettingsChallengeSecurityPreferenceEnum enum to its proto representation.
func RecaptchaenterpriseKeyWebSettingsChallengeSecurityPreferenceEnumToProto(e *recaptchaenterprise.KeyWebSettingsChallengeSecurityPreferenceEnum) recaptchaenterprisepb.RecaptchaenterpriseKeyWebSettingsChallengeSecurityPreferenceEnum {
	if e == nil {
		return recaptchaenterprisepb.RecaptchaenterpriseKeyWebSettingsChallengeSecurityPreferenceEnum(0)
	}
	if v, ok := recaptchaenterprisepb.RecaptchaenterpriseKeyWebSettingsChallengeSecurityPreferenceEnum_value["KeyWebSettingsChallengeSecurityPreferenceEnum"+string(*e)]; ok {
		return recaptchaenterprisepb.RecaptchaenterpriseKeyWebSettingsChallengeSecurityPreferenceEnum(v)
	}
	return recaptchaenterprisepb.RecaptchaenterpriseKeyWebSettingsChallengeSecurityPreferenceEnum(0)
}

// KeyTestingOptionsTestingChallengeEnumToProto converts a KeyTestingOptionsTestingChallengeEnum enum to its proto representation.
func RecaptchaenterpriseKeyTestingOptionsTestingChallengeEnumToProto(e *recaptchaenterprise.KeyTestingOptionsTestingChallengeEnum) recaptchaenterprisepb.RecaptchaenterpriseKeyTestingOptionsTestingChallengeEnum {
	if e == nil {
		return recaptchaenterprisepb.RecaptchaenterpriseKeyTestingOptionsTestingChallengeEnum(0)
	}
	if v, ok := recaptchaenterprisepb.RecaptchaenterpriseKeyTestingOptionsTestingChallengeEnum_value["KeyTestingOptionsTestingChallengeEnum"+string(*e)]; ok {
		return recaptchaenterprisepb.RecaptchaenterpriseKeyTestingOptionsTestingChallengeEnum(v)
	}
	return recaptchaenterprisepb.RecaptchaenterpriseKeyTestingOptionsTestingChallengeEnum(0)
}

// KeyWafSettingsWafServiceEnumToProto converts a KeyWafSettingsWafServiceEnum enum to its proto representation.
func RecaptchaenterpriseKeyWafSettingsWafServiceEnumToProto(e *recaptchaenterprise.KeyWafSettingsWafServiceEnum) recaptchaenterprisepb.RecaptchaenterpriseKeyWafSettingsWafServiceEnum {
	if e == nil {
		return recaptchaenterprisepb.RecaptchaenterpriseKeyWafSettingsWafServiceEnum(0)
	}
	if v, ok := recaptchaenterprisepb.RecaptchaenterpriseKeyWafSettingsWafServiceEnum_value["KeyWafSettingsWafServiceEnum"+string(*e)]; ok {
		return recaptchaenterprisepb.RecaptchaenterpriseKeyWafSettingsWafServiceEnum(v)
	}
	return recaptchaenterprisepb.RecaptchaenterpriseKeyWafSettingsWafServiceEnum(0)
}

// KeyWafSettingsWafFeatureEnumToProto converts a KeyWafSettingsWafFeatureEnum enum to its proto representation.
func RecaptchaenterpriseKeyWafSettingsWafFeatureEnumToProto(e *recaptchaenterprise.KeyWafSettingsWafFeatureEnum) recaptchaenterprisepb.RecaptchaenterpriseKeyWafSettingsWafFeatureEnum {
	if e == nil {
		return recaptchaenterprisepb.RecaptchaenterpriseKeyWafSettingsWafFeatureEnum(0)
	}
	if v, ok := recaptchaenterprisepb.RecaptchaenterpriseKeyWafSettingsWafFeatureEnum_value["KeyWafSettingsWafFeatureEnum"+string(*e)]; ok {
		return recaptchaenterprisepb.RecaptchaenterpriseKeyWafSettingsWafFeatureEnum(v)
	}
	return recaptchaenterprisepb.RecaptchaenterpriseKeyWafSettingsWafFeatureEnum(0)
}

// KeyWebSettingsToProto converts a KeyWebSettings object to its proto representation.
func RecaptchaenterpriseKeyWebSettingsToProto(o *recaptchaenterprise.KeyWebSettings) *recaptchaenterprisepb.RecaptchaenterpriseKeyWebSettings {
	if o == nil {
		return nil
	}
	p := &recaptchaenterprisepb.RecaptchaenterpriseKeyWebSettings{}
	p.SetAllowAllDomains(dcl.ValueOrEmptyBool(o.AllowAllDomains))
	p.SetAllowAmpTraffic(dcl.ValueOrEmptyBool(o.AllowAmpTraffic))
	p.SetIntegrationType(RecaptchaenterpriseKeyWebSettingsIntegrationTypeEnumToProto(o.IntegrationType))
	p.SetChallengeSecurityPreference(RecaptchaenterpriseKeyWebSettingsChallengeSecurityPreferenceEnumToProto(o.ChallengeSecurityPreference))
	sAllowedDomains := make([]string, len(o.AllowedDomains))
	for i, r := range o.AllowedDomains {
		sAllowedDomains[i] = r
	}
	p.SetAllowedDomains(sAllowedDomains)
	return p
}

// KeyAndroidSettingsToProto converts a KeyAndroidSettings object to its proto representation.
func RecaptchaenterpriseKeyAndroidSettingsToProto(o *recaptchaenterprise.KeyAndroidSettings) *recaptchaenterprisepb.RecaptchaenterpriseKeyAndroidSettings {
	if o == nil {
		return nil
	}
	p := &recaptchaenterprisepb.RecaptchaenterpriseKeyAndroidSettings{}
	p.SetAllowAllPackageNames(dcl.ValueOrEmptyBool(o.AllowAllPackageNames))
	sAllowedPackageNames := make([]string, len(o.AllowedPackageNames))
	for i, r := range o.AllowedPackageNames {
		sAllowedPackageNames[i] = r
	}
	p.SetAllowedPackageNames(sAllowedPackageNames)
	return p
}

// KeyIosSettingsToProto converts a KeyIosSettings object to its proto representation.
func RecaptchaenterpriseKeyIosSettingsToProto(o *recaptchaenterprise.KeyIosSettings) *recaptchaenterprisepb.RecaptchaenterpriseKeyIosSettings {
	if o == nil {
		return nil
	}
	p := &recaptchaenterprisepb.RecaptchaenterpriseKeyIosSettings{}
	p.SetAllowAllBundleIds(dcl.ValueOrEmptyBool(o.AllowAllBundleIds))
	sAllowedBundleIds := make([]string, len(o.AllowedBundleIds))
	for i, r := range o.AllowedBundleIds {
		sAllowedBundleIds[i] = r
	}
	p.SetAllowedBundleIds(sAllowedBundleIds)
	return p
}

// KeyTestingOptionsToProto converts a KeyTestingOptions object to its proto representation.
func RecaptchaenterpriseKeyTestingOptionsToProto(o *recaptchaenterprise.KeyTestingOptions) *recaptchaenterprisepb.RecaptchaenterpriseKeyTestingOptions {
	if o == nil {
		return nil
	}
	p := &recaptchaenterprisepb.RecaptchaenterpriseKeyTestingOptions{}
	p.SetTestingScore(dcl.ValueOrEmptyDouble(o.TestingScore))
	p.SetTestingChallenge(RecaptchaenterpriseKeyTestingOptionsTestingChallengeEnumToProto(o.TestingChallenge))
	return p
}

// KeyWafSettingsToProto converts a KeyWafSettings object to its proto representation.
func RecaptchaenterpriseKeyWafSettingsToProto(o *recaptchaenterprise.KeyWafSettings) *recaptchaenterprisepb.RecaptchaenterpriseKeyWafSettings {
	if o == nil {
		return nil
	}
	p := &recaptchaenterprisepb.RecaptchaenterpriseKeyWafSettings{}
	p.SetWafService(RecaptchaenterpriseKeyWafSettingsWafServiceEnumToProto(o.WafService))
	p.SetWafFeature(RecaptchaenterpriseKeyWafSettingsWafFeatureEnumToProto(o.WafFeature))
	return p
}

// KeyToProto converts a Key resource to its proto representation.
func KeyToProto(resource *recaptchaenterprise.Key) *recaptchaenterprisepb.RecaptchaenterpriseKey {
	p := &recaptchaenterprisepb.RecaptchaenterpriseKey{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetWebSettings(RecaptchaenterpriseKeyWebSettingsToProto(resource.WebSettings))
	p.SetAndroidSettings(RecaptchaenterpriseKeyAndroidSettingsToProto(resource.AndroidSettings))
	p.SetIosSettings(RecaptchaenterpriseKeyIosSettingsToProto(resource.IosSettings))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetTestingOptions(RecaptchaenterpriseKeyTestingOptionsToProto(resource.TestingOptions))
	p.SetWafSettings(RecaptchaenterpriseKeyWafSettingsToProto(resource.WafSettings))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyKey handles the gRPC request by passing it to the underlying Key Apply() method.
func (s *KeyServer) applyKey(ctx context.Context, c *recaptchaenterprise.Client, request *recaptchaenterprisepb.ApplyRecaptchaenterpriseKeyRequest) (*recaptchaenterprisepb.RecaptchaenterpriseKey, error) {
	p := ProtoToKey(request.GetResource())
	res, err := c.ApplyKey(ctx, p)
	if err != nil {
		return nil, err
	}
	r := KeyToProto(res)
	return r, nil
}

// applyRecaptchaenterpriseKey handles the gRPC request by passing it to the underlying Key Apply() method.
func (s *KeyServer) ApplyRecaptchaenterpriseKey(ctx context.Context, request *recaptchaenterprisepb.ApplyRecaptchaenterpriseKeyRequest) (*recaptchaenterprisepb.RecaptchaenterpriseKey, error) {
	cl, err := createConfigKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyKey(ctx, cl, request)
}

// DeleteKey handles the gRPC request by passing it to the underlying Key Delete() method.
func (s *KeyServer) DeleteRecaptchaenterpriseKey(ctx context.Context, request *recaptchaenterprisepb.DeleteRecaptchaenterpriseKeyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteKey(ctx, ProtoToKey(request.GetResource()))

}

// ListRecaptchaenterpriseKey handles the gRPC request by passing it to the underlying KeyList() method.
func (s *KeyServer) ListRecaptchaenterpriseKey(ctx context.Context, request *recaptchaenterprisepb.ListRecaptchaenterpriseKeyRequest) (*recaptchaenterprisepb.ListRecaptchaenterpriseKeyResponse, error) {
	cl, err := createConfigKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListKey(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*recaptchaenterprisepb.RecaptchaenterpriseKey
	for _, r := range resources.Items {
		rp := KeyToProto(r)
		protos = append(protos, rp)
	}
	p := &recaptchaenterprisepb.ListRecaptchaenterpriseKeyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigKey(ctx context.Context, service_account_file string) (*recaptchaenterprise.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return recaptchaenterprise.NewClient(conf), nil
}
