package bigtable

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/bigtable/v1beta1"

	pb "google.golang.org/genproto/googleapis/bigtable/admin/v2"
)

//	func CreateInstanceRequest_FromProto(ctx *MapContext, in *pb.CreateInstanceRequest) *krm.CreateInstanceRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.CreateInstanceRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.InstanceId = LazyPtr(in.InstanceId)
//		out.Instance = Instance_FromProto(ctx, in.Instance)
//		out.Clusters = CreateInstanceRequest_ClustersEntry_FromProto(ctx, in.Clusters)
//		return out
//	}
//
//	func CreateInstanceRequest_ToProto(ctx *MapContext, in *krm.CreateInstanceRequest) *pb.CreateInstanceRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.CreateInstanceRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.InstanceId = LazyPtr(in.InstanceId)
//		out.Instance = Instance_ToProto(ctx, in.Instance)
//		out.Clusters = CreateInstanceRequest_ClustersEntry_ToProto(ctx, in.Clusters)
//		return out
//	}
//
//	func CreateInstanceRequest_ClustersEntry_FromProto(ctx *MapContext, in *pb.CreateInstanceRequest_ClustersEntry) *krm.CreateInstanceRequest_ClustersEntry {
//		if in == nil {
//			return nil
//		}
//		out := &krm.CreateInstanceRequest_ClustersEntry{}
//		out.Key = LazyPtr(in.Key)
//		out.Value = Cluster_FromProto(ctx, in.Value)
//		return out
//	}
//
//	func CreateInstanceRequest_ClustersEntry_ToProto(ctx *MapContext, in *krm.CreateInstanceRequest_ClustersEntry) *pb.CreateInstanceRequest_ClustersEntry {
//		if in == nil {
//			return nil
//		}
//		out := &pb.CreateInstanceRequest_ClustersEntry{}
//		out.Key = LazyPtr(in.Key)
//		out.Value = Cluster_ToProto(ctx, in.Value)
//		return out
//	}
//
//	func GetInstanceRequest_FromProto(ctx *MapContext, in *pb.GetInstanceRequest) *krm.GetInstanceRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.GetInstanceRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func GetInstanceRequest_ToProto(ctx *MapContext, in *krm.GetInstanceRequest) *pb.GetInstanceRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.GetInstanceRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func ListInstancesRequest_FromProto(ctx *MapContext, in *pb.ListInstancesRequest) *krm.ListInstancesRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.ListInstancesRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.PageToken = LazyPtr(in.PageToken)
//		return out
//	}
//
//	func ListInstancesRequest_ToProto(ctx *MapContext, in *krm.ListInstancesRequest) *pb.ListInstancesRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.ListInstancesRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.PageToken = LazyPtr(in.PageToken)
//		return out
//	}
//
//	func ListInstancesResponse_FromProto(ctx *MapContext, in *pb.ListInstancesResponse) *krm.ListInstancesResponse {
//		if in == nil {
//			return nil
//		}
//		out := &krm.ListInstancesResponse{}
//		out.Instances = Instance_FromProto(ctx, in.Instances)
//		out.FailedLocations = LazyPtr(in.FailedLocations)
//		out.NextPageToken = LazyPtr(in.NextPageToken)
//		return out
//	}
//
//	func ListInstancesResponse_ToProto(ctx *MapContext, in *krm.ListInstancesResponse) *pb.ListInstancesResponse {
//		if in == nil {
//			return nil
//		}
//		out := &pb.ListInstancesResponse{}
//		out.Instances = Instance_ToProto(ctx, in.Instances)
//		out.FailedLocations = LazyPtr(in.FailedLocations)
//		out.NextPageToken = LazyPtr(in.NextPageToken)
//		return out
//	}
//
//	func PartialUpdateInstanceRequest_FromProto(ctx *MapContext, in *pb.PartialUpdateInstanceRequest) *krm.PartialUpdateInstanceRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.PartialUpdateInstanceRequest{}
//		out.Instance = Instance_FromProto(ctx, in.Instance)
//		out.UpdateMask = FieldMask_FromProto(ctx, in.UpdateMask)
//		return out
//	}
//
//	func PartialUpdateInstanceRequest_ToProto(ctx *MapContext, in *krm.PartialUpdateInstanceRequest) *pb.PartialUpdateInstanceRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.PartialUpdateInstanceRequest{}
//		out.Instance = Instance_ToProto(ctx, in.Instance)
//		out.UpdateMask = FieldMask_ToProto(ctx, in.UpdateMask)
//		return out
//	}
//
//	func DeleteInstanceRequest_FromProto(ctx *MapContext, in *pb.DeleteInstanceRequest) *krm.DeleteInstanceRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.DeleteInstanceRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func DeleteInstanceRequest_ToProto(ctx *MapContext, in *krm.DeleteInstanceRequest) *pb.DeleteInstanceRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.DeleteInstanceRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func CreateClusterRequest_FromProto(ctx *MapContext, in *pb.CreateClusterRequest) *krm.CreateClusterRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.CreateClusterRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.ClusterId = LazyPtr(in.ClusterId)
//		out.Cluster = Cluster_FromProto(ctx, in.Cluster)
//		return out
//	}
//
//	func CreateClusterRequest_ToProto(ctx *MapContext, in *krm.CreateClusterRequest) *pb.CreateClusterRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.CreateClusterRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.ClusterId = LazyPtr(in.ClusterId)
//		out.Cluster = Cluster_ToProto(ctx, in.Cluster)
//		return out
//	}
//
//	func GetClusterRequest_FromProto(ctx *MapContext, in *pb.GetClusterRequest) *krm.GetClusterRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.GetClusterRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func GetClusterRequest_ToProto(ctx *MapContext, in *krm.GetClusterRequest) *pb.GetClusterRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.GetClusterRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func ListClustersRequest_FromProto(ctx *MapContext, in *pb.ListClustersRequest) *krm.ListClustersRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.ListClustersRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.PageToken = LazyPtr(in.PageToken)
//		return out
//	}
//
//	func ListClustersRequest_ToProto(ctx *MapContext, in *krm.ListClustersRequest) *pb.ListClustersRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.ListClustersRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.PageToken = LazyPtr(in.PageToken)
//		return out
//	}
//
//	func ListClustersResponse_FromProto(ctx *MapContext, in *pb.ListClustersResponse) *krm.ListClustersResponse {
//		if in == nil {
//			return nil
//		}
//		out := &krm.ListClustersResponse{}
//		out.Clusters = Cluster_FromProto(ctx, in.Clusters)
//		out.FailedLocations = LazyPtr(in.FailedLocations)
//		out.NextPageToken = LazyPtr(in.NextPageToken)
//		return out
//	}
//
//	func ListClustersResponse_ToProto(ctx *MapContext, in *krm.ListClustersResponse) *pb.ListClustersResponse {
//		if in == nil {
//			return nil
//		}
//		out := &pb.ListClustersResponse{}
//		out.Clusters = Cluster_ToProto(ctx, in.Clusters)
//		out.FailedLocations = LazyPtr(in.FailedLocations)
//		out.NextPageToken = LazyPtr(in.NextPageToken)
//		return out
//	}
//
//	func DeleteClusterRequest_FromProto(ctx *MapContext, in *pb.DeleteClusterRequest) *krm.DeleteClusterRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.DeleteClusterRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func DeleteClusterRequest_ToProto(ctx *MapContext, in *krm.DeleteClusterRequest) *pb.DeleteClusterRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.DeleteClusterRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func CreateInstanceMetadata_FromProto(ctx *MapContext, in *pb.CreateInstanceMetadata) *krm.CreateInstanceMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &krm.CreateInstanceMetadata{}
//		out.OriginalRequest = CreateInstanceRequest_FromProto(ctx, in.OriginalRequest)
//		out.RequestTime = Timestamp_FromProto(ctx, in.RequestTime)
//		out.FinishTime = Timestamp_FromProto(ctx, in.FinishTime)
//		return out
//	}
//
//	func CreateInstanceMetadata_ToProto(ctx *MapContext, in *krm.CreateInstanceMetadata) *pb.CreateInstanceMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &pb.CreateInstanceMetadata{}
//		out.OriginalRequest = CreateInstanceRequest_ToProto(ctx, in.OriginalRequest)
//		out.RequestTime = Timestamp_ToProto(ctx, in.RequestTime)
//		out.FinishTime = Timestamp_ToProto(ctx, in.FinishTime)
//		return out
//	}
//
//	func UpdateInstanceMetadata_FromProto(ctx *MapContext, in *pb.UpdateInstanceMetadata) *krm.UpdateInstanceMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &krm.UpdateInstanceMetadata{}
//		out.OriginalRequest = PartialUpdateInstanceRequest_FromProto(ctx, in.OriginalRequest)
//		out.RequestTime = Timestamp_FromProto(ctx, in.RequestTime)
//		out.FinishTime = Timestamp_FromProto(ctx, in.FinishTime)
//		return out
//	}
//
//	func UpdateInstanceMetadata_ToProto(ctx *MapContext, in *krm.UpdateInstanceMetadata) *pb.UpdateInstanceMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &pb.UpdateInstanceMetadata{}
//		out.OriginalRequest = PartialUpdateInstanceRequest_ToProto(ctx, in.OriginalRequest)
//		out.RequestTime = Timestamp_ToProto(ctx, in.RequestTime)
//		out.FinishTime = Timestamp_ToProto(ctx, in.FinishTime)
//		return out
//	}
//
//	func CreateClusterMetadata_FromProto(ctx *MapContext, in *pb.CreateClusterMetadata) *krm.CreateClusterMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &krm.CreateClusterMetadata{}
//		out.OriginalRequest = CreateClusterRequest_FromProto(ctx, in.OriginalRequest)
//		out.RequestTime = Timestamp_FromProto(ctx, in.RequestTime)
//		out.FinishTime = Timestamp_FromProto(ctx, in.FinishTime)
//		out.Tables = CreateClusterMetadata_TablesEntry_FromProto(ctx, in.Tables)
//		return out
//	}
//
//	func CreateClusterMetadata_ToProto(ctx *MapContext, in *krm.CreateClusterMetadata) *pb.CreateClusterMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &pb.CreateClusterMetadata{}
//		out.OriginalRequest = CreateClusterRequest_ToProto(ctx, in.OriginalRequest)
//		out.RequestTime = Timestamp_ToProto(ctx, in.RequestTime)
//		out.FinishTime = Timestamp_ToProto(ctx, in.FinishTime)
//		out.Tables = CreateClusterMetadata_TablesEntry_ToProto(ctx, in.Tables)
//		return out
//	}
//
//	func CreateClusterMetadata_TableProgress_FromProto(ctx *MapContext, in *pb.CreateClusterMetadata_TableProgress) *krm.CreateClusterMetadata_TableProgress {
//		if in == nil {
//			return nil
//		}
//		out := &krm.CreateClusterMetadata_TableProgress{}
//		out.EstimatedSizeBytes = LazyPtr(in.EstimatedSizeBytes)
//		out.EstimatedCopiedBytes = LazyPtr(in.EstimatedCopiedBytes)
//		out.State = Enum_FromProto(ctx, &in.State)
//		return out
//	}
//
//	func CreateClusterMetadata_TableProgress_ToProto(ctx *MapContext, in *krm.CreateClusterMetadata_TableProgress) *pb.CreateClusterMetadata_TableProgress {
//		if in == nil {
//			return nil
//		}
//		out := &pb.CreateClusterMetadata_TableProgress{}
//		out.EstimatedSizeBytes = LazyPtr(in.EstimatedSizeBytes)
//		out.EstimatedCopiedBytes = LazyPtr(in.EstimatedCopiedBytes)
//		out.State = Enum_ToProto(ctx, &in.State)
//		return out
//	}
//
//	func CreateClusterMetadata_TablesEntry_FromProto(ctx *MapContext, in *pb.CreateClusterMetadata_TablesEntry) *krm.CreateClusterMetadata_TablesEntry {
//		if in == nil {
//			return nil
//		}
//		out := &krm.CreateClusterMetadata_TablesEntry{}
//		out.Key = LazyPtr(in.Key)
//		out.Value = CreateClusterMetadata_TableProgress_FromProto(ctx, in.Value)
//		return out
//	}
//
//	func CreateClusterMetadata_TablesEntry_ToProto(ctx *MapContext, in *krm.CreateClusterMetadata_TablesEntry) *pb.CreateClusterMetadata_TablesEntry {
//		if in == nil {
//			return nil
//		}
//		out := &pb.CreateClusterMetadata_TablesEntry{}
//		out.Key = LazyPtr(in.Key)
//		out.Value = CreateClusterMetadata_TableProgress_ToProto(ctx, in.Value)
//		return out
//	}
//
//	func UpdateClusterMetadata_FromProto(ctx *MapContext, in *pb.UpdateClusterMetadata) *krm.UpdateClusterMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &krm.UpdateClusterMetadata{}
//		out.OriginalRequest = Cluster_FromProto(ctx, in.OriginalRequest)
//		out.RequestTime = Timestamp_FromProto(ctx, in.RequestTime)
//		out.FinishTime = Timestamp_FromProto(ctx, in.FinishTime)
//		return out
//	}
//
//	func UpdateClusterMetadata_ToProto(ctx *MapContext, in *krm.UpdateClusterMetadata) *pb.UpdateClusterMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &pb.UpdateClusterMetadata{}
//		out.OriginalRequest = Cluster_ToProto(ctx, in.OriginalRequest)
//		out.RequestTime = Timestamp_ToProto(ctx, in.RequestTime)
//		out.FinishTime = Timestamp_ToProto(ctx, in.FinishTime)
//		return out
//	}
//
//	func PartialUpdateClusterMetadata_FromProto(ctx *MapContext, in *pb.PartialUpdateClusterMetadata) *krm.PartialUpdateClusterMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &krm.PartialUpdateClusterMetadata{}
//		out.RequestTime = Timestamp_FromProto(ctx, in.RequestTime)
//		out.FinishTime = Timestamp_FromProto(ctx, in.FinishTime)
//		out.OriginalRequest = PartialUpdateClusterRequest_FromProto(ctx, in.OriginalRequest)
//		return out
//	}
//
//	func PartialUpdateClusterMetadata_ToProto(ctx *MapContext, in *krm.PartialUpdateClusterMetadata) *pb.PartialUpdateClusterMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &pb.PartialUpdateClusterMetadata{}
//		out.RequestTime = Timestamp_ToProto(ctx, in.RequestTime)
//		out.FinishTime = Timestamp_ToProto(ctx, in.FinishTime)
//		out.OriginalRequest = PartialUpdateClusterRequest_ToProto(ctx, in.OriginalRequest)
//		return out
//	}
//
//	func PartialUpdateClusterRequest_FromProto(ctx *MapContext, in *pb.PartialUpdateClusterRequest) *krm.PartialUpdateClusterRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.PartialUpdateClusterRequest{}
//		out.Cluster = Cluster_FromProto(ctx, in.Cluster)
//		out.UpdateMask = FieldMask_FromProto(ctx, in.UpdateMask)
//		return out
//	}
//
//	func PartialUpdateClusterRequest_ToProto(ctx *MapContext, in *krm.PartialUpdateClusterRequest) *pb.PartialUpdateClusterRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.PartialUpdateClusterRequest{}
//		out.Cluster = Cluster_ToProto(ctx, in.Cluster)
//		out.UpdateMask = FieldMask_ToProto(ctx, in.UpdateMask)
//		return out
//	}
//
//	func CreateAppProfileRequest_FromProto(ctx *MapContext, in *pb.CreateAppProfileRequest) *krm.CreateAppProfileRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.CreateAppProfileRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.AppProfileId = LazyPtr(in.AppProfileId)
//		out.AppProfile = AppProfile_FromProto(ctx, in.AppProfile)
//		out.IgnoreWarnings = LazyPtr(in.IgnoreWarnings)
//		return out
//	}
//
//	func CreateAppProfileRequest_ToProto(ctx *MapContext, in *krm.CreateAppProfileRequest) *pb.CreateAppProfileRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.CreateAppProfileRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.AppProfileId = LazyPtr(in.AppProfileId)
//		out.AppProfile = AppProfile_ToProto(ctx, in.AppProfile)
//		out.IgnoreWarnings = LazyPtr(in.IgnoreWarnings)
//		return out
//	}
//
//	func GetAppProfileRequest_FromProto(ctx *MapContext, in *pb.GetAppProfileRequest) *krm.GetAppProfileRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.GetAppProfileRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func GetAppProfileRequest_ToProto(ctx *MapContext, in *krm.GetAppProfileRequest) *pb.GetAppProfileRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.GetAppProfileRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func ListAppProfilesRequest_FromProto(ctx *MapContext, in *pb.ListAppProfilesRequest) *krm.ListAppProfilesRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.ListAppProfilesRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.PageSize = LazyPtr(in.PageSize)
//		out.PageToken = LazyPtr(in.PageToken)
//		return out
//	}
//
//	func ListAppProfilesRequest_ToProto(ctx *MapContext, in *krm.ListAppProfilesRequest) *pb.ListAppProfilesRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.ListAppProfilesRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.PageSize = LazyPtr(in.PageSize)
//		out.PageToken = LazyPtr(in.PageToken)
//		return out
//	}
//
//	func ListAppProfilesResponse_FromProto(ctx *MapContext, in *pb.ListAppProfilesResponse) *krm.ListAppProfilesResponse {
//		if in == nil {
//			return nil
//		}
//		out := &krm.ListAppProfilesResponse{}
//		out.AppProfiles = AppProfile_FromProto(ctx, in.AppProfiles)
//		out.NextPageToken = LazyPtr(in.NextPageToken)
//		out.FailedLocations = LazyPtr(in.FailedLocations)
//		return out
//	}
//
//	func ListAppProfilesResponse_ToProto(ctx *MapContext, in *krm.ListAppProfilesResponse) *pb.ListAppProfilesResponse {
//		if in == nil {
//			return nil
//		}
//		out := &pb.ListAppProfilesResponse{}
//		out.AppProfiles = AppProfile_ToProto(ctx, in.AppProfiles)
//		out.NextPageToken = LazyPtr(in.NextPageToken)
//		out.FailedLocations = LazyPtr(in.FailedLocations)
//		return out
//	}
//
//	func UpdateAppProfileRequest_FromProto(ctx *MapContext, in *pb.UpdateAppProfileRequest) *krm.UpdateAppProfileRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.UpdateAppProfileRequest{}
//		out.AppProfile = AppProfile_FromProto(ctx, in.AppProfile)
//		out.UpdateMask = FieldMask_FromProto(ctx, in.UpdateMask)
//		out.IgnoreWarnings = LazyPtr(in.IgnoreWarnings)
//		return out
//	}
//
//	func UpdateAppProfileRequest_ToProto(ctx *MapContext, in *krm.UpdateAppProfileRequest) *pb.UpdateAppProfileRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.UpdateAppProfileRequest{}
//		out.AppProfile = AppProfile_ToProto(ctx, in.AppProfile)
//		out.UpdateMask = FieldMask_ToProto(ctx, in.UpdateMask)
//		out.IgnoreWarnings = LazyPtr(in.IgnoreWarnings)
//		return out
//	}
//
//	func DeleteAppProfileRequest_FromProto(ctx *MapContext, in *pb.DeleteAppProfileRequest) *krm.DeleteAppProfileRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.DeleteAppProfileRequest{}
//		out.Name = LazyPtr(in.Name)
//		out.IgnoreWarnings = LazyPtr(in.IgnoreWarnings)
//		return out
//	}
//
//	func DeleteAppProfileRequest_ToProto(ctx *MapContext, in *krm.DeleteAppProfileRequest) *pb.DeleteAppProfileRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.DeleteAppProfileRequest{}
//		out.Name = LazyPtr(in.Name)
//		out.IgnoreWarnings = LazyPtr(in.IgnoreWarnings)
//		return out
//	}
//
//	func UpdateAppProfileMetadata_FromProto(ctx *MapContext, in *pb.UpdateAppProfileMetadata) *krm.UpdateAppProfileMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &krm.UpdateAppProfileMetadata{}
//		return out
//	}
//
//	func UpdateAppProfileMetadata_ToProto(ctx *MapContext, in *krm.UpdateAppProfileMetadata) *pb.UpdateAppProfileMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &pb.UpdateAppProfileMetadata{}
//		return out
//	}
//
//	func ListHotTabletsRequest_FromProto(ctx *MapContext, in *pb.ListHotTabletsRequest) *krm.ListHotTabletsRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.ListHotTabletsRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.StartTime = Timestamp_FromProto(ctx, in.StartTime)
//		out.EndTime = Timestamp_FromProto(ctx, in.EndTime)
//		out.PageSize = LazyPtr(in.PageSize)
//		out.PageToken = LazyPtr(in.PageToken)
//		return out
//	}
//
//	func ListHotTabletsRequest_ToProto(ctx *MapContext, in *krm.ListHotTabletsRequest) *pb.ListHotTabletsRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.ListHotTabletsRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.StartTime = Timestamp_ToProto(ctx, in.StartTime)
//		out.EndTime = Timestamp_ToProto(ctx, in.EndTime)
//		out.PageSize = LazyPtr(in.PageSize)
//		out.PageToken = LazyPtr(in.PageToken)
//		return out
//	}
//
//	func ListHotTabletsResponse_FromProto(ctx *MapContext, in *pb.ListHotTabletsResponse) *krm.ListHotTabletsResponse {
//		if in == nil {
//			return nil
//		}
//		out := &krm.ListHotTabletsResponse{}
//		out.HotTablets = HotTablet_FromProto(ctx, in.HotTablets)
//		out.NextPageToken = LazyPtr(in.NextPageToken)
//		return out
//	}
//
//	func ListHotTabletsResponse_ToProto(ctx *MapContext, in *krm.ListHotTabletsResponse) *pb.ListHotTabletsResponse {
//		if in == nil {
//			return nil
//		}
//		out := &pb.ListHotTabletsResponse{}
//		out.HotTablets = HotTablet_ToProto(ctx, in.HotTablets)
//		out.NextPageToken = LazyPtr(in.NextPageToken)
//		return out
//	}
//
//	func OperationProgress_FromProto(ctx *MapContext, in *pb.OperationProgress) *krm.OperationProgress {
//		if in == nil {
//			return nil
//		}
//		out := &krm.OperationProgress{}
//		out.ProgressPercent = LazyPtr(in.ProgressPercent)
//		out.StartTime = Timestamp_FromProto(ctx, in.StartTime)
//		out.EndTime = Timestamp_FromProto(ctx, in.EndTime)
//		return out
//	}
//
//	func OperationProgress_ToProto(ctx *MapContext, in *krm.OperationProgress) *pb.OperationProgress {
//		if in == nil {
//			return nil
//		}
//		out := &pb.OperationProgress{}
//		out.ProgressPercent = LazyPtr(in.ProgressPercent)
//		out.StartTime = Timestamp_ToProto(ctx, in.StartTime)
//		out.EndTime = Timestamp_ToProto(ctx, in.EndTime)
//		return out
//	}
//
//	func RestoreTableRequest_FromProto(ctx *MapContext, in *pb.RestoreTableRequest) *krm.RestoreTableRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.RestoreTableRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.TableId = LazyPtr(in.TableId)
//		out.Backup = LazyPtr(in.Backup)
//		return out
//	}
//
//	func RestoreTableRequest_ToProto(ctx *MapContext, in *krm.RestoreTableRequest) *pb.RestoreTableRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.RestoreTableRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.TableId = LazyPtr(in.TableId)
//		out.Backup = LazyPtr(in.Backup)
//		return out
//	}
//
//	func RestoreTableMetadata_FromProto(ctx *MapContext, in *pb.RestoreTableMetadata) *krm.RestoreTableMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &krm.RestoreTableMetadata{}
//		out.Name = LazyPtr(in.Name)
//		out.SourceType = Enum_FromProto(ctx, &in.SourceType)
//		out.BackupInfo = BackupInfo_FromProto(ctx, in.BackupInfo)
//		out.OptimizeTableOperationName = LazyPtr(in.OptimizeTableOperationName)
//		out.Progress = OperationProgress_FromProto(ctx, in.Progress)
//		return out
//	}
//
//	func RestoreTableMetadata_ToProto(ctx *MapContext, in *krm.RestoreTableMetadata) *pb.RestoreTableMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &pb.RestoreTableMetadata{}
//		out.Name = LazyPtr(in.Name)
//		out.SourceType = Enum_ToProto(ctx, &in.SourceType)
//		out.BackupInfo = BackupInfo_ToProto(ctx, in.BackupInfo)
//		out.OptimizeTableOperationName = LazyPtr(in.OptimizeTableOperationName)
//		out.Progress = OperationProgress_ToProto(ctx, in.Progress)
//		return out
//	}
//
//	func OptimizeRestoredTableMetadata_FromProto(ctx *MapContext, in *pb.OptimizeRestoredTableMetadata) *krm.OptimizeRestoredTableMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &krm.OptimizeRestoredTableMetadata{}
//		out.Name = LazyPtr(in.Name)
//		out.Progress = OperationProgress_FromProto(ctx, in.Progress)
//		return out
//	}
//
//	func OptimizeRestoredTableMetadata_ToProto(ctx *MapContext, in *krm.OptimizeRestoredTableMetadata) *pb.OptimizeRestoredTableMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &pb.OptimizeRestoredTableMetadata{}
//		out.Name = LazyPtr(in.Name)
//		out.Progress = OperationProgress_ToProto(ctx, in.Progress)
//		return out
//	}
//
//	func CreateTableRequest_FromProto(ctx *MapContext, in *pb.CreateTableRequest) *krm.CreateTableRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.CreateTableRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.TableId = LazyPtr(in.TableId)
//		out.Table = Table_FromProto(ctx, in.Table)
//		out.InitialSplits = CreateTableRequest_Split_FromProto(ctx, in.InitialSplits)
//		return out
//	}
//
//	func CreateTableRequest_ToProto(ctx *MapContext, in *krm.CreateTableRequest) *pb.CreateTableRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.CreateTableRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.TableId = LazyPtr(in.TableId)
//		out.Table = Table_ToProto(ctx, in.Table)
//		out.InitialSplits = CreateTableRequest_Split_ToProto(ctx, in.InitialSplits)
//		return out
//	}
//
//	func CreateTableRequest_Split_FromProto(ctx *MapContext, in *pb.CreateTableRequest_Split) *krm.CreateTableRequest_Split {
//		if in == nil {
//			return nil
//		}
//		out := &krm.CreateTableRequest_Split{}
//		out.Key = LazyPtr(in.Key)
//		return out
//	}
//
//	func CreateTableRequest_Split_ToProto(ctx *MapContext, in *krm.CreateTableRequest_Split) *pb.CreateTableRequest_Split {
//		if in == nil {
//			return nil
//		}
//		out := &pb.CreateTableRequest_Split{}
//		out.Key = LazyPtr(in.Key)
//		return out
//	}
//
//	func CreateTableFromSnapshotRequest_FromProto(ctx *MapContext, in *pb.CreateTableFromSnapshotRequest) *krm.CreateTableFromSnapshotRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.CreateTableFromSnapshotRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.TableId = LazyPtr(in.TableId)
//		out.SourceSnapshot = LazyPtr(in.SourceSnapshot)
//		return out
//	}
//
//	func CreateTableFromSnapshotRequest_ToProto(ctx *MapContext, in *krm.CreateTableFromSnapshotRequest) *pb.CreateTableFromSnapshotRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.CreateTableFromSnapshotRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.TableId = LazyPtr(in.TableId)
//		out.SourceSnapshot = LazyPtr(in.SourceSnapshot)
//		return out
//	}
//
//	func DropRowRangeRequest_FromProto(ctx *MapContext, in *pb.DropRowRangeRequest) *krm.DropRowRangeRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.DropRowRangeRequest{}
//		out.Name = LazyPtr(in.Name)
//		out.RowKeyPrefix = LazyPtr(in.RowKeyPrefix)
//		out.DeleteAllDataFromTable = LazyPtr(in.DeleteAllDataFromTable)
//		return out
//	}
//
//	func DropRowRangeRequest_ToProto(ctx *MapContext, in *krm.DropRowRangeRequest) *pb.DropRowRangeRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.DropRowRangeRequest{}
//		out.Name = LazyPtr(in.Name)
//		out.RowKeyPrefix = LazyPtr(in.RowKeyPrefix)
//		out.DeleteAllDataFromTable = LazyPtr(in.DeleteAllDataFromTable)
//		return out
//	}
//
//	func ListTablesRequest_FromProto(ctx *MapContext, in *pb.ListTablesRequest) *krm.ListTablesRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.ListTablesRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.View = Enum_FromProto(ctx, &in.View)
//		out.PageSize = LazyPtr(in.PageSize)
//		out.PageToken = LazyPtr(in.PageToken)
//		return out
//	}
//
//	func ListTablesRequest_ToProto(ctx *MapContext, in *krm.ListTablesRequest) *pb.ListTablesRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.ListTablesRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.View = Enum_ToProto(ctx, &in.View)
//		out.PageSize = LazyPtr(in.PageSize)
//		out.PageToken = LazyPtr(in.PageToken)
//		return out
//	}
//
//	func ListTablesResponse_FromProto(ctx *MapContext, in *pb.ListTablesResponse) *krm.ListTablesResponse {
//		if in == nil {
//			return nil
//		}
//		out := &krm.ListTablesResponse{}
//		out.Tables = Table_FromProto(ctx, in.Tables)
//		out.NextPageToken = LazyPtr(in.NextPageToken)
//		return out
//	}
//
//	func ListTablesResponse_ToProto(ctx *MapContext, in *krm.ListTablesResponse) *pb.ListTablesResponse {
//		if in == nil {
//			return nil
//		}
//		out := &pb.ListTablesResponse{}
//		out.Tables = Table_ToProto(ctx, in.Tables)
//		out.NextPageToken = LazyPtr(in.NextPageToken)
//		return out
//	}
//
//	func GetTableRequest_FromProto(ctx *MapContext, in *pb.GetTableRequest) *krm.GetTableRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.GetTableRequest{}
//		out.Name = LazyPtr(in.Name)
//		out.View = Enum_FromProto(ctx, &in.View)
//		return out
//	}
//
//	func GetTableRequest_ToProto(ctx *MapContext, in *krm.GetTableRequest) *pb.GetTableRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.GetTableRequest{}
//		out.Name = LazyPtr(in.Name)
//		out.View = Enum_ToProto(ctx, &in.View)
//		return out
//	}
//
//	func UpdateTableRequest_FromProto(ctx *MapContext, in *pb.UpdateTableRequest) *krm.UpdateTableRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.UpdateTableRequest{}
//		out.Table = Table_FromProto(ctx, in.Table)
//		out.UpdateMask = FieldMask_FromProto(ctx, in.UpdateMask)
//		return out
//	}
//
//	func UpdateTableRequest_ToProto(ctx *MapContext, in *krm.UpdateTableRequest) *pb.UpdateTableRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.UpdateTableRequest{}
//		out.Table = Table_ToProto(ctx, in.Table)
//		out.UpdateMask = FieldMask_ToProto(ctx, in.UpdateMask)
//		return out
//	}
//
//	func UpdateTableMetadata_FromProto(ctx *MapContext, in *pb.UpdateTableMetadata) *krm.UpdateTableMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &krm.UpdateTableMetadata{}
//		out.Name = LazyPtr(in.Name)
//		out.StartTime = Timestamp_FromProto(ctx, in.StartTime)
//		out.EndTime = Timestamp_FromProto(ctx, in.EndTime)
//		return out
//	}
//
//	func UpdateTableMetadata_ToProto(ctx *MapContext, in *krm.UpdateTableMetadata) *pb.UpdateTableMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &pb.UpdateTableMetadata{}
//		out.Name = LazyPtr(in.Name)
//		out.StartTime = Timestamp_ToProto(ctx, in.StartTime)
//		out.EndTime = Timestamp_ToProto(ctx, in.EndTime)
//		return out
//	}
//
//	func DeleteTableRequest_FromProto(ctx *MapContext, in *pb.DeleteTableRequest) *krm.DeleteTableRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.DeleteTableRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func DeleteTableRequest_ToProto(ctx *MapContext, in *krm.DeleteTableRequest) *pb.DeleteTableRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.DeleteTableRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func UndeleteTableRequest_FromProto(ctx *MapContext, in *pb.UndeleteTableRequest) *krm.UndeleteTableRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.UndeleteTableRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func UndeleteTableRequest_ToProto(ctx *MapContext, in *krm.UndeleteTableRequest) *pb.UndeleteTableRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.UndeleteTableRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func UndeleteTableMetadata_FromProto(ctx *MapContext, in *pb.UndeleteTableMetadata) *krm.UndeleteTableMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &krm.UndeleteTableMetadata{}
//		out.Name = LazyPtr(in.Name)
//		out.StartTime = Timestamp_FromProto(ctx, in.StartTime)
//		out.EndTime = Timestamp_FromProto(ctx, in.EndTime)
//		return out
//	}
//
//	func UndeleteTableMetadata_ToProto(ctx *MapContext, in *krm.UndeleteTableMetadata) *pb.UndeleteTableMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &pb.UndeleteTableMetadata{}
//		out.Name = LazyPtr(in.Name)
//		out.StartTime = Timestamp_ToProto(ctx, in.StartTime)
//		out.EndTime = Timestamp_ToProto(ctx, in.EndTime)
//		return out
//	}
//
//	func ModifyColumnFamiliesRequest_FromProto(ctx *MapContext, in *pb.ModifyColumnFamiliesRequest) *krm.ModifyColumnFamiliesRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.ModifyColumnFamiliesRequest{}
//		out.Name = LazyPtr(in.Name)
//		out.Modifications = ModifyColumnFamiliesRequest_Modification_FromProto(ctx, in.Modifications)
//		out.IgnoreWarnings = LazyPtr(in.IgnoreWarnings)
//		return out
//	}
//
//	func ModifyColumnFamiliesRequest_ToProto(ctx *MapContext, in *krm.ModifyColumnFamiliesRequest) *pb.ModifyColumnFamiliesRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.ModifyColumnFamiliesRequest{}
//		out.Name = LazyPtr(in.Name)
//		out.Modifications = ModifyColumnFamiliesRequest_Modification_ToProto(ctx, in.Modifications)
//		out.IgnoreWarnings = LazyPtr(in.IgnoreWarnings)
//		return out
//	}
//
//	func ModifyColumnFamiliesRequest_Modification_FromProto(ctx *MapContext, in *pb.ModifyColumnFamiliesRequest_Modification) *krm.ModifyColumnFamiliesRequest_Modification {
//		if in == nil {
//			return nil
//		}
//		out := &krm.ModifyColumnFamiliesRequest_Modification{}
//		out.Id = LazyPtr(in.Id)
//		out.Create = ColumnFamily_FromProto(ctx, in.Create)
//		out.Update = ColumnFamily_FromProto(ctx, in.Update)
//		out.Drop = LazyPtr(in.Drop)
//		out.UpdateMask = FieldMask_FromProto(ctx, in.UpdateMask)
//		return out
//	}
//
//	func ModifyColumnFamiliesRequest_Modification_ToProto(ctx *MapContext, in *krm.ModifyColumnFamiliesRequest_Modification) *pb.ModifyColumnFamiliesRequest_Modification {
//		if in == nil {
//			return nil
//		}
//		out := &pb.ModifyColumnFamiliesRequest_Modification{}
//		out.Id = LazyPtr(in.Id)
//		out.Create = ColumnFamily_ToProto(ctx, in.Create)
//		out.Update = ColumnFamily_ToProto(ctx, in.Update)
//		out.Drop = LazyPtr(in.Drop)
//		out.UpdateMask = FieldMask_ToProto(ctx, in.UpdateMask)
//		return out
//	}
//
//	func GenerateConsistencyTokenRequest_FromProto(ctx *MapContext, in *pb.GenerateConsistencyTokenRequest) *krm.GenerateConsistencyTokenRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.GenerateConsistencyTokenRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func GenerateConsistencyTokenRequest_ToProto(ctx *MapContext, in *krm.GenerateConsistencyTokenRequest) *pb.GenerateConsistencyTokenRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.GenerateConsistencyTokenRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func GenerateConsistencyTokenResponse_FromProto(ctx *MapContext, in *pb.GenerateConsistencyTokenResponse) *krm.GenerateConsistencyTokenResponse {
//		if in == nil {
//			return nil
//		}
//		out := &krm.GenerateConsistencyTokenResponse{}
//		out.ConsistencyToken = LazyPtr(in.ConsistencyToken)
//		return out
//	}
//
//	func GenerateConsistencyTokenResponse_ToProto(ctx *MapContext, in *krm.GenerateConsistencyTokenResponse) *pb.GenerateConsistencyTokenResponse {
//		if in == nil {
//			return nil
//		}
//		out := &pb.GenerateConsistencyTokenResponse{}
//		out.ConsistencyToken = LazyPtr(in.ConsistencyToken)
//		return out
//	}
//
//	func CheckConsistencyRequest_FromProto(ctx *MapContext, in *pb.CheckConsistencyRequest) *krm.CheckConsistencyRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.CheckConsistencyRequest{}
//		out.Name = LazyPtr(in.Name)
//		out.ConsistencyToken = LazyPtr(in.ConsistencyToken)
//		out.StandardReadRemoteWrites = StandardReadRemoteWrites_FromProto(ctx, in.StandardReadRemoteWrites)
//		out.DataBoostReadLocalWrites = DataBoostReadLocalWrites_FromProto(ctx, in.DataBoostReadLocalWrites)
//		return out
//	}
//
//	func CheckConsistencyRequest_ToProto(ctx *MapContext, in *krm.CheckConsistencyRequest) *pb.CheckConsistencyRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.CheckConsistencyRequest{}
//		out.Name = LazyPtr(in.Name)
//		out.ConsistencyToken = LazyPtr(in.ConsistencyToken)
//		out.StandardReadRemoteWrites = StandardReadRemoteWrites_ToProto(ctx, in.StandardReadRemoteWrites)
//		out.DataBoostReadLocalWrites = DataBoostReadLocalWrites_ToProto(ctx, in.DataBoostReadLocalWrites)
//		return out
//	}
//
//	func StandardReadRemoteWrites_FromProto(ctx *MapContext, in *pb.StandardReadRemoteWrites) *krm.StandardReadRemoteWrites {
//		if in == nil {
//			return nil
//		}
//		out := &krm.StandardReadRemoteWrites{}
//		return out
//	}
//
//	func StandardReadRemoteWrites_ToProto(ctx *MapContext, in *krm.StandardReadRemoteWrites) *pb.StandardReadRemoteWrites {
//		if in == nil {
//			return nil
//		}
//		out := &pb.StandardReadRemoteWrites{}
//		return out
//	}
//
//	func DataBoostReadLocalWrites_FromProto(ctx *MapContext, in *pb.DataBoostReadLocalWrites) *krm.DataBoostReadLocalWrites {
//		if in == nil {
//			return nil
//		}
//		out := &krm.DataBoostReadLocalWrites{}
//		return out
//	}
//
//	func DataBoostReadLocalWrites_ToProto(ctx *MapContext, in *krm.DataBoostReadLocalWrites) *pb.DataBoostReadLocalWrites {
//		if in == nil {
//			return nil
//		}
//		out := &pb.DataBoostReadLocalWrites{}
//		return out
//	}
//
//	func CheckConsistencyResponse_FromProto(ctx *MapContext, in *pb.CheckConsistencyResponse) *krm.CheckConsistencyResponse {
//		if in == nil {
//			return nil
//		}
//		out := &krm.CheckConsistencyResponse{}
//		out.Consistent = LazyPtr(in.Consistent)
//		return out
//	}
//
//	func CheckConsistencyResponse_ToProto(ctx *MapContext, in *krm.CheckConsistencyResponse) *pb.CheckConsistencyResponse {
//		if in == nil {
//			return nil
//		}
//		out := &pb.CheckConsistencyResponse{}
//		out.Consistent = LazyPtr(in.Consistent)
//		return out
//	}
//
//	func SnapshotTableRequest_FromProto(ctx *MapContext, in *pb.SnapshotTableRequest) *krm.SnapshotTableRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.SnapshotTableRequest{}
//		out.Name = LazyPtr(in.Name)
//		out.Cluster = LazyPtr(in.Cluster)
//		out.SnapshotId = LazyPtr(in.SnapshotId)
//		out.Ttl = Duration_FromProto(ctx, in.Ttl)
//		out.Description = LazyPtr(in.Description)
//		return out
//	}
//
//	func SnapshotTableRequest_ToProto(ctx *MapContext, in *krm.SnapshotTableRequest) *pb.SnapshotTableRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.SnapshotTableRequest{}
//		out.Name = LazyPtr(in.Name)
//		out.Cluster = LazyPtr(in.Cluster)
//		out.SnapshotId = LazyPtr(in.SnapshotId)
//		out.Ttl = Duration_ToProto(ctx, in.Ttl)
//		out.Description = LazyPtr(in.Description)
//		return out
//	}
//
//	func GetSnapshotRequest_FromProto(ctx *MapContext, in *pb.GetSnapshotRequest) *krm.GetSnapshotRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.GetSnapshotRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func GetSnapshotRequest_ToProto(ctx *MapContext, in *krm.GetSnapshotRequest) *pb.GetSnapshotRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.GetSnapshotRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func ListSnapshotsRequest_FromProto(ctx *MapContext, in *pb.ListSnapshotsRequest) *krm.ListSnapshotsRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.ListSnapshotsRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.PageSize = LazyPtr(in.PageSize)
//		out.PageToken = LazyPtr(in.PageToken)
//		return out
//	}
//
//	func ListSnapshotsRequest_ToProto(ctx *MapContext, in *krm.ListSnapshotsRequest) *pb.ListSnapshotsRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.ListSnapshotsRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.PageSize = LazyPtr(in.PageSize)
//		out.PageToken = LazyPtr(in.PageToken)
//		return out
//	}
//
//	func ListSnapshotsResponse_FromProto(ctx *MapContext, in *pb.ListSnapshotsResponse) *krm.ListSnapshotsResponse {
//		if in == nil {
//			return nil
//		}
//		out := &krm.ListSnapshotsResponse{}
//		out.Snapshots = Snapshot_FromProto(ctx, in.Snapshots)
//		out.NextPageToken = LazyPtr(in.NextPageToken)
//		return out
//	}
//
//	func ListSnapshotsResponse_ToProto(ctx *MapContext, in *krm.ListSnapshotsResponse) *pb.ListSnapshotsResponse {
//		if in == nil {
//			return nil
//		}
//		out := &pb.ListSnapshotsResponse{}
//		out.Snapshots = Snapshot_ToProto(ctx, in.Snapshots)
//		out.NextPageToken = LazyPtr(in.NextPageToken)
//		return out
//	}
//
//	func DeleteSnapshotRequest_FromProto(ctx *MapContext, in *pb.DeleteSnapshotRequest) *krm.DeleteSnapshotRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.DeleteSnapshotRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func DeleteSnapshotRequest_ToProto(ctx *MapContext, in *krm.DeleteSnapshotRequest) *pb.DeleteSnapshotRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.DeleteSnapshotRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func SnapshotTableMetadata_FromProto(ctx *MapContext, in *pb.SnapshotTableMetadata) *krm.SnapshotTableMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &krm.SnapshotTableMetadata{}
//		out.OriginalRequest = SnapshotTableRequest_FromProto(ctx, in.OriginalRequest)
//		out.RequestTime = Timestamp_FromProto(ctx, in.RequestTime)
//		out.FinishTime = Timestamp_FromProto(ctx, in.FinishTime)
//		return out
//	}
//
//	func SnapshotTableMetadata_ToProto(ctx *MapContext, in *krm.SnapshotTableMetadata) *pb.SnapshotTableMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &pb.SnapshotTableMetadata{}
//		out.OriginalRequest = SnapshotTableRequest_ToProto(ctx, in.OriginalRequest)
//		out.RequestTime = Timestamp_ToProto(ctx, in.RequestTime)
//		out.FinishTime = Timestamp_ToProto(ctx, in.FinishTime)
//		return out
//	}
//
//	func CreateTableFromSnapshotMetadata_FromProto(ctx *MapContext, in *pb.CreateTableFromSnapshotMetadata) *krm.CreateTableFromSnapshotMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &krm.CreateTableFromSnapshotMetadata{}
//		out.OriginalRequest = CreateTableFromSnapshotRequest_FromProto(ctx, in.OriginalRequest)
//		out.RequestTime = Timestamp_FromProto(ctx, in.RequestTime)
//		out.FinishTime = Timestamp_FromProto(ctx, in.FinishTime)
//		return out
//	}
//
//	func CreateTableFromSnapshotMetadata_ToProto(ctx *MapContext, in *krm.CreateTableFromSnapshotMetadata) *pb.CreateTableFromSnapshotMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &pb.CreateTableFromSnapshotMetadata{}
//		out.OriginalRequest = CreateTableFromSnapshotRequest_ToProto(ctx, in.OriginalRequest)
//		out.RequestTime = Timestamp_ToProto(ctx, in.RequestTime)
//		out.FinishTime = Timestamp_ToProto(ctx, in.FinishTime)
//		return out
//	}
//
//	func CreateBackupRequest_FromProto(ctx *MapContext, in *pb.CreateBackupRequest) *krm.CreateBackupRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.CreateBackupRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.BackupId = LazyPtr(in.BackupId)
//		out.Backup = Backup_FromProto(ctx, in.Backup)
//		return out
//	}
//
//	func CreateBackupRequest_ToProto(ctx *MapContext, in *krm.CreateBackupRequest) *pb.CreateBackupRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.CreateBackupRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.BackupId = LazyPtr(in.BackupId)
//		out.Backup = Backup_ToProto(ctx, in.Backup)
//		return out
//	}
//
//	func CreateBackupMetadata_FromProto(ctx *MapContext, in *pb.CreateBackupMetadata) *krm.CreateBackupMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &krm.CreateBackupMetadata{}
//		out.Name = LazyPtr(in.Name)
//		out.SourceTable = LazyPtr(in.SourceTable)
//		out.StartTime = Timestamp_FromProto(ctx, in.StartTime)
//		out.EndTime = Timestamp_FromProto(ctx, in.EndTime)
//		return out
//	}
//
//	func CreateBackupMetadata_ToProto(ctx *MapContext, in *krm.CreateBackupMetadata) *pb.CreateBackupMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &pb.CreateBackupMetadata{}
//		out.Name = LazyPtr(in.Name)
//		out.SourceTable = LazyPtr(in.SourceTable)
//		out.StartTime = Timestamp_ToProto(ctx, in.StartTime)
//		out.EndTime = Timestamp_ToProto(ctx, in.EndTime)
//		return out
//	}
//
//	func UpdateBackupRequest_FromProto(ctx *MapContext, in *pb.UpdateBackupRequest) *krm.UpdateBackupRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.UpdateBackupRequest{}
//		out.Backup = Backup_FromProto(ctx, in.Backup)
//		out.UpdateMask = FieldMask_FromProto(ctx, in.UpdateMask)
//		return out
//	}
//
//	func UpdateBackupRequest_ToProto(ctx *MapContext, in *krm.UpdateBackupRequest) *pb.UpdateBackupRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.UpdateBackupRequest{}
//		out.Backup = Backup_ToProto(ctx, in.Backup)
//		out.UpdateMask = FieldMask_ToProto(ctx, in.UpdateMask)
//		return out
//	}
//
//	func GetBackupRequest_FromProto(ctx *MapContext, in *pb.GetBackupRequest) *krm.GetBackupRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.GetBackupRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func GetBackupRequest_ToProto(ctx *MapContext, in *krm.GetBackupRequest) *pb.GetBackupRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.GetBackupRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func DeleteBackupRequest_FromProto(ctx *MapContext, in *pb.DeleteBackupRequest) *krm.DeleteBackupRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.DeleteBackupRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func DeleteBackupRequest_ToProto(ctx *MapContext, in *krm.DeleteBackupRequest) *pb.DeleteBackupRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.DeleteBackupRequest{}
//		out.Name = LazyPtr(in.Name)
//		return out
//	}
//
//	func ListBackupsRequest_FromProto(ctx *MapContext, in *pb.ListBackupsRequest) *krm.ListBackupsRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.ListBackupsRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.Filter = LazyPtr(in.Filter)
//		out.OrderBy = LazyPtr(in.OrderBy)
//		out.PageSize = LazyPtr(in.PageSize)
//		out.PageToken = LazyPtr(in.PageToken)
//		return out
//	}
//
//	func ListBackupsRequest_ToProto(ctx *MapContext, in *krm.ListBackupsRequest) *pb.ListBackupsRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.ListBackupsRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.Filter = LazyPtr(in.Filter)
//		out.OrderBy = LazyPtr(in.OrderBy)
//		out.PageSize = LazyPtr(in.PageSize)
//		out.PageToken = LazyPtr(in.PageToken)
//		return out
//	}
//
//	func ListBackupsResponse_FromProto(ctx *MapContext, in *pb.ListBackupsResponse) *krm.ListBackupsResponse {
//		if in == nil {
//			return nil
//		}
//		out := &krm.ListBackupsResponse{}
//		out.Backups = Backup_FromProto(ctx, in.Backups)
//		out.NextPageToken = LazyPtr(in.NextPageToken)
//		return out
//	}
//
//	func ListBackupsResponse_ToProto(ctx *MapContext, in *krm.ListBackupsResponse) *pb.ListBackupsResponse {
//		if in == nil {
//			return nil
//		}
//		out := &pb.ListBackupsResponse{}
//		out.Backups = Backup_ToProto(ctx, in.Backups)
//		out.NextPageToken = LazyPtr(in.NextPageToken)
//		return out
//	}
//
//	func CopyBackupRequest_FromProto(ctx *MapContext, in *pb.CopyBackupRequest) *krm.CopyBackupRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.CopyBackupRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.BackupId = LazyPtr(in.BackupId)
//		out.SourceBackup = LazyPtr(in.SourceBackup)
//		out.ExpireTime = Timestamp_FromProto(ctx, in.ExpireTime)
//		return out
//	}
//
//	func CopyBackupRequest_ToProto(ctx *MapContext, in *krm.CopyBackupRequest) *pb.CopyBackupRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.CopyBackupRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.BackupId = LazyPtr(in.BackupId)
//		out.SourceBackup = LazyPtr(in.SourceBackup)
//		out.ExpireTime = Timestamp_ToProto(ctx, in.ExpireTime)
//		return out
//	}
//
//	func CopyBackupMetadata_FromProto(ctx *MapContext, in *pb.CopyBackupMetadata) *krm.CopyBackupMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &krm.CopyBackupMetadata{}
//		out.Name = LazyPtr(in.Name)
//		out.SourceBackupInfo = BackupInfo_FromProto(ctx, in.SourceBackupInfo)
//		out.Progress = OperationProgress_FromProto(ctx, in.Progress)
//		return out
//	}
//
//	func CopyBackupMetadata_ToProto(ctx *MapContext, in *krm.CopyBackupMetadata) *pb.CopyBackupMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &pb.CopyBackupMetadata{}
//		out.Name = LazyPtr(in.Name)
//		out.SourceBackupInfo = BackupInfo_ToProto(ctx, in.SourceBackupInfo)
//		out.Progress = OperationProgress_ToProto(ctx, in.Progress)
//		return out
//	}
//
//	func CreateAuthorizedViewRequest_FromProto(ctx *MapContext, in *pb.CreateAuthorizedViewRequest) *krm.CreateAuthorizedViewRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.CreateAuthorizedViewRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.AuthorizedViewId = LazyPtr(in.AuthorizedViewId)
//		out.AuthorizedView = AuthorizedView_FromProto(ctx, in.AuthorizedView)
//		return out
//	}
//
//	func CreateAuthorizedViewRequest_ToProto(ctx *MapContext, in *krm.CreateAuthorizedViewRequest) *pb.CreateAuthorizedViewRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.CreateAuthorizedViewRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.AuthorizedViewId = LazyPtr(in.AuthorizedViewId)
//		out.AuthorizedView = AuthorizedView_ToProto(ctx, in.AuthorizedView)
//		return out
//	}
//
//	func CreateAuthorizedViewMetadata_FromProto(ctx *MapContext, in *pb.CreateAuthorizedViewMetadata) *krm.CreateAuthorizedViewMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &krm.CreateAuthorizedViewMetadata{}
//		out.OriginalRequest = CreateAuthorizedViewRequest_FromProto(ctx, in.OriginalRequest)
//		out.RequestTime = Timestamp_FromProto(ctx, in.RequestTime)
//		out.FinishTime = Timestamp_FromProto(ctx, in.FinishTime)
//		return out
//	}
//
//	func CreateAuthorizedViewMetadata_ToProto(ctx *MapContext, in *krm.CreateAuthorizedViewMetadata) *pb.CreateAuthorizedViewMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &pb.CreateAuthorizedViewMetadata{}
//		out.OriginalRequest = CreateAuthorizedViewRequest_ToProto(ctx, in.OriginalRequest)
//		out.RequestTime = Timestamp_ToProto(ctx, in.RequestTime)
//		out.FinishTime = Timestamp_ToProto(ctx, in.FinishTime)
//		return out
//	}
//
//	func ListAuthorizedViewsRequest_FromProto(ctx *MapContext, in *pb.ListAuthorizedViewsRequest) *krm.ListAuthorizedViewsRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.ListAuthorizedViewsRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.PageSize = LazyPtr(in.PageSize)
//		out.PageToken = LazyPtr(in.PageToken)
//		out.View = Enum_FromProto(ctx, &in.View)
//		return out
//	}
//
//	func ListAuthorizedViewsRequest_ToProto(ctx *MapContext, in *krm.ListAuthorizedViewsRequest) *pb.ListAuthorizedViewsRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.ListAuthorizedViewsRequest{}
//		out.Parent = LazyPtr(in.Parent)
//		out.PageSize = LazyPtr(in.PageSize)
//		out.PageToken = LazyPtr(in.PageToken)
//		out.View = Enum_ToProto(ctx, &in.View)
//		return out
//	}
//
//	func ListAuthorizedViewsResponse_FromProto(ctx *MapContext, in *pb.ListAuthorizedViewsResponse) *krm.ListAuthorizedViewsResponse {
//		if in == nil {
//			return nil
//		}
//		out := &krm.ListAuthorizedViewsResponse{}
//		out.AuthorizedViews = AuthorizedView_FromProto(ctx, in.AuthorizedViews)
//		out.NextPageToken = LazyPtr(in.NextPageToken)
//		return out
//	}
//
//	func ListAuthorizedViewsResponse_ToProto(ctx *MapContext, in *krm.ListAuthorizedViewsResponse) *pb.ListAuthorizedViewsResponse {
//		if in == nil {
//			return nil
//		}
//		out := &pb.ListAuthorizedViewsResponse{}
//		out.AuthorizedViews = AuthorizedView_ToProto(ctx, in.AuthorizedViews)
//		out.NextPageToken = LazyPtr(in.NextPageToken)
//		return out
//	}
//
//	func GetAuthorizedViewRequest_FromProto(ctx *MapContext, in *pb.GetAuthorizedViewRequest) *krm.GetAuthorizedViewRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.GetAuthorizedViewRequest{}
//		out.Name = LazyPtr(in.Name)
//		out.View = Enum_FromProto(ctx, &in.View)
//		return out
//	}
//
//	func GetAuthorizedViewRequest_ToProto(ctx *MapContext, in *krm.GetAuthorizedViewRequest) *pb.GetAuthorizedViewRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.GetAuthorizedViewRequest{}
//		out.Name = LazyPtr(in.Name)
//		out.View = Enum_ToProto(ctx, &in.View)
//		return out
//	}
//
//	func UpdateAuthorizedViewRequest_FromProto(ctx *MapContext, in *pb.UpdateAuthorizedViewRequest) *krm.UpdateAuthorizedViewRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.UpdateAuthorizedViewRequest{}
//		out.AuthorizedView = AuthorizedView_FromProto(ctx, in.AuthorizedView)
//		out.UpdateMask = FieldMask_FromProto(ctx, in.UpdateMask)
//		out.IgnoreWarnings = LazyPtr(in.IgnoreWarnings)
//		return out
//	}
//
//	func UpdateAuthorizedViewRequest_ToProto(ctx *MapContext, in *krm.UpdateAuthorizedViewRequest) *pb.UpdateAuthorizedViewRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.UpdateAuthorizedViewRequest{}
//		out.AuthorizedView = AuthorizedView_ToProto(ctx, in.AuthorizedView)
//		out.UpdateMask = FieldMask_ToProto(ctx, in.UpdateMask)
//		out.IgnoreWarnings = LazyPtr(in.IgnoreWarnings)
//		return out
//	}
//
//	func UpdateAuthorizedViewMetadata_FromProto(ctx *MapContext, in *pb.UpdateAuthorizedViewMetadata) *krm.UpdateAuthorizedViewMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &krm.UpdateAuthorizedViewMetadata{}
//		out.OriginalRequest = UpdateAuthorizedViewRequest_FromProto(ctx, in.OriginalRequest)
//		out.RequestTime = Timestamp_FromProto(ctx, in.RequestTime)
//		out.FinishTime = Timestamp_FromProto(ctx, in.FinishTime)
//		return out
//	}
//
//	func UpdateAuthorizedViewMetadata_ToProto(ctx *MapContext, in *krm.UpdateAuthorizedViewMetadata) *pb.UpdateAuthorizedViewMetadata {
//		if in == nil {
//			return nil
//		}
//		out := &pb.UpdateAuthorizedViewMetadata{}
//		out.OriginalRequest = UpdateAuthorizedViewRequest_ToProto(ctx, in.OriginalRequest)
//		out.RequestTime = Timestamp_ToProto(ctx, in.RequestTime)
//		out.FinishTime = Timestamp_ToProto(ctx, in.FinishTime)
//		return out
//	}
//
//	func DeleteAuthorizedViewRequest_FromProto(ctx *MapContext, in *pb.DeleteAuthorizedViewRequest) *krm.DeleteAuthorizedViewRequest {
//		if in == nil {
//			return nil
//		}
//		out := &krm.DeleteAuthorizedViewRequest{}
//		out.Name = LazyPtr(in.Name)
//		out.Etag = LazyPtr(in.Etag)
//		return out
//	}
//
//	func DeleteAuthorizedViewRequest_ToProto(ctx *MapContext, in *krm.DeleteAuthorizedViewRequest) *pb.DeleteAuthorizedViewRequest {
//		if in == nil {
//			return nil
//		}
//		out := &pb.DeleteAuthorizedViewRequest{}
//		out.Name = LazyPtr(in.Name)
//		out.Etag = LazyPtr(in.Etag)
//		return out
//	}
//
//	func Type_FromProto(ctx *MapContext, in *pb.Type) *krm.Type {
//		if in == nil {
//			return nil
//		}
//		out := &krm.Type{}
//		out.BytesType = Type_Bytes_FromProto(ctx, in.BytesType)
//		out.Int64Type = Type_Int64_FromProto(ctx, in.Int64Type)
//		out.AggregateType = Type_Aggregate_FromProto(ctx, in.AggregateType)
//		return out
//	}
//
//	func Type_ToProto(ctx *MapContext, in *krm.Type) *pb.Type {
//		if in == nil {
//			return nil
//		}
//		out := &pb.Type{}
//		out.BytesType = Type_Bytes_ToProto(ctx, in.BytesType)
//		out.Int64Type = Type_Int64_ToProto(ctx, in.Int64Type)
//		out.AggregateType = Type_Aggregate_ToProto(ctx, in.AggregateType)
//		return out
//	}
//
//	func Type_Bytes_FromProto(ctx *MapContext, in *pb.Type_Bytes) *krm.Type_Bytes {
//		if in == nil {
//			return nil
//		}
//		out := &krm.Type_Bytes{}
//		out.Encoding = Type_Bytes_Encoding_FromProto(ctx, in.Encoding)
//		return out
//	}
//
//	func Type_Bytes_ToProto(ctx *MapContext, in *krm.Type_Bytes) *pb.Type_Bytes {
//		if in == nil {
//			return nil
//		}
//		out := &pb.Type_Bytes{}
//		out.Encoding = Type_Bytes_Encoding_ToProto(ctx, in.Encoding)
//		return out
//	}
//
//	func Type_Bytes_Encoding_FromProto(ctx *MapContext, in *pb.Type_Bytes_Encoding) *krm.Type_Bytes_Encoding {
//		if in == nil {
//			return nil
//		}
//		out := &krm.Type_Bytes_Encoding{}
//		out.Raw = Type_Bytes_Encoding_Raw_FromProto(ctx, in.Raw)
//		return out
//	}
//
//	func Type_Bytes_Encoding_ToProto(ctx *MapContext, in *krm.Type_Bytes_Encoding) *pb.Type_Bytes_Encoding {
//		if in == nil {
//			return nil
//		}
//		out := &pb.Type_Bytes_Encoding{}
//		out.Raw = Type_Bytes_Encoding_Raw_ToProto(ctx, in.Raw)
//		return out
//	}
//
//	func Type_Bytes_Encoding_Raw_FromProto(ctx *MapContext, in *pb.Type_Bytes_Encoding_Raw) *krm.Type_Bytes_Encoding_Raw {
//		if in == nil {
//			return nil
//		}
//		out := &krm.Type_Bytes_Encoding_Raw{}
//		return out
//	}
//
//	func Type_Bytes_Encoding_Raw_ToProto(ctx *MapContext, in *krm.Type_Bytes_Encoding_Raw) *pb.Type_Bytes_Encoding_Raw {
//		if in == nil {
//			return nil
//		}
//		out := &pb.Type_Bytes_Encoding_Raw{}
//		return out
//	}
//
//	func Type_Int64_FromProto(ctx *MapContext, in *pb.Type_Int64) *krm.Type_Int64 {
//		if in == nil {
//			return nil
//		}
//		out := &krm.Type_Int64{}
//		out.Encoding = Type_Int64_Encoding_FromProto(ctx, in.Encoding)
//		return out
//	}
//
//	func Type_Int64_ToProto(ctx *MapContext, in *krm.Type_Int64) *pb.Type_Int64 {
//		if in == nil {
//			return nil
//		}
//		out := &pb.Type_Int64{}
//		out.Encoding = Type_Int64_Encoding_ToProto(ctx, in.Encoding)
//		return out
//	}
//
//	func Type_Int64_Encoding_FromProto(ctx *MapContext, in *pb.Type_Int64_Encoding) *krm.Type_Int64_Encoding {
//		if in == nil {
//			return nil
//		}
//		out := &krm.Type_Int64_Encoding{}
//		out.BigEndianBytes = Type_Int64_Encoding_BigEndianBytes_FromProto(ctx, in.BigEndianBytes)
//		return out
//	}
//
//	func Type_Int64_Encoding_ToProto(ctx *MapContext, in *krm.Type_Int64_Encoding) *pb.Type_Int64_Encoding {
//		if in == nil {
//			return nil
//		}
//		out := &pb.Type_Int64_Encoding{}
//		out.BigEndianBytes = Type_Int64_Encoding_BigEndianBytes_ToProto(ctx, in.BigEndianBytes)
//		return out
//	}
//
//	func Type_Int64_Encoding_BigEndianBytes_FromProto(ctx *MapContext, in *pb.Type_Int64_Encoding_BigEndianBytes) *krm.Type_Int64_Encoding_BigEndianBytes {
//		if in == nil {
//			return nil
//		}
//		out := &krm.Type_Int64_Encoding_BigEndianBytes{}
//		out.BytesType = Type_Bytes_FromProto(ctx, in.BytesType)
//		return out
//	}
//
//	func Type_Int64_Encoding_BigEndianBytes_ToProto(ctx *MapContext, in *krm.Type_Int64_Encoding_BigEndianBytes) *pb.Type_Int64_Encoding_BigEndianBytes {
//		if in == nil {
//			return nil
//		}
//		out := &pb.Type_Int64_Encoding_BigEndianBytes{}
//		out.BytesType = Type_Bytes_ToProto(ctx, in.BytesType)
//		return out
//	}
//
//	func Type_Aggregate_FromProto(ctx *MapContext, in *pb.Type_Aggregate) *krm.Type_Aggregate {
//		if in == nil {
//			return nil
//		}
//		out := &krm.Type_Aggregate{}
//		out.InputType = Type_FromProto(ctx, in.InputType)
//		out.StateType = Type_FromProto(ctx, in.StateType)
//		out.Sum = Type_Aggregate_Sum_FromProto(ctx, in.Sum)
//		return out
//	}
//
//	func Type_Aggregate_ToProto(ctx *MapContext, in *krm.Type_Aggregate) *pb.Type_Aggregate {
//		if in == nil {
//			return nil
//		}
//		out := &pb.Type_Aggregate{}
//		out.InputType = Type_ToProto(ctx, in.InputType)
//		out.StateType = Type_ToProto(ctx, in.StateType)
//		out.Sum = Type_Aggregate_Sum_ToProto(ctx, in.Sum)
//		return out
//	}
//
//	func Type_Aggregate_Sum_FromProto(ctx *MapContext, in *pb.Type_Aggregate_Sum) *krm.Type_Aggregate_Sum {
//		if in == nil {
//			return nil
//		}
//		out := &krm.Type_Aggregate_Sum{}
//		return out
//	}
//
//	func Type_Aggregate_Sum_ToProto(ctx *MapContext, in *krm.Type_Aggregate_Sum) *pb.Type_Aggregate_Sum {
//		if in == nil {
//			return nil
//		}
//		out := &pb.Type_Aggregate_Sum{}
//		return out
//	}
//
//	func RestoreInfo_FromProto(ctx *MapContext, in *pb.RestoreInfo) *krm.RestoreInfo {
//		if in == nil {
//			return nil
//		}
//		out := &krm.RestoreInfo{}
//		out.SourceType = Enum_FromProto(ctx, &in.SourceType)
//		out.BackupInfo = BackupInfo_FromProto(ctx, in.BackupInfo)
//		return out
//	}
//
//	func RestoreInfo_ToProto(ctx *MapContext, in *krm.RestoreInfo) *pb.RestoreInfo {
//		if in == nil {
//			return nil
//		}
//		out := &pb.RestoreInfo{}
//		out.SourceType = Enum_ToProto(ctx, &in.SourceType)
//		out.BackupInfo = BackupInfo_ToProto(ctx, in.BackupInfo)
//		return out
//	}
//
//	func ChangeStreamConfig_FromProto(ctx *MapContext, in *pb.ChangeStreamConfig) *krm.ChangeStreamConfig {
//		if in == nil {
//			return nil
//		}
//		out := &krm.ChangeStreamConfig{}
//		out.RetentionPeriod = Duration_FromProto(ctx, in.RetentionPeriod)
//		return out
//	}
//
//	func ChangeStreamConfig_ToProto(ctx *MapContext, in *krm.ChangeStreamConfig) *pb.ChangeStreamConfig {
//		if in == nil {
//			return nil
//		}
//		out := &pb.ChangeStreamConfig{}
//		out.RetentionPeriod = Duration_ToProto(ctx, in.RetentionPeriod)
//		return out
//	}
//
//	func Table_FromProto(ctx *MapContext, in *pb.Table) *krm.Table {
//		if in == nil {
//			return nil
//		}
//		out := &krm.Table{}
//		out.Name = LazyPtr(in.Name)
//		out.ClusterStates = Table_ClusterStatesEntry_FromProto(ctx, in.ClusterStates)
//		out.ColumnFamilies = Table_ColumnFamiliesEntry_FromProto(ctx, in.ColumnFamilies)
//		out.Granularity = Enum_FromProto(ctx, &in.Granularity)
//		out.RestoreInfo = RestoreInfo_FromProto(ctx, in.RestoreInfo)
//		out.ChangeStreamConfig = ChangeStreamConfig_FromProto(ctx, in.ChangeStreamConfig)
//		out.DeletionProtection = LazyPtr(in.DeletionProtection)
//		out.AutomatedBackupPolicy = Table_AutomatedBackupPolicy_FromProto(ctx, in.AutomatedBackupPolicy)
//		return out
//	}
//
//	func Table_ToProto(ctx *MapContext, in *krm.Table) *pb.Table {
//		if in == nil {
//			return nil
//		}
//		out := &pb.Table{}
//		out.Name = LazyPtr(in.Name)
//		out.ClusterStates = Table_ClusterStatesEntry_ToProto(ctx, in.ClusterStates)
//		out.ColumnFamilies = Table_ColumnFamiliesEntry_ToProto(ctx, in.ColumnFamilies)
//		out.Granularity = Enum_ToProto(ctx, &in.Granularity)
//		out.RestoreInfo = RestoreInfo_ToProto(ctx, in.RestoreInfo)
//		out.ChangeStreamConfig = ChangeStreamConfig_ToProto(ctx, in.ChangeStreamConfig)
//		out.DeletionProtection = LazyPtr(in.DeletionProtection)
//		out.AutomatedBackupPolicy = Table_AutomatedBackupPolicy_ToProto(ctx, in.AutomatedBackupPolicy)
//		return out
//	}
//
//	func Table_ClusterState_FromProto(ctx *MapContext, in *pb.Table_ClusterState) *krm.Table_ClusterState {
//		if in == nil {
//			return nil
//		}
//		out := &krm.Table_ClusterState{}
//		out.ReplicationState = Enum_FromProto(ctx, &in.ReplicationState)
//		out.EncryptionInfo = EncryptionInfo_FromProto(ctx, in.EncryptionInfo)
//		return out
//	}
//
//	func Table_ClusterState_ToProto(ctx *MapContext, in *krm.Table_ClusterState) *pb.Table_ClusterState {
//		if in == nil {
//			return nil
//		}
//		out := &pb.Table_ClusterState{}
//		out.ReplicationState = Enum_ToProto(ctx, &in.ReplicationState)
//		out.EncryptionInfo = EncryptionInfo_ToProto(ctx, in.EncryptionInfo)
//		return out
//	}
//
//	func Table_AutomatedBackupPolicy_FromProto(ctx *MapContext, in *pb.Table_AutomatedBackupPolicy) *krm.Table_AutomatedBackupPolicy {
//		if in == nil {
//			return nil
//		}
//		out := &krm.Table_AutomatedBackupPolicy{}
//		out.RetentionPeriod = Duration_FromProto(ctx, in.RetentionPeriod)
//		out.Frequency = Duration_FromProto(ctx, in.Frequency)
//		return out
//	}
//
//	func Table_AutomatedBackupPolicy_ToProto(ctx *MapContext, in *krm.Table_AutomatedBackupPolicy) *pb.Table_AutomatedBackupPolicy {
//		if in == nil {
//			return nil
//		}
//		out := &pb.Table_AutomatedBackupPolicy{}
//		out.RetentionPeriod = Duration_ToProto(ctx, in.RetentionPeriod)
//		out.Frequency = Duration_ToProto(ctx, in.Frequency)
//		return out
//	}
//
//	func Table_ClusterStatesEntry_FromProto(ctx *MapContext, in *pb.Table_ClusterStatesEntry) *krm.Table_ClusterStatesEntry {
//		if in == nil {
//			return nil
//		}
//		out := &krm.Table_ClusterStatesEntry{}
//		out.Key = LazyPtr(in.Key)
//		out.Value = Table_ClusterState_FromProto(ctx, in.Value)
//		return out
//	}
//
//	func Table_ClusterStatesEntry_ToProto(ctx *MapContext, in *krm.Table_ClusterStatesEntry) *pb.Table_ClusterStatesEntry {
//		if in == nil {
//			return nil
//		}
//		out := &pb.Table_ClusterStatesEntry{}
//		out.Key = LazyPtr(in.Key)
//		out.Value = Table_ClusterState_ToProto(ctx, in.Value)
//		return out
//	}
//
//	func Table_ColumnFamiliesEntry_FromProto(ctx *MapContext, in *pb.Table_ColumnFamiliesEntry) *krm.Table_ColumnFamiliesEntry {
//		if in == nil {
//			return nil
//		}
//		out := &krm.Table_ColumnFamiliesEntry{}
//		out.Key = LazyPtr(in.Key)
//		out.Value = ColumnFamily_FromProto(ctx, in.Value)
//		return out
//	}
//
//	func Table_ColumnFamiliesEntry_ToProto(ctx *MapContext, in *krm.Table_ColumnFamiliesEntry) *pb.Table_ColumnFamiliesEntry {
//		if in == nil {
//			return nil
//		}
//		out := &pb.Table_ColumnFamiliesEntry{}
//		out.Key = LazyPtr(in.Key)
//		out.Value = ColumnFamily_ToProto(ctx, in.Value)
//		return out
//	}
//
//	func AuthorizedView_FromProto(ctx *MapContext, in *pb.AuthorizedView) *krm.AuthorizedView {
//		if in == nil {
//			return nil
//		}
//		out := &krm.AuthorizedView{}
//		out.Name = LazyPtr(in.Name)
//		out.SubsetView = AuthorizedView_SubsetView_FromProto(ctx, in.SubsetView)
//		out.Etag = LazyPtr(in.Etag)
//		out.DeletionProtection = LazyPtr(in.DeletionProtection)
//		return out
//	}
//
//	func AuthorizedView_ToProto(ctx *MapContext, in *krm.AuthorizedView) *pb.AuthorizedView {
//		if in == nil {
//			return nil
//		}
//		out := &pb.AuthorizedView{}
//		out.Name = LazyPtr(in.Name)
//		out.SubsetView = AuthorizedView_SubsetView_ToProto(ctx, in.SubsetView)
//		out.Etag = LazyPtr(in.Etag)
//		out.DeletionProtection = LazyPtr(in.DeletionProtection)
//		return out
//	}
//
//	func AuthorizedView_FamilySubsets_FromProto(ctx *MapContext, in *pb.AuthorizedView_FamilySubsets) *krm.AuthorizedView_FamilySubsets {
//		if in == nil {
//			return nil
//		}
//		out := &krm.AuthorizedView_FamilySubsets{}
//		out.Qualifiers = LazyPtr(in.Qualifiers)
//		out.QualifierPrefixes = LazyPtr(in.QualifierPrefixes)
//		return out
//	}
//
//	func AuthorizedView_FamilySubsets_ToProto(ctx *MapContext, in *krm.AuthorizedView_FamilySubsets) *pb.AuthorizedView_FamilySubsets {
//		if in == nil {
//			return nil
//		}
//		out := &pb.AuthorizedView_FamilySubsets{}
//		out.Qualifiers = LazyPtr(in.Qualifiers)
//		out.QualifierPrefixes = LazyPtr(in.QualifierPrefixes)
//		return out
//	}
//
//	func AuthorizedView_SubsetView_FromProto(ctx *MapContext, in *pb.AuthorizedView_SubsetView) *krm.AuthorizedView_SubsetView {
//		if in == nil {
//			return nil
//		}
//		out := &krm.AuthorizedView_SubsetView{}
//		out.RowPrefixes = LazyPtr(in.RowPrefixes)
//		out.FamilySubsets = AuthorizedView_SubsetView_FamilySubsetsEntry_FromProto(ctx, in.FamilySubsets)
//		return out
//	}
//
//	func AuthorizedView_SubsetView_ToProto(ctx *MapContext, in *krm.AuthorizedView_SubsetView) *pb.AuthorizedView_SubsetView {
//		if in == nil {
//			return nil
//		}
//		out := &pb.AuthorizedView_SubsetView{}
//		out.RowPrefixes = LazyPtr(in.RowPrefixes)
//		out.FamilySubsets = AuthorizedView_SubsetView_FamilySubsetsEntry_ToProto(ctx, in.FamilySubsets)
//		return out
//	}
//
//	func AuthorizedView_SubsetView_FamilySubsetsEntry_FromProto(ctx *MapContext, in *pb.AuthorizedView_SubsetView_FamilySubsetsEntry) *krm.AuthorizedView_SubsetView_FamilySubsetsEntry {
//		if in == nil {
//			return nil
//		}
//		out := &krm.AuthorizedView_SubsetView_FamilySubsetsEntry{}
//		out.Key = LazyPtr(in.Key)
//		out.Value = AuthorizedView_FamilySubsets_FromProto(ctx, in.Value)
//		return out
//	}
//
//	func AuthorizedView_SubsetView_FamilySubsetsEntry_ToProto(ctx *MapContext, in *krm.AuthorizedView_SubsetView_FamilySubsetsEntry) *pb.AuthorizedView_SubsetView_FamilySubsetsEntry {
//		if in == nil {
//			return nil
//		}
//		out := &pb.AuthorizedView_SubsetView_FamilySubsetsEntry{}
//		out.Key = LazyPtr(in.Key)
//		out.Value = AuthorizedView_FamilySubsets_ToProto(ctx, in.Value)
//		return out
//	}
//
//	func ColumnFamily_FromProto(ctx *MapContext, in *pb.ColumnFamily) *krm.ColumnFamily {
//		if in == nil {
//			return nil
//		}
//		out := &krm.ColumnFamily{}
//		out.GcRule = GcRule_FromProto(ctx, in.GcRule)
//		out.ValueType = Type_FromProto(ctx, in.ValueType)
//		return out
//	}
//
//	func ColumnFamily_ToProto(ctx *MapContext, in *krm.ColumnFamily) *pb.ColumnFamily {
//		if in == nil {
//			return nil
//		}
//		out := &pb.ColumnFamily{}
//		out.GcRule = GcRule_ToProto(ctx, in.GcRule)
//		out.ValueType = Type_ToProto(ctx, in.ValueType)
//		return out
//	}
//
//	func GcRule_FromProto(ctx *MapContext, in *pb.GcRule) *krm.GcRule {
//		if in == nil {
//			return nil
//		}
//		out := &krm.GcRule{}
//		out.MaxNumVersions = LazyPtr(in.MaxNumVersions)
//		out.MaxAge = Duration_FromProto(ctx, in.MaxAge)
//		out.Intersection = GcRule_Intersection_FromProto(ctx, in.Intersection)
//		out.Union = GcRule_Union_FromProto(ctx, in.Union)
//		return out
//	}
//
//	func GcRule_ToProto(ctx *MapContext, in *krm.GcRule) *pb.GcRule {
//		if in == nil {
//			return nil
//		}
//		out := &pb.GcRule{}
//		out.MaxNumVersions = LazyPtr(in.MaxNumVersions)
//		out.MaxAge = Duration_ToProto(ctx, in.MaxAge)
//		out.Intersection = GcRule_Intersection_ToProto(ctx, in.Intersection)
//		out.Union = GcRule_Union_ToProto(ctx, in.Union)
//		return out
//	}
//
//	func GcRule_Intersection_FromProto(ctx *MapContext, in *pb.GcRule_Intersection) *krm.GcRule_Intersection {
//		if in == nil {
//			return nil
//		}
//		out := &krm.GcRule_Intersection{}
//		out.Rules = GcRule_FromProto(ctx, in.Rules)
//		return out
//	}
//
//	func GcRule_Intersection_ToProto(ctx *MapContext, in *krm.GcRule_Intersection) *pb.GcRule_Intersection {
//		if in == nil {
//			return nil
//		}
//		out := &pb.GcRule_Intersection{}
//		out.Rules = GcRule_ToProto(ctx, in.Rules)
//		return out
//	}
//
//	func GcRule_Union_FromProto(ctx *MapContext, in *pb.GcRule_Union) *krm.GcRule_Union {
//		if in == nil {
//			return nil
//		}
//		out := &krm.GcRule_Union{}
//		out.Rules = GcRule_FromProto(ctx, in.Rules)
//		return out
//	}
//
//	func GcRule_Union_ToProto(ctx *MapContext, in *krm.GcRule_Union) *pb.GcRule_Union {
//		if in == nil {
//			return nil
//		}
//		out := &pb.GcRule_Union{}
//		out.Rules = GcRule_ToProto(ctx, in.Rules)
//		return out
//	}
//
//	func EncryptionInfo_FromProto(ctx *MapContext, in *pb.EncryptionInfo) *krm.EncryptionInfo {
//		if in == nil {
//			return nil
//		}
//		out := &krm.EncryptionInfo{}
//		out.EncryptionType = Enum_FromProto(ctx, &in.EncryptionType)
//		out.EncryptionStatus = Status_FromProto(ctx, in.EncryptionStatus)
//		out.KmsKeyVersion = LazyPtr(in.KmsKeyVersion)
//		return out
//	}
//
//	func EncryptionInfo_ToProto(ctx *MapContext, in *krm.EncryptionInfo) *pb.EncryptionInfo {
//		if in == nil {
//			return nil
//		}
//		out := &pb.EncryptionInfo{}
//		out.EncryptionType = Enum_ToProto(ctx, &in.EncryptionType)
//		out.EncryptionStatus = Status_ToProto(ctx, in.EncryptionStatus)
//		out.KmsKeyVersion = LazyPtr(in.KmsKeyVersion)
//		return out
//	}
//
//	func Snapshot_FromProto(ctx *MapContext, in *pb.Snapshot) *krm.Snapshot {
//		if in == nil {
//			return nil
//		}
//		out := &krm.Snapshot{}
//		out.Name = LazyPtr(in.Name)
//		out.SourceTable = Table_FromProto(ctx, in.SourceTable)
//		out.DataSizeBytes = LazyPtr(in.DataSizeBytes)
//		out.CreateTime = Timestamp_FromProto(ctx, in.CreateTime)
//		out.DeleteTime = Timestamp_FromProto(ctx, in.DeleteTime)
//		out.State = Enum_FromProto(ctx, &in.State)
//		out.Description = LazyPtr(in.Description)
//		return out
//	}
//
//	func Snapshot_ToProto(ctx *MapContext, in *krm.Snapshot) *pb.Snapshot {
//		if in == nil {
//			return nil
//		}
//		out := &pb.Snapshot{}
//		out.Name = LazyPtr(in.Name)
//		out.SourceTable = Table_ToProto(ctx, in.SourceTable)
//		out.DataSizeBytes = LazyPtr(in.DataSizeBytes)
//		out.CreateTime = Timestamp_ToProto(ctx, in.CreateTime)
//		out.DeleteTime = Timestamp_ToProto(ctx, in.DeleteTime)
//		out.State = Enum_ToProto(ctx, &in.State)
//		out.Description = LazyPtr(in.Description)
//		return out
//	}
//
//	func Backup_FromProto(ctx *MapContext, in *pb.Backup) *krm.Backup {
//		if in == nil {
//			return nil
//		}
//		out := &krm.Backup{}
//		out.Name = LazyPtr(in.Name)
//		out.SourceTable = LazyPtr(in.SourceTable)
//		out.SourceBackup = LazyPtr(in.SourceBackup)
//		out.ExpireTime = Timestamp_FromProto(ctx, in.ExpireTime)
//		out.StartTime = Timestamp_FromProto(ctx, in.StartTime)
//		out.EndTime = Timestamp_FromProto(ctx, in.EndTime)
//		out.SizeBytes = LazyPtr(in.SizeBytes)
//		out.State = Enum_FromProto(ctx, &in.State)
//		out.EncryptionInfo = EncryptionInfo_FromProto(ctx, in.EncryptionInfo)
//		return out
//	}
//
//	func Backup_ToProto(ctx *MapContext, in *krm.Backup) *pb.Backup {
//		if in == nil {
//			return nil
//		}
//		out := &pb.Backup{}
//		out.Name = LazyPtr(in.Name)
//		out.SourceTable = LazyPtr(in.SourceTable)
//		out.SourceBackup = LazyPtr(in.SourceBackup)
//		out.ExpireTime = Timestamp_ToProto(ctx, in.ExpireTime)
//		out.StartTime = Timestamp_ToProto(ctx, in.StartTime)
//		out.EndTime = Timestamp_ToProto(ctx, in.EndTime)
//		out.SizeBytes = LazyPtr(in.SizeBytes)
//		out.State = Enum_ToProto(ctx, &in.State)
//		out.EncryptionInfo = EncryptionInfo_ToProto(ctx, in.EncryptionInfo)
//		return out
//	}
//
//	func BackupInfo_FromProto(ctx *MapContext, in *pb.BackupInfo) *krm.BackupInfo {
//		if in == nil {
//			return nil
//		}
//		out := &krm.BackupInfo{}
//		out.Backup = LazyPtr(in.Backup)
//		out.StartTime = Timestamp_FromProto(ctx, in.StartTime)
//		out.EndTime = Timestamp_FromProto(ctx, in.EndTime)
//		out.SourceTable = LazyPtr(in.SourceTable)
//		out.SourceBackup = LazyPtr(in.SourceBackup)
//		return out
//	}
//
//	func BackupInfo_ToProto(ctx *MapContext, in *krm.BackupInfo) *pb.BackupInfo {
//		if in == nil {
//			return nil
//		}
//		out := &pb.BackupInfo{}
//		out.Backup = LazyPtr(in.Backup)
//		out.StartTime = Timestamp_ToProto(ctx, in.StartTime)
//		out.EndTime = Timestamp_ToProto(ctx, in.EndTime)
//		out.SourceTable = LazyPtr(in.SourceTable)
//		out.SourceBackup = LazyPtr(in.SourceBackup)
//		return out
//	}
//
//	func Instance_FromProto(ctx *MapContext, in *pb.Instance) *krm.Instance {
//		if in == nil {
//			return nil
//		}
//		out := &krm.Instance{}
//		out.Name = LazyPtr(in.Name)
//		out.DisplayName = LazyPtr(in.DisplayName)
//		out.State = Enum_FromProto(ctx, &in.State)
//		out.Type = Enum_FromProto(ctx, &in.Type)
//		out.Labels = Instance_LabelsEntry_FromProto(ctx, in.Labels)
//		out.CreateTime = Timestamp_FromProto(ctx, in.CreateTime)
//		out.SatisfiesPzs = LazyPtr(in.SatisfiesPzs)
//		return out
//	}
func Instance_ToProto(ctx *MapContext, in *krm.BigtableInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.DisplayName = ValueOf(in.DisplayName)
	out.Type = Enum_ToProto[pb.Instance_Type](ctx, in.InstanceType)
	// out.Labels = Instance_LabelsEntry_ToProto(ctx, in.Labels)
	return out
}

//	func Instance_LabelsEntry_FromProto(ctx *MapContext, in *pb.Instance_LabelsEntry) *krm.Instance_LabelsEntry {
//		if in == nil {
//			return nil
//		}
//		out := &krm.Instance_LabelsEntry{}
//		out.Key = LazyPtr(in.Key)
//		out.Value = LazyPtr(in.Value)
//		return out
//	}
//
//	func Instance_LabelsEntry_ToProto(ctx *MapContext, in *krm.Instance_LabelsEntry) *pb.Instance_LabelsEntry {
//		if in == nil {
//			return nil
//		}
//		out := &pb.Instance_LabelsEntry{}
//		out.Key = LazyPtr(in.Key)
//		out.Value = LazyPtr(in.Value)
//		return out
//	}
//
//	func AutoscalingTargets_FromProto(ctx *MapContext, in *pb.AutoscalingTargets) *krm.AutoscalingTargets {
//		if in == nil {
//			return nil
//		}
//		out := &krm.AutoscalingTargets{}
//		out.CpuUtilizationPercent = LazyPtr(in.CpuUtilizationPercent)
//		out.StorageUtilizationGibPerNode = LazyPtr(in.StorageUtilizationGibPerNode)
//		return out
//	}
//
//	func AutoscalingTargets_ToProto(ctx *MapContext, in *krm.AutoscalingTargets) *pb.AutoscalingTargets {
//		if in == nil {
//			return nil
//		}
//		out := &pb.AutoscalingTargets{}
//		out.CpuUtilizationPercent = LazyPtr(in.CpuUtilizationPercent)
//		out.StorageUtilizationGibPerNode = LazyPtr(in.StorageUtilizationGibPerNode)
//		return out
//	}
//
//	func AutoscalingLimits_FromProto(ctx *MapContext, in *pb.AutoscalingLimits) *krm.AutoscalingLimits {
//		if in == nil {
//			return nil
//		}
//		out := &krm.AutoscalingLimits{}
//		out.MinServeNodes = LazyPtr(in.MinServeNodes)
//		out.MaxServeNodes = LazyPtr(in.MaxServeNodes)
//		return out
//	}
//
//	func AutoscalingLimits_ToProto(ctx *MapContext, in *krm.AutoscalingLimits) *pb.AutoscalingLimits {
//		if in == nil {
//			return nil
//		}
//		out := &pb.AutoscalingLimits{}
//		out.MinServeNodes = LazyPtr(in.MinServeNodes)
//		out.MaxServeNodes = LazyPtr(in.MaxServeNodes)
//		return out
//	}
//
//	func Cluster_FromProto(ctx *MapContext, in *pb.Cluster) *krm.Cluster {
//		if in == nil {
//			return nil
//		}
//		out := &krm.Cluster{}
//		out.Name = LazyPtr(in.Name)
//		out.Location = LazyPtr(in.Location)
//		out.State = Enum_FromProto(ctx, &in.State)
//		out.ServeNodes = LazyPtr(in.ServeNodes)
//		out.ClusterConfig = Cluster_ClusterConfig_FromProto(ctx, in.ClusterConfig)
//		out.DefaultStorageType = Enum_FromProto(ctx, &in.DefaultStorageType)
//		out.EncryptionConfig = Cluster_EncryptionConfig_FromProto(ctx, in.EncryptionConfig)
//		return out
//	}

// func Cluster_ClusterAutoscalingConfig_FromProto(ctx *MapContext, in *pb.Cluster_ClusterAutoscalingConfig) *krm.Cluster_ClusterAutoscalingConfig {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.Cluster_ClusterAutoscalingConfig{}
// 	out.AutoscalingLimits = AutoscalingLimits_FromProto(ctx, in.AutoscalingLimits)
// 	out.AutoscalingTargets = AutoscalingTargets_FromProto(ctx, in.AutoscalingTargets)
// 	return out
// }
// func Cluster_ClusterAutoscalingConfig_ToProto(ctx *MapContext, in *krm.Cluster_ClusterAutoscalingConfig) *pb.Cluster_ClusterAutoscalingConfig {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.Cluster_ClusterAutoscalingConfig{}
// 	out.AutoscalingLimits = AutoscalingLimits_ToProto(ctx, in.AutoscalingLimits)
// 	out.AutoscalingTargets = AutoscalingTargets_ToProto(ctx, in.AutoscalingTargets)
// 	return out
// }
// func Cluster_ClusterConfig_FromProto(ctx *MapContext, in *pb.Cluster_ClusterConfig) *krm.Cluster_ClusterConfig {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.Cluster_ClusterConfig{}
// 	out.ClusterAutoscalingConfig = Cluster_ClusterAutoscalingConfig_FromProto(ctx, in.ClusterAutoscalingConfig)
// 	return out
// }
// func Cluster_ClusterConfig_ToProto(ctx *MapContext, in *krm.Cluster_ClusterConfig) *pb.Cluster_ClusterConfig {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.Cluster_ClusterConfig{}
// 	out.ClusterAutoscalingConfig = Cluster_ClusterAutoscalingConfig_ToProto(ctx, in.ClusterAutoscalingConfig)
// 	return out
// }
// func Cluster_EncryptionConfig_FromProto(ctx *MapContext, in *pb.Cluster_EncryptionConfig) *krm.Cluster_EncryptionConfig {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.Cluster_EncryptionConfig{}
// 	out.KmsKeyName = LazyPtr(in.KmsKeyName)
// 	return out
// }
// func Cluster_EncryptionConfig_ToProto(ctx *MapContext, in *krm.Cluster_EncryptionConfig) *pb.Cluster_EncryptionConfig {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.Cluster_EncryptionConfig{}
// 	out.KmsKeyName = LazyPtr(in.KmsKeyName)
// 	return out
// }
// func AppProfile_FromProto(ctx *MapContext, in *pb.AppProfile) *krm.AppProfile {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.AppProfile{}
// 	out.Name = LazyPtr(in.Name)
// 	out.Etag = LazyPtr(in.Etag)
// 	out.Description = LazyPtr(in.Description)
// 	out.MultiClusterRoutingUseAny = AppProfile_MultiClusterRoutingUseAny_FromProto(ctx, in.MultiClusterRoutingUseAny)
// 	out.SingleClusterRouting = AppProfile_SingleClusterRouting_FromProto(ctx, in.SingleClusterRouting)
// 	out.Priority = Enum_FromProto(ctx, &in.Priority)
// 	out.StandardIsolation = AppProfile_StandardIsolation_FromProto(ctx, in.StandardIsolation)
// 	out.DataBoostIsolationReadOnly = AppProfile_DataBoostIsolationReadOnly_FromProto(ctx, in.DataBoostIsolationReadOnly)
// 	return out
// }
// func AppProfile_ToProto(ctx *MapContext, in *krm.AppProfile) *pb.AppProfile {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.AppProfile{}
// 	out.Name = LazyPtr(in.Name)
// 	out.Etag = LazyPtr(in.Etag)
// 	out.Description = LazyPtr(in.Description)
// 	out.MultiClusterRoutingUseAny = AppProfile_MultiClusterRoutingUseAny_ToProto(ctx, in.MultiClusterRoutingUseAny)
// 	out.SingleClusterRouting = AppProfile_SingleClusterRouting_ToProto(ctx, in.SingleClusterRouting)
// 	out.Priority = Enum_ToProto(ctx, &in.Priority)
// 	out.StandardIsolation = AppProfile_StandardIsolation_ToProto(ctx, in.StandardIsolation)
// 	out.DataBoostIsolationReadOnly = AppProfile_DataBoostIsolationReadOnly_ToProto(ctx, in.DataBoostIsolationReadOnly)
// 	return out
// }
// func AppProfile_MultiClusterRoutingUseAny_FromProto(ctx *MapContext, in *pb.AppProfile_MultiClusterRoutingUseAny) *krm.AppProfile_MultiClusterRoutingUseAny {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.AppProfile_MultiClusterRoutingUseAny{}
// 	out.ClusterIds = LazyPtr(in.ClusterIds)
// 	return out
// }
// func AppProfile_MultiClusterRoutingUseAny_ToProto(ctx *MapContext, in *krm.AppProfile_MultiClusterRoutingUseAny) *pb.AppProfile_MultiClusterRoutingUseAny {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.AppProfile_MultiClusterRoutingUseAny{}
// 	out.ClusterIds = LazyPtr(in.ClusterIds)
// 	return out
// }
// func AppProfile_SingleClusterRouting_FromProto(ctx *MapContext, in *pb.AppProfile_SingleClusterRouting) *krm.AppProfile_SingleClusterRouting {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.AppProfile_SingleClusterRouting{}
// 	out.ClusterId = LazyPtr(in.ClusterId)
// 	out.AllowTransactionalWrites = LazyPtr(in.AllowTransactionalWrites)
// 	return out
// }
// func AppProfile_SingleClusterRouting_ToProto(ctx *MapContext, in *krm.AppProfile_SingleClusterRouting) *pb.AppProfile_SingleClusterRouting {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.AppProfile_SingleClusterRouting{}
// 	out.ClusterId = LazyPtr(in.ClusterId)
// 	out.AllowTransactionalWrites = LazyPtr(in.AllowTransactionalWrites)
// 	return out
// }
// func AppProfile_StandardIsolation_FromProto(ctx *MapContext, in *pb.AppProfile_StandardIsolation) *krm.AppProfile_StandardIsolation {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.AppProfile_StandardIsolation{}
// 	out.Priority = Enum_FromProto(ctx, &in.Priority)
// 	return out
// }
// func AppProfile_StandardIsolation_ToProto(ctx *MapContext, in *krm.AppProfile_StandardIsolation) *pb.AppProfile_StandardIsolation {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.AppProfile_StandardIsolation{}
// 	out.Priority = Enum_ToProto(ctx, &in.Priority)
// 	return out
// }
// func AppProfile_DataBoostIsolationReadOnly_FromProto(ctx *MapContext, in *pb.AppProfile_DataBoostIsolationReadOnly) *krm.AppProfile_DataBoostIsolationReadOnly {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.AppProfile_DataBoostIsolationReadOnly{}
// 	out.ComputeBillingOwner = Enum_FromProto(ctx, &in.ComputeBillingOwner)
// 	return out
// }
// func AppProfile_DataBoostIsolationReadOnly_ToProto(ctx *MapContext, in *krm.AppProfile_DataBoostIsolationReadOnly) *pb.AppProfile_DataBoostIsolationReadOnly {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.AppProfile_DataBoostIsolationReadOnly{}
// 	out.ComputeBillingOwner = Enum_ToProto(ctx, &in.ComputeBillingOwner)
// 	return out
// }
// func HotTablet_FromProto(ctx *MapContext, in *pb.HotTablet) *krm.HotTablet {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.HotTablet{}
// 	out.Name = LazyPtr(in.Name)
// 	out.TableName = LazyPtr(in.TableName)
// 	out.StartTime = Timestamp_FromProto(ctx, in.StartTime)
// 	out.EndTime = Timestamp_FromProto(ctx, in.EndTime)
// 	out.StartKey = LazyPtr(in.StartKey)
// 	out.EndKey = LazyPtr(in.EndKey)
// 	out.NodeCpuUsagePercent = LazyPtr(in.NodeCpuUsagePercent)
// 	return out
// }
// func HotTablet_ToProto(ctx *MapContext, in *krm.HotTablet) *pb.HotTablet {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.HotTablet{}
// 	out.Name = LazyPtr(in.Name)
// 	out.TableName = LazyPtr(in.TableName)
// 	out.StartTime = Timestamp_ToProto(ctx, in.StartTime)
// 	out.EndTime = Timestamp_ToProto(ctx, in.EndTime)
// 	out.StartKey = LazyPtr(in.StartKey)
// 	out.EndKey = LazyPtr(in.EndKey)
// 	out.NodeCpuUsagePercent = LazyPtr(in.NodeCpuUsagePercent)
// 	return out
// }
