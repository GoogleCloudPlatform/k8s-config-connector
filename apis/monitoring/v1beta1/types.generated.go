// Copyright 2024 Google LLC
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

package v1beta1

// +kcc:proto=google.monitoring.v3.Service.AppEngine
type Service_AppEngine struct {
	// The ID of the App Engine module underlying this service. Corresponds to
	//  the `module_id` resource label in the [`gae_app` monitored
	//  resource](https://cloud.google.com/monitoring/api/resources#tag_gae_app).
	ModuleID *string `json:"moduleID,omitempty"`
}

// +kcc:proto=google.monitoring.v3.Service.BasicService
type Service_BasicService struct {
	// The type of service that this basic service defines, e.g.
	//  APP_ENGINE service type.
	//  Documentation and valid values
	//  [here](https://cloud.google.com/stackdriver/docs/solutions/slo-monitoring/api/api-structures#basic-svc-w-basic-sli).
	ServiceType *string `json:"serviceType,omitempty"`

	// Labels that specify the resource that emits the monitoring data which
	//  is used for SLO reporting of this `Service`.
	//  Documentation and valid values for given service types
	//  [here](https://cloud.google.com/stackdriver/docs/solutions/slo-monitoring/api/api-structures#basic-svc-w-basic-sli).
	ServiceLabels map[string]string `json:"serviceLabels,omitempty"`
}

// +kcc:proto=google.monitoring.v3.Service.CloudEndpoints
type Service_CloudEndpoints struct {
	// The name of the Cloud Endpoints service underlying this service.
	//  Corresponds to the `service` resource label in the [`api` monitored
	//  resource](https://cloud.google.com/monitoring/api/resources#tag_api).
	Service *string `json:"service,omitempty"`
}

// +kcc:proto=google.monitoring.v3.Service.CloudRun
type Service_CloudRun struct {
	// The name of the Cloud Run service. Corresponds to the `service_name`
	//  resource label in the [`cloud_run_revision` monitored
	//  resource](https://cloud.google.com/monitoring/api/resources#tag_cloud_run_revision).
	ServiceName *string `json:"serviceName,omitempty"`

	// The location the service is run. Corresponds to the `location`
	//  resource label in the [`cloud_run_revision` monitored
	//  resource](https://cloud.google.com/monitoring/api/resources#tag_cloud_run_revision).
	Location *string `json:"location,omitempty"`
}

// +kcc:proto=google.monitoring.v3.Service.ClusterIstio
type Service_ClusterIstio struct {
	// The location of the Kubernetes cluster in which this Istio service is
	//  defined. Corresponds to the `location` resource label in `k8s_cluster`
	//  resources.
	Location *string `json:"location,omitempty"`

	// The name of the Kubernetes cluster in which this Istio service is
	//  defined. Corresponds to the `cluster_name` resource label in
	//  `k8s_cluster` resources.
	ClusterName *string `json:"clusterName,omitempty"`

	// The namespace of the Istio service underlying this service. Corresponds
	//  to the `destination_service_namespace` metric label in Istio metrics.
	ServiceNamespace *string `json:"serviceNamespace,omitempty"`

	// The name of the Istio service underlying this service. Corresponds to the
	//  `destination_service_name` metric label in Istio metrics.
	ServiceName *string `json:"serviceName,omitempty"`
}

// +kcc:proto=google.monitoring.v3.Service.Custom
type Service_Custom struct {
}

// +kcc:proto=google.monitoring.v3.Service.GkeNamespace
type Service_GkeNamespace struct {
	// Output only. The project this resource lives in. For legacy services
	//  migrated from the `Custom` type, this may be a distinct project from the
	//  one parenting the service itself.
	ProjectID *string `json:"projectID,omitempty"`

	// The location of the parent cluster. This may be a zone or region.
	Location *string `json:"location,omitempty"`

	// The name of the parent cluster.
	ClusterName *string `json:"clusterName,omitempty"`

	// The name of this namespace.
	NamespaceName *string `json:"namespaceName,omitempty"`
}

// +kcc:proto=google.monitoring.v3.Service.GkeService
type Service_GkeService struct {
	// Output only. The project this resource lives in. For legacy services
	//  migrated from the `Custom` type, this may be a distinct project from the
	//  one parenting the service itself.
	ProjectID *string `json:"projectID,omitempty"`

	// The location of the parent cluster. This may be a zone or region.
	Location *string `json:"location,omitempty"`

	// The name of the parent cluster.
	ClusterName *string `json:"clusterName,omitempty"`

	// The name of the parent namespace.
	NamespaceName *string `json:"namespaceName,omitempty"`

	// The name of this service.
	ServiceName *string `json:"serviceName,omitempty"`
}

// +kcc:proto=google.monitoring.v3.Service.GkeWorkload
type Service_GkeWorkload struct {
	// Output only. The project this resource lives in. For legacy services
	//  migrated from the `Custom` type, this may be a distinct project from the
	//  one parenting the service itself.
	ProjectID *string `json:"projectID,omitempty"`

	// The location of the parent cluster. This may be a zone or region.
	Location *string `json:"location,omitempty"`

	// The name of the parent cluster.
	ClusterName *string `json:"clusterName,omitempty"`

	// The name of the parent namespace.
	NamespaceName *string `json:"namespaceName,omitempty"`

	// The type of this workload (for example, "Deployment" or "DaemonSet")
	TopLevelControllerType *string `json:"topLevelControllerType,omitempty"`

	// The name of this workload.
	TopLevelControllerName *string `json:"topLevelControllerName,omitempty"`
}

// +kcc:proto=google.monitoring.v3.Service.IstioCanonicalService
type Service_IstioCanonicalService struct {
	// Identifier for the Istio mesh in which this canonical service is defined.
	//  Corresponds to the `mesh_uid` metric label in
	//  [Istio metrics](https://cloud.google.com/monitoring/api/metrics_istio).
	MeshUid *string `json:"meshUid,omitempty"`

	// The namespace of the canonical service underlying this service.
	//  Corresponds to the `destination_canonical_service_namespace` metric
	//  label in [Istio
	//  metrics](https://cloud.google.com/monitoring/api/metrics_istio).
	CanonicalServiceNamespace *string `json:"canonicalServiceNamespace,omitempty"`

	// The name of the canonical service underlying this service.
	//  Corresponds to the `destination_canonical_service_name` metric label in
	//  label in [Istio
	//  metrics](https://cloud.google.com/monitoring/api/metrics_istio).
	CanonicalService *string `json:"canonicalService,omitempty"`
}

// +kcc:proto=google.monitoring.v3.Service.MeshIstio
type Service_MeshIstio struct {
	// Identifier for the mesh in which this Istio service is defined.
	//  Corresponds to the `mesh_uid` metric label in Istio metrics.
	MeshUid *string `json:"meshUid,omitempty"`

	// The namespace of the Istio service underlying this service. Corresponds
	//  to the `destination_service_namespace` metric label in Istio metrics.
	ServiceNamespace *string `json:"serviceNamespace,omitempty"`

	// The name of the Istio service underlying this service. Corresponds to the
	//  `destination_service_name` metric label in Istio metrics.
	ServiceName *string `json:"serviceName,omitempty"`
}

// +kcc:proto=google.monitoring.v3.Service.Telemetry
type Service_Telemetry struct {
	// The full name of the resource that defines this service. Formatted as
	//  described in https://cloud.google.com/apis/design/resource_names.
	ResourceName *string `json:"resourceName,omitempty"`
}
