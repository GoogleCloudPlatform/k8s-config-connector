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

// +mockgcp-support
// apiVersion: kms.cnrm.cloud.google.com/v1beta1
// kind: KMSAutokeyConfig
// service: google.cloud.kms.v1.AutokeyAdmin
// resource: AutokeyConfig

package mockkms

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/kms/v1"
)

type autokeyAdminServer struct {
	*MockService
	pb.UnimplementedAutokeyAdminServer
}

func (r *autokeyAdminServer) GetAutokeyConfig(ctx context.Context, req *pb.GetAutokeyConfigRequest) (*pb.AutokeyConfig, error) {
	name, err := r.parseAutokeyConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AutokeyConfig{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "AutokeyConfig %s not found.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (r *autokeyAdminServer) UpdateAutokeyConfig(ctx context.Context, req *pb.UpdateAutokeyConfigRequest) (*pb.AutokeyConfig, error) {
	reqName := req.GetAutokeyConfig().GetName()
	name, err := r.parseAutokeyConfigName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetAutokeyConfig()).(*pb.AutokeyConfig)
	obj.Name = fqn

	if err := r.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *autokeyAdminServer) ShowEffectiveAutokeyConfig(ctx context.Context, req *pb.ShowEffectiveAutokeyConfigRequest) (*pb.ShowEffectiveAutokeyConfigResponse, error) {
	project := req.Parent
	obj := &pb.ShowEffectiveAutokeyConfigResponse{}
	obj.KeyProject = project

	return obj, nil
}

type autokeyConfigName struct {
	folder    string 
}

func (a *autokeyConfigName) String() string {
	return "folders/" + a.folder + "/autokeyConfig";
}

// parseAutokeyConfigName parses a string into an AutoKeyConfig name.
// The expected form is `folders/{FOLDER_NUMBER}/autokeyConfig`. 
func (r *autokeyAdminServer) parseAutokeyConfigName(name string) (*autokeyConfigName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 3 && tokens[0] == "folders" && tokens[2] == "autokeyConfig" {
		name := &autokeyConfigName{
			folder: tokens[1],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

