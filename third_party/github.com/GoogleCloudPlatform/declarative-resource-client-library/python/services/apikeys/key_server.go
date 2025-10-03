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
	apikeyspb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/apikeys/apikeys_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apikeys"
)

// KeyServer implements the gRPC interface for Key.
type KeyServer struct{}

// ProtoToKeyRestrictions converts a KeyRestrictions object from its proto representation.
func ProtoToApikeysKeyRestrictions(p *apikeyspb.ApikeysKeyRestrictions) *apikeys.KeyRestrictions {
	if p == nil {
		return nil
	}
	obj := &apikeys.KeyRestrictions{
		BrowserKeyRestrictions: ProtoToApikeysKeyRestrictionsBrowserKeyRestrictions(p.GetBrowserKeyRestrictions()),
		ServerKeyRestrictions:  ProtoToApikeysKeyRestrictionsServerKeyRestrictions(p.GetServerKeyRestrictions()),
		AndroidKeyRestrictions: ProtoToApikeysKeyRestrictionsAndroidKeyRestrictions(p.GetAndroidKeyRestrictions()),
		IosKeyRestrictions:     ProtoToApikeysKeyRestrictionsIosKeyRestrictions(p.GetIosKeyRestrictions()),
	}
	for _, r := range p.GetApiTargets() {
		obj.ApiTargets = append(obj.ApiTargets, *ProtoToApikeysKeyRestrictionsApiTargets(r))
	}
	return obj
}

// ProtoToKeyRestrictionsBrowserKeyRestrictions converts a KeyRestrictionsBrowserKeyRestrictions object from its proto representation.
func ProtoToApikeysKeyRestrictionsBrowserKeyRestrictions(p *apikeyspb.ApikeysKeyRestrictionsBrowserKeyRestrictions) *apikeys.KeyRestrictionsBrowserKeyRestrictions {
	if p == nil {
		return nil
	}
	obj := &apikeys.KeyRestrictionsBrowserKeyRestrictions{}
	for _, r := range p.GetAllowedReferrers() {
		obj.AllowedReferrers = append(obj.AllowedReferrers, r)
	}
	return obj
}

// ProtoToKeyRestrictionsServerKeyRestrictions converts a KeyRestrictionsServerKeyRestrictions object from its proto representation.
func ProtoToApikeysKeyRestrictionsServerKeyRestrictions(p *apikeyspb.ApikeysKeyRestrictionsServerKeyRestrictions) *apikeys.KeyRestrictionsServerKeyRestrictions {
	if p == nil {
		return nil
	}
	obj := &apikeys.KeyRestrictionsServerKeyRestrictions{}
	for _, r := range p.GetAllowedIps() {
		obj.AllowedIps = append(obj.AllowedIps, r)
	}
	return obj
}

// ProtoToKeyRestrictionsAndroidKeyRestrictions converts a KeyRestrictionsAndroidKeyRestrictions object from its proto representation.
func ProtoToApikeysKeyRestrictionsAndroidKeyRestrictions(p *apikeyspb.ApikeysKeyRestrictionsAndroidKeyRestrictions) *apikeys.KeyRestrictionsAndroidKeyRestrictions {
	if p == nil {
		return nil
	}
	obj := &apikeys.KeyRestrictionsAndroidKeyRestrictions{}
	for _, r := range p.GetAllowedApplications() {
		obj.AllowedApplications = append(obj.AllowedApplications, *ProtoToApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplications(r))
	}
	return obj
}

// ProtoToKeyRestrictionsAndroidKeyRestrictionsAllowedApplications converts a KeyRestrictionsAndroidKeyRestrictionsAllowedApplications object from its proto representation.
func ProtoToApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplications(p *apikeyspb.ApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplications) *apikeys.KeyRestrictionsAndroidKeyRestrictionsAllowedApplications {
	if p == nil {
		return nil
	}
	obj := &apikeys.KeyRestrictionsAndroidKeyRestrictionsAllowedApplications{
		Sha1Fingerprint: dcl.StringOrNil(p.GetSha1Fingerprint()),
		PackageName:     dcl.StringOrNil(p.GetPackageName()),
	}
	return obj
}

// ProtoToKeyRestrictionsIosKeyRestrictions converts a KeyRestrictionsIosKeyRestrictions object from its proto representation.
func ProtoToApikeysKeyRestrictionsIosKeyRestrictions(p *apikeyspb.ApikeysKeyRestrictionsIosKeyRestrictions) *apikeys.KeyRestrictionsIosKeyRestrictions {
	if p == nil {
		return nil
	}
	obj := &apikeys.KeyRestrictionsIosKeyRestrictions{}
	for _, r := range p.GetAllowedBundleIds() {
		obj.AllowedBundleIds = append(obj.AllowedBundleIds, r)
	}
	return obj
}

// ProtoToKeyRestrictionsApiTargets converts a KeyRestrictionsApiTargets object from its proto representation.
func ProtoToApikeysKeyRestrictionsApiTargets(p *apikeyspb.ApikeysKeyRestrictionsApiTargets) *apikeys.KeyRestrictionsApiTargets {
	if p == nil {
		return nil
	}
	obj := &apikeys.KeyRestrictionsApiTargets{
		Service: dcl.StringOrNil(p.GetService()),
	}
	for _, r := range p.GetMethods() {
		obj.Methods = append(obj.Methods, r)
	}
	return obj
}

// ProtoToKey converts a Key resource from its proto representation.
func ProtoToKey(p *apikeyspb.ApikeysKey) *apikeys.Key {
	obj := &apikeys.Key{
		Name:         dcl.StringOrNil(p.GetName()),
		DisplayName:  dcl.StringOrNil(p.GetDisplayName()),
		KeyString:    dcl.StringOrNil(p.GetKeyString()),
		Uid:          dcl.StringOrNil(p.GetUid()),
		Restrictions: ProtoToApikeysKeyRestrictions(p.GetRestrictions()),
		Project:      dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// KeyRestrictionsToProto converts a KeyRestrictions object to its proto representation.
func ApikeysKeyRestrictionsToProto(o *apikeys.KeyRestrictions) *apikeyspb.ApikeysKeyRestrictions {
	if o == nil {
		return nil
	}
	p := &apikeyspb.ApikeysKeyRestrictions{}
	p.SetBrowserKeyRestrictions(ApikeysKeyRestrictionsBrowserKeyRestrictionsToProto(o.BrowserKeyRestrictions))
	p.SetServerKeyRestrictions(ApikeysKeyRestrictionsServerKeyRestrictionsToProto(o.ServerKeyRestrictions))
	p.SetAndroidKeyRestrictions(ApikeysKeyRestrictionsAndroidKeyRestrictionsToProto(o.AndroidKeyRestrictions))
	p.SetIosKeyRestrictions(ApikeysKeyRestrictionsIosKeyRestrictionsToProto(o.IosKeyRestrictions))
	sApiTargets := make([]*apikeyspb.ApikeysKeyRestrictionsApiTargets, len(o.ApiTargets))
	for i, r := range o.ApiTargets {
		sApiTargets[i] = ApikeysKeyRestrictionsApiTargetsToProto(&r)
	}
	p.SetApiTargets(sApiTargets)
	return p
}

// KeyRestrictionsBrowserKeyRestrictionsToProto converts a KeyRestrictionsBrowserKeyRestrictions object to its proto representation.
func ApikeysKeyRestrictionsBrowserKeyRestrictionsToProto(o *apikeys.KeyRestrictionsBrowserKeyRestrictions) *apikeyspb.ApikeysKeyRestrictionsBrowserKeyRestrictions {
	if o == nil {
		return nil
	}
	p := &apikeyspb.ApikeysKeyRestrictionsBrowserKeyRestrictions{}
	sAllowedReferrers := make([]string, len(o.AllowedReferrers))
	for i, r := range o.AllowedReferrers {
		sAllowedReferrers[i] = r
	}
	p.SetAllowedReferrers(sAllowedReferrers)
	return p
}

// KeyRestrictionsServerKeyRestrictionsToProto converts a KeyRestrictionsServerKeyRestrictions object to its proto representation.
func ApikeysKeyRestrictionsServerKeyRestrictionsToProto(o *apikeys.KeyRestrictionsServerKeyRestrictions) *apikeyspb.ApikeysKeyRestrictionsServerKeyRestrictions {
	if o == nil {
		return nil
	}
	p := &apikeyspb.ApikeysKeyRestrictionsServerKeyRestrictions{}
	sAllowedIps := make([]string, len(o.AllowedIps))
	for i, r := range o.AllowedIps {
		sAllowedIps[i] = r
	}
	p.SetAllowedIps(sAllowedIps)
	return p
}

// KeyRestrictionsAndroidKeyRestrictionsToProto converts a KeyRestrictionsAndroidKeyRestrictions object to its proto representation.
func ApikeysKeyRestrictionsAndroidKeyRestrictionsToProto(o *apikeys.KeyRestrictionsAndroidKeyRestrictions) *apikeyspb.ApikeysKeyRestrictionsAndroidKeyRestrictions {
	if o == nil {
		return nil
	}
	p := &apikeyspb.ApikeysKeyRestrictionsAndroidKeyRestrictions{}
	sAllowedApplications := make([]*apikeyspb.ApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplications, len(o.AllowedApplications))
	for i, r := range o.AllowedApplications {
		sAllowedApplications[i] = ApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsToProto(&r)
	}
	p.SetAllowedApplications(sAllowedApplications)
	return p
}

// KeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsToProto converts a KeyRestrictionsAndroidKeyRestrictionsAllowedApplications object to its proto representation.
func ApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsToProto(o *apikeys.KeyRestrictionsAndroidKeyRestrictionsAllowedApplications) *apikeyspb.ApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplications {
	if o == nil {
		return nil
	}
	p := &apikeyspb.ApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplications{}
	p.SetSha1Fingerprint(dcl.ValueOrEmptyString(o.Sha1Fingerprint))
	p.SetPackageName(dcl.ValueOrEmptyString(o.PackageName))
	return p
}

// KeyRestrictionsIosKeyRestrictionsToProto converts a KeyRestrictionsIosKeyRestrictions object to its proto representation.
func ApikeysKeyRestrictionsIosKeyRestrictionsToProto(o *apikeys.KeyRestrictionsIosKeyRestrictions) *apikeyspb.ApikeysKeyRestrictionsIosKeyRestrictions {
	if o == nil {
		return nil
	}
	p := &apikeyspb.ApikeysKeyRestrictionsIosKeyRestrictions{}
	sAllowedBundleIds := make([]string, len(o.AllowedBundleIds))
	for i, r := range o.AllowedBundleIds {
		sAllowedBundleIds[i] = r
	}
	p.SetAllowedBundleIds(sAllowedBundleIds)
	return p
}

// KeyRestrictionsApiTargetsToProto converts a KeyRestrictionsApiTargets object to its proto representation.
func ApikeysKeyRestrictionsApiTargetsToProto(o *apikeys.KeyRestrictionsApiTargets) *apikeyspb.ApikeysKeyRestrictionsApiTargets {
	if o == nil {
		return nil
	}
	p := &apikeyspb.ApikeysKeyRestrictionsApiTargets{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	sMethods := make([]string, len(o.Methods))
	for i, r := range o.Methods {
		sMethods[i] = r
	}
	p.SetMethods(sMethods)
	return p
}

// KeyToProto converts a Key resource to its proto representation.
func KeyToProto(resource *apikeys.Key) *apikeyspb.ApikeysKey {
	p := &apikeyspb.ApikeysKey{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetKeyString(dcl.ValueOrEmptyString(resource.KeyString))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetRestrictions(ApikeysKeyRestrictionsToProto(resource.Restrictions))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))

	return p
}

// applyKey handles the gRPC request by passing it to the underlying Key Apply() method.
func (s *KeyServer) applyKey(ctx context.Context, c *apikeys.Client, request *apikeyspb.ApplyApikeysKeyRequest) (*apikeyspb.ApikeysKey, error) {
	p := ProtoToKey(request.GetResource())
	res, err := c.ApplyKey(ctx, p)
	if err != nil {
		return nil, err
	}
	r := KeyToProto(res)
	return r, nil
}

// applyApikeysKey handles the gRPC request by passing it to the underlying Key Apply() method.
func (s *KeyServer) ApplyApikeysKey(ctx context.Context, request *apikeyspb.ApplyApikeysKeyRequest) (*apikeyspb.ApikeysKey, error) {
	cl, err := createConfigKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyKey(ctx, cl, request)
}

// DeleteKey handles the gRPC request by passing it to the underlying Key Delete() method.
func (s *KeyServer) DeleteApikeysKey(ctx context.Context, request *apikeyspb.DeleteApikeysKeyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteKey(ctx, ProtoToKey(request.GetResource()))

}

// ListApikeysKey handles the gRPC request by passing it to the underlying KeyList() method.
func (s *KeyServer) ListApikeysKey(ctx context.Context, request *apikeyspb.ListApikeysKeyRequest) (*apikeyspb.ListApikeysKeyResponse, error) {
	cl, err := createConfigKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListKey(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*apikeyspb.ApikeysKey
	for _, r := range resources.Items {
		rp := KeyToProto(r)
		protos = append(protos, rp)
	}
	p := &apikeyspb.ListApikeysKeyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigKey(ctx context.Context, service_account_file string) (*apikeys.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return apikeys.NewClient(conf), nil
}
