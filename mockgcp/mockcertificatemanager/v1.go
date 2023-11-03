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

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/klog/v2"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/certificatemanager/v1"
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
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "certificate %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading certificate: %v", err)
		}
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

	obj := proto.Clone(req.Certificate).(*pb.Certificate)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating certificate: %v", err)
	}

	return s.operations.NewLRO(ctx)
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
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "certificate %q not found", reqName)
		}
		return nil, status.Errorf(codes.Internal, "error reading certificate: %v", err)
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
		return nil, status.Errorf(codes.Internal, "error updating certificate: %v", err)
	}

	return s.operations.NewLRO(ctx)
}

func (s *CertificateManagerV1) DeleteCertificate(ctx context.Context, req *pb.DeleteCertificateRequest) (*longrunning.Operation, error) {
	name, err := s.parseCertificateName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb.Certificate{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "certificate %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting certificate: %v", err)
		}
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
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "certificateMap %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading certificateMap: %v", err)
		}
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

	obj := proto.Clone(req.CertificateMap).(*pb.CertificateMap)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating certificate map: %v", err)
	}

	return s.operations.NewLRO(ctx)
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
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "certificateMap %q not found", reqName)
		}
		return nil, status.Errorf(codes.Internal, "error reading certificateMap: %v", err)
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
		return nil, status.Errorf(codes.Internal, "error updating certificateMap: %v", err)
	}

	return s.operations.NewLRO(ctx)
}

func (s *CertificateManagerV1) DeleteCertificateMap(ctx context.Context, req *pb.DeleteCertificateMapRequest) (*longrunning.Operation, error) {
	name, err := s.parseCertificateMapName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.CertificateMap{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "certificate map %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting certificate map: %v", err)
		}
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
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "dns authorization %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading dns authorization: %v", err)
		}
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
		return nil, status.Errorf(codes.Internal, "error creating dns authorization: %v", err)
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
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "dnsAuthorization %q not found", reqName)
		}
		return nil, status.Errorf(codes.Internal, "error reading dnsAuthorization: %v", err)
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
		return nil, status.Errorf(codes.Internal, "error updating dnsAuthorization: %v", err)
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
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "dns authorization %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting dns authorization: %v", err)
		}
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
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "certificate map entry %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading certificate map entry: %v", err)
		}
	}

	return obj, nil
}

func (s *CertificateManagerV1) CreateCertificateMapEntry(ctx context.Context, req *pb.CreateCertificateMapEntryRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/certificateMapEntries/" + req.CertificateMapEntryId
	name, err := s.parseCertificateMapEntryName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.CertificateMapEntry).(*pb.CertificateMapEntry)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating certificate map entry: %v", err)
	}

	return s.operations.NewLRO(ctx)
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
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "certificateMapEntry %q not found", reqName)
		}
		return nil, status.Errorf(codes.Internal, "error reading certificateMapEntry: %v", err)
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

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating certificateMapEntry: %v", err)
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
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "certificate map entry %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting certificate map entry: %v", err)
		}
	}

	return s.operations.NewLRO(ctx)
}
