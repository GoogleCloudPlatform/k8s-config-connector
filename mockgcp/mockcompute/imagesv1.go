// Copyright 2022 Google LLC
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

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ImagesV1 struct {
	*MockService
	pb.UnimplementedImagesServer
}

func (s *ImagesV1) GetFromFamily(ctx context.Context, req *pb.GetFromFamilyImageRequest) (*pb.Image, error) {
	obj := &pb.Image{}

	// Details from gcloud compute images describe-from-family debian-11 --project debian-cloud --log-http
	obj.Kind = PtrTo("compute#image")
	obj.Name = PtrTo("debian-11-bullseye-v20231010")
	obj.Description = PtrTo("Debian, Debian GNU/Linux, 11 (bullseye), amd64 built on 20231010")
	obj.SelfLink = PtrTo("https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-11-bullseye-v20231010")
	obj.Family = PtrTo("debian-11")
	obj.Status = PtrTo("UP")

	return obj, nil
}

func (s *ImagesV1) Get(ctx context.Context, req *pb.GetImageRequest) (*pb.Image, error) {
	// {
	// 	"error": {
	// 	  "code": 404,
	// 	  "message": "The resource 'projects/debian-cloud/global/images/debian-11' was not found",
	// 	  "errors": [
	// 		{
	// 		  "message": "The resource 'projects/debian-cloud/global/images/debian-11' was not found",
	// 		  "domain": "global",
	// 		  "reason": "notFound"
	// 		}
	// 	  ]
	// 	}
	//   }

	return nil, status.Errorf(codes.NotFound, "image not found")
}
