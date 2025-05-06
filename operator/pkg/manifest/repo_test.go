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

package manifest_test

import (
	"context"
	"path"
	"reflect"
	"strings"
	"testing"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/manifest"
	testpaths "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/util/paths"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/loaders"
)

var clusterModeWIManifest = `# Copyright 2022 Google LLC
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
---
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
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        type: object
    served: true
    storage: true
---
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

apiVersion: v1
kind: Namespace
metadata:
  name: cnrm-system
---
apiVersion: v1
kind: ServiceAccount
metadata:
  annotations:
    iam.gke.io/gcp-service-account: cnrm-system@${PROJECT_ID?}.iam.gserviceaccount.com
  name: cnrm-controller-manager
  namespace: cnrm-system
---
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
---
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
`

var clusterModeGcpManifest = `# Copyright 2022 Google LLC
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
---
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
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        type: object
    served: true
    storage: true
---
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

apiVersion: v1
kind: Namespace
metadata:
  name: cnrm-system
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: cnrm-controller-manager
  namespace: cnrm-system
---
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
---
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
`

var namespacedManifest = `# Copyright 2022 Google LLC
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
---
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
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        type: object
    served: true
    storage: true
---
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

apiVersion: v1
kind: Namespace
metadata:
  name: cnrm-system
`

var namespacedComponentsOnly = `# Copyright 2022 Google LLC
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

apiVersion: v1
kind: ServiceAccount
metadata:
  annotations:
    iam.gke.io/gcp-service-account: cnrm-system-${NAMESPACE?}@${PROJECT_ID?}.iam.gserviceaccount.com
  labels:
    cnrm.cloud.google.com/scoped-namespace: ${NAMESPACE?}
    cnrm.cloud.google.com/system: "true"
  name: cnrm-controller-manager-${NAMESPACE?}
  namespace: cnrm-system
---
apiVersion: v1
kind: Service
metadata:
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
---
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
`

func TestNewLocalRepository_LoadManifest(t *testing.T) {
	t.Parallel()
	basedir, repo := newTestNewLocalRepository()
	tests := []struct {
		name   string
		cc     *corev1beta1.ConfigConnector
		result map[string]string
	}{
		{
			name: "cluster mode, workload identity",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					Mode:                 "cluster",
				},
			},
			result: map[string]string{strings.Join([]string{basedir, "packages", "configconnector", "0.0.0-test", "cluster", "workload-identity"}, "/"): clusterModeWIManifest},
		},
		{
			name: "cluster mode, gcp identity",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					CredentialSecretName: "my-key",
					Mode:                 "cluster",
				},
			},
			result: map[string]string{strings.Join([]string{basedir, "packages", "configconnector", "0.0.0-test", "cluster", "gcp-identity"}, "/"): clusterModeGcpManifest},
		},
		{
			name: "namespaced mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					Mode:                 "namespaced",
				},
			},
			result: map[string]string{strings.Join([]string{basedir, "packages", "configconnector", "0.0.0-test", "namespaced"}, "/"): namespacedManifest},
		},
		{
			name: "namespaced mode by default",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
				},
			},
			result: map[string]string{strings.Join([]string{basedir, "packages", "configconnector", "0.0.0-test", "namespaced"}, "/"): namespacedManifest},
		},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			manifestStrs, err := repo.LoadManifest(context.TODO(), "configconnector", "0.0.0-test", tc.cc)
			if err != nil {
				t.Fatalf("unexpected error while loadding the manifest: %v", err)
			}
			if !reflect.DeepEqual(manifestStrs, tc.result) {
				t.Fatalf("unexpected diff: %v", cmp.Diff(manifestStrs, tc.result))
			}
		})
	}
}

func TestNewLocalRepository_LoadNamespacedComponents(t *testing.T) {
	t.Parallel()
	baseDir, repo := newTestNewLocalRepository()
	manifestPath := path.Join(baseDir, "packages/configconnector/0.0.0-test/namespaced/per-namespace-components.yaml")
	tests := []struct {
		name   string
		result map[string]string
	}{
		{
			name:   "load namespaced component",
			result: map[string]string{manifestPath: namespacedComponentsOnly},
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			manifests, err := repo.LoadNamespacedComponents(context.TODO(), "configconnector", "0.0.0-test")
			if err != nil {
				t.Fatalf("unexpected error while loadding the manifest for namespaced components: %v", err)
			}
			if !reflect.DeepEqual(manifests, tc.result) {
				t.Fatalf("unexpected diff: %v", cmp.Diff(manifests, tc.result))
			}
		})
	}
}

func TestNewLocalRepository_LoadChannel(t *testing.T) {
	t.Parallel()
	_, repo := newTestNewLocalRepository()
	res, err := repo.LoadChannel(context.TODO(), "stable")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := &loaders.Channel{
		Manifests: []loaders.Version{
			{
				Version: "0.0.0-test",
			},
		},
	}
	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("unexpected diff: %v", cmp.Diff(res, expected))
	}
}

func newTestNewLocalRepository() (string, manifest.Repository) {
	root := testpaths.GetOperatorSrcRootOrLogFatal()
	basedir := root + "/" + "pkg/manifest/testchannel"
	return basedir, manifest.NewLocalRepository(basedir)
}
