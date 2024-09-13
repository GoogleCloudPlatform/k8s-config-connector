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

package bigqueryconnection

import (
	pb "cloud.google.com/go/bigquery/connection/apiv1/connectionpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryconnection/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AwsAccessRole_FromProto(mapCtx *direct.MapContext, in *pb.AwsAccessRole) *krm.AwsAccessRole {
	if in == nil {
		return nil
	}
	out := &krm.AwsAccessRole{}
	out.IamRoleID = direct.LazyPtr(in.GetIamRoleId())
	out.Identity = direct.LazyPtr(in.GetIdentity())
	return out
}
func AwsAccessRole_ToProto(mapCtx *direct.MapContext, in *krm.AwsAccessRole) *pb.AwsAccessRole {
	if in == nil {
		return nil
	}
	out := &pb.AwsAccessRole{}
	out.IamRoleId = direct.ValueOf(in.IamRoleID)
	out.Identity = direct.ValueOf(in.Identity)
	return out
}
func AwsCrossAccountRole_FromProto(mapCtx *direct.MapContext, in *pb.AwsCrossAccountRole) *krm.AwsCrossAccountRole {
	if in == nil {
		return nil
	}
	out := &krm.AwsCrossAccountRole{}
	out.IamRoleID = direct.LazyPtr(in.GetIamRoleId())
	out.IamUserID = direct.LazyPtr(in.GetIamUserId())
	out.ExternalID = direct.LazyPtr(in.GetExternalId())
	return out
}
func AwsCrossAccountRole_ToProto(mapCtx *direct.MapContext, in *krm.AwsCrossAccountRole) *pb.AwsCrossAccountRole {
	if in == nil {
		return nil
	}
	out := &pb.AwsCrossAccountRole{}
	out.IamRoleId = direct.ValueOf(in.IamRoleID)
	out.IamUserId = direct.ValueOf(in.IamUserID)
	out.ExternalId = direct.ValueOf(in.ExternalID)
	return out
}
func AwsProperties_FromProto(mapCtx *direct.MapContext, in *pb.AwsProperties) *krm.AwsProperties {
	if in == nil {
		return nil
	}
	out := &krm.AwsProperties{}
	out.CrossAccountRole = AwsCrossAccountRole_FromProto(mapCtx, in.GetCrossAccountRole())
	out.AccessRole = AwsAccessRole_FromProto(mapCtx, in.GetAccessRole())
	return out
}
func AwsProperties_ToProto(mapCtx *direct.MapContext, in *krm.AwsProperties) *pb.AwsProperties {
	if in == nil {
		return nil
	}
	out := &pb.AwsProperties{}
	if oneof := AwsCrossAccountRole_ToProto(mapCtx, in.CrossAccountRole); oneof != nil {
		out.AuthenticationMethod = &pb.AwsProperties_CrossAccountRole{CrossAccountRole: oneof}
	}
	if oneof := AwsAccessRole_ToProto(mapCtx, in.AccessRole); oneof != nil {
		out.AuthenticationMethod = &pb.AwsProperties_AccessRole{AccessRole: oneof}
	}
	return out
}
func AzureProperties_FromProto(mapCtx *direct.MapContext, in *pb.AzureProperties) *krm.AzureProperties {
	if in == nil {
		return nil
	}
	out := &krm.AzureProperties{}
	out.Application = direct.LazyPtr(in.GetApplication())
	out.ClientID = direct.LazyPtr(in.GetClientId())
	out.ObjectID = direct.LazyPtr(in.GetObjectId())
	out.CustomerTenantID = direct.LazyPtr(in.GetCustomerTenantId())
	out.RedirectUri = direct.LazyPtr(in.GetRedirectUri())
	out.FederatedApplicationClientID = direct.LazyPtr(in.GetFederatedApplicationClientId())
	out.Identity = direct.LazyPtr(in.GetIdentity())
	return out
}
func AzureProperties_ToProto(mapCtx *direct.MapContext, in *krm.AzureProperties) *pb.AzureProperties {
	if in == nil {
		return nil
	}
	out := &pb.AzureProperties{}
	out.Application = direct.ValueOf(in.Application)
	out.ClientId = direct.ValueOf(in.ClientID)
	out.ObjectId = direct.ValueOf(in.ObjectID)
	out.CustomerTenantId = direct.ValueOf(in.CustomerTenantID)
	out.RedirectUri = direct.ValueOf(in.RedirectUri)
	out.FederatedApplicationClientId = direct.ValueOf(in.FederatedApplicationClientID)
	out.Identity = direct.ValueOf(in.Identity)
	return out
}
func BigQueryConnectionConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Connection) *krm.BigQueryConnectionConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryConnectionConnectionObservedState{}
	// MISSING: Name
	out.FriendlyName = direct.LazyPtr(in.GetFriendlyName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CloudSql
	// MISSING: Aws
	// MISSING: Azure
	// MISSING: CloudSpanner
	out.CloudResource = CloudResourcePropertiesStatus_FromProto(mapCtx, in.GetCloudResource())
	// MISSING: Spark
	// MISSING: SalesforceDataCloud
	out.HasCredential = direct.LazyPtr(in.GetHasCredential())
	return out
}
func BigQueryConnectionConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryConnectionConnectionObservedState) *pb.Connection {
	if in == nil {
		return nil
	}
	out := &pb.Connection{}
	// MISSING: Name
	out.FriendlyName = direct.ValueOf(in.FriendlyName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CloudSql
	// MISSING: Aws
	// MISSING: Azure
	// MISSING: CloudSpanner
	if oneof := CloudResourcePropertiesStatus_ToProto(mapCtx, in.CloudResource); oneof != nil {
		out.Properties = &pb.Connection_CloudResource{CloudResource: oneof}
	}
	// MISSING: Spark
	// MISSING: SalesforceDataCloud
	out.HasCredential = direct.ValueOf(in.HasCredential)
	return out
}
func BigQueryConnectionConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Connection) *krm.BigQueryConnectionConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryConnectionConnectionSpec{}
	// MISSING: Name
	out.FriendlyName = direct.LazyPtr(in.GetFriendlyName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CloudSql
	// MISSING: Aws
	// MISSING: Azure
	// MISSING: CloudSpanner
	// MISSING: Spark
	// MISSING: SalesforceDataCloud
	return out
}
func CloudResourceProperties_FromProto(mapCtx *direct.MapContext, in *pb.CloudResourceProperties) *krm.CloudResourceProperties {
	if in == nil {
		return nil
	}
	out := &krm.CloudResourceProperties{}
	out.ServiceAccountID = direct.LazyPtr(in.GetServiceAccountId())
	return out
}
func CloudResourceProperties_ToProto(mapCtx *direct.MapContext, in *krm.CloudResourceProperties) *pb.CloudResourceProperties {
	if in == nil {
		return nil
	}
	out := &pb.CloudResourceProperties{}
	out.ServiceAccountId = direct.ValueOf(in.ServiceAccountID)
	return out
}
func CloudResourcePropertiesStatus_FromProto(mapCtx *direct.MapContext, in *pb.CloudResourceProperties) *krm.CloudResourcePropertiesStatus {
	if in == nil {
		return nil
	}
	out := &krm.CloudResourcePropertiesStatus{}
	out.ServiceAccountID = direct.LazyPtr(in.GetServiceAccountId())
	return out
}
func CloudResourcePropertiesStatus_ToProto(mapCtx *direct.MapContext, in *krm.CloudResourcePropertiesStatus) *pb.CloudResourceProperties {
	if in == nil {
		return nil
	}
	out := &pb.CloudResourceProperties{}
	out.ServiceAccountId = direct.ValueOf(in.ServiceAccountID)
	return out
}
func CloudSpannerProperties_FromProto(mapCtx *direct.MapContext, in *pb.CloudSpannerProperties) *krm.CloudSpannerProperties {
	if in == nil {
		return nil
	}
	out := &krm.CloudSpannerProperties{}
	out.Database = direct.LazyPtr(in.GetDatabase())
	out.UseParallelism = direct.LazyPtr(in.GetUseParallelism())
	out.MaxParallelism = direct.LazyPtr(in.GetMaxParallelism())
	out.UseServerlessAnalytics = direct.LazyPtr(in.GetUseServerlessAnalytics())
	out.UseDataBoost = direct.LazyPtr(in.GetUseDataBoost())
	out.DatabaseRole = direct.LazyPtr(in.GetDatabaseRole())
	return out
}
func CloudSpannerProperties_ToProto(mapCtx *direct.MapContext, in *krm.CloudSpannerProperties) *pb.CloudSpannerProperties {
	if in == nil {
		return nil
	}
	out := &pb.CloudSpannerProperties{}
	out.Database = direct.ValueOf(in.Database)
	out.UseParallelism = direct.ValueOf(in.UseParallelism)
	out.MaxParallelism = direct.ValueOf(in.MaxParallelism)
	out.UseServerlessAnalytics = direct.ValueOf(in.UseServerlessAnalytics)
	out.UseDataBoost = direct.ValueOf(in.UseDataBoost)
	out.DatabaseRole = direct.ValueOf(in.DatabaseRole)
	return out
}
func CloudSqlCredential_FromProto(mapCtx *direct.MapContext, in *pb.CloudSqlCredential) *krm.CloudSqlCredential {
	if in == nil {
		return nil
	}
	out := &krm.CloudSqlCredential{}
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	return out
}
func CloudSqlCredential_ToProto(mapCtx *direct.MapContext, in *krm.CloudSqlCredential) *pb.CloudSqlCredential {
	if in == nil {
		return nil
	}
	out := &pb.CloudSqlCredential{}
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	return out
}
func CloudSqlProperties_FromProto(mapCtx *direct.MapContext, in *pb.CloudSqlProperties) *krm.CloudSqlProperties {
	if in == nil {
		return nil
	}
	out := &krm.CloudSqlProperties{}
	out.InstanceID = direct.LazyPtr(in.GetInstanceId())
	out.Database = direct.LazyPtr(in.GetDatabase())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Credential = CloudSqlCredential_FromProto(mapCtx, in.GetCredential())
	out.ServiceAccountID = direct.LazyPtr(in.GetServiceAccountId())
	return out
}
func CloudSqlProperties_ToProto(mapCtx *direct.MapContext, in *krm.CloudSqlProperties) *pb.CloudSqlProperties {
	if in == nil {
		return nil
	}
	out := &pb.CloudSqlProperties{}
	out.InstanceId = direct.ValueOf(in.InstanceID)
	out.Database = direct.ValueOf(in.Database)
	out.Type = direct.Enum_ToProto[pb.CloudSqlProperties_DatabaseType](mapCtx, in.Type)
	out.Credential = CloudSqlCredential_ToProto(mapCtx, in.Credential)
	out.ServiceAccountId = direct.ValueOf(in.ServiceAccountID)
	return out
}
func Connection_FromProto(mapCtx *direct.MapContext, in *pb.Connection) *krm.Connection {
	if in == nil {
		return nil
	}
	out := &krm.Connection{}
	out.Name = direct.LazyPtr(in.GetName())
	out.FriendlyName = direct.LazyPtr(in.GetFriendlyName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.CloudSql = CloudSqlProperties_FromProto(mapCtx, in.GetCloudSql())
	out.Aws = AwsProperties_FromProto(mapCtx, in.GetAws())
	out.Azure = AzureProperties_FromProto(mapCtx, in.GetAzure())
	out.CloudSpanner = CloudSpannerProperties_FromProto(mapCtx, in.GetCloudSpanner())
	out.CloudResource = CloudResourceProperties_FromProto(mapCtx, in.GetCloudResource())
	out.Spark = SparkProperties_FromProto(mapCtx, in.GetSpark())
	out.SalesforceDataCloud = SalesforceDataCloudProperties_FromProto(mapCtx, in.GetSalesforceDataCloud())
	out.CreationTime = direct.LazyPtr(in.GetCreationTime())
	out.LastModifiedTime = direct.LazyPtr(in.GetLastModifiedTime())
	out.HasCredential = direct.LazyPtr(in.GetHasCredential())
	return out
}
func Connection_ToProto(mapCtx *direct.MapContext, in *krm.Connection) *pb.Connection {
	if in == nil {
		return nil
	}
	out := &pb.Connection{}
	out.Name = direct.ValueOf(in.Name)
	out.FriendlyName = direct.ValueOf(in.FriendlyName)
	out.Description = direct.ValueOf(in.Description)
	if oneof := CloudSqlProperties_ToProto(mapCtx, in.CloudSql); oneof != nil {
		out.Properties = &pb.Connection_CloudSql{CloudSql: oneof}
	}
	if oneof := AwsProperties_ToProto(mapCtx, in.Aws); oneof != nil {
		out.Properties = &pb.Connection_Aws{Aws: oneof}
	}
	if oneof := AzureProperties_ToProto(mapCtx, in.Azure); oneof != nil {
		out.Properties = &pb.Connection_Azure{Azure: oneof}
	}
	if oneof := CloudSpannerProperties_ToProto(mapCtx, in.CloudSpanner); oneof != nil {
		out.Properties = &pb.Connection_CloudSpanner{CloudSpanner: oneof}
	}
	if oneof := CloudResourceProperties_ToProto(mapCtx, in.CloudResource); oneof != nil {
		out.Properties = &pb.Connection_CloudResource{CloudResource: oneof}
	}
	if oneof := SparkProperties_ToProto(mapCtx, in.Spark); oneof != nil {
		out.Properties = &pb.Connection_Spark{Spark: oneof}
	}
	if oneof := SalesforceDataCloudProperties_ToProto(mapCtx, in.SalesforceDataCloud); oneof != nil {
		out.Properties = &pb.Connection_SalesforceDataCloud{SalesforceDataCloud: oneof}
	}
	out.CreationTime = direct.ValueOf(in.CreationTime)
	out.LastModifiedTime = direct.ValueOf(in.LastModifiedTime)
	out.HasCredential = direct.ValueOf(in.HasCredential)
	return out
}
func MetastoreServiceConfig_FromProto(mapCtx *direct.MapContext, in *pb.MetastoreServiceConfig) *krm.MetastoreServiceConfig {
	if in == nil {
		return nil
	}
	out := &krm.MetastoreServiceConfig{}
	out.MetastoreService = direct.LazyPtr(in.GetMetastoreService())
	return out
}
func MetastoreServiceConfig_ToProto(mapCtx *direct.MapContext, in *krm.MetastoreServiceConfig) *pb.MetastoreServiceConfig {
	if in == nil {
		return nil
	}
	out := &pb.MetastoreServiceConfig{}
	out.MetastoreService = direct.ValueOf(in.MetastoreService)
	return out
}
func SalesforceDataCloudProperties_FromProto(mapCtx *direct.MapContext, in *pb.SalesforceDataCloudProperties) *krm.SalesforceDataCloudProperties {
	if in == nil {
		return nil
	}
	out := &krm.SalesforceDataCloudProperties{}
	out.InstanceUri = direct.LazyPtr(in.GetInstanceUri())
	out.Identity = direct.LazyPtr(in.GetIdentity())
	out.TenantID = direct.LazyPtr(in.GetTenantId())
	return out
}
func SalesforceDataCloudProperties_ToProto(mapCtx *direct.MapContext, in *krm.SalesforceDataCloudProperties) *pb.SalesforceDataCloudProperties {
	if in == nil {
		return nil
	}
	out := &pb.SalesforceDataCloudProperties{}
	out.InstanceUri = direct.ValueOf(in.InstanceUri)
	out.Identity = direct.ValueOf(in.Identity)
	out.TenantId = direct.ValueOf(in.TenantID)
	return out
}
func SparkHistoryServerConfig_FromProto(mapCtx *direct.MapContext, in *pb.SparkHistoryServerConfig) *krm.SparkHistoryServerConfig {
	if in == nil {
		return nil
	}
	out := &krm.SparkHistoryServerConfig{}
	out.DataprocCluster = direct.LazyPtr(in.GetDataprocCluster())
	return out
}
func SparkHistoryServerConfig_ToProto(mapCtx *direct.MapContext, in *krm.SparkHistoryServerConfig) *pb.SparkHistoryServerConfig {
	if in == nil {
		return nil
	}
	out := &pb.SparkHistoryServerConfig{}
	out.DataprocCluster = direct.ValueOf(in.DataprocCluster)
	return out
}
func SparkProperties_FromProto(mapCtx *direct.MapContext, in *pb.SparkProperties) *krm.SparkProperties {
	if in == nil {
		return nil
	}
	out := &krm.SparkProperties{}
	out.ServiceAccountID = direct.LazyPtr(in.GetServiceAccountId())
	out.MetastoreServiceConfig = MetastoreServiceConfig_FromProto(mapCtx, in.GetMetastoreServiceConfig())
	out.SparkHistoryServerConfig = SparkHistoryServerConfig_FromProto(mapCtx, in.GetSparkHistoryServerConfig())
	return out
}
func SparkProperties_ToProto(mapCtx *direct.MapContext, in *krm.SparkProperties) *pb.SparkProperties {
	if in == nil {
		return nil
	}
	out := &pb.SparkProperties{}
	out.ServiceAccountId = direct.ValueOf(in.ServiceAccountID)
	out.MetastoreServiceConfig = MetastoreServiceConfig_ToProto(mapCtx, in.MetastoreServiceConfig)
	out.SparkHistoryServerConfig = SparkHistoryServerConfig_ToProto(mapCtx, in.SparkHistoryServerConfig)
	return out
}
