// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockspanner

import (
	"context"
	"regexp"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/spanner/admin/database/v1"
)

var _ pb.DatabaseAdminServer = &SpannerDatabaseV1{}

type SpannerDatabaseV1 struct {
	*MockService
	pb.UnimplementedDatabaseAdminServer
}

type spannerDatabaseName struct {
	Project      *projects.ProjectData
	InstanceName string
	DatabaseName string
}

func (s *SpannerDatabaseV1) GetDatabase(ctx context.Context, req *pb.GetDatabaseRequest) (*pb.Database, error) {
	name, err := s.parseDatabaseName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Database{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *SpannerDatabaseV1) CreateDatabase(ctx context.Context, req *pb.CreateDatabaseRequest) (*longrunningpb.Operation, error) {
	databaseName := databaseNameFromCreateStatement(req.GetCreateStatement())
	fqn := req.GetParent() + "/databases/" + databaseName
	if isValid := validateDatabaseName(databaseName); !isValid {
		return nil, status.Errorf(codes.InvalidArgument,
			"Expected projects/{project ID}/instances/{instance ID}/databases/{database name}\n"+
				"Got: %v\nError: Invalid database name: %q. Database name must start with a lowercase letter, "+
				"be 2-30 characters long, contain only lowercase letters, numbers, underscores or hyphens, "+
				"and not end with an underscore or hyphen.", fqn, databaseName)
	}

	obj := &pb.Database{}
	obj.Name = fqn
	obj.DatabaseDialect = req.DatabaseDialect
	if obj.DatabaseDialect == pb.DatabaseDialect_DATABASE_DIALECT_UNSPECIFIED {
		obj.DatabaseDialect = pb.DatabaseDialect_GOOGLE_STANDARD_SQL
	}
	obj.VersionRetentionPeriod = versionRetentionPeriodFromDDL(req.GetExtraStatements())
	obj.EncryptionConfig = req.EncryptionConfig
	obj.State = pb.Database_READY
	obj.CreateTime = timestamppb.Now()

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.CreateDatabaseMetadata{
		Database: fqn,
	}
	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		// Many fields are not populated in the LRO result
		result := proto.Clone(obj).(*pb.Database)
		return result, nil
	})
}

// validateDatabaseName verifies that the databaseName matches the API's
// requirement:
// "Database name must start with a lowercase letter, be 2-30 characters long,
// contain only lowercase letters, numbers, underscores or hyphens, and not end
// with an underscore or hyphen.
func validateDatabaseName(databaseName string) bool {
	validDatabaseName := regexp.MustCompile(`^[a-z][-a-z0-9_]*[a-z0-9]$`)
	return validDatabaseName.MatchString(databaseName) && len(databaseName) <= 30 && len(databaseName) >= 2
}

func (s *SpannerDatabaseV1) DropDatabase(ctx context.Context, req *pb.DropDatabaseRequest) (*emptypb.Empty, error) {
	name, err := s.parseDatabaseName(req.Database)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Database{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *SpannerDatabaseV1) UpdateDatabaseDdl(ctx context.Context, req *pb.UpdateDatabaseDdlRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseDatabaseName(req.Database)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Database{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.UpdateDatabaseDdlMetadata{
		Database:   fqn,
		Statements: req.Statements,
	}
	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		// Many fields are not populated in the LRO result
		result := proto.Clone(obj).(*pb.Database)
		result.Name = fqn
		return result, nil
	})
}

// Given a createStatement string of the format "CREATE DATABASE `<dbname>`",
// return the database name as a string.
func databaseNameFromCreateStatement(createStatement string) string {
	dbname := strings.TrimPrefix(createStatement, "CREATE DATABASE `")
	dbname = strings.TrimSuffix(dbname, "`")

	if strings.ContainsAny(dbname, " \t\n\r") {
		klog.Fatalf("invalid database name: %q", dbname)
	}

	if dbname == "" {
		klog.Fatalf("invalid database name: %q", dbname)
	}

	return dbname
}

func (s *SpannerDatabaseV1) ListBackups(ctx context.Context, req *pb.ListBackupsRequest) (*pb.ListBackupsResponse, error) {
	return &pb.ListBackupsResponse{
		Backups: []*pb.Backup{},
	}, nil
}

// Given a list of DDL extra statement strings, find one with the format
// "ALTER DATABASE `%s` SET OPTIONS (version_retention_period = '%s')"
// and return the version_retention_period, or "1h" if no match.
func versionRetentionPeriodFromDDL(ddl []string) string {
	re := regexp.MustCompile(`ALTER DATABASE \x60.*\x60 SET OPTIONS \(version_retention_period = '(.*)'\)`)
	for _, s := range ddl {
		if matches := re.FindStringSubmatch(s); len(matches) > 0 {
			return matches[1]
		}
	}
	return "1h"
}
