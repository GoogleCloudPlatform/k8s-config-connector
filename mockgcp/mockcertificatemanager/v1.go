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
	"time"

	pb "google.golang.org/genproto/googleapis/cloud/certificatemanager/v1"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
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

func (s *CertificateManagerV1) DeleteCertificate(ctx context.Context, req *pb.DeleteCertificateRequest) (*longrunning.Operation, error) {
	name, err := s.parseCertificateName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	kind := (&pb.Certificate{}).ProtoReflect().Descriptor()
	if err := s.storage.Delete(ctx, kind, fqn); err != nil {
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

	if time.Now().UnixNano()%2 == 0 {
		return nil, status.Errorf(codes.Internal, "try again!")
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

func (s *CertificateManagerV1) DeleteCertificateMap(ctx context.Context, req *pb.DeleteCertificateMapRequest) (*longrunning.Operation, error) {
	name, err := s.parseCertificateMapName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	kind := (&pb.CertificateMap{}).ProtoReflect().Descriptor()
	if err := s.storage.Delete(ctx, kind, fqn); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "certificate map %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting certificate map: %v", err)
		}
	}

	return s.operations.NewLRO(ctx)
}
