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
	"fmt"
	"strings"

	customizev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/customize/v1beta1"

	"google.golang.org/protobuf/proto"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	ControllerResourceCRForControllerManagerResources = &customizev1beta1.ControllerResource{
		ObjectMeta: metav1.ObjectMeta{
			Name: "cnrm-controller-manager",
		},
		Spec: customizev1beta1.ControllerResourceSpec{
			Containers: []customizev1beta1.ContainerResourceSpec{
				{
					Name: "manager",
					Resources: customizev1beta1.ResourceRequirements{
						Limits: corev1.ResourceList{
							corev1.ResourceCPU: resource.MustParse("400m"),
						},
						Requests: corev1.ResourceList{
							corev1.ResourceMemory: resource.MustParse("512Mi"),
						},
					},
				},
			},
		},
	}
	ControllerResourceCRForObservedControllerManagerResources = &customizev1beta1.ControllerResource{
		ObjectMeta: metav1.ObjectMeta{
			Name: "cnrm-controller-manager",
		},
		Spec: customizev1beta1.ControllerResourceSpec{
			Containers: []customizev1beta1.ContainerResourceSpec{
				{
					Name: "manager",
					Resources: customizev1beta1.ResourceRequirements{
						Limits: corev1.ResourceList{
							corev1.ResourceCPU: resource.MustParse("400m"),
						},
						Requests: corev1.ResourceList{
							corev1.ResourceMemory: resource.MustParse("512Mi"),
						},
					},
				},
			},
		},
	}
	ControllerResourceCRForControllerManagerReplicas = &customizev1beta1.ControllerResource{
		ObjectMeta: metav1.ObjectMeta{
			Name: "cnrm-controller-manager",
		},
		Spec: customizev1beta1.ControllerResourceSpec{
			Replicas: proto.Int64(int64(4)),
			Containers: []customizev1beta1.ContainerResourceSpec{
				{
					Name: "manager",
					Resources: customizev1beta1.ResourceRequirements{
						Limits:   corev1.ResourceList{},
						Requests: corev1.ResourceList{},
					},
				},
			},
		},
	}
	ControllerResourceCRForWebhookManagerResourcesAndReplicas = &customizev1beta1.ControllerResource{
		ObjectMeta: metav1.ObjectMeta{
			Name: "cnrm-webhook-manager",
		},
		Spec: customizev1beta1.ControllerResourceSpec{
			Replicas: proto.Int64(int64(4)),
			Containers: []customizev1beta1.ContainerResourceSpec{
				{
					Name: "webhook",
					Resources: customizev1beta1.ResourceRequirements{
						Limits: corev1.ResourceList{
							corev1.ResourceMemory: resource.MustParse("512Mi"),
						},
						Requests: corev1.ResourceList{
							corev1.ResourceMemory: resource.MustParse("256Mi"),
						},
					},
				},
			},
		},
	}
	ControllerResourceCRForWebhookManagerWithLargeReplicas = &customizev1beta1.ControllerResource{
		ObjectMeta: metav1.ObjectMeta{
			Name: "cnrm-webhook-manager",
		},
		Spec: customizev1beta1.ControllerResourceSpec{
			Replicas: proto.Int64(int64(30)), // this value is larger than the default value of "maxReplicas" of HPA in KCC's manifests
			Containers: []customizev1beta1.ContainerResourceSpec{
				{
					Name: "webhook",
					Resources: customizev1beta1.ResourceRequirements{
						Limits: corev1.ResourceList{
							corev1.ResourceMemory: resource.MustParse("512Mi"),
						},
						Requests: corev1.ResourceList{
							corev1.ResourceMemory: resource.MustParse("256Mi"),
						},
					},
				},
			},
		},
	}
	NamespacedControllerResourceCRForControllerManagerResources = &customizev1beta1.NamespacedControllerResource{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cnrm-controller-manager",
			Namespace: "foo-ns",
		},
		Spec: customizev1beta1.NamespacedControllerResourceSpec{
			Containers: []customizev1beta1.ContainerResourceSpec{
				{
					Name: "manager",
					Resources: customizev1beta1.ResourceRequirements{
						Limits: corev1.ResourceList{
							corev1.ResourceCPU: resource.MustParse("400m"),
						},
						Requests: corev1.ResourceList{
							corev1.ResourceMemory: resource.MustParse("512Mi"),
						},
					},
				},
			},
		},
	}
	NamespacedControllerReconcilerCR = &customizev1beta1.NamespacedControllerReconciler{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cnrm-controller-manager",
			Namespace: "foo-ns",
		},
		Spec: customizev1beta1.NamespacedControllerReconcilerSpec{
			RateLimit: &customizev1beta1.RateLimit{
				Burst: 30,
				QPS:   80,
			},
		},
	}
	ControllerReconcilerCR = &customizev1beta1.ControllerReconciler{
		ObjectMeta: metav1.ObjectMeta{
			Name: "cnrm-controller-manager",
		},
		Spec: customizev1beta1.ControllerReconcilerSpec{
			RateLimit: &customizev1beta1.RateLimit{
				Burst: 30,
				QPS:   80,
			},
		},
	}
)

var (
	nonExistingControllerName                    = "controller-does-not-exist"
	ControllerResourceCRForNonExistingController = &customizev1beta1.ControllerResource{
		ObjectMeta: metav1.ObjectMeta{
			Name: nonExistingControllerName,
		},
		Spec: customizev1beta1.ControllerResourceSpec{
			Containers: []customizev1beta1.ContainerResourceSpec{},
		},
	}
	NamespacedControllerResourceCRForNonExistingController = &customizev1beta1.NamespacedControllerResource{
		ObjectMeta: metav1.ObjectMeta{
			Name:      nonExistingControllerName,
			Namespace: "foo-ns",
		},
		Spec: customizev1beta1.NamespacedControllerResourceSpec{
			Containers: []customizev1beta1.ContainerResourceSpec{},
		},
	}
	ErrNonExistingController = fmt.Sprintf("resource customization for controller %s is not supported", nonExistingControllerName)
)

var (
	nonExistingContainerName                    = "recorder" // there is no "recorder" container in "cnrm-controller-manager".
	ControllerResourceCRForNonExistingContainer = &customizev1beta1.ControllerResource{
		ObjectMeta: metav1.ObjectMeta{
			Name: "cnrm-controller-manager",
		},
		Spec: customizev1beta1.ControllerResourceSpec{
			Containers: []customizev1beta1.ContainerResourceSpec{
				{
					Name: nonExistingContainerName,
				},
			},
		},
	}
	NamespacedControllerResourceCRForNonExistingContainer = &customizev1beta1.NamespacedControllerResource{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cnrm-controller-manager",
			Namespace: "foo-ns",
		},
		Spec: customizev1beta1.NamespacedControllerResourceSpec{
			Containers: []customizev1beta1.ContainerResourceSpec{
				{
					Name: nonExistingContainerName,
				},
			},
		},
	}
	ErrNonExistingContainer = fmt.Sprintf("failed to apply customization cnrm-controller-manager: resource customization failed for the following containers because there are no matching containers in the manifest: %s", nonExistingContainerName)
)

var (
	ControllerResourceCRForDuplicatedContainer = &customizev1beta1.ControllerResource{
		ObjectMeta: metav1.ObjectMeta{
			Name: "cnrm-controller-manager",
		},
		Spec: customizev1beta1.ControllerResourceSpec{
			Containers: []customizev1beta1.ContainerResourceSpec{
				{
					Name: "manager", // a valid container name
				},
				{
					Name: "prom-to-sd", // another valid container name
				},
				{
					Name: "manager", // a valid container name but duplicated
				},
			},
		},
	}
	ErrDuplicatedContainer = fmt.Sprintf("failed to apply customization cnrm-controller-manager: the following containers are specified multiple times in the Spec: manager")
)

var (
	NamespacedControllerResourceCRWrongNamespace = &customizev1beta1.NamespacedControllerResource{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cnrm-controller-manager",
			Namespace: "does-not-match",
		},
		Spec: customizev1beta1.NamespacedControllerResourceSpec{
			Containers: []customizev1beta1.ContainerResourceSpec{
				{
					Name: "manager",
					Resources: customizev1beta1.ResourceRequirements{
						Limits: corev1.ResourceList{
							corev1.ResourceCPU: resource.MustParse("400m"),
						},
						Requests: corev1.ResourceList{
							corev1.ResourceMemory: resource.MustParse("512Mi"),
						},
					},
				},
			},
		},
	}
	NamespacedControllerReconcilerCRWrongNamespace = &customizev1beta1.NamespacedControllerReconciler{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cnrm-controller-manager",
			Namespace: "does-not-match",
		},
		Spec: customizev1beta1.NamespacedControllerReconcilerSpec{
			RateLimit: &customizev1beta1.RateLimit{
				Burst: 30,
				QPS:   80,
			},
		},
	}
)

var (
	ValidatingWebhookCRForDuplicatedWebhook = &customizev1beta1.ValidatingWebhookConfigurationCustomization{
		ObjectMeta: metav1.ObjectMeta{
			Name: "validating-webhook",
		},
		Spec: customizev1beta1.WebhookConfigurationCustomizationSpec{
			Webhooks: []customizev1beta1.WebhookCustomizationSpec{
				{
					Name: "deny-immutable-field-updates", // a valid webhook name
				},
				{
					Name: "resource-validation", // another valid webhook name
				},
				{
					Name: "deny-immutable-field-updates", // a valid webhook name but duplicated
				},
			},
		},
	}
	MutatingWebhookCRForDuplicatedWebhook = &customizev1beta1.MutatingWebhookConfigurationCustomization{
		ObjectMeta: metav1.ObjectMeta{
			Name: "mutating-webhook",
		},
		Spec: customizev1beta1.WebhookConfigurationCustomizationSpec{
			Webhooks: []customizev1beta1.WebhookCustomizationSpec{
				{
					Name: "container-annotation-handler", // a valid webhook name
				},
				{
					Name: "iam-defaulter", // another valid webhook name
				},
				{
					Name: "container-annotation-handler", // a valid webhook name but duplicated
				},
			},
		},
	}
	ErrDuplicatedWebhookForValidatingWebhookCR = fmt.Sprintf("invalid webhook configuration customization: the following webhooks are specified multiple times in the Spec: deny-immutable-field-updates")
	ErrDuplicatedWebhookForMutatingWebhookCR   = fmt.Sprintf("invalid webhook configuration customization: the following webhooks are specified multiple times in the Spec: container-annotation-handler")
)

var (
	unsupportedControllerName                                = "cnrm-webhook-manager" // a valid KCC controller but its rate limit customization is not currently supported.
	NamespacedControllerReconcilerCRForUnsupportedController = &customizev1beta1.NamespacedControllerReconciler{
		ObjectMeta: metav1.ObjectMeta{
			Name:      unsupportedControllerName,
			Namespace: "foo-ns",
		},
		Spec: customizev1beta1.NamespacedControllerReconcilerSpec{
			RateLimit: &customizev1beta1.RateLimit{
				Burst: 30,
				QPS:   80,
			},
		},
	}
	ControllerReconcilerCRForUnsupportedController = &customizev1beta1.ControllerReconciler{
		ObjectMeta: metav1.ObjectMeta{
			Name: unsupportedControllerName,
		},
		Spec: customizev1beta1.ControllerReconcilerSpec{
			RateLimit: &customizev1beta1.RateLimit{
				Burst: 30,
				QPS:   80,
			},
		},
	}
	ErrUnsupportedController = fmt.Sprintf("failed to apply rate limit customization %s: "+
		"rate limit customization for %s is not supported. "+
		"Supported controllers: %s",
		unsupportedControllerName, unsupportedControllerName, strings.Join(customizev1beta1.ValidRateLimitControllers, ", "))
)

var ClusterModeComponents = []string{`
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
    spec:
      containers:
      - args: ["--scoped-namespace=${NAMESPACE?}", "--stderrthreshold=INFO", "--prometheus-scrape-endpoint=:8888"]
        command: ["/configconnector/manager"]
        image: gcr.io/gke-release/cnrm/controller:4af93f1
        name: manager
        resources:
          limits:
            cpu: 200m
          requests:
            memory: 256Mi
      - command: ["/monitor", "--source=configconnector:http://localhost:8888?whitelisted=reconcile_requests_total,reconcile_request_duration_seconds,reconcile_workers_total,reconcile_occupied_workers_total,internal_errors_total&customResourceType=k8s_container&customLabels[container_name]&customLabels[project_id]&customLabels[location]&customLabels[cluster_name]&customLabels[namespace_name]&customLabels[pod_name]", "--stackdriver-prefix=kubernetes.io/internal/addons"]
        image: gke.gcr.io/prometheus-to-sd:v0.11.12-gke.11
        name: prom-to-sd
`, `
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    cnrm.cloud.google.com/component: cnrm-webhook-manager
    cnrm.cloud.google.com/system: "true"
  name: cnrm-webhook-manager
  namespace: cnrm-system
spec:
  revisionHistoryLimit: 1
  selector:
    matchLabels:
      cnrm.cloud.google.com/component: cnrm-webhook-manager
      cnrm.cloud.google.com/system: "true"
  template:
    metadata:
      labels:
        cnrm.cloud.google.com/component: cnrm-webhook-manager
        cnrm.cloud.google.com/system: "true"
    spec:
      containers:
      - command:
        - /configconnector/webhook
        env:
        - name: GOMEMLIMIT
          value: 110MiB
        - name: NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        image: gcr.io/gke-release/cnrm/webhook:54aab28
        imagePullPolicy: Always
        name: webhook
        ports:
        - containerPort: 23232
        readinessProbe:
          httpGet:
            path: /ready
            port: 23232
          initialDelaySeconds: 7
          periodSeconds: 3
        resources:
          limits:
            memory: 128Mi
          requests:
            cpu: 250m
            memory: 128Mi
        securityContext:
          allowPrivilegeEscalation: false
          privileged: false
          runAsNonRoot: true
          runAsUser: 1000
      enableServiceLinks: false
      serviceAccountName: cnrm-webhook-manager
      terminationGracePeriodSeconds: 10
`, `
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
  annotations:
    autoscaling.alpha.kubernetes.io/metrics: '[{"type":"Resource","resource":{"name":"memory","targetAverageUtilization":70}}]'
  labels:
    cnrm.cloud.google.com/system: "true"
  name: cnrm-webhook
  namespace: cnrm-system
spec:
  maxReplicas: 20
  minReplicas: 2
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: cnrm-webhook-manager
  targetCPUUtilizationPercentage: 90
`}

// ClusterModeComponentsWithCustomizedControllerManager is the same as ClusterModeComponents
// with the following differences:
// - the "resources" section for cnrm-controller-manager/manager container.
//
// Note that the GOMEMLIMIT env for the webhook manager deployment still has the default "110MiB" value,
// because there was no memory customization on the webhook manager.
var ClusterModeComponentsWithCustomizedControllerManager = []string{`
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
    spec:
      containers:
      - args: ["--scoped-namespace=${NAMESPACE?}", "--stderrthreshold=INFO", "--prometheus-scrape-endpoint=:8888"]
        command: ["/configconnector/manager"]
        image: gcr.io/gke-release/cnrm/controller:4af93f1
        name: manager
        resources:
          limits:
            cpu: 400m
          requests:
            memory: 512Mi
      - command: ["/monitor", "--source=configconnector:http://localhost:8888?whitelisted=reconcile_requests_total,reconcile_request_duration_seconds,reconcile_workers_total,reconcile_occupied_workers_total,internal_errors_total&customResourceType=k8s_container&customLabels[container_name]&customLabels[project_id]&customLabels[location]&customLabels[cluster_name]&customLabels[namespace_name]&customLabels[pod_name]", "--stackdriver-prefix=kubernetes.io/internal/addons"]
        image: gke.gcr.io/prometheus-to-sd:v0.11.12-gke.11
        name: prom-to-sd
`, `
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    cnrm.cloud.google.com/component: cnrm-webhook-manager
    cnrm.cloud.google.com/system: "true"
  name: cnrm-webhook-manager
  namespace: cnrm-system
spec:
  revisionHistoryLimit: 1
  selector:
    matchLabels:
      cnrm.cloud.google.com/component: cnrm-webhook-manager
      cnrm.cloud.google.com/system: "true"
  template:
    metadata:
      labels:
        cnrm.cloud.google.com/component: cnrm-webhook-manager
        cnrm.cloud.google.com/system: "true"
    spec:
      containers:
      - command:
        - /configconnector/webhook
        env:
        - name: GOMEMLIMIT
          value: 110MiB
        - name: NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        image: gcr.io/gke-release/cnrm/webhook:54aab28
        imagePullPolicy: Always
        name: webhook
        ports:
        - containerPort: 23232
        readinessProbe:
          httpGet:
            path: /ready
            port: 23232
          initialDelaySeconds: 7
          periodSeconds: 3
        resources:
          limits:
            memory: 128Mi
          requests:
            cpu: 250m
            memory: 128Mi
        securityContext:
          allowPrivilegeEscalation: false
          privileged: false
          runAsNonRoot: true
          runAsUser: 1000
      enableServiceLinks: false
      serviceAccountName: cnrm-webhook-manager
      terminationGracePeriodSeconds: 10
`, `
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
  annotations:
    autoscaling.alpha.kubernetes.io/metrics: '[{"type":"Resource","resource":{"name":"memory","targetAverageUtilization":70}}]'
  labels:
    cnrm.cloud.google.com/system: "true"
  name: cnrm-webhook
  namespace: cnrm-system
spec:
  maxReplicas: 20
  minReplicas: 2
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: cnrm-webhook-manager
  targetCPUUtilizationPercentage: 90
`}

// ClusterModeComponentsWithCustomizedWebhookManager is the same as ClusterModeComponents
// with the following differences:
// - the "resources" section for cnrm-webhook-manager/webhook container.
// - the "replicas" field for cnrm-webhook-manger deployment.
// - the "minReplicas" field for HorizontalPodAutoscaler.
// - the "GOMEMLIMIT" environment variable.
var ClusterModeComponentsWithCustomizedWebhookManager = []string{`
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
    spec:
      containers:
      - args: ["--scoped-namespace=${NAMESPACE?}", "--stderrthreshold=INFO", "--prometheus-scrape-endpoint=:8888"]
        command: ["/configconnector/manager"]
        image: gcr.io/gke-release/cnrm/controller:4af93f1
        name: manager
        resources:
          limits:
            cpu: 200m
          requests:
            memory: 256Mi
      - command: ["/monitor", "--source=configconnector:http://localhost:8888?whitelisted=reconcile_requests_total,reconcile_request_duration_seconds,reconcile_workers_total,reconcile_occupied_workers_total,internal_errors_total&customResourceType=k8s_container&customLabels[container_name]&customLabels[project_id]&customLabels[location]&customLabels[cluster_name]&customLabels[namespace_name]&customLabels[pod_name]", "--stackdriver-prefix=kubernetes.io/internal/addons"]
        image: gke.gcr.io/prometheus-to-sd:v0.11.12-gke.11
        name: prom-to-sd
`, `
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    cnrm.cloud.google.com/component: cnrm-webhook-manager
    cnrm.cloud.google.com/system: "true"
  name: cnrm-webhook-manager
  namespace: cnrm-system
spec:
  replicas: 4
  revisionHistoryLimit: 1
  selector:
    matchLabels:
      cnrm.cloud.google.com/component: cnrm-webhook-manager
      cnrm.cloud.google.com/system: "true"
  template:
    metadata:
      labels:
        cnrm.cloud.google.com/component: cnrm-webhook-manager
        cnrm.cloud.google.com/system: "true"
    spec:
      containers:
      - command:
        - /configconnector/webhook
        env:
        - name: GOMEMLIMIT
          value: "228170137B"
        - name: NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        image: gcr.io/gke-release/cnrm/webhook:54aab28
        imagePullPolicy: Always
        name: webhook
        ports:
        - containerPort: 23232
        readinessProbe:
          httpGet:
            path: /ready
            port: 23232
          initialDelaySeconds: 7
          periodSeconds: 3
        resources:
          limits:
            memory: 512Mi
          requests:
            cpu: 250m
            memory: 256Mi
        securityContext:
          allowPrivilegeEscalation: false
          privileged: false
          runAsNonRoot: true
          runAsUser: 1000
      enableServiceLinks: false
      serviceAccountName: cnrm-webhook-manager
      terminationGracePeriodSeconds: 10
`, `
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
  annotations:
    autoscaling.alpha.kubernetes.io/metrics: '[{"type":"Resource","resource":{"name":"memory","targetAverageUtilization":70}}]'
  labels:
    cnrm.cloud.google.com/system: "true"
  name: cnrm-webhook
  namespace: cnrm-system
spec:
  maxReplicas: 20
  minReplicas: 4
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: cnrm-webhook-manager
  targetCPUUtilizationPercentage: 90
`}

// ClusterModeComponentsWithCustomizedWebhookManagerWithLargeReplicas is the same as ClusterModeComponents
// with the following differences:
// - the "resources" section for cnrm-webhook-manager/webhook container.
// - the "replicas" field for cnrm-webhook-manger deployment.
// - the "minReplicas" field for HorizontalPodAutoscaler.
// - the "maxReplicas" field for HorizontalPodAutoscaler is also updated to match the value of "minReplcias".
// - the "GOMEMLIMIT" environment variable.
var ClusterModeComponentsWithCustomizedWebhookManagerWithLargeReplicas = []string{`
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
    spec:
      containers:
      - args: ["--scoped-namespace=${NAMESPACE?}", "--stderrthreshold=INFO", "--prometheus-scrape-endpoint=:8888"]
        command: ["/configconnector/manager"]
        image: gcr.io/gke-release/cnrm/controller:4af93f1
        name: manager
        resources:
          limits:
            cpu: 200m
          requests:
            memory: 256Mi
      - command: ["/monitor", "--source=configconnector:http://localhost:8888?whitelisted=reconcile_requests_total,reconcile_request_duration_seconds,reconcile_workers_total,reconcile_occupied_workers_total,internal_errors_total&customResourceType=k8s_container&customLabels[container_name]&customLabels[project_id]&customLabels[location]&customLabels[cluster_name]&customLabels[namespace_name]&customLabels[pod_name]", "--stackdriver-prefix=kubernetes.io/internal/addons"]
        image: gke.gcr.io/prometheus-to-sd:v0.11.12-gke.11
        name: prom-to-sd
`, `
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    cnrm.cloud.google.com/component: cnrm-webhook-manager
    cnrm.cloud.google.com/system: "true"
  name: cnrm-webhook-manager
  namespace: cnrm-system
spec:
  replicas: 30
  revisionHistoryLimit: 1
  selector:
    matchLabels:
      cnrm.cloud.google.com/component: cnrm-webhook-manager
      cnrm.cloud.google.com/system: "true"
  template:
    metadata:
      labels:
        cnrm.cloud.google.com/component: cnrm-webhook-manager
        cnrm.cloud.google.com/system: "true"
    spec:
      containers:
      - command:
        - /configconnector/webhook
        env:
        - name: GOMEMLIMIT
          value: "228170137B"
        - name: NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        image: gcr.io/gke-release/cnrm/webhook:54aab28
        imagePullPolicy: Always
        name: webhook
        ports:
        - containerPort: 23232
        readinessProbe:
          httpGet:
            path: /ready
            port: 23232
          initialDelaySeconds: 7
          periodSeconds: 3
        resources:
          limits:
            memory: 512Mi
          requests:
            cpu: 250m
            memory: 256Mi
        securityContext:
          allowPrivilegeEscalation: false
          privileged: false
          runAsNonRoot: true
          runAsUser: 1000
      enableServiceLinks: false
      serviceAccountName: cnrm-webhook-manager
      terminationGracePeriodSeconds: 10
`, `
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
  annotations:
    autoscaling.alpha.kubernetes.io/metrics: '[{"type":"Resource","resource":{"name":"memory","targetAverageUtilization":70}}]'
  labels:
    cnrm.cloud.google.com/system: "true"
  name: cnrm-webhook
  namespace: cnrm-system
spec:
  maxReplicas: 30
  minReplicas: 30
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: cnrm-webhook-manager
  targetCPUUtilizationPercentage: 90
`}

var NamespacedComponents = []string{`
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

// NamespacedComponentsWithCustomizedControllerManager is the same as NamespacedComponents
// with the following differences:
// - the "resources" section for cnrm-controller-manager/manager container.
var NamespacedComponentsWithCustomizedControllerManager = []string{`
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
        resources:
          limits:
            cpu: 400m
          requests:
            memory: 512Mi
      - command: ["/monitor", "--source=configconnector:http://localhost:8888?whitelisted=reconcile_requests_total,reconcile_request_duration_seconds,reconcile_workers_total,reconcile_occupied_workers_total,internal_errors_total&customResourceType=k8s_container&customLabels[container_name]&customLabels[project_id]&customLabels[location]&customLabels[cluster_name]&customLabels[namespace_name]&customLabels[pod_name]", "--stackdriver-prefix=kubernetes.io/internal/addons"]
        image: gke.gcr.io/prometheus-to-sd:v0.11.12-gke.11
        name: prom-to-sd
`}

// NamespacedComponentsWithRatLimitCustomization is the same as NamespacedComponents
// with the following differences:
// - the "args" for cnrm-controller-manager/manager container.
var NamespacedComponentsWithRatLimitCustomization = []string{`
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
      - args: ["--qps=80", "--burst=30", "--scoped-namespace=${NAMESPACE?}", "--stderrthreshold=INFO", "--prometheus-scrape-endpoint=:8888"]
        command: ["/configconnector/manager"]
        image: gcr.io/gke-release/cnrm/controller:4af93f1
        name: manager
      - command: ["/monitor", "--source=configconnector:http://localhost:8888?whitelisted=reconcile_requests_total,reconcile_request_duration_seconds,reconcile_workers_total,reconcile_occupied_workers_total,internal_errors_total&customResourceType=k8s_container&customLabels[container_name]&customLabels[project_id]&customLabels[location]&customLabels[cluster_name]&customLabels[namespace_name]&customLabels[pod_name]", "--stackdriver-prefix=kubernetes.io/internal/addons"]
        image: gke.gcr.io/prometheus-to-sd:v0.11.12-gke.11
        name: prom-to-sd
`}

// ClusterModeComponentsWithRatLimitCustomization is the same as ClusterModeComponents
// with the following differences:
// - the "args" for cnrm-controller-manager/manager container.
var ClusterModeComponentsWithRatLimitCustomization = []string{`
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
    spec:
      containers:
      - args: ["--qps=80", "--burst=30", "--scoped-namespace=${NAMESPACE?}", "--stderrthreshold=INFO", "--prometheus-scrape-endpoint=:8888"]
        command: ["/configconnector/manager"]
        image: gcr.io/gke-release/cnrm/controller:4af93f1
        name: manager
        resources:
          limits:
            cpu: 200m
          requests:
            memory: 256Mi
      - command: ["/monitor", "--source=configconnector:http://localhost:8888?whitelisted=reconcile_requests_total,reconcile_request_duration_seconds,reconcile_workers_total,reconcile_occupied_workers_total,internal_errors_total&customResourceType=k8s_container&customLabels[container_name]&customLabels[project_id]&customLabels[location]&customLabels[cluster_name]&customLabels[namespace_name]&customLabels[pod_name]", "--stackdriver-prefix=kubernetes.io/internal/addons"]
        image: gke.gcr.io/prometheus-to-sd:v0.11.12-gke.11
        name: prom-to-sd
`, `
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    cnrm.cloud.google.com/component: cnrm-webhook-manager
    cnrm.cloud.google.com/system: "true"
  name: cnrm-webhook-manager
  namespace: cnrm-system
spec:
  revisionHistoryLimit: 1
  selector:
    matchLabels:
      cnrm.cloud.google.com/component: cnrm-webhook-manager
      cnrm.cloud.google.com/system: "true"
  template:
    metadata:
      labels:
        cnrm.cloud.google.com/component: cnrm-webhook-manager
        cnrm.cloud.google.com/system: "true"
    spec:
      containers:
      - command:
        - /configconnector/webhook
        env:
        - name: GOMEMLIMIT
          value: 110MiB
        - name: NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        image: gcr.io/gke-release/cnrm/webhook:54aab28
        imagePullPolicy: Always
        name: webhook
        ports:
        - containerPort: 23232
        readinessProbe:
          httpGet:
            path: /ready
            port: 23232
          initialDelaySeconds: 7
          periodSeconds: 3
        resources:
          limits:
            memory: 128Mi
          requests:
            cpu: 250m
            memory: 128Mi
        securityContext:
          allowPrivilegeEscalation: false
          privileged: false
          runAsNonRoot: true
          runAsUser: 1000
      enableServiceLinks: false
      serviceAccountName: cnrm-webhook-manager
      terminationGracePeriodSeconds: 10
`, `
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
  annotations:
    autoscaling.alpha.kubernetes.io/metrics: '[{"type":"Resource","resource":{"name":"memory","targetAverageUtilization":70}}]'
  labels:
    cnrm.cloud.google.com/system: "true"
  name: cnrm-webhook
  namespace: cnrm-system
spec:
  maxReplicas: 20
  minReplicas: 2
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: cnrm-webhook-manager
  targetCPUUtilizationPercentage: 90
`}
