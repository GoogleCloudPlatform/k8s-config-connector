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

	customizev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/customize/v1alpha1"

	"google.golang.org/protobuf/proto"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	ControllerResourceCRForControllerManagerResources = &customizev1alpha1.ControllerResource{
		ObjectMeta: metav1.ObjectMeta{
			Name: "cnrm-controller-manager",
		},
		Spec: customizev1alpha1.ControllerResourceSpec{
			Containers: []customizev1alpha1.ContainerResourceSpec{
				{
					Name: "manager",
					Resources: corev1.ResourceRequirements{
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
	ControllerResourceCRForControllerManagerReplicas = &customizev1alpha1.ControllerResource{
		ObjectMeta: metav1.ObjectMeta{
			Name: "cnrm-controller-manager",
		},
		Spec: customizev1alpha1.ControllerResourceSpec{
			Replicas: proto.Int64(int64(4)),
			Containers: []customizev1alpha1.ContainerResourceSpec{
				{
					Name: "manager",
					Resources: corev1.ResourceRequirements{
						Limits:   corev1.ResourceList{},
						Requests: corev1.ResourceList{},
					},
				},
			},
		},
	}
	ControllerResourceCRForWebhookManagerResourcesAndReplicas = &customizev1alpha1.ControllerResource{
		ObjectMeta: metav1.ObjectMeta{
			Name: "cnrm-webhook-manager",
		},
		Spec: customizev1alpha1.ControllerResourceSpec{
			Replicas: proto.Int64(int64(4)),
			Containers: []customizev1alpha1.ContainerResourceSpec{
				{
					Name: "webhook",
					Resources: corev1.ResourceRequirements{
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
	ControllerResourceCRForWebhookManagerWithLargeReplicas = &customizev1alpha1.ControllerResource{
		ObjectMeta: metav1.ObjectMeta{
			Name: "cnrm-webhook-manager",
		},
		Spec: customizev1alpha1.ControllerResourceSpec{
			Replicas: proto.Int64(int64(30)), // this value is larger than the default value of "maxReplicas" of HPA in KCC's manifests
			Containers: []customizev1alpha1.ContainerResourceSpec{
				{
					Name: "webhook",
					Resources: corev1.ResourceRequirements{
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
)

var (
	nonExistingControllerName                    = "controller-does-not-exist"
	ControllerResourceCRForNonExistingController = &customizev1alpha1.ControllerResource{
		ObjectMeta: metav1.ObjectMeta{
			Name: nonExistingControllerName,
		},
		Spec: customizev1alpha1.ControllerResourceSpec{
			Containers: []customizev1alpha1.ContainerResourceSpec{},
		},
	}
	ErrNonExistingController = fmt.Sprintf("resource customization for controller %s is not supported", nonExistingControllerName)
)

var (
	nonExistingContainerName                    = "recorder" // there is no "recorder" container in "cnrm-controller-manager".
	ControllerResourceCRForNonExistingContainer = &customizev1alpha1.ControllerResource{
		ObjectMeta: metav1.ObjectMeta{
			Name: "cnrm-controller-manager",
		},
		Spec: customizev1alpha1.ControllerResourceSpec{
			Containers: []customizev1alpha1.ContainerResourceSpec{
				{
					Name: nonExistingContainerName,
				},
			},
		},
	}
	ErrNonExistingContainer = fmt.Sprintf("failed to apply customization cnrm-controller-manager: resource customization failed for the following containers because there are no matching containers in the manifest: %s", nonExistingContainerName)
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
        image: gke.gcr.io/prometheus-to-sd:v0.9.1
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
        image: gke.gcr.io/prometheus-to-sd:v0.9.1
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
        image: gke.gcr.io/prometheus-to-sd:v0.9.1
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
        image: gke.gcr.io/prometheus-to-sd:v0.9.1
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
