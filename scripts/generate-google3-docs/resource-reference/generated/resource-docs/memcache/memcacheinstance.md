{% extends "config-connector/_base.html" %}

{% block page_title %}MemcacheInstance{% endblock %}
{% block body %}


<table>
<thead>
<tr>
<th><strong>Property</strong></th>
<th><strong>Value</strong></th>
</tr>
</thead>
<tbody>
<tr>
<td>{{gcp_name_short}} Service Name</td>
<td>Cloud Memorystore for Memcached</td>
</tr>
<tr>
<td>{{gcp_name_short}} Service Documentation</td>
<td><a href="/memorystore/docs/memcached/">/memorystore/docs/memcached/</a></td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Name</td>
<td>v1beta2.projects.locations.instances</td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Documentation</td>
<td><a href="/memorystore/docs/memcached/reference/rest/v1beta2/projects.locations.instances">/memorystore/docs/memcached/reference/rest/v1beta2/projects.locations.instances</a></td>
</tr>
<tr>
<td>{{product_name_short}} Resource Short Names</td>
<td>gcpmemcacheinstance<br>gcpmemcacheinstances<br>memcacheinstance</td>
</tr>
<tr>
<td>{{product_name_short}} Service Name</td>
<td>memcache.googleapis.com</td>
</tr>
<tr>
<td>{{product_name_short}} Resource Fully Qualified Name</td>
<td>memcacheinstances.memcache.cnrm.cloud.google.com</td>
</tr>

<tr>
    <td>Can Be Referenced by IAMPolicy/IAMPolicyMember</td>
    <td>No</td>
</tr>


</tbody>
</table>

## Custom Resource Definition Properties


### Annotations
<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>cnrm.cloud.google.com/project-id</code></td>
    </tr>
</tbody>
</table>


### Spec
#### Schema
  ```yaml
  displayName: string
  memcacheParameters:
    id: string
    params:
      string: string
  memcacheVersion: string
  networkRef:
    external: string
    name: string
    namespace: string
  nodeConfig:
    cpuCount: integer
    memorySizeMb: integer
  nodeCount: integer
  region: string
  resourceID: string
  zones:
  - string
  ```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td>
            <p><code>displayName</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}A user-visible name for the instance.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>memcacheParameters</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Immutable. User-specified parameters for this memcache instance.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>memcacheParameters.id</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}This is a unique ID associated with this set of parameters.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>memcacheParameters.params</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">map (key: string, value: string)</code></p>
            <p>{% verbatim %}User-defined set of parameters to use in the memcache process.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>memcacheVersion</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The major version of Memcached software. If not provided, latest supported version will be used.
Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
determined by our system based on the latest supported minor version. Default value: "MEMCACHE_1_5" Possible values: ["MEMCACHE_1_5"].{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>networkRef</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}The full name of the network to connect the instance to.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>networkRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: string of the format `projects/{{project}}/global/networks/{{value}}`, where {{value}} is the `name` field of a `ComputeNetwork` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>networkRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>networkRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>nodeConfig</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Immutable. Configuration for memcache nodes.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>nodeConfig.cpuCount</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Number of CPUs per node.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>nodeConfig.memorySizeMb</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Memory size in Mebibytes for each memcache node.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>nodeCount</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Number of nodes in the memcache instance.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>region</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. The region of the Memcache instance. If it is not provided, the provider region is used.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>resourceID</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>zones</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Immutable. Zones where memcache nodes should be provisioned.  If not
provided, all zones will be used.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>zones[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>



### Status
#### Schema
  ```yaml
  conditions:
  - lastTransitionTime: string
    message: string
    reason: string
    status: string
    type: string
  createTime: string
  discoveryEndpoint: string
  memcacheFullVersion: string
  memcacheNodes:
  - host: string
    nodeId: string
    port: integer
    state: string
    zone: string
  observedGeneration: integer
  ```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>conditions</code></td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Conditions represent the latest available observation of the resource's current state.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[]</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].lastTransitionTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Last time the condition transitioned from one status to another.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].message</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Human-readable message indicating details about last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].reason</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Unique, one-word, CamelCase reason for the condition's last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].status</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Status is the status of the condition. Can be True, False, Unknown.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].type</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Type is the type of the condition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>createTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Creation timestamp in RFC3339 text format.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>discoveryEndpoint</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Endpoint for Discovery API.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>memcacheFullVersion</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The full version of memcached server running on this instance.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>memcacheNodes</code></td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Additional information about the instance state, if available.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>memcacheNodes[]</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>memcacheNodes[].host</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Hostname or IP address of the Memcached node used by the clients to connect to the Memcached server on this node.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>memcacheNodes[].nodeId</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Identifier of the Memcached node. The node id does not include project or location like the Memcached instance name.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>memcacheNodes[].port</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}The port number of the Memcached server on this node.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>memcacheNodes[].state</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Current state of the Memcached node.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>memcacheNodes[].zone</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Location (GCP Zone) for the Memcached node.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedGeneration</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>

## Sample YAML(s)

### Typical Use Case
  ```yaml
  # Copyright 2020 Google LLC
  #
  # Licensed under the Apache License, Version 2.0 (the "License");
  # you may not use this file except in compliance with the License.
  # You may obtain a copy of the License at
  #
  #     http://www.apache.org/licenses/LICENSE-2.0
  #
  # Unless required by applicable law or agreed to in writing, software
  # distributed under the License is distributed on an "AS IS" BASIS,
  # WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  # See the License for the specific language governing permissions and
  # limitations under the License.
  
  apiVersion: memcache.cnrm.cloud.google.com/v1beta1
  kind: MemcacheInstance
  metadata:
    labels:
      size: small
      process-type: long-queue
    name: memcacheinstance-sample
  spec:
    networkRef:
      name: memcacheinstance-dep
    nodeConfig:
      memorySizeMb: 1024
      cpuCount: 1
    nodeCount: 2
    region: us-central1
    zones:
      - us-central1-a
      - us-central1-c
    displayName: Sample Memcache Instance
    memcacheParameters:
      params:
        listen-backlog: "10000"
        max-item-size: "524288"
        max-reqs-per-event: "1"
  ---
  apiVersion: compute.cnrm.cloud.google.com/v1beta1
  kind: ComputeAddress
  metadata:
    name: memcacheinstance-dep
  spec:
    addressType: INTERNAL
    location: global
    purpose: VPC_PEERING
    prefixLength: 16
    networkRef:
      name: memcacheinstance-dep
  ---
  apiVersion: compute.cnrm.cloud.google.com/v1beta1
  kind: ComputeNetwork
  metadata:
    name: memcacheinstance-dep
  spec:
    autoCreateSubnetworks: false
  ---
  apiVersion: servicenetworking.cnrm.cloud.google.com/v1beta1
  kind: ServiceNetworkingConnection
  metadata:
    name: memcacheinstance-dep
  spec:
    networkRef:
      name: memcacheinstance-dep
    reservedPeeringRanges:
      - name: memcacheinstance-dep
    service: "servicenetworking.googleapis.com"
  ```


{% endblock %}
