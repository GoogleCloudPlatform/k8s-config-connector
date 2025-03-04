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
// proto.service: google.cloud.apphub.v1
// proto.message: google.cloud.apphub.v1.DiscoveredService

package mockapphub

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/apphub/v1"
)

type appHubServer struct {
	*MockService
	pb.UnimplementedAppHubServer
}

func (s *appHubServer) ListDiscoveredServices(ctx context.Context, req *pb.ListDiscoveredServicesRequest) (*pb.ListDiscoveredServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDiscoveredServices not implemented")
}

func (s *appHubServer) GetDiscoveredService(ctx context.Context, req *pb.GetDiscoveredServiceRequest) (*pb.DiscoveredService, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDiscoveredService not implemented")
}



