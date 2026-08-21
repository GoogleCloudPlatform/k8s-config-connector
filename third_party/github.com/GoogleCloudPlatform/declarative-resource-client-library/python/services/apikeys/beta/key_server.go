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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/apikeys/beta/apikeys_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apikeys/beta"
)

// KeyServer implements the gRPC interface for Key.
type KeyServer struct{}

// ProtoToKeyRestrictions converts a KeyRestrictions object from its proto representation.
func ProtoToApikeysBetaKeyRestrictions(p *betapb.ApikeysBetaKeyRestrictions) *beta.KeyRestrictions {
	if p == nil {
		return nil
	}
	obj := &beta.KeyRestrictions{
		BrowserKeyRestrictions: ProtoToApikeysBetaKeyRestrictionsBrowserKeyRestrictions(p.GetBrowserKeyRestrictions()),
		ServerKeyRestrictions:  ProtoToApikeysBetaKeyRestrictionsServerKeyRestrictions(p.GetServerKeyRestrictions()),
		AndroidKeyRestrictions: ProtoToApikeysBetaKeyRestrictionsAndroidKeyRestrictions(p.GetAndroidKeyRestrictions()),
		IosKeyRestrictions:     ProtoToApikeysBetaKeyRestrictionsIosKeyRestrictions(p.GetIosKeyRestrictions()),
	}
	for _, r := range p.GetApiTargets() {
		obj.ApiTargets = append(obj.ApiTargets, *ProtoToApikeysBetaKeyRestrictionsApiTargets(r))
	}
	return obj
}

// ProtoToKeyRestrictionsBrowserKeyRestrictions converts a KeyRestrictionsBrowserKeyRestrictions object from its proto representation.
func ProtoToApikeysBetaKeyRestrictionsBrowserKeyRestrictions(p *betapb.ApikeysBetaKeyRestrictionsBrowserKeyRestrictions) *beta.KeyRestrictionsBrowserKeyRestrictions {
	if p == nil {
		return nil
	}
	obj := &beta.KeyRestrictionsBrowserKeyRestrictions{}
	for _, r := range p.GetAllowedReferrers() {
		obj.AllowedReferrers = append(obj.AllowedReferrers, r)
	}
	return obj
}

// ProtoToKeyRestrictionsServerKeyRestrictions converts a KeyRestrictionsServerKeyRestrictions object from its proto representation.
func ProtoToApikeysBetaKeyRestrictionsServerKeyRestrictions(p *betapb.ApikeysBetaKeyRestrictionsServerKeyRestrictions) *beta.KeyRestrictionsServerKeyRestrictions {
	if p == nil {
		return nil
	}
	obj := &beta.KeyRestrictionsServerKeyRestrictions{}
	for _, r := range p.GetAllowedIps() {
		obj.AllowedIps = append(obj.AllowedIps, r)
	}
	return obj
}

// ProtoToKeyRestrictionsAndroidKeyRestrictions converts a KeyRestrictionsAndroidKeyRestrictions object from its proto representation.
func ProtoToApikeysBetaKeyRestrictionsAndroidKeyRestrictions(p *betapb.ApikeysBetaKeyRestrictionsAndroidKeyRestrictions) *beta.KeyRestrictionsAndroidKeyRestrictions {
	if p == nil {
		return nil
	}
	obj := &beta.KeyRestrictionsAndroidKeyRestrictions{}
	for _, r := range p.GetAllowedApplications() {
		obj.AllowedApplications = append(obj.AllowedApplications, *ProtoToApikeysBetaKeyRestrictionsAndroidKeyRestrictionsAllowedApplications(r))
	}
	return obj
}

// ProtoToKeyRestrictionsAndroidKeyRestrictionsAllowedApplications converts a KeyRestrictionsAndroidKeyRestrictionsAllowedApplications object from its proto representation.
func ProtoToApikeysBetaKeyRestrictionsAndroidKeyRestrictionsAllowedApplications(p *betapb.ApikeysBetaKeyRestrictionsAndroidKeyRestrictionsAllowedApplications) *beta.KeyRestrictionsAndroidKeyRestrictionsAllowedApplications {
	if p == nil {
		return nil
	}
	obj := &beta.KeyRestrictionsAndroidKeyRestrictionsAllowedApplications{
		Sha1Fingerprint: dcl.StringOrNil(p.GetSha1Fingerprint()),
		PackageName:     dcl.StringOrNil(p.GetPackageName()),
	}
	return obj
}

// ProtoToKeyRestrictionsIosKeyRestrictions converts a KeyRestrictionsIosKeyRestrictions object from its proto representation.
func ProtoToApikeysBetaKeyRestrictionsIosKeyRestrictions(p *betapb.ApikeysBetaKeyRestrictionsIosKeyRestrictions) *beta.KeyRestrictionsIosKeyRestrictions {
	if p == nil {
		return nil
	}
	obj := &beta.KeyRestrictionsIosKeyRestrictions{}
	for _, r := range p.GetAllowedBundleIds() {
		obj.AllowedBundleIds = append(obj.AllowedBundleIds, r)
	}
	return obj
}

// ProtoToKeyRestrictionsApiTargets converts a KeyRestrictionsApiTargets object from its proto representation.
func ProtoToApikeysBetaKeyRestrictionsApiTargets(p *betapb.ApikeysBetaKeyRestrictionsApiTargets) *beta.KeyRestrictionsApiTargets {
	if p == nil {
		return nil
	}
	obj := &beta.KeyRestrictionsApiTargets{
		Service: dcl.StringOrNil(p.GetService()),
	}
	for _, r := range p.GetMethods() {
		obj.Methods = append(obj.Methods, r)
	}
	return obj
}

// ProtoToKey converts a Key resource from its proto representation.
func ProtoToKey(p *betapb.ApikeysBetaKey) *beta.Key {
	obj := &beta.Key{
		Name:         dcl.StringOrNil(p.GetName()),
		DisplayName:  dcl.StringOrNil(p.GetDisplayName()),
		KeyString:    dcl.StringOrNil(p.GetKeyString()),
		Uid:          dcl.StringOrNil(p.GetUid()),
		Restrictions: ProtoToApikeysBetaKeyRestrictions(p.GetRestrictions()),
		Project:      dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// KeyRestrictionsToProto converts a KeyRestrictions object to its proto representation.
func ApikeysBetaKeyRestrictionsToProto(o *beta.KeyRestrictions) *betapb.ApikeysBetaKeyRestrictions {
	if o == nil {
		return nil
	}
	p := &betapb.ApikeysBetaKeyRestrictions{}
	p.SetBrowserKeyRestrictions(ApikeysBetaKeyRestrictionsBrowserKeyRestrictionsToProto(o.BrowserKeyRestrictions))
	p.SetServerKeyRestrictions(ApikeysBetaKeyRestrictionsServerKeyRestrictionsToProto(o.ServerKeyRestrictions))
	p.SetAndroidKeyRestrictions(ApikeysBetaKeyRestrictionsAndroidKeyRestrictionsToProto(o.AndroidKeyRestrictions))
	p.SetIosKeyRestrictions(ApikeysBetaKeyRestrictionsIosKeyRestrictionsToProto(o.IosKeyRestrictions))
	sApiTargets := make([]*betapb.ApikeysBetaKeyRestrictionsApiTargets, len(o.ApiTargets))
	for i, r := range o.ApiTargets {
		sApiTargets[i] = ApikeysBetaKeyRestrictionsApiTargetsToProto(&r)
	}
	p.SetApiTargets(sApiTargets)
	return p
}

// KeyRestrictionsBrowserKeyRestrictionsToProto converts a KeyRestrictionsBrowserKeyRestrictions object to its proto representation.
func ApikeysBetaKeyRestrictionsBrowserKeyRestrictionsToProto(o *beta.KeyRestrictionsBrowserKeyRestrictions) *betapb.ApikeysBetaKeyRestrictionsBrowserKeyRestrictions {
	if o == nil {
		return nil
	}
	p := &betapb.ApikeysBetaKeyRestrictionsBrowserKeyRestrictions{}
	sAllowedReferrers := make([]string, len(o.AllowedReferrers))
	for i, r := range o.AllowedReferrers {
		sAllowedReferrers[i] = r
	}
	p.SetAllowedReferrers(sAllowedReferrers)
	return p
}

// KeyRestrictionsServerKeyRestrictionsToProto converts a KeyRestrictionsServerKeyRestrictions object to its proto representation.
func ApikeysBetaKeyRestrictionsServerKeyRestrictionsToProto(o *beta.KeyRestrictionsServerKeyRestrictions) *betapb.ApikeysBetaKeyRestrictionsServerKeyRestrictions {
	if o == nil {
		return nil
	}
	p := &betapb.ApikeysBetaKeyRestrictionsServerKeyRestrictions{}
	sAllowedIps := make([]string, len(o.AllowedIps))
	for i, r := range o.AllowedIps {
		sAllowedIps[i] = r
	}
	p.SetAllowedIps(sAllowedIps)
	return p
}

// KeyRestrictionsAndroidKeyRestrictionsToProto converts a KeyRestrictionsAndroidKeyRestrictions object to its proto representation.
func ApikeysBetaKeyRestrictionsAndroidKeyRestrictionsToProto(o *beta.KeyRestrictionsAndroidKeyRestrictions) *betapb.ApikeysBetaKeyRestrictionsAndroidKeyRestrictions {
	if o == nil {
		return nil
	}
	p := &betapb.ApikeysBetaKeyRestrictionsAndroidKeyRestrictions{}
	sAllowedApplications := make([]*betapb.ApikeysBetaKeyRestrictionsAndroidKeyRestrictionsAllowedApplications, len(o.AllowedApplications))
	for i, r := range o.AllowedApplications {
		sAllowedApplications[i] = ApikeysBetaKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsToProto(&r)
	}
	p.SetAllowedApplications(sAllowedApplications)
	return p
}

// KeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsToProto converts a KeyRestrictionsAndroidKeyRestrictionsAllowedApplications object to its proto representation.
func ApikeysBetaKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsToProto(o *beta.KeyRestrictionsAndroidKeyRestrictionsAllowedApplications) *betapb.ApikeysBetaKeyRestrictionsAndroidKeyRestrictionsAllowedApplications {
	if o == nil {
		return nil
	}
	p := &betapb.ApikeysBetaKeyRestrictionsAndroidKeyRestrictionsAllowedApplications{}
	p.SetSha1Fingerprint(dcl.ValueOrEmptyString(o.Sha1Fingerprint))
	p.SetPackageName(dcl.ValueOrEmptyString(o.PackageName))
	return p
}

// KeyRestrictionsIosKeyRestrictionsToProto converts a KeyRestrictionsIosKeyRestrictions object to its proto representation.
func ApikeysBetaKeyRestrictionsIosKeyRestrictionsToProto(o *beta.KeyRestrictionsIosKeyRestrictions) *betapb.ApikeysBetaKeyRestrictionsIosKeyRestrictions {
	if o == nil {
		return nil
	}
	p := &betapb.ApikeysBetaKeyRestrictionsIosKeyRestrictions{}
	sAllowedBundleIds := make([]string, len(o.AllowedBundleIds))
	for i, r := range o.AllowedBundleIds {
		sAllowedBundleIds[i] = r
	}
	p.SetAllowedBundleIds(sAllowedBundleIds)
	return p
}

// KeyRestrictionsApiTargetsToProto converts a KeyRestrictionsApiTargets object to its proto representation.
func ApikeysBetaKeyRestrictionsApiTargetsToProto(o *beta.KeyRestrictionsApiTargets) *betapb.ApikeysBetaKeyRestrictionsApiTargets {
	if o == nil {
		return nil
	}
	p := &betapb.ApikeysBetaKeyRestrictionsApiTargets{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	sMethods := make([]string, len(o.Methods))
	for i, r := range o.Methods {
		sMethods[i] = r
	}
	p.SetMethods(sMethods)
	return p
}

// KeyToProto converts a Key resource to its proto representation.
func KeyToProto(resource *beta.Key) *betapb.ApikeysBetaKey {
	p := &betapb.ApikeysBetaKey{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetKeyString(dcl.ValueOrEmptyString(resource.KeyString))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetRestrictions(ApikeysBetaKeyRestrictionsToProto(resource.Restrictions))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))

	return p
}

// applyKey handles the gRPC request by passing it to the underlying Key Apply() method.
func (s *KeyServer) applyKey(ctx context.Context, c *beta.Client, request *betapb.ApplyApikeysBetaKeyRequest) (*betapb.ApikeysBetaKey, error) {
	p := ProtoToKey(request.GetResource())
	res, err := c.ApplyKey(ctx, p)
	if err != nil {
		return nil, err
	}
	r := KeyToProto(res)
	return r, nil
}

// applyApikeysBetaKey handles the gRPC request by passing it to the underlying Key Apply() method.
func (s *KeyServer) ApplyApikeysBetaKey(ctx context.Context, request *betapb.ApplyApikeysBetaKeyRequest) (*betapb.ApikeysBetaKey, error) {
	cl, err := createConfigKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyKey(ctx, cl, request)
}

// DeleteKey handles the gRPC request by passing it to the underlying Key Delete() method.
func (s *KeyServer) DeleteApikeysBetaKey(ctx context.Context, request *betapb.DeleteApikeysBetaKeyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteKey(ctx, ProtoToKey(request.GetResource()))

}

// ListApikeysBetaKey handles the gRPC request by passing it to the underlying KeyList() method.
func (s *KeyServer) ListApikeysBetaKey(ctx context.Context, request *betapb.ListApikeysBetaKeyRequest) (*betapb.ListApikeysBetaKeyResponse, error) {
	cl, err := createConfigKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListKey(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ApikeysBetaKey
	for _, r := range resources.Items {
		rp := KeyToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListApikeysBetaKeyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigKey(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
