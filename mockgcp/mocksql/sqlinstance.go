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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocks"
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

func (s *sqlInstancesService) Clone(ctx context.Context, req *pb.SqlInstancesCloneRequest) (*pb.Operation, error) {
	sourceFQN, err := s.buildInstanceName(req.GetProject(), req.GetInstance())
	if err != nil {
		return nil, err
	}

	source := &pb.DatabaseInstance{}
	if err := s.storage.Get(ctx, sourceFQN.String(), source); err != nil {
		return nil, err
	}

	cloneName := req.Body.CloneContext.DestinationInstanceName
	clone := proto.Clone(source).(*pb.DatabaseInstance)
	clone.Name = cloneName

	insertReq := &pb.SqlInstancesInsertRequest{
		Project: req.GetProject(),
		Body:    clone,
	}

	insertOp, err := s.Insert(ctx, insertReq)
	if err != nil {
		return nil, err
	}

	cloneOp := &pb.Operation{
		TargetProject: insertOp.TargetProject,
		OperationType: pb.Operation_CLONE,
	}

	return s.operations.startLRO(ctx, cloneOp, clone, func() (proto.Message, error) {
		return clone, nil
	})
}

func (s *sqlInstancesService) Insert(ctx context.Context, req *pb.SqlInstancesInsertRequest) (*pb.Operation, error) {
	name, err := s.buildInstanceName(req.GetProject(), req.GetBody().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetBody()).(*pb.DatabaseInstance)
	obj.Name = name.InstanceName
	obj.Project = name.Project.ID

	obj.SelfLink = fmt.Sprintf("https://sqladmin.googleapis.com/sql/v1beta4/projects/%s/instances/%s",
		name.Project.ID, name.InstanceName)
	obj.ConnectionName = obj.Project + ":" + obj.Region + ":" + obj.Name
	obj.CreateTime = timestamppb.New(now)

	if err := setDatabaseVersionDefaults(obj); err != nil {
		return nil, err
	}

	// By default, allocate a public IP for the instance.
	shouldAllocatePublicIP := true
	// By default, do not allocate a private IP for the instance.
	shouldAllocatePrivateIP := false

	ipConfigurationSpecified := obj.Settings != nil && obj.Settings.IpConfiguration != nil
	ipv4Specified := ipConfigurationSpecified && obj.Settings.IpConfiguration.Ipv4Enabled != nil

	if ipv4Specified && !obj.Settings.IpConfiguration.Ipv4Enabled.Value {
		shouldAllocatePublicIP = false
	}
	if ipConfigurationSpecified && obj.Settings.IpConfiguration.PrivateNetwork != "" {
		shouldAllocatePrivateIP = true
	}

	if shouldAllocatePublicIP {
		obj.IpAddresses = []*pb.IpMapping{
			{
				IpAddress: "10.10.10.10",
				Type:      pb.SqlIpAddressType_PRIMARY,
			},
		}
		if isPostgres(obj) {
			obj.IpAddresses = append(obj.IpAddresses, &pb.IpMapping{
				IpAddress: "10.10.10.11",
				Type:      pb.SqlIpAddressType_OUTGOING,
			})
		}
	}

	if shouldAllocatePrivateIP {
		obj.IpAddresses = append(obj.IpAddresses, &pb.IpMapping{
			IpAddress: "192.168.0.3",
			Type:      pb.SqlIpAddressType_PRIVATE,
		})
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

	populateDefaults(obj)

	obj.GceZone = obj.Settings.LocationPreference.Zone

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
		} else if isPostgres(obj) {
			if _, err := s.users.Insert(ctx, &pb.SqlUsersInsertRequest{
				Instance: name.InstanceName,
				Project:  name.Project.ID,
				Body: &pb.User{
					Name: "postgres",
				},
			}); err != nil {
				return nil, fmt.Errorf("creating postgres user: %w", err)
			}
		}
	}

	op := &pb.Operation{
		TargetProject: name.Project.ID,
		OperationType: pb.Operation_CREATE,
	}
	if obj.InstanceType == pb.SqlInstanceType_READ_REPLICA_INSTANCE {
		op.OperationType = pb.Operation_CREATE_REPLICA
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

func setDatabaseVersionDefaults(obj *pb.DatabaseInstance) error {
	switch obj.DatabaseVersion {
	case pb.SqlDatabaseVersion_MYSQL_5_7:
		obj.DatabaseInstalledVersion = "MYSQL_5_7_44"
		obj.MaintenanceVersion = "MYSQL_5_7_44.R20241020.00_00"
		obj.UpgradableDatabaseVersions = []*pb.AvailableDatabaseVersion{
			{
				DisplayName:  asRef("MySQL 8.0"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.18"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_18"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.26"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_26"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.27"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_27"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.28"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_28"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.29"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_29"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.30"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_30"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.31"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_31"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.32"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_32"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.33"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_33"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.34"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_34"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.35"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_35"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.36"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_36"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.37"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_37"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.39"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_39"),
			},
		}
	case pb.SqlDatabaseVersion_MYSQL_8_0:
		obj.DatabaseInstalledVersion = "MYSQL_8_0_31"
		obj.MaintenanceVersion = "MYSQL_8_0_31.R20240527.01_00"
		obj.UpgradableDatabaseVersions = []*pb.AvailableDatabaseVersion{
			{
				DisplayName:  asRef("MySQL 8.0.18"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_18"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.26"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_26"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.27"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_27"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.28"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_28"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.29"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_29"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.30"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_30"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.32"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_32"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.33"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_33"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.34"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_34"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.35"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_35"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.36"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_36"),
			},
			{
				DisplayName:  asRef("MySQL 8.0.37"),
				MajorVersion: asRef("MYSQL_8_0"),
				Name:         asRef("MYSQL_8_0_37"),
			},
		}
	case pb.SqlDatabaseVersion_SQLSERVER_2017_EXPRESS:
		obj.DatabaseInstalledVersion = "SQLSERVER_2017_EXPRESS_CU31_GDR"
		obj.MaintenanceVersion = "SQLSERVER_2017_EXPRESS_CU31_GDR.R20231029.00_02"
	case pb.SqlDatabaseVersion_SQLSERVER_2019_EXPRESS:
		obj.DatabaseInstalledVersion = "SQLSERVER_2019_EXPRESS_CU26"
		obj.MaintenanceVersion = "SQLSERVER_2019_EXPRESS_CU26.R20240501.00_05"
		obj.UpgradableDatabaseVersions = []*pb.AvailableDatabaseVersion{
			{
				MajorVersion: asRef("SQLSERVER_2019_STANDARD"),
				Name:         asRef("SQLSERVER_2019_STANDARD"),
				DisplayName:  asRef("SQL Server 2019 Standard"),
			},
			{
				MajorVersion: asRef("SQLSERVER_2019_ENTERPRISE"),
				Name:         asRef("SQLSERVER_2019_ENTERPRISE"),
				DisplayName:  asRef("SQL Server 2019 Enterprise"),
			},
			{
				MajorVersion: asRef("SQLSERVER_2019_WEB"),
				Name:         asRef("SQLSERVER_2019_WEB"),
				DisplayName:  asRef("SQL Server 2019 Web"),
			},
			{
				MajorVersion: asRef("SQLSERVER_2022_STANDARD"),
				Name:         asRef("SQLSERVER_2022_STANDARD"),
				DisplayName:  asRef("SQL Server 2022 Standard"),
			},
			{
				MajorVersion: asRef("SQLSERVER_2022_ENTERPRISE"),
				Name:         asRef("SQLSERVER_2022_ENTERPRISE"),
				DisplayName:  asRef("SQL Server 2022 Enterprise"),
			},
			{
				MajorVersion: asRef("SQLSERVER_2022_EXPRESS"),
				Name:         asRef("SQLSERVER_2022_EXPRESS"),
				DisplayName:  asRef("SQL Server 2022 Express"),
			},
			{
				MajorVersion: asRef("SQLSERVER_2022_WEB"),
				Name:         asRef("SQLSERVER_2022_WEB"),
				DisplayName:  asRef("SQL Server 2022 Web"),
			},
		}
	case pb.SqlDatabaseVersion_SQLSERVER_2022_EXPRESS:
		obj.DatabaseInstalledVersion = "SQLSERVER_2022_EXPRESS_CU12_GDR"
		obj.MaintenanceVersion = "SQLSERVER_2022_EXPRESS_CU12_GDR.R20240501.00_05"
		obj.UpgradableDatabaseVersions = []*pb.AvailableDatabaseVersion{
			{
				MajorVersion: asRef("SQLSERVER_2022_STANDARD"),
				Name:         asRef("SQLSERVER_2022_STANDARD"),
				DisplayName:  asRef("SQL Server 2022 Standard"),
			},
			{
				MajorVersion: asRef("SQLSERVER_2022_ENTERPRISE"),
				Name:         asRef("SQLSERVER_2022_ENTERPRISE"),
				DisplayName:  asRef("SQL Server 2022 Enterprise"),
			},
			{
				MajorVersion: asRef("SQLSERVER_2022_WEB"),
				Name:         asRef("SQLSERVER_2022_WEB"),
				DisplayName:  asRef("SQL Server 2022 Web"),
			},
		}
	case pb.SqlDatabaseVersion_POSTGRES_9_6:
		obj.DatabaseInstalledVersion = "POSTGRES_9_6"
	case pb.SqlDatabaseVersion_POSTGRES_15:
		obj.DatabaseInstalledVersion = "POSTGRES_15_7"
		obj.MaintenanceVersion = "POSTGRES_15_7.R20240514.00_12"
		obj.UpgradableDatabaseVersions = []*pb.AvailableDatabaseVersion{
			{
				MajorVersion: asRef("POSTGRES_16"),
				Name:         asRef("POSTGRES_16"),
				DisplayName:  asRef("PostgreSQL 16"),
			},
		}
	case pb.SqlDatabaseVersion_POSTGRES_16:
		obj.DatabaseInstalledVersion = "POSTGRES_16_3"
		obj.MaintenanceVersion = "POSTGRES_16_3.R20240527.01_10"
		obj.UpgradableDatabaseVersions = nil
	default:
		return fmt.Errorf("database version %s not yet supported by mock", obj.DatabaseVersion)
	}
	return nil
}

func populateDefaults(obj *pb.DatabaseInstance) {
	if obj.InstanceType == pb.SqlInstanceType_SQL_INSTANCE_TYPE_UNSPECIFIED {
		obj.InstanceType = pb.SqlInstanceType_CLOUD_SQL_INSTANCE
	}

	if obj.GeminiConfig == nil {
		if isMysql(obj) {
			obj.GeminiConfig = &pb.GeminiInstanceConfig{
				Entitled:               asRef(false),
				FlagRecommenderEnabled: asRef(false),
				IndexAdvisorEnabled:    asRef(false),
				ActiveQueryEnabled:     asRef(false),
			}
		} else if isPostgres(obj) {
			obj.GeminiConfig = &pb.GeminiInstanceConfig{
				Entitled:                asRef(false),
				GoogleVacuumMgmtEnabled: asRef(false),
				OomSessionCancelEnabled: asRef(false),
				ActiveQueryEnabled:      asRef(false),
				IndexAdvisorEnabled:     asRef(false),
			}
		}
	}

	// This field is input only.
	obj.RootPassword = ""

	settings := obj.Settings
	settings.Kind = "sql#settings"
	if settings.AuthorizedGaeApplications == nil {
		settings.AuthorizedGaeApplications = []string{}
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
		ipConfiguration.AuthorizedNetworks = []*pb.AclEntry{}
	}
	setDefaultBool(&ipConfiguration.Ipv4Enabled, true)
	setDefaultBool(&ipConfiguration.RequireSsl, false)
	if ipConfiguration.SslMode == pb.IpConfiguration_SSL_MODE_UNSPECIFIED {
		if ipConfiguration.RequireSsl.Value {
			ipConfiguration.SslMode = pb.IpConfiguration_TRUSTED_CLIENT_CERTIFICATE_REQUIRED
		} else {
			ipConfiguration.SslMode = pb.IpConfiguration_ALLOW_UNENCRYPTED_AND_ENCRYPTED
		}
	}

	if settings.LocationPreference == nil {
		settings.LocationPreference = &pb.LocationPreference{
			Kind: "sql#locationPreference",
			Zone: obj.Region + "-a",
		}
	}

	backupConfiguration := settings.BackupConfiguration
	if backupConfiguration == nil {
		backupConfiguration = &pb.BackupConfiguration{}
		settings.BackupConfiguration = backupConfiguration
	} else {
		if isPostgres(obj) {
			setDefaultBool(&backupConfiguration.ReplicationLogArchivingEnabled, false)
		}

		if backupConfiguration.BinaryLogEnabled != nil && backupConfiguration.BinaryLogEnabled.Value {
			if isPostgres(obj) || isMysql(obj) {
				backupConfiguration.TransactionalLogStorageState = mocks.PtrTo(pb.BackupConfiguration_CLOUD_STORAGE)
			}
		}
	}
	backupConfiguration.Kind = "sql#backupConfiguration"

	backupRetentionSettings := backupConfiguration.BackupRetentionSettings
	if backupRetentionSettings == nil {
		backupRetentionSettings = &pb.BackupRetentionSettings{}
		backupConfiguration.BackupRetentionSettings = backupRetentionSettings
	}
	setDefaultInt32(&backupRetentionSettings.RetainedBackups, 7)
	if backupRetentionSettings.RetentionUnit == 0 {
		backupRetentionSettings.RetentionUnit = pb.BackupRetentionSettings_COUNT
	}

	if backupConfiguration.BinaryLogEnabled != nil && !backupConfiguration.BinaryLogEnabled.Value {
		if !isMysql(obj) {
			backupConfiguration.BinaryLogEnabled = nil
		}
	}

	if backupConfiguration.PointInTimeRecoveryEnabled != nil && isMysql(obj) {
		backupConfiguration.PointInTimeRecoveryEnabled = nil
	}

	setDefaultBool(&backupConfiguration.Enabled, false)
	setDefaultInt32(&backupConfiguration.TransactionLogRetentionDays, 7)
	if backupConfiguration.StartTime == "" {
		backupConfiguration.StartTime = "12:00"
	}
	if backupConfiguration.TransactionalLogStorageState == nil {
		backupConfiguration.TransactionalLogStorageState = asRef(pb.BackupConfiguration_TRANSACTIONAL_LOG_STORAGE_STATE_UNSPECIFIED)
	}

}

func isMysql(obj *pb.DatabaseInstance) bool {
	return strings.HasPrefix(obj.GetDatabaseVersion().String(), "MYSQL_")
}

func isPostgres(obj *pb.DatabaseInstance) bool {
	return strings.HasPrefix(obj.GetDatabaseVersion().String(), "POSTGRES_")
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
		if settings.Edition != pb.Settings_EDITION_UNSPECIFIED {
			obj.Settings.Edition = settings.Edition
		}
		if settings.Tier != "" {
			obj.Settings.Tier = settings.Tier
		}
	}
	if body := req.GetBody(); body != nil {
		if body.DatabaseVersion != pb.SqlDatabaseVersion_SQL_DATABASE_VERSION_UNSPECIFIED {
			obj.DatabaseVersion = body.DatabaseVersion
		}
		if err := setDatabaseVersionDefaults(obj); err != nil {
			return nil, err
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
	obj.UpgradableDatabaseVersions = existing.UpgradableDatabaseVersions

	populateDefaults(obj)

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

func asRef[T any](v T) *T {
	return &v
}
