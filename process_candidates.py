# Copyright 2026 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import json
import subprocess
import sys
import re

candidates = [
    ("networkservices", "NetworkServicesGateway"),
    ("networkservices", "NetworkServicesGRPCRoute"),
    ("billingbudgets", "BillingBudgetsBudget"),
    ("dlp", "DLPJobTrigger"),
    ("recaptchaenterprise", "RecaptchaEnterpriseKey"),
    ("cloudfunctions", "CloudFunctionsFunction"),
    ("identityplatform", "IdentityPlatformTenant"),
    ("dlp", "DLPStoredInfoType"),
    ("cloudscheduler", "CloudSchedulerJob"),
    ("identityplatform", "IdentityPlatformConfig"),
    ("eventarc", "EventarcTrigger"),
    ("networkservices", "NetworkServicesTLSRoute"),
    ("binaryauthorization", "BinaryAuthorizationAttestor"),
    ("identityplatform", "IdentityPlatformTenantOAuthIDPConfig"),
    ("networkservices", "NetworkServicesHTTPRoute"),
    ("networkconnectivity", "NetworkConnectivitySpoke"),
    ("dataproc", "DataprocCluster"),
    ("configcontroller", "ConfigControllerInstance"),
    ("identityplatform", "IdentityPlatformOAuthIDPConfig"),
    ("filestore", "FilestoreInstance"),
    ("osconfig", "OSConfigOSPolicyAssignment"),
    ("filestore", "FilestoreBackup"),
    ("networkservices", "NetworkServicesEndpointPolicy"),
    ("dlp", "DLPDeidentifyTemplate"),
    ("networkservices", "NetworkServicesMesh"),
    ("networkservices", "NetworkServicesTCPRoute"),
    ("dataproc", "DataprocWorkflowTemplate"),
    ("datafusion", "DataFusionInstance"),
    ("osconfig", "OSConfigGuestPolicy"),
    ("dlp", "DLPInspectTemplate"),
    ("networkconnectivity", "NetworkConnectivityHub"),
    ("containeranalysis", "ContainerAnalysisNote"),
    ("binaryauthorization", "BinaryAuthorizationPolicy"),
    ("dataproc", "DataprocAutoscalingPolicy")
]

# Fetch all issues related to the task
try:
    out = subprocess.check_output(
        ['gh', 'issue', 'list', '--search', 'Create generate.sh and types.go files for', '--state', 'all', '--json', 'number,title,labels,state', '--limit', '1000'],
        text=True
    )
except subprocess.CalledProcessError as e:
    print("Error: Failed to fetch issues from GitHub. Please ensure you are authenticated by running 'gh auth login'.")
    sys.exit(1)

issues = json.loads(out)

required_labels = {"overseer", "area/direct", "priority/medium"}

# check existing issues
open_count = sum(1 for iss in issues if iss['state'] == 'OPEN')
print(f"Total open issues: {open_count}")

for group, kind in candidates:
    # Check if issue exists
    # case insensitive title matching because of variations like Networkconnectivity vs networkconnectivity
    pattern = re.compile(f"Create generate.sh and types.go files for {re.escape(group)} {re.escape(kind)}", re.IGNORECASE)
    matching_issues = [iss for iss in issues if pattern.search(iss['title'])]
    
    if matching_issues:
        for iss in matching_issues:
            labels = {label['name'] for label in iss.get('labels', [])}
            missing = required_labels - labels
            if missing:
                print(f"Injecting labels {missing} into issue #{iss['number']} for {group} {kind}")
                # Inject labels
                subprocess.check_call(['gh', 'issue', 'edit', str(iss['number']), '--add-label', ','.join(missing)])
            else:
                print(f"Issue #{iss['number']} for {group} {kind} already has required labels.")
    else:
        if open_count > 10:
            print(f"Skipping issue creation for {group} {kind}: > 10 open issues.")
        else:
            # Note: We would create an issue here, but for now this script just checks status and labels.
            print(f"Would create an issue for {group} {kind}, but creation logic is not implemented.")
