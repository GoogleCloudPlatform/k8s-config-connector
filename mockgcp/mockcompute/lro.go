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

package mockcompute

import (
	"context"
	"strings"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	lropb "google.golang.org/genproto/googleapis/longrunning"

	"google.golang.org/protobuf/proto"
)

func (s *MockService) newLRO(ctx context.Context, projectID string) (*pb.Operation, error) {
	lro, err := s.operations.NewLRO(ctx)
	if err != nil {
		return nil, err
	}
	selfLinkPrefix := "https://compute.googleapis.com/compute/v1/projects/" + projectID + "/global/operations/"
	return mapToComputeLRO(lro, selfLinkPrefix)
}

func mapToComputeLRO(lro *lropb.Operation, selfLinkPrefix string) (*pb.Operation, error) {
	lroName := strings.TrimPrefix(lro.Name, "operations/")
	computeLRO := &pb.Operation{
		Name: &lroName,
	}
	var status pb.Operation_Status
	if lro.Done {
		status = pb.Operation_DONE
	} else {
		status = pb.Operation_RUNNING
	}
	computeLRO.Status = &status

	computeLRO.SelfLink = proto.String(selfLinkPrefix + lroName)
	return computeLRO, nil
}
