# Direct Resource API Validation


# TL;DR

Config Connector Direct Resource shall have each `spec` field validated.


# Basic Rules

Each Config Connector resource `spec` field shall be validated by at least one of the following approach:


## 1. CRD validation

* This throws standard OpenAPI validation errors without requiring the Config Connector controller to run.
* This validation shall take the responsibility to give users a cheap and self-guiding check. It does not contain complicated logic. 


## 2. Config Connector controller validation


* This throws the Config Connector-defined errors in runtime.
* This validation shall take the responsibility to guardrail config issues that otherwise could cause ambiguous or unexpected Config Connector behavior.
* We shall build and publish a Config Connector error code table to future explain the issue and give a triage guide.


## 3. GCP Service Validation


* This throws GCP defined errors in runtime. Config Connector shall propagate the GCP response errors to the Config Connector object’s `status` field. 
* This validation shall take the responsibility to guide users on fixing the issues due to GCP service requirements, like `SpannerInstance` has `spec.processingUnit` whose value can only increase but not decrease. 


# CRD validation


## Rule 1: Required/Optional

Required field should use tag `+required`. 

Optional field does not need the tag, make sure it has `omitempty` json tag.


```
type CloudBuildWorkerPoolSpec struct {
  // +required   
  PrivatePoolConfig *PrivatePoolV1Config `json:"privatePoolV1Config,omitempty"`
}
```


## Rule 2: Immutable field


### Option 1

Validates the immutability in the controller runtime, by comparing the difference between `spec` and `status`, together with the corresponding proto field "IMMUTABLE" **field_behavior**. (default behavior if using the direct auto-generated code)

All google services are located in [googleapis Github repo](https://github.com/googleapis/googleapis/tree/master/google),
refer to your resource's API documentation to identify the service name, for example: [compute](https://cloud.google.com/compute/docs/reference/rest/v1).
Once you identify the service, find the proper path to the proto files, for example:
[`google/cloud/compute/v1/compute.proto`](https://github.com/googleapis/googleapis/blob/master/google/cloud/compute/v1/compute.proto).
To see if you can use this approach, check if the resource proto supports **field_behavior** (See [AIP 203](https://google.aip.dev/203)).

Search for the resource name, like `message TargetTcpProxy`, this resource does not support immutable field_behavior.
SecureSourceManagerInstance resource, for example `message Instance` in [secure_source_manager.proto](https://github.com/googleapis/googleapis/blob/master/google/cloud/securesourcemanager/v1/secure_source_manager.proto),
field [`kms_key`](https://github.com/googleapis/googleapis/blob/master/google/cloud/securesourcemanager/v1/secure_source_manager.proto#L398)(KCC type KmsKeyRef) 
has `(google.api.field_behavior) = IMMUTABLE`.

If the resource proto supports field_behavior, you need to make sure
1. Those immutable fields are added in both `spec` and `status.observedState`.
1. Your controller calls the `CompareProtoMessage` function in `AdapterForObject.Update`. 

### Option 2

Use CEL rule `+kubebuilder:validation:XValidation:rule="self == oldSelf",message="the field is immutable"` 
to validate field immutability in the CRD level. This field needs to be *required*.

If the field is *optional*, add another CEL rule`//+kubebuilder:validation:Optional` to make sure field must be allowed to be initially unset,
and immutable once it has been first set.
```
type PrivatePoolV1Config_NetworkConfigSpec struct {
        // +kubebuilder:validation:Optional
        // +kubebuilder:validation:XValidation:rule="self == oldSelf",message="the field is immutable"
        // Immutable. The network definition that the workers are peered
        //  to. If this section is left empty, the workers will be peered to
        //  `WorkerPool.project_id` on the service producer network.
        PeeredNetworkRef refv1beta1.ComputeNetworkRef `json:"peeredNetworkRef,omitempty"`
}
```

Note:

* Kubernetes supports CEL since 1.26 which is the oldest GKE cluster version (written at ). Using CEL shall not cause GKE cluster version issues, but we shall reevaluate this when the condition changes.
* Some Config Connector resources use webhook to validate immutable fields. We can continue using that, but for external contributors CEL is much easier for self-learning. 

* We only apply CEL on immutable fields.** More complicated CEL validation requires further discussion.**

### Option 3

If the field contains more complicated GCP behavior, for example the GCP server assigns a default value if the field is not configured, Config Connector may not check the immutability but rely on the GCP server to give more user facing message (reflected in the `status.conditions`). This is a case-by-case analysis. Please consult the Config Connector team by filing an issue if you encounter a special case.

## Rule 3: Parent

Parent field should be required and immutable.

Note that parent is not a field name. It refers to the pre-requisite of the resource in question. You can identify the parent field through to the cloud resource's parent segment in the URL. For example, for GKE clusters, the [GET request URL](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.zones.clusters/get) `https://container.googleapis.com/v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}` has `projectId` and `zone` in the parent segment. Therefore, for GKE clusters, projectRef and zone will be the parent fields.


### Required

Suggest using the following `inline` struct


```
type CloudBuildWorkerPoolSpec struct {
  // +required   
  Parent `json:",inline"`
}

type Parent struct {
    // +required      
    ProjectRef *refv1beta1.ProjectRef `json:"projectRef"`

    // +required
    Location string `json:"location"`
}
```

### Immutable

The parent immutability should be guaranteed by the `status.externalRef`, where the direct controller checks the parents between `spec` and `status.externalRef` and make sure they match in each reconciliation. 

This is the default behavior if using the direct autogenerated code. The code logic is in the `apis/<service>/<version>/ <resource>_reference.go` when you run the "generate-types" command to generate the APIs.

### Exceptions

Some existing *v1beta1* resources allow optional `spec.projectRef`, where the project information comes from either the `cnrm.cloud.google.com/project-id` annotation, or the namespace annotation.

When migrating a Terraform-based or DCL-based resource to the direct approach, we want to continue supporitng this behavior for backward compatibility. You can check if the resource uses cnrm.cloud.google.com/project-id` annotation in its Google reference doc.

## Rule 4: No `anyOf`, `oneOf`, or` allOf`

Short answer: Not now.

Full answer: 

OpenAPI supports `oneOf`, `anyOf` and `allOf`. Thus, in theory these should be the CRD level rules. The Direct Resource CRD is auto-generated by controller-gen, which only reads the kubebuilder tag rules. But the kubebuilder tags [do not (and most likely won’t)](https://github.com/kubernetes-sigs/controller-tools/issues/461) support `oneOf`, `anyOf` or `allOf`.

To set this as CRD level rule, Config Connector shall build its own CRD transformer.  Ideally, this could be a Config Connector tag similar to kubebuilder, but this requires integrating with controller-gen which is a non-trivial amount of work. Other options like webhook, OpenAPI schema modifier are not self-explaining and easy enough to use, considering that Direct Resource shall be open to external contributors. Thus, Config Connector does not use CRD level validation for `oneOf`, `anyOf` and `allOf`.

One future improvement can be having a Direct Resource tag that adds CRD level checks for   `oneOf`, `anyOf` or `allOf`.


```
// +Config Connector:validation:oneOf:serviceAttachmentRef,targetGRPCProxyRef,targetHTTPProxyRef,targetHTTPSProxyRef"
type target struct {
        ServiceAttachmentRef *refs.ComputeServiceAttachmentRef `json:"serviceAttachmentRef,omitempty"`

        TargetGRPCProxyRef *refs.ComputeTargetGrpcProxyRef `json:"targetGRPCProxyRef,omitempty"`

        TargetHTTPProxyRef *refs.ComputeTargetHTTPProxyRef `json:"targetHTTPProxyRef,omitempty"`

        TargetHTTPSProxyRef *refs.ComputeTargetHTTPSProxyRef `json:"targetHTTPSProxyRef,omitempty"`
}
```



# Controller validation

Advanced Kubernetes contributors can use webhook to validate the resource CRs. 

To easily ramp-up external contributors, we recommend using the controller-level validation instead. This validation mostly happens on Adapter [initialization](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/91aac5186eb0aa2c5c3def94d7a7c79f948ac9a9/pkg/controller/direct/directbase/directbase_controller.go#L226), GCP object [creation](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/91aac5186eb0aa2c5c3def94d7a7c79f948ac9a9/pkg/controller/direct/directbase/directbase_controller.go#L283) and [update](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/91aac5186eb0aa2c5c3def94d7a7c79f948ac9a9/pkg/controller/direct/directbase/directbase_controller.go#L291) steps. 


## Rule 1: Resource Reference

Config Connector shall validate the resource reference based on the [resource reference guide](./resource-reference.md). 


## Rule 2: Resource ID

`spec.resourceID` is an optional and  immutable field. 

Config Connector shall validate the `resourceID` (if present) matches the `status.externalRef` (if present) based on the  reconciliation 4 steps.
