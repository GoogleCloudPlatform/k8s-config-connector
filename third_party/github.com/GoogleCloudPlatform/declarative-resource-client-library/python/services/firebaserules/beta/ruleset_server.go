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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/firebaserules/beta/firebaserules_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/firebaserules/beta"
)

// RulesetServer implements the gRPC interface for Ruleset.
type RulesetServer struct{}

// ProtoToRulesetSourceLanguageEnum converts a RulesetSourceLanguageEnum enum from its proto representation.
func ProtoToFirebaserulesBetaRulesetSourceLanguageEnum(e betapb.FirebaserulesBetaRulesetSourceLanguageEnum) *beta.RulesetSourceLanguageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.FirebaserulesBetaRulesetSourceLanguageEnum_name[int32(e)]; ok {
		e := beta.RulesetSourceLanguageEnum(n[len("FirebaserulesBetaRulesetSourceLanguageEnum"):])
		return &e
	}
	return nil
}

// ProtoToRulesetSource converts a RulesetSource object from its proto representation.
func ProtoToFirebaserulesBetaRulesetSource(p *betapb.FirebaserulesBetaRulesetSource) *beta.RulesetSource {
	if p == nil {
		return nil
	}
	obj := &beta.RulesetSource{
		Language: ProtoToFirebaserulesBetaRulesetSourceLanguageEnum(p.GetLanguage()),
	}
	for _, r := range p.GetFiles() {
		obj.Files = append(obj.Files, *ProtoToFirebaserulesBetaRulesetSourceFiles(r))
	}
	return obj
}

// ProtoToRulesetSourceFiles converts a RulesetSourceFiles object from its proto representation.
func ProtoToFirebaserulesBetaRulesetSourceFiles(p *betapb.FirebaserulesBetaRulesetSourceFiles) *beta.RulesetSourceFiles {
	if p == nil {
		return nil
	}
	obj := &beta.RulesetSourceFiles{
		Content:     dcl.StringOrNil(p.GetContent()),
		Name:        dcl.StringOrNil(p.GetName()),
		Fingerprint: dcl.StringOrNil(p.GetFingerprint()),
	}
	return obj
}

// ProtoToRulesetMetadata converts a RulesetMetadata object from its proto representation.
func ProtoToFirebaserulesBetaRulesetMetadata(p *betapb.FirebaserulesBetaRulesetMetadata) *beta.RulesetMetadata {
	if p == nil {
		return nil
	}
	obj := &beta.RulesetMetadata{}
	for _, r := range p.GetServices() {
		obj.Services = append(obj.Services, r)
	}
	return obj
}

// ProtoToRuleset converts a Ruleset resource from its proto representation.
func ProtoToRuleset(p *betapb.FirebaserulesBetaRuleset) *beta.Ruleset {
	obj := &beta.Ruleset{
		Name:       dcl.StringOrNil(p.GetName()),
		Source:     ProtoToFirebaserulesBetaRulesetSource(p.GetSource()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
		Metadata:   ProtoToFirebaserulesBetaRulesetMetadata(p.GetMetadata()),
		Project:    dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// RulesetSourceLanguageEnumToProto converts a RulesetSourceLanguageEnum enum to its proto representation.
func FirebaserulesBetaRulesetSourceLanguageEnumToProto(e *beta.RulesetSourceLanguageEnum) betapb.FirebaserulesBetaRulesetSourceLanguageEnum {
	if e == nil {
		return betapb.FirebaserulesBetaRulesetSourceLanguageEnum(0)
	}
	if v, ok := betapb.FirebaserulesBetaRulesetSourceLanguageEnum_value["RulesetSourceLanguageEnum"+string(*e)]; ok {
		return betapb.FirebaserulesBetaRulesetSourceLanguageEnum(v)
	}
	return betapb.FirebaserulesBetaRulesetSourceLanguageEnum(0)
}

// RulesetSourceToProto converts a RulesetSource object to its proto representation.
func FirebaserulesBetaRulesetSourceToProto(o *beta.RulesetSource) *betapb.FirebaserulesBetaRulesetSource {
	if o == nil {
		return nil
	}
	p := &betapb.FirebaserulesBetaRulesetSource{}
	p.SetLanguage(FirebaserulesBetaRulesetSourceLanguageEnumToProto(o.Language))
	sFiles := make([]*betapb.FirebaserulesBetaRulesetSourceFiles, len(o.Files))
	for i, r := range o.Files {
		sFiles[i] = FirebaserulesBetaRulesetSourceFilesToProto(&r)
	}
	p.SetFiles(sFiles)
	return p
}

// RulesetSourceFilesToProto converts a RulesetSourceFiles object to its proto representation.
func FirebaserulesBetaRulesetSourceFilesToProto(o *beta.RulesetSourceFiles) *betapb.FirebaserulesBetaRulesetSourceFiles {
	if o == nil {
		return nil
	}
	p := &betapb.FirebaserulesBetaRulesetSourceFiles{}
	p.SetContent(dcl.ValueOrEmptyString(o.Content))
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetFingerprint(dcl.ValueOrEmptyString(o.Fingerprint))
	return p
}

// RulesetMetadataToProto converts a RulesetMetadata object to its proto representation.
func FirebaserulesBetaRulesetMetadataToProto(o *beta.RulesetMetadata) *betapb.FirebaserulesBetaRulesetMetadata {
	if o == nil {
		return nil
	}
	p := &betapb.FirebaserulesBetaRulesetMetadata{}
	sServices := make([]string, len(o.Services))
	for i, r := range o.Services {
		sServices[i] = r
	}
	p.SetServices(sServices)
	return p
}

// RulesetToProto converts a Ruleset resource to its proto representation.
func RulesetToProto(resource *beta.Ruleset) *betapb.FirebaserulesBetaRuleset {
	p := &betapb.FirebaserulesBetaRuleset{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetSource(FirebaserulesBetaRulesetSourceToProto(resource.Source))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetMetadata(FirebaserulesBetaRulesetMetadataToProto(resource.Metadata))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))

	return p
}

// applyRuleset handles the gRPC request by passing it to the underlying Ruleset Apply() method.
func (s *RulesetServer) applyRuleset(ctx context.Context, c *beta.Client, request *betapb.ApplyFirebaserulesBetaRulesetRequest) (*betapb.FirebaserulesBetaRuleset, error) {
	p := ProtoToRuleset(request.GetResource())
	res, err := c.ApplyRuleset(ctx, p)
	if err != nil {
		return nil, err
	}
	r := RulesetToProto(res)
	return r, nil
}

// applyFirebaserulesBetaRuleset handles the gRPC request by passing it to the underlying Ruleset Apply() method.
func (s *RulesetServer) ApplyFirebaserulesBetaRuleset(ctx context.Context, request *betapb.ApplyFirebaserulesBetaRulesetRequest) (*betapb.FirebaserulesBetaRuleset, error) {
	cl, err := createConfigRuleset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyRuleset(ctx, cl, request)
}

// DeleteRuleset handles the gRPC request by passing it to the underlying Ruleset Delete() method.
func (s *RulesetServer) DeleteFirebaserulesBetaRuleset(ctx context.Context, request *betapb.DeleteFirebaserulesBetaRulesetRequest) (*emptypb.Empty, error) {

	cl, err := createConfigRuleset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteRuleset(ctx, ProtoToRuleset(request.GetResource()))

}

// ListFirebaserulesBetaRuleset handles the gRPC request by passing it to the underlying RulesetList() method.
func (s *RulesetServer) ListFirebaserulesBetaRuleset(ctx context.Context, request *betapb.ListFirebaserulesBetaRulesetRequest) (*betapb.ListFirebaserulesBetaRulesetResponse, error) {
	cl, err := createConfigRuleset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListRuleset(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.FirebaserulesBetaRuleset
	for _, r := range resources.Items {
		rp := RulesetToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListFirebaserulesBetaRulesetResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigRuleset(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
