// Copyright 2024 Google LLC
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

// +tool:mockgcp-support
// proto.service: google.cloud.kms.v1.KeyManagementService
// proto.message: google.cloud.kms.v1.ImportJob

package mockkms

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/kms/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *kmsServer) GetImportJob(ctx context.Context, req *pb.GetImportJobRequest) (*pb.ImportJob, error) {
	name, err := s.parseImportJobName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ImportJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "ImportJob %s not found.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *kmsServer) ListImportJobs(ctx context.Context, req *pb.ListImportJobsRequest) (*pb.ListImportJobsResponse, error) {
	var importJobs []*pb.ImportJob

	importJobKind := (&pb.ImportJob{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, importJobKind, storage.ListOptions{}, func(obj proto.Message) error {
		importJob := obj.(*pb.ImportJob)
		if strings.HasPrefix(importJob.GetName(), req.Parent) {
			importJobs = append(importJobs, importJob)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListImportJobsResponse{
		ImportJobs: importJobs,
		TotalSize:  int32(len(importJobs)),
	}, nil
}

func (s *kmsServer) CreateImportJob(ctx context.Context, req *pb.CreateImportJobRequest) (*pb.ImportJob, error) {
	reqName := fmt.Sprintf("%s/importJobs/%s", req.GetParent(), req.GetImportJobId())
	name, err := s.parseImportJobName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetImportJob()).(*pb.ImportJob)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.ExpireTime = timestamppb.New(now)
	obj.ImportMethod = pb.ImportJob_RSA_OAEP_3072_SHA1_AES_256
	obj.State = pb.ImportJob_PENDING_GENERATION

	result := proto.Clone(obj).(*pb.ImportJob)

	obj.GenerateTime = timestamppb.New(now)
	obj.State = pb.ImportJob_ACTIVE
	obj.Attestation = &pb.KeyOperationAttestation{
		CertChains: &pb.KeyOperationAttestation_CertificateChains{
			CaviumCerts: []string{
				"-----BEGIN CERTIFICATE-----\ncertificate 1\n-----END CERTIFICATE-----\n",
				"-----BEGIN CERTIFICATE-----\ncertificate 2\n-----END CERTIFICATE-----\n",
			},
			GoogleCardCerts: []string{
				"-----BEGIN CERTIFICATE-----\ncertificate 3\n-----END CERTIFICATE-----\n",
			},
			GooglePartitionCerts: []string{
				"-----BEGIN CERTIFICATE-----\ncertificate 4\n-----END CERTIFICATE-----\n",
			},
		},
		Content: []byte("content"),
		Format:  pb.KeyOperationAttestation_CAVIUM_V2_COMPRESSED,
	}
	obj.PublicKey = &pb.ImportJob_WrappingPublicKey{
		Pem: "-----BEGIN PUBLIC KEY-----\npublic key\n-----END PUBLIC KEY-----\n",
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return result, nil
}

type importJobName struct {
	KeyRingName
	ImportJobID string
}

func (n *importJobName) String() string {
	return n.KeyRingName.String() + "/importJobs/" + n.ImportJobID
}

// parseImportJobName parses a string into an ImportJobName.
// The expected form is `projects/*/locations/*/keyRings/*/importJobs/*`.
func (s *kmsServer) parseImportJobName(name string) (*importJobName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[6] == "importJobs" {
		keyRingName, err := s.parseKeyRingName(strings.Join(tokens[0:6], "/"))
		if err != nil {
			return nil, err
		}

		name := &importJobName{
			KeyRingName: *keyRingName,
			ImportJobID: tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
