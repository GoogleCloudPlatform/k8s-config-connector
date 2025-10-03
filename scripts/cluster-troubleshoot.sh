#!/usr/bin/env bash
# Copyright 2022 Google LLC
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

# This script is intended as a diagnostic tool to:
#  - catch common issues with Config Connector
#  - extract diagnostic data that helps with debugging
#
# It assumes the following:
# - kubectl is configured to operate on the cluster running Config Connector
#
# Conventions:
# - section headers ARE ALL CAPS
#
# Example output for successful run:
#
#    KUBERNETES VERSION:
#    Client Version: v1.16.4
#    Server Version: v1.17.14-gke.1600
#
#    CONFIG CONNECTOR VERSION:
#    1.32.0
#
#    CONFIG CONNECTOR MODE:
#    namespaced
#
#    installation verified.

CONTACT_SUPPORT="
If you have followed troubleshooting instructions and are still unable to
resolve the issue, contact GCP Support (https://cloud.google.com/support-hub) or
file an issue on GitHub (https://github.com/GoogleCloudPlatform/k8s-config-connector/issues)
including the output of this script for further assistance."

function main {
    print_kubernetes_version
    echo
    print_config_connector_version
    echo
    print_config_connector_mode
    echo
    echo

    if ! installation_is_correct; then
        echo "$CONTACT_SUPPORT"
        return 1
    else
        echo "installation verified."
    fi
}

function print_kubernetes_version {
    echo "KUBERNETES VERSION:"
    kubectl version --short
}

function print_config_connector_version {
    echo "CONFIG CONNECTOR VERSION:"
    kubectl get ns cnrm-system -o jsonpath='{.metadata.annotations.cnrm\.cloud\.google\.com/version}'
    echo
}

function print_config_connector_mode {
    echo "CONFIG CONNECTOR MODE:"
    kubectl get ConfigConnector "configconnector.core.cnrm.cloud.google.com" -o=jsonpath="{@.spec.mode}"
}

function installation_is_correct {
    verify_cnrm_system_namespace_exists && \
    verify_config_connector_operator_installed && \
    verify_config_connector_kind_exists && \
    verify_cnrm_system_status
}

function verify_cnrm_system_namespace_exists {
    local error_message="
\"cnrm-system\" NAMESPACE NOT FOUND.

The cnrm-system namespace should be created as part of the installation process.
This namespace is where Config Connector controllers run.

to troubleshoot: run through the installation instructions again."

    if ! kubectl get namespace cnrm-system > /dev/null; then
        echo "$error_message"
        return 1
    fi
}

function verify_config_connector_kind_exists {
    local error_message="
ConfigConnector KIND NOT FOUND.

A resource of kind ConfigConnector must be applied in order for the Config Controller
operator to spawn it's resources, and start to reconcile resources.

remediation:

See \"Configuring Config Connector\":
  https://cloud.google.com/config-connector/docs/how-to/install-upgrade-uninstall#addon-configuring".

    if [ "$(kubectl get ConfigConnector)" = "" ]; then
        echo "$error_message"
        return 1
    fi
}

function verify_config_connector_operator_installed {
    local error_message="
CONFIG CONNECTOR OPERATOR NOT FOUND

As of 1.33, the operator (or the GKE add-on) is the only supported
method to install Config Connector.

remediation:

If you have previously installed and verified your Config Connector installation, upgrade to the config connector
operator as stated in https://cloud.google.com/config-connector/docs/how-to/advanced-install#upgrading_from_non-operator_installations.

Otherwise, you may not have installed Config Connector to the cluster.
See https://cloud.google.com/config-connector/docs/how-to/install-upgrade-uninstall."
    if [[ "$(kubectl get pod -n configconnector-operator-system | wc -l)" == 0 ]]; then
        echo "$error_message"
        return 1
    fi
}

function verify_cnrm_system_status {
    local error_message="
NON-RUNNING PODS DETECTED IN \"cnrm-system\" NAMESPACE

The cnrm-system namespace is where Config Connector resources
required to properly operate are installed.

If there are resources that are unhealthy in this namespace,
Config Connector will not operate properly.

to troubleshoot: examine the output below."
    local output
    output="$(kubectl get pods -n cnrm-system --field-selector status.phase!=Running 2>/dev/null)"
    if [ "$output" != "" ]; then
        echo \
"$error_message

non-running pods:
$output"
        return 1
    fi
}

main
