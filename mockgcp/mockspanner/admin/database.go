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

// +tool:mockgcp-support-spanner
// proto.service: google.spanner.admin.database.v1.DatabaseAdmin
// proto.message: google.spanner.admin.database.v1.Database
package mockspanner

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/spanner/admin/database/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
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

func (s *SpannerDatabaseV1) CreateBackupSchedule(ctx context.Context, req *pb.CreateBackupScheduleRequest) (*pb.BackupSchedule, error) {

	reqName := fmt.Sprintf("%s/backupSchedules/%s", req.GetParent(), req.GetBackupScheduleId())
	name, err := s.parseBackupScheduleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()
	obj := req.GetBackupSchedule()
	obj.Name = fqn
	cronSpecText := ""
	if obj.Spec.ScheduleSpec != nil {
		cronSpecText = obj.Spec.GetCronSpec().Text
	}
	obj.Spec.ScheduleSpec = &pb.BackupScheduleSpec_CronSpec{
		CronSpec: &pb.CrontabSpec{
			TimeZone:       "UTC",
			CreationWindow: durationpb.New(14400 * time.Second),
			Text:           cronSpecText,
		},
	}
	obj.UpdateTime = timestamppb.New(now)
	if obj.EncryptionConfig == nil {
		obj.EncryptionConfig = &pb.CreateBackupEncryptionConfig{
			EncryptionType: 1,
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *SpannerDatabaseV1) DeleteBackupSchedule(ctx context.Context, req *pb.DeleteBackupScheduleRequest) (*emptypb.Empty, error) {
	name, err := s.parseBackupScheduleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.BackupSchedule{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *SpannerDatabaseV1) GetBackupSchedule(ctx context.Context, req *pb.GetBackupScheduleRequest) (*pb.BackupSchedule, error) {
	name, err := s.parseBackupScheduleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.BackupSchedule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	obj.Name = fqn
	cronSpecText := ""
	if obj.Spec.ScheduleSpec != nil {
		cronSpecText = obj.Spec.GetCronSpec().Text
	}
	obj.Spec.ScheduleSpec = &pb.BackupScheduleSpec_CronSpec{
		CronSpec: &pb.CrontabSpec{
			TimeZone:       "UTC",
			CreationWindow: durationpb.New(14400 * time.Second),
			Text:           cronSpecText,
		},
	}
	return obj, nil
}

func (s *SpannerDatabaseV1) ListBackupSchedules(ctx context.Context, req *pb.ListBackupSchedulesRequest) (*pb.ListBackupSchedulesResponse, error) {
	name, err := s.parseBackupScheduleName(req.Parent + "/backupSchedules/dummy")
	if err != nil {
		return nil, err
	}
	response := &pb.ListBackupSchedulesResponse{}
	findPrefix := fmt.Sprintf("projects/%s/instances/%s/databases/%s/backupSchedules/%s", name.Project.ID, name.Instance, name.Database, name.BackupName)

	metadataStoreKind := (&pb.BackupSchedule{}).ProtoReflect().Descriptor()

	if err := s.storage.List(ctx, metadataStoreKind, storage.ListOptions{}, func(obj proto.Message) error {
		budget := obj.(*pb.BackupSchedule)
		if strings.HasPrefix(budget.GetName(), findPrefix) {
			response.BackupSchedules = append(response.BackupSchedules, budget)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *SpannerDatabaseV1) UpdateBackupSchedule(ctx context.Context, req *pb.UpdateBackupScheduleRequest) (*pb.BackupSchedule, error) {
	name, err := s.parseBackupScheduleName(req.GetBackupSchedule().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.BackupSchedule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Name = fqn
	obj.UpdateTime = timestamppb.New(now)

	for _, path := range req.GetUpdateMask().GetPaths() {
		switch camelToUnderscore(path) {
		case "spec":
			obj.Spec = req.GetBackupSchedule().GetSpec()
		case "retention_duration":
			obj.RetentionDuration = req.GetBackupSchedule().GetRetentionDuration()
		case "encryption_config":
			obj.EncryptionConfig = req.GetBackupSchedule().GetEncryptionConfig()
		case "backup_type_spec":
			obj.BackupTypeSpec = req.GetBackupSchedule().GetBackupTypeSpec()
		case "spec.cron_spec.text":
			obj.Spec.ScheduleSpec = &pb.BackupScheduleSpec_CronSpec{
				CronSpec: &pb.CrontabSpec{
					Text: req.GetBackupSchedule().GetSpec().GetCronSpec().Text,
				},
			}
		default:
			return nil, status.Errorf(codes.InvalidArgument, "UpdateBackupSchedule does not support field mask path: %q", path)
		}
	}
	cronSpecText := ""
	if obj.Spec.ScheduleSpec != nil {
		cronSpecText = obj.Spec.GetCronSpec().Text
	}
	obj.Spec.ScheduleSpec = &pb.BackupScheduleSpec_CronSpec{
		CronSpec: &pb.CrontabSpec{
			TimeZone:       "UTC",
			CreationWindow: durationpb.New(14400 * time.Second),
			Text:           cronSpecText,
		},
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

type backupScheduleName struct {
	Project    *projects.ProjectData
	Instance   string
	Database   string
	BackupName string
}

func (n *backupScheduleName) String() string {
	return fmt.Sprintf("projects/%s/instances/%s/databases/%s/backupSchedules/%s", n.Project.ID, n.Instance, n.Database, n.BackupName)
}

// parseBackupScheduleName parses a string into a backupScheduleName.
// The expected form is `projects/*/instances/*/databases/*/backupSchedules/*`.
func (s *MockService) parseBackupScheduleName(name string) (*backupScheduleName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "instances" && tokens[4] == "databases" && tokens[6] == "backupSchedules" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &backupScheduleName{
			Project:    project,
			Instance:   tokens[3],
			Database:   tokens[5],
			BackupName: tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func camelToUnderscore(input string) string {
	// Split by dots first to handle paths
	parts := strings.Split(input, ".")
	for i, part := range parts {
		if part == "" {
			continue
		}
		var result strings.Builder
		for j, r := range part {
			if j > 0 && unicode.IsUpper(r) {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(r))
		}
		parts[i] = result.String()
	}
	return strings.Join(parts, ".")
}
