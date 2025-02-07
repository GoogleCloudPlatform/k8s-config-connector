// Copyright 2025 Google LLC
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

package v1alpha1


// +kcc:proto=google.appengine.v1.Instance
type Instance struct {
}

// +kcc:proto=google.appengine.v1.Instance
type InstanceObservedState struct {
	// Output only. Full path to the Instance resource in the API.
	//  Example: `apps/myapp/services/default/versions/v1/instances/instance-1`.
	// +kcc:proto:field=google.appengine.v1.Instance.name
	Name *string `json:"name,omitempty"`

	// Output only. Relative name of the instance within the version.
	//  Example: `instance-1`.
	// +kcc:proto:field=google.appengine.v1.Instance.id
	ID *string `json:"id,omitempty"`

	// Output only. App Engine release this instance is running on.
	// +kcc:proto:field=google.appengine.v1.Instance.app_engine_release
	AppEngineRelease *string `json:"appEngineRelease,omitempty"`

	// Output only. Availability of the instance.
	// +kcc:proto:field=google.appengine.v1.Instance.availability
	Availability *string `json:"availability,omitempty"`

	// Output only. Name of the virtual machine where this instance lives. Only applicable
	//  for instances in App Engine flexible environment.
	// +kcc:proto:field=google.appengine.v1.Instance.vm_name
	VmName *string `json:"vmName,omitempty"`

	// Output only. Zone where the virtual machine is located. Only applicable for instances
	//  in App Engine flexible environment.
	// +kcc:proto:field=google.appengine.v1.Instance.vm_zone_name
	VmZoneName *string `json:"vmZoneName,omitempty"`

	// Output only. Virtual machine ID of this instance. Only applicable for instances in
	//  App Engine flexible environment.
	// +kcc:proto:field=google.appengine.v1.Instance.vm_id
	VmID *string `json:"vmID,omitempty"`

	// Output only. Time that this instance was started.
	//
	//  @OutputOnly
	// +kcc:proto:field=google.appengine.v1.Instance.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Number of requests since this instance was started.
	// +kcc:proto:field=google.appengine.v1.Instance.requests
	Requests *int32 `json:"requests,omitempty"`

	// Output only. Number of errors since this instance was started.
	// +kcc:proto:field=google.appengine.v1.Instance.errors
	Errors *int32 `json:"errors,omitempty"`

	// Output only. Average queries per second (QPS) over the last minute.
	// +kcc:proto:field=google.appengine.v1.Instance.qps
	Qps *float32 `json:"qps,omitempty"`

	// Output only. Average latency (ms) over the last minute.
	// +kcc:proto:field=google.appengine.v1.Instance.average_latency
	AverageLatency *int32 `json:"averageLatency,omitempty"`

	// Output only. Total memory in use (bytes).
	// +kcc:proto:field=google.appengine.v1.Instance.memory_usage
	MemoryUsage *int64 `json:"memoryUsage,omitempty"`

	// Output only. Status of the virtual machine where this instance lives. Only applicable
	//  for instances in App Engine flexible environment.
	// +kcc:proto:field=google.appengine.v1.Instance.vm_status
	VmStatus *string `json:"vmStatus,omitempty"`

	// Output only. Whether this instance is in debug mode. Only applicable for instances in
	//  App Engine flexible environment.
	// +kcc:proto:field=google.appengine.v1.Instance.vm_debug_enabled
	VmDebugEnabled *bool `json:"vmDebugEnabled,omitempty"`

	// Output only. The IP address of this instance. Only applicable for instances in App
	//  Engine flexible environment.
	// +kcc:proto:field=google.appengine.v1.Instance.vm_ip
	VmIP *string `json:"vmIP,omitempty"`

	// Output only. The liveness health check of this instance. Only applicable for instances
	//  in App Engine flexible environment.
	// +kcc:proto:field=google.appengine.v1.Instance.vm_liveness
	VmLiveness *string `json:"vmLiveness,omitempty"`
}
