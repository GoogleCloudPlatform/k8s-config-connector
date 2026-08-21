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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/binaryauthorization/alpha/binaryauthorization_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization/alpha"
)

// AttestorServer implements the gRPC interface for Attestor.
type AttestorServer struct{}

// ProtoToAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum converts a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum enum from its proto representation.
func ProtoToBinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(e alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum) *alpha.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum_name[int32(e)]; ok {
		e := alpha.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(n[len("BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToAttestorUserOwnedDrydockNote converts a AttestorUserOwnedDrydockNote object from its proto representation.
func ProtoToBinaryauthorizationAlphaAttestorUserOwnedDrydockNote(p *alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNote) *alpha.AttestorUserOwnedDrydockNote {
	if p == nil {
		return nil
	}
	obj := &alpha.AttestorUserOwnedDrydockNote{
		NoteReference:                 dcl.StringOrNil(p.GetNoteReference()),
		DelegationServiceAccountEmail: dcl.StringOrNil(p.GetDelegationServiceAccountEmail()),
	}
	for _, r := range p.GetPublicKeys() {
		obj.PublicKeys = append(obj.PublicKeys, *ProtoToBinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeys(r))
	}
	return obj
}

// ProtoToAttestorUserOwnedDrydockNotePublicKeys converts a AttestorUserOwnedDrydockNotePublicKeys object from its proto representation.
func ProtoToBinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeys(p *alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeys) *alpha.AttestorUserOwnedDrydockNotePublicKeys {
	if p == nil {
		return nil
	}
	obj := &alpha.AttestorUserOwnedDrydockNotePublicKeys{
		Comment:                  dcl.StringOrNil(p.GetComment()),
		Id:                       dcl.StringOrNil(p.GetId()),
		AsciiArmoredPgpPublicKey: dcl.StringOrNil(p.GetAsciiArmoredPgpPublicKey()),
		PkixPublicKey:            ProtoToBinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(p.GetPkixPublicKey()),
	}
	return obj
}

// ProtoToAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey converts a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey object from its proto representation.
func ProtoToBinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(p *alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) *alpha.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
	if p == nil {
		return nil
	}
	obj := &alpha.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{
		PublicKeyPem:       dcl.StringOrNil(p.GetPublicKeyPem()),
		SignatureAlgorithm: ProtoToBinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(p.GetSignatureAlgorithm()),
	}
	return obj
}

// ProtoToAttestor converts a Attestor resource from its proto representation.
func ProtoToAttestor(p *alphapb.BinaryauthorizationAlphaAttestor) *alpha.Attestor {
	obj := &alpha.Attestor{
		Name:                 dcl.StringOrNil(p.GetName()),
		Description:          dcl.StringOrNil(p.GetDescription()),
		UserOwnedDrydockNote: ProtoToBinaryauthorizationAlphaAttestorUserOwnedDrydockNote(p.GetUserOwnedDrydockNote()),
		UpdateTime:           dcl.StringOrNil(p.GetUpdateTime()),
		Project:              dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumToProto converts a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum enum to its proto representation.
func BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumToProto(e *alpha.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum) alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum {
	if e == nil {
		return alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(0)
	}
	if v, ok := alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum_value["AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum"+string(*e)]; ok {
		return alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(v)
	}
	return alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(0)
}

// AttestorUserOwnedDrydockNoteToProto converts a AttestorUserOwnedDrydockNote object to its proto representation.
func BinaryauthorizationAlphaAttestorUserOwnedDrydockNoteToProto(o *alpha.AttestorUserOwnedDrydockNote) *alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNote {
	if o == nil {
		return nil
	}
	p := &alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNote{}
	p.SetNoteReference(dcl.ValueOrEmptyString(o.NoteReference))
	p.SetDelegationServiceAccountEmail(dcl.ValueOrEmptyString(o.DelegationServiceAccountEmail))
	sPublicKeys := make([]*alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeys, len(o.PublicKeys))
	for i, r := range o.PublicKeys {
		sPublicKeys[i] = BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysToProto(&r)
	}
	p.SetPublicKeys(sPublicKeys)
	return p
}

// AttestorUserOwnedDrydockNotePublicKeysToProto converts a AttestorUserOwnedDrydockNotePublicKeys object to its proto representation.
func BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysToProto(o *alpha.AttestorUserOwnedDrydockNotePublicKeys) *alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeys {
	if o == nil {
		return nil
	}
	p := &alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeys{}
	p.SetComment(dcl.ValueOrEmptyString(o.Comment))
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetAsciiArmoredPgpPublicKey(dcl.ValueOrEmptyString(o.AsciiArmoredPgpPublicKey))
	p.SetPkixPublicKey(BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyToProto(o.PkixPublicKey))
	return p
}

// AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyToProto converts a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey object to its proto representation.
func BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyToProto(o *alpha.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) *alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
	if o == nil {
		return nil
	}
	p := &alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{}
	p.SetPublicKeyPem(dcl.ValueOrEmptyString(o.PublicKeyPem))
	p.SetSignatureAlgorithm(BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumToProto(o.SignatureAlgorithm))
	return p
}

// AttestorToProto converts a Attestor resource to its proto representation.
func AttestorToProto(resource *alpha.Attestor) *alphapb.BinaryauthorizationAlphaAttestor {
	p := &alphapb.BinaryauthorizationAlphaAttestor{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetUserOwnedDrydockNote(BinaryauthorizationAlphaAttestorUserOwnedDrydockNoteToProto(resource.UserOwnedDrydockNote))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))

	return p
}

// applyAttestor handles the gRPC request by passing it to the underlying Attestor Apply() method.
func (s *AttestorServer) applyAttestor(ctx context.Context, c *alpha.Client, request *alphapb.ApplyBinaryauthorizationAlphaAttestorRequest) (*alphapb.BinaryauthorizationAlphaAttestor, error) {
	p := ProtoToAttestor(request.GetResource())
	res, err := c.ApplyAttestor(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AttestorToProto(res)
	return r, nil
}

// applyBinaryauthorizationAlphaAttestor handles the gRPC request by passing it to the underlying Attestor Apply() method.
func (s *AttestorServer) ApplyBinaryauthorizationAlphaAttestor(ctx context.Context, request *alphapb.ApplyBinaryauthorizationAlphaAttestorRequest) (*alphapb.BinaryauthorizationAlphaAttestor, error) {
	cl, err := createConfigAttestor(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyAttestor(ctx, cl, request)
}

// DeleteAttestor handles the gRPC request by passing it to the underlying Attestor Delete() method.
func (s *AttestorServer) DeleteBinaryauthorizationAlphaAttestor(ctx context.Context, request *alphapb.DeleteBinaryauthorizationAlphaAttestorRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAttestor(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAttestor(ctx, ProtoToAttestor(request.GetResource()))

}

// ListBinaryauthorizationAlphaAttestor handles the gRPC request by passing it to the underlying AttestorList() method.
func (s *AttestorServer) ListBinaryauthorizationAlphaAttestor(ctx context.Context, request *alphapb.ListBinaryauthorizationAlphaAttestorRequest) (*alphapb.ListBinaryauthorizationAlphaAttestorResponse, error) {
	cl, err := createConfigAttestor(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAttestor(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.BinaryauthorizationAlphaAttestor
	for _, r := range resources.Items {
		rp := AttestorToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListBinaryauthorizationAlphaAttestorResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigAttestor(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
