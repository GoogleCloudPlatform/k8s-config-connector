#!/bin/bash
candidates=(
  "binaryauthorization BinaryAuthorizationAttestor"
  "binaryauthorization BinaryAuthorizationPolicy"
  "cloudfunctions CloudFunctionsFunction"
  "cloudscheduler CloudSchedulerJob"
  "configcontroller ConfigControllerInstance"
  "containeranalysis ContainerAnalysisNote"
  "datafusion DataFusionInstance"
  "dataproc DataprocAutoscalingPolicy"
  "dataproc DataprocCluster"
  "dataproc DataprocWorkflowTemplate"
  "dlp DLPDeidentifyTemplate"
  "dlp DLPInspectTemplate"
  "dlp DLPJobTrigger"
  "dlp DLPStoredInfoType"
  "eventarc EventarcTrigger"
  "filestore FilestoreBackup"
  "filestore FilestoreInstance"
  "identityplatform IdentityPlatformConfig"
  "identityplatform IdentityPlatformOAuthIDPConfig"
  "identityplatform IdentityPlatformTenant"
  "identityplatform IdentityPlatformTenantOAuthIDPConfig"
  "networkconnectivity NetworkConnectivityHub"
  "networkconnectivity NetworkConnectivitySpoke"
  "networkservices NetworkServicesEndpointPolicy"
  "networkservices NetworkServicesGRPCRoute"
  "networkservices NetworkServicesGateway"
  "networkservices NetworkServicesHTTPRoute"
  "networkservices NetworkServicesMesh"
  "networkservices NetworkServicesTCPRoute"
  "networkservices NetworkServicesTLSRoute"
  "osconfig OSConfigGuestPolicy"
  "osconfig OSConfigOSPolicyAssignment"
  "recaptchaenterprise RecaptchaEnterpriseKey"
)

for c in "${candidates[@]}"; do
  group=$(echo $c | awk '{print $1}')
  kind=$(echo $c | awk '{print $2}')
  
  # check if issue exists
  output=$(gh issue list --search "in:title Create generate.sh and types.go files for $group $kind" --state all --json number,labels 2>/dev/null)
  
  if [ "$output" == "[]" ] || [ -z "$output" ]; then
    echo "NEED_CREATE $group $kind"
    exit 0
  else
    # Issue exists, check labels
    has_overseer=$(echo $output | grep '"name":"overseer"')
    has_priority=$(echo $output | grep '"name":"priority/medium"')
    has_area=$(echo $output | grep '"name":"area/direct"')
    
    number=$(echo $output | jq -r '.[0].number')
    
    missing_labels=""
    if [ -z "$has_overseer" ]; then missing_labels="overseer,$missing_labels"; fi
    if [ -z "$has_priority" ]; then missing_labels="priority/medium,$missing_labels"; fi
    if [ -z "$has_area" ]; then missing_labels="area/direct,$missing_labels"; fi
    
    if [ -n "$missing_labels" ]; then
      missing_labels=${missing_labels%,}
      echo "NEED_LABELS $number $missing_labels"
      # According to instructions, we can inject labels here and then skip to next
      gh issue edit $number --add-label "$missing_labels"
    fi
  fi
done
echo "DONE"
