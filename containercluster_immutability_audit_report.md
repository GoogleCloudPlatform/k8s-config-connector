# KCC vs TF vs Backend Implementation Immutability Audit for ContainerCluster

| Field | KCC | TF | Backend Implementation                | Verdict |
|---|---|---|---------------------------------------|---|
| addonsConfig | False | False | Mutable                               | ✅ Match |
| addonsConfig.cloudrunConfig | False | False | Mutable                               | ✅ Match |
| addonsConfig.cloudrunConfig.disabled | False | False | Mutable                               | ✅ Match |
| addonsConfig.cloudrunConfig.loadBalancer<br>Type | False | False | Mutable                               | ✅ Match |
| addonsConfig.configConnectorConfig | False | False | Mutable                               | ✅ Match |
| addonsConfig.configConnectorConfig.enabl<br>ed | False | False | Mutable                               | ✅ Match |
| addonsConfig.dnsCacheConfig | False | False | Mutable                               | ✅ Match |
| addonsConfig.dnsCacheConfig.enabled | False | False | Mutable                               | ✅ Match |
| addonsConfig.gcePersistentDiskCsiDriverC<br>onfig | False | False | Mutable                               | ✅ Match |
| addonsConfig.gcePersistentDiskCsiDriverC<br>onfig.enabled | False | False | Mutable                               | ✅ Match |
| addonsConfig.gcpFilestoreCsiDriverConfig | False | False | Mutable                               | ✅ Match |
| addonsConfig.gcpFilestoreCsiDriverConfig<br>.enabled | False | False | Mutable                               | ✅ Match |
| addonsConfig.gcsFuseCsiDriverConfig | False | False | Mutable                               | ✅ Match |
| addonsConfig.gcsFuseCsiDriverConfig.enab<br>led | False | False | Mutable                               | ✅ Match |
| addonsConfig.gkeBackupAgentConfig | False | False | Mutable                               | ✅ Match |
| addonsConfig.gkeBackupAgentConfig.enable<br>d | False | False | Mutable                               | ✅ Match |
| addonsConfig.horizontalPodAutoscaling | False | False | Mutable                               | ✅ Match |
| addonsConfig.horizontalPodAutoscaling.di<br>sabled | False | False | Mutable                               | ✅ Match |
| addonsConfig.httpLoadBalancing | False | False | Mutable                               | ✅ Match |
| addonsConfig.httpLoadBalancing.disabled | False | False | Mutable                               | ✅ Match |
| addonsConfig.istioConfig | False | False | Mutable                               | ✅ Match |
| addonsConfig.istioConfig.auth | False | False | Mutable                               | ✅ Match |
| addonsConfig.istioConfig.disabled | False | False | Mutable                               | ✅ Match |
| addonsConfig.kalmConfig | False | False | Mutable                               | ✅ Match |
| addonsConfig.kalmConfig.enabled | False | False | Mutable                               | ✅ Match |
| addonsConfig.networkPolicyConfig | False | False | Mutable                               | ✅ Match |
| addonsConfig.networkPolicyConfig.disable<br>d | False | False | Mutable                               | ✅ Match |
| allowNetAdmin | False | False | N/A                                   | ❓ (No Internal Data) |
| authenticatorGroupsConfig | False | False | Mutable                               | ✅ Match |
| authenticatorGroupsConfig.securityGroup | False | False | Mutable                               | ✅ Match |
| binaryAuthorization | False | False | Mutable                               | ✅ Match |
| binaryAuthorization.enabled | False | False | Mutable                               | ✅ Match |
| binaryAuthorization.evaluationMode | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.bootDiskKMSKeyRef | True | True | Mutable                               | ⚠️ KCC Stricter |
| clusterAutoscaling.autoProvisioningDefau<br>lts.diskSize | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.diskType | N/A | False | Mutable                               | N/A |
| clusterAutoscaling.autoProvisioningDefau<br>lts.imageType | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.management | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.management.autoRepair | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.management.autoUpgrade | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.management.upgradeOptions | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.management.upgradeOptions[].autoUpgr<br>adeStartTime | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.management.upgradeOptions[].descript<br>ion | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.minCpuPlatform | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.oauthScopes | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.serviceAccountRef | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.shieldedInstanceConfig | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.shieldedInstanceConfig.enableIntegri<br>tyMonitoring | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.shieldedInstanceConfig.enableSecureB<br>oot | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.upgradeSettings | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.upgradeSettings.blueGreenSettings | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.upgradeSettings.blueGreenSettings.no<br>dePoolSoakDuration | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.upgradeSettings.blueGreenSettings.st<br>andardRolloutPolicy | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.upgradeSettings.blueGreenSettings.st<br>andardRolloutPolicy.batchNodeCount | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.upgradeSettings.blueGreenSettings.st<br>andardRolloutPolicy.batchPercentage | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.upgradeSettings.blueGreenSettings.st<br>andardRolloutPolicy.batchSoakDuration | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.upgradeSettings.maxSurge | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.upgradeSettings.maxUnavailable | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoProvisioningDefau<br>lts.upgradeSettings.strategy | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.autoscalingProfile | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.defaultComputeClassCo<br>nfig | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.defaultComputeClassCo<br>nfig.enabled | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.enabled | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.resourceLimits | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.resourceLimits[].maxi<br>mum | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.resourceLimits[].mini<br>mum | False | False | Mutable                               | ✅ Match |
| clusterAutoscaling.resourceLimits[].reso<br>urceType | False | False | Mutable                               | ✅ Match |
| clusterIpv4Cidr | True | True | Immutable                             | ✅ Match |
| clusterTelemetry | False | False | Mutable                               | ✅ Match |
| clusterTelemetry.type | False | False | Mutable                               | ✅ Match |
| confidentialNodes | True | True | N/A                                   | ❓ (No Internal Data) |
| confidentialNodes.enabled | True | True | N/A                                   | ❓ (No Internal Data) |
| controlPlaneEndpointsConfig | False | False | Mutable                               | ✅ Match |
| controlPlaneEndpointsConfig.dnsEndpointC<br>onfig | False | False | Mutable                               | ✅ Match |
| controlPlaneEndpointsConfig.dnsEndpointC<br>onfig.allowExternalTraffic | False | False | Mutable                               | ✅ Match |
| controlPlaneEndpointsConfig.dnsEndpointC<br>onfig.endpoint | N/A | False | Mutable                               | N/A |
| controlPlaneEndpointsConfig.ipEndpointsC<br>onfig | False | False | Mutable                               | ✅ Match |
| controlPlaneEndpointsConfig.ipEndpointsC<br>onfig.enabled | False | False | Mutable                               | ✅ Match |
| costManagementConfig | False | False | Mutable                               | ✅ Match |
| costManagementConfig.enabled | False | False | Mutable                               | ✅ Match |
| databaseEncryption | False | False | Mutable                               | ✅ Match |
| databaseEncryption.keyName | False | False | Mutable                               | ✅ Match |
| databaseEncryption.state | False | False | Mutable                               | ✅ Match |
| datapathProvider | True | True | Mutable                               | ⚠️ KCC Stricter |
| defaultMaxPodsPerNode | True | True | N/A                                   | ❓ (No Internal Data) |
| defaultSnatStatus | False | False | Mutable                               | ✅ Match |
| defaultSnatStatus.disabled | False | False | Mutable                               | ✅ Match |
| description | True | True | Immutable                             | ✅ Match |
| dnsConfig | True | True | Mutable                               | ⚠️ KCC Stricter |
| dnsConfig.clusterDns | False | False | Mutable                               | ✅ Match |
| dnsConfig.clusterDnsDomain | False | False | Mutable                               | ✅ Match |
| dnsConfig.clusterDnsScope | False | False | Mutable                               | ✅ Match |
| enableAutopilot | True | True | Mutable                               | ⚠️ KCC Stricter |
| enableBinaryAuthorization | False | False | Mutable                               | ✅ Match |
| enableCiliumClusterwideNetworkPolicy | False | False | Mutable                               | ✅ Match |
| enableFqdnNetworkPolicy | False | False | Mutable                               | ✅ Match |
| enableIntranodeVisibility | False | False | Pending Verification (Parent Mutable) | ⚠️ Pending Verification |
| enableK8sBetaApis | False | False | Mutable                               | ✅ Match |
| enableK8sBetaApis.enabledApis | False | False | N/A                                   | ❓ (No Internal Data) |
| enableKubernetesAlpha | True | True | Immutable                             | ✅ Match |
| enableL4IlbSubsetting | False | False | N/A                                   | ❓ (No Internal Data) |
| enableLegacyAbac | False | False | Mutable                               | ✅ Match |
| enableMultiNetworking | True | True | Mutable                               | ⚠️ KCC Stricter |
| enableShieldedNodes | False | False | Mutable                               | ✅ Match |
| enableTpu | True | True | Immutable                             | ✅ Match |
| endpoint | N/A | False | Immutable                             | N/A |
| gatewayApiConfig | False | False | Mutable                               | ✅ Match |
| gatewayApiConfig.channel | False | False | Mutable                               | ✅ Match |
| identityServiceConfig | False | False | Mutable                               | ✅ Match |
| identityServiceConfig.enabled | False | False | Mutable                               | ✅ Match |
| initialNodeCount | True | True | Immutable                             | ✅ Match |
| ipAllocationPolicy | True | True | Immutable                             | ✅ Match |
| ipAllocationPolicy.additionalPodRangesCo<br>nfig | False | False | Immutable                             | ❌ KCC Lax (Risk) |
| ipAllocationPolicy.additionalPodRangesCo<br>nfig.podRangeNames | False | False | Immutable                             | ❌ KCC Lax (Risk) |
| ipAllocationPolicy.clusterIpv4CidrBlock | True | True | Immutable                             | ✅ Match |
| ipAllocationPolicy.clusterSecondaryRange<br>Name | True | True | Immutable                             | ✅ Match |
| ipAllocationPolicy.podCidrOverprovisionC<br>onfig | True | True | Immutable                             | ✅ Match |
| ipAllocationPolicy.podCidrOverprovisionC<br>onfig.disabled | False | False | Immutable                             | ❌ KCC Lax (Risk) |
| ipAllocationPolicy.servicesIpv4CidrBlock | True | True | Immutable                             | ✅ Match |
| ipAllocationPolicy.servicesSecondaryRang<br>eName | True | True | Immutable                             | ✅ Match |
| ipAllocationPolicy.stackType | True | True | Immutable                             | ✅ Match |
| labelFingerprint | N/A | False | N/A                                   | ❓ (No Internal Data) |
| location | True | True | Immutable                             | ✅ Match |
| loggingConfig | False | False | Mutable                               | ✅ Match |
| loggingConfig.enableComponents | False | False | Mutable                               | ✅ Match |
| loggingService | False | False | Mutable                               | ✅ Match |
| maintenancePolicy | False | False | Mutable                               | ✅ Match |
| maintenancePolicy.dailyMaintenanceWindow | False | False | Mutable                               | ✅ Match |
| maintenancePolicy.dailyMaintenanceWindow<br>.duration | False | False | Mutable                               | ✅ Match |
| maintenancePolicy.dailyMaintenanceWindow<br>.startTime | False | False | Mutable                               | ✅ Match |
| maintenancePolicy.maintenanceExclusion | False | False | Mutable                               | ✅ Match |
| maintenancePolicy.maintenanceExclusion[]<br>.endTime | False | False | Mutable                               | ✅ Match |
| maintenancePolicy.maintenanceExclusion[]<br>.exclusionName | False | False | Mutable                               | ✅ Match |
| maintenancePolicy.maintenanceExclusion[]<br>.exclusionOptions | False | False | Mutable                               | ✅ Match |
| maintenancePolicy.maintenanceExclusion[]<br>.exclusionOptions.scope | False | False | Mutable                               | ✅ Match |
| maintenancePolicy.maintenanceExclusion[]<br>.startTime | False | False | Mutable                               | ✅ Match |
| maintenancePolicy.recurringWindow | False | False | Mutable                               | ✅ Match |
| maintenancePolicy.recurringWindow.endTim<br>e | False | False | Mutable                               | ✅ Match |
| maintenancePolicy.recurringWindow.recurr<br>ence | False | False | Mutable                               | ✅ Match |
| maintenancePolicy.recurringWindow.startT<br>ime | False | False | Mutable                               | ✅ Match |
| masterAuth | False | False | Mutable                               | ✅ Match |
| masterAuth.clientCertificate | False | False | Mutable                               | ✅ Match |
| masterAuth.clientCertificateConfig | True | True | Mutable                               | ⚠️ KCC Stricter |
| masterAuth.clientCertificateConfig.issue<br>ClientCertificate | True | True | Mutable                               | ⚠️ KCC Stricter |
| masterAuth.clientKey | False | False | Mutable                               | ✅ Match |
| masterAuth.clusterCaCertificate | False | False | Mutable                               | ✅ Match |
| masterAuth.password | False | False | Mutable                               | ✅ Match |
| masterAuth.username | False | False | Mutable                               | ✅ Match |
| masterAuthorizedNetworksConfig | False | False | Mutable                               | ✅ Match |
| masterAuthorizedNetworksConfig.cidrBlock<br>s | False | False | Mutable                               | ✅ Match |
| masterAuthorizedNetworksConfig.cidrBlock<br>s[].cidrBlock | False | False | Mutable                               | ✅ Match |
| masterAuthorizedNetworksConfig.cidrBlock<br>s[].displayName | False | False | Mutable                               | ✅ Match |
| masterAuthorizedNetworksConfig.gcpPublic<br>CidrsAccessEnabled | False | False | Mutable                               | ✅ Match |
| masterVersion | N/A | False | Mutable                               | N/A |
| meshCertificates | False | False | Mutable                               | ✅ Match |
| meshCertificates.enableCertificates | False | False | Mutable                               | ✅ Match |
| minMasterVersion | False | False | N/A                                   | ❓ (No Internal Data) |
| monitoringConfig | False | False | Mutable                               | ✅ Match |
| monitoringConfig.advancedDatapathObserva<br>bilityConfig | False | False | Mutable                               | ✅ Match |
| monitoringConfig.advancedDatapathObserva<br>bilityConfig[].enableMetrics | False | False | Mutable                               | ✅ Match |
| monitoringConfig.advancedDatapathObserva<br>bilityConfig[].relayMode | False | False | Mutable                               | ✅ Match |
| monitoringConfig.enableComponents | False | False | Mutable                               | ✅ Match |
| monitoringConfig.managedPrometheus | False | False | Mutable                               | ✅ Match |
| monitoringConfig.managedPrometheus.enabl<br>ed | False | False | Mutable                               | ✅ Match |
| monitoringService | False | False | Mutable                               | ✅ Match |
| networkPolicy | False | False | Mutable                               | ✅ Match |
| networkPolicy.enabled | False | False | Mutable                               | ✅ Match |
| networkPolicy.provider | False | False | Mutable                               | ✅ Match |
| networkRef | False | True | Immutable                             | ❌ KCC Lax (Risk) |
| networkingMode | True | True | N/A                                   | ❓ (No Internal Data) |
| nodeConfig | True | True | Immutable                             | ✅ Match |
| nodeConfig.advancedMachineFeatures | True | True | Immutable                             | ✅ Match |
| nodeConfig.advancedMachineFeatures.enabl<br>eNestedVirtualization | True | True | Immutable                             | ✅ Match |
| nodeConfig.advancedMachineFeatures.threa<br>dsPerCore | True | True | Immutable                             | ✅ Match |
| nodeConfig.bootDiskKMSCryptoKeyRef | False | True | Immutable                             | ❌ KCC Lax (Risk) |
| nodeConfig.confidentialNodes | True | True | Immutable                             | ✅ Match |
| nodeConfig.confidentialNodes.enabled | True | True | Immutable                             | ✅ Match |
| nodeConfig.diskSizeGb | True | True | Immutable                             | ✅ Match |
| nodeConfig.diskType | True | True | Immutable                             | ✅ Match |
| nodeConfig.ephemeralStorageConfig | True | True | Immutable                             | ✅ Match |
| nodeConfig.ephemeralStorageConfig.localS<br>sdCount | True | True | Immutable                             | ✅ Match |
| nodeConfig.ephemeralStorageLocalSsdConfi<br>g | True | True | Immutable                             | ✅ Match |
| nodeConfig.ephemeralStorageLocalSsdConfi<br>g.localSsdCount | True | True | Immutable                             | ✅ Match |
| nodeConfig.fastSocket | False | False | Immutable                             | ❌ KCC Lax (Risk) |
| nodeConfig.fastSocket.enabled | False | False | Immutable                             | ❌ KCC Lax (Risk) |
| nodeConfig.gcfsConfig | True | N/A | Immutable                             | ✅ Match |
| nodeConfig.gcfsConfig.enabled | True | N/A | Immutable                             | ✅ Match |
| nodeConfig.guestAccelerator | True | True | Immutable                             | ✅ Match |
| nodeConfig.guestAccelerator[].count | True | True | Immutable                             | ✅ Match |
| nodeConfig.guestAccelerator[].gpuDriverI<br>nstallationConfig | True | True | Immutable                             | ✅ Match |
| nodeConfig.guestAccelerator[].gpuDriverI<br>nstallationConfig.gpuDriverVersion | True | True | Immutable                             | ✅ Match |
| nodeConfig.guestAccelerator[].gpuPartiti<br>onSize | True | True | Immutable                             | ✅ Match |
| nodeConfig.guestAccelerator[].gpuSharing<br>Config | True | True | Immutable                             | ✅ Match |
| nodeConfig.guestAccelerator[].gpuSharing<br>Config.gpuSharingStrategy | True | True | Immutable                             | ✅ Match |
| nodeConfig.guestAccelerator[].gpuSharing<br>Config.maxSharedClientsPerGpu | True | True | Immutable                             | ✅ Match |
| nodeConfig.guestAccelerator[].type | True | True | Immutable                             | ✅ Match |
| nodeConfig.gvnic | True | True | Immutable                             | ✅ Match |
| nodeConfig.gvnic.enabled | True | True | Immutable                             | ✅ Match |
| nodeConfig.hostMaintenancePolicy | True | True | Immutable                             | ✅ Match |
| nodeConfig.hostMaintenancePolicy.mainten<br>anceInterval | True | True | Immutable                             | ✅ Match |
| nodeConfig.imageType | False | False | Immutable                             | ❌ KCC Lax (Risk) |
| nodeConfig.kubeletConfig | False | False | Immutable                             | ❌ KCC Lax (Risk) |
| nodeConfig.kubeletConfig.cpuCfsQuota | False | False | Immutable                             | ❌ KCC Lax (Risk) |
| nodeConfig.kubeletConfig.cpuCfsQuotaPeri<br>od | False | False | Immutable                             | ❌ KCC Lax (Risk) |
| nodeConfig.kubeletConfig.cpuManagerPolic<br>y | False | False | Immutable                             | ❌ KCC Lax (Risk) |
| nodeConfig.kubeletConfig.podPidsLimit | False | False | Immutable                             | ❌ KCC Lax (Risk) |
| nodeConfig.labels | True | False | Immutable                             | ✅ Match |
| nodeConfig.linuxNodeConfig | False | False | Immutable                             | ❌ KCC Lax (Risk) |
| nodeConfig.linuxNodeConfig.cgroupMode | False | False | Immutable                             | ❌ KCC Lax (Risk) |
| nodeConfig.linuxNodeConfig.sysctls | False | False | Immutable                             | ❌ KCC Lax (Risk) |
| nodeConfig.localNvmeSsdBlockConfig | True | True | Immutable                             | ✅ Match |
| nodeConfig.localNvmeSsdBlockConfig.local<br>SsdCount | True | True | Immutable                             | ✅ Match |
| nodeConfig.localSsdCount | True | True | Immutable                             | ✅ Match |
| nodeConfig.loggingVariant | False | N/A | Immutable                             | ❌ KCC Lax (Risk) |
| nodeConfig.machineType | True | True | Immutable                             | ✅ Match |
| nodeConfig.metadata | True | True | Immutable                             | ✅ Match |
| nodeConfig.minCpuPlatform | True | True | Immutable                             | ✅ Match |
| nodeConfig.nodeGroupRef | True | True | Immutable                             | ✅ Match |
| nodeConfig.oauthScopes | True | True | Immutable                             | ✅ Match |
| nodeConfig.preemptible | True | True | Immutable                             | ✅ Match |
| nodeConfig.reservationAffinity | True | True | Immutable                             | ✅ Match |
| nodeConfig.reservationAffinity.consumeRe<br>servationType | True | True | Immutable                             | ✅ Match |
| nodeConfig.reservationAffinity.key | True | True | Immutable                             | ✅ Match |
| nodeConfig.reservationAffinity.values | True | True | Immutable                             | ✅ Match |
| nodeConfig.resourceLabels | False | False | Immutable                             | ❌ KCC Lax (Risk) |
| nodeConfig.sandboxConfig | True | True | Immutable                             | ✅ Match |
| nodeConfig.sandboxConfig.sandboxType | False | False | Immutable                             | ❌ KCC Lax (Risk) |
| nodeConfig.serviceAccountRef | False | True | Immutable                             | ❌ KCC Lax (Risk) |
| nodeConfig.shieldedInstanceConfig | True | True | Immutable                             | ✅ Match |
| nodeConfig.shieldedInstanceConfig.enable<br>IntegrityMonitoring | True | True | Immutable                             | ✅ Match |
| nodeConfig.shieldedInstanceConfig.enable<br>SecureBoot | True | True | Immutable                             | ✅ Match |
| nodeConfig.soleTenantConfig | True | True | Immutable                             | ✅ Match |
| nodeConfig.soleTenantConfig.nodeAffinity | True | True | Immutable                             | ✅ Match |
| nodeConfig.soleTenantConfig.nodeAffinity<br>[].key | True | True | Immutable                             | ✅ Match |
| nodeConfig.soleTenantConfig.nodeAffinity<br>[].operator | True | True | Immutable                             | ✅ Match |
| nodeConfig.soleTenantConfig.nodeAffinity<br>[].values | True | True | Immutable                             | ✅ Match |
| nodeConfig.spot | True | True | Immutable                             | ✅ Match |
| nodeConfig.tags | False | False | Immutable                             | ❌ KCC Lax (Risk) |
| nodeConfig.taint | False | False | Immutable                             | ❌ KCC Lax (Risk) |
| nodeConfig.taint[].effect | False | False | Immutable                             | ❌ KCC Lax (Risk) |
| nodeConfig.taint[].key | False | False | Immutable                             | ❌ KCC Lax (Risk) |
| nodeConfig.taint[].value | False | False | Immutable                             | ❌ KCC Lax (Risk) |
| nodeConfig.workloadMetadataConfig | True | False | Immutable                             | ✅ Match |
| nodeConfig.workloadMetadataConfig.mode | False | False | Immutable                             | ❌ KCC Lax (Risk) |
| nodeConfig.workloadMetadataConfig.nodeMe<br>tadata | False | False | Immutable                             | ❌ KCC Lax (Risk) |
| nodeLocations | False | False | N/A                                   | ❓ (No Internal Data) |
| nodePool | N/A | True | N/A                                   | ❓ (No Internal Data) |
| nodePoolAutoConfig | False | False | N/A                                   | ❓ (No Internal Data) |
| nodePoolAutoConfig.networkTags | False | False | Mutable                               | ✅ Match |
| nodePoolAutoConfig.networkTags.tags | False | False | Mutable                               | ✅ Match |
| nodePoolDefaults | False | False | N/A                                   | ❓ (No Internal Data) |
| nodePoolDefaults.nodeConfigDefaults | False | False | N/A                                   | ❓ (No Internal Data) |
| nodePoolDefaults.nodeConfigDefaults.gcfs<br>Config | False | False | N/A                                   | ❓ (No Internal Data) |
| nodePoolDefaults.nodeConfigDefaults.gcfs<br>Config.enabled | False | False | N/A                                   | ❓ (No Internal Data) |
| nodePoolDefaults.nodeConfigDefaults.logg<br>ingVariant | False | False | N/A                                   | ❓ (No Internal Data) |
| nodeVersion | False | False | Mutable                               | ✅ Match |
| notificationConfig | False | False | Mutable                               | ✅ Match |
| notificationConfig.pubsub | False | False | Mutable                               | ✅ Match |
| notificationConfig.pubsub.enabled | False | False | Mutable                               | ✅ Match |
| notificationConfig.pubsub.filter | False | False | Mutable                               | ✅ Match |
| notificationConfig.pubsub.filter.eventTy<br>pe | False | False | Mutable                               | ✅ Match |
| notificationConfig.pubsub.topicRef | False | False | Mutable                               | ✅ Match |
| operation | N/A | False | N/A                                   | ❓ (No Internal Data) |
| podSecurityPolicyConfig | False | False | Mutable                               | ✅ Match |
| podSecurityPolicyConfig.enabled | False | False | Mutable                               | ✅ Match |
| privateClusterConfig | False | False | N/A                                   | ❓ (No Internal Data) |
| privateClusterConfig.enablePrivateEndpoi<br>nt | False | False | N/A                                   | ❓ (No Internal Data) |
| privateClusterConfig.enablePrivateNodes | False | False | N/A                                   | ❓ (No Internal Data) |
| privateClusterConfig.masterGlobalAccessC<br>onfig | False | False | N/A                                   | ❓ (No Internal Data) |
| privateClusterConfig.masterGlobalAccessC<br>onfig.enabled | False | False | N/A                                   | ❓ (No Internal Data) |
| privateClusterConfig.masterIpv4CidrBlock | True | True | N/A                                   | ❓ (No Internal Data) |
| privateClusterConfig.peeringName | False | False | N/A                                   | ❓ (No Internal Data) |
| privateClusterConfig.privateEndpoint | False | False | N/A                                   | ❓ (No Internal Data) |
| privateClusterConfig.privateEndpointSubn<br>etworkRef | True | True | N/A                                   | ❓ (No Internal Data) |
| privateClusterConfig.publicEndpoint | False | False | N/A                                   | ❓ (No Internal Data) |
| privateIpv6GoogleAccess | False | False | Mutable                               | ✅ Match |
| project | N/A | True | N/A                                   | ❓ (No Internal Data) |
| protectConfig | False | False | N/A                                   | ❓ (No Internal Data) |
| protectConfig.workloadConfig | False | False | N/A                                   | ❓ (No Internal Data) |
| protectConfig.workloadConfig.auditMode | False | False | N/A                                   | ❓ (No Internal Data) |
| protectConfig.workloadVulnerabilityMode | False | False | N/A                                   | ❓ (No Internal Data) |
| releaseChannel | False | False | Mutable                               | ✅ Match |
| releaseChannel.channel | False | False | Mutable                               | ✅ Match |
| removeDefaultNodePool | N/A | False | N/A                                   | ❓ (No Internal Data) |
| resourceID | True | True | Immutable                             | ✅ Match |
| resourceLabels | N/A | False | Mutable                               | N/A |
| resourceUsageExportConfig | False | False | Mutable                               | ✅ Match |
| resourceUsageExportConfig.bigqueryDestin<br>ation | False | False | Mutable                               | ✅ Match |
| resourceUsageExportConfig.bigqueryDestin<br>ation.datasetId | False | False | Mutable                               | ✅ Match |
| resourceUsageExportConfig.enableNetworkE<br>gressMetering | False | False | Mutable                               | ✅ Match |
| resourceUsageExportConfig.enableResource<br>ConsumptionMetering | False | False | Mutable                               | ✅ Match |
| securityPostureConfig | False | False | Mutable                               | ✅ Match |
| securityPostureConfig.mode | False | False | Mutable                               | ✅ Match |
| securityPostureConfig.vulnerabilityMode | False | False | Mutable                               | ✅ Match |
| selfLink | N/A | False | Immutable                             | N/A |
| serviceExternalIpsConfig | False | False | Mutable                               | ✅ Match |
| serviceExternalIpsConfig.enabled | False | False | Mutable                               | ✅ Match |
| servicesIpv4Cidr | N/A | False | Immutable                             | N/A |
| subnetworkRef | False | True | Immutable                             | ❌ KCC Lax (Risk) |
| tpuIpv4CidrBlock | N/A | False | N/A                                   | ❓ (No Internal Data) |
| verticalPodAutoscaling | False | False | Mutable                               | ✅ Match |
| verticalPodAutoscaling.enabled | False | False | Mutable                               | ✅ Match |
| workloadIdentityConfig | False | False | Mutable                               | ✅ Match |
| workloadIdentityConfig.identityNamespace | False | N/A | Mutable                               | ✅ Match |
| workloadIdentityConfig.workloadPool | False | False | Mutable                               | ✅ Match |
