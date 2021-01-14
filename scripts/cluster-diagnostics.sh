#!/bin/bash
# This script is intended to gather diagnostic data
# that helps with debugging Config Connector
#
#
function main {
    echo "DIAGNOSTIC INFORMATION:
"
    print_kubernetes_version
    echo
    print_config_connector_version
    echo
    print_cnrm_system_status
    echo
    print_namespace_annotations
    echo
    # after this point, we have to do specific commands for namespace / non-namespace mode
    mode=$(kubectl get ConfigConnector -o jsonpath='{.items[0].spec.mode}')
    echo "CONFIG CONNECTOR MODE:
$mode"
    echo
    if [ "$mode" == "cluster" ] ; then
        print_cluster_mode_diagnostics
    elif [ "$mode" == "namespaced" ] ; then
        print_namespace_mode_diagnostics
    else
        echo "
invalid cluster mode: $mode. Please apply a ConfigConnector spec.mode that is one of
(cluster, namespaced)
"
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

function print_cnrm_system_status {
    echo "CNRM-SYSTEM RESOURCE STATUS:"
    kubectl get all -n cnrm-system
    echo
}

function print_namespace_annotations {
    echo "CONFIG CONNECTOR NAMESPACE ANNOTATIONS:
List of Kubernetes namespaces and the Config Connector project / folder / org annotations.
If no such annotations exist, Config Connector defaults to the project whose id is the same
as the namespace.
"

    kubectl get namespace -o jsonpath="
{range .items[*]}{.metadata.name}
    {'project-id: '}{.metadata.annotations['cnrm\.cloud\.google\.com/project-id']}
    {'folder-id: '}{.metadata.annotations['cnrm\.cloud\.google\.com/folder-id']}
    {'organization-id: '}{.metadata.annotations['cnrm\.cloud\.google\.com/organization-id']}
"
}

# CLUSTER MODE DIAGNOSTICS

function print_cluster_mode_diagnostics {
    echo "CLUSTER MODE DIAGNOSTICS

CONFIG CONNECTOR KIND:
"
    kubectl describe ConfigConnector
}

# NAMESPACE MODE DIAGNOSTICS

function print_namespace_mode_diagnostics {
    echo "NAMESPACE MODE DIAGNOSTICS

CONFIG CONNECTOR CONTEXTS:
"
    kubectl describe ConfigConnectorContext --all-namespaces
}

main