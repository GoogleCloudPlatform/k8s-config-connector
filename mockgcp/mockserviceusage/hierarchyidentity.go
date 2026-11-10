// Copyright 2026 High Five Labs
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

// A folder's or an
// organization's service identity (services.generateServiceIdentity with a
// parent of folders/{id}/services/{service} or
// organizations/{id}/services/{service}). The folder or organization is not
// looked up. The email follows the form of Google's folder and organization
// service agents, service-folder-{id}@gcp-sa-{service}.iam.gserviceaccount.com
// and service-org-{id}@..., where {service} is the first label of the
// service's name; real agents of some services have other domains.

package mockserviceusage

import (
	"context"
	"strings"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/api/serviceusage/v1beta1"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

// hierarchyServiceIdentity answers generateServiceIdentity for a folder or an
// organization; ok is false for any other parent.
func (s *ServiceUsageV1Beta1) hierarchyServiceIdentity(ctx context.Context, parent string) (*longrunning.Operation, bool, error) {
	t := strings.Split(parent, "/")
	if len(t) != 4 || (t[0] != "folders" && t[0] != "organizations") || t[2] != "services" {
		return nil, false, nil
	}
	if t[1] == "" || strings.Trim(t[1], "0123456789") != "" {
		return nil, true, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
	}
	service := t[3]
	short, _, ok := strings.Cut(service, ".")
	if !ok || short == "" {
		return nil, true, status.Errorf(codes.InvalidArgument, "service %q is not valid", service)
	}
	kind := "folder"
	if t[0] == "organizations" {
		kind = "org"
	}
	identity := &pb.ServiceIdentity{
		Email:    "service-" + kind + "-" + t[1] + "@gcp-sa-" + short + ".iam.gserviceaccount.com",
		UniqueId: "1" + t[1],
	}
	op, err := s.operations.NewLRO(ctx)
	if err != nil {
		return nil, true, err
	}
	response, err := anypb.New(identity)
	if err != nil {
		return nil, true, err
	}
	op.Done = true
	op.Result = &longrunning.Operation_Response{Response: response}
	if err := s.storage.Update(ctx, op.Name, op); err != nil {
		return nil, true, err
	}
	return op, true, nil
}
