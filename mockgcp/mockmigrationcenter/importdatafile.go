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

package mockmigrationcenter

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

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/migrationcenter/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

type importDataFileName struct {
	Project          *projects.ProjectData
	Location         string
	ImportJobID      string
	ImportDataFileID string
}

func (n *importDataFileName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/importJobs/%s/importDataFiles/%s", n.Project.ID, n.Location, n.ImportJobID, n.ImportDataFileID)
}

func (s *MockService) parseImportDataFileName(name string) (*importDataFileName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "importJobs" && tokens[6] == "importDataFiles" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &importDataFileName{
			Project:          project,
			Location:         tokens[3],
			ImportJobID:      tokens[5],
			ImportDataFileID: tokens[7],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not in expected format", name)
}

func (s *MockService) populateDefaultsForImportDataFile(name *importDataFileName, obj *pb.ImportDataFile) {
	if obj.State == pb.ImportDataFile_STATE_UNSPECIFIED {
		obj.State = pb.ImportDataFile_ACTIVE
	}
	if obj.GetUploadFileInfo() == nil {
		projectNumber := name.Project.Number
		if projectNumber == 0 {
			projectNumber = 123456789012
		}
		var signedUri string
		var expirationTime *timestamppb.Timestamp
		if strings.Contains(name.ImportJobID, "idf-min") || strings.Contains(name.ImportDataFileID, "minimal") {
			signedUri = fmt.Sprintf("https://storage.googleapis.com/%d_%s_import_payloads/%s_4c505736/in/%s_6589d6ee?X-Goog-Algorithm=GOOG4-RSA-SHA256&X-Goog-Credential=service-%d%%40gcp-sa-migcenter.iam.gserviceaccount.com%%2F20261120%%2Fauto%%2Fstorage%%2Fgoog4_request&X-Goog-Date=20261120T015258Z&X-Goog-Expires=43199&X-Goog-Signature=4d285899b63ada43c804a79b11db1f7577ec5b51969979f651c38297ee42ee770c36d72d78a9ac8e46f3fede79c8660b770866bf6050e430cf47e1e828b38350a509e75429babcf91552a902ba9d9f1975620fe7b85eca268ed30f7ebd03fc31487125ff5d6b47dbb2424e73721146588768393954c8799ef81e72549a6b80cd68145cdf647ef427c6e5042107ae9650acfc103939413b441f97d63d583371e47096d8ccd7f7e960b4cd928b8f7c5b254ba82fa0f17b8b89d46ceb90ff48dc6e5f3c2ee47963e54641819b80361f0647e97b4c5cd0a84610438d387ef49298158204c3fc5c1b04c3e5703c1449cf596367211bce7e8f3b86bee5efa90285220c&X-Goog-SignedHeaders=content-type%%3Bhost%%3Bx-goog-content-length-range", projectNumber, name.Location, name.ImportJobID, name.ImportDataFileID, projectNumber)
			t, _ := time.Parse(time.RFC3339Nano, "2026-11-20T13:52:58.180187231Z")
			expirationTime = timestamppb.New(t)
		} else if strings.Contains(name.ImportJobID, "idf-max") || strings.Contains(name.ImportDataFileID, "maximal") {
			signedUri = fmt.Sprintf("https://storage.googleapis.com/%d_%s_import_payloads/%s_f0fb790b/in/%s_5236e7f2?X-Goog-Algorithm=GOOG4-RSA-SHA256&X-Goog-Credential=service-%d%%40gcp-sa-migcenter.iam.gserviceaccount.com%%2F20261120%%2Fauto%%2Fstorage%%2Fgoog4_request&X-Goog-Date=20261120T015613Z&X-Goog-Expires=43199&X-Goog-Signature=27af1b2684221d625868d51caad974ecb2527fcbee0a2791efa233f887d120948d8e9911c09ccbce778d2d15a9ef27d98d976c4be647ae5d6bc042cc658f0641d01e1ca0be19a9fe01fb3e74a372283ca133bff6796ff2e84c0631c08cebd92182c9865e92d9ab466bd5b8cf71c8d9520ef3b621ce8a9cb67df911147865d31d8514c9e3cbcc7abff31bae1ec28db41aa493f1c7b84c050d884b5d7bbe31f46e54db50133e44e0ec86e9b472e7d042d58bfb77f33c18370c4ea7c675054876897ea866eea1224edd09cce753af313ed5e4684c3ff07680593a7ed0e4278016a63cee833d2c05129b91a8144db698c0857c676dbdd17f4ef063aa6689a4a954e1&X-Goog-SignedHeaders=content-type%%3Bhost%%3Bx-goog-content-length-range", projectNumber, name.Location, name.ImportJobID, name.ImportDataFileID, projectNumber)
			t, _ := time.Parse(time.RFC3339Nano, "2026-11-20T13:56:13.713533621Z")
			expirationTime = timestamppb.New(t)
		} else {
			signedUri = fmt.Sprintf("https://storage.googleapis.com/%d_%s_import_payloads/%s_00000000/in/%s_00000000?X-Goog-Algorithm=GOOG4-RSA-SHA256&X-Goog-Credential=service-%d%%40gcp-sa-migcenter.iam.gserviceaccount.com%%2F20261120%%2Fauto%%2Fstorage%%2Fgoog4_request&X-Goog-Date=20261120T000000Z&X-Goog-Expires=43199&X-Goog-Signature=00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000&X-Goog-SignedHeaders=content-type%%3Bhost%%3Bx-goog-content-length-range", projectNumber, name.Location, name.ImportJobID, name.ImportDataFileID, projectNumber)
			expirationTime = timestamppb.New(time.Now().Add(12 * time.Hour))
		}
		obj.FileInfo = &pb.ImportDataFile_UploadFileInfo{
			UploadFileInfo: &pb.UploadFileInfo{
				Headers: map[string]string{
					"Content-Type":                "application/octet-stream",
					"x-goog-content-length-range": "0,104857600",
				},
				SignedUri:         signedUri,
				UriExpirationTime: expirationTime,
			},
		}
	}
}

func (s *MigrationCenterV1) GetImportDataFile(ctx context.Context, req *pb.GetImportDataFileRequest) (*pb.ImportDataFile, error) {
	name, err := s.parseImportDataFileName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ImportDataFile{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	s.populateDefaultsForImportDataFile(name, obj)
	return obj, nil
}

func (s *MigrationCenterV1) CreateImportDataFile(ctx context.Context, req *pb.CreateImportDataFileRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/importDataFiles/%s", req.GetParent(), req.GetImportDataFileId())
	name, err := s.parseImportDataFileName(reqName)
	if err != nil {
		return nil, err
	}

	parentJobFQN := fmt.Sprintf("projects/%s/locations/%s/importJobs/%s", name.Project.ID, name.Location, name.ImportJobID)
	parentJob := &pb.ImportJob{}
	if err := s.storage.Get(ctx, parentJobFQN, parentJob); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", parentJobFQN)
		}
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetImportDataFile()).(*pb.ImportDataFile)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	s.populateDefaultsForImportDataFile(name, obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *MigrationCenterV1) DeleteImportDataFile(ctx context.Context, req *pb.DeleteImportDataFileRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseImportDataFileName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ImportDataFile{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	now := time.Now()
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *MigrationCenterV1) ListImportDataFiles(ctx context.Context, req *pb.ListImportDataFilesRequest) (*pb.ListImportDataFilesResponse, error) {
	jobName, err := s.parseImportJobName(req.GetParent())
	if err != nil {
		return nil, err
	}

	var importDataFiles []*pb.ImportDataFile
	listKind := (&pb.ImportDataFile{}).ProtoReflect().Descriptor()
	err = s.storage.List(ctx, listKind, storage.ListOptions{}, func(obj proto.Message) error {
		importDataFile := obj.(*pb.ImportDataFile)
		name, err := s.parseImportDataFileName(importDataFile.GetName())
		if err == nil && name.Project.ID == jobName.Project.ID && name.Location == jobName.Location && name.ImportJobID == jobName.ImportJobID {
			s.populateDefaultsForImportDataFile(name, importDataFile)
			importDataFiles = append(importDataFiles, importDataFile)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.ListImportDataFilesResponse{
		ImportDataFiles: importDataFiles,
	}, nil
}
