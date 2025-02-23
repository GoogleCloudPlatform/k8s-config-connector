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

package mockiam

import (
	"context"
	"crypto/md5"
	"regexp"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/iam/admin/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

const ServiceAccountSuffix = ".iam.gserviceaccount.com"

// Gets a [ServiceAccount][google.iam.admin.v1.ServiceAccount].
func (s *IAMServer) GetServiceAccount(ctx context.Context, req *pb.GetServiceAccountRequest) (*pb.ServiceAccount, error) {
	name, err := s.parseServiceAccountName(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	if isNumber(name.Email) {
		uniqueID := name.Email

		// TODO: Some sort of index on uniqueid
		var found *pb.ServiceAccount
		serviceAccountKind := (&pb.ServiceAccount{}).ProtoReflect().Descriptor()
		if err := s.storage.List(ctx, serviceAccountKind, storage.ListOptions{
			Prefix: "projects/" + name.Project.ID + "/",
		}, func(obj proto.Message) error {
			sa := obj.(*pb.ServiceAccount)
			if sa.UniqueId == uniqueID {
				found = sa
			}
			return nil
		}); err != nil {
			return nil, status.Errorf(codes.Internal, "error reading serviceaccounts: %v", err)
		}

		if found == nil {
			return nil, status.Errorf(codes.NotFound, "Unknown service account")
		}

		return found, nil
	}

	sa := &pb.ServiceAccount{}
	fqn := name.String()
	if err := s.storage.Get(ctx, fqn, sa); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Unknown service account")
		}
		return nil, err
	}

	return sa, nil
}

func isNumber(s string) bool {
	match, _ := regexp.MatchString("^[0-9]+$", s)
	return match
}

// Creates a [ServiceAccount][google.iam.admin.v1.ServiceAccount].
func (s *IAMServer) CreateServiceAccount(ctx context.Context, req *pb.CreateServiceAccountRequest) (*pb.ServiceAccount, error) {
	accountID := req.AccountId
	if accountID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "AccountId is required")
	}
	if len(accountID) < 6 || len(accountID) > 30 {
		return nil, status.Errorf(codes.InvalidArgument, "AccountId  must be 6-30 characters long")
	}

	projectName, err := projects.ParseProjectName(req.Name)
	if err != nil {
		return nil, err
	}
	project, err := s.Projects.GetProject(projectName)
	if err != nil {
		return nil, err
	}

	name := &serviceAccountName{
		Project: project,
		Email:   accountID + "@" + project.ID + ServiceAccountSuffix,
	}

	// TODO: Something more real
	id := time.Now().UnixNano()
	id = id & 0xffffffff
	projectNumber := project.Number
	uniqueID := int64(projectNumber)
	uniqueID <<= 32
	uniqueID |= id

	sa := proto.Clone(req.GetServiceAccount()).(*pb.ServiceAccount)
	sa.Name = name.String()
	sa.ProjectId = project.ID
	sa.UniqueId = strconv.FormatInt(uniqueID, 10)
	sa.Email = name.Email
	sa.Oauth2ClientId = sa.UniqueId

	sa.Etag = computeEtag(sa)

	fqn := name.String()
	if err := s.storage.Create(ctx, fqn, sa); err != nil {
		return nil, err
	}

	return sa, nil
}

func (s *IAMServer) DeleteServiceAccount(ctx context.Context, req *pb.DeleteServiceAccountRequest) (*emptypb.Empty, error) {
	name, err := s.parseServiceAccountName(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	deletedObj := &pb.ServiceAccount{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *IAMServer) PatchServiceAccount(ctx context.Context, req *pb.PatchServiceAccountRequest) (*pb.ServiceAccount, error) {
	reqName := req.GetServiceAccount().GetName()

	name, err := s.parseServiceAccountName(ctx, reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	sa := &pb.ServiceAccount{}
	if err := s.storage.Get(ctx, fqn, sa); err != nil {
		return nil, err
	}

	// Unclear exactly what's going on here, but it seems to only return the fields we patch
	retVal := &pb.ServiceAccount{
		Name: sa.Name,
	}

	// You can patch only the `display_name` and `description` fields.
	// You must use the `update_mask` field to specify which of these fields you want to patch.
	paths := req.GetUpdateMask().GetPaths()
	for _, path := range paths {
		switch path {
		case "display_name":
			sa.DisplayName = req.GetServiceAccount().GetDisplayName()
			retVal.DisplayName = sa.DisplayName
		case "description":
			sa.Description = req.GetServiceAccount().GetDescription()
			retVal.Description = sa.Description
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, sa); err != nil {
		return nil, err
	}

	return retVal, nil
}

func (s *IAMServer) DisableServiceAccount(ctx context.Context, req *pb.DisableServiceAccountRequest) (*emptypb.Empty, error) {
	name, err := s.parseServiceAccountName(ctx, req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	sa := &pb.ServiceAccount{}
	if err := s.storage.Get(ctx, fqn, sa); err != nil {
		return nil, err
	}

	sa.Disabled = true

	if err := s.storage.Update(ctx, fqn, sa); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *IAMServer) EnableServiceAccount(ctx context.Context, req *pb.EnableServiceAccountRequest) (*emptypb.Empty, error) {
	name, err := s.parseServiceAccountName(ctx, req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	sa := &pb.ServiceAccount{}
	if err := s.storage.Get(ctx, fqn, sa); err != nil {
		return nil, err
	}

	sa.Disabled = false

	if err := s.storage.Update(ctx, fqn, sa); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func computeEtag(obj proto.Message) []byte {
	// TODO: Do we risk exposing internal fields?  Doesn't matter on a mock, I guess
	b, err := proto.Marshal(obj)
	if err != nil {
		klog.Fatalf("failed to marshal proto object: %v", err)
	}
	hash := md5.Sum(b)
	return hash[:]
}

type serviceAccountName struct {
	Project *projects.ProjectData
	Email   string
}

func (n *serviceAccountName) String() string {
	return "projects/" + n.Project.ID + "/serviceAccounts/" + n.Email
}

func (s *MockService) parseServiceAccountName(ctx context.Context, name string) (*serviceAccountName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "serviceAccounts" {
		projectID := tokens[1]
		email := tokens[3]

		// Using `-` as a wildcard for the `PROJECT_ID` will infer the project from
		// the account. The `ACCOUNT` value can be the `email` address or the
		// `unique_id` of the service account.
		if projectID == "-" {
			tokens := strings.Split(email, "@")
			if len(tokens) == 2 && strings.HasSuffix(tokens[1], ServiceAccountSuffix) {
				projectID = strings.TrimSuffix(tokens[1], ServiceAccountSuffix)
			} else {
				// Infer from the account
				uniqueID, err := strconv.ParseInt(email, 10, 64)
				if err != nil {
					return nil, status.Errorf(codes.InvalidArgument, "name %q not known", name)
				}

				projectNumber := uniqueID >> 32
				project, err := s.Projects.GetProjectByNumber(strconv.FormatInt(projectNumber, 10))
				if err != nil {
					return nil, err
				}

				return &serviceAccountName{
					Project: project,
					Email:   email,
				}, nil
			}
		}

		project, err := s.Projects.GetProjectByID(projectID)
		if err != nil {
			return nil, err
		}

		name := &serviceAccountName{
			Project: project,
			Email:   tokens[3],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
