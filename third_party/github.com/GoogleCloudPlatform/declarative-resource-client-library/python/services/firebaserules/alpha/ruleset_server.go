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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/firebaserules/alpha/firebaserules_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/firebaserules/alpha"
)

// RulesetServer implements the gRPC interface for Ruleset.
type RulesetServer struct{}

// ProtoToRulesetSourceLanguageEnum converts a RulesetSourceLanguageEnum enum from its proto representation.
func ProtoToFirebaserulesAlphaRulesetSourceLanguageEnum(e alphapb.FirebaserulesAlphaRulesetSourceLanguageEnum) *alpha.RulesetSourceLanguageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.FirebaserulesAlphaRulesetSourceLanguageEnum_name[int32(e)]; ok {
		e := alpha.RulesetSourceLanguageEnum(n[len("FirebaserulesAlphaRulesetSourceLanguageEnum"):])
		return &e
	}
	return nil
}

// ProtoToRulesetSource converts a RulesetSource object from its proto representation.
func ProtoToFirebaserulesAlphaRulesetSource(p *alphapb.FirebaserulesAlphaRulesetSource) *alpha.RulesetSource {
	if p == nil {
		return nil
	}
	obj := &alpha.RulesetSource{
		Language: ProtoToFirebaserulesAlphaRulesetSourceLanguageEnum(p.GetLanguage()),
	}
	for _, r := range p.GetFiles() {
		obj.Files = append(obj.Files, *ProtoToFirebaserulesAlphaRulesetSourceFiles(r))
	}
	return obj
}

// ProtoToRulesetSourceFiles converts a RulesetSourceFiles object from its proto representation.
func ProtoToFirebaserulesAlphaRulesetSourceFiles(p *alphapb.FirebaserulesAlphaRulesetSourceFiles) *alpha.RulesetSourceFiles {
	if p == nil {
		return nil
	}
	obj := &alpha.RulesetSourceFiles{
		Content:     dcl.StringOrNil(p.GetContent()),
		Name:        dcl.StringOrNil(p.GetName()),
		Fingerprint: dcl.StringOrNil(p.GetFingerprint()),
	}
	return obj
}

// ProtoToRulesetMetadata converts a RulesetMetadata object from its proto representation.
func ProtoToFirebaserulesAlphaRulesetMetadata(p *alphapb.FirebaserulesAlphaRulesetMetadata) *alpha.RulesetMetadata {
	if p == nil {
		return nil
	}
	obj := &alpha.RulesetMetadata{}
	for _, r := range p.GetServices() {
		obj.Services = append(obj.Services, r)
	}
	return obj
}

// ProtoToRuleset converts a Ruleset resource from its proto representation.
func ProtoToRuleset(p *alphapb.FirebaserulesAlphaRuleset) *alpha.Ruleset {
	obj := &alpha.Ruleset{
		Name:       dcl.StringOrNil(p.GetName()),
		Source:     ProtoToFirebaserulesAlphaRulesetSource(p.GetSource()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
		Metadata:   ProtoToFirebaserulesAlphaRulesetMetadata(p.GetMetadata()),
		Project:    dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// RulesetSourceLanguageEnumToProto converts a RulesetSourceLanguageEnum enum to its proto representation.
func FirebaserulesAlphaRulesetSourceLanguageEnumToProto(e *alpha.RulesetSourceLanguageEnum) alphapb.FirebaserulesAlphaRulesetSourceLanguageEnum {
	if e == nil {
		return alphapb.FirebaserulesAlphaRulesetSourceLanguageEnum(0)
	}
	if v, ok := alphapb.FirebaserulesAlphaRulesetSourceLanguageEnum_value["RulesetSourceLanguageEnum"+string(*e)]; ok {
		return alphapb.FirebaserulesAlphaRulesetSourceLanguageEnum(v)
	}
	return alphapb.FirebaserulesAlphaRulesetSourceLanguageEnum(0)
}

// RulesetSourceToProto converts a RulesetSource object to its proto representation.
func FirebaserulesAlphaRulesetSourceToProto(o *alpha.RulesetSource) *alphapb.FirebaserulesAlphaRulesetSource {
	if o == nil {
		return nil
	}
	p := &alphapb.FirebaserulesAlphaRulesetSource{}
	p.SetLanguage(FirebaserulesAlphaRulesetSourceLanguageEnumToProto(o.Language))
	sFiles := make([]*alphapb.FirebaserulesAlphaRulesetSourceFiles, len(o.Files))
	for i, r := range o.Files {
		sFiles[i] = FirebaserulesAlphaRulesetSourceFilesToProto(&r)
	}
	p.SetFiles(sFiles)
	return p
}

// RulesetSourceFilesToProto converts a RulesetSourceFiles object to its proto representation.
func FirebaserulesAlphaRulesetSourceFilesToProto(o *alpha.RulesetSourceFiles) *alphapb.FirebaserulesAlphaRulesetSourceFiles {
	if o == nil {
		return nil
	}
	p := &alphapb.FirebaserulesAlphaRulesetSourceFiles{}
	p.SetContent(dcl.ValueOrEmptyString(o.Content))
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetFingerprint(dcl.ValueOrEmptyString(o.Fingerprint))
	return p
}

// RulesetMetadataToProto converts a RulesetMetadata object to its proto representation.
func FirebaserulesAlphaRulesetMetadataToProto(o *alpha.RulesetMetadata) *alphapb.FirebaserulesAlphaRulesetMetadata {
	if o == nil {
		return nil
	}
	p := &alphapb.FirebaserulesAlphaRulesetMetadata{}
	sServices := make([]string, len(o.Services))
	for i, r := range o.Services {
		sServices[i] = r
	}
	p.SetServices(sServices)
	return p
}

// RulesetToProto converts a Ruleset resource to its proto representation.
func RulesetToProto(resource *alpha.Ruleset) *alphapb.FirebaserulesAlphaRuleset {
	p := &alphapb.FirebaserulesAlphaRuleset{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetSource(FirebaserulesAlphaRulesetSourceToProto(resource.Source))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetMetadata(FirebaserulesAlphaRulesetMetadataToProto(resource.Metadata))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))

	return p
}

// applyRuleset handles the gRPC request by passing it to the underlying Ruleset Apply() method.
func (s *RulesetServer) applyRuleset(ctx context.Context, c *alpha.Client, request *alphapb.ApplyFirebaserulesAlphaRulesetRequest) (*alphapb.FirebaserulesAlphaRuleset, error) {
	p := ProtoToRuleset(request.GetResource())
	res, err := c.ApplyRuleset(ctx, p)
	if err != nil {
		return nil, err
	}
	r := RulesetToProto(res)
	return r, nil
}

// applyFirebaserulesAlphaRuleset handles the gRPC request by passing it to the underlying Ruleset Apply() method.
func (s *RulesetServer) ApplyFirebaserulesAlphaRuleset(ctx context.Context, request *alphapb.ApplyFirebaserulesAlphaRulesetRequest) (*alphapb.FirebaserulesAlphaRuleset, error) {
	cl, err := createConfigRuleset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyRuleset(ctx, cl, request)
}

// DeleteRuleset handles the gRPC request by passing it to the underlying Ruleset Delete() method.
func (s *RulesetServer) DeleteFirebaserulesAlphaRuleset(ctx context.Context, request *alphapb.DeleteFirebaserulesAlphaRulesetRequest) (*emptypb.Empty, error) {

	cl, err := createConfigRuleset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteRuleset(ctx, ProtoToRuleset(request.GetResource()))

}

// ListFirebaserulesAlphaRuleset handles the gRPC request by passing it to the underlying RulesetList() method.
func (s *RulesetServer) ListFirebaserulesAlphaRuleset(ctx context.Context, request *alphapb.ListFirebaserulesAlphaRulesetRequest) (*alphapb.ListFirebaserulesAlphaRulesetResponse, error) {
	cl, err := createConfigRuleset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListRuleset(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.FirebaserulesAlphaRuleset
	for _, r := range resources.Items {
		rp := RulesetToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListFirebaserulesAlphaRulesetResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigRuleset(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
