// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockcertificatemanager

import (
	"context"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/klog/v2"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/certificatemanager/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type CertificateManagerV1 struct {
	*MockService
	pb.UnimplementedCertificateManagerServer
}

func (s *CertificateManagerV1) GetCertificate(ctx context.Context, req *pb.GetCertificateRequest) (*pb.Certificate, error) {
	name, err := s.parseCertificateName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Certificate{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *CertificateManagerV1) CreateCertificate(ctx context.Context, req *pb.CreateCertificateRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/certificates/" + req.CertificateId
	name, err := s.parseCertificateName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.Certificate).(*pb.Certificate)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if selfManaged := obj.GetSelfManaged(); selfManaged != nil {
		obj.PemCertificate = selfManaged.PemCertificate
	}
	// TODO: Handle self-managed certificates
	// if obj.PemCertificate != "" {
	// 	obj.Type = &pb.Certificate_SelfManaged{
	// 		SelfManaged: &pb.Certificate_SelfManagedCertificate{
	// 			PemCertificate: obj.PemCertificate,
	// 			PemPrivateKey:  obj.PemPrivateKey,
	// 		},
	// 	}
	// }

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opPrefix := "projects/" + name.Project.ID + "/locations/" + name.Location
	op := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
	}
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *CertificateManagerV1) UpdateCertificate(ctx context.Context, req *pb.UpdateCertificateRequest) (*longrunning.Operation, error) {
	reqName := req.GetCertificate().GetName()

	name, err := s.parseCertificateName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Certificate{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		klog.Warningf("update_mask was not provided in request, should be required")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetCertificate().GetDescription()
		case "labels":
			obj.Labels = req.GetCertificate().GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *CertificateManagerV1) DeleteCertificate(ctx context.Context, req *pb.DeleteCertificateRequest) (*longrunning.Operation, error) {
	name, err := s.parseCertificateName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	// Cannot have any dependencies
	if err := storage.List(ctx, s.storage, storage.ListOptions{}, func(obj *pb.CertificateMapEntry) error {
		for _, cert := range obj.GetCertificates() {
			if name.WithNumber() == cert {
				return status.Errorf(codes.FailedPrecondition,
					"can't delete certificate that is referenced by a CertificateMapEntry or other resources")
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	deletedObj := &pb.Certificate{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *CertificateManagerV1) GetCertificateMap(ctx context.Context, req *pb.GetCertificateMapRequest) (*pb.CertificateMap, error) {
	name, err := s.parseCertificateMapName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.CertificateMap{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *CertificateManagerV1) CreateCertificateMap(ctx context.Context, req *pb.CreateCertificateMapRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/certificateMaps/" + req.CertificateMapId
	name, err := s.parseCertificateMapName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.CertificateMap).(*pb.CertificateMap)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opPrefix := "projects/" + name.Project.ID + "/locations/" + name.Location
	op := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
	}
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *CertificateManagerV1) UpdateCertificateMap(ctx context.Context, req *pb.UpdateCertificateMapRequest) (*longrunning.Operation, error) {
	reqName := req.GetCertificateMap().GetName()

	name, err := s.parseCertificateMapName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.CertificateMap{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		klog.Warningf("update_mask was not provided in request, should be required")
	}
	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetCertificateMap().GetDescription()
		case "labels":
			obj.Labels = req.GetCertificateMap().GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *CertificateManagerV1) DeleteCertificateMap(ctx context.Context, req *pb.DeleteCertificateMapRequest) (*longrunning.Operation, error) {
	name, err := s.parseCertificateMapName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	// Cannot have any child CertificateMapEntries
	if err := storage.List(ctx, s.storage, storage.ListOptions{}, func(obj *pb.CertificateMapEntry) error {
		if strings.HasPrefix(obj.GetName(), fqn+"/") {
			return status.Errorf(codes.FailedPrecondition,
				"Resource %q has nested resources. If the API supports cascading delete, set 'force' to true to delete it and its nested resouces.",
				name.String())
		}
		return nil
	}); err != nil {
		return nil, err
	}

	oldObj := &pb.CertificateMap{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *CertificateManagerV1) GetDnsAuthorization(ctx context.Context, req *pb.GetDnsAuthorizationRequest) (*pb.DnsAuthorization, error) {
	name, err := s.parseDNSAuthorizationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DnsAuthorization{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *CertificateManagerV1) CreateDnsAuthorization(ctx context.Context, req *pb.CreateDnsAuthorizationRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/dnsAuthorizations/" + req.DnsAuthorizationId
	name, err := s.parseDNSAuthorizationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.DnsAuthorization).(*pb.DnsAuthorization)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *CertificateManagerV1) UpdateDnsAuthorization(ctx context.Context, req *pb.UpdateDnsAuthorizationRequest) (*longrunning.Operation, error) {
	reqName := req.GetDnsAuthorization().GetName()

	name, err := s.parseDNSAuthorizationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.DnsAuthorization{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		klog.Warningf("update_mask was not provided in request, should be required")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetDnsAuthorization().GetDescription()
		case "labels":
			obj.Labels = req.GetDnsAuthorization().GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *CertificateManagerV1) DeleteDnsAuthorization(ctx context.Context, req *pb.DeleteDnsAuthorizationRequest) (*longrunning.Operation, error) {
	name, err := s.parseDNSAuthorizationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.DnsAuthorization{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *CertificateManagerV1) GetCertificateMapEntry(ctx context.Context, req *pb.GetCertificateMapEntryRequest) (*pb.CertificateMapEntry, error) {
	name, err := s.parseCertificateMapEntryName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.CertificateMapEntry{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *CertificateManagerV1) normalizeCertificateMapEntry(obj *pb.CertificateMapEntry) error {
	for i, cert := range obj.Certificates {
		certName, err := s.parseCertificateName(cert)
		if err != nil {
			return err
		}
		// Store in project-number form
		obj.Certificates[i] = certName.WithNumber()
	}
	return nil
}

func (s *CertificateManagerV1) CreateCertificateMapEntry(ctx context.Context, req *pb.CreateCertificateMapEntryRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/certificateMapEntries/" + req.CertificateMapEntryId
	name, err := s.parseCertificateMapEntryName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.CertificateMapEntry).(*pb.CertificateMapEntry)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	obj.State = pb.ServingState_PENDING

	if err := s.normalizeCertificateMapEntry(obj); err != nil {
		return nil, err
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opPrefix := "projects/" + name.Project.ID + "/locations/" + name.Location
	op := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
	}
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *CertificateManagerV1) UpdateCertificateMapEntry(ctx context.Context, req *pb.UpdateCertificateMapEntryRequest) (*longrunning.Operation, error) {
	reqName := req.GetCertificateMapEntry().GetName()

	name, err := s.parseCertificateMapEntryName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.CertificateMapEntry{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		klog.Warningf("update_mask was not provided in request, should be required")
	}
	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetCertificateMapEntry().GetDescription()
		case "labels":
			obj.Labels = req.GetCertificateMapEntry().GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.normalizeCertificateMapEntry(obj); err != nil {
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *CertificateManagerV1) DeleteCertificateMapEntry(ctx context.Context, req *pb.DeleteCertificateMapEntryRequest) (*longrunning.Operation, error) {
	name, err := s.parseCertificateMapEntryName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb.CertificateMapEntry{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}
