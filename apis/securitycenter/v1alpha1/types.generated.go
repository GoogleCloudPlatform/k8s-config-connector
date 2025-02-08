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

package v1alpha1


// +kcc:proto=google.cloud.securitycenter.v1.Access
type Access struct {
	// Associated email, such as "foo@google.com".
	//
	//  The email address of the authenticated user or a service account acting on
	//  behalf of a third party principal making the request. For third party
	//  identity callers, the `principal_subject` field is populated instead of
	//  this field. For privacy reasons, the principal email address is sometimes
	//  redacted. For more information, see [Caller identities in audit
	//  logs](https://cloud.google.com/logging/docs/audit#user-id).
	// +kcc:proto:field=google.cloud.securitycenter.v1.Access.principal_email
	PrincipalEmail *string `json:"principalEmail,omitempty"`

	// Caller's IP address, such as "1.1.1.1".
	// +kcc:proto:field=google.cloud.securitycenter.v1.Access.caller_ip
	CallerIP *string `json:"callerIP,omitempty"`

	// The caller IP's geolocation, which identifies where the call came from.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Access.caller_ip_geo
	CallerIPGeo *Geolocation `json:"callerIPGeo,omitempty"`

	// Type of user agent associated with the finding. For example, an operating
	//  system shell or an embedded or standalone application.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Access.user_agent_family
	UserAgentFamily *string `json:"userAgentFamily,omitempty"`

	// The caller's user agent string associated with the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Access.user_agent
	UserAgent *string `json:"userAgent,omitempty"`

	// This is the API service that the service account made a call to, e.g.
	//  "iam.googleapis.com"
	// +kcc:proto:field=google.cloud.securitycenter.v1.Access.service_name
	ServiceName *string `json:"serviceName,omitempty"`

	// The method that the service account called, e.g. "SetIamPolicy".
	// +kcc:proto:field=google.cloud.securitycenter.v1.Access.method_name
	MethodName *string `json:"methodName,omitempty"`

	// A string that represents the principal_subject that is associated with the
	//  identity. Unlike `principal_email`, `principal_subject` supports principals
	//  that aren't associated with email addresses, such as third party
	//  principals. For most identities, the format is
	//  `principal://iam.googleapis.com/{identity pool name}/subject/{subject}`.
	//  Some GKE identities, such as GKE_WORKLOAD, FREEFORM, and GKE_HUB_WORKLOAD,
	//  still use the legacy format `serviceAccount:{identity pool
	//  name}[{subject}]`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Access.principal_subject
	PrincipalSubject *string `json:"principalSubject,omitempty"`

	// The name of the service account key that was used to create or exchange
	//  credentials when authenticating the service account that made the request.
	//  This is a scheme-less URI full resource name. For example:
	//
	//  "//iam.googleapis.com/projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}/keys/{key}".
	// +kcc:proto:field=google.cloud.securitycenter.v1.Access.service_account_key_name
	ServiceAccountKeyName *string `json:"serviceAccountKeyName,omitempty"`

	// The identity delegation history of an authenticated service account that
	//  made the request. The `serviceAccountDelegationInfo[]` object contains
	//  information about the real authorities that try to access Google Cloud
	//  resources by delegating on a service account. When multiple authorities are
	//  present, they are guaranteed to be sorted based on the original ordering of
	//  the identity delegation events.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Access.service_account_delegation_info
	ServiceAccountDelegationInfo []ServiceAccountDelegationInfo `json:"serviceAccountDelegationInfo,omitempty"`

	// A string that represents a username. The username provided depends on the
	//  type of the finding and is likely not an IAM principal. For example, this
	//  can be a system username if the finding is related to a virtual machine, or
	//  it can be an application login username.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Access.user_name
	UserName *string `json:"userName,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.AdaptiveProtection
type AdaptiveProtection struct {
	// A score of 0 means that there is low confidence that the detected event is
	//  an actual attack. A score of 1 means that there is high confidence that the
	//  detected event is an attack. See the [Adaptive Protection
	//  documentation](https://cloud.google.com/armor/docs/adaptive-protection-overview#configure-alert-tuning)
	//  for further explanation.
	// +kcc:proto:field=google.cloud.securitycenter.v1.AdaptiveProtection.confidence
	Confidence *float64 `json:"confidence,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Application
type Application struct {
	// The base URI that identifies the network location of the application in
	//  which the vulnerability was detected. For example, `http://example.com`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Application.base_uri
	BaseURI *string `json:"baseURI,omitempty"`

	// The full URI with payload that can be used to reproduce the
	//  vulnerability. For example, `http://example.com?p=aMmYgI6H`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Application.full_uri
	FullURI *string `json:"fullURI,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Attack
type Attack struct {
	// Total PPS (packets per second) volume of attack.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Attack.volume_pps
	VolumePps *int32 `json:"volumePps,omitempty"`

	// Total BPS (bytes per second) volume of attack.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Attack.volume_bps
	VolumeBps *int32 `json:"volumeBps,omitempty"`

	// Type of attack, for example, 'SYN-flood', 'NTP-udp', or 'CHARGEN-udp'.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Attack.classification
	Classification *string `json:"classification,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.AttackExposure
type AttackExposure struct {
	// A number between 0 (inclusive) and infinity that represents how important
	//  this finding is to remediate. The higher the score, the more important it
	//  is to remediate.
	// +kcc:proto:field=google.cloud.securitycenter.v1.AttackExposure.score
	Score *float64 `json:"score,omitempty"`

	// The most recent time the attack exposure was updated on this finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.AttackExposure.latest_calculation_time
	LatestCalculationTime *string `json:"latestCalculationTime,omitempty"`

	// The resource name of the attack path simulation result that contains the
	//  details regarding this attack exposure score.
	//  Example: `organizations/123/simulations/456/attackExposureResults/789`
	// +kcc:proto:field=google.cloud.securitycenter.v1.AttackExposure.attack_exposure_result
	AttackExposureResult *string `json:"attackExposureResult,omitempty"`

	// What state this AttackExposure is in. This captures whether or not an
	//  attack exposure has been calculated or not.
	// +kcc:proto:field=google.cloud.securitycenter.v1.AttackExposure.state
	State *string `json:"state,omitempty"`

	// The number of high value resources that are exposed as a result of this
	//  finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.AttackExposure.exposed_high_value_resources_count
	ExposedHighValueResourcesCount *int32 `json:"exposedHighValueResourcesCount,omitempty"`

	// The number of medium value resources that are exposed as a result of this
	//  finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.AttackExposure.exposed_medium_value_resources_count
	ExposedMediumValueResourcesCount *int32 `json:"exposedMediumValueResourcesCount,omitempty"`

	// The number of high value resources that are exposed as a result of this
	//  finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.AttackExposure.exposed_low_value_resources_count
	ExposedLowValueResourcesCount *int32 `json:"exposedLowValueResourcesCount,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.BackupDisasterRecovery
type BackupDisasterRecovery struct {
	// The name of a Backup and DR template which comprises one or more backup
	//  policies. See the [Backup and DR
	//  documentation](https://cloud.google.com/backup-disaster-recovery/docs/concepts/backup-plan#temp)
	//  for more information. For example, `snap-ov`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.BackupDisasterRecovery.backup_template
	BackupTemplate *string `json:"backupTemplate,omitempty"`

	// The names of Backup and DR policies that are associated with a template
	//  and that define when to run a backup, how frequently to run a backup, and
	//  how long to retain the backup image. For example, `onvaults`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.BackupDisasterRecovery.policies
	Policies []string `json:"policies,omitempty"`

	// The name of a Backup and DR host, which is managed by the backup and
	//  recovery appliance and known to the management console. The host can be of
	//  type Generic (for example, Compute Engine, SQL Server, Oracle DB, SMB file
	//  system, etc.), vCenter, or an ESX server. See the [Backup and DR
	//  documentation on
	//  hosts](https://cloud.google.com/backup-disaster-recovery/docs/configuration/manage-hosts-and-their-applications)
	//  for more information. For example, `centos7-01`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.BackupDisasterRecovery.host
	Host *string `json:"host,omitempty"`

	// The names of Backup and DR applications. An application is a VM, database,
	//  or file system on a managed host monitored by a backup and recovery
	//  appliance. For example, `centos7-01-vol00`, `centos7-01-vol01`,
	//  `centos7-01-vol02`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.BackupDisasterRecovery.applications
	Applications []string `json:"applications,omitempty"`

	// The name of the Backup and DR storage pool that the backup and recovery
	//  appliance is storing data in. The storage pool could be of type Cloud,
	//  Primary, Snapshot, or OnVault. See the [Backup and DR documentation on
	//  storage
	//  pools](https://cloud.google.com/backup-disaster-recovery/docs/concepts/storage-pools).
	//  For example, `DiskPoolOne`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.BackupDisasterRecovery.storage_pool
	StoragePool *string `json:"storagePool,omitempty"`

	// The names of Backup and DR advanced policy options of a policy applying to
	//  an application. See the [Backup and DR documentation on policy
	//  options](https://cloud.google.com/backup-disaster-recovery/docs/create-plan/policy-settings).
	//  For example, `skipofflineappsincongrp, nounmap`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.BackupDisasterRecovery.policy_options
	PolicyOptions []string `json:"policyOptions,omitempty"`

	// The name of the Backup and DR resource profile that specifies the storage
	//  media for backups of application and VM data. See the [Backup and DR
	//  documentation on
	//  profiles](https://cloud.google.com/backup-disaster-recovery/docs/concepts/backup-plan#profile).
	//  For example, `GCP`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.BackupDisasterRecovery.profile
	Profile *string `json:"profile,omitempty"`

	// The name of the Backup and DR appliance that captures, moves, and manages
	//  the lifecycle of backup data. For example, `backup-server-57137`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.BackupDisasterRecovery.appliance
	Appliance *string `json:"appliance,omitempty"`

	// The backup type of the Backup and DR image.
	//  For example, `Snapshot`, `Remote Snapshot`, `OnVault`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.BackupDisasterRecovery.backup_type
	BackupType *string `json:"backupType,omitempty"`

	// The timestamp at which the Backup and DR backup was created.
	// +kcc:proto:field=google.cloud.securitycenter.v1.BackupDisasterRecovery.backup_create_time
	BackupCreateTime *string `json:"backupCreateTime,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.CloudArmor
type CloudArmor struct {
	// Information about the [Google Cloud Armor security
	//  policy](https://cloud.google.com/armor/docs/security-policy-overview)
	//  relevant to the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.CloudArmor.security_policy
	SecurityPolicy *SecurityPolicy `json:"securityPolicy,omitempty"`

	// Information about incoming requests evaluated by [Google Cloud Armor
	//  security
	//  policies](https://cloud.google.com/armor/docs/security-policy-overview).
	// +kcc:proto:field=google.cloud.securitycenter.v1.CloudArmor.requests
	Requests *Requests `json:"requests,omitempty"`

	// Information about potential Layer 7 DDoS attacks identified by [Google
	//  Cloud Armor Adaptive
	//  Protection](https://cloud.google.com/armor/docs/adaptive-protection-overview).
	// +kcc:proto:field=google.cloud.securitycenter.v1.CloudArmor.adaptive_protection
	AdaptiveProtection *AdaptiveProtection `json:"adaptiveProtection,omitempty"`

	// Information about DDoS attack volume and classification.
	// +kcc:proto:field=google.cloud.securitycenter.v1.CloudArmor.attack
	Attack *Attack `json:"attack,omitempty"`

	// Distinguish between volumetric & protocol DDoS attack and
	//  application layer attacks. For example, "L3_4" for Layer 3 and Layer 4 DDoS
	//  attacks, or "L_7" for Layer 7 DDoS attacks.
	// +kcc:proto:field=google.cloud.securitycenter.v1.CloudArmor.threat_vector
	ThreatVector *string `json:"threatVector,omitempty"`

	// Duration of attack from the start until the current moment (updated every 5
	//  minutes).
	// +kcc:proto:field=google.cloud.securitycenter.v1.CloudArmor.duration
	Duration *string `json:"duration,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.CloudDlpDataProfile
type CloudDlpDataProfile struct {
	// Name of the data profile, for example,
	//  `projects/123/locations/europe/tableProfiles/8383929`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.CloudDlpDataProfile.data_profile
	DataProfile *string `json:"dataProfile,omitempty"`

	// The resource hierarchy level at which the data profile was generated.
	// +kcc:proto:field=google.cloud.securitycenter.v1.CloudDlpDataProfile.parent_type
	ParentType *string `json:"parentType,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.CloudDlpInspection
type CloudDlpInspection struct {
	// Name of the inspection job, for example,
	//  `projects/123/locations/europe/dlpJobs/i-8383929`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.CloudDlpInspection.inspect_job
	InspectJob *string `json:"inspectJob,omitempty"`

	// The type of information (or
	//  *[infoType](https://cloud.google.com/dlp/docs/infotypes-reference)*) found,
	//  for example, `EMAIL_ADDRESS` or `STREET_ADDRESS`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.CloudDlpInspection.info_type
	InfoType *string `json:"infoType,omitempty"`

	// The number of times Cloud DLP found this infoType within this job
	//  and resource.
	// +kcc:proto:field=google.cloud.securitycenter.v1.CloudDlpInspection.info_type_count
	InfoTypeCount *int64 `json:"infoTypeCount,omitempty"`

	// Whether Cloud DLP scanned the complete resource or a sampled subset.
	// +kcc:proto:field=google.cloud.securitycenter.v1.CloudDlpInspection.full_scan
	FullScan *bool `json:"fullScan,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.CloudLoggingEntry
type CloudLoggingEntry struct {
	// A unique identifier for the log entry.
	// +kcc:proto:field=google.cloud.securitycenter.v1.CloudLoggingEntry.insert_id
	InsertID *string `json:"insertID,omitempty"`

	// The type of the log (part of `log_name`. `log_name` is the resource name of
	//  the log to which this log entry belongs). For example:
	//  `cloudresourcemanager.googleapis.com/activity`. Note that this field is not
	//  URL-encoded, unlike the `LOG_ID` field in `LogEntry`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.CloudLoggingEntry.log_id
	LogID *string `json:"logID,omitempty"`

	// The organization, folder, or project of the monitored resource that
	//  produced this log entry.
	// +kcc:proto:field=google.cloud.securitycenter.v1.CloudLoggingEntry.resource_container
	ResourceContainer *string `json:"resourceContainer,omitempty"`

	// The time the event described by the log entry occurred.
	// +kcc:proto:field=google.cloud.securitycenter.v1.CloudLoggingEntry.timestamp
	Timestamp *string `json:"timestamp,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Compliance
type Compliance struct {
	// Industry-wide compliance standards or benchmarks, such as CIS, PCI, and
	//  OWASP.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Compliance.standard
	Standard *string `json:"standard,omitempty"`

	// Version of the standard or benchmark, for example, 1.1
	// +kcc:proto:field=google.cloud.securitycenter.v1.Compliance.version
	Version *string `json:"version,omitempty"`

	// Policies within the standard or benchmark, for example, A.12.4.1
	// +kcc:proto:field=google.cloud.securitycenter.v1.Compliance.ids
	Ids []string `json:"ids,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Connection
type Connection struct {
	// Destination IP address. Not present for sockets that are listening and not
	//  connected.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Connection.destination_ip
	DestinationIP *string `json:"destinationIP,omitempty"`

	// Destination port. Not present for sockets that are listening and not
	//  connected.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Connection.destination_port
	DestinationPort *int32 `json:"destinationPort,omitempty"`

	// Source IP address.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Connection.source_ip
	SourceIP *string `json:"sourceIP,omitempty"`

	// Source port.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Connection.source_port
	SourcePort *int32 `json:"sourcePort,omitempty"`

	// IANA Internet Protocol Number such as TCP(6) and UDP(17).
	// +kcc:proto:field=google.cloud.securitycenter.v1.Connection.protocol
	Protocol *string `json:"protocol,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Contact
type Contact struct {
	// An email address. For example, "`person123@company.com`".
	// +kcc:proto:field=google.cloud.securitycenter.v1.Contact.email
	Email *string `json:"email,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.ContactDetails
type ContactDetails struct {
	// A list of contacts
	// +kcc:proto:field=google.cloud.securitycenter.v1.ContactDetails.contacts
	Contacts []Contact `json:"contacts,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Container
type Container struct {
	// Name of the container.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Container.name
	Name *string `json:"name,omitempty"`

	// Container image URI provided when configuring a pod or container. This
	//  string can identify a container image version using mutable tags.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Container.uri
	URI *string `json:"uri,omitempty"`

	// Optional container image ID, if provided by the container runtime. Uniquely
	//  identifies the container image launched using a container image digest.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Container.image_id
	ImageID *string `json:"imageID,omitempty"`

	// Container labels, as provided by the container runtime.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Container.labels
	Labels []Label `json:"labels,omitempty"`

	// The time that the container was created.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Container.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Cve
type Cve struct {
	// The unique identifier for the vulnerability. e.g. CVE-2021-34527
	// +kcc:proto:field=google.cloud.securitycenter.v1.Cve.id
	ID *string `json:"id,omitempty"`

	// Additional information about the CVE.
	//  e.g. https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2021-34527
	// +kcc:proto:field=google.cloud.securitycenter.v1.Cve.references
	References []Reference `json:"references,omitempty"`

	// Describe Common Vulnerability Scoring System specified at
	//  https://www.first.org/cvss/v3.1/specification-document
	// +kcc:proto:field=google.cloud.securitycenter.v1.Cve.cvssv3
	Cvssv3 *Cvssv3 `json:"cvssv3,omitempty"`

	// Whether upstream fix is available for the CVE.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Cve.upstream_fix_available
	UpstreamFixAvailable *bool `json:"upstreamFixAvailable,omitempty"`

	// The potential impact of the vulnerability if it was to be exploited.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Cve.impact
	Impact *string `json:"impact,omitempty"`

	// The exploitation activity of the vulnerability in the wild.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Cve.exploitation_activity
	ExploitationActivity *string `json:"exploitationActivity,omitempty"`

	// Whether or not the vulnerability has been observed in the wild.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Cve.observed_in_the_wild
	ObservedInTheWild *bool `json:"observedInTheWild,omitempty"`

	// Whether or not the vulnerability was zero day when the finding was
	//  published.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Cve.zero_day
	ZeroDay *bool `json:"zeroDay,omitempty"`

	// Date the first publicly available exploit or PoC was released.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Cve.exploit_release_date
	ExploitReleaseDate *string `json:"exploitReleaseDate,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Cvssv3
type Cvssv3 struct {
	// The base score is a function of the base metric scores.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Cvssv3.base_score
	BaseScore *float64 `json:"baseScore,omitempty"`

	// Base Metrics
	//  Represents the intrinsic characteristics of a vulnerability that are
	//  constant over time and across user environments.
	//  This metric reflects the context by which vulnerability exploitation is
	//  possible.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Cvssv3.attack_vector
	AttackVector *string `json:"attackVector,omitempty"`

	// This metric describes the conditions beyond the attacker's control that
	//  must exist in order to exploit the vulnerability.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Cvssv3.attack_complexity
	AttackComplexity *string `json:"attackComplexity,omitempty"`

	// This metric describes the level of privileges an attacker must possess
	//  before successfully exploiting the vulnerability.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Cvssv3.privileges_required
	PrivilegesRequired *string `json:"privilegesRequired,omitempty"`

	// This metric captures the requirement for a human user, other than the
	//  attacker, to participate in the successful compromise of the vulnerable
	//  component.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Cvssv3.user_interaction
	UserInteraction *string `json:"userInteraction,omitempty"`

	// The Scope metric captures whether a vulnerability in one vulnerable
	//  component impacts resources in components beyond its security scope.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Cvssv3.scope
	Scope *string `json:"scope,omitempty"`

	// This metric measures the impact to the confidentiality of the information
	//  resources managed by a software component due to a successfully exploited
	//  vulnerability.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Cvssv3.confidentiality_impact
	ConfidentialityImpact *string `json:"confidentialityImpact,omitempty"`

	// This metric measures the impact to integrity of a successfully exploited
	//  vulnerability.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Cvssv3.integrity_impact
	IntegrityImpact *string `json:"integrityImpact,omitempty"`

	// This metric measures the impact to the availability of the impacted
	//  component resulting from a successfully exploited vulnerability.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Cvssv3.availability_impact
	AvailabilityImpact *string `json:"availabilityImpact,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Database
type Database struct {
	// Some database resources may not have the [full resource
	//  name](https://google.aip.dev/122#full-resource-names) populated because
	//  these resource types are not yet supported by Cloud Asset Inventory (e.g.
	//  Cloud SQL databases). In these cases only the display name will be
	//  provided.
	//  The [full resource name](https://google.aip.dev/122#full-resource-names) of
	//  the database that the user connected to, if it is supported by Cloud Asset
	//  Inventory.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Database.name
	Name *string `json:"name,omitempty"`

	// The human-readable name of the database that the user connected to.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Database.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The username used to connect to the database. The username might not be an
	//  IAM principal and does not have a set format.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Database.user_name
	UserName *string `json:"userName,omitempty"`

	// The SQL statement that is associated with the database access.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Database.query
	Query *string `json:"query,omitempty"`

	// The target usernames, roles, or groups of an SQL privilege grant, which is
	//  not an IAM policy change.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Database.grantees
	Grantees []string `json:"grantees,omitempty"`

	// The version of the database, for example, POSTGRES_14.
	//  See [the complete
	//  list](https://cloud.google.com/sql/docs/mysql/admin-api/rest/v1/SqlDatabaseVersion).
	// +kcc:proto:field=google.cloud.securitycenter.v1.Database.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.EnvironmentVariable
type EnvironmentVariable struct {
	// Environment variable name as a JSON encoded string.
	// +kcc:proto:field=google.cloud.securitycenter.v1.EnvironmentVariable.name
	Name *string `json:"name,omitempty"`

	// Environment variable value as a JSON encoded string.
	// +kcc:proto:field=google.cloud.securitycenter.v1.EnvironmentVariable.val
	Val *string `json:"val,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.ExfilResource
type ExfilResource struct {
	// The resource's [full resource
	//  name](https://cloud.google.com/apis/design/resource_names#full_resource_name).
	// +kcc:proto:field=google.cloud.securitycenter.v1.ExfilResource.name
	Name *string `json:"name,omitempty"`

	// Subcomponents of the asset that was exfiltrated, like URIs used during
	//  exfiltration, table names, databases, and filenames. For example, multiple
	//  tables might have been exfiltrated from the same Cloud SQL instance, or
	//  multiple files might have been exfiltrated from the same Cloud Storage
	//  bucket.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ExfilResource.components
	Components []string `json:"components,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Exfiltration
type Exfiltration struct {
	// If there are multiple sources, then the data is considered "joined" between
	//  them. For instance, BigQuery can join multiple tables, and each
	//  table would be considered a source.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Exfiltration.sources
	Sources []ExfilResource `json:"sources,omitempty"`

	// If there are multiple targets, each target would get a complete copy of the
	//  "joined" source data.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Exfiltration.targets
	Targets []ExfilResource `json:"targets,omitempty"`

	// Total exfiltrated bytes processed for the entire job.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Exfiltration.total_exfiltrated_bytes
	TotalExfiltratedBytes *int64 `json:"totalExfiltratedBytes,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.ExternalSystem
type ExternalSystem struct {
	// Full resource name of the external system, for example:
	//  "organizations/1234/sources/5678/findings/123456/externalSystems/jira",
	//  "folders/1234/sources/5678/findings/123456/externalSystems/jira",
	//  "projects/1234/sources/5678/findings/123456/externalSystems/jira"
	// +kcc:proto:field=google.cloud.securitycenter.v1.ExternalSystem.name
	Name *string `json:"name,omitempty"`

	// References primary/secondary etc assignees in the external system.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ExternalSystem.assignees
	Assignees []string `json:"assignees,omitempty"`

	// The identifier that's used to track the finding's corresponding case in the
	//  external system.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ExternalSystem.external_uid
	ExternalUid *string `json:"externalUid,omitempty"`

	// The most recent status of the finding's corresponding case, as reported by
	//  the external system.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ExternalSystem.status
	Status *string `json:"status,omitempty"`

	// The time when the case was last updated, as reported by the external
	//  system.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ExternalSystem.external_system_update_time
	ExternalSystemUpdateTime *string `json:"externalSystemUpdateTime,omitempty"`

	// The link to the finding's corresponding case in the external system.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ExternalSystem.case_uri
	CaseURI *string `json:"caseURI,omitempty"`

	// The priority of the finding's corresponding case in the external system.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ExternalSystem.case_priority
	CasePriority *string `json:"casePriority,omitempty"`

	// The SLA of the finding's corresponding case in the external system.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ExternalSystem.case_sla
	CaseSla *string `json:"caseSla,omitempty"`

	// The time when the case was created, as reported by the external system.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ExternalSystem.case_create_time
	CaseCreateTime *string `json:"caseCreateTime,omitempty"`

	// The time when the case was closed, as reported by the external system.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ExternalSystem.case_close_time
	CaseCloseTime *string `json:"caseCloseTime,omitempty"`

	// Information about the ticket, if any, that is being used to track the
	//  resolution of the issue that is identified by this finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ExternalSystem.ticket_info
	TicketInfo *ExternalSystem_TicketInfo `json:"ticketInfo,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.ExternalSystem.TicketInfo
type ExternalSystem_TicketInfo struct {
	// The identifier of the ticket in the ticket system.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ExternalSystem.TicketInfo.id
	ID *string `json:"id,omitempty"`

	// The assignee of the ticket in the ticket system.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ExternalSystem.TicketInfo.assignee
	Assignee *string `json:"assignee,omitempty"`

	// The description of the ticket in the ticket system.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ExternalSystem.TicketInfo.description
	Description *string `json:"description,omitempty"`

	// The link to the ticket in the ticket system.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ExternalSystem.TicketInfo.uri
	URI *string `json:"uri,omitempty"`

	// The latest status of the ticket, as reported by the ticket system.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ExternalSystem.TicketInfo.status
	Status *string `json:"status,omitempty"`

	// The time when the ticket was last updated, as reported by the ticket
	//  system.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ExternalSystem.TicketInfo.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.File
type File struct {
	// Absolute path of the file as a JSON encoded string.
	// +kcc:proto:field=google.cloud.securitycenter.v1.File.path
	Path *string `json:"path,omitempty"`

	// Size of the file in bytes.
	// +kcc:proto:field=google.cloud.securitycenter.v1.File.size
	Size *int64 `json:"size,omitempty"`

	// SHA256 hash of the first hashed_size bytes of the file encoded as a
	//  hex string.  If hashed_size == size, sha256 represents the SHA256 hash
	//  of the entire file.
	// +kcc:proto:field=google.cloud.securitycenter.v1.File.sha256
	Sha256 *string `json:"sha256,omitempty"`

	// The length in bytes of the file prefix that was hashed.  If
	//  hashed_size == size, any hashes reported represent the entire
	//  file.
	// +kcc:proto:field=google.cloud.securitycenter.v1.File.hashed_size
	HashedSize *int64 `json:"hashedSize,omitempty"`

	// True when the hash covers only a prefix of the file.
	// +kcc:proto:field=google.cloud.securitycenter.v1.File.partially_hashed
	PartiallyHashed *bool `json:"partiallyHashed,omitempty"`

	// Prefix of the file contents as a JSON-encoded string.
	// +kcc:proto:field=google.cloud.securitycenter.v1.File.contents
	Contents *string `json:"contents,omitempty"`

	// Path of the file in terms of underlying disk/partition identifiers.
	// +kcc:proto:field=google.cloud.securitycenter.v1.File.disk_path
	DiskPath *File_DiskPath `json:"diskPath,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.File.DiskPath
type File_DiskPath struct {
	// UUID of the partition (format
	//  https://wiki.archlinux.org/title/persistent_block_device_naming#by-uuid)
	// +kcc:proto:field=google.cloud.securitycenter.v1.File.DiskPath.partition_uuid
	PartitionUuid *string `json:"partitionUuid,omitempty"`

	// Relative path of the file in the partition as a JSON encoded string.
	//  Example: /home/user1/executable_file.sh
	// +kcc:proto:field=google.cloud.securitycenter.v1.File.DiskPath.relative_path
	RelativePath *string `json:"relativePath,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Finding
type Finding struct {
	// The [relative resource
	//  name](https://cloud.google.com/apis/design/resource_names#relative_resource_name)
	//  of the finding. Example:
	//  "organizations/{organization_id}/sources/{source_id}/findings/{finding_id}",
	//  "folders/{folder_id}/sources/{source_id}/findings/{finding_id}",
	//  "projects/{project_id}/sources/{source_id}/findings/{finding_id}".
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.name
	Name *string `json:"name,omitempty"`

	// The relative resource name of the source the finding belongs to. See:
	//  https://cloud.google.com/apis/design/resource_names#relative_resource_name
	//  This field is immutable after creation time.
	//  For example:
	//  "organizations/{organization_id}/sources/{source_id}"
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.parent
	Parent *string `json:"parent,omitempty"`

	// For findings on Google Cloud resources, the full resource
	//  name of the Google Cloud resource this finding is for. See:
	//  https://cloud.google.com/apis/design/resource_names#full_resource_name
	//  When the finding is for a non-Google Cloud resource, the resourceName can
	//  be a customer or partner defined string. This field is immutable after
	//  creation time.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.resource_name
	ResourceName *string `json:"resourceName,omitempty"`

	// The state of the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.state
	State *string `json:"state,omitempty"`

	// The additional taxonomy group within findings from a given source.
	//  This field is immutable after creation time.
	//  Example: "XSS_FLASH_INJECTION"
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.category
	Category *string `json:"category,omitempty"`

	// The URI that, if available, points to a web page outside of Security
	//  Command Center where additional information about the finding can be found.
	//  This field is guaranteed to be either empty or a well formed URL.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.external_uri
	ExternalURI *string `json:"externalURI,omitempty"`

	// TODO: unsupported map type with key string and value message


	// The time the finding was first detected. If an existing finding is updated,
	//  then this is the time the update occurred.
	//  For example, if the finding represents an open firewall, this property
	//  captures the time the detector believes the firewall became open. The
	//  accuracy is determined by the detector. If the finding is later resolved,
	//  then this time reflects when the finding was resolved. This must not
	//  be set to a value greater than the current timestamp.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.event_time
	EventTime *string `json:"eventTime,omitempty"`

	// The time at which the finding was created in Security Command Center.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// The severity of the finding. This field is managed by the source that
	//  writes the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.severity
	Severity *string `json:"severity,omitempty"`

	// The canonical name of the finding. It's either
	//  "organizations/{organization_id}/sources/{source_id}/findings/{finding_id}",
	//  "folders/{folder_id}/sources/{source_id}/findings/{finding_id}" or
	//  "projects/{project_number}/sources/{source_id}/findings/{finding_id}",
	//  depending on the closest CRM ancestor of the resource associated with the
	//  finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.canonical_name
	CanonicalName *string `json:"canonicalName,omitempty"`

	// Indicates the mute state of a finding (either muted, unmuted
	//  or undefined). Unlike other attributes of a finding, a finding provider
	//  shouldn't set the value of mute.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.mute
	Mute *string `json:"mute,omitempty"`

	// The class of the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.finding_class
	FindingClass *string `json:"findingClass,omitempty"`

	// Represents what's commonly known as an *indicator of compromise* (IoC) in
	//  computer forensics. This is an artifact observed on a network or in an
	//  operating system that, with high confidence, indicates a computer
	//  intrusion. For more information, see [Indicator of
	//  compromise](https://en.wikipedia.org/wiki/Indicator_of_compromise).
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.indicator
	Indicator *Indicator `json:"indicator,omitempty"`

	// Represents vulnerability-specific fields like CVE and CVSS scores.
	//  CVE stands for Common Vulnerabilities and Exposures
	//  (https://cve.mitre.org/about/)
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.vulnerability
	Vulnerability *Vulnerability `json:"vulnerability,omitempty"`

	// MITRE ATT&CK tactics and techniques related to this finding.
	//  See: https://attack.mitre.org
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.mitre_attack
	MitreAttack *MitreAttack `json:"mitreAttack,omitempty"`

	// Access details associated with the finding, such as more information on the
	//  caller, which method was accessed, and from where.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.access
	Access *Access `json:"access,omitempty"`

	// Contains information about the IP connection associated with the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.connections
	Connections []Connection `json:"connections,omitempty"`

	// Records additional information about the mute operation, for example, the
	//  [mute configuration](/security-command-center/docs/how-to-mute-findings)
	//  that muted the finding and the user who muted the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.mute_initiator
	MuteInitiator *string `json:"muteInitiator,omitempty"`

	// Represents operating system processes associated with the Finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.processes
	Processes []Process `json:"processes,omitempty"`

	// Contains compliance information for security standards associated to the
	//  finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.compliances
	Compliances []Compliance `json:"compliances,omitempty"`

	// Contains more details about the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.description
	Description *string `json:"description,omitempty"`

	// Represents exfiltrations associated with the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.exfiltration
	Exfiltration *Exfiltration `json:"exfiltration,omitempty"`

	// Represents IAM bindings associated with the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.iam_bindings
	IamBindings []IamBinding `json:"iamBindings,omitempty"`

	// Steps to address the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.next_steps
	NextSteps *string `json:"nextSteps,omitempty"`

	// Unique identifier of the module which generated the finding.
	//  Example:
	//  folders/598186756061/securityHealthAnalyticsSettings/customModules/56799441161885
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.module_name
	ModuleName *string `json:"moduleName,omitempty"`

	// Containers associated with the finding. This field provides information for
	//  both Kubernetes and non-Kubernetes containers.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.containers
	Containers []Container `json:"containers,omitempty"`

	// Kubernetes resources associated with the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.kubernetes
	Kubernetes *Kubernetes `json:"kubernetes,omitempty"`

	// Database associated with the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.database
	Database *Database `json:"database,omitempty"`

	// The results of an attack path simulation relevant to this finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.attack_exposure
	AttackExposure *AttackExposure `json:"attackExposure,omitempty"`

	// File associated with the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.files
	Files []File `json:"files,omitempty"`

	// Cloud Data Loss Prevention (Cloud DLP) inspection results that are
	//  associated with the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.cloud_dlp_inspection
	CloudDlpInspection *CloudDlpInspection `json:"cloudDlpInspection,omitempty"`

	// Cloud DLP data profile that is associated with the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.cloud_dlp_data_profile
	CloudDlpDataProfile *CloudDlpDataProfile `json:"cloudDlpDataProfile,omitempty"`

	// Signature of the kernel rootkit.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.kernel_rootkit
	KernelRootkit *KernelRootkit `json:"kernelRootkit,omitempty"`

	// Contains information about the org policies associated with the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.org_policies
	OrgPolicies []OrgPolicy `json:"orgPolicies,omitempty"`

	// Represents an application associated with the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.application
	Application *Application `json:"application,omitempty"`

	// Fields related to Backup and DR findings.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.backup_disaster_recovery
	BackupDisasterRecovery *BackupDisasterRecovery `json:"backupDisasterRecovery,omitempty"`

	// The security posture associated with the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.security_posture
	SecurityPosture *SecurityPosture `json:"securityPosture,omitempty"`

	// Log entries that are relevant to the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.log_entries
	LogEntries []LogEntry `json:"logEntries,omitempty"`

	// The load balancers associated with the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.load_balancers
	LoadBalancers []LoadBalancer `json:"loadBalancers,omitempty"`

	// Fields related to Cloud Armor findings.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.cloud_armor
	CloudArmor *CloudArmor `json:"cloudArmor,omitempty"`

	// Notebook associated with the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.notebook
	Notebook *Notebook `json:"notebook,omitempty"`

	// Contains details about a group of security issues that, when the issues
	//  occur together, represent a greater risk than when the issues occur
	//  independently. A group of such issues is referred to as a toxic
	//  combination.
	//  This field cannot be updated. Its value is ignored in all update requests.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.toxic_combination
	ToxicCombination *ToxicCombination `json:"toxicCombination,omitempty"`

	// Contains details about groups of which this finding is a member. A group is
	//  a collection of findings that are related in some way.
	//  This field cannot be updated. Its value is ignored in all update requests.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.group_memberships
	GroupMemberships []GroupMembership `json:"groupMemberships,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Finding.MuteInfo
type Finding_MuteInfo struct {
	// If set, the static mute applied to this finding. Static mutes override
	//  dynamic mutes. If unset, there is no static mute.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.MuteInfo.static_mute
	StaticMute *Finding_MuteInfo_StaticMute `json:"staticMute,omitempty"`

	// The list of dynamic mute rules that currently match the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.MuteInfo.dynamic_mute_records
	DynamicMuteRecords []Finding_MuteInfo_DynamicMuteRecord `json:"dynamicMuteRecords,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Finding.MuteInfo.DynamicMuteRecord
type Finding_MuteInfo_DynamicMuteRecord struct {
	// The relative resource name of the mute rule, represented by a mute
	//  config, that created this record, for example
	//  `organizations/123/muteConfigs/mymuteconfig` or
	//  `organizations/123/locations/global/muteConfigs/mymuteconfig`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.MuteInfo.DynamicMuteRecord.mute_config
	MuteConfig *string `json:"muteConfig,omitempty"`

	// When the dynamic mute rule first matched the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.MuteInfo.DynamicMuteRecord.match_time
	MatchTime *string `json:"matchTime,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Finding.MuteInfo.StaticMute
type Finding_MuteInfo_StaticMute struct {
	// The static mute state. If the value is `MUTED` or `UNMUTED`, then the
	//  finding's overall mute state will have the same value.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.MuteInfo.StaticMute.state
	State *string `json:"state,omitempty"`

	// When the static mute was applied.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.MuteInfo.StaticMute.apply_time
	ApplyTime *string `json:"applyTime,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Geolocation
type Geolocation struct {
	// A CLDR.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Geolocation.region_code
	RegionCode *string `json:"regionCode,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.GroupMembership
type GroupMembership struct {
	// Type of group.
	// +kcc:proto:field=google.cloud.securitycenter.v1.GroupMembership.group_type
	GroupType *string `json:"groupType,omitempty"`

	// ID of the group.
	// +kcc:proto:field=google.cloud.securitycenter.v1.GroupMembership.group_id
	GroupID *string `json:"groupID,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.IamBinding
type IamBinding struct {
	// The action that was performed on a Binding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.IamBinding.action
	Action *string `json:"action,omitempty"`

	// Role that is assigned to "members".
	//  For example, "roles/viewer", "roles/editor", or "roles/owner".
	// +kcc:proto:field=google.cloud.securitycenter.v1.IamBinding.role
	Role *string `json:"role,omitempty"`

	// A single identity requesting access for a Cloud Platform resource, for
	//  example, "foo@google.com".
	// +kcc:proto:field=google.cloud.securitycenter.v1.IamBinding.member
	Member *string `json:"member,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Indicator
type Indicator struct {
	// The list of IP addresses that are associated with the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Indicator.ip_addresses
	IPAddresses []string `json:"ipAddresses,omitempty"`

	// List of domains associated to the Finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Indicator.domains
	Domains []string `json:"domains,omitempty"`

	// The list of matched signatures indicating that the given
	//  process is present in the environment.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Indicator.signatures
	Signatures []Indicator_ProcessSignature `json:"signatures,omitempty"`

	// The list of URIs associated to the Findings.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Indicator.uris
	Uris []string `json:"uris,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Indicator.ProcessSignature
type Indicator_ProcessSignature struct {
	// Signature indicating that a binary family was matched.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Indicator.ProcessSignature.memory_hash_signature
	MemoryHashSignature *Indicator_ProcessSignature_MemoryHashSignature `json:"memoryHashSignature,omitempty"`

	// Signature indicating that a YARA rule was matched.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Indicator.ProcessSignature.yara_rule_signature
	YaraRuleSignature *Indicator_ProcessSignature_YaraRuleSignature `json:"yaraRuleSignature,omitempty"`

	// Describes the type of resource associated with the signature.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Indicator.ProcessSignature.signature_type
	SignatureType *string `json:"signatureType,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Indicator.ProcessSignature.MemoryHashSignature
type Indicator_ProcessSignature_MemoryHashSignature struct {
	// The binary family.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Indicator.ProcessSignature.MemoryHashSignature.binary_family
	BinaryFamily *string `json:"binaryFamily,omitempty"`

	// The list of memory hash detections contributing to the binary family
	//  match.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Indicator.ProcessSignature.MemoryHashSignature.detections
	Detections []Indicator_ProcessSignature_MemoryHashSignature_Detection `json:"detections,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Indicator.ProcessSignature.MemoryHashSignature.Detection
type Indicator_ProcessSignature_MemoryHashSignature_Detection struct {
	// The name of the binary associated with the memory hash
	//  signature detection.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Indicator.ProcessSignature.MemoryHashSignature.Detection.binary
	Binary *string `json:"binary,omitempty"`

	// The percentage of memory page hashes in the signature
	//  that were matched.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Indicator.ProcessSignature.MemoryHashSignature.Detection.percent_pages_matched
	PercentPagesMatched *float64 `json:"percentPagesMatched,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Indicator.ProcessSignature.YaraRuleSignature
type Indicator_ProcessSignature_YaraRuleSignature struct {
	// The name of the YARA rule.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Indicator.ProcessSignature.YaraRuleSignature.yara_rule
	YaraRule *string `json:"yaraRule,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.KernelRootkit
type KernelRootkit struct {
	// Rootkit name, when available.
	// +kcc:proto:field=google.cloud.securitycenter.v1.KernelRootkit.name
	Name *string `json:"name,omitempty"`

	// True if unexpected modifications of kernel code memory are present.
	// +kcc:proto:field=google.cloud.securitycenter.v1.KernelRootkit.unexpected_code_modification
	UnexpectedCodeModification *bool `json:"unexpectedCodeModification,omitempty"`

	// True if unexpected modifications of kernel read-only data memory are
	//  present.
	// +kcc:proto:field=google.cloud.securitycenter.v1.KernelRootkit.unexpected_read_only_data_modification
	UnexpectedReadOnlyDataModification *bool `json:"unexpectedReadOnlyDataModification,omitempty"`

	// True if `ftrace` points are present with callbacks pointing to regions
	//  that are not in the expected kernel or module code range.
	// +kcc:proto:field=google.cloud.securitycenter.v1.KernelRootkit.unexpected_ftrace_handler
	UnexpectedFtraceHandler *bool `json:"unexpectedFtraceHandler,omitempty"`

	// True if `kprobe` points are present with callbacks pointing to regions
	//  that are not in the expected kernel or module code range.
	// +kcc:proto:field=google.cloud.securitycenter.v1.KernelRootkit.unexpected_kprobe_handler
	UnexpectedKprobeHandler *bool `json:"unexpectedKprobeHandler,omitempty"`

	// True if kernel code pages that are not in the expected kernel or module
	//  code regions are present.
	// +kcc:proto:field=google.cloud.securitycenter.v1.KernelRootkit.unexpected_kernel_code_pages
	UnexpectedKernelCodePages *bool `json:"unexpectedKernelCodePages,omitempty"`

	// True if system call handlers that are are not in the expected kernel or
	//  module code regions are present.
	// +kcc:proto:field=google.cloud.securitycenter.v1.KernelRootkit.unexpected_system_call_handler
	UnexpectedSystemCallHandler *bool `json:"unexpectedSystemCallHandler,omitempty"`

	// True if interrupt handlers that are are not in the expected kernel or
	//  module code regions are present.
	// +kcc:proto:field=google.cloud.securitycenter.v1.KernelRootkit.unexpected_interrupt_handler
	UnexpectedInterruptHandler *bool `json:"unexpectedInterruptHandler,omitempty"`

	// True if unexpected processes in the scheduler run queue are present. Such
	//  processes are in the run queue, but not in the process task list.
	// +kcc:proto:field=google.cloud.securitycenter.v1.KernelRootkit.unexpected_processes_in_runqueue
	UnexpectedProcessesInRunqueue *bool `json:"unexpectedProcessesInRunqueue,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Kubernetes
type Kubernetes struct {
	// Kubernetes
	//  [Pods](https://cloud.google.com/kubernetes-engine/docs/concepts/pod)
	//  associated with the finding. This field contains Pod records for each
	//  container that is owned by a Pod.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.pods
	Pods []Kubernetes_Pod `json:"pods,omitempty"`

	// Provides Kubernetes
	//  [node](https://cloud.google.com/kubernetes-engine/docs/concepts/cluster-architecture#nodes)
	//  information.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.nodes
	Nodes []Kubernetes_Node `json:"nodes,omitempty"`

	// GKE [node
	//  pools](https://cloud.google.com/kubernetes-engine/docs/concepts/node-pools)
	//  associated with the finding. This field contains node pool information for
	//  each node, when it is available.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.node_pools
	NodePools []Kubernetes_NodePool `json:"nodePools,omitempty"`

	// Provides Kubernetes role information for findings that involve [Roles or
	//  ClusterRoles](https://cloud.google.com/kubernetes-engine/docs/how-to/role-based-access-control).
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.roles
	Roles []Kubernetes_Role `json:"roles,omitempty"`

	// Provides Kubernetes role binding information for findings that involve
	//  [RoleBindings or
	//  ClusterRoleBindings](https://cloud.google.com/kubernetes-engine/docs/how-to/role-based-access-control).
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.bindings
	Bindings []Kubernetes_Binding `json:"bindings,omitempty"`

	// Provides information on any Kubernetes access reviews (privilege checks)
	//  relevant to the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.access_reviews
	AccessReviews []Kubernetes_AccessReview `json:"accessReviews,omitempty"`

	// Kubernetes objects related to the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.objects
	Objects []Kubernetes_Object `json:"objects,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Kubernetes.AccessReview
type Kubernetes_AccessReview struct {
	// The API group of the resource. "*" means all.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.AccessReview.group
	Group *string `json:"group,omitempty"`

	// Namespace of the action being requested. Currently, there is no
	//  distinction between no namespace and all namespaces.  Both
	//  are represented by "" (empty).
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.AccessReview.ns
	Ns *string `json:"ns,omitempty"`

	// The name of the resource being requested. Empty means all.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.AccessReview.name
	Name *string `json:"name,omitempty"`

	// The optional resource type requested. "*" means all.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.AccessReview.resource
	Resource *string `json:"resource,omitempty"`

	// The optional subresource type.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.AccessReview.subresource
	Subresource *string `json:"subresource,omitempty"`

	// A Kubernetes resource API verb, like get, list, watch, create, update,
	//  delete, proxy. "*" means all.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.AccessReview.verb
	Verb *string `json:"verb,omitempty"`

	// The API version of the resource. "*" means all.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.AccessReview.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Kubernetes.Binding
type Kubernetes_Binding struct {
	// Namespace for the binding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.Binding.ns
	Ns *string `json:"ns,omitempty"`

	// Name for the binding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.Binding.name
	Name *string `json:"name,omitempty"`

	// The Role or ClusterRole referenced by the binding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.Binding.role
	Role *Kubernetes_Role `json:"role,omitempty"`

	// Represents one or more subjects that are bound to the role. Not always
	//  available for PATCH requests.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.Binding.subjects
	Subjects []Kubernetes_Subject `json:"subjects,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Kubernetes.Node
type Kubernetes_Node struct {
	// [Full resource name](https://google.aip.dev/122#full-resource-names) of
	//  the Compute Engine VM running the cluster node.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.Node.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Kubernetes.NodePool
type Kubernetes_NodePool struct {
	// Kubernetes node pool name.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.NodePool.name
	Name *string `json:"name,omitempty"`

	// Nodes associated with the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.NodePool.nodes
	Nodes []Kubernetes_Node `json:"nodes,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Kubernetes.Object
type Kubernetes_Object struct {
	// Kubernetes object group, such as "policy.k8s.io/v1".
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.Object.group
	Group *string `json:"group,omitempty"`

	// Kubernetes object kind, such as "Namespace".
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.Object.kind
	Kind *string `json:"kind,omitempty"`

	// Kubernetes object namespace. Must be a valid DNS label. Named
	//  "ns" to avoid collision with C++ namespace keyword. For details see
	//  https://kubernetes.io/docs/tasks/administer-cluster/namespaces/.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.Object.ns
	Ns *string `json:"ns,omitempty"`

	// Kubernetes object name. For details see
	//  https://kubernetes.io/docs/concepts/overview/working-with-objects/names/.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.Object.name
	Name *string `json:"name,omitempty"`

	// Pod containers associated with this finding, if any.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.Object.containers
	Containers []Container `json:"containers,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Kubernetes.Pod
type Kubernetes_Pod struct {
	// Kubernetes Pod namespace.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.Pod.ns
	Ns *string `json:"ns,omitempty"`

	// Kubernetes Pod name.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.Pod.name
	Name *string `json:"name,omitempty"`

	// Pod labels.  For Kubernetes containers, these are applied to the
	//  container.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.Pod.labels
	Labels []Label `json:"labels,omitempty"`

	// Pod containers associated with this finding, if any.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.Pod.containers
	Containers []Container `json:"containers,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Kubernetes.Role
type Kubernetes_Role struct {
	// Role type.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.Role.kind
	Kind *string `json:"kind,omitempty"`

	// Role namespace.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.Role.ns
	Ns *string `json:"ns,omitempty"`

	// Role name.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.Role.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Kubernetes.Subject
type Kubernetes_Subject struct {
	// Authentication type for the subject.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.Subject.kind
	Kind *string `json:"kind,omitempty"`

	// Namespace for the subject.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.Subject.ns
	Ns *string `json:"ns,omitempty"`

	// Name for the subject.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Kubernetes.Subject.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Label
type Label struct {
	// Name of the label.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Label.name
	Name *string `json:"name,omitempty"`

	// Value that corresponds to the label's name.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Label.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.LoadBalancer
type LoadBalancer struct {
	// The name of the load balancer associated with the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.LoadBalancer.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.LogEntry
type LogEntry struct {
	// An individual entry in a log stored in Cloud Logging.
	// +kcc:proto:field=google.cloud.securitycenter.v1.LogEntry.cloud_logging_entry
	CloudLoggingEntry *CloudLoggingEntry `json:"cloudLoggingEntry,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.MitreAttack
type MitreAttack struct {
	// The MITRE ATT&CK tactic most closely represented by this finding, if any.
	// +kcc:proto:field=google.cloud.securitycenter.v1.MitreAttack.primary_tactic
	PrimaryTactic *string `json:"primaryTactic,omitempty"`

	// The MITRE ATT&CK technique most closely represented by this finding, if
	//  any. primary_techniques is a repeated field because there are multiple
	//  levels of MITRE ATT&CK techniques.  If the technique most closely
	//  represented by this finding is a sub-technique (e.g. `SCANNING_IP_BLOCKS`),
	//  both the sub-technique and its parent technique(s) will be listed (e.g.
	//  `SCANNING_IP_BLOCKS`, `ACTIVE_SCANNING`).
	// +kcc:proto:field=google.cloud.securitycenter.v1.MitreAttack.primary_techniques
	PrimaryTechniques []string `json:"primaryTechniques,omitempty"`

	// Additional MITRE ATT&CK tactics related to this finding, if any.
	// +kcc:proto:field=google.cloud.securitycenter.v1.MitreAttack.additional_tactics
	AdditionalTactics []string `json:"additionalTactics,omitempty"`

	// Additional MITRE ATT&CK techniques related to this finding, if any, along
	//  with any of their respective parent techniques.
	// +kcc:proto:field=google.cloud.securitycenter.v1.MitreAttack.additional_techniques
	AdditionalTechniques []string `json:"additionalTechniques,omitempty"`

	// The MITRE ATT&CK version referenced by the above fields. E.g. "8".
	// +kcc:proto:field=google.cloud.securitycenter.v1.MitreAttack.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Notebook
type Notebook struct {
	// The name of the notebook.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Notebook.name
	Name *string `json:"name,omitempty"`

	// The source notebook service, for example, "Colab Enterprise".
	// +kcc:proto:field=google.cloud.securitycenter.v1.Notebook.service
	Service *string `json:"service,omitempty"`

	// The user ID of the latest author to modify the notebook.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Notebook.last_author
	LastAuthor *string `json:"lastAuthor,omitempty"`

	// The most recent time the notebook was updated.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Notebook.notebook_update_time
	NotebookUpdateTime *string `json:"notebookUpdateTime,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.OrgPolicy
type OrgPolicy struct {
	// The resource name of the org policy.
	//  Example:
	//  "organizations/{organization_id}/policies/{constraint_name}"
	// +kcc:proto:field=google.cloud.securitycenter.v1.OrgPolicy.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Package
type Package struct {
	// The name of the package where the vulnerability was detected.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Package.package_name
	PackageName *string `json:"packageName,omitempty"`

	// The CPE URI where the vulnerability was detected.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Package.cpe_uri
	CpeURI *string `json:"cpeURI,omitempty"`

	// Type of package, for example, os, maven, or go.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Package.package_type
	PackageType *string `json:"packageType,omitempty"`

	// The version of the package.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Package.package_version
	PackageVersion *string `json:"packageVersion,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Process
type Process struct {
	// The process name, as displayed in utilities like `top` and `ps`. This name
	//  can be accessed through `/proc/[pid]/comm` and changed with
	//  `prctl(PR_SET_NAME)`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Process.name
	Name *string `json:"name,omitempty"`

	// File information for the process executable.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Process.binary
	Binary *File `json:"binary,omitempty"`

	// File information for libraries loaded by the process.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Process.libraries
	Libraries []File `json:"libraries,omitempty"`

	// When the process represents the invocation of a script, `binary` provides
	//  information about the interpreter, while `script` provides information
	//  about the script file provided to the interpreter.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Process.script
	Script *File `json:"script,omitempty"`

	// Process arguments as JSON encoded strings.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Process.args
	Args []string `json:"args,omitempty"`

	// True if `args` is incomplete.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Process.arguments_truncated
	ArgumentsTruncated *bool `json:"argumentsTruncated,omitempty"`

	// Process environment variables.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Process.env_variables
	EnvVariables []EnvironmentVariable `json:"envVariables,omitempty"`

	// True if `env_variables` is incomplete.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Process.env_variables_truncated
	EnvVariablesTruncated *bool `json:"envVariablesTruncated,omitempty"`

	// The process ID.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Process.pid
	Pid *int64 `json:"pid,omitempty"`

	// The parent process ID.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Process.parent_pid
	ParentPid *int64 `json:"parentPid,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Reference
type Reference struct {
	// Source of the reference e.g. NVD
	// +kcc:proto:field=google.cloud.securitycenter.v1.Reference.source
	Source *string `json:"source,omitempty"`

	// Uri for the mentioned source e.g.
	//  https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2021-34527.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Reference.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Requests
type Requests struct {
	// For 'Increasing deny ratio', the ratio is the denied traffic divided by the
	//  allowed traffic. For 'Allowed traffic spike', the ratio is the allowed
	//  traffic in the short term divided by allowed traffic in the long term.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Requests.ratio
	Ratio *float64 `json:"ratio,omitempty"`

	// Allowed RPS (requests per second) in the short term.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Requests.short_term_allowed
	ShortTermAllowed *int32 `json:"shortTermAllowed,omitempty"`

	// Allowed RPS (requests per second) over the long term.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Requests.long_term_allowed
	LongTermAllowed *int32 `json:"longTermAllowed,omitempty"`

	// Denied RPS (requests per second) over the long term.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Requests.long_term_denied
	LongTermDenied *int32 `json:"longTermDenied,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.SecurityBulletin
type SecurityBulletin struct {
	// ID of the bulletin corresponding to the vulnerability.
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityBulletin.bulletin_id
	BulletinID *string `json:"bulletinID,omitempty"`

	// Submission time of this Security Bulletin.
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityBulletin.submission_time
	SubmissionTime *string `json:"submissionTime,omitempty"`

	// This represents a version that the cluster receiving this notification
	//  should be upgraded to, based on its current version. For example, 1.15.0
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityBulletin.suggested_upgrade_version
	SuggestedUpgradeVersion *string `json:"suggestedUpgradeVersion,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.SecurityMarks
type SecurityMarks struct {
	// The relative resource name of the SecurityMarks. See:
	//  https://cloud.google.com/apis/design/resource_names#relative_resource_name
	//  Examples:
	//  "organizations/{organization_id}/assets/{asset_id}/securityMarks"
	//  "organizations/{organization_id}/sources/{source_id}/findings/{finding_id}/securityMarks".
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityMarks.name
	Name *string `json:"name,omitempty"`

	// Mutable user specified security marks belonging to the parent resource.
	//  Constraints are as follows:
	//
	//    * Keys and values are treated as case insensitive
	//    * Keys must be between 1 - 256 characters (inclusive)
	//    * Keys must be letters, numbers, underscores, or dashes
	//    * Values have leading and trailing whitespace trimmed, remaining
	//      characters must be between 1 - 4096 characters (inclusive)
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityMarks.marks
	Marks map[string]string `json:"marks,omitempty"`

	// The canonical name of the marks.
	//  Examples:
	//  "organizations/{organization_id}/assets/{asset_id}/securityMarks"
	//  "folders/{folder_id}/assets/{asset_id}/securityMarks"
	//  "projects/{project_number}/assets/{asset_id}/securityMarks"
	//  "organizations/{organization_id}/sources/{source_id}/findings/{finding_id}/securityMarks"
	//  "folders/{folder_id}/sources/{source_id}/findings/{finding_id}/securityMarks"
	//  "projects/{project_number}/sources/{source_id}/findings/{finding_id}/securityMarks"
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityMarks.canonical_name
	CanonicalName *string `json:"canonicalName,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.SecurityPolicy
type SecurityPolicy struct {
	// The name of the Google Cloud Armor security policy, for example,
	//  "my-security-policy".
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityPolicy.name
	Name *string `json:"name,omitempty"`

	// The type of Google Cloud Armor security policy for example, 'backend
	//  security policy', 'edge security policy', 'network edge security policy',
	//  or 'always-on DDoS protection'.
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityPolicy.type
	Type *string `json:"type,omitempty"`

	// Whether or not the associated rule or policy is in preview mode.
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityPolicy.preview
	Preview *bool `json:"preview,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.SecurityPosture
type SecurityPosture struct {
	// Name of the posture, for example, `CIS-Posture`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityPosture.name
	Name *string `json:"name,omitempty"`

	// The version of the posture, for example, `c7cfa2a8`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityPosture.revision_id
	RevisionID *string `json:"revisionID,omitempty"`

	// The project, folder, or organization on which the posture is deployed,
	//  for example, `projects/{project_number}`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityPosture.posture_deployment_resource
	PostureDeploymentResource *string `json:"postureDeploymentResource,omitempty"`

	// The name of the posture deployment, for example,
	//  `organizations/{org_id}/posturedeployments/{posture_deployment_id}`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityPosture.posture_deployment
	PostureDeployment *string `json:"postureDeployment,omitempty"`

	// The name of the updated policy, for example,
	//  `projects/{project_id}/policies/{constraint_name}`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityPosture.changed_policy
	ChangedPolicy *string `json:"changedPolicy,omitempty"`

	// The name of the updated policyset, for example, `cis-policyset`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityPosture.policy_set
	PolicySet *string `json:"policySet,omitempty"`

	// The ID of the updated policy, for example, `compute-policy-1`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityPosture.policy
	Policy *string `json:"policy,omitempty"`

	// The details about a change in an updated policy that violates the deployed
	//  posture.
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityPosture.policy_drift_details
	PolicyDriftDetails []SecurityPosture_PolicyDriftDetails `json:"policyDriftDetails,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.SecurityPosture.PolicyDriftDetails
type SecurityPosture_PolicyDriftDetails struct {
	// The name of the updated field, for example
	//  constraint.implementation.policy_rules[0].enforce
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityPosture.PolicyDriftDetails.field
	Field *string `json:"field,omitempty"`

	// The value of this field that was configured in a posture, for example,
	//  `true` or `allowed_values={"projects/29831892"}`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityPosture.PolicyDriftDetails.expected_value
	ExpectedValue *string `json:"expectedValue,omitempty"`

	// The detected value that violates the deployed posture, for example,
	//  `false` or `allowed_values={"projects/22831892"}`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityPosture.PolicyDriftDetails.detected_value
	DetectedValue *string `json:"detectedValue,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.ServiceAccountDelegationInfo
type ServiceAccountDelegationInfo struct {
	// The email address of a Google account.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ServiceAccountDelegationInfo.principal_email
	PrincipalEmail *string `json:"principalEmail,omitempty"`

	// A string representing the principal_subject associated with the identity.
	//  As compared to `principal_email`, supports principals that aren't
	//  associated with email addresses, such as third party principals. For most
	//  identities, the format will be `principal://iam.googleapis.com/{identity
	//  pool name}/subjects/{subject}` except for some GKE identities
	//  (GKE_WORKLOAD, FREEFORM, GKE_HUB_WORKLOAD) that are still in the legacy
	//  format `serviceAccount:{identity pool name}[{subject}]`
	// +kcc:proto:field=google.cloud.securitycenter.v1.ServiceAccountDelegationInfo.principal_subject
	PrincipalSubject *string `json:"principalSubject,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.ToxicCombination
type ToxicCombination struct {
	// The
	//  [Attack exposure
	//  score](https://cloud.google.com/security-command-center/docs/attack-exposure-learn#attack_exposure_scores)
	//  of this toxic combination. The score is a measure of how much this toxic
	//  combination exposes one or more high-value resources to potential attack.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ToxicCombination.attack_exposure_score
	AttackExposureScore *float64 `json:"attackExposureScore,omitempty"`

	// List of resource names of findings associated with this toxic combination.
	//  For example, `organizations/123/sources/456/findings/789`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ToxicCombination.related_findings
	RelatedFindings []string `json:"relatedFindings,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Vulnerability
type Vulnerability struct {
	// CVE stands for Common Vulnerabilities and Exposures
	//  (https://cve.mitre.org/about/)
	// +kcc:proto:field=google.cloud.securitycenter.v1.Vulnerability.cve
	Cve *Cve `json:"cve,omitempty"`

	// The offending package is relevant to the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Vulnerability.offending_package
	OffendingPackage *Package `json:"offendingPackage,omitempty"`

	// The fixed package is relevant to the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Vulnerability.fixed_package
	FixedPackage *Package `json:"fixedPackage,omitempty"`

	// The security bulletin is relevant to this finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Vulnerability.security_bulletin
	SecurityBulletin *SecurityBulletin `json:"securityBulletin,omitempty"`
}

// +kcc:proto=google.protobuf.ListValue
type ListValue struct {
	// Repeated field of dynamically typed values.
	// +kcc:proto:field=google.protobuf.ListValue.values
	Values []Value `json:"values,omitempty"`
}

// +kcc:proto=google.protobuf.Value
type Value struct {
	// Represents a null value.
	// +kcc:proto:field=google.protobuf.Value.null_value
	NullValue *string `json:"nullValue,omitempty"`

	// Represents a double value.
	// +kcc:proto:field=google.protobuf.Value.number_value
	NumberValue *float64 `json:"numberValue,omitempty"`

	// Represents a string value.
	// +kcc:proto:field=google.protobuf.Value.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Represents a boolean value.
	// +kcc:proto:field=google.protobuf.Value.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// Represents a structured value.
	// +kcc:proto:field=google.protobuf.Value.struct_value
	StructValue map[string]string `json:"structValue,omitempty"`

	// Represents a repeated `Value`.
	// +kcc:proto:field=google.protobuf.Value.list_value
	ListValue *ListValue `json:"listValue,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.Finding
type FindingObservedState struct {
	// Output only. User specified security marks. These marks are entirely
	//  managed by the user and come from the SecurityMarks resource that belongs
	//  to the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.security_marks
	SecurityMarks *SecurityMarks `json:"securityMarks,omitempty"`

	// Output only. The most recent time this finding was muted or unmuted.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.mute_update_time
	MuteUpdateTime *string `json:"muteUpdateTime,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Output only. The mute information regarding this finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.mute_info
	MuteInfo *Finding_MuteInfo `json:"muteInfo,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Output only. The human readable display name of the finding source such as
	//  "Event Threat Detection" or "Security Health Analytics".
	// +kcc:proto:field=google.cloud.securitycenter.v1.Finding.parent_display_name
	ParentDisplayName *string `json:"parentDisplayName,omitempty"`
}
