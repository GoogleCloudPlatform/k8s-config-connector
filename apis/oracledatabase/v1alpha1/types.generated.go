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


// +kcc:proto=google.cloud.oracledatabase.v1.CloudVmCluster
type CloudVmCluster struct {
	// Identifier. The name of the VM Cluster resource with the format:
	//  projects/{project}/locations/{region}/cloudVmClusters/{cloud_vm_cluster}
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.name
	Name *string `json:"name,omitempty"`

	// Required. The name of the Exadata Infrastructure resource on which VM
	//  cluster resource is created, in the following format:
	//  projects/{project}/locations/{region}/cloudExadataInfrastuctures/{cloud_extradata_infrastructure}
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.exadata_infrastructure
	ExadataInfrastructure *string `json:"exadataInfrastructure,omitempty"`

	// Optional. User friendly name for this resource.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Various properties of the VM Cluster.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.properties
	Properties *CloudVmClusterProperties `json:"properties,omitempty"`

	// Optional. Labels or tags associated with the VM Cluster.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Network settings. CIDR to use for cluster IP allocation.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.cidr
	Cidr *string `json:"cidr,omitempty"`

	// Required. CIDR range of the backup subnet.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.backup_subnet_cidr
	BackupSubnetCidr *string `json:"backupSubnetCidr,omitempty"`

	// Required. The name of the VPC network.
	//  Format: projects/{project}/global/networks/{network}
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.network
	Network *string `json:"network,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.CloudVmClusterProperties
type CloudVmClusterProperties struct {

	// Required. License type of VM Cluster.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.license_type
	LicenseType *string `json:"licenseType,omitempty"`

	// Optional. Grid Infrastructure Version.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.gi_version
	GiVersion *string `json:"giVersion,omitempty"`

	// Optional. Time zone of VM Cluster to set. Defaults to UTC if not specified.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.time_zone
	TimeZone *TimeZone `json:"timeZone,omitempty"`

	// Optional. SSH public keys to be stored with cluster.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.ssh_public_keys
	SSHPublicKeys []string `json:"sshPublicKeys,omitempty"`

	// Optional. Number of database servers.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.node_count
	NodeCount *int32 `json:"nodeCount,omitempty"`

	// Optional. OCPU count per VM. Minimum is 0.1.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.ocpu_count
	OcpuCount *float32 `json:"ocpuCount,omitempty"`

	// Optional. Memory allocated in GBs.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.memory_size_gb
	MemorySizeGB *int32 `json:"memorySizeGB,omitempty"`

	// Optional. Local storage per VM.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.db_node_storage_size_gb
	DbNodeStorageSizeGB *int32 `json:"dbNodeStorageSizeGB,omitempty"`

	// Optional. The data disk group size to be allocated in TBs.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.data_storage_size_tb
	DataStorageSizeTb *float64 `json:"dataStorageSizeTb,omitempty"`

	// Optional. The type of redundancy.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.disk_redundancy
	DiskRedundancy *string `json:"diskRedundancy,omitempty"`

	// Optional. Use exadata sparse snapshots.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.sparse_diskgroup_enabled
	SparseDiskgroupEnabled *bool `json:"sparseDiskgroupEnabled,omitempty"`

	// Optional. Use local backup.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.local_backup_enabled
	LocalBackupEnabled *bool `json:"localBackupEnabled,omitempty"`

	// Optional. Prefix for VM cluster host names.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.hostname_prefix
	HostnamePrefix *string `json:"hostnamePrefix,omitempty"`

	// Optional. Data collection options for diagnostics.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.diagnostics_data_collection_options
	DiagnosticsDataCollectionOptions *DataCollectionOptions `json:"diagnosticsDataCollectionOptions,omitempty"`

	// Required. Number of enabled CPU cores.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.cpu_core_count
	CpuCoreCount *int32 `json:"cpuCoreCount,omitempty"`

	// Optional. Operating system version of the image.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.system_version
	SystemVersion *string `json:"systemVersion,omitempty"`

	// Optional. OCID of database servers.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.db_server_ocids
	DbServerOcids []string `json:"dbServerOcids,omitempty"`

	// Optional. OCI Cluster name.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.cluster_name
	ClusterName *string `json:"clusterName,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.DataCollectionOptions
type DataCollectionOptions struct {
	// Optional. Indicates whether diagnostic collection is enabled for the VM
	//  cluster
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DataCollectionOptions.diagnostics_events_enabled
	DiagnosticsEventsEnabled *bool `json:"diagnosticsEventsEnabled,omitempty"`

	// Optional. Indicates whether health monitoring is enabled for the VM cluster
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DataCollectionOptions.health_monitoring_enabled
	HealthMonitoringEnabled *bool `json:"healthMonitoringEnabled,omitempty"`

	// Optional. Indicates whether incident logs and trace collection are enabled
	//  for the VM cluster
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DataCollectionOptions.incident_logs_enabled
	IncidentLogsEnabled *bool `json:"incidentLogsEnabled,omitempty"`
}

// +kcc:proto=google.type.TimeZone
type TimeZone struct {
	// IANA Time Zone Database time zone, e.g. "America/New_York".
	// +kcc:proto:field=google.type.TimeZone.id
	ID *string `json:"id,omitempty"`

	// Optional. IANA Time Zone Database version number, e.g. "2019a".
	// +kcc:proto:field=google.type.TimeZone.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.CloudVmCluster
type CloudVmClusterObservedState struct {
	// Output only. Google Cloud Platform location where Oracle Exadata is hosted.
	//  It is same as Google Cloud Platform Oracle zone of Exadata infrastructure.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.gcp_oracle_zone
	GcpOracleZone *string `json:"gcpOracleZone,omitempty"`

	// Optional. Various properties of the VM Cluster.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.properties
	Properties *CloudVmClusterPropertiesObservedState `json:"properties,omitempty"`

	// Output only. The date and time that the VM cluster was created.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.CloudVmClusterProperties
type CloudVmClusterPropertiesObservedState struct {
	// Output only. Oracle Cloud Infrastructure ID of VM Cluster.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.ocid
	Ocid *string `json:"ocid,omitempty"`

	// Output only. Shape of VM Cluster.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.shape
	Shape *string `json:"shape,omitempty"`

	// Output only. The storage allocation for the disk group, in gigabytes (GB).
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.storage_size_gb
	StorageSizeGB *int32 `json:"storageSizeGB,omitempty"`

	// Output only. State of the cluster.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.state
	State *string `json:"state,omitempty"`

	// Output only. SCAN listener port - TCP
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.scan_listener_port_tcp
	ScanListenerPortTcp *int32 `json:"scanListenerPortTcp,omitempty"`

	// Output only. SCAN listener port - TLS
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.scan_listener_port_tcp_ssl
	ScanListenerPortTcpSsl *int32 `json:"scanListenerPortTcpSsl,omitempty"`

	// Output only. Parent DNS domain where SCAN DNS and hosts names are
	//  qualified. ex: ocispdelegated.ocisp10jvnet.oraclevcn.com
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.domain
	Domain *string `json:"domain,omitempty"`

	// Output only. SCAN DNS name.
	//  ex: sp2-yi0xq-scan.ocispdelegated.ocisp10jvnet.oraclevcn.com
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.scan_dns
	ScanDns *string `json:"scanDns,omitempty"`

	// Output only. host name without domain.
	//  format: "<hostname_prefix>-" with some suffix.
	//  ex: sp2-yi0xq where "sp2" is the hostname_prefix.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.hostname
	Hostname *string `json:"hostname,omitempty"`

	// Output only. OCIDs of scan IPs.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.scan_ip_ids
	ScanIPIds []string `json:"scanIPIds,omitempty"`

	// Output only. OCID of scan DNS record.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.scan_dns_record_id
	ScanDnsRecordID *string `json:"scanDnsRecordID,omitempty"`

	// Output only. Deep link to the OCI console to view this resource.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.oci_url
	OciURL *string `json:"ociURL,omitempty"`

	// Output only. Compartment ID of cluster.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.compartment_id
	CompartmentID *string `json:"compartmentID,omitempty"`

	// Output only. DNS listener IP.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmClusterProperties.dns_listener_ip
	DnsListenerIP *string `json:"dnsListenerIP,omitempty"`
}
