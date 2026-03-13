# Config Connector on OpenShift
Pre-requisites:
- gsutil: https://cloud.google.com/storage/docs/gsutil_install
- gcloud: https://cloud.google.com/sdk/docs/install
- oc + kubectl: https://access.redhat.com/documentation/en-us/openshift_container_platform/4.12/html-single/cli_tools/index#installing-openshift-cli

Install ConfigConnector, the following steps are based on the official documentation:
https://cloud.google.com/config-connector/docs/how-to/install-other-kubernetes#installing

If you are using existing service accounts, you can skip the step of creating a service account for ConfigConnector.

Applying credentials to cluster
```
kubectl create namespace cnrm-system
```

Create secret with the service account key. Note the secret name for the ConfigConnector resource creation.
```
kubectl create secret generic $SECRET_NAME \
    --from-file key.json \
    --namespace cnrm-system
```

```
gsutil cp gs://configconnector-operator/latest/release-bundle.tar.gz release-bundle.tar.gz
tar zxvf release-bundle.tar.gz
kubectl apply -f operator-system/configconnector-operator.yaml
```

## Patching Operator to run on OpenShift
We want to remove runAsUser and runAsGroup from the StatefulSet on OpenShift due to the fact that the default SCC does not allow these fields to be set.
We also want to set securityContext.capabilities.drop=["ALL"]
We can do this by applying the following patch:

```
kubectl patch statefulset.apps configconnector-operator \
  --namespace configconnector-operator-system \
  --type json \
  --patch '[{"op": "remove", "path": "/spec/template/spec/containers/0/securityContext/runAsUser"}, {"op": "remove", "path": "/spec/template/spec/containers/0/securityContext/runAsGroup"}, {"op": "add", "path": "/spec/template/spec/containers/0/securityContext/capabilities/drop", "value": ["ALL"]}]'
```
Note: could omit the patch if we create an SCC for the Operator SA.

### Create OpenShift SCC for ConfigConnector
Without modifying ConfigConnector Operator container images, it deploys a workload which uses 1000 UID which is disallowed on OpenShift. We will override these protections by assigning ServiceAccounts to SecurityContextConstraints that give high enough permissions.

SA names that needs anyuid SCC:
- cnrm-webhook-manager
- cnrm-controller-manager
- cnrm-deletiondefender

```
echo "cnrm-webhook-manager cnrm-controller-manager cnrm-deletiondefender" | xargs -n 1 \
  -I{} sh -c 'oc adm policy add-scc-to-user anyuid system:serviceaccount:cnrm-system:{} | sed "s/: /:/g"'
```

Validate results
```
oc get clusterrolebinding.rbac.authorization.k8s.io/system:openshift:scc:anyuid -ojsonpath='{.subjects}'
```

`cnrm-resource-stats-recorder` SA requires a combination of *anyuid* and *hostnetwork-v2* SCC. Since only the most permissive SCC takes effect, we need to create a new SCC that combines the two.

Create a new SCC with the following content:
```bash
echo "allowHostPorts: true # Modified from the anyuid SCC
priority: 10
requiredDropCapabilities:
  - MKNOD
allowPrivilegedContainer: false
runAsUser:
  type: RunAsAny
users:
  - system:serviceaccount:cnrm-system:cnrm-resource-stats-recorder
allowHostDirVolumePlugin: false
allowHostIPC: false
seLinuxContext:
  type: MustRunAs
readOnlyRootFilesystem: false
metadata:
  name: cnrm-resource-stats-recorder-scc
fsGroup:
  type: RunAsAny
groups:
  - 'system:cluster-admins'
kind: SecurityContextConstraints
defaultAddCapabilities: null
supplementalGroups:
  type: RunAsAny
volumes:
  - configMap
  - downwardAPI
  - emptyDir
  - ephemeral
  - persistentVolumeClaim
  - projected
  - seret
allowHostPID: false
allowHostNetwork: true # Modified from the anyuid SCC
allowPrivilegeEscalation: true
apiVersion: security.openshift.io/v1
allowedCapabilities: null
" | oc apply -f -
```

## [Configuring ConfigConnector](https://cloud.google.com/config-connector/docs/how-to/install-other-kubernetes#configuring)
You can now follow the official documentation to configure ConfigConnector.

## Monitoring for errors
```
oc get pods -n cnrm-system -oname | xargs -I {} -n 1 -P 10 sh -c "oc logs -f {} | grep --line-buffered error | sed \"s#^#{}: #\""
```

# Cleaup
```
echo "Uninstalling"
kubectl delete ConfigConnector configconnector.core.cnrm.cloud.google.com \
    --wait=true && \
kubectl delete -f operator-system/configconnector-operator.yaml  --wait=true && \
echo "Config Connector Operator uninstalled"
oc delete clusterrolebinding.rbac.authorization.k8s.io/openshift-configconnector-manager-rolebinding
oc delete clusterrole.rbac.authorization.k8s.io/openshift-configconnector-manager-role
```

# Troubleshooting

If you see [`is forbidden: cannot set blockOwnerDeletion if an ownerReference refers to a resource you can't set finalizers on`](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/434).

Add more permissions to the service account for OpenShift to resolve. The [PR has been merged](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/797)
```
echo "kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: openshift-configconnector-manager-role
rules:
  - verbs:
    - update
    apiGroups:
      - core.cnrm.cloud.google.com
    resources:
      - configconnectors/finalizers
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: openshift-configconnector-manager-rolebinding
subjects:
  - kind: ServiceAccount
    name: configconnector-operator
    namespace: configconnector-operator-system
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: openshift-configconnector-manager-role
" | oc apply -f -
```
