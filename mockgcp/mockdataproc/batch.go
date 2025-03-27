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
// proto.service: google.cloud.dataproc.v1
// proto.message: google.cloud.dataproc.v1.Batch

package mockdataproc

import (
	"context"
	"fmt"
	"strings"
	"time"

	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type batchControllerServer struct {
	*MockService
	pb.UnimplementedBatchControllerServer
}

func (s *batchControllerServer) GetBatch(ctx context.Context, req *pb.GetBatchRequest) (*pb.Batch, error) {
	name, err := s.parseBatchName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Batch{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *batchControllerServer) CreateBatch(ctx context.Context, req *pb.CreateBatchRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseBatchName(req.Parent + "/batches/" + req.BatchId)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()
	labels := make(map[string]string)
	labels["cnrm-test"] = "true"
	labels["managed-by-cnrm"] = "true"
	labels["goog-dataproc-batch-id"] = req.BatchId
	labels["goog-dataproc-location"] = name.Location
	labels["goog-dataproc-batch-uuid"] = req.RequestId
	labels["goog-dataproc-drz-resource-uuid"] = "batch-" + req.RequestId
	properties := make(map[string]string)
	properties["spark:spark.app.name"] = fqn
	properties["spark:spark.dataproc.scaling.version"] = "2"
	properties["spark:spark.driver.cores"] = "4"
	properties["spark:spark.driver.memory"] = "9600m"
	properties["spark:spark.dynamicAllocation.executorAllocationRatio"] = "0.3"
	properties["spark:spark.executor.cores"] = "4"
	properties["spark:spark.executor.instances"] = "2"
	properties["spark:spark.executor.memory"] = "9600m"

	obj := &pb.Batch{}
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.EnvironmentConfig = &pb.EnvironmentConfig{
		ExecutionConfig: &pb.ExecutionConfig{
			AuthenticationConfig: &pb.AuthenticationConfig{},
			Ttl:                  durationpb.New(14400 * time.Second),
		},
		PeripheralsConfig: &pb.PeripheralsConfig{
			SparkHistoryServerConfig: &pb.SparkHistoryServerConfig{},
		},
	}
	obj.BatchConfig = &pb.Batch_PysparkBatch{
		PysparkBatch: &pb.PySparkBatch{
			MainPythonFileUri: "gs://config-connector-samples/dataproc/spark.py",
		},
	}
	obj.Labels = labels
	obj.RuntimeConfig = &pb.RuntimeConfig{
		Properties: properties,
		Version:    "2.2.39",
	}
	obj.RuntimeInfo = &pb.RuntimeInfo{
		ApproximateUsage: &pb.UsageMetrics{
			AcceleratorType:         "UNSPECIFIED",
			MilliDcuSeconds:         711000,
			ShuffleStorageGbSeconds: 72000,
		},
		OutputUri: "gs://dataproc-staging-us-central1-${projectNumber}-hyktfe58/google-cloud-dataproc-metainfo/fffc30c2-/jobs/srvls-batch-27e5b2b4/driveroutput",
	}
	obj.State = pb.Batch_SUCCEEDED
	obj.StateTime = timestamppb.New(now)
	obj.Uuid = req.RequestId
	obj.StateHistory = []*pb.Batch_StateHistory{
		{
			State:          pb.Batch_PENDING,
			StateStartTime: timestamppb.New(now),
		},
	}
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/regions/%s", name.Project, name.Location)
	lroMetadata := &pb.BatchOperationMetadata{
		Batch:         fqn,
		OperationType: pb.BatchOperationMetadata_BATCH,
		Description:   "Batch",
		BatchUuid:     req.RequestId,
		CreateTime:    timestamppb.New(now),
		DoneTime:      timestamppb.New(now),
		Labels:        labels,
		Warnings: []string{
			"No runtime version specified. Using the default runtime version. It is recommended to select a specific major.minor runtime version in production, as the default runtime version may change at any time.",
		},
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		obj.StateHistory = append(obj.StateHistory, &pb.Batch_StateHistory{
			State:          pb.Batch_RUNNING,
			StateStartTime: timestamppb.New(now),
		})
		return obj, nil
	})

}

func (s *batchControllerServer) DeleteBatch(ctx context.Context, req *pb.DeleteBatchRequest) (*emptypb.Empty, error) {
	name, err := s.parseBatchName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Batch{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type batchName struct {
	Project  string
	Batch    string
	Location string
}

func (n *batchName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/batches/%s", n.Project, n.Location, n.Batch)
}

// parseBatchName parses a string into a batchName.
// The expected form is `projects/*/locations/*/batches/*`.
func (s *MockService) parseBatchName(name string) (*batchName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "batches" {
		name := &batchName{
			Project:  tokens[1],
			Location: tokens[3],
			Batch:    tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
