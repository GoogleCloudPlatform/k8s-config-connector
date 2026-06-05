/*
Copyright 2024 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package declarative

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

func Test_ImageRegistryTransform(t *testing.T) {
	inputManifest := `---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: test-app
  name: frontend
spec:
  replicas: 1
  selector:
    matchLabels:
      app: test-app
  strategy: {}
  template:
    metadata:
      labels:
        app: test-app
    spec:
      containers:
      - image: busybox
        name: busybox
---
apiVersion: batch/v1
kind: CronJob
metadata:
  name: hello
spec:
  schedule: "* * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: hello
            image: busybox:1.28
            imagePullPolicy: IfNotPresent
            command:
            - /bin/sh
            - -c
            - date; echo Hello from the Kubernetes cluster
          restartPolicy: OnFailure`
	var testCases = []struct {
		name             string
		registry         string
		imagePullSecret  string
		inputManifest    string
		expectedManifest string
	}{
		{
			name:            "replace registry only",
			registry:        "gcr.io/foo/bar",
			imagePullSecret: "",
			inputManifest:   inputManifest,
			expectedManifest: `---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: test-app
  name: frontend
spec:
  replicas: 1
  selector:
    matchLabels:
      app: test-app
  strategy: {}
  template:
    metadata:
      labels:
        app: test-app
    spec:
      containers:
      - image: gcr.io/foo/bar/busybox
        name: busybox
---
apiVersion: batch/v1
kind: CronJob
metadata:
  name: hello
spec:
  schedule: "* * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: hello
            image: gcr.io/foo/bar/busybox:1.28
            imagePullPolicy: IfNotPresent
            command:
            - /bin/sh
            - -c
            - date; echo Hello from the Kubernetes cluster
          restartPolicy: OnFailure`,
		},
		{
			name:            "replace imagePullSecrets only",
			registry:        "",
			imagePullSecret: "some-secret",
			inputManifest:   inputManifest,
			expectedManifest: `---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: test-app
  name: frontend
spec:
  replicas: 1
  selector:
    matchLabels:
      app: test-app
  strategy: {}
  template:
    metadata:
      labels:
        app: test-app
    spec:
      containers:
      - image: busybox
        name: busybox
      imagePullSecrets:
      - name: some-secret
---
apiVersion: batch/v1
kind: CronJob
metadata:
  name: hello
spec:
  schedule: "* * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: hello
            image: busybox:1.28
            imagePullPolicy: IfNotPresent
            command:
            - /bin/sh
            - -c
            - date; echo Hello from the Kubernetes cluster
          imagePullSecrets:
          - name: some-secret
          restartPolicy: OnFailure`,
		},
		{
			name:            "replace registry and imagePullSecrets",
			registry:        "gcr.io/foo/bar",
			imagePullSecret: "some-secret",
			inputManifest:   inputManifest,
			expectedManifest: `---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: test-app
  name: frontend
spec:
  replicas: 1
  selector:
    matchLabels:
      app: test-app
  strategy: {}
  template:
    metadata:
      labels:
        app: test-app
    spec:
      containers:
      - image: gcr.io/foo/bar/busybox
        name: busybox
      imagePullSecrets:
      - name: some-secret
---
apiVersion: batch/v1
kind: CronJob
metadata:
  name: hello
spec:
  schedule: "* * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: hello
            image: gcr.io/foo/bar/busybox:1.28
            imagePullPolicy: IfNotPresent
            command:
            - /bin/sh
            - -c
            - date; echo Hello from the Kubernetes cluster
          imagePullSecrets:
          - name: some-secret
          restartPolicy: OnFailure`,
		},
		{
			name:             "replace nothing",
			registry:         "",
			imagePullSecret:  "",
			inputManifest:    inputManifest,
			expectedManifest: inputManifest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			dummyDeclarative := &TestResource{
				TypeMeta: metav1.TypeMeta{
					Kind:       "TestResource",
					APIVersion: "addons.example.org/v1alpha1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-instance",
				},
			}

			ctx := context.Background()

			objects, err := manifest.ParseObjects(ctx, tc.inputManifest)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			fn := ImageRegistryTransform(tc.registry, tc.imagePullSecret)
			err = fn(ctx, dummyDeclarative, objects)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			expectedObjects, err := manifest.ParseObjects(ctx, tc.expectedManifest)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if len(expectedObjects.Items) != len(objects.Items) {
				t.Fatal("expected number of objects does not equal number of objects")
			}

			for idx := range expectedObjects.Items {
				diff := cmp.Diff(
					expectedObjects.Items[idx].UnstructuredObject().Object,
					objects.Items[idx].UnstructuredObject().Object)
				if diff != "" {
					t.Errorf("result mismatch (-want +got):\n%s", diff)
				}
			}

		})
	}
}
