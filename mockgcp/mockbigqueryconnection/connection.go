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

package mockbigqueryconnection

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/bigquery/connection/v1"
)

type ConnectionV1 struct {
	*MockService
	pb.UnimplementedConnectionServiceServer
}

func (s *ConnectionV1) GetConnection(ctx context.Context, req *pb.GetConnectionRequest) (*pb.Connection, error) {
	name, err := s.parseConnectionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Connection{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: Connection %s", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *ConnectionV1) CreateConnection(ctx context.Context, req *pb.CreateConnectionRequest) (*pb.Connection, error) {
	var reqName string
	if req.ConnectionId != "" {
		reqName = req.Parent + "/connections/" + req.ConnectionId
	} else if req.Connection.Name != "" {
		reqName = req.Connection.Name
	} else {
		// reqName = req.Parent + "/connections/" + uuid.New().String()
		// Using fixed UUID to test "acquire" in spec.resourceID. This also fix the dynamic uuid value in the `x-goog-request-params` header.
		reqName = req.Parent + "/connections/" + "71389360-831c-431d-8975-837aee2153be"
	}
	name, err := s.parseConnectionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.Connection).(*pb.Connection)

	obj.Name = name.String()
	obj.CreationTime = now.Unix()
	obj.LastModifiedTime = now.Unix()

	// Bigqueryconnections supports two types of SA.
	// Primary SA is attached with pre-defined IAM roles, while the delegation SA doesn't include any roles.
	// see https://cloud.google.com/iam/docs/service-agents#bigquery-connection-delegation-service-agent.

	buildDelegationServiceAccountId := func() string {
		letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		b := make([]rune, 4)
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}
		return fmt.Sprintf("bqcx-%s-%s@gcp-sa-bigquery-condel.iam.gserviceaccount.com", req.GetParent(), string(b))
	}

	buildPrimaryServiceAccountId := func() string {
		return fmt.Sprintf("service-%s@gcp-sa-bigqueryconnection.iam.gserviceaccount.com", req.GetParent())
	}

	buildGoogleIdentity := func() string {
		letterRunes := []rune("0123456789")
		b := make([]rune, 21)
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}
		return string(b)
	}

	if _, ok := (req.Connection.Properties).(*pb.Connection_Aws); ok {
		if aws := req.Connection.GetAws(); aws != nil {
			obj.Properties = &pb.Connection_Aws{
				Aws: &pb.AwsProperties{
					AuthenticationMethod: &pb.AwsProperties_AccessRole{
						AccessRole: &pb.AwsAccessRole{
							IamRoleId: aws.GetAccessRole().GetIamRoleId(),
							Identity:  buildGoogleIdentity(),
						},
					},
				},
			}
		}
	}

	if _, ok := (req.Connection.Properties).(*pb.Connection_Azure); ok {
		if azure := req.Connection.GetAzure(); azure != nil {
			obj.Properties = &pb.Connection_Azure{
				Azure: &pb.AzureProperties{
					CustomerTenantId:             azure.GetCustomerTenantId(),
					FederatedApplicationClientId: azure.GetFederatedApplicationClientId(),
					Identity:                     buildGoogleIdentity(),
				},
			}
		}
	}

	if _, ok := (req.Connection.Properties).(*pb.Connection_CloudResource); ok {
		obj.Properties = &pb.Connection_CloudResource{
			CloudResource: &pb.CloudResourceProperties{
				ServiceAccountId: buildDelegationServiceAccountId(),
			},
		}
	}

	if _, ok := (req.Connection.Properties).(*pb.Connection_CloudSql); ok {
		obj.HasCredential = true
		if sql := req.Connection.GetCloudSql(); sql != nil {
			obj.Properties = &pb.Connection_CloudSql{
				CloudSql: &pb.CloudSqlProperties{
					InstanceId:       sql.InstanceId,
					Database:         sql.Database,
					Type:             sql.Type,
					ServiceAccountId: buildPrimaryServiceAccountId(),
				},
			}
		}
	}

	if _, ok := (req.Connection.Properties).(*pb.Connection_CloudSpanner); ok {
		if spanner := req.Connection.GetCloudSpanner(); spanner != nil {
			obj.Properties = &pb.Connection_CloudSpanner{
				CloudSpanner: &pb.CloudSpannerProperties{
					Database:       spanner.Database,
					UseParallelism: spanner.UseParallelism,
					UseDataBoost:   spanner.UseDataBoost,
					MaxParallelism: spanner.MaxParallelism,
					DatabaseRole:   spanner.DatabaseRole,
				},
			}
		}
	}

	if _, ok := (req.Connection.Properties).(*pb.Connection_Spark); ok {
		if spark := req.Connection.GetSpark(); spark != nil {
			obj.Properties = &pb.Connection_Spark{
				Spark: &pb.SparkProperties{
					MetastoreServiceConfig:   spark.GetMetastoreServiceConfig(),
					SparkHistoryServerConfig: spark.GetSparkHistoryServerConfig(),
					ServiceAccountId:         buildDelegationServiceAccountId(),
				},
			}
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *ConnectionV1) UpdateConnection(ctx context.Context, req *pb.UpdateConnectionRequest) (*pb.Connection, error) {
	name, err := s.parseConnectionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.Connection{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	for _, path := range paths {
		switch path {
		case "friendlyName":
			obj.FriendlyName = req.GetConnection().GetFriendlyName()
		case "description":
			obj.Description = req.GetConnection().GetDescription()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}
	obj.LastModifiedTime = now.Unix()

	if _, ok := (req.Connection.Properties).(*pb.Connection_Aws); ok {
		if mod := req.Connection.GetAws(); mod != nil {
			obj.GetAws().GetAccessRole().IamRoleId = mod.GetAccessRole().IamRoleId
		}
	}

	if _, ok := (req.Connection.Properties).(*pb.Connection_CloudSpanner); ok {
		if mod := req.Connection.GetCloudSpanner(); mod != nil {
			obj.GetCloudSpanner().Database = mod.Database
			obj.GetCloudSpanner().UseDataBoost = mod.UseDataBoost
			obj.GetCloudSpanner().UseParallelism = mod.UseParallelism
			obj.GetCloudSpanner().MaxParallelism = mod.MaxParallelism
			obj.GetCloudSpanner().DatabaseRole = mod.DatabaseRole
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *ConnectionV1) DeleteConnection(ctx context.Context, req *pb.DeleteConnectionRequest) (*empty.Empty, error) {
	name, err := s.parseConnectionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Connection{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

type connectionName struct {
	Project    *projects.ProjectData
	Location   string
	ResourceID string
}

func (n *connectionName) String() string {
	return "projects/" + strconv.FormatInt(n.Project.Number, 10) + "/locations/" + n.Location + "/connections/" + n.ResourceID
}

// parseConnectionName parses a string into a connectionName.
// The expected form is projects/<projectNum>/locations/<location>/connections/<connectionID>
func (s *MockService) parseConnectionName(name string) (*connectionName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "connections" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &connectionName{
			Project:    project,
			Location:   tokens[3],
			ResourceID: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
