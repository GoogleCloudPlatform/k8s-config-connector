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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/apikeys/alpha/apikeys_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apikeys/alpha"
)

// KeyServer implements the gRPC interface for Key.
type KeyServer struct{}

// ProtoToKeyRestrictions converts a KeyRestrictions object from its proto representation.
func ProtoToApikeysAlphaKeyRestrictions(p *alphapb.ApikeysAlphaKeyRestrictions) *alpha.KeyRestrictions {
	if p == nil {
		return nil
	}
	obj := &alpha.KeyRestrictions{
		BrowserKeyRestrictions: ProtoToApikeysAlphaKeyRestrictionsBrowserKeyRestrictions(p.GetBrowserKeyRestrictions()),
		ServerKeyRestrictions:  ProtoToApikeysAlphaKeyRestrictionsServerKeyRestrictions(p.GetServerKeyRestrictions()),
		AndroidKeyRestrictions: ProtoToApikeysAlphaKeyRestrictionsAndroidKeyRestrictions(p.GetAndroidKeyRestrictions()),
		IosKeyRestrictions:     ProtoToApikeysAlphaKeyRestrictionsIosKeyRestrictions(p.GetIosKeyRestrictions()),
	}
	for _, r := range p.GetApiTargets() {
		obj.ApiTargets = append(obj.ApiTargets, *ProtoToApikeysAlphaKeyRestrictionsApiTargets(r))
	}
	return obj
}

// ProtoToKeyRestrictionsBrowserKeyRestrictions converts a KeyRestrictionsBrowserKeyRestrictions object from its proto representation.
func ProtoToApikeysAlphaKeyRestrictionsBrowserKeyRestrictions(p *alphapb.ApikeysAlphaKeyRestrictionsBrowserKeyRestrictions) *alpha.KeyRestrictionsBrowserKeyRestrictions {
	if p == nil {
		return nil
	}
	obj := &alpha.KeyRestrictionsBrowserKeyRestrictions{}
	for _, r := range p.GetAllowedReferrers() {
		obj.AllowedReferrers = append(obj.AllowedReferrers, r)
	}
	return obj
}

// ProtoToKeyRestrictionsServerKeyRestrictions converts a KeyRestrictionsServerKeyRestrictions object from its proto representation.
func ProtoToApikeysAlphaKeyRestrictionsServerKeyRestrictions(p *alphapb.ApikeysAlphaKeyRestrictionsServerKeyRestrictions) *alpha.KeyRestrictionsServerKeyRestrictions {
	if p == nil {
		return nil
	}
	obj := &alpha.KeyRestrictionsServerKeyRestrictions{}
	for _, r := range p.GetAllowedIps() {
		obj.AllowedIps = append(obj.AllowedIps, r)
	}
	return obj
}

// ProtoToKeyRestrictionsAndroidKeyRestrictions converts a KeyRestrictionsAndroidKeyRestrictions object from its proto representation.
func ProtoToApikeysAlphaKeyRestrictionsAndroidKeyRestrictions(p *alphapb.ApikeysAlphaKeyRestrictionsAndroidKeyRestrictions) *alpha.KeyRestrictionsAndroidKeyRestrictions {
	if p == nil {
		return nil
	}
	obj := &alpha.KeyRestrictionsAndroidKeyRestrictions{}
	for _, r := range p.GetAllowedApplications() {
		obj.AllowedApplications = append(obj.AllowedApplications, *ProtoToApikeysAlphaKeyRestrictionsAndroidKeyRestrictionsAllowedApplications(r))
	}
	return obj
}

// ProtoToKeyRestrictionsAndroidKeyRestrictionsAllowedApplications converts a KeyRestrictionsAndroidKeyRestrictionsAllowedApplications object from its proto representation.
func ProtoToApikeysAlphaKeyRestrictionsAndroidKeyRestrictionsAllowedApplications(p *alphapb.ApikeysAlphaKeyRestrictionsAndroidKeyRestrictionsAllowedApplications) *alpha.KeyRestrictionsAndroidKeyRestrictionsAllowedApplications {
	if p == nil {
		return nil
	}
	obj := &alpha.KeyRestrictionsAndroidKeyRestrictionsAllowedApplications{
		Sha1Fingerprint: dcl.StringOrNil(p.GetSha1Fingerprint()),
		PackageName:     dcl.StringOrNil(p.GetPackageName()),
	}
	return obj
}

// ProtoToKeyRestrictionsIosKeyRestrictions converts a KeyRestrictionsIosKeyRestrictions object from its proto representation.
func ProtoToApikeysAlphaKeyRestrictionsIosKeyRestrictions(p *alphapb.ApikeysAlphaKeyRestrictionsIosKeyRestrictions) *alpha.KeyRestrictionsIosKeyRestrictions {
	if p == nil {
		return nil
	}
	obj := &alpha.KeyRestrictionsIosKeyRestrictions{}
	for _, r := range p.GetAllowedBundleIds() {
		obj.AllowedBundleIds = append(obj.AllowedBundleIds, r)
	}
	return obj
}

// ProtoToKeyRestrictionsApiTargets converts a KeyRestrictionsApiTargets object from its proto representation.
func ProtoToApikeysAlphaKeyRestrictionsApiTargets(p *alphapb.ApikeysAlphaKeyRestrictionsApiTargets) *alpha.KeyRestrictionsApiTargets {
	if p == nil {
		return nil
	}
	obj := &alpha.KeyRestrictionsApiTargets{
		Service: dcl.StringOrNil(p.GetService()),
	}
	for _, r := range p.GetMethods() {
		obj.Methods = append(obj.Methods, r)
	}
	return obj
}

// ProtoToKey converts a Key resource from its proto representation.
func ProtoToKey(p *alphapb.ApikeysAlphaKey) *alpha.Key {
	obj := &alpha.Key{
		Name:         dcl.StringOrNil(p.GetName()),
		DisplayName:  dcl.StringOrNil(p.GetDisplayName()),
		KeyString:    dcl.StringOrNil(p.GetKeyString()),
		Uid:          dcl.StringOrNil(p.GetUid()),
		Restrictions: ProtoToApikeysAlphaKeyRestrictions(p.GetRestrictions()),
		Project:      dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// KeyRestrictionsToProto converts a KeyRestrictions object to its proto representation.
func ApikeysAlphaKeyRestrictionsToProto(o *alpha.KeyRestrictions) *alphapb.ApikeysAlphaKeyRestrictions {
	if o == nil {
		return nil
	}
	p := &alphapb.ApikeysAlphaKeyRestrictions{}
	p.SetBrowserKeyRestrictions(ApikeysAlphaKeyRestrictionsBrowserKeyRestrictionsToProto(o.BrowserKeyRestrictions))
	p.SetServerKeyRestrictions(ApikeysAlphaKeyRestrictionsServerKeyRestrictionsToProto(o.ServerKeyRestrictions))
	p.SetAndroidKeyRestrictions(ApikeysAlphaKeyRestrictionsAndroidKeyRestrictionsToProto(o.AndroidKeyRestrictions))
	p.SetIosKeyRestrictions(ApikeysAlphaKeyRestrictionsIosKeyRestrictionsToProto(o.IosKeyRestrictions))
	sApiTargets := make([]*alphapb.ApikeysAlphaKeyRestrictionsApiTargets, len(o.ApiTargets))
	for i, r := range o.ApiTargets {
		sApiTargets[i] = ApikeysAlphaKeyRestrictionsApiTargetsToProto(&r)
	}
	p.SetApiTargets(sApiTargets)
	return p
}

// KeyRestrictionsBrowserKeyRestrictionsToProto converts a KeyRestrictionsBrowserKeyRestrictions object to its proto representation.
func ApikeysAlphaKeyRestrictionsBrowserKeyRestrictionsToProto(o *alpha.KeyRestrictionsBrowserKeyRestrictions) *alphapb.ApikeysAlphaKeyRestrictionsBrowserKeyRestrictions {
	if o == nil {
		return nil
	}
	p := &alphapb.ApikeysAlphaKeyRestrictionsBrowserKeyRestrictions{}
	sAllowedReferrers := make([]string, len(o.AllowedReferrers))
	for i, r := range o.AllowedReferrers {
		sAllowedReferrers[i] = r
	}
	p.SetAllowedReferrers(sAllowedReferrers)
	return p
}

// KeyRestrictionsServerKeyRestrictionsToProto converts a KeyRestrictionsServerKeyRestrictions object to its proto representation.
func ApikeysAlphaKeyRestrictionsServerKeyRestrictionsToProto(o *alpha.KeyRestrictionsServerKeyRestrictions) *alphapb.ApikeysAlphaKeyRestrictionsServerKeyRestrictions {
	if o == nil {
		return nil
	}
	p := &alphapb.ApikeysAlphaKeyRestrictionsServerKeyRestrictions{}
	sAllowedIps := make([]string, len(o.AllowedIps))
	for i, r := range o.AllowedIps {
		sAllowedIps[i] = r
	}
	p.SetAllowedIps(sAllowedIps)
	return p
}

// KeyRestrictionsAndroidKeyRestrictionsToProto converts a KeyRestrictionsAndroidKeyRestrictions object to its proto representation.
func ApikeysAlphaKeyRestrictionsAndroidKeyRestrictionsToProto(o *alpha.KeyRestrictionsAndroidKeyRestrictions) *alphapb.ApikeysAlphaKeyRestrictionsAndroidKeyRestrictions {
	if o == nil {
		return nil
	}
	p := &alphapb.ApikeysAlphaKeyRestrictionsAndroidKeyRestrictions{}
	sAllowedApplications := make([]*alphapb.ApikeysAlphaKeyRestrictionsAndroidKeyRestrictionsAllowedApplications, len(o.AllowedApplications))
	for i, r := range o.AllowedApplications {
		sAllowedApplications[i] = ApikeysAlphaKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsToProto(&r)
	}
	p.SetAllowedApplications(sAllowedApplications)
	return p
}

// KeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsToProto converts a KeyRestrictionsAndroidKeyRestrictionsAllowedApplications object to its proto representation.
func ApikeysAlphaKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsToProto(o *alpha.KeyRestrictionsAndroidKeyRestrictionsAllowedApplications) *alphapb.ApikeysAlphaKeyRestrictionsAndroidKeyRestrictionsAllowedApplications {
	if o == nil {
		return nil
	}
	p := &alphapb.ApikeysAlphaKeyRestrictionsAndroidKeyRestrictionsAllowedApplications{}
	p.SetSha1Fingerprint(dcl.ValueOrEmptyString(o.Sha1Fingerprint))
	p.SetPackageName(dcl.ValueOrEmptyString(o.PackageName))
	return p
}

// KeyRestrictionsIosKeyRestrictionsToProto converts a KeyRestrictionsIosKeyRestrictions object to its proto representation.
func ApikeysAlphaKeyRestrictionsIosKeyRestrictionsToProto(o *alpha.KeyRestrictionsIosKeyRestrictions) *alphapb.ApikeysAlphaKeyRestrictionsIosKeyRestrictions {
	if o == nil {
		return nil
	}
	p := &alphapb.ApikeysAlphaKeyRestrictionsIosKeyRestrictions{}
	sAllowedBundleIds := make([]string, len(o.AllowedBundleIds))
	for i, r := range o.AllowedBundleIds {
		sAllowedBundleIds[i] = r
	}
	p.SetAllowedBundleIds(sAllowedBundleIds)
	return p
}

// KeyRestrictionsApiTargetsToProto converts a KeyRestrictionsApiTargets object to its proto representation.
func ApikeysAlphaKeyRestrictionsApiTargetsToProto(o *alpha.KeyRestrictionsApiTargets) *alphapb.ApikeysAlphaKeyRestrictionsApiTargets {
	if o == nil {
		return nil
	}
	p := &alphapb.ApikeysAlphaKeyRestrictionsApiTargets{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	sMethods := make([]string, len(o.Methods))
	for i, r := range o.Methods {
		sMethods[i] = r
	}
	p.SetMethods(sMethods)
	return p
}

// KeyToProto converts a Key resource to its proto representation.
func KeyToProto(resource *alpha.Key) *alphapb.ApikeysAlphaKey {
	p := &alphapb.ApikeysAlphaKey{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetKeyString(dcl.ValueOrEmptyString(resource.KeyString))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetRestrictions(ApikeysAlphaKeyRestrictionsToProto(resource.Restrictions))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))

	return p
}

// applyKey handles the gRPC request by passing it to the underlying Key Apply() method.
func (s *KeyServer) applyKey(ctx context.Context, c *alpha.Client, request *alphapb.ApplyApikeysAlphaKeyRequest) (*alphapb.ApikeysAlphaKey, error) {
	p := ProtoToKey(request.GetResource())
	res, err := c.ApplyKey(ctx, p)
	if err != nil {
		return nil, err
	}
	r := KeyToProto(res)
	return r, nil
}

// applyApikeysAlphaKey handles the gRPC request by passing it to the underlying Key Apply() method.
func (s *KeyServer) ApplyApikeysAlphaKey(ctx context.Context, request *alphapb.ApplyApikeysAlphaKeyRequest) (*alphapb.ApikeysAlphaKey, error) {
	cl, err := createConfigKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyKey(ctx, cl, request)
}

// DeleteKey handles the gRPC request by passing it to the underlying Key Delete() method.
func (s *KeyServer) DeleteApikeysAlphaKey(ctx context.Context, request *alphapb.DeleteApikeysAlphaKeyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteKey(ctx, ProtoToKey(request.GetResource()))

}

// ListApikeysAlphaKey handles the gRPC request by passing it to the underlying KeyList() method.
func (s *KeyServer) ListApikeysAlphaKey(ctx context.Context, request *alphapb.ListApikeysAlphaKeyRequest) (*alphapb.ListApikeysAlphaKeyResponse, error) {
	cl, err := createConfigKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListKey(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ApikeysAlphaKey
	for _, r := range resources.Items {
		rp := KeyToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListApikeysAlphaKeyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigKey(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
