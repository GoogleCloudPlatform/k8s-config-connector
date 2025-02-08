// Copyright 2025 Google LLC
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

package securitycenter

import (
	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func Access_FromProto(mapCtx *direct.MapContext, in *pb.Access) *krm.Access {
	if in == nil {
		return nil
	}
	out := &krm.Access{}
	out.PrincipalEmail = direct.LazyPtr(in.GetPrincipalEmail())
	out.CallerIP = direct.LazyPtr(in.GetCallerIp())
	out.CallerIPGeo = Geolocation_FromProto(mapCtx, in.GetCallerIpGeo())
	out.UserAgentFamily = direct.LazyPtr(in.GetUserAgentFamily())
	out.UserAgent = direct.LazyPtr(in.GetUserAgent())
	out.ServiceName = direct.LazyPtr(in.GetServiceName())
	out.MethodName = direct.LazyPtr(in.GetMethodName())
	out.PrincipalSubject = direct.LazyPtr(in.GetPrincipalSubject())
	out.ServiceAccountKeyName = direct.LazyPtr(in.GetServiceAccountKeyName())
	out.ServiceAccountDelegationInfo = direct.Slice_FromProto(mapCtx, in.ServiceAccountDelegationInfo, ServiceAccountDelegationInfo_FromProto)
	out.UserName = direct.LazyPtr(in.GetUserName())
	return out
}
func Access_ToProto(mapCtx *direct.MapContext, in *krm.Access) *pb.Access {
	if in == nil {
		return nil
	}
	out := &pb.Access{}
	out.PrincipalEmail = direct.ValueOf(in.PrincipalEmail)
	out.CallerIp = direct.ValueOf(in.CallerIP)
	out.CallerIpGeo = Geolocation_ToProto(mapCtx, in.CallerIPGeo)
	out.UserAgentFamily = direct.ValueOf(in.UserAgentFamily)
	out.UserAgent = direct.ValueOf(in.UserAgent)
	out.ServiceName = direct.ValueOf(in.ServiceName)
	out.MethodName = direct.ValueOf(in.MethodName)
	out.PrincipalSubject = direct.ValueOf(in.PrincipalSubject)
	out.ServiceAccountKeyName = direct.ValueOf(in.ServiceAccountKeyName)
	out.ServiceAccountDelegationInfo = direct.Slice_ToProto(mapCtx, in.ServiceAccountDelegationInfo, ServiceAccountDelegationInfo_ToProto)
	out.UserName = direct.ValueOf(in.UserName)
	return out
}
func AdaptiveProtection_FromProto(mapCtx *direct.MapContext, in *pb.AdaptiveProtection) *krm.AdaptiveProtection {
	if in == nil {
		return nil
	}
	out := &krm.AdaptiveProtection{}
	out.Confidence = direct.LazyPtr(in.GetConfidence())
	return out
}
func AdaptiveProtection_ToProto(mapCtx *direct.MapContext, in *krm.AdaptiveProtection) *pb.AdaptiveProtection {
	if in == nil {
		return nil
	}
	out := &pb.AdaptiveProtection{}
	out.Confidence = direct.ValueOf(in.Confidence)
	return out
}
func Application_FromProto(mapCtx *direct.MapContext, in *pb.Application) *krm.Application {
	if in == nil {
		return nil
	}
	out := &krm.Application{}
	out.BaseURI = direct.LazyPtr(in.GetBaseUri())
	out.FullURI = direct.LazyPtr(in.GetFullUri())
	return out
}
func Application_ToProto(mapCtx *direct.MapContext, in *krm.Application) *pb.Application {
	if in == nil {
		return nil
	}
	out := &pb.Application{}
	out.BaseUri = direct.ValueOf(in.BaseURI)
	out.FullUri = direct.ValueOf(in.FullURI)
	return out
}
func Attack_FromProto(mapCtx *direct.MapContext, in *pb.Attack) *krm.Attack {
	if in == nil {
		return nil
	}
	out := &krm.Attack{}
	out.VolumePps = direct.LazyPtr(in.GetVolumePps())
	out.VolumeBps = direct.LazyPtr(in.GetVolumeBps())
	out.Classification = direct.LazyPtr(in.GetClassification())
	return out
}
func Attack_ToProto(mapCtx *direct.MapContext, in *krm.Attack) *pb.Attack {
	if in == nil {
		return nil
	}
	out := &pb.Attack{}
	out.VolumePps = direct.ValueOf(in.VolumePps)
	out.VolumeBps = direct.ValueOf(in.VolumeBps)
	out.Classification = direct.ValueOf(in.Classification)
	return out
}
func AttackExposure_FromProto(mapCtx *direct.MapContext, in *pb.AttackExposure) *krm.AttackExposure {
	if in == nil {
		return nil
	}
	out := &krm.AttackExposure{}
	out.Score = direct.LazyPtr(in.GetScore())
	out.LatestCalculationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLatestCalculationTime())
	out.AttackExposureResult = direct.LazyPtr(in.GetAttackExposureResult())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ExposedHighValueResourcesCount = direct.LazyPtr(in.GetExposedHighValueResourcesCount())
	out.ExposedMediumValueResourcesCount = direct.LazyPtr(in.GetExposedMediumValueResourcesCount())
	out.ExposedLowValueResourcesCount = direct.LazyPtr(in.GetExposedLowValueResourcesCount())
	return out
}
func AttackExposure_ToProto(mapCtx *direct.MapContext, in *krm.AttackExposure) *pb.AttackExposure {
	if in == nil {
		return nil
	}
	out := &pb.AttackExposure{}
	out.Score = direct.ValueOf(in.Score)
	out.LatestCalculationTime = direct.StringTimestamp_ToProto(mapCtx, in.LatestCalculationTime)
	out.AttackExposureResult = direct.ValueOf(in.AttackExposureResult)
	out.State = direct.Enum_ToProto[pb.AttackExposure_State](mapCtx, in.State)
	out.ExposedHighValueResourcesCount = direct.ValueOf(in.ExposedHighValueResourcesCount)
	out.ExposedMediumValueResourcesCount = direct.ValueOf(in.ExposedMediumValueResourcesCount)
	out.ExposedLowValueResourcesCount = direct.ValueOf(in.ExposedLowValueResourcesCount)
	return out
}
func BackupDisasterRecovery_FromProto(mapCtx *direct.MapContext, in *pb.BackupDisasterRecovery) *krm.BackupDisasterRecovery {
	if in == nil {
		return nil
	}
	out := &krm.BackupDisasterRecovery{}
	out.BackupTemplate = direct.LazyPtr(in.GetBackupTemplate())
	out.Policies = in.Policies
	out.Host = direct.LazyPtr(in.GetHost())
	out.Applications = in.Applications
	out.StoragePool = direct.LazyPtr(in.GetStoragePool())
	out.PolicyOptions = in.PolicyOptions
	out.Profile = direct.LazyPtr(in.GetProfile())
	out.Appliance = direct.LazyPtr(in.GetAppliance())
	out.BackupType = direct.LazyPtr(in.GetBackupType())
	out.BackupCreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetBackupCreateTime())
	return out
}
func BackupDisasterRecovery_ToProto(mapCtx *direct.MapContext, in *krm.BackupDisasterRecovery) *pb.BackupDisasterRecovery {
	if in == nil {
		return nil
	}
	out := &pb.BackupDisasterRecovery{}
	out.BackupTemplate = direct.ValueOf(in.BackupTemplate)
	out.Policies = in.Policies
	out.Host = direct.ValueOf(in.Host)
	out.Applications = in.Applications
	out.StoragePool = direct.ValueOf(in.StoragePool)
	out.PolicyOptions = in.PolicyOptions
	out.Profile = direct.ValueOf(in.Profile)
	out.Appliance = direct.ValueOf(in.Appliance)
	out.BackupType = direct.ValueOf(in.BackupType)
	out.BackupCreateTime = direct.StringTimestamp_ToProto(mapCtx, in.BackupCreateTime)
	return out
}
func CloudArmor_FromProto(mapCtx *direct.MapContext, in *pb.CloudArmor) *krm.CloudArmor {
	if in == nil {
		return nil
	}
	out := &krm.CloudArmor{}
	out.SecurityPolicy = SecurityPolicy_FromProto(mapCtx, in.GetSecurityPolicy())
	out.Requests = Requests_FromProto(mapCtx, in.GetRequests())
	out.AdaptiveProtection = AdaptiveProtection_FromProto(mapCtx, in.GetAdaptiveProtection())
	out.Attack = Attack_FromProto(mapCtx, in.GetAttack())
	out.ThreatVector = direct.LazyPtr(in.GetThreatVector())
	out.Duration = direct.StringDuration_FromProto(mapCtx, in.GetDuration())
	return out
}
func CloudArmor_ToProto(mapCtx *direct.MapContext, in *krm.CloudArmor) *pb.CloudArmor {
	if in == nil {
		return nil
	}
	out := &pb.CloudArmor{}
	out.SecurityPolicy = SecurityPolicy_ToProto(mapCtx, in.SecurityPolicy)
	out.Requests = Requests_ToProto(mapCtx, in.Requests)
	out.AdaptiveProtection = AdaptiveProtection_ToProto(mapCtx, in.AdaptiveProtection)
	out.Attack = Attack_ToProto(mapCtx, in.Attack)
	out.ThreatVector = direct.ValueOf(in.ThreatVector)
	out.Duration = direct.StringDuration_ToProto(mapCtx, in.Duration)
	return out
}
func CloudDlpDataProfile_FromProto(mapCtx *direct.MapContext, in *pb.CloudDlpDataProfile) *krm.CloudDlpDataProfile {
	if in == nil {
		return nil
	}
	out := &krm.CloudDlpDataProfile{}
	out.DataProfile = direct.LazyPtr(in.GetDataProfile())
	out.ParentType = direct.Enum_FromProto(mapCtx, in.GetParentType())
	return out
}
func CloudDlpDataProfile_ToProto(mapCtx *direct.MapContext, in *krm.CloudDlpDataProfile) *pb.CloudDlpDataProfile {
	if in == nil {
		return nil
	}
	out := &pb.CloudDlpDataProfile{}
	out.DataProfile = direct.ValueOf(in.DataProfile)
	out.ParentType = direct.Enum_ToProto[pb.CloudDlpDataProfile_ParentType](mapCtx, in.ParentType)
	return out
}
func CloudDlpInspection_FromProto(mapCtx *direct.MapContext, in *pb.CloudDlpInspection) *krm.CloudDlpInspection {
	if in == nil {
		return nil
	}
	out := &krm.CloudDlpInspection{}
	out.InspectJob = direct.LazyPtr(in.GetInspectJob())
	out.InfoType = direct.LazyPtr(in.GetInfoType())
	out.InfoTypeCount = direct.LazyPtr(in.GetInfoTypeCount())
	out.FullScan = direct.LazyPtr(in.GetFullScan())
	return out
}
func CloudDlpInspection_ToProto(mapCtx *direct.MapContext, in *krm.CloudDlpInspection) *pb.CloudDlpInspection {
	if in == nil {
		return nil
	}
	out := &pb.CloudDlpInspection{}
	out.InspectJob = direct.ValueOf(in.InspectJob)
	out.InfoType = direct.ValueOf(in.InfoType)
	out.InfoTypeCount = direct.ValueOf(in.InfoTypeCount)
	out.FullScan = direct.ValueOf(in.FullScan)
	return out
}
func CloudLoggingEntry_FromProto(mapCtx *direct.MapContext, in *pb.CloudLoggingEntry) *krm.CloudLoggingEntry {
	if in == nil {
		return nil
	}
	out := &krm.CloudLoggingEntry{}
	out.InsertID = direct.LazyPtr(in.GetInsertId())
	out.LogID = direct.LazyPtr(in.GetLogId())
	out.ResourceContainer = direct.LazyPtr(in.GetResourceContainer())
	out.Timestamp = direct.StringTimestamp_FromProto(mapCtx, in.GetTimestamp())
	return out
}
func CloudLoggingEntry_ToProto(mapCtx *direct.MapContext, in *krm.CloudLoggingEntry) *pb.CloudLoggingEntry {
	if in == nil {
		return nil
	}
	out := &pb.CloudLoggingEntry{}
	out.InsertId = direct.ValueOf(in.InsertID)
	out.LogId = direct.ValueOf(in.LogID)
	out.ResourceContainer = direct.ValueOf(in.ResourceContainer)
	out.Timestamp = direct.StringTimestamp_ToProto(mapCtx, in.Timestamp)
	return out
}
func Compliance_FromProto(mapCtx *direct.MapContext, in *pb.Compliance) *krm.Compliance {
	if in == nil {
		return nil
	}
	out := &krm.Compliance{}
	out.Standard = direct.LazyPtr(in.GetStandard())
	out.Version = direct.LazyPtr(in.GetVersion())
	out.Ids = in.Ids
	return out
}
func Compliance_ToProto(mapCtx *direct.MapContext, in *krm.Compliance) *pb.Compliance {
	if in == nil {
		return nil
	}
	out := &pb.Compliance{}
	out.Standard = direct.ValueOf(in.Standard)
	out.Version = direct.ValueOf(in.Version)
	out.Ids = in.Ids
	return out
}
func Connection_FromProto(mapCtx *direct.MapContext, in *pb.Connection) *krm.Connection {
	if in == nil {
		return nil
	}
	out := &krm.Connection{}
	out.DestinationIP = direct.LazyPtr(in.GetDestinationIp())
	out.DestinationPort = direct.LazyPtr(in.GetDestinationPort())
	out.SourceIP = direct.LazyPtr(in.GetSourceIp())
	out.SourcePort = direct.LazyPtr(in.GetSourcePort())
	out.Protocol = direct.Enum_FromProto(mapCtx, in.GetProtocol())
	return out
}
func Connection_ToProto(mapCtx *direct.MapContext, in *krm.Connection) *pb.Connection {
	if in == nil {
		return nil
	}
	out := &pb.Connection{}
	out.DestinationIp = direct.ValueOf(in.DestinationIP)
	out.DestinationPort = direct.ValueOf(in.DestinationPort)
	out.SourceIp = direct.ValueOf(in.SourceIP)
	out.SourcePort = direct.ValueOf(in.SourcePort)
	out.Protocol = direct.Enum_ToProto[pb.Connection_Protocol](mapCtx, in.Protocol)
	return out
}
func Contact_FromProto(mapCtx *direct.MapContext, in *pb.Contact) *krm.Contact {
	if in == nil {
		return nil
	}
	out := &krm.Contact{}
	out.Email = direct.LazyPtr(in.GetEmail())
	return out
}
func Contact_ToProto(mapCtx *direct.MapContext, in *krm.Contact) *pb.Contact {
	if in == nil {
		return nil
	}
	out := &pb.Contact{}
	out.Email = direct.ValueOf(in.Email)
	return out
}
func ContactDetails_FromProto(mapCtx *direct.MapContext, in *pb.ContactDetails) *krm.ContactDetails {
	if in == nil {
		return nil
	}
	out := &krm.ContactDetails{}
	out.Contacts = direct.Slice_FromProto(mapCtx, in.Contacts, Contact_FromProto)
	return out
}
func ContactDetails_ToProto(mapCtx *direct.MapContext, in *krm.ContactDetails) *pb.ContactDetails {
	if in == nil {
		return nil
	}
	out := &pb.ContactDetails{}
	out.Contacts = direct.Slice_ToProto(mapCtx, in.Contacts, Contact_ToProto)
	return out
}
func Container_FromProto(mapCtx *direct.MapContext, in *pb.Container) *krm.Container {
	if in == nil {
		return nil
	}
	out := &krm.Container{}
	out.Name = direct.LazyPtr(in.GetName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.ImageID = direct.LazyPtr(in.GetImageId())
	out.Labels = direct.Slice_FromProto(mapCtx, in.Labels, Label_FromProto)
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	return out
}
func Container_ToProto(mapCtx *direct.MapContext, in *krm.Container) *pb.Container {
	if in == nil {
		return nil
	}
	out := &pb.Container{}
	out.Name = direct.ValueOf(in.Name)
	out.Uri = direct.ValueOf(in.URI)
	out.ImageId = direct.ValueOf(in.ImageID)
	out.Labels = direct.Slice_ToProto(mapCtx, in.Labels, Label_ToProto)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	return out
}
func Cve_FromProto(mapCtx *direct.MapContext, in *pb.Cve) *krm.Cve {
	if in == nil {
		return nil
	}
	out := &krm.Cve{}
	out.ID = direct.LazyPtr(in.GetId())
	out.References = direct.Slice_FromProto(mapCtx, in.References, Reference_FromProto)
	out.Cvssv3 = Cvssv3_FromProto(mapCtx, in.GetCvssv3())
	out.UpstreamFixAvailable = direct.LazyPtr(in.GetUpstreamFixAvailable())
	out.Impact = direct.Enum_FromProto(mapCtx, in.GetImpact())
	out.ExploitationActivity = direct.Enum_FromProto(mapCtx, in.GetExploitationActivity())
	out.ObservedInTheWild = direct.LazyPtr(in.GetObservedInTheWild())
	out.ZeroDay = direct.LazyPtr(in.GetZeroDay())
	out.ExploitReleaseDate = direct.StringTimestamp_FromProto(mapCtx, in.GetExploitReleaseDate())
	return out
}
func Cve_ToProto(mapCtx *direct.MapContext, in *krm.Cve) *pb.Cve {
	if in == nil {
		return nil
	}
	out := &pb.Cve{}
	out.Id = direct.ValueOf(in.ID)
	out.References = direct.Slice_ToProto(mapCtx, in.References, Reference_ToProto)
	out.Cvssv3 = Cvssv3_ToProto(mapCtx, in.Cvssv3)
	out.UpstreamFixAvailable = direct.ValueOf(in.UpstreamFixAvailable)
	out.Impact = direct.Enum_ToProto[pb.Cve_RiskRating](mapCtx, in.Impact)
	out.ExploitationActivity = direct.Enum_ToProto[pb.Cve_ExploitationActivity](mapCtx, in.ExploitationActivity)
	out.ObservedInTheWild = direct.ValueOf(in.ObservedInTheWild)
	out.ZeroDay = direct.ValueOf(in.ZeroDay)
	out.ExploitReleaseDate = direct.StringTimestamp_ToProto(mapCtx, in.ExploitReleaseDate)
	return out
}
func Cvssv3_FromProto(mapCtx *direct.MapContext, in *pb.Cvssv3) *krm.Cvssv3 {
	if in == nil {
		return nil
	}
	out := &krm.Cvssv3{}
	out.BaseScore = direct.LazyPtr(in.GetBaseScore())
	out.AttackVector = direct.Enum_FromProto(mapCtx, in.GetAttackVector())
	out.AttackComplexity = direct.Enum_FromProto(mapCtx, in.GetAttackComplexity())
	out.PrivilegesRequired = direct.Enum_FromProto(mapCtx, in.GetPrivilegesRequired())
	out.UserInteraction = direct.Enum_FromProto(mapCtx, in.GetUserInteraction())
	out.Scope = direct.Enum_FromProto(mapCtx, in.GetScope())
	out.ConfidentialityImpact = direct.Enum_FromProto(mapCtx, in.GetConfidentialityImpact())
	out.IntegrityImpact = direct.Enum_FromProto(mapCtx, in.GetIntegrityImpact())
	out.AvailabilityImpact = direct.Enum_FromProto(mapCtx, in.GetAvailabilityImpact())
	return out
}
func Cvssv3_ToProto(mapCtx *direct.MapContext, in *krm.Cvssv3) *pb.Cvssv3 {
	if in == nil {
		return nil
	}
	out := &pb.Cvssv3{}
	out.BaseScore = direct.ValueOf(in.BaseScore)
	out.AttackVector = direct.Enum_ToProto[pb.Cvssv3_AttackVector](mapCtx, in.AttackVector)
	out.AttackComplexity = direct.Enum_ToProto[pb.Cvssv3_AttackComplexity](mapCtx, in.AttackComplexity)
	out.PrivilegesRequired = direct.Enum_ToProto[pb.Cvssv3_PrivilegesRequired](mapCtx, in.PrivilegesRequired)
	out.UserInteraction = direct.Enum_ToProto[pb.Cvssv3_UserInteraction](mapCtx, in.UserInteraction)
	out.Scope = direct.Enum_ToProto[pb.Cvssv3_Scope](mapCtx, in.Scope)
	out.ConfidentialityImpact = direct.Enum_ToProto[pb.Cvssv3_Impact](mapCtx, in.ConfidentialityImpact)
	out.IntegrityImpact = direct.Enum_ToProto[pb.Cvssv3_Impact](mapCtx, in.IntegrityImpact)
	out.AvailabilityImpact = direct.Enum_ToProto[pb.Cvssv3_Impact](mapCtx, in.AvailabilityImpact)
	return out
}
func Database_FromProto(mapCtx *direct.MapContext, in *pb.Database) *krm.Database {
	if in == nil {
		return nil
	}
	out := &krm.Database{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.UserName = direct.LazyPtr(in.GetUserName())
	out.Query = direct.LazyPtr(in.GetQuery())
	out.Grantees = in.Grantees
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}
func Database_ToProto(mapCtx *direct.MapContext, in *krm.Database) *pb.Database {
	if in == nil {
		return nil
	}
	out := &pb.Database{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.UserName = direct.ValueOf(in.UserName)
	out.Query = direct.ValueOf(in.Query)
	out.Grantees = in.Grantees
	out.Version = direct.ValueOf(in.Version)
	return out
}
func EnvironmentVariable_FromProto(mapCtx *direct.MapContext, in *pb.EnvironmentVariable) *krm.EnvironmentVariable {
	if in == nil {
		return nil
	}
	out := &krm.EnvironmentVariable{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Val = direct.LazyPtr(in.GetVal())
	return out
}
func EnvironmentVariable_ToProto(mapCtx *direct.MapContext, in *krm.EnvironmentVariable) *pb.EnvironmentVariable {
	if in == nil {
		return nil
	}
	out := &pb.EnvironmentVariable{}
	out.Name = direct.ValueOf(in.Name)
	out.Val = direct.ValueOf(in.Val)
	return out
}
func ExfilResource_FromProto(mapCtx *direct.MapContext, in *pb.ExfilResource) *krm.ExfilResource {
	if in == nil {
		return nil
	}
	out := &krm.ExfilResource{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Components = in.Components
	return out
}
func ExfilResource_ToProto(mapCtx *direct.MapContext, in *krm.ExfilResource) *pb.ExfilResource {
	if in == nil {
		return nil
	}
	out := &pb.ExfilResource{}
	out.Name = direct.ValueOf(in.Name)
	out.Components = in.Components
	return out
}
func Exfiltration_FromProto(mapCtx *direct.MapContext, in *pb.Exfiltration) *krm.Exfiltration {
	if in == nil {
		return nil
	}
	out := &krm.Exfiltration{}
	out.Sources = direct.Slice_FromProto(mapCtx, in.Sources, ExfilResource_FromProto)
	out.Targets = direct.Slice_FromProto(mapCtx, in.Targets, ExfilResource_FromProto)
	out.TotalExfiltratedBytes = direct.LazyPtr(in.GetTotalExfiltratedBytes())
	return out
}
func Exfiltration_ToProto(mapCtx *direct.MapContext, in *krm.Exfiltration) *pb.Exfiltration {
	if in == nil {
		return nil
	}
	out := &pb.Exfiltration{}
	out.Sources = direct.Slice_ToProto(mapCtx, in.Sources, ExfilResource_ToProto)
	out.Targets = direct.Slice_ToProto(mapCtx, in.Targets, ExfilResource_ToProto)
	out.TotalExfiltratedBytes = direct.ValueOf(in.TotalExfiltratedBytes)
	return out
}
func ExternalSystem_FromProto(mapCtx *direct.MapContext, in *pb.ExternalSystem) *krm.ExternalSystem {
	if in == nil {
		return nil
	}
	out := &krm.ExternalSystem{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Assignees = in.Assignees
	out.ExternalUid = direct.LazyPtr(in.GetExternalUid())
	out.Status = direct.LazyPtr(in.GetStatus())
	out.ExternalSystemUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExternalSystemUpdateTime())
	out.CaseURI = direct.LazyPtr(in.GetCaseUri())
	out.CasePriority = direct.LazyPtr(in.GetCasePriority())
	out.CaseSla = direct.StringTimestamp_FromProto(mapCtx, in.GetCaseSla())
	out.CaseCreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCaseCreateTime())
	out.CaseCloseTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCaseCloseTime())
	out.TicketInfo = ExternalSystem_TicketInfo_FromProto(mapCtx, in.GetTicketInfo())
	return out
}
func ExternalSystem_ToProto(mapCtx *direct.MapContext, in *krm.ExternalSystem) *pb.ExternalSystem {
	if in == nil {
		return nil
	}
	out := &pb.ExternalSystem{}
	out.Name = direct.ValueOf(in.Name)
	out.Assignees = in.Assignees
	out.ExternalUid = direct.ValueOf(in.ExternalUid)
	out.Status = direct.ValueOf(in.Status)
	out.ExternalSystemUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.ExternalSystemUpdateTime)
	out.CaseUri = direct.ValueOf(in.CaseURI)
	out.CasePriority = direct.ValueOf(in.CasePriority)
	out.CaseSla = direct.StringTimestamp_ToProto(mapCtx, in.CaseSla)
	out.CaseCreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CaseCreateTime)
	out.CaseCloseTime = direct.StringTimestamp_ToProto(mapCtx, in.CaseCloseTime)
	out.TicketInfo = ExternalSystem_TicketInfo_ToProto(mapCtx, in.TicketInfo)
	return out
}
func ExternalSystem_TicketInfo_FromProto(mapCtx *direct.MapContext, in *pb.ExternalSystem_TicketInfo) *krm.ExternalSystem_TicketInfo {
	if in == nil {
		return nil
	}
	out := &krm.ExternalSystem_TicketInfo{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Assignee = direct.LazyPtr(in.GetAssignee())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.URI = direct.LazyPtr(in.GetUri())
	out.Status = direct.LazyPtr(in.GetStatus())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func ExternalSystem_TicketInfo_ToProto(mapCtx *direct.MapContext, in *krm.ExternalSystem_TicketInfo) *pb.ExternalSystem_TicketInfo {
	if in == nil {
		return nil
	}
	out := &pb.ExternalSystem_TicketInfo{}
	out.Id = direct.ValueOf(in.ID)
	out.Assignee = direct.ValueOf(in.Assignee)
	out.Description = direct.ValueOf(in.Description)
	out.Uri = direct.ValueOf(in.URI)
	out.Status = direct.ValueOf(in.Status)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func File_FromProto(mapCtx *direct.MapContext, in *pb.File) *krm.File {
	if in == nil {
		return nil
	}
	out := &krm.File{}
	out.Path = direct.LazyPtr(in.GetPath())
	out.Size = direct.LazyPtr(in.GetSize())
	out.Sha256 = direct.LazyPtr(in.GetSha256())
	out.HashedSize = direct.LazyPtr(in.GetHashedSize())
	out.PartiallyHashed = direct.LazyPtr(in.GetPartiallyHashed())
	out.Contents = direct.LazyPtr(in.GetContents())
	out.DiskPath = File_DiskPath_FromProto(mapCtx, in.GetDiskPath())
	return out
}
func File_ToProto(mapCtx *direct.MapContext, in *krm.File) *pb.File {
	if in == nil {
		return nil
	}
	out := &pb.File{}
	out.Path = direct.ValueOf(in.Path)
	out.Size = direct.ValueOf(in.Size)
	out.Sha256 = direct.ValueOf(in.Sha256)
	out.HashedSize = direct.ValueOf(in.HashedSize)
	out.PartiallyHashed = direct.ValueOf(in.PartiallyHashed)
	out.Contents = direct.ValueOf(in.Contents)
	out.DiskPath = File_DiskPath_ToProto(mapCtx, in.DiskPath)
	return out
}
func File_DiskPath_FromProto(mapCtx *direct.MapContext, in *pb.File_DiskPath) *krm.File_DiskPath {
	if in == nil {
		return nil
	}
	out := &krm.File_DiskPath{}
	out.PartitionUuid = direct.LazyPtr(in.GetPartitionUuid())
	out.RelativePath = direct.LazyPtr(in.GetRelativePath())
	return out
}
func File_DiskPath_ToProto(mapCtx *direct.MapContext, in *krm.File_DiskPath) *pb.File_DiskPath {
	if in == nil {
		return nil
	}
	out := &pb.File_DiskPath{}
	out.PartitionUuid = direct.ValueOf(in.PartitionUuid)
	out.RelativePath = direct.ValueOf(in.RelativePath)
	return out
}
func Finding_FromProto(mapCtx *direct.MapContext, in *pb.Finding) *krm.Finding {
	if in == nil {
		return nil
	}
	out := &krm.Finding{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Parent = direct.LazyPtr(in.GetParent())
	out.ResourceName = direct.LazyPtr(in.GetResourceName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Category = direct.LazyPtr(in.GetCategory())
	out.ExternalURI = direct.LazyPtr(in.GetExternalUri())
	// MISSING: SourceProperties
	// MISSING: SecurityMarks
	out.EventTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEventTime())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Severity = direct.Enum_FromProto(mapCtx, in.GetSeverity())
	out.CanonicalName = direct.LazyPtr(in.GetCanonicalName())
	out.Mute = direct.Enum_FromProto(mapCtx, in.GetMute())
	out.FindingClass = direct.Enum_FromProto(mapCtx, in.GetFindingClass())
	out.Indicator = Indicator_FromProto(mapCtx, in.GetIndicator())
	out.Vulnerability = Vulnerability_FromProto(mapCtx, in.GetVulnerability())
	// MISSING: MuteUpdateTime
	// MISSING: ExternalSystems
	out.MitreAttack = MitreAttack_FromProto(mapCtx, in.GetMitreAttack())
	out.Access = Access_FromProto(mapCtx, in.GetAccess())
	out.Connections = direct.Slice_FromProto(mapCtx, in.Connections, Connection_FromProto)
	out.MuteInitiator = direct.LazyPtr(in.GetMuteInitiator())
	// MISSING: MuteInfo
	out.Processes = direct.Slice_FromProto(mapCtx, in.Processes, Process_FromProto)
	// MISSING: Contacts
	out.Compliances = direct.Slice_FromProto(mapCtx, in.Compliances, Compliance_FromProto)
	// MISSING: ParentDisplayName
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Exfiltration = Exfiltration_FromProto(mapCtx, in.GetExfiltration())
	out.IamBindings = direct.Slice_FromProto(mapCtx, in.IamBindings, IamBinding_FromProto)
	out.NextSteps = direct.LazyPtr(in.GetNextSteps())
	out.ModuleName = direct.LazyPtr(in.GetModuleName())
	out.Containers = direct.Slice_FromProto(mapCtx, in.Containers, Container_FromProto)
	out.Kubernetes = Kubernetes_FromProto(mapCtx, in.GetKubernetes())
	out.Database = Database_FromProto(mapCtx, in.GetDatabase())
	out.AttackExposure = AttackExposure_FromProto(mapCtx, in.GetAttackExposure())
	out.Files = direct.Slice_FromProto(mapCtx, in.Files, File_FromProto)
	out.CloudDlpInspection = CloudDlpInspection_FromProto(mapCtx, in.GetCloudDlpInspection())
	out.CloudDlpDataProfile = CloudDlpDataProfile_FromProto(mapCtx, in.GetCloudDlpDataProfile())
	out.KernelRootkit = KernelRootkit_FromProto(mapCtx, in.GetKernelRootkit())
	out.OrgPolicies = direct.Slice_FromProto(mapCtx, in.OrgPolicies, OrgPolicy_FromProto)
	out.Application = Application_FromProto(mapCtx, in.GetApplication())
	out.BackupDisasterRecovery = BackupDisasterRecovery_FromProto(mapCtx, in.GetBackupDisasterRecovery())
	out.SecurityPosture = SecurityPosture_FromProto(mapCtx, in.GetSecurityPosture())
	out.LogEntries = direct.Slice_FromProto(mapCtx, in.LogEntries, LogEntry_FromProto)
	out.LoadBalancers = direct.Slice_FromProto(mapCtx, in.LoadBalancers, LoadBalancer_FromProto)
	out.CloudArmor = CloudArmor_FromProto(mapCtx, in.GetCloudArmor())
	out.Notebook = Notebook_FromProto(mapCtx, in.GetNotebook())
	out.ToxicCombination = ToxicCombination_FromProto(mapCtx, in.GetToxicCombination())
	out.GroupMemberships = direct.Slice_FromProto(mapCtx, in.GroupMemberships, GroupMembership_FromProto)
	return out
}
func Finding_ToProto(mapCtx *direct.MapContext, in *krm.Finding) *pb.Finding {
	if in == nil {
		return nil
	}
	out := &pb.Finding{}
	out.Name = direct.ValueOf(in.Name)
	out.Parent = direct.ValueOf(in.Parent)
	out.ResourceName = direct.ValueOf(in.ResourceName)
	out.State = direct.Enum_ToProto[pb.Finding_State](mapCtx, in.State)
	out.Category = direct.ValueOf(in.Category)
	out.ExternalUri = direct.ValueOf(in.ExternalURI)
	// MISSING: SourceProperties
	// MISSING: SecurityMarks
	out.EventTime = direct.StringTimestamp_ToProto(mapCtx, in.EventTime)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.Severity = direct.Enum_ToProto[pb.Finding_Severity](mapCtx, in.Severity)
	out.CanonicalName = direct.ValueOf(in.CanonicalName)
	out.Mute = direct.Enum_ToProto[pb.Finding_Mute](mapCtx, in.Mute)
	out.FindingClass = direct.Enum_ToProto[pb.Finding_FindingClass](mapCtx, in.FindingClass)
	out.Indicator = Indicator_ToProto(mapCtx, in.Indicator)
	out.Vulnerability = Vulnerability_ToProto(mapCtx, in.Vulnerability)
	// MISSING: MuteUpdateTime
	// MISSING: ExternalSystems
	out.MitreAttack = MitreAttack_ToProto(mapCtx, in.MitreAttack)
	out.Access = Access_ToProto(mapCtx, in.Access)
	out.Connections = direct.Slice_ToProto(mapCtx, in.Connections, Connection_ToProto)
	out.MuteInitiator = direct.ValueOf(in.MuteInitiator)
	// MISSING: MuteInfo
	out.Processes = direct.Slice_ToProto(mapCtx, in.Processes, Process_ToProto)
	// MISSING: Contacts
	out.Compliances = direct.Slice_ToProto(mapCtx, in.Compliances, Compliance_ToProto)
	// MISSING: ParentDisplayName
	out.Description = direct.ValueOf(in.Description)
	out.Exfiltration = Exfiltration_ToProto(mapCtx, in.Exfiltration)
	out.IamBindings = direct.Slice_ToProto(mapCtx, in.IamBindings, IamBinding_ToProto)
	out.NextSteps = direct.ValueOf(in.NextSteps)
	out.ModuleName = direct.ValueOf(in.ModuleName)
	out.Containers = direct.Slice_ToProto(mapCtx, in.Containers, Container_ToProto)
	out.Kubernetes = Kubernetes_ToProto(mapCtx, in.Kubernetes)
	out.Database = Database_ToProto(mapCtx, in.Database)
	out.AttackExposure = AttackExposure_ToProto(mapCtx, in.AttackExposure)
	out.Files = direct.Slice_ToProto(mapCtx, in.Files, File_ToProto)
	out.CloudDlpInspection = CloudDlpInspection_ToProto(mapCtx, in.CloudDlpInspection)
	out.CloudDlpDataProfile = CloudDlpDataProfile_ToProto(mapCtx, in.CloudDlpDataProfile)
	out.KernelRootkit = KernelRootkit_ToProto(mapCtx, in.KernelRootkit)
	out.OrgPolicies = direct.Slice_ToProto(mapCtx, in.OrgPolicies, OrgPolicy_ToProto)
	out.Application = Application_ToProto(mapCtx, in.Application)
	out.BackupDisasterRecovery = BackupDisasterRecovery_ToProto(mapCtx, in.BackupDisasterRecovery)
	out.SecurityPosture = SecurityPosture_ToProto(mapCtx, in.SecurityPosture)
	out.LogEntries = direct.Slice_ToProto(mapCtx, in.LogEntries, LogEntry_ToProto)
	out.LoadBalancers = direct.Slice_ToProto(mapCtx, in.LoadBalancers, LoadBalancer_ToProto)
	out.CloudArmor = CloudArmor_ToProto(mapCtx, in.CloudArmor)
	out.Notebook = Notebook_ToProto(mapCtx, in.Notebook)
	out.ToxicCombination = ToxicCombination_ToProto(mapCtx, in.ToxicCombination)
	out.GroupMemberships = direct.Slice_ToProto(mapCtx, in.GroupMemberships, GroupMembership_ToProto)
	return out
}
func FindingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Finding) *krm.FindingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FindingObservedState{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ResourceName
	// MISSING: State
	// MISSING: Category
	// MISSING: ExternalURI
	// MISSING: SourceProperties
	out.SecurityMarks = SecurityMarks_FromProto(mapCtx, in.GetSecurityMarks())
	// MISSING: EventTime
	// MISSING: CreateTime
	// MISSING: Severity
	// MISSING: CanonicalName
	// MISSING: Mute
	// MISSING: FindingClass
	// MISSING: Indicator
	// MISSING: Vulnerability
	out.MuteUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetMuteUpdateTime())
	// MISSING: ExternalSystems
	// MISSING: MitreAttack
	// MISSING: Access
	// MISSING: Connections
	// MISSING: MuteInitiator
	out.MuteInfo = Finding_MuteInfo_FromProto(mapCtx, in.GetMuteInfo())
	// MISSING: Processes
	// MISSING: Contacts
	// MISSING: Compliances
	out.ParentDisplayName = direct.LazyPtr(in.GetParentDisplayName())
	// MISSING: Description
	// MISSING: Exfiltration
	// MISSING: IamBindings
	// MISSING: NextSteps
	// MISSING: ModuleName
	// MISSING: Containers
	// MISSING: Kubernetes
	// MISSING: Database
	// MISSING: AttackExposure
	// MISSING: Files
	// MISSING: CloudDlpInspection
	// MISSING: CloudDlpDataProfile
	// MISSING: KernelRootkit
	// MISSING: OrgPolicies
	// MISSING: Application
	// MISSING: BackupDisasterRecovery
	// MISSING: SecurityPosture
	// MISSING: LogEntries
	// MISSING: LoadBalancers
	// MISSING: CloudArmor
	// MISSING: Notebook
	// MISSING: ToxicCombination
	// MISSING: GroupMemberships
	return out
}
func FindingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FindingObservedState) *pb.Finding {
	if in == nil {
		return nil
	}
	out := &pb.Finding{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ResourceName
	// MISSING: State
	// MISSING: Category
	// MISSING: ExternalURI
	// MISSING: SourceProperties
	out.SecurityMarks = SecurityMarks_ToProto(mapCtx, in.SecurityMarks)
	// MISSING: EventTime
	// MISSING: CreateTime
	// MISSING: Severity
	// MISSING: CanonicalName
	// MISSING: Mute
	// MISSING: FindingClass
	// MISSING: Indicator
	// MISSING: Vulnerability
	out.MuteUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.MuteUpdateTime)
	// MISSING: ExternalSystems
	// MISSING: MitreAttack
	// MISSING: Access
	// MISSING: Connections
	// MISSING: MuteInitiator
	out.MuteInfo = Finding_MuteInfo_ToProto(mapCtx, in.MuteInfo)
	// MISSING: Processes
	// MISSING: Contacts
	// MISSING: Compliances
	out.ParentDisplayName = direct.ValueOf(in.ParentDisplayName)
	// MISSING: Description
	// MISSING: Exfiltration
	// MISSING: IamBindings
	// MISSING: NextSteps
	// MISSING: ModuleName
	// MISSING: Containers
	// MISSING: Kubernetes
	// MISSING: Database
	// MISSING: AttackExposure
	// MISSING: Files
	// MISSING: CloudDlpInspection
	// MISSING: CloudDlpDataProfile
	// MISSING: KernelRootkit
	// MISSING: OrgPolicies
	// MISSING: Application
	// MISSING: BackupDisasterRecovery
	// MISSING: SecurityPosture
	// MISSING: LogEntries
	// MISSING: LoadBalancers
	// MISSING: CloudArmor
	// MISSING: Notebook
	// MISSING: ToxicCombination
	// MISSING: GroupMemberships
	return out
}
func Finding_MuteInfo_FromProto(mapCtx *direct.MapContext, in *pb.Finding_MuteInfo) *krm.Finding_MuteInfo {
	if in == nil {
		return nil
	}
	out := &krm.Finding_MuteInfo{}
	out.StaticMute = Finding_MuteInfo_StaticMute_FromProto(mapCtx, in.GetStaticMute())
	out.DynamicMuteRecords = direct.Slice_FromProto(mapCtx, in.DynamicMuteRecords, Finding_MuteInfo_DynamicMuteRecord_FromProto)
	return out
}
func Finding_MuteInfo_ToProto(mapCtx *direct.MapContext, in *krm.Finding_MuteInfo) *pb.Finding_MuteInfo {
	if in == nil {
		return nil
	}
	out := &pb.Finding_MuteInfo{}
	out.StaticMute = Finding_MuteInfo_StaticMute_ToProto(mapCtx, in.StaticMute)
	out.DynamicMuteRecords = direct.Slice_ToProto(mapCtx, in.DynamicMuteRecords, Finding_MuteInfo_DynamicMuteRecord_ToProto)
	return out
}
func Finding_MuteInfo_DynamicMuteRecord_FromProto(mapCtx *direct.MapContext, in *pb.Finding_MuteInfo_DynamicMuteRecord) *krm.Finding_MuteInfo_DynamicMuteRecord {
	if in == nil {
		return nil
	}
	out := &krm.Finding_MuteInfo_DynamicMuteRecord{}
	out.MuteConfig = direct.LazyPtr(in.GetMuteConfig())
	out.MatchTime = direct.StringTimestamp_FromProto(mapCtx, in.GetMatchTime())
	return out
}
func Finding_MuteInfo_DynamicMuteRecord_ToProto(mapCtx *direct.MapContext, in *krm.Finding_MuteInfo_DynamicMuteRecord) *pb.Finding_MuteInfo_DynamicMuteRecord {
	if in == nil {
		return nil
	}
	out := &pb.Finding_MuteInfo_DynamicMuteRecord{}
	out.MuteConfig = direct.ValueOf(in.MuteConfig)
	out.MatchTime = direct.StringTimestamp_ToProto(mapCtx, in.MatchTime)
	return out
}
func Finding_MuteInfo_StaticMute_FromProto(mapCtx *direct.MapContext, in *pb.Finding_MuteInfo_StaticMute) *krm.Finding_MuteInfo_StaticMute {
	if in == nil {
		return nil
	}
	out := &krm.Finding_MuteInfo_StaticMute{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ApplyTime = direct.StringTimestamp_FromProto(mapCtx, in.GetApplyTime())
	return out
}
func Finding_MuteInfo_StaticMute_ToProto(mapCtx *direct.MapContext, in *krm.Finding_MuteInfo_StaticMute) *pb.Finding_MuteInfo_StaticMute {
	if in == nil {
		return nil
	}
	out := &pb.Finding_MuteInfo_StaticMute{}
	out.State = direct.Enum_ToProto[pb.Finding_Mute](mapCtx, in.State)
	out.ApplyTime = direct.StringTimestamp_ToProto(mapCtx, in.ApplyTime)
	return out
}
func Geolocation_FromProto(mapCtx *direct.MapContext, in *pb.Geolocation) *krm.Geolocation {
	if in == nil {
		return nil
	}
	out := &krm.Geolocation{}
	out.RegionCode = direct.LazyPtr(in.GetRegionCode())
	return out
}
func Geolocation_ToProto(mapCtx *direct.MapContext, in *krm.Geolocation) *pb.Geolocation {
	if in == nil {
		return nil
	}
	out := &pb.Geolocation{}
	out.RegionCode = direct.ValueOf(in.RegionCode)
	return out
}
func GroupMembership_FromProto(mapCtx *direct.MapContext, in *pb.GroupMembership) *krm.GroupMembership {
	if in == nil {
		return nil
	}
	out := &krm.GroupMembership{}
	out.GroupType = direct.Enum_FromProto(mapCtx, in.GetGroupType())
	out.GroupID = direct.LazyPtr(in.GetGroupId())
	return out
}
func GroupMembership_ToProto(mapCtx *direct.MapContext, in *krm.GroupMembership) *pb.GroupMembership {
	if in == nil {
		return nil
	}
	out := &pb.GroupMembership{}
	out.GroupType = direct.Enum_ToProto[pb.GroupMembership_GroupType](mapCtx, in.GroupType)
	out.GroupId = direct.ValueOf(in.GroupID)
	return out
}
func IamBinding_FromProto(mapCtx *direct.MapContext, in *pb.IamBinding) *krm.IamBinding {
	if in == nil {
		return nil
	}
	out := &krm.IamBinding{}
	out.Action = direct.Enum_FromProto(mapCtx, in.GetAction())
	out.Role = direct.LazyPtr(in.GetRole())
	out.Member = direct.LazyPtr(in.GetMember())
	return out
}
func IamBinding_ToProto(mapCtx *direct.MapContext, in *krm.IamBinding) *pb.IamBinding {
	if in == nil {
		return nil
	}
	out := &pb.IamBinding{}
	out.Action = direct.Enum_ToProto[pb.IamBinding_Action](mapCtx, in.Action)
	out.Role = direct.ValueOf(in.Role)
	out.Member = direct.ValueOf(in.Member)
	return out
}
func Indicator_FromProto(mapCtx *direct.MapContext, in *pb.Indicator) *krm.Indicator {
	if in == nil {
		return nil
	}
	out := &krm.Indicator{}
	out.IPAddresses = in.IpAddresses
	out.Domains = in.Domains
	out.Signatures = direct.Slice_FromProto(mapCtx, in.Signatures, Indicator_ProcessSignature_FromProto)
	out.Uris = in.Uris
	return out
}
func Indicator_ToProto(mapCtx *direct.MapContext, in *krm.Indicator) *pb.Indicator {
	if in == nil {
		return nil
	}
	out := &pb.Indicator{}
	out.IpAddresses = in.IPAddresses
	out.Domains = in.Domains
	out.Signatures = direct.Slice_ToProto(mapCtx, in.Signatures, Indicator_ProcessSignature_ToProto)
	out.Uris = in.Uris
	return out
}
func Indicator_ProcessSignature_FromProto(mapCtx *direct.MapContext, in *pb.Indicator_ProcessSignature) *krm.Indicator_ProcessSignature {
	if in == nil {
		return nil
	}
	out := &krm.Indicator_ProcessSignature{}
	out.MemoryHashSignature = Indicator_ProcessSignature_MemoryHashSignature_FromProto(mapCtx, in.GetMemoryHashSignature())
	out.YaraRuleSignature = Indicator_ProcessSignature_YaraRuleSignature_FromProto(mapCtx, in.GetYaraRuleSignature())
	out.SignatureType = direct.Enum_FromProto(mapCtx, in.GetSignatureType())
	return out
}
func Indicator_ProcessSignature_ToProto(mapCtx *direct.MapContext, in *krm.Indicator_ProcessSignature) *pb.Indicator_ProcessSignature {
	if in == nil {
		return nil
	}
	out := &pb.Indicator_ProcessSignature{}
	if oneof := Indicator_ProcessSignature_MemoryHashSignature_ToProto(mapCtx, in.MemoryHashSignature); oneof != nil {
		out.Signature = &pb.Indicator_ProcessSignature_MemoryHashSignature_{MemoryHashSignature: oneof}
	}
	if oneof := Indicator_ProcessSignature_YaraRuleSignature_ToProto(mapCtx, in.YaraRuleSignature); oneof != nil {
		out.Signature = &pb.Indicator_ProcessSignature_YaraRuleSignature_{YaraRuleSignature: oneof}
	}
	out.SignatureType = direct.Enum_ToProto[pb.Indicator_ProcessSignature_SignatureType](mapCtx, in.SignatureType)
	return out
}
func Indicator_ProcessSignature_MemoryHashSignature_FromProto(mapCtx *direct.MapContext, in *pb.Indicator_ProcessSignature_MemoryHashSignature) *krm.Indicator_ProcessSignature_MemoryHashSignature {
	if in == nil {
		return nil
	}
	out := &krm.Indicator_ProcessSignature_MemoryHashSignature{}
	out.BinaryFamily = direct.LazyPtr(in.GetBinaryFamily())
	out.Detections = direct.Slice_FromProto(mapCtx, in.Detections, Indicator_ProcessSignature_MemoryHashSignature_Detection_FromProto)
	return out
}
func Indicator_ProcessSignature_MemoryHashSignature_ToProto(mapCtx *direct.MapContext, in *krm.Indicator_ProcessSignature_MemoryHashSignature) *pb.Indicator_ProcessSignature_MemoryHashSignature {
	if in == nil {
		return nil
	}
	out := &pb.Indicator_ProcessSignature_MemoryHashSignature{}
	out.BinaryFamily = direct.ValueOf(in.BinaryFamily)
	out.Detections = direct.Slice_ToProto(mapCtx, in.Detections, Indicator_ProcessSignature_MemoryHashSignature_Detection_ToProto)
	return out
}
func Indicator_ProcessSignature_MemoryHashSignature_Detection_FromProto(mapCtx *direct.MapContext, in *pb.Indicator_ProcessSignature_MemoryHashSignature_Detection) *krm.Indicator_ProcessSignature_MemoryHashSignature_Detection {
	if in == nil {
		return nil
	}
	out := &krm.Indicator_ProcessSignature_MemoryHashSignature_Detection{}
	out.Binary = direct.LazyPtr(in.GetBinary())
	out.PercentPagesMatched = direct.LazyPtr(in.GetPercentPagesMatched())
	return out
}
func Indicator_ProcessSignature_MemoryHashSignature_Detection_ToProto(mapCtx *direct.MapContext, in *krm.Indicator_ProcessSignature_MemoryHashSignature_Detection) *pb.Indicator_ProcessSignature_MemoryHashSignature_Detection {
	if in == nil {
		return nil
	}
	out := &pb.Indicator_ProcessSignature_MemoryHashSignature_Detection{}
	out.Binary = direct.ValueOf(in.Binary)
	out.PercentPagesMatched = direct.ValueOf(in.PercentPagesMatched)
	return out
}
func Indicator_ProcessSignature_YaraRuleSignature_FromProto(mapCtx *direct.MapContext, in *pb.Indicator_ProcessSignature_YaraRuleSignature) *krm.Indicator_ProcessSignature_YaraRuleSignature {
	if in == nil {
		return nil
	}
	out := &krm.Indicator_ProcessSignature_YaraRuleSignature{}
	out.YaraRule = direct.LazyPtr(in.GetYaraRule())
	return out
}
func Indicator_ProcessSignature_YaraRuleSignature_ToProto(mapCtx *direct.MapContext, in *krm.Indicator_ProcessSignature_YaraRuleSignature) *pb.Indicator_ProcessSignature_YaraRuleSignature {
	if in == nil {
		return nil
	}
	out := &pb.Indicator_ProcessSignature_YaraRuleSignature{}
	out.YaraRule = direct.ValueOf(in.YaraRule)
	return out
}
func KernelRootkit_FromProto(mapCtx *direct.MapContext, in *pb.KernelRootkit) *krm.KernelRootkit {
	if in == nil {
		return nil
	}
	out := &krm.KernelRootkit{}
	out.Name = direct.LazyPtr(in.GetName())
	out.UnexpectedCodeModification = direct.LazyPtr(in.GetUnexpectedCodeModification())
	out.UnexpectedReadOnlyDataModification = direct.LazyPtr(in.GetUnexpectedReadOnlyDataModification())
	out.UnexpectedFtraceHandler = direct.LazyPtr(in.GetUnexpectedFtraceHandler())
	out.UnexpectedKprobeHandler = direct.LazyPtr(in.GetUnexpectedKprobeHandler())
	out.UnexpectedKernelCodePages = direct.LazyPtr(in.GetUnexpectedKernelCodePages())
	out.UnexpectedSystemCallHandler = direct.LazyPtr(in.GetUnexpectedSystemCallHandler())
	out.UnexpectedInterruptHandler = direct.LazyPtr(in.GetUnexpectedInterruptHandler())
	out.UnexpectedProcessesInRunqueue = direct.LazyPtr(in.GetUnexpectedProcessesInRunqueue())
	return out
}
func KernelRootkit_ToProto(mapCtx *direct.MapContext, in *krm.KernelRootkit) *pb.KernelRootkit {
	if in == nil {
		return nil
	}
	out := &pb.KernelRootkit{}
	out.Name = direct.ValueOf(in.Name)
	out.UnexpectedCodeModification = direct.ValueOf(in.UnexpectedCodeModification)
	out.UnexpectedReadOnlyDataModification = direct.ValueOf(in.UnexpectedReadOnlyDataModification)
	out.UnexpectedFtraceHandler = direct.ValueOf(in.UnexpectedFtraceHandler)
	out.UnexpectedKprobeHandler = direct.ValueOf(in.UnexpectedKprobeHandler)
	out.UnexpectedKernelCodePages = direct.ValueOf(in.UnexpectedKernelCodePages)
	out.UnexpectedSystemCallHandler = direct.ValueOf(in.UnexpectedSystemCallHandler)
	out.UnexpectedInterruptHandler = direct.ValueOf(in.UnexpectedInterruptHandler)
	out.UnexpectedProcessesInRunqueue = direct.ValueOf(in.UnexpectedProcessesInRunqueue)
	return out
}
func Kubernetes_FromProto(mapCtx *direct.MapContext, in *pb.Kubernetes) *krm.Kubernetes {
	if in == nil {
		return nil
	}
	out := &krm.Kubernetes{}
	out.Pods = direct.Slice_FromProto(mapCtx, in.Pods, Kubernetes_Pod_FromProto)
	out.Nodes = direct.Slice_FromProto(mapCtx, in.Nodes, Kubernetes_Node_FromProto)
	out.NodePools = direct.Slice_FromProto(mapCtx, in.NodePools, Kubernetes_NodePool_FromProto)
	out.Roles = direct.Slice_FromProto(mapCtx, in.Roles, Kubernetes_Role_FromProto)
	out.Bindings = direct.Slice_FromProto(mapCtx, in.Bindings, Kubernetes_Binding_FromProto)
	out.AccessReviews = direct.Slice_FromProto(mapCtx, in.AccessReviews, Kubernetes_AccessReview_FromProto)
	out.Objects = direct.Slice_FromProto(mapCtx, in.Objects, Kubernetes_Object_FromProto)
	return out
}
func Kubernetes_ToProto(mapCtx *direct.MapContext, in *krm.Kubernetes) *pb.Kubernetes {
	if in == nil {
		return nil
	}
	out := &pb.Kubernetes{}
	out.Pods = direct.Slice_ToProto(mapCtx, in.Pods, Kubernetes_Pod_ToProto)
	out.Nodes = direct.Slice_ToProto(mapCtx, in.Nodes, Kubernetes_Node_ToProto)
	out.NodePools = direct.Slice_ToProto(mapCtx, in.NodePools, Kubernetes_NodePool_ToProto)
	out.Roles = direct.Slice_ToProto(mapCtx, in.Roles, Kubernetes_Role_ToProto)
	out.Bindings = direct.Slice_ToProto(mapCtx, in.Bindings, Kubernetes_Binding_ToProto)
	out.AccessReviews = direct.Slice_ToProto(mapCtx, in.AccessReviews, Kubernetes_AccessReview_ToProto)
	out.Objects = direct.Slice_ToProto(mapCtx, in.Objects, Kubernetes_Object_ToProto)
	return out
}
func Kubernetes_AccessReview_FromProto(mapCtx *direct.MapContext, in *pb.Kubernetes_AccessReview) *krm.Kubernetes_AccessReview {
	if in == nil {
		return nil
	}
	out := &krm.Kubernetes_AccessReview{}
	out.Group = direct.LazyPtr(in.GetGroup())
	out.Ns = direct.LazyPtr(in.GetNs())
	out.Name = direct.LazyPtr(in.GetName())
	out.Resource = direct.LazyPtr(in.GetResource())
	out.Subresource = direct.LazyPtr(in.GetSubresource())
	out.Verb = direct.LazyPtr(in.GetVerb())
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}
func Kubernetes_AccessReview_ToProto(mapCtx *direct.MapContext, in *krm.Kubernetes_AccessReview) *pb.Kubernetes_AccessReview {
	if in == nil {
		return nil
	}
	out := &pb.Kubernetes_AccessReview{}
	out.Group = direct.ValueOf(in.Group)
	out.Ns = direct.ValueOf(in.Ns)
	out.Name = direct.ValueOf(in.Name)
	out.Resource = direct.ValueOf(in.Resource)
	out.Subresource = direct.ValueOf(in.Subresource)
	out.Verb = direct.ValueOf(in.Verb)
	out.Version = direct.ValueOf(in.Version)
	return out
}
func Kubernetes_Binding_FromProto(mapCtx *direct.MapContext, in *pb.Kubernetes_Binding) *krm.Kubernetes_Binding {
	if in == nil {
		return nil
	}
	out := &krm.Kubernetes_Binding{}
	out.Ns = direct.LazyPtr(in.GetNs())
	out.Name = direct.LazyPtr(in.GetName())
	out.Role = Kubernetes_Role_FromProto(mapCtx, in.GetRole())
	out.Subjects = direct.Slice_FromProto(mapCtx, in.Subjects, Kubernetes_Subject_FromProto)
	return out
}
func Kubernetes_Binding_ToProto(mapCtx *direct.MapContext, in *krm.Kubernetes_Binding) *pb.Kubernetes_Binding {
	if in == nil {
		return nil
	}
	out := &pb.Kubernetes_Binding{}
	out.Ns = direct.ValueOf(in.Ns)
	out.Name = direct.ValueOf(in.Name)
	out.Role = Kubernetes_Role_ToProto(mapCtx, in.Role)
	out.Subjects = direct.Slice_ToProto(mapCtx, in.Subjects, Kubernetes_Subject_ToProto)
	return out
}
func Kubernetes_Node_FromProto(mapCtx *direct.MapContext, in *pb.Kubernetes_Node) *krm.Kubernetes_Node {
	if in == nil {
		return nil
	}
	out := &krm.Kubernetes_Node{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func Kubernetes_Node_ToProto(mapCtx *direct.MapContext, in *krm.Kubernetes_Node) *pb.Kubernetes_Node {
	if in == nil {
		return nil
	}
	out := &pb.Kubernetes_Node{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
func Kubernetes_NodePool_FromProto(mapCtx *direct.MapContext, in *pb.Kubernetes_NodePool) *krm.Kubernetes_NodePool {
	if in == nil {
		return nil
	}
	out := &krm.Kubernetes_NodePool{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Nodes = direct.Slice_FromProto(mapCtx, in.Nodes, Kubernetes_Node_FromProto)
	return out
}
func Kubernetes_NodePool_ToProto(mapCtx *direct.MapContext, in *krm.Kubernetes_NodePool) *pb.Kubernetes_NodePool {
	if in == nil {
		return nil
	}
	out := &pb.Kubernetes_NodePool{}
	out.Name = direct.ValueOf(in.Name)
	out.Nodes = direct.Slice_ToProto(mapCtx, in.Nodes, Kubernetes_Node_ToProto)
	return out
}
func Kubernetes_Object_FromProto(mapCtx *direct.MapContext, in *pb.Kubernetes_Object) *krm.Kubernetes_Object {
	if in == nil {
		return nil
	}
	out := &krm.Kubernetes_Object{}
	out.Group = direct.LazyPtr(in.GetGroup())
	out.Kind = direct.LazyPtr(in.GetKind())
	out.Ns = direct.LazyPtr(in.GetNs())
	out.Name = direct.LazyPtr(in.GetName())
	out.Containers = direct.Slice_FromProto(mapCtx, in.Containers, Container_FromProto)
	return out
}
func Kubernetes_Object_ToProto(mapCtx *direct.MapContext, in *krm.Kubernetes_Object) *pb.Kubernetes_Object {
	if in == nil {
		return nil
	}
	out := &pb.Kubernetes_Object{}
	out.Group = direct.ValueOf(in.Group)
	out.Kind = direct.ValueOf(in.Kind)
	out.Ns = direct.ValueOf(in.Ns)
	out.Name = direct.ValueOf(in.Name)
	out.Containers = direct.Slice_ToProto(mapCtx, in.Containers, Container_ToProto)
	return out
}
func Kubernetes_Pod_FromProto(mapCtx *direct.MapContext, in *pb.Kubernetes_Pod) *krm.Kubernetes_Pod {
	if in == nil {
		return nil
	}
	out := &krm.Kubernetes_Pod{}
	out.Ns = direct.LazyPtr(in.GetNs())
	out.Name = direct.LazyPtr(in.GetName())
	out.Labels = direct.Slice_FromProto(mapCtx, in.Labels, Label_FromProto)
	out.Containers = direct.Slice_FromProto(mapCtx, in.Containers, Container_FromProto)
	return out
}
func Kubernetes_Pod_ToProto(mapCtx *direct.MapContext, in *krm.Kubernetes_Pod) *pb.Kubernetes_Pod {
	if in == nil {
		return nil
	}
	out := &pb.Kubernetes_Pod{}
	out.Ns = direct.ValueOf(in.Ns)
	out.Name = direct.ValueOf(in.Name)
	out.Labels = direct.Slice_ToProto(mapCtx, in.Labels, Label_ToProto)
	out.Containers = direct.Slice_ToProto(mapCtx, in.Containers, Container_ToProto)
	return out
}
func Kubernetes_Role_FromProto(mapCtx *direct.MapContext, in *pb.Kubernetes_Role) *krm.Kubernetes_Role {
	if in == nil {
		return nil
	}
	out := &krm.Kubernetes_Role{}
	out.Kind = direct.Enum_FromProto(mapCtx, in.GetKind())
	out.Ns = direct.LazyPtr(in.GetNs())
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func Kubernetes_Role_ToProto(mapCtx *direct.MapContext, in *krm.Kubernetes_Role) *pb.Kubernetes_Role {
	if in == nil {
		return nil
	}
	out := &pb.Kubernetes_Role{}
	out.Kind = direct.Enum_ToProto[pb.Kubernetes_Role_Kind](mapCtx, in.Kind)
	out.Ns = direct.ValueOf(in.Ns)
	out.Name = direct.ValueOf(in.Name)
	return out
}
func Kubernetes_Subject_FromProto(mapCtx *direct.MapContext, in *pb.Kubernetes_Subject) *krm.Kubernetes_Subject {
	if in == nil {
		return nil
	}
	out := &krm.Kubernetes_Subject{}
	out.Kind = direct.Enum_FromProto(mapCtx, in.GetKind())
	out.Ns = direct.LazyPtr(in.GetNs())
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func Kubernetes_Subject_ToProto(mapCtx *direct.MapContext, in *krm.Kubernetes_Subject) *pb.Kubernetes_Subject {
	if in == nil {
		return nil
	}
	out := &pb.Kubernetes_Subject{}
	out.Kind = direct.Enum_ToProto[pb.Kubernetes_Subject_AuthType](mapCtx, in.Kind)
	out.Ns = direct.ValueOf(in.Ns)
	out.Name = direct.ValueOf(in.Name)
	return out
}
func Label_FromProto(mapCtx *direct.MapContext, in *pb.Label) *krm.Label {
	if in == nil {
		return nil
	}
	out := &krm.Label{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func Label_ToProto(mapCtx *direct.MapContext, in *krm.Label) *pb.Label {
	if in == nil {
		return nil
	}
	out := &pb.Label{}
	out.Name = direct.ValueOf(in.Name)
	out.Value = direct.ValueOf(in.Value)
	return out
}
func LoadBalancer_FromProto(mapCtx *direct.MapContext, in *pb.LoadBalancer) *krm.LoadBalancer {
	if in == nil {
		return nil
	}
	out := &krm.LoadBalancer{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func LoadBalancer_ToProto(mapCtx *direct.MapContext, in *krm.LoadBalancer) *pb.LoadBalancer {
	if in == nil {
		return nil
	}
	out := &pb.LoadBalancer{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
func LogEntry_FromProto(mapCtx *direct.MapContext, in *pb.LogEntry) *krm.LogEntry {
	if in == nil {
		return nil
	}
	out := &krm.LogEntry{}
	out.CloudLoggingEntry = CloudLoggingEntry_FromProto(mapCtx, in.GetCloudLoggingEntry())
	return out
}
func LogEntry_ToProto(mapCtx *direct.MapContext, in *krm.LogEntry) *pb.LogEntry {
	if in == nil {
		return nil
	}
	out := &pb.LogEntry{}
	if oneof := CloudLoggingEntry_ToProto(mapCtx, in.CloudLoggingEntry); oneof != nil {
		out.LogEntry = &pb.LogEntry_CloudLoggingEntry{CloudLoggingEntry: oneof}
	}
	return out
}
func MitreAttack_FromProto(mapCtx *direct.MapContext, in *pb.MitreAttack) *krm.MitreAttack {
	if in == nil {
		return nil
	}
	out := &krm.MitreAttack{}
	out.PrimaryTactic = direct.Enum_FromProto(mapCtx, in.GetPrimaryTactic())
	out.PrimaryTechniques = direct.EnumSlice_FromProto(mapCtx, in.PrimaryTechniques)
	out.AdditionalTactics = direct.EnumSlice_FromProto(mapCtx, in.AdditionalTactics)
	out.AdditionalTechniques = direct.EnumSlice_FromProto(mapCtx, in.AdditionalTechniques)
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}
func MitreAttack_ToProto(mapCtx *direct.MapContext, in *krm.MitreAttack) *pb.MitreAttack {
	if in == nil {
		return nil
	}
	out := &pb.MitreAttack{}
	out.PrimaryTactic = direct.Enum_ToProto[pb.MitreAttack_Tactic](mapCtx, in.PrimaryTactic)
	out.PrimaryTechniques = direct.EnumSlice_ToProto[pb.MitreAttack_Technique](mapCtx, in.PrimaryTechniques)
	out.AdditionalTactics = direct.EnumSlice_ToProto[pb.MitreAttack_Tactic](mapCtx, in.AdditionalTactics)
	out.AdditionalTechniques = direct.EnumSlice_ToProto[pb.MitreAttack_Technique](mapCtx, in.AdditionalTechniques)
	out.Version = direct.ValueOf(in.Version)
	return out
}
func Notebook_FromProto(mapCtx *direct.MapContext, in *pb.Notebook) *krm.Notebook {
	if in == nil {
		return nil
	}
	out := &krm.Notebook{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Service = direct.LazyPtr(in.GetService())
	out.LastAuthor = direct.LazyPtr(in.GetLastAuthor())
	out.NotebookUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetNotebookUpdateTime())
	return out
}
func Notebook_ToProto(mapCtx *direct.MapContext, in *krm.Notebook) *pb.Notebook {
	if in == nil {
		return nil
	}
	out := &pb.Notebook{}
	out.Name = direct.ValueOf(in.Name)
	out.Service = direct.ValueOf(in.Service)
	out.LastAuthor = direct.ValueOf(in.LastAuthor)
	out.NotebookUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.NotebookUpdateTime)
	return out
}
func OrgPolicy_FromProto(mapCtx *direct.MapContext, in *pb.OrgPolicy) *krm.OrgPolicy {
	if in == nil {
		return nil
	}
	out := &krm.OrgPolicy{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func OrgPolicy_ToProto(mapCtx *direct.MapContext, in *krm.OrgPolicy) *pb.OrgPolicy {
	if in == nil {
		return nil
	}
	out := &pb.OrgPolicy{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
func Package_FromProto(mapCtx *direct.MapContext, in *pb.Package) *krm.Package {
	if in == nil {
		return nil
	}
	out := &krm.Package{}
	out.PackageName = direct.LazyPtr(in.GetPackageName())
	out.CpeURI = direct.LazyPtr(in.GetCpeUri())
	out.PackageType = direct.LazyPtr(in.GetPackageType())
	out.PackageVersion = direct.LazyPtr(in.GetPackageVersion())
	return out
}
func Package_ToProto(mapCtx *direct.MapContext, in *krm.Package) *pb.Package {
	if in == nil {
		return nil
	}
	out := &pb.Package{}
	out.PackageName = direct.ValueOf(in.PackageName)
	out.CpeUri = direct.ValueOf(in.CpeURI)
	out.PackageType = direct.ValueOf(in.PackageType)
	out.PackageVersion = direct.ValueOf(in.PackageVersion)
	return out
}
func Process_FromProto(mapCtx *direct.MapContext, in *pb.Process) *krm.Process {
	if in == nil {
		return nil
	}
	out := &krm.Process{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Binary = File_FromProto(mapCtx, in.GetBinary())
	out.Libraries = direct.Slice_FromProto(mapCtx, in.Libraries, File_FromProto)
	out.Script = File_FromProto(mapCtx, in.GetScript())
	out.Args = in.Args
	out.ArgumentsTruncated = direct.LazyPtr(in.GetArgumentsTruncated())
	out.EnvVariables = direct.Slice_FromProto(mapCtx, in.EnvVariables, EnvironmentVariable_FromProto)
	out.EnvVariablesTruncated = direct.LazyPtr(in.GetEnvVariablesTruncated())
	out.Pid = direct.LazyPtr(in.GetPid())
	out.ParentPid = direct.LazyPtr(in.GetParentPid())
	return out
}
func Process_ToProto(mapCtx *direct.MapContext, in *krm.Process) *pb.Process {
	if in == nil {
		return nil
	}
	out := &pb.Process{}
	out.Name = direct.ValueOf(in.Name)
	out.Binary = File_ToProto(mapCtx, in.Binary)
	out.Libraries = direct.Slice_ToProto(mapCtx, in.Libraries, File_ToProto)
	out.Script = File_ToProto(mapCtx, in.Script)
	out.Args = in.Args
	out.ArgumentsTruncated = direct.ValueOf(in.ArgumentsTruncated)
	out.EnvVariables = direct.Slice_ToProto(mapCtx, in.EnvVariables, EnvironmentVariable_ToProto)
	out.EnvVariablesTruncated = direct.ValueOf(in.EnvVariablesTruncated)
	out.Pid = direct.ValueOf(in.Pid)
	out.ParentPid = direct.ValueOf(in.ParentPid)
	return out
}
func Reference_FromProto(mapCtx *direct.MapContext, in *pb.Reference) *krm.Reference {
	if in == nil {
		return nil
	}
	out := &krm.Reference{}
	out.Source = direct.LazyPtr(in.GetSource())
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func Reference_ToProto(mapCtx *direct.MapContext, in *krm.Reference) *pb.Reference {
	if in == nil {
		return nil
	}
	out := &pb.Reference{}
	out.Source = direct.ValueOf(in.Source)
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func Requests_FromProto(mapCtx *direct.MapContext, in *pb.Requests) *krm.Requests {
	if in == nil {
		return nil
	}
	out := &krm.Requests{}
	out.Ratio = direct.LazyPtr(in.GetRatio())
	out.ShortTermAllowed = direct.LazyPtr(in.GetShortTermAllowed())
	out.LongTermAllowed = direct.LazyPtr(in.GetLongTermAllowed())
	out.LongTermDenied = direct.LazyPtr(in.GetLongTermDenied())
	return out
}
func Requests_ToProto(mapCtx *direct.MapContext, in *krm.Requests) *pb.Requests {
	if in == nil {
		return nil
	}
	out := &pb.Requests{}
	out.Ratio = direct.ValueOf(in.Ratio)
	out.ShortTermAllowed = direct.ValueOf(in.ShortTermAllowed)
	out.LongTermAllowed = direct.ValueOf(in.LongTermAllowed)
	out.LongTermDenied = direct.ValueOf(in.LongTermDenied)
	return out
}
func SecurityBulletin_FromProto(mapCtx *direct.MapContext, in *pb.SecurityBulletin) *krm.SecurityBulletin {
	if in == nil {
		return nil
	}
	out := &krm.SecurityBulletin{}
	out.BulletinID = direct.LazyPtr(in.GetBulletinId())
	out.SubmissionTime = direct.StringTimestamp_FromProto(mapCtx, in.GetSubmissionTime())
	out.SuggestedUpgradeVersion = direct.LazyPtr(in.GetSuggestedUpgradeVersion())
	return out
}
func SecurityBulletin_ToProto(mapCtx *direct.MapContext, in *krm.SecurityBulletin) *pb.SecurityBulletin {
	if in == nil {
		return nil
	}
	out := &pb.SecurityBulletin{}
	out.BulletinId = direct.ValueOf(in.BulletinID)
	out.SubmissionTime = direct.StringTimestamp_ToProto(mapCtx, in.SubmissionTime)
	out.SuggestedUpgradeVersion = direct.ValueOf(in.SuggestedUpgradeVersion)
	return out
}
func SecurityMarks_FromProto(mapCtx *direct.MapContext, in *pb.SecurityMarks) *krm.SecurityMarks {
	if in == nil {
		return nil
	}
	out := &krm.SecurityMarks{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Marks = in.Marks
	out.CanonicalName = direct.LazyPtr(in.GetCanonicalName())
	return out
}
func SecurityMarks_ToProto(mapCtx *direct.MapContext, in *krm.SecurityMarks) *pb.SecurityMarks {
	if in == nil {
		return nil
	}
	out := &pb.SecurityMarks{}
	out.Name = direct.ValueOf(in.Name)
	out.Marks = in.Marks
	out.CanonicalName = direct.ValueOf(in.CanonicalName)
	return out
}
func SecurityPolicy_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicy) *krm.SecurityPolicy {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPolicy{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Type = direct.LazyPtr(in.GetType())
	out.Preview = direct.LazyPtr(in.GetPreview())
	return out
}
func SecurityPolicy_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPolicy) *pb.SecurityPolicy {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicy{}
	out.Name = direct.ValueOf(in.Name)
	out.Type = direct.ValueOf(in.Type)
	out.Preview = direct.ValueOf(in.Preview)
	return out
}
func SecurityPosture_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPosture) *krm.SecurityPosture {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPosture{}
	out.Name = direct.LazyPtr(in.GetName())
	out.RevisionID = direct.LazyPtr(in.GetRevisionId())
	out.PostureDeploymentResource = direct.LazyPtr(in.GetPostureDeploymentResource())
	out.PostureDeployment = direct.LazyPtr(in.GetPostureDeployment())
	out.ChangedPolicy = direct.LazyPtr(in.GetChangedPolicy())
	out.PolicySet = direct.LazyPtr(in.GetPolicySet())
	out.Policy = direct.LazyPtr(in.GetPolicy())
	out.PolicyDriftDetails = direct.Slice_FromProto(mapCtx, in.PolicyDriftDetails, SecurityPosture_PolicyDriftDetails_FromProto)
	return out
}
func SecurityPosture_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPosture) *pb.SecurityPosture {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPosture{}
	out.Name = direct.ValueOf(in.Name)
	out.RevisionId = direct.ValueOf(in.RevisionID)
	out.PostureDeploymentResource = direct.ValueOf(in.PostureDeploymentResource)
	out.PostureDeployment = direct.ValueOf(in.PostureDeployment)
	out.ChangedPolicy = direct.ValueOf(in.ChangedPolicy)
	out.PolicySet = direct.ValueOf(in.PolicySet)
	out.Policy = direct.ValueOf(in.Policy)
	out.PolicyDriftDetails = direct.Slice_ToProto(mapCtx, in.PolicyDriftDetails, SecurityPosture_PolicyDriftDetails_ToProto)
	return out
}
func SecurityPosture_PolicyDriftDetails_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPosture_PolicyDriftDetails) *krm.SecurityPosture_PolicyDriftDetails {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPosture_PolicyDriftDetails{}
	out.Field = direct.LazyPtr(in.GetField())
	out.ExpectedValue = direct.LazyPtr(in.GetExpectedValue())
	out.DetectedValue = direct.LazyPtr(in.GetDetectedValue())
	return out
}
func SecurityPosture_PolicyDriftDetails_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPosture_PolicyDriftDetails) *pb.SecurityPosture_PolicyDriftDetails {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPosture_PolicyDriftDetails{}
	out.Field = direct.ValueOf(in.Field)
	out.ExpectedValue = direct.ValueOf(in.ExpectedValue)
	out.DetectedValue = direct.ValueOf(in.DetectedValue)
	return out
}
func SecuritycenterFindingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Finding) *krm.SecuritycenterFindingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterFindingObservedState{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ResourceName
	// MISSING: State
	// MISSING: Category
	// MISSING: ExternalURI
	// MISSING: SourceProperties
	// MISSING: SecurityMarks
	// MISSING: EventTime
	// MISSING: CreateTime
	// MISSING: Severity
	// MISSING: CanonicalName
	// MISSING: Mute
	// MISSING: FindingClass
	// MISSING: Indicator
	// MISSING: Vulnerability
	// MISSING: MuteUpdateTime
	// MISSING: ExternalSystems
	// MISSING: MitreAttack
	// MISSING: Access
	// MISSING: Connections
	// MISSING: MuteInitiator
	// MISSING: MuteInfo
	// MISSING: Processes
	// MISSING: Contacts
	// MISSING: Compliances
	// MISSING: ParentDisplayName
	// MISSING: Description
	// MISSING: Exfiltration
	// MISSING: IamBindings
	// MISSING: NextSteps
	// MISSING: ModuleName
	// MISSING: Containers
	// MISSING: Kubernetes
	// MISSING: Database
	// MISSING: AttackExposure
	// MISSING: Files
	// MISSING: CloudDlpInspection
	// MISSING: CloudDlpDataProfile
	// MISSING: KernelRootkit
	// MISSING: OrgPolicies
	// MISSING: Application
	// MISSING: BackupDisasterRecovery
	// MISSING: SecurityPosture
	// MISSING: LogEntries
	// MISSING: LoadBalancers
	// MISSING: CloudArmor
	// MISSING: Notebook
	// MISSING: ToxicCombination
	// MISSING: GroupMemberships
	return out
}
func SecuritycenterFindingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterFindingObservedState) *pb.Finding {
	if in == nil {
		return nil
	}
	out := &pb.Finding{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ResourceName
	// MISSING: State
	// MISSING: Category
	// MISSING: ExternalURI
	// MISSING: SourceProperties
	// MISSING: SecurityMarks
	// MISSING: EventTime
	// MISSING: CreateTime
	// MISSING: Severity
	// MISSING: CanonicalName
	// MISSING: Mute
	// MISSING: FindingClass
	// MISSING: Indicator
	// MISSING: Vulnerability
	// MISSING: MuteUpdateTime
	// MISSING: ExternalSystems
	// MISSING: MitreAttack
	// MISSING: Access
	// MISSING: Connections
	// MISSING: MuteInitiator
	// MISSING: MuteInfo
	// MISSING: Processes
	// MISSING: Contacts
	// MISSING: Compliances
	// MISSING: ParentDisplayName
	// MISSING: Description
	// MISSING: Exfiltration
	// MISSING: IamBindings
	// MISSING: NextSteps
	// MISSING: ModuleName
	// MISSING: Containers
	// MISSING: Kubernetes
	// MISSING: Database
	// MISSING: AttackExposure
	// MISSING: Files
	// MISSING: CloudDlpInspection
	// MISSING: CloudDlpDataProfile
	// MISSING: KernelRootkit
	// MISSING: OrgPolicies
	// MISSING: Application
	// MISSING: BackupDisasterRecovery
	// MISSING: SecurityPosture
	// MISSING: LogEntries
	// MISSING: LoadBalancers
	// MISSING: CloudArmor
	// MISSING: Notebook
	// MISSING: ToxicCombination
	// MISSING: GroupMemberships
	return out
}
func SecuritycenterFindingSpec_FromProto(mapCtx *direct.MapContext, in *pb.Finding) *krm.SecuritycenterFindingSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterFindingSpec{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ResourceName
	// MISSING: State
	// MISSING: Category
	// MISSING: ExternalURI
	// MISSING: SourceProperties
	// MISSING: SecurityMarks
	// MISSING: EventTime
	// MISSING: CreateTime
	// MISSING: Severity
	// MISSING: CanonicalName
	// MISSING: Mute
	// MISSING: FindingClass
	// MISSING: Indicator
	// MISSING: Vulnerability
	// MISSING: MuteUpdateTime
	// MISSING: ExternalSystems
	// MISSING: MitreAttack
	// MISSING: Access
	// MISSING: Connections
	// MISSING: MuteInitiator
	// MISSING: MuteInfo
	// MISSING: Processes
	// MISSING: Contacts
	// MISSING: Compliances
	// MISSING: ParentDisplayName
	// MISSING: Description
	// MISSING: Exfiltration
	// MISSING: IamBindings
	// MISSING: NextSteps
	// MISSING: ModuleName
	// MISSING: Containers
	// MISSING: Kubernetes
	// MISSING: Database
	// MISSING: AttackExposure
	// MISSING: Files
	// MISSING: CloudDlpInspection
	// MISSING: CloudDlpDataProfile
	// MISSING: KernelRootkit
	// MISSING: OrgPolicies
	// MISSING: Application
	// MISSING: BackupDisasterRecovery
	// MISSING: SecurityPosture
	// MISSING: LogEntries
	// MISSING: LoadBalancers
	// MISSING: CloudArmor
	// MISSING: Notebook
	// MISSING: ToxicCombination
	// MISSING: GroupMemberships
	return out
}
func SecuritycenterFindingSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterFindingSpec) *pb.Finding {
	if in == nil {
		return nil
	}
	out := &pb.Finding{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ResourceName
	// MISSING: State
	// MISSING: Category
	// MISSING: ExternalURI
	// MISSING: SourceProperties
	// MISSING: SecurityMarks
	// MISSING: EventTime
	// MISSING: CreateTime
	// MISSING: Severity
	// MISSING: CanonicalName
	// MISSING: Mute
	// MISSING: FindingClass
	// MISSING: Indicator
	// MISSING: Vulnerability
	// MISSING: MuteUpdateTime
	// MISSING: ExternalSystems
	// MISSING: MitreAttack
	// MISSING: Access
	// MISSING: Connections
	// MISSING: MuteInitiator
	// MISSING: MuteInfo
	// MISSING: Processes
	// MISSING: Contacts
	// MISSING: Compliances
	// MISSING: ParentDisplayName
	// MISSING: Description
	// MISSING: Exfiltration
	// MISSING: IamBindings
	// MISSING: NextSteps
	// MISSING: ModuleName
	// MISSING: Containers
	// MISSING: Kubernetes
	// MISSING: Database
	// MISSING: AttackExposure
	// MISSING: Files
	// MISSING: CloudDlpInspection
	// MISSING: CloudDlpDataProfile
	// MISSING: KernelRootkit
	// MISSING: OrgPolicies
	// MISSING: Application
	// MISSING: BackupDisasterRecovery
	// MISSING: SecurityPosture
	// MISSING: LogEntries
	// MISSING: LoadBalancers
	// MISSING: CloudArmor
	// MISSING: Notebook
	// MISSING: ToxicCombination
	// MISSING: GroupMemberships
	return out
}
func ServiceAccountDelegationInfo_FromProto(mapCtx *direct.MapContext, in *pb.ServiceAccountDelegationInfo) *krm.ServiceAccountDelegationInfo {
	if in == nil {
		return nil
	}
	out := &krm.ServiceAccountDelegationInfo{}
	out.PrincipalEmail = direct.LazyPtr(in.GetPrincipalEmail())
	out.PrincipalSubject = direct.LazyPtr(in.GetPrincipalSubject())
	return out
}
func ServiceAccountDelegationInfo_ToProto(mapCtx *direct.MapContext, in *krm.ServiceAccountDelegationInfo) *pb.ServiceAccountDelegationInfo {
	if in == nil {
		return nil
	}
	out := &pb.ServiceAccountDelegationInfo{}
	out.PrincipalEmail = direct.ValueOf(in.PrincipalEmail)
	out.PrincipalSubject = direct.ValueOf(in.PrincipalSubject)
	return out
}
func ToxicCombination_FromProto(mapCtx *direct.MapContext, in *pb.ToxicCombination) *krm.ToxicCombination {
	if in == nil {
		return nil
	}
	out := &krm.ToxicCombination{}
	out.AttackExposureScore = direct.LazyPtr(in.GetAttackExposureScore())
	out.RelatedFindings = in.RelatedFindings
	return out
}
func ToxicCombination_ToProto(mapCtx *direct.MapContext, in *krm.ToxicCombination) *pb.ToxicCombination {
	if in == nil {
		return nil
	}
	out := &pb.ToxicCombination{}
	out.AttackExposureScore = direct.ValueOf(in.AttackExposureScore)
	out.RelatedFindings = in.RelatedFindings
	return out
}
func Vulnerability_FromProto(mapCtx *direct.MapContext, in *pb.Vulnerability) *krm.Vulnerability {
	if in == nil {
		return nil
	}
	out := &krm.Vulnerability{}
	out.Cve = Cve_FromProto(mapCtx, in.GetCve())
	out.OffendingPackage = Package_FromProto(mapCtx, in.GetOffendingPackage())
	out.FixedPackage = Package_FromProto(mapCtx, in.GetFixedPackage())
	out.SecurityBulletin = SecurityBulletin_FromProto(mapCtx, in.GetSecurityBulletin())
	return out
}
func Vulnerability_ToProto(mapCtx *direct.MapContext, in *krm.Vulnerability) *pb.Vulnerability {
	if in == nil {
		return nil
	}
	out := &pb.Vulnerability{}
	out.Cve = Cve_ToProto(mapCtx, in.Cve)
	out.OffendingPackage = Package_ToProto(mapCtx, in.OffendingPackage)
	out.FixedPackage = Package_ToProto(mapCtx, in.FixedPackage)
	out.SecurityBulletin = SecurityBulletin_ToProto(mapCtx, in.SecurityBulletin)
	return out
}
