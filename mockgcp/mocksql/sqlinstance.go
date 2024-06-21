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

package mocksql

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/sql/v1beta4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type sqlInstancesService struct {
	*MockService
	pb.UnimplementedSqlInstancesServiceServer
}

func (s *sqlInstancesService) Get(ctx context.Context, req *pb.SqlInstancesGetRequest) (*pb.DatabaseInstance, error) {
	name, err := s.buildInstanceName(req.GetProject(), req.GetInstance())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DatabaseInstance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *sqlInstancesService) Insert(ctx context.Context, req *pb.SqlInstancesInsertRequest) (*pb.Operation, error) {
	name, err := s.buildInstanceName(req.GetProject(), req.GetBody().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	region := "us-central1"
	zone := "us-central1-c"

	obj := proto.Clone(req.GetBody()).(*pb.DatabaseInstance)
	obj.Name = name.InstanceName
	obj.Project = name.Project.ID
	obj.Region = region

	obj.SelfLink = fmt.Sprintf("https://sqladmin.googleapis.com/sql/v1beta4/projects/%s/instances/%s",
		name.Project.ID, name.InstanceName)
	obj.ConnectionName = obj.Project + ":" + obj.Region + ":" + obj.Name
	obj.CreateTime = timestamppb.New(now)
	switch obj.DatabaseVersion {
	case pb.SqlDatabaseVersion_MYSQL_5_7:
		obj.DatabaseInstalledVersion = "MYSQL_5_7_44"
		obj.MaintenanceVersion = "MYSQL_5_7_44.R20231105.01_03"
	case pb.SqlDatabaseVersion_SQLSERVER_2017_EXPRESS:
		obj.DatabaseInstalledVersion = "SQLSERVER_2017_EXPRESS_CU31_GDR"
		obj.MaintenanceVersion = "SQLSERVER_2017_EXPRESS_CU31_GDR.R20231029.00_02"
	case pb.SqlDatabaseVersion_POSTGRES_9_6:
		obj.DatabaseInstalledVersion = "POSTGRES_9_6"
	default:
		return nil, fmt.Errorf("database version %s not yet supported by mock", obj.DatabaseVersion)
	}
	obj.GceZone = zone
	obj.IpAddresses = []*pb.IpMapping{
		{
			IpAddress: "10.10.10.10",
			Type:      pb.SqlIpAddressType_PRIMARY,
		},
	}
	obj.Kind = "sql#instance"

	obj.BackendType = pb.SqlBackendType_SECOND_GEN

	if obj.SqlNetworkArchitecture == nil {
		obj.SqlNetworkArchitecture = pb.DatabaseInstance_NEW_NETWORK_ARCHITECTURE.Enum()
	}
	obj.State = pb.DatabaseInstance_RUNNABLE

	obj.ServerCaCert = &pb.SslCert{
		CertSerialNumber: "0",
		Cert:             "-----BEGIN CERTIFICATE-----\n-----END CERTIFICATE-----\n",
		CommonName:       "common-name",
		CreateTime:       timestamppb.New(now),
		ExpirationTime:   timestamppb.New(now.Add(time.Hour * 24 * 365)),
		Sha1Fingerprint:  "12345678",
		Instance:         name.InstanceName,
		Kind:             "sql#sslCert",
	}

	obj.ServiceAccountEmailAddress = fmt.Sprintf("p%d-abcdef@gcp-sa-cloud-sql.iam.gserviceaccount.com", name.Project.Number)

	populateDefaults(obj, zone)

	obj.Settings.SettingsVersion = wrapperspb.Int64(1)

	obj.Etag = fields.ComputeWeakEtag(obj)

	if err := validateDatabaseInstance(obj); err != nil {
		return nil, err
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Move to workflow
	{
		if isMysql(obj) {
			if _, err := s.users.Insert(ctx, &pb.SqlUsersInsertRequest{
				Instance: name.InstanceName,
				Project:  name.Project.ID,
				Body: &pb.User{
					Name: "root",
					Host: "%",
				},
			}); err != nil {
				return nil, fmt.Errorf("creating root user: %w", err)
			}
		} else if isSqlServer(obj) {
			users := []*pb.User{
				{
					Name: "##MS_PolicyEventProcessingLogin##",
					UserDetails: &pb.User_SqlserverUserDetails{
						SqlserverUserDetails: &pb.SqlServerUserDetails{
							Disabled: true,
						},
					},
				},
				{
					Name: "##MS_PolicyTsqlExecutionLogin##",
					UserDetails: &pb.User_SqlserverUserDetails{
						SqlserverUserDetails: &pb.SqlServerUserDetails{
							Disabled: true,
						},
					},
				},
				{
					Name: "sqlserver",
					UserDetails: &pb.User_SqlserverUserDetails{
						SqlserverUserDetails: &pb.SqlServerUserDetails{
							ServerRoles: []string{"CustomerDbRootRole"},
						},
					},
				},
			}

			for _, user := range users {
				if _, err := s.users.Insert(ctx, &pb.SqlUsersInsertRequest{
					Instance: name.InstanceName,
					Project:  name.Project.ID,
					Body:     user,
				}); err != nil {
					return nil, fmt.Errorf("creating initial user: %w", err)
				}
			}
		}
	}

	op := &pb.Operation{
		TargetProject: name.Project.ID,
		OperationType: pb.Operation_CREATE,
	}

	return s.operations.startLRO(ctx, op, obj, func() (proto.Message, error) {
		return obj, nil
	})
}

func setDefaultInt64(pp **wrapperspb.Int64Value, defaultValue int64) {
	if *pp == nil {
		*pp = &wrapperspb.Int64Value{
			Value: defaultValue,
		}
	}
}

func setDefaultInt32(pp **wrapperspb.Int32Value, defaultValue int32) {
	if *pp == nil {
		*pp = &wrapperspb.Int32Value{
			Value: defaultValue,
		}
	}
}

func setDefaultBool(pp **wrapperspb.BoolValue, defaultValue bool) {
	if *pp == nil {
		*pp = &wrapperspb.BoolValue{
			Value: defaultValue,
		}
	}
}
func populateDefaults(obj *pb.DatabaseInstance, zone string) {
	settings := obj.Settings
	settings.Kind = "sql#settings"
	if settings.AuthorizedGaeApplications == nil {
		settings.AuthorizedGaeApplications = make([]string, 0)
	}
	setDefaultInt64(&settings.DataDiskSizeGb, 10)
	setDefaultBool(&settings.DeletionProtectionEnabled, false)
	if settings.ConnectorEnforcement == 0 {
		settings.ConnectorEnforcement = pb.Settings_NOT_REQUIRED
	}
	if settings.DataDiskType == 0 {
		settings.DataDiskType = pb.SqlDataDiskType_PD_SSD
	}
	if settings.PricingPlan == 0 {
		settings.PricingPlan = pb.SqlPricingPlan_PER_USE
	}
	if settings.ReplicationType == 0 {
		settings.ReplicationType = pb.SqlReplicationType_SYNCHRONOUS
	}
	setDefaultInt64(&settings.StorageAutoResizeLimit, 0)
	setDefaultBool(&settings.StorageAutoResize, false)

	if settings.IpConfiguration == nil {
		settings.IpConfiguration = &pb.IpConfiguration{}
	}
	ipConfiguration := settings.IpConfiguration
	if ipConfiguration.AuthorizedNetworks == nil {
		ipConfiguration.AuthorizedNetworks = make([]*pb.AclEntry, 0)
	}
	setDefaultBool(&ipConfiguration.Ipv4Enabled, true)
	setDefaultBool(&ipConfiguration.RequireSsl, false)
	if ipConfiguration.SslMode == 0 {
		ipConfiguration.SslMode = pb.IpConfiguration_ALLOW_UNENCRYPTED_AND_ENCRYPTED
	}

	locationPreference := settings.LocationPreference
	if locationPreference == nil {
		locationPreference = &pb.LocationPreference{}
		settings.LocationPreference = locationPreference
	}
	locationPreference.Kind = "sql#locationPreference"
	locationPreference.Zone = zone

	backupConfiguration := settings.BackupConfiguration
	if backupConfiguration == nil {
		backupConfiguration = &pb.BackupConfiguration{}
		settings.BackupConfiguration = backupConfiguration
	}
	backupConfiguration.Kind = "sql#backupConfiguration"
	setDefaultBool(&backupConfiguration.Enabled, false)
	setDefaultBool(&backupConfiguration.PointInTimeRecoveryEnabled, false)
	setDefaultInt32(&backupConfiguration.TransactionLogRetentionDays, 7)
	if backupConfiguration.StartTime == "" {
		backupConfiguration.StartTime = "21:00"
	}
	backupRetentionSettings := backupConfiguration.BackupRetentionSettings
	if backupRetentionSettings == nil {
		backupRetentionSettings = &pb.BackupRetentionSettings{}
		backupConfiguration.BackupRetentionSettings = backupRetentionSettings
	}
	setDefaultInt32(&backupRetentionSettings.RetainedBackups, 7)
	if backupRetentionSettings.RetentionUnit == 0 {
		backupRetentionSettings.RetentionUnit = pb.BackupRetentionSettings_COUNT
	}

}

func isMysql(obj *pb.DatabaseInstance) bool {
	return strings.HasPrefix(obj.GetDatabaseVersion().String(), "MYSQL_")
}

func isSqlServer(obj *pb.DatabaseInstance) bool {
	return strings.HasPrefix(obj.GetDatabaseVersion().String(), "SQLSERVER_")
}

func validateDatabaseInstance(obj *pb.DatabaseInstance) error {
	// Validate some things we hit when running tests
	if isMysql(obj) {
		if obj.GetSettings().GetAvailabilityType() == pb.SqlAvailabilityType_REGIONAL {
			if !obj.GetSettings().GetBackupConfiguration().GetBinaryLogEnabled().GetValue() {
				return status.Errorf(codes.InvalidArgument, "Invalid flag for instance role: MySQL HA non-replica instances need to have binary logging enabled.")
			}
		}

		if !obj.GetSettings().GetBackupConfiguration().GetEnabled().GetValue() {
			if obj.GetSettings().GetBackupConfiguration().GetBinaryLogEnabled().GetValue() {
				return status.Errorf(codes.InvalidArgument, "Binary log must be disabled when backup is disabled or the instance must be a replica instance with a MySQL 5.7 or above version.")
			}
		}
	}
	return nil
}

func (s *sqlInstancesService) Patch(ctx context.Context, req *pb.SqlInstancesPatchRequest) (*pb.Operation, error) {
	name, err := s.buildInstanceName(req.GetProject(), req.GetInstance())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DatabaseInstance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if settings := req.GetBody().GetSettings(); settings != nil {
		if settings.Tier != "" {
			obj.Settings.Tier = settings.Tier
		}
	}
	if body := req.GetBody(); body != nil {
		if body.DatabaseVersion != 0 {
			obj.DatabaseVersion = body.DatabaseVersion
		}
	}

	obj.Settings.SettingsVersion = wrapperspb.Int64(obj.GetSettings().GetSettingsVersion().GetValue() + 1)

	obj.Etag = fields.ComputeWeakEtag(obj)

	if err := validateDatabaseInstance(obj); err != nil {
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetProject: name.Project.ID,
		OperationType: pb.Operation_UPDATE,
	}

	return s.operations.startLRO(ctx, op, obj, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *sqlInstancesService) Update(ctx context.Context, req *pb.SqlInstancesUpdateRequest) (*pb.Operation, error) {
	name, err := s.buildInstanceName(req.GetProject(), req.GetInstance())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.DatabaseInstance{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	obj := proto.Clone(req.GetBody()).(*pb.DatabaseInstance)
	obj.Name = existing.Name
	obj.Region = existing.Region
	obj.Project = existing.Project
	obj.SelfLink = existing.SelfLink
	obj.Kind = existing.Kind

	obj.BackendType = existing.BackendType
	obj.ConnectionName = existing.ConnectionName
	obj.CreateTime = existing.CreateTime
	obj.DatabaseInstalledVersion = existing.DatabaseInstalledVersion
	obj.DatabaseVersion = existing.DatabaseVersion
	obj.GceZone = existing.GceZone
	obj.IpAddresses = existing.IpAddresses
	obj.ServerCaCert = existing.ServerCaCert
	obj.ServiceAccountEmailAddress = existing.ServiceAccountEmailAddress
	obj.MaintenanceVersion = existing.MaintenanceVersion
	obj.ServiceAccountEmailAddress = existing.ServiceAccountEmailAddress
	obj.SqlNetworkArchitecture = existing.SqlNetworkArchitecture
	obj.State = existing.State

	populateDefaults(obj, existing.GetSettings().GetLocationPreference().GetZone())

	obj.Settings.SettingsVersion = wrapperspb.Int64(existing.GetSettings().GetSettingsVersion().GetValue() + 1)

	obj.Etag = fields.ComputeWeakEtag(obj)

	if err := validateDatabaseInstance(obj); err != nil {
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetProject: name.Project.ID,
		OperationType: pb.Operation_UPDATE,
	}

	return s.operations.startLRO(ctx, op, obj, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *sqlInstancesService) Delete(ctx context.Context, req *pb.SqlInstancesDeleteRequest) (*pb.Operation, error) {
	name, err := s.buildInstanceName(req.GetProject(), req.GetInstance())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.DatabaseInstance{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetProject: name.Project.ID,
		OperationType: pb.Operation_DELETE,
	}

	return s.operations.startLRO(ctx, op, deleted, func() (proto.Message, error) {
		return deleted, nil
	})
}

type InstanceName struct {
	Project      *projects.ProjectData
	InstanceName string
}

func (n *InstanceName) String() string {
	return "projects/" + n.Project.ID + "/SQLInstances/" + n.InstanceName
}

// parseSQLInstanceName parses a string into a InstanceName.
// The expected form is projects/<projectID>/instances/<SQLInstanceName>
func (s *MockService) parseInstanceName(name string) (*InstanceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "instances" {
		return s.buildInstanceName(tokens[1], tokens[3])
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *MockService) buildInstanceName(projectID, instanceName string) (*InstanceName, error) {
	project, err := s.projects.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}

	return &InstanceName{
		Project:      project,
		InstanceName: instanceName,
	}, nil
}
