package clouddms

import (
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	alloydbv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddms/v1alpha1"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func MySQLConnectionProfile_FromProto(mapCtx *direct.MapContext, in *pb.MySqlConnectionProfile) *krm.MySQLConnectionProfile {
	if in == nil {
		return nil
	}
	out := &krm.MySQLConnectionProfile{}
	out.Host = direct.LazyPtr(in.GetHost())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	out.SSL = SSLConfig_FromProto(mapCtx, in.GetSsl())
	if in.GetCloudSqlId() != "" {
		out.CloudSQLInstanceRef = &refsv1beta1.SQLInstanceRef{
			External: in.GetCloudSqlId(),
		}
	}
	return out
}

func MySQLConnectionProfile_ToProto(mapCtx *direct.MapContext, in *krm.MySQLConnectionProfile) *pb.MySqlConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.MySqlConnectionProfile{}
	out.Host = direct.ValueOf(in.Host)
	out.Port = direct.ValueOf(in.Port)
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	out.Ssl = SSLConfig_ToProto(mapCtx, in.SSL)
	if in.CloudSQLInstanceRef != nil {
		out.CloudSqlId = in.CloudSQLInstanceRef.External
	}
	return out
}

func PostgreSQLConnectionProfile_FromProto(mapCtx *direct.MapContext, in *pb.PostgreSqlConnectionProfile) *krm.PostgreSQLConnectionProfile {
	if in == nil {
		return nil
	}
	out := &krm.PostgreSQLConnectionProfile{}
	out.Host = direct.LazyPtr(in.GetHost())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	out.SSL = SSLConfig_FromProto(mapCtx, in.GetSsl())
	if in.GetCloudSqlId() != "" {
		out.CloudSQLInstanceRef = &refsv1beta1.SQLInstanceRef{
			External: in.GetCloudSqlId(),
		}
	}
	if in.GetStaticIpConnectivity() != nil {
		out.StaticIPConnectivity = StaticIPConnectivity_FromProto(mapCtx, in.GetStaticIpConnectivity())
	}
	if in.GetPrivateServiceConnectConnectivity() != nil {
		out.PrivateServiceConnectConnectivity = PrivateServiceConnectConnectivity_FromProto(mapCtx, in.GetPrivateServiceConnectConnectivity())
	}
	return out
}

func PostgreSQLConnectionProfile_ToProto(mapCtx *direct.MapContext, in *krm.PostgreSQLConnectionProfile) *pb.PostgreSqlConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.PostgreSqlConnectionProfile{}
	out.Host = direct.ValueOf(in.Host)
	out.Port = direct.ValueOf(in.Port)
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	out.Ssl = SSLConfig_ToProto(mapCtx, in.SSL)
	if in.CloudSQLInstanceRef != nil {
		out.CloudSqlId = in.CloudSQLInstanceRef.External
	}
	if in.StaticIPConnectivity != nil {
		out.Connectivity = &pb.PostgreSqlConnectionProfile_StaticIpConnectivity{
			StaticIpConnectivity: StaticIPConnectivity_ToProto(mapCtx, in.StaticIPConnectivity),
		}
	} else if in.PrivateServiceConnectConnectivity != nil {
		out.Connectivity = &pb.PostgreSqlConnectionProfile_PrivateServiceConnectConnectivity{
			PrivateServiceConnectConnectivity: PrivateServiceConnectConnectivity_ToProto(mapCtx, in.PrivateServiceConnectConnectivity),
		}
	}
	return out
}

func AlloyDbConnectionProfile_FromProto(mapCtx *direct.MapContext, in *pb.AlloyDbConnectionProfile) *krm.AlloyDbConnectionProfile {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDbConnectionProfile{}
	if in.GetClusterId() != "" {
		out.ClusterRef = &alloydbv1beta1.ClusterRef{
			External: in.GetClusterId(),
		}
	}
	out.Settings = AlloyDbSettings_FromProto(mapCtx, in.GetSettings())
	return out
}

func AlloyDbConnectionProfile_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDbConnectionProfile) *pb.AlloyDbConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.AlloyDbConnectionProfile{}
	if in.ClusterRef != nil {
		out.ClusterId = in.ClusterRef.External
	}
	out.Settings = AlloyDbSettings_ToProto(mapCtx, in.Settings)
	return out
}

func AlloyDbSettings_FromProto(mapCtx *direct.MapContext, in *pb.AlloyDbSettings) *krm.AlloyDbSettings {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDbSettings{}
	out.InitialUser = AlloyDbSettings_UserPassword_FromProto(mapCtx, in.GetInitialUser())
	if in.GetVpcNetwork() != "" {
		out.VPCNetworkRef = &computev1beta1.ComputeNetworkRef{
			External: in.GetVpcNetwork(),
		}
	}
	out.Labels = in.GetLabels()
	out.PrimaryInstanceSettings = AlloyDbSettings_PrimaryInstanceSettings_FromProto(mapCtx, in.GetPrimaryInstanceSettings())
	out.EncryptionConfig = AlloyDbSettings_EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	return out
}

func AlloyDbSettings_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDbSettings) *pb.AlloyDbSettings {
	if in == nil {
		return nil
	}
	out := &pb.AlloyDbSettings{}
	out.InitialUser = AlloyDbSettings_UserPassword_ToProto(mapCtx, in.InitialUser)
	if in.VPCNetworkRef != nil {
		out.VpcNetwork = in.VPCNetworkRef.External
	}
	out.Labels = in.Labels
	out.PrimaryInstanceSettings = AlloyDbSettings_PrimaryInstanceSettings_ToProto(mapCtx, in.PrimaryInstanceSettings)
	out.EncryptionConfig = AlloyDbSettings_EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	return out
}

func CloudSQLSettings_FromProto(mapCtx *direct.MapContext, in *pb.CloudSqlSettings) *krm.CloudSQLSettings {
	if in == nil {
		return nil
	}
	out := &krm.CloudSQLSettings{}
	if in.GetDatabaseVersion() != pb.CloudSqlSettings_SQL_DATABASE_VERSION_UNSPECIFIED {
		out.DatabaseVersion = direct.LazyPtr(in.GetDatabaseVersion().String())
	}
	out.UserLabels = in.GetUserLabels()
	out.Tier = direct.LazyPtr(in.GetTier())
	if in.GetStorageAutoResizeLimit() != nil {
		out.StorageAutoResizeLimit = direct.LazyPtr(in.GetStorageAutoResizeLimit().GetValue())
	}
	if in.GetActivationPolicy() != pb.CloudSqlSettings_SQL_ACTIVATION_POLICY_UNSPECIFIED {
		out.ActivationPolicy = direct.LazyPtr(in.GetActivationPolicy().String())
	}
	out.IPConfig = SQLIPConfig_FromProto(mapCtx, in.GetIpConfig())
	if in.GetAutoStorageIncrease() != nil {
		out.AutoStorageIncrease = direct.LazyPtr(in.GetAutoStorageIncrease().GetValue())
	}
	out.DatabaseFlags = in.GetDatabaseFlags()
	if in.GetDataDiskType() != pb.CloudSqlSettings_SQL_DATA_DISK_TYPE_UNSPECIFIED {
		out.DataDiskType = direct.LazyPtr(in.GetDataDiskType().String())
	}
	if in.GetDataDiskSizeGb() != nil {
		out.DataDiskSizeGB = direct.LazyPtr(in.GetDataDiskSizeGb().GetValue())
	}
	out.Zone = direct.LazyPtr(in.GetZone())
	out.SecondaryZone = direct.LazyPtr(in.GetSecondaryZone())
	if in.GetSourceId() != "" {
		out.SourceRef = &krm.CloudDMSConnectionProfileRef{
			External: in.GetSourceId(),
		}
	}
	out.RootPassword = direct.LazyPtr(in.GetRootPassword())
	out.Collation = direct.LazyPtr(in.GetCollation())
	if in.GetCmekKeyName() != "" {
		out.KMSKeyRef = &refsv1beta1.KMSCryptoKeyRef{
			External: in.GetCmekKeyName(),
		}
	}
	if in.GetAvailabilityType() != pb.CloudSqlSettings_SQL_AVAILABILITY_TYPE_UNSPECIFIED {
		out.AvailabilityType = direct.LazyPtr(in.GetAvailabilityType().String())
	}
	if in.GetEdition() != pb.CloudSqlSettings_EDITION_UNSPECIFIED {
		out.Edition = direct.LazyPtr(in.GetEdition().String())
	}
	return out
}

func CloudSQLSettings_ToProto(mapCtx *direct.MapContext, in *krm.CloudSQLSettings) *pb.CloudSqlSettings {
	if in == nil {
		return nil
	}
	out := &pb.CloudSqlSettings{}
	if in.DatabaseVersion != nil {
		out.DatabaseVersion = pb.CloudSqlSettings_SqlDatabaseVersion(pb.CloudSqlSettings_SqlDatabaseVersion_value[*in.DatabaseVersion])
	}
	out.UserLabels = in.UserLabels
	out.Tier = direct.ValueOf(in.Tier)
	if in.StorageAutoResizeLimit != nil {
		out.StorageAutoResizeLimit = &wrapperspb.Int64Value{Value: *in.StorageAutoResizeLimit}
	}
	if in.ActivationPolicy != nil {
		out.ActivationPolicy = pb.CloudSqlSettings_SqlActivationPolicy(pb.CloudSqlSettings_SqlActivationPolicy_value[*in.ActivationPolicy])
	}
	out.IpConfig = SQLIPConfig_ToProto(mapCtx, in.IPConfig)
	if in.AutoStorageIncrease != nil {
		out.AutoStorageIncrease = &wrapperspb.BoolValue{Value: *in.AutoStorageIncrease}
	}
	out.DatabaseFlags = in.DatabaseFlags
	if in.DataDiskType != nil {
		out.DataDiskType = pb.CloudSqlSettings_SqlDataDiskType(pb.CloudSqlSettings_SqlDataDiskType_value[*in.DataDiskType])
	}
	if in.DataDiskSizeGB != nil {
		out.DataDiskSizeGb = &wrapperspb.Int64Value{Value: *in.DataDiskSizeGB}
	}
	out.Zone = direct.ValueOf(in.Zone)
	out.SecondaryZone = direct.ValueOf(in.SecondaryZone)
	if in.SourceRef != nil {
		out.SourceId = in.SourceRef.External
	}
	out.RootPassword = direct.ValueOf(in.RootPassword)
	out.Collation = direct.ValueOf(in.Collation)
	if in.KMSKeyRef != nil {
		out.CmekKeyName = in.KMSKeyRef.External
	}
	if in.AvailabilityType != nil {
		out.AvailabilityType = pb.CloudSqlSettings_SqlAvailabilityType(pb.CloudSqlSettings_SqlAvailabilityType_value[*in.AvailabilityType])
	}
	if in.Edition != nil {
		out.Edition = pb.CloudSqlSettings_Edition(pb.CloudSqlSettings_Edition_value[*in.Edition])
	}
	return out
}

func PrivateConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnectivity) *krm.PrivateConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.PrivateConnectivity{}
	if in.GetPrivateConnection() != "" {
		out.PrivateConnectionRef = &krm.PrivateConnectionRef{
			External: in.GetPrivateConnection(),
		}
	}
	return out
}

func PrivateConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.PrivateConnectivity) *pb.PrivateConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnectivity{}
	if in.PrivateConnectionRef != nil {
		out.PrivateConnection = in.PrivateConnectionRef.External
	}
	return out
}

func PrivateServiceConnectConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.PrivateServiceConnectConnectivity) *krm.PrivateServiceConnectConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.PrivateServiceConnectConnectivity{}
	if in.GetServiceAttachment() != "" {
		out.ServiceAttachmentRef = &refsv1beta1.ComputeServiceAttachmentRef{
			External: in.GetServiceAttachment(),
		}
	}
	return out
}

func PrivateServiceConnectConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.PrivateServiceConnectConnectivity) *pb.PrivateServiceConnectConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.PrivateServiceConnectConnectivity{}
	if in.ServiceAttachmentRef != nil {
		out.ServiceAttachment = in.ServiceAttachmentRef.External
	}
	return out
}

func SQLIPConfig_FromProto(mapCtx *direct.MapContext, in *pb.SqlIpConfig) *krm.SQLIPConfig {
	if in == nil {
		return nil
	}
	out := &krm.SQLIPConfig{}
	if in.GetEnableIpv4() != nil {
		out.EnableIPV4 = direct.LazyPtr(in.GetEnableIpv4().GetValue())
	}
	if in.GetPrivateNetwork() != "" {
		out.PrivateNetworkRef = &computev1beta1.ComputeNetworkRef{
			External: in.GetPrivateNetwork(),
		}
	}
	out.AllocatedIPRange = direct.LazyPtr(in.GetAllocatedIpRange())
	if in.GetRequireSsl() != nil {
		out.RequireSSL = direct.LazyPtr(in.GetRequireSsl().GetValue())
	}
	if in.GetAuthorizedNetworks() != nil {
		for _, v := range in.GetAuthorizedNetworks() {
			out.AuthorizedNetworks = append(out.AuthorizedNetworks, *SQLAclEntry_FromProto(mapCtx, v))
		}
	}
	return out
}

func SQLIPConfig_ToProto(mapCtx *direct.MapContext, in *krm.SQLIPConfig) *pb.SqlIpConfig {
	if in == nil {
		return nil
	}
	out := &pb.SqlIpConfig{}
	if in.EnableIPV4 != nil {
		out.EnableIpv4 = &wrapperspb.BoolValue{Value: *in.EnableIPV4}
	}
	if in.PrivateNetworkRef != nil {
		out.PrivateNetwork = in.PrivateNetworkRef.External
	}
	out.AllocatedIpRange = direct.ValueOf(in.AllocatedIPRange)
	if in.RequireSSL != nil {
		out.RequireSsl = &wrapperspb.BoolValue{Value: *in.RequireSSL}
	}
	if in.AuthorizedNetworks != nil {
		for _, v := range in.AuthorizedNetworks {
			out.AuthorizedNetworks = append(out.AuthorizedNetworks, SQLAclEntry_ToProto(mapCtx, &v))
		}
	}
	return out
}

func AlloyDbSettings_EncryptionConfig_FromProto(mapCtx *direct.MapContext, in *pb.AlloyDbSettings_EncryptionConfig) *krm.AlloyDbSettings_EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDbSettings_EncryptionConfig{}
	if in.GetKmsKeyName() != "" {
		out.KMSKeyRef = &refsv1beta1.KMSCryptoKeyRef{
			External: in.GetKmsKeyName(),
		}
	}
	return out
}

func AlloyDbSettings_EncryptionConfig_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDbSettings_EncryptionConfig) *pb.AlloyDbSettings_EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.AlloyDbSettings_EncryptionConfig{}
	if in.KMSKeyRef != nil {
		out.KmsKeyName = in.KMSKeyRef.External
	}
	return out
}
