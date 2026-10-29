// Copyright 2026 Google LLC
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
// proto.service: google.cloud.netapp.v1.NetApp
// proto.message: google.cloud.netapp.v1.KmsConfig

package mocknetapp

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	pb "cloud.google.com/go/netapp/apiv1/netapppb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *backupVaultsService) GetKmsConfig(ctx context.Context, req *pb.GetKmsConfigRequest) (*pb.KmsConfig, error) {
	name, err := s.parseKmsConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.KmsConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%v' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *backupVaultsService) ListKmsConfigs(ctx context.Context, req *pb.ListKmsConfigsRequest) (*pb.ListKmsConfigsResponse, error) {
	parent := req.GetParent()
	response := &pb.ListKmsConfigsResponse{}

	findPrefix := fmt.Sprintf("%s/kmsConfigs/", parent)

	kind := (&pb.KmsConfig{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		cloned := proto.Clone(obj).(*pb.KmsConfig)
		response.KmsConfigs = append(response.KmsConfigs, cloned)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *backupVaultsService) CreateKmsConfig(ctx context.Context, req *pb.CreateKmsConfigRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/kmsConfigs/%s", req.GetParent(), req.GetKmsConfigId())
	name, err := s.parseKmsConfigName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetKmsConfig()).(*pb.KmsConfig)
	obj.Name = fqn
	now := time.Now()
	obj.CreateTime = timestamppb.New(now)

	s.populateDefaultsForKmsConfig(name, obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.Now(),
		Target:                name.String(),
		Verb:                  "create",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *backupVaultsService) UpdateKmsConfig(ctx context.Context, req *pb.UpdateKmsConfigRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseKmsConfigName(req.GetKmsConfig().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.KmsConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%v' was not found", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}
	if err := fields.UpdateByFieldMask(obj, req.GetKmsConfig(), req.UpdateMask.Paths); err != nil {
		return nil, fmt.Errorf("update field_mask.paths: %w", err)
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.Now(),
		Target:                name.String(),
		Verb:                  "update",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *backupVaultsService) DeleteKmsConfig(ctx context.Context, req *pb.DeleteKmsConfigRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseKmsConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.KmsConfig{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.Now(),
		Target:                fqn,
		Verb:                  "delete",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *backupVaultsService) VerifyKmsConfig(ctx context.Context, req *pb.VerifyKmsConfigRequest) (*pb.VerifyKmsConfigResponse, error) {
	name, err := s.parseKmsConfigName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.KmsConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%v' was not found", fqn)
		}
		return nil, err
	}

	return &pb.VerifyKmsConfigResponse{
		Healthy:      true,
		Instructions: obj.Instructions,
	}, nil
}

type kmsConfigName struct {
	Project     *projects.ProjectData
	Location    string
	KMSConfigID string
}

func (n *kmsConfigName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/kmsConfigs/%s", n.Project.ID, n.Location, n.KMSConfigID)
}

func (s *MockService) parseKmsConfigName(name string) (*kmsConfigName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "kmsConfigs" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &kmsConfigName{
			Project:     project,
			Location:    tokens[3],
			KMSConfigID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *MockService) populateDefaultsForKmsConfig(name *kmsConfigName, obj *pb.KmsConfig) {
	serviceAccount := fmt.Sprintf("n-cmek-%s-%d@netapp-cmek-prod.iam.gserviceaccount.com", locationCode(name.Location), name.Project.Number)
	obj.ServiceAccount = serviceAccount
	obj.State = pb.KmsConfig_KEY_CHECK_PENDING
	obj.StateDetails = "Credentials created and key check pending"

	cryptoKeyName := obj.GetCryptoKeyName()
	tokens := strings.Split(cryptoKeyName, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "keyRings" && tokens[6] == "cryptoKeys" {
		keyProject := tokens[1]
		keyLocation := tokens[3]
		keyRing := tokens[5]
		keyCryptoKey := tokens[7]
		obj.Instructions = fmt.Sprintf("Please copy and paste the commands listed below into Google Cloud Shell in the project that contains the key ring. The commands create a KMS role and assign it to the CVS service account so that it can access the key.\n## CREATE KMS role ## gcloud iam roles create cmekNetAppVolumesRole --project=%s --title='cmekNetAppVolumesRole' --description='custom cmek cvs role' --permissions=cloudkms.cryptoKeyVersions.get,cloudkms.cryptoKeyVersions.list,cloudkms.cryptoKeyVersions.useToDecrypt,cloudkms.cryptoKeyVersions.useToEncrypt,cloudkms.cryptoKeys.get,cloudkms.keyRings.get,cloudkms.locations.get,cloudkms.locations.list,resourcemanager.projects.get --stage=GA\n ## ASSIGN role and give KEY ACCESS to CVS service account ## gcloud kms keys add-iam-policy-binding %s --project=%s --keyring %s --location %s --member serviceAccount:%s --role projects/%s/roles/cmekNetAppVolumesRole",
			keyProject, keyCryptoKey, keyProject, keyRing, keyLocation, serviceAccount, keyProject)
	}
}

func locationCode(loc string) string {
	parts := strings.Split(loc, "-")
	if len(parts) == 2 {
		region := parts[0]
		sub := parts[1]
		if len(region) >= 2 && len(sub) >= 3 {
			return region[:2] + sub[:2] + string(sub[len(sub)-1])
		}
	}
	return loc
}
