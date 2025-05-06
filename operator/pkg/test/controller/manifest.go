// Copyright 2022 Google LLC
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

package controller

import (
	"context"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cluster"

	"github.com/ghodss/yaml" //nolint:depguard
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

var NamespacedComponentsTemplate = []string{`
apiVersion: v1
kind: ServiceAccount
metadata:
  annotations:
    iam.gke.io/gcp-service-account: ${SERVICE_ACCOUNT?}
  labels:
    cnrm.cloud.google.com/scoped-namespace: ${NAMESPACE?}
    cnrm.cloud.google.com/system: "true"
  name: cnrm-controller-manager-${NAMESPACE?}
  namespace: cnrm-system
`, `
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  labels:
    cnrm.cloud.google.com/scoped-namespace: ${NAMESPACE?}
    cnrm.cloud.google.com/system: "true"
  name: cnrm-admin-binding-${NAMESPACE?}
  namespace: ${NAMESPACE?}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cnrm-admin
subjects:
- kind: ServiceAccount
  name: cnrm-controller-manager-${NAMESPACE?}
  namespace: cnrm-system
`, `
apiVersion: v1
kind: Service
metadata:
  labels:
    cnrm.cloud.google.com/monitored: "true"
    cnrm.cloud.google.com/scoped-namespace: ${NAMESPACE?}
    cnrm.cloud.google.com/system: "true"
  name: cnrm-manager-${NAMESPACE?}
  namespace: cnrm-system
spec:
  ports:
  - name: controller-manager
    port: 443
  - name: metrics
    port: 8888
  selector:
    cnrm.cloud.google.com/component: cnrm-controller-manager
    cnrm.cloud.google.com/scoped-namespace: ${NAMESPACE?}
    cnrm.cloud.google.com/system: "true"
`, `
apiVersion: apps/v1
kind: StatefulSet
metadata:
  labels:
    cnrm.cloud.google.com/component: cnrm-controller-manager
    cnrm.cloud.google.com/scoped-namespace: ${NAMESPACE?}
    cnrm.cloud.google.com/system: "true"
  name: cnrm-controller-manager-${NAMESPACE?}
  namespace: cnrm-system
spec:
  selector:
    matchLabels:
      cnrm.cloud.google.com/component: cnrm-controller-manager
      cnrm.cloud.google.com/scoped-namespace: ${NAMESPACE?}
      cnrm.cloud.google.com/system: "true"
  serviceName: cnrm-manager-${NAMESPACE?}
  template:
    metadata:
      labels:
        cnrm.cloud.google.com/component: cnrm-controller-manager
        cnrm.cloud.google.com/scoped-namespace: ${NAMESPACE?}
        cnrm.cloud.google.com/system: "true"
    spec:
      containers:
      - args: ["--scoped-namespace=${NAMESPACE?}", "--stderrthreshold=INFO", "--prometheus-scrape-endpoint=:8888"]
        command: ["/configconnector/manager"]
        image: gcr.io/gke-release/cnrm/controller:4af93f1
        name: manager
      - command: ["/monitor", "--source=configconnector:http://localhost:8888?whitelisted=reconcile_requests_total,reconcile_request_duration_seconds,reconcile_workers_total,reconcile_occupied_workers_total,internal_errors_total&customResourceType=k8s_container&customLabels[container_name]&customLabels[project_id]&customLabels[location]&customLabels[cluster_name]&customLabels[namespace_name]&customLabels[pod_name]", "--stackdriver-prefix=kubernetes.io/internal/addons"]
        image: gke.gcr.io/prometheus-to-sd:v0.11.12-gke.11
        name: prom-to-sd
`}

var PerNamespaceComponentsTemplate = []string{`
apiVersion: v1
kind: ServiceAccount
metadata:
  annotations:
    iam.gke.io/gcp-service-account: ${SERVICE_ACCOUNT?}
  finalizers:
    - configconnector.cnrm.cloud.google.com/finalizer
  labels:
    cnrm.cloud.google.com/scoped-namespace: ${NAMESPACE?}
    cnrm.cloud.google.com/system: "true"
    tenancy.gke.io/access-level: supervisor
    tenancy.gke.io/project: t1234
    tenancy.gke.io/tenant: t1234-tenant0
  name: cnrm-controller-manager-${NAMESPACE?}
  namespace: t1234-tenant0-supervisor
`, `
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  finalizers:
    - configconnector.cnrm.cloud.google.com/finalizer
  labels:
    cnrm.cloud.google.com/scoped-namespace: ${NAMESPACE?}
    cnrm.cloud.google.com/system: "true"
    tenancy.gke.io/access-level: supervisor
    tenancy.gke.io/project: t1234
    tenancy.gke.io/tenant: t1234-tenant0
  name: cnrm-admin-binding-${NAMESPACE?}
  namespace: ${NAMESPACE?}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cnrm-admin
subjects:
- kind: ServiceAccount
  name: cnrm-controller-manager-${NAMESPACE?}
  namespace: t1234-tenant0-supervisor
`, `
apiVersion: v1
kind: Service
metadata:
  labels:
    cnrm.cloud.google.com/monitored: "true"
    cnrm.cloud.google.com/scoped-namespace: ${NAMESPACE?}
    cnrm.cloud.google.com/system: "true"
    tenancy.gke.io/access-level: supervisor
    tenancy.gke.io/project: t1234
    tenancy.gke.io/tenant: t1234-tenant0
  name: cnrm-manager-${NAMESPACE?}
  namespace: t1234-tenant0-supervisor
spec:
  ports:
  - name: controller-manager
    port: 443
  - name: metrics
    port: 8888
  selector:
    cnrm.cloud.google.com/component: cnrm-controller-manager
    cnrm.cloud.google.com/scoped-namespace: ${NAMESPACE?}
    cnrm.cloud.google.com/system: "true"
`, `
apiVersion: apps/v1
kind: StatefulSet
metadata:
  finalizers:
    - configconnector.cnrm.cloud.google.com/finalizer
  labels:
    cnrm.cloud.google.com/component: cnrm-controller-manager
    cnrm.cloud.google.com/scoped-namespace: ${NAMESPACE?}
    cnrm.cloud.google.com/system: "true"
    tenancy.gke.io/access-level: supervisor
    tenancy.gke.io/project: t1234
    tenancy.gke.io/tenant: t1234-tenant0
  name: cnrm-controller-manager-${NAMESPACE?}
  namespace: t1234-tenant0-supervisor
spec:
  selector:
    matchLabels:
      cnrm.cloud.google.com/component: cnrm-controller-manager
      cnrm.cloud.google.com/scoped-namespace: ${NAMESPACE?}
      cnrm.cloud.google.com/system: "true"
  serviceName: cnrm-manager-${NAMESPACE?}
  template:
    metadata:
      finalizers:
        - configconnector.cnrm.cloud.google.com/finalizer
      labels:
        cnrm.cloud.google.com/component: cnrm-controller-manager
        cnrm.cloud.google.com/scoped-namespace: ${NAMESPACE?}
        cnrm.cloud.google.com/system: "true"
    spec:
      containers:
      - args: ["--scoped-namespace=${NAMESPACE?}", "--stderrthreshold=INFO", "--prometheus-scrape-endpoint=:8888"]
        command: ["/configconnector/manager"]
        image: gcr.io/gke-release/cnrm/controller:4af93f1
        name: manager
      - command: ["/monitor", "--source=configconnector:http://localhost:8888?whitelisted=reconcile_requests_total,reconcile_request_duration_seconds,reconcile_workers_total,reconcile_occupied_workers_total,internal_errors_total&customResourceType=k8s_container&customLabels[container_name]&customLabels[project_id]&customLabels[location]&customLabels[cluster_name]&customLabels[namespace_name]&customLabels[pod_name]", "--stackdriver-prefix=kubernetes.io/internal/addons"]
        image: gke.gcr.io/prometheus-to-sd:v0.11.12-gke.11
        name: prom-to-sd
`}

var FooCRD = `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.test.cnrm.cloud.google.com
  labels:
    cnrm.cloud.google.com/system: "true"
    cnrm.cloud.google.com/managed-by-kcc: "true"
spec:
  scope: Namespaced
  group: test.cnrm.cloud.google.com
  names:
    kind: foo
    plural: foos
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        type: object
    served: true
    storage: true
`

var barCRD = `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: bars.test.cnrm.cloud.google.com
  labels:
    cnrm.cloud.google.com/system: "true"
    cnrm.cloud.google.com/managed-by-kcc: "true"
spec:
  scope: Namespaced
  group: test.cnrm.cloud.google.com
  names:
    kind: bar
    plural: bars
  versions:
  - name: v1beta1
    schema:
      openAPIV3Schema:
        type: object
    served: true
    storage: true
`

var nonKCCCRD = `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: bars.test.nonkcc.cloud.google.com
  labels:
    cnrm.cloud.google.com/system: "true"
    cnrm.cloud.google.com/managed-by-kcc: "true"
spec:
  scope: Namespaced
  group: test.nonkcc.cloud.google.com
  names:
    kind: bar
    plural: bars
  versions:
  - name: v1beta1
    schema:
      openAPIV3Schema:
        type: object
    served: true
    storage: true
`

var defectiveCRD = `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: bars.test.nonkcc.cloud.google.com
  labels:
    cnrm.cloud.google.com/system: "true"
    cnrm.cloud.google.com/managed-by-kcc: "true"
spec:
  scope: Namespaced
  group: test.nonkcc.cloud.google.com
  names:
    kind: bar
    plural: bars
  versions: nil
`

var SystemNs = `apiVersion: v1
kind: Namespace
metadata:
  name: cnrm-system
`

var ClusterModeOnlyWorkloadIdentityComponents = []string{`
apiVersion: v1
kind: ServiceAccount
metadata:
  annotations:
    iam.gke.io/gcp-service-account: ${SERVICE_ACCOUNT?}
  name: cnrm-controller-manager
  namespace: cnrm-system
`, `
apiVersion: v1
kind: Service
metadata:
  name: cnrm-manager
  namespace: cnrm-system
spec:
  ports:
  - name: controller-manager
    port: 443
  - name: metrics
    port: 8888
  selector:
    cnrm.cloud.google.com/component: cnrm-controller-manager
    cnrm.cloud.google.com/system: "true"
`, `
apiVersion: apps/v1
kind: StatefulSet
metadata:
  labels:
    cnrm.cloud.google.com/component: cnrm-controller-manager
    cnrm.cloud.google.com/system: "true"
  name: cnrm-controller-manager
  namespace: cnrm-system
spec:
  selector:
    matchLabels:
      cnrm.cloud.google.com/component: cnrm-controller-manager
      cnrm.cloud.google.com/system: "true"
  serviceName: cnrm-manager
  template:
    metadata:
      labels:
        cnrm.cloud.google.com/component: cnrm-controller-manager
        cnrm.cloud.google.com/system: "true"
`}

var ClusterModeOnlyGCPComponents = []string{`
apiVersion: v1
kind: ServiceAccount
metadata:
  name: cnrm-controller-manager
  namespace: cnrm-system
`, `
apiVersion: v1
kind: Service
metadata:
  name: cnrm-manager
  namespace: cnrm-system
spec:
  ports:
  - name: controller-manager
    port: 443
  - name: metrics
    port: 8888
  selector:
    cnrm.cloud.google.com/component: cnrm-controller-manager
    cnrm.cloud.google.com/system: "true"
`, `
apiVersion: apps/v1
kind: StatefulSet
metadata:
  labels:
    cnrm.cloud.google.com/component: cnrm-controller-manager
    cnrm.cloud.google.com/system: "true"
  name: cnrm-controller-manager
  namespace: cnrm-system
spec:
  selector:
    matchLabels:
      cnrm.cloud.google.com/component: cnrm-controller-manager
      cnrm.cloud.google.com/system: "true"
  serviceName: cnrm-manager
  template:
    metadata:
      labels:
        cnrm.cloud.google.com/component: cnrm-controller-manager
        cnrm.cloud.google.com/system: "true"
    spec:
      volumes:
      - name: gcp-service-account
        secret:
          secretName: gcp-key
`}

var PerNamespaceControllerManagerPod = `apiVersion: v1
kind: Pod
metadata:
  labels:
    cnrm.cloud.google.com/component: cnrm-controller-manager
    cnrm.cloud.google.com/system: "true"
  name: cnrm-controller-manager-12345-0
  namespace: cnrm-system
spec:
  containers:
  - name: manager
    image: test-image
`

var NamespacedControllerManagerPod = `apiVersion: v1
kind: Pod
metadata:
  labels:
    cnrm.cloud.google.com/component: cnrm-controller-manager
    cnrm.cloud.google.com/system: "true"
  name: cnrm-controller-manager-12345-0
  namespace: t1234-tenant0-supervisor
spec:
  containers:
  - name: manager
    image: test-image
`

func GetSharedComponentsManifest() []string {
	res := make([]string, 0)
	res = append(res, FooCRD, SystemNs)
	return res
}

func GetManifestsWithAlphaAndBetaCRDs() []string {
	res := make([]string, 0)
	res = append(res, FooCRD, barCRD, SystemNs)
	return res
}

func GetManifestsWithAlphaCRD() []string {
	res := make([]string, 0)
	res = append(res, FooCRD, SystemNs)
	return res
}

func GetManifestsWithBetaCRD() []string {
	res := make([]string, 0)
	res = append(res, barCRD, SystemNs)
	return res
}

func GetManifestsWithNoCRD() []string {
	res := make([]string, 0)
	res = append(res, SystemNs)
	return res
}

func GetManifestsWithNonKCCCRD() []string {
	res := make([]string, 0)
	res = append(res, nonKCCCRD, FooCRD, SystemNs)
	return res
}

func GetManifestsWithDefectiveCRD() []string {
	res := make([]string, 0)
	res = append(res, defectiveCRD, SystemNs)
	return res
}

func GetClusterModeGCPManifest() []string {
	res := make([]string, 0)
	res = append(res, GetSharedComponentsManifest()...)
	res = append(res, ClusterModeOnlyGCPComponents...)
	return res
}

func GetClusterModeWorkloadIdentityManifest() []string {
	res := make([]string, 0)
	res = append(res, GetSharedComponentsManifest()...)
	res = append(res, ClusterModeOnlyWorkloadIdentityComponents...)
	return res
}

func GetPerNamespaceManifest() []string {
	res := make([]string, 0)
	res = append(res, NamespacedComponentsTemplate...)
	return res
}

func ManuallyReplaceGSA(components []string, saName string) []string {
	res := make([]string, 0)
	for _, s := range components {
		s = strings.ReplaceAll(s, "${SERVICE_ACCOUNT?}", saName)
		res = append(res, s)
	}
	return res
}

func ManuallyReplaceSecretVolume(components []string, secretName string) []string {
	res := make([]string, 0)
	for _, s := range components {
		s = strings.ReplaceAll(s, "gcp-key", secretName)
		res = append(res, s)
	}
	return res
}

func ManuallyModifyNamespaceTemplates(t *testing.T, template []string, nsName, saName string, userProjectOverride bool, billingProject string, c client.Client) []string {
	var res []string
	nsID, err := cluster.GetNamespaceID(context.TODO(), k8s.OperatorNamespaceIDConfigMapNN, c, nsName)
	if err != nil {
		t.Fatalf("error getting the id for namespace %v", err)
	}
	for _, s := range template {
		applied := s
		if strings.Contains(s, "kind: StatefulSet") {
			if billingProject != "" {
				applied = strings.ReplaceAll(applied,
					`args: ["--scoped-namespace=${NAMESPACE?}", "--stderrthreshold=INFO", "--prometheus-scrape-endpoint=:8888"`,
					`args: ["--scoped-namespace=${NAMESPACE?}", "--stderrthreshold=INFO", "--prometheus-scrape-endpoint=:8888", "--billing-project=`+billingProject+`"`,
				)
			}

			if userProjectOverride {
				applied = strings.ReplaceAll(applied,
					`args: ["--scoped-namespace=${NAMESPACE?}", "--stderrthreshold=INFO", "--prometheus-scrape-endpoint=:8888"`,
					`args: ["--scoped-namespace=${NAMESPACE?}", "--stderrthreshold=INFO", "--prometheus-scrape-endpoint=:8888", "--user-project-override=true"`,
				)
			}

			applied = strings.ReplaceAll(applied, "cnrm-controller-manager-${NAMESPACE?}", "cnrm-controller-manager-"+nsID)
			applied = strings.ReplaceAll(applied, k8s.NamespacedManagerServiceTmpl, k8s.NamespacedManagerServicePrefix+nsID)
		}
		if strings.Contains(s, "name: cnrm-manager-${NAMESPACE?}") {
			applied = strings.ReplaceAll(applied, k8s.NamespacedManagerServiceTmpl, k8s.NamespacedManagerServicePrefix+nsID)
		}
		applied = strings.ReplaceAll(applied, "${SERVICE_ACCOUNT?}", saName)
		applied = strings.ReplaceAll(applied, "${NAMESPACE?}", nsName)
		u := ToUnstructured(t, applied)
		labels := u.GetLabels()
		labels[k8s.ConfigConnectorContextNamespaceLabel] = nsName
		u.SetLabels(labels)
		applied = ToString(t, u)
		res = append(res, applied)
	}
	return res
}

func ToUnstructured(t *testing.T, objStr string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	b := []byte(objStr)
	err := yaml.Unmarshal(b, u)
	if err != nil {
		t.Fatalf("error unmarshalling bytes to unstruct: %v", err)
	}
	return u
}

func ToString(t *testing.T, u *unstructured.Unstructured) string {
	json, err := u.MarshalJSON()
	if err != nil {
		t.Fatalf("error marshalling unstructured to json: %v", err)
	}
	y, err := yaml.JSONToYAML(json)
	if err != nil {
		t.Fatalf("error converting json to yaml: %v", err)
	}
	return string(y)
}

func ParseObjects(ctx context.Context, t *testing.T, objects []string) *manifest.Objects {
	objs := strings.Join(objects, "---\n")
	m, err := manifest.ParseObjects(ctx, objs)
	if err != nil {
		t.Fatalf("while parsing objects: %v", err)
	}
	return m
}
