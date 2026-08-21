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
	firebaserulespb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/firebaserules/firebaserules_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/firebaserules"
)

// RulesetServer implements the gRPC interface for Ruleset.
type RulesetServer struct{}

// ProtoToRulesetSourceLanguageEnum converts a RulesetSourceLanguageEnum enum from its proto representation.
func ProtoToFirebaserulesRulesetSourceLanguageEnum(e firebaserulespb.FirebaserulesRulesetSourceLanguageEnum) *firebaserules.RulesetSourceLanguageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := firebaserulespb.FirebaserulesRulesetSourceLanguageEnum_name[int32(e)]; ok {
		e := firebaserules.RulesetSourceLanguageEnum(n[len("FirebaserulesRulesetSourceLanguageEnum"):])
		return &e
	}
	return nil
}

// ProtoToRulesetSource converts a RulesetSource object from its proto representation.
func ProtoToFirebaserulesRulesetSource(p *firebaserulespb.FirebaserulesRulesetSource) *firebaserules.RulesetSource {
	if p == nil {
		return nil
	}
	obj := &firebaserules.RulesetSource{
		Language: ProtoToFirebaserulesRulesetSourceLanguageEnum(p.GetLanguage()),
	}
	for _, r := range p.GetFiles() {
		obj.Files = append(obj.Files, *ProtoToFirebaserulesRulesetSourceFiles(r))
	}
	return obj
}

// ProtoToRulesetSourceFiles converts a RulesetSourceFiles object from its proto representation.
func ProtoToFirebaserulesRulesetSourceFiles(p *firebaserulespb.FirebaserulesRulesetSourceFiles) *firebaserules.RulesetSourceFiles {
	if p == nil {
		return nil
	}
	obj := &firebaserules.RulesetSourceFiles{
		Content:     dcl.StringOrNil(p.GetContent()),
		Name:        dcl.StringOrNil(p.GetName()),
		Fingerprint: dcl.StringOrNil(p.GetFingerprint()),
	}
	return obj
}

// ProtoToRulesetMetadata converts a RulesetMetadata object from its proto representation.
func ProtoToFirebaserulesRulesetMetadata(p *firebaserulespb.FirebaserulesRulesetMetadata) *firebaserules.RulesetMetadata {
	if p == nil {
		return nil
	}
	obj := &firebaserules.RulesetMetadata{}
	for _, r := range p.GetServices() {
		obj.Services = append(obj.Services, r)
	}
	return obj
}

// ProtoToRuleset converts a Ruleset resource from its proto representation.
func ProtoToRuleset(p *firebaserulespb.FirebaserulesRuleset) *firebaserules.Ruleset {
	obj := &firebaserules.Ruleset{
		Name:       dcl.StringOrNil(p.GetName()),
		Source:     ProtoToFirebaserulesRulesetSource(p.GetSource()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
		Metadata:   ProtoToFirebaserulesRulesetMetadata(p.GetMetadata()),
		Project:    dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// RulesetSourceLanguageEnumToProto converts a RulesetSourceLanguageEnum enum to its proto representation.
func FirebaserulesRulesetSourceLanguageEnumToProto(e *firebaserules.RulesetSourceLanguageEnum) firebaserulespb.FirebaserulesRulesetSourceLanguageEnum {
	if e == nil {
		return firebaserulespb.FirebaserulesRulesetSourceLanguageEnum(0)
	}
	if v, ok := firebaserulespb.FirebaserulesRulesetSourceLanguageEnum_value["RulesetSourceLanguageEnum"+string(*e)]; ok {
		return firebaserulespb.FirebaserulesRulesetSourceLanguageEnum(v)
	}
	return firebaserulespb.FirebaserulesRulesetSourceLanguageEnum(0)
}

// RulesetSourceToProto converts a RulesetSource object to its proto representation.
func FirebaserulesRulesetSourceToProto(o *firebaserules.RulesetSource) *firebaserulespb.FirebaserulesRulesetSource {
	if o == nil {
		return nil
	}
	p := &firebaserulespb.FirebaserulesRulesetSource{}
	p.SetLanguage(FirebaserulesRulesetSourceLanguageEnumToProto(o.Language))
	sFiles := make([]*firebaserulespb.FirebaserulesRulesetSourceFiles, len(o.Files))
	for i, r := range o.Files {
		sFiles[i] = FirebaserulesRulesetSourceFilesToProto(&r)
	}
	p.SetFiles(sFiles)
	return p
}

// RulesetSourceFilesToProto converts a RulesetSourceFiles object to its proto representation.
func FirebaserulesRulesetSourceFilesToProto(o *firebaserules.RulesetSourceFiles) *firebaserulespb.FirebaserulesRulesetSourceFiles {
	if o == nil {
		return nil
	}
	p := &firebaserulespb.FirebaserulesRulesetSourceFiles{}
	p.SetContent(dcl.ValueOrEmptyString(o.Content))
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetFingerprint(dcl.ValueOrEmptyString(o.Fingerprint))
	return p
}

// RulesetMetadataToProto converts a RulesetMetadata object to its proto representation.
func FirebaserulesRulesetMetadataToProto(o *firebaserules.RulesetMetadata) *firebaserulespb.FirebaserulesRulesetMetadata {
	if o == nil {
		return nil
	}
	p := &firebaserulespb.FirebaserulesRulesetMetadata{}
	sServices := make([]string, len(o.Services))
	for i, r := range o.Services {
		sServices[i] = r
	}
	p.SetServices(sServices)
	return p
}

// RulesetToProto converts a Ruleset resource to its proto representation.
func RulesetToProto(resource *firebaserules.Ruleset) *firebaserulespb.FirebaserulesRuleset {
	p := &firebaserulespb.FirebaserulesRuleset{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetSource(FirebaserulesRulesetSourceToProto(resource.Source))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetMetadata(FirebaserulesRulesetMetadataToProto(resource.Metadata))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))

	return p
}

// applyRuleset handles the gRPC request by passing it to the underlying Ruleset Apply() method.
func (s *RulesetServer) applyRuleset(ctx context.Context, c *firebaserules.Client, request *firebaserulespb.ApplyFirebaserulesRulesetRequest) (*firebaserulespb.FirebaserulesRuleset, error) {
	p := ProtoToRuleset(request.GetResource())
	res, err := c.ApplyRuleset(ctx, p)
	if err != nil {
		return nil, err
	}
	r := RulesetToProto(res)
	return r, nil
}

// applyFirebaserulesRuleset handles the gRPC request by passing it to the underlying Ruleset Apply() method.
func (s *RulesetServer) ApplyFirebaserulesRuleset(ctx context.Context, request *firebaserulespb.ApplyFirebaserulesRulesetRequest) (*firebaserulespb.FirebaserulesRuleset, error) {
	cl, err := createConfigRuleset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyRuleset(ctx, cl, request)
}

// DeleteRuleset handles the gRPC request by passing it to the underlying Ruleset Delete() method.
func (s *RulesetServer) DeleteFirebaserulesRuleset(ctx context.Context, request *firebaserulespb.DeleteFirebaserulesRulesetRequest) (*emptypb.Empty, error) {

	cl, err := createConfigRuleset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteRuleset(ctx, ProtoToRuleset(request.GetResource()))

}

// ListFirebaserulesRuleset handles the gRPC request by passing it to the underlying RulesetList() method.
func (s *RulesetServer) ListFirebaserulesRuleset(ctx context.Context, request *firebaserulespb.ListFirebaserulesRulesetRequest) (*firebaserulespb.ListFirebaserulesRulesetResponse, error) {
	cl, err := createConfigRuleset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListRuleset(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*firebaserulespb.FirebaserulesRuleset
	for _, r := range resources.Items {
		rp := RulesetToProto(r)
		protos = append(protos, rp)
	}
	p := &firebaserulespb.ListFirebaserulesRulesetResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigRuleset(ctx context.Context, service_account_file string) (*firebaserules.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return firebaserules.NewClient(conf), nil
}
