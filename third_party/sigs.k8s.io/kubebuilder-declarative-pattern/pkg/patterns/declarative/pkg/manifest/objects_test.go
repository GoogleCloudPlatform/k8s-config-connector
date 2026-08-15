package manifest

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func Test_Object(t *testing.T) {
	tests := []struct {
		name           string
		inputManifest  string
		expectedObject []*Object
		expectedBlobs  []string
	}{
		{
			name: "simple applied manifest",
			inputManifest: `---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: foo-operator
  namespace: kube-system`,
			expectedObject: []*Object{
				{
					object: &unstructured.Unstructured{
						Object: map[string]interface{}{
							"apiVersion": "v1",
							"kind":       "ServiceAccount",
							"metadata": map[string]interface{}{
								"name":      "foo-operator",
								"namespace": "kube-system",
							},
						},
					},
				},
			},
			expectedBlobs: []string{},
		},
		{
			name: "simple kustomization manifest",
			inputManifest: `---
resources:
	- services.yaml
	- deployment.yaml
configMapGenerator:
- name: coredns
	namespace: kube-system
	files:
	- Corefile`,
			expectedObject: []*Object{},
			expectedBlobs: []string{
				`resources:
	- services.yaml
	- deployment.yaml
configMapGenerator:
- name: coredns
	namespace: kube-system
	files:
	- Corefile
`,
			},
		},
		{
			name: "a simple and kustomization manifest",
			inputManifest: `---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: foo-operator
  namespace: kube-system
---
resources:
	- services.yaml
	- deployment.yaml
configMapGenerator:
- name: coredns
	namespace: kube-system
	files:
	- Corefile`,
			expectedObject: []*Object{
				{
					object: &unstructured.Unstructured{
						Object: map[string]interface{}{
							"apiVersion": "v1",
							"kind":       "ServiceAccount",
							"metadata": map[string]interface{}{
								"name":      "foo-operator",
								"namespace": "kube-system",
							},
						},
					},
				},
			},
			expectedBlobs: []string{
				`resources:
	- services.yaml
	- deployment.yaml
configMapGenerator:
- name: coredns
	namespace: kube-system
	files:
	- Corefile
`,
			},
		},
		{
			name: "multi doc with comment",
			inputManifest: `--- # first one
apiVersion: v1
kind: ServiceAccount
metadata:
  name: foo-operator
  namespace: kube-system
--- # empty doc
# comments only
--- # second one
apiVersion: v1
kind: ServiceAccount
metadata:
  name: foo-operator
  namespace: kube-system`,
			expectedObject: []*Object{
				{
					object: &unstructured.Unstructured{
						Object: map[string]interface{}{
							"apiVersion": "v1",
							"kind":       "ServiceAccount",
							"metadata": map[string]interface{}{
								"name":      "foo-operator",
								"namespace": "kube-system",
							},
						},
					},
				},
				{
					object: &unstructured.Unstructured{
						Object: map[string]interface{}{
							"apiVersion": "v1",
							"kind":       "ServiceAccount",
							"metadata": map[string]interface{}{
								"name":      "foo-operator",
								"namespace": "kube-system",
							},
						},
					},
				},
			},
			expectedBlobs: []string{},
		},
		{
			name: "empty doc",
			inputManifest: `
`,
			expectedObject: []*Object{},
			expectedBlobs:  []string{},
		},
		{
			name: "empty objects",
			inputManifest: `
---
null
---
---
`,
			expectedObject: []*Object{},
			expectedBlobs:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			returnedObj, err := ParseObjects(ctx, tt.inputManifest)
			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}

			if len(tt.expectedObject) != len(returnedObj.Items) {
				t.Errorf("Expected length of %v to be %v but is %v", returnedObj.Items, len(tt.expectedObject),
					len(returnedObj.Items))
			}

			if len(tt.expectedBlobs) != len(returnedObj.Blobs) {
				t.Errorf("Expected length of %v to be %v but is %v", returnedObj.Blobs, len(tt.expectedBlobs),
					len(returnedObj.Blobs))
			}

			for i, actual := range returnedObj.Blobs {
				actualStr := string(actual)
				expectedStr := tt.expectedBlobs[i]
				if expectedStr != actualStr {
					t.Fatalf("unexpected result, expected ========\n%v\n\nactual ========\n%v\n", expectedStr, actualStr)
				}
			}

			for i, actual := range returnedObj.Items {
				actualBytes, err := actual.JSON()
				if err != nil {
					t.Fatalf("unexpected err: %v", err)
				}
				expectedBytes, err := tt.expectedObject[i].JSON()
				if err != nil {
					t.Fatalf("unexpected err: %v", err)
				}
				actualStr := string(actualBytes)
				expectedStr := string(expectedBytes)
				if expectedStr != actualStr {
					t.Fatalf("unexpected result, expected ========\n%v\n\nactual ========\n%v\n", expectedStr, actualStr)
				}
			}
		})
	}
}

func Test_AddLabels(t *testing.T) {
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
        name: busybox`
	tests := []struct {
		name           string
		inputManifest  string
		inputLabels    map[string]string
		expectedLabels map[string]string
	}{
		{
			name:           "add labels which are all new one",
			inputManifest:  inputManifest,
			inputLabels:    map[string]string{"sample-key1": "sample-value1", "sample-key2": "sample-value2"},
			expectedLabels: map[string]string{"app": "test-app", "sample-key1": "sample-value1", "sample-key2": "sample-value2"},
		},
		{
			// If call AddLabels with a key which has exists already, value will be overwritten.
			name:           "add label which has already exist in manifest",
			inputManifest:  inputManifest,
			inputLabels:    map[string]string{"app": "test-app2"},
			expectedLabels: map[string]string{"app": "test-app2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			objects, err := ParseObjects(ctx, tt.inputManifest)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			for _, o := range objects.Items {
				o.AddLabels(tt.inputLabels)
				if len(tt.expectedLabels) != len(o.object.GetLabels()) {
					t.Errorf("Expected length of labels to be %v but is %v", len(tt.expectedLabels), len(o.object.GetLabels()))
				}
				if diff := cmp.Diff(tt.expectedLabels, o.object.GetLabels()); diff != "" {
					t.Fatalf("result mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func Test_AddAnnotations(t *testing.T) {
	inputManifest := `---
apiVersion: apps/v1
kind: Deployment
metadata:
  annotations:
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
        name: busybox`
	tests := []struct {
		name                string
		inputManifest       string
		inputAnnotations    map[string]string
		expectedAnnotations map[string]string
	}{
		{
			name:                "add annotations which are all new one",
			inputManifest:       inputManifest,
			inputAnnotations:    map[string]string{"sample-key1": "sample-value1", "sample-key2": "sample-value2"},
			expectedAnnotations: map[string]string{"app": "test-app", "sample-key1": "sample-value1", "sample-key2": "sample-value2"},
		},
		{
			// If call AddAnnotations with a key which has exists already, value will be overwritten.
			name:                "add annotations which has already exist in manifest",
			inputManifest:       inputManifest,
			inputAnnotations:    map[string]string{"app": "test-app2"},
			expectedAnnotations: map[string]string{"app": "test-app2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			objects, err := ParseObjects(ctx, tt.inputManifest)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			for _, o := range objects.Items {
				o.AddAnnotations(tt.inputAnnotations)
				if len(tt.expectedAnnotations) != len(o.object.GetAnnotations()) {
					t.Errorf("Expected length of labels to be %v but is %v", len(tt.expectedAnnotations), len(o.object.GetAnnotations()))
				}
				if diff := cmp.Diff(tt.expectedAnnotations, o.object.GetAnnotations()); diff != "" {
					t.Fatalf("result mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func Test_ParseJSONToObject(t *testing.T) {
	tests := []struct {
		name           string
		inputManifest  string
		expectedObject *Object
		error          bool
	}{
		{
			name: "valid json manifest",
			inputManifest: `{
  "apiVersion": "v1",
  "kind": "ServiceAccount",
  "metadata": {
    "name": "foo-operator",
    "namespace": "kube-system"
  }
}`,
			expectedObject: &Object{
				object: &unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "v1",
						"kind":       "ServiceAccount",
						"metadata": map[string]interface{}{
							"name":      "foo-operator",
							"namespace": "kube-system",
						},
					},
				},
			},
			error: false,
		},
		{
			name: "parse json error will occur",
			inputManifest: `{
  "apiVersion": "v1",
  "kind": "ServiceAccount",
  "metadata": {
    "name": "foo-operator",
    "invalid-key": 
  }
}`,
			expectedObject: nil,
			error:          true,
		},
		{
			name: "unexpected type error will occur",
			inputManifest: `{
  "apiVersion": "v1",
  "invalid-kind": "ServiceAccount",
  "metadata": {
    "name": "foo-operator",
    "namespace": "kube-system"
  }
}`,
			expectedObject: nil,
			error:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.error == true {
				_, err := ParseJSONToObject([]byte(tt.inputManifest))
				if err == nil {
					t.Fatalf("expect error occur, but error doesn't occur")
				}
			} else {
				object, err := ParseJSONToObject([]byte(tt.inputManifest))
				if err != nil {
					t.Fatalf("unexpected err: %v", err)
				}
				actual, _ := object.object.MarshalJSON()
				expected, _ := tt.expectedObject.object.MarshalJSON()
				if string(actual) != string(expected) {
					t.Fatalf("unexpected result, expected ========\n%v\n\nactual ========\n%v\n", tt.expectedObject.object, object.object)
				}
			}
		})
	}
}

func Test_SetNestedStringMap(t *testing.T) {
	tests := []struct {
		name           string
		inputManifest  string
		inputMap       map[string]string
		expectedObject []*Object
	}{
		{
			name: "normal pattern",
			inputManifest: `---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: foo-operator
  namespace: kube-system`,
			inputMap: map[string]string{"foo": "bar"},
			expectedObject: []*Object{
				{
					object: &unstructured.Unstructured{
						Object: map[string]interface{}{
							"apiVersion": "v1",
							"kind":       "ServiceAccount",
							"metadata": map[string]interface{}{
								"name":      "foo-operator",
								"namespace": "kube-system",
								"labels": map[string]interface{}{
									"foo": "bar",
								},
							},
						},
					},
				},
			},
		},
		{
			name:          "nil object pattern",
			inputManifest: "",
			inputMap:      map[string]string{"foo": "bar"},
			expectedObject: []*Object{
				{
					object: &unstructured.Unstructured{
						Object: map[string]interface{}{
							"metadata": map[string]interface{}{
								"labels": map[string]interface{}{
									"foo": "bar",
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			if len(tt.inputManifest) != 0 {
				objects, err := ParseObjects(ctx, tt.inputManifest)
				if err != nil {
					t.Fatalf("unexpected err: %v", err)
				}
				for _, o := range objects.Items {
					o.SetNestedStringMap(tt.inputMap, "metadata", "labels")
					actualBytes, _ := o.JSON()
					actualStr := string(actualBytes)

					expectedBytes, _ := tt.expectedObject[0].JSON()
					expectedStr := string(expectedBytes)

					if expectedStr != actualStr {
						t.Fatalf("unexpected result, expected ========\n%v\n\nactual ========\n%v\n", expectedStr, actualStr)
					}
				}
			} else { // Test for object.Object == nil pattern
				o, _ := NewObject(&unstructured.Unstructured{})
				o.SetNestedStringMap(tt.inputMap, "metadata", "labels")
				actualBytes, _ := o.JSON()
				actualStr := string(actualBytes)

				expectedBytes, _ := tt.expectedObject[0].JSON()
				expectedStr := string(expectedBytes)

				if expectedStr != actualStr {
					t.Fatalf("unexpected result, expected ========\n%v\n\nactual ========\n%v\n", expectedStr, actualStr)
				}
			}
		})
	}
}

func Test_MutateContainers(t *testing.T) {
	tests := []struct {
		name           string
		inputManifest  string
		expectedObject []*Object
		error          bool
		errorString    string
		fn             func(map[string]interface{}) error
	}{
		{
			name: "normal success pattern",
			inputManifest: `---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend
  labels:
    app: test-app
spec:
  selector:
    matchLabels:
      app: guestbook
      tier: frontend
  replicas: 3
  template:
    metadata:
      labels:
        app: guestbook
        tier: frontend
    spec:
      containers:
      - name: php-redis
        image: gcr.io/google-samples/gb-frontend:v4`,
			expectedObject: []*Object{
				{
					object: &unstructured.Unstructured{
						Object: map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"labels": map[string]interface{}{
									"app": "test-app",
								},
								"name": "frontend",
							},
							"spec": map[string]interface{}{
								"replicas": 3,
								"selector": map[string]interface{}{
									"matchLabels": map[string]interface{}{
										"app":  "guestbook",
										"tier": "frontend",
									},
								},
								"template": map[string]interface{}{
									"metadata": map[string]interface{}{
										"labels": map[string]interface{}{
											"app":  "guestbook",
											"tier": "frontend",
										},
									},
									"spec": map[string]interface{}{
										"containers": []interface{}{
											map[string]interface{}{
												"image": "gcr.io/google-samples/gb-frontend:v4",
												"name":  "php-redis",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			error: false,
			fn: func(m map[string]interface{}) error {
				return nil
			},
		},
		{
			name: "normal success pattern without init container",
			inputManifest: `---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend
  labels:
    app: test-app
spec:
  selector:
    matchLabels:
      app: guestbook
      tier: frontend
  replicas: 3
  template:
    metadata:
      labels:
        app: guestbook
        tier: frontend
    spec:
      containers:
      - name: php-redis
        image: gcr.io/google-samples/gb-frontend:v4`,
			expectedObject: []*Object{
				{
					object: &unstructured.Unstructured{
						Object: map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"labels": map[string]interface{}{
									"app": "test-app",
								},
								"name": "frontend",
							},
							"spec": map[string]interface{}{
								"replicas": 3,
								"selector": map[string]interface{}{
									"matchLabels": map[string]interface{}{
										"app":  "guestbook",
										"tier": "frontend",
									},
								},
								"template": map[string]interface{}{
									"metadata": map[string]interface{}{
										"labels": map[string]interface{}{
											"app":  "guestbook",
											"tier": "frontend",
										},
									},
									"spec": map[string]interface{}{
										"containers": []interface{}{
											map[string]interface{}{
												"image": "mutated_image",
												"name":  "php-redis",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			error: false,
			fn: func(container map[string]interface{}) error {
				container["image"] = "mutated_image"
				return nil
			},
		},
		{
			name: "normal success pattern with init containers",
			inputManifest: `---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend
  labels:
    app: test-app
spec:
  selector:
    matchLabels:
      app: guestbook
      tier: frontend
  replicas: 3
  template:
    metadata:
      labels:
        app: guestbook
        tier: frontend
    spec:
      initContainers:
      - name: init
        image: gcr.io/google-samples/gb-frontend:v4
      containers:
      - name: php-redis
        image: gcr.io/google-samples/gb-frontend:v4`,
			expectedObject: []*Object{
				{
					object: &unstructured.Unstructured{
						Object: map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"labels": map[string]interface{}{
									"app": "test-app",
								},
								"name": "frontend",
							},
							"spec": map[string]interface{}{
								"replicas": 3,
								"selector": map[string]interface{}{
									"matchLabels": map[string]interface{}{
										"app":  "guestbook",
										"tier": "frontend",
									},
								},
								"template": map[string]interface{}{
									"metadata": map[string]interface{}{
										"labels": map[string]interface{}{
											"app":  "guestbook",
											"tier": "frontend",
										},
									},
									"spec": map[string]interface{}{
										"initContainers": []interface{}{
											map[string]interface{}{
												"image": "mutated_image",
												"name":  "init",
											},
										},
										"containers": []interface{}{
											map[string]interface{}{
												"image": "mutated_image",
												"name":  "php-redis",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			error: false,
			fn: func(container map[string]interface{}) error {
				container["image"] = "mutated_image"
				return nil
			},
		},
		{
			name: "object has no containers key",
			inputManifest: `---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend
  labels:
    app: test-app
spec:
  selector:
    matchLabels:
      app: guestbook
      tier: frontend
  replicas: 3
  template:
    metadata:
      labels:
        app: guestbook
        tier: frontend
    spec: invalid-value`,
			expectedObject: nil,
			error:          true,
			errorString:    "error reading containers: spec.template.spec.containers accessor error: invalid-value is of the type string, expected map[string]interface{}",
			fn: func(m map[string]interface{}) error {
				return nil
			},
		},
		{
			name:           "no manifest",
			inputManifest:  "",
			expectedObject: nil,
			error:          true,
			errorString:    "containers not found",
			fn: func(m map[string]interface{}) error {
				return nil
			},
		},
		{
			name: "object has no containers list",
			inputManifest: `---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend
  labels:
    app: test-app
spec:
  selector:
    matchLabels:
      app: guestbook
      tier: frontend
  replicas: 3
  template:
    metadata:
      labels:
        app: guestbook
        tier: frontend
    spec:
      containers: invalid-value`,
			expectedObject: nil,
			error:          true,
			errorString:    "containers was not a list",
			fn: func(m map[string]interface{}) error {
				return nil
			},
		},
		{
			name: "object has no containers normal structure",
			inputManifest: `---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend
  labels:
    app: test-app
spec:
  selector:
    matchLabels:
      app: guestbook
      tier: frontend
  replicas: 3
  template:
    metadata:
      labels:
        app: guestbook
        tier: frontend
    spec:
      containers:
      - dummy-value`,
			expectedObject: nil,
			error:          true,
			errorString:    "container was not an object",
			fn: func(m map[string]interface{}) error {
				return nil
			},
		},
		{
			name: "mutate function return error",
			inputManifest: `---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend
  labels:
    app: test-app
spec:
  selector:
    matchLabels:
      app: guestbook
      tier: frontend
  replicas: 3
  template:
    metadata:
      labels:
        app: guestbook
        tier: frontend
    spec:
      containers:
      - name: php-redis
        image: gcr.io/google-samples/gb-frontend:v4`,
			expectedObject: nil,
			error:          true,
			errorString:    "error occures in mutate function",
			fn: func(m map[string]interface{}) error {
				return fmt.Errorf("error occures in mutate function")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			objects, err := ParseObjects(ctx, tt.inputManifest)
			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			if tt.error == false {
				for _, o := range objects.Items {
					err = o.MutateContainers(tt.fn)
					actualBytes, _ := o.JSON()
					actualStr := string(actualBytes)

					expectedBytes, _ := tt.expectedObject[0].JSON()
					expectedStr := string(expectedBytes)

					if expectedStr != actualStr {
						t.Fatalf("unexpected result, expected ========\n%v\n\nactual ========\n%v\n", expectedStr, actualStr)
					}
				}
			} else {
				if tt.inputManifest == "" {
					o, _ := NewObject(&unstructured.Unstructured{})
					err = o.MutateContainers(tt.fn)
					if diff := cmp.Diff(tt.errorString, err.Error()); diff != "" {
						t.Errorf("error mismatch (-want +got):\n%s", diff)
					}

				} else {
					for _, o := range objects.Items {
						err = o.MutateContainers(tt.fn)
						if diff := cmp.Diff(tt.errorString, err.Error()); diff != "" {
							t.Errorf("error mismatch (-want +got):\n%s", diff)
						}
					}
				}
			}

		})
	}
}

func Test_MutatePodSpec(t *testing.T) {
	tests := []struct {
		name           string
		inputManifest  string
		expectedObject []*Object
		error          bool
		errorString    string
		fn             func(map[string]interface{}) error
	}{
		{
			name: "normal success pattern",
			inputManifest: `---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend
  labels:
    app: test-app
spec:
  selector:
    matchLabels:
      app: guestbook
      tier: frontend
  replicas: 3
  template:
    metadata:
      labels:
        app: guestbook
        tier: frontend
    spec:
      containers:
      - name: php-redis
        image: gcr.io/google-samples/gb-frontend:v4`,
			expectedObject: []*Object{
				{
					object: &unstructured.Unstructured{
						Object: map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"labels": map[string]interface{}{
									"app": "test-app",
								},
								"name": "frontend",
							},
							"spec": map[string]interface{}{
								"replicas": 3,
								"selector": map[string]interface{}{
									"matchLabels": map[string]interface{}{
										"app":  "guestbook",
										"tier": "frontend",
									},
								},
								"template": map[string]interface{}{
									"metadata": map[string]interface{}{
										"labels": map[string]interface{}{
											"app":  "guestbook",
											"tier": "frontend",
										},
									},
									"spec": map[string]interface{}{
										"containers": []interface{}{
											map[string]interface{}{
												"image": "gcr.io/google-samples/gb-frontend:v4",
												"name":  "php-redis",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			error: false,
		},
		{
			name: "object has no spec key in template key",
			inputManifest: `---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend
  labels:
    app: test-app
spec:
  selector:
    matchLabels:
      app: guestbook
      tier: frontend
  replicas: 3
  template: invalid-value`,
			expectedObject: nil,
			error:          true,
			errorString:    "error reading containers: spec.template.spec accessor error: invalid-value is of the type string, expected map[string]interface{}",
			fn: func(m map[string]interface{}) error {
				return nil
			},
		},
		{
			name:           "no manifest",
			inputManifest:  "",
			expectedObject: nil,
			error:          true,
			errorString:    "pod spec not found",
			fn: func(m map[string]interface{}) error {
				return nil
			},
		},
		{
			name: "object has no containers normal structure",
			inputManifest: `---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend
  labels:
    app: test-app
spec:
  selector:
    matchLabels:
      app: guestbook
      tier: frontend
  replicas: 3
  template:
    metadata:
      labels:
        app: guestbook
        tier: frontend
    spec: invalid-value`,
			expectedObject: nil,
			error:          true,
			errorString:    "pod spec was not an object",
			fn: func(m map[string]interface{}) error {
				return nil
			},
		},
		{
			name: "mutate function return error",
			inputManifest: `---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend
  labels:
    app: test-app
spec:
  selector:
    matchLabels:
      app: guestbook
      tier: frontend
  replicas: 3
  template:
    metadata:
      labels:
        app: guestbook
        tier: frontend
    spec:
      containers:
      - name: php-redis
        image: gcr.io/google-samples/gb-frontend:v4`,
			expectedObject: nil,
			error:          true,
			errorString:    "error occures in mutate function",
			fn: func(m map[string]interface{}) error {
				return fmt.Errorf("error occures in mutate function")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			objects, err := ParseObjects(ctx, tt.inputManifest)
			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			if tt.error == false {
				for _, o := range objects.Items {
					err = o.MutatePodSpec(func(m map[string]interface{}) error {
						return nil
					})
					actualBytes, _ := o.JSON()
					actualStr := string(actualBytes)

					expectedBytes, _ := tt.expectedObject[0].JSON()
					expectedStr := string(expectedBytes)

					if expectedStr != actualStr {
						t.Fatalf("unexpected result, expected ========\n%v\n\nactual ========\n%v\n", expectedStr, actualStr)
					}
				}
			} else {
				if tt.inputManifest == "" {
					o, _ := NewObject(&unstructured.Unstructured{})
					err = o.MutatePodSpec(tt.fn)
					if diff := cmp.Diff(tt.errorString, err.Error()); diff != "" {
						t.Errorf("error mismatch (-want +got):\n%s", diff)
					}
				} else {
					for _, o := range objects.Items {
						err = o.MutatePodSpec(tt.fn)
						if diff := cmp.Diff(tt.errorString, err.Error()); diff != "" {
							t.Errorf("error mismatch (-want +got):\n%s", diff)
						}
					}
				}
			}

		})
	}
}

func Test_Sort(t *testing.T) {
	deployment1 := &Object{
		object: &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "apps/v1",
				"kind":       "Deployment",
				"metadata": map[string]interface{}{
					"name": "frontend111",
				},
			},
		},
		name:  "frontend111",
		Kind:  "Deployment",
		Group: "apps",
	}
	deployment2 := &Object{
		object: &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "apps/v1",
				"kind":       "Deployment",
				"metadata": map[string]interface{}{
					"name": "frontend22222",
				},
			},
		},
		name:  "frontend22222",
		Kind:  "Deployment",
		Group: "apps",
	}
	service := &Object{
		object: &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "v1",
				"kind":       "Service",
				"metadata": map[string]interface{}{
					"name": "frontend-service",
				},
			},
		},
		name:  "frontend-service",
		Kind:  "Service",
		Group: "",
	}
	serviceAccount := &Object{
		object: &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "v1",
				"kind":       "ServiceAccount",
				"metadata": map[string]interface{}{
					"name": "serviceaccount",
				},
			},
		},
		name:  "serviceaccount",
		Kind:  "ServiceAccount",
		Group: "",
	}
	tests := []struct {
		name            string
		inputObjects    *Objects
		expectedObjects *Objects
		error           bool
		scoreFunc       func(*Object) int
	}{
		{
			name: "sort with score function's result",
			inputObjects: &Objects{
				Items: []*Object{
					deployment2,
					deployment1,
				},
			},
			expectedObjects: &Objects{
				Items: []*Object{
					deployment1,
					deployment2,
				},
			},
			error:     false,
			scoreFunc: func(o *Object) int { return len(o.GetName()) },
		},
		{
			name: "sort with Group",
			inputObjects: &Objects{
				Items: []*Object{
					deployment1,
					service,
				},
			},
			expectedObjects: &Objects{
				Items: []*Object{
					service,
					deployment1,
				},
			},
			error:     false,
			scoreFunc: func(o *Object) int { return 0 },
		},
		{
			name: "sort with Kind",
			inputObjects: &Objects{
				Items: []*Object{
					serviceAccount,
					service,
				},
			},
			expectedObjects: &Objects{
				Items: []*Object{
					service,
					serviceAccount,
				},
			},
			error:     false,
			scoreFunc: func(o *Object) int { return 0 },
		},
		{
			name: "sort with Name",
			inputObjects: &Objects{
				Items: []*Object{
					deployment2,
					deployment1,
				},
			},
			expectedObjects: &Objects{
				Items: []*Object{
					deployment1,
					deployment2,
				},
			},
			error:     false,
			scoreFunc: func(o *Object) int { return 0 },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.inputObjects.Sort(tt.scoreFunc)
			if diff := cmp.Diff(tt.expectedObjects, tt.inputObjects, cmpopts.IgnoreUnexported(Object{})); diff != "" {
				t.Errorf("objects mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
