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
	binaryauthorizationpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/binaryauthorization/binaryauthorization_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization"
)

// AttestorServer implements the gRPC interface for Attestor.
type AttestorServer struct{}

// ProtoToAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum converts a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum enum from its proto representation.
func ProtoToBinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(e binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum) *binaryauthorization.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum_name[int32(e)]; ok {
		e := binaryauthorization.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(n[len("BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToAttestorUserOwnedDrydockNote converts a AttestorUserOwnedDrydockNote object from its proto representation.
func ProtoToBinaryauthorizationAttestorUserOwnedDrydockNote(p *binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNote) *binaryauthorization.AttestorUserOwnedDrydockNote {
	if p == nil {
		return nil
	}
	obj := &binaryauthorization.AttestorUserOwnedDrydockNote{
		NoteReference:                 dcl.StringOrNil(p.GetNoteReference()),
		DelegationServiceAccountEmail: dcl.StringOrNil(p.GetDelegationServiceAccountEmail()),
	}
	for _, r := range p.GetPublicKeys() {
		obj.PublicKeys = append(obj.PublicKeys, *ProtoToBinaryauthorizationAttestorUserOwnedDrydockNotePublicKeys(r))
	}
	return obj
}

// ProtoToAttestorUserOwnedDrydockNotePublicKeys converts a AttestorUserOwnedDrydockNotePublicKeys object from its proto representation.
func ProtoToBinaryauthorizationAttestorUserOwnedDrydockNotePublicKeys(p *binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeys) *binaryauthorization.AttestorUserOwnedDrydockNotePublicKeys {
	if p == nil {
		return nil
	}
	obj := &binaryauthorization.AttestorUserOwnedDrydockNotePublicKeys{
		Comment:                  dcl.StringOrNil(p.GetComment()),
		Id:                       dcl.StringOrNil(p.GetId()),
		AsciiArmoredPgpPublicKey: dcl.StringOrNil(p.GetAsciiArmoredPgpPublicKey()),
		PkixPublicKey:            ProtoToBinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(p.GetPkixPublicKey()),
	}
	return obj
}

// ProtoToAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey converts a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey object from its proto representation.
func ProtoToBinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(p *binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) *binaryauthorization.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
	if p == nil {
		return nil
	}
	obj := &binaryauthorization.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{
		PublicKeyPem:       dcl.StringOrNil(p.GetPublicKeyPem()),
		SignatureAlgorithm: ProtoToBinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(p.GetSignatureAlgorithm()),
	}
	return obj
}

// ProtoToAttestor converts a Attestor resource from its proto representation.
func ProtoToAttestor(p *binaryauthorizationpb.BinaryauthorizationAttestor) *binaryauthorization.Attestor {
	obj := &binaryauthorization.Attestor{
		Name:                 dcl.StringOrNil(p.GetName()),
		Description:          dcl.StringOrNil(p.GetDescription()),
		UserOwnedDrydockNote: ProtoToBinaryauthorizationAttestorUserOwnedDrydockNote(p.GetUserOwnedDrydockNote()),
		UpdateTime:           dcl.StringOrNil(p.GetUpdateTime()),
		Project:              dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumToProto converts a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum enum to its proto representation.
func BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumToProto(e *binaryauthorization.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum) binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum {
	if e == nil {
		return binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(0)
	}
	if v, ok := binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum_value["AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum"+string(*e)]; ok {
		return binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(v)
	}
	return binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(0)
}

// AttestorUserOwnedDrydockNoteToProto converts a AttestorUserOwnedDrydockNote object to its proto representation.
func BinaryauthorizationAttestorUserOwnedDrydockNoteToProto(o *binaryauthorization.AttestorUserOwnedDrydockNote) *binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNote {
	if o == nil {
		return nil
	}
	p := &binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNote{}
	p.SetNoteReference(dcl.ValueOrEmptyString(o.NoteReference))
	p.SetDelegationServiceAccountEmail(dcl.ValueOrEmptyString(o.DelegationServiceAccountEmail))
	sPublicKeys := make([]*binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeys, len(o.PublicKeys))
	for i, r := range o.PublicKeys {
		sPublicKeys[i] = BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysToProto(&r)
	}
	p.SetPublicKeys(sPublicKeys)
	return p
}

// AttestorUserOwnedDrydockNotePublicKeysToProto converts a AttestorUserOwnedDrydockNotePublicKeys object to its proto representation.
func BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysToProto(o *binaryauthorization.AttestorUserOwnedDrydockNotePublicKeys) *binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeys {
	if o == nil {
		return nil
	}
	p := &binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeys{}
	p.SetComment(dcl.ValueOrEmptyString(o.Comment))
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetAsciiArmoredPgpPublicKey(dcl.ValueOrEmptyString(o.AsciiArmoredPgpPublicKey))
	p.SetPkixPublicKey(BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyToProto(o.PkixPublicKey))
	return p
}

// AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyToProto converts a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey object to its proto representation.
func BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyToProto(o *binaryauthorization.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) *binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
	if o == nil {
		return nil
	}
	p := &binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{}
	p.SetPublicKeyPem(dcl.ValueOrEmptyString(o.PublicKeyPem))
	p.SetSignatureAlgorithm(BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumToProto(o.SignatureAlgorithm))
	return p
}

// AttestorToProto converts a Attestor resource to its proto representation.
func AttestorToProto(resource *binaryauthorization.Attestor) *binaryauthorizationpb.BinaryauthorizationAttestor {
	p := &binaryauthorizationpb.BinaryauthorizationAttestor{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetUserOwnedDrydockNote(BinaryauthorizationAttestorUserOwnedDrydockNoteToProto(resource.UserOwnedDrydockNote))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))

	return p
}

// applyAttestor handles the gRPC request by passing it to the underlying Attestor Apply() method.
func (s *AttestorServer) applyAttestor(ctx context.Context, c *binaryauthorization.Client, request *binaryauthorizationpb.ApplyBinaryauthorizationAttestorRequest) (*binaryauthorizationpb.BinaryauthorizationAttestor, error) {
	p := ProtoToAttestor(request.GetResource())
	res, err := c.ApplyAttestor(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AttestorToProto(res)
	return r, nil
}

// applyBinaryauthorizationAttestor handles the gRPC request by passing it to the underlying Attestor Apply() method.
func (s *AttestorServer) ApplyBinaryauthorizationAttestor(ctx context.Context, request *binaryauthorizationpb.ApplyBinaryauthorizationAttestorRequest) (*binaryauthorizationpb.BinaryauthorizationAttestor, error) {
	cl, err := createConfigAttestor(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyAttestor(ctx, cl, request)
}

// DeleteAttestor handles the gRPC request by passing it to the underlying Attestor Delete() method.
func (s *AttestorServer) DeleteBinaryauthorizationAttestor(ctx context.Context, request *binaryauthorizationpb.DeleteBinaryauthorizationAttestorRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAttestor(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAttestor(ctx, ProtoToAttestor(request.GetResource()))

}

// ListBinaryauthorizationAttestor handles the gRPC request by passing it to the underlying AttestorList() method.
func (s *AttestorServer) ListBinaryauthorizationAttestor(ctx context.Context, request *binaryauthorizationpb.ListBinaryauthorizationAttestorRequest) (*binaryauthorizationpb.ListBinaryauthorizationAttestorResponse, error) {
	cl, err := createConfigAttestor(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAttestor(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*binaryauthorizationpb.BinaryauthorizationAttestor
	for _, r := range resources.Items {
		rp := AttestorToProto(r)
		protos = append(protos, rp)
	}
	p := &binaryauthorizationpb.ListBinaryauthorizationAttestorResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigAttestor(ctx context.Context, service_account_file string) (*binaryauthorization.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return binaryauthorization.NewClient(conf), nil
}
