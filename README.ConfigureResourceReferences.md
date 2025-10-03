# Overview

The following guidelines will help you identify potential reference fields and
configure them in service mappings.

## What is a reference field?

Config Connector uses the concept of
[resource references](https://cloud.google.com/config-connector/docs/how-to/creating-resource-references),
which lets you specify the value of a resource reference field by referencing other resources.

A reference field is a field whose value comes from another Google Cloud resource.
The common cases of such values are:

*   Resource ID of another GCP resource
*   Relative resource name of another GCP resource
*   Self link of another GCP resource
*   A service-generated value of another GCP resource

## Identify which fields should be turned into reference fields

You want to look at the
[Terraform resource documentation](https://registry.terraform.io/providers/hashicorp/google-beta/latest/docs)
when researching the resource. You can also look at the API documentation
if needed.

For example, to learn the fields in PubSubTopic, you should check the
fields under the Arguments Reference section in
[google_pubsub_topic](https://registry.terraform.io/providers/hashicorp/google-beta/latest/docs/resources/pubsub_topic),
You might also need to check the
[REST API doc](https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.topics)
and the
[tutorial doc](https://cloud.google.com/pubsub/docs/create-topic#managing_topics).

After you have a basic understanding of the fields for the target resource, these
approaches can help you identify whether the field maps to a Google Cloud
resource:

1.  In the GCP tutorial for the target resource, if another GCP resource is
    required before creating the target resource, and if the value of a field in
    the other GCP resource is needed when creating the target resource, then
    the field is likely a reference field. Make a note of the field name and the
    other GCP resource name.
1.  In the TF sample for the target resource, if the value of a field is an
    attribute reference of another TF resource, then the field is likely
    a reference field. Make a note of the field name and the other TF resource type
    name, for example the `key_ring` fields mentioned in the
    [cmek sample for PubSubTopic](https://registry.terraform.io/providers/hashicorp/google-beta/latest/docs/resources/pubsub_topic#example-usage---pubsub-topic-cmek).
1.  If the field name includes: `name`, `id`, `link`, `email`,
    `service_account`, resource names under the same API of the target resource,
    or field descriptions include `refer`, `GCP`/`google`/`cloud` resources,
    `url`, `uri`, and so on. These keywords indicate it's likely that these
    fields reference Google Cloud resources. Double-check the Google Cloud tutorial
    and the Terraform documentation to confirm.

Figure out the referenced GCP resource through reading the field description.
Once it's determined, note down the field name and the referenced GCP resource name.

## Identify whether the referenced resource is TF-based or DCL-based

### Turn GCP resource name or TF type name into KCC kind name

Before you can identify what the resource is based on, you must convert
the Google Cloud resource name or Terraform type name into a
Config Connector kind name. Create a GitHub question in this repo for help
if you still can't find the Config Connector kind name after following these steps:

1.  If you know the TF type name, you can do project-scoped search in your local
    repository with the TF type name. If you find a hit in a service mapping
    yaml file (under either the `config/servicemappings/` folder or the
    `scripts/resource-autogen/generated/servicemappings/`), the corresponding
    `kind` field in the same resource config should be the KCC kind.

1.  If you know the GCP resource name, you can either use the search widget to
    find the corresponding TF type in
    [Terraform doc](https://registry.terraform.io/providers/hashicorp/google-beta/latest/docs)
    with the GCP resource name. Then you can find the KCC kind name by following
    the first step above. Alternatively, you can use the search widget to find the
    corresponding KCC kind directly in
    [KCC reference doc](https://cloud.google.com/config-connector/docs/reference/overview).

1.  If you can't find a KCC kind with the steps above, it's possible that the
    referenced resource is not supported in KCC. Create a GitHub question in this repo
    to clarify. If you've determined that the KCC kind is not supported in
    KCC, skip ahead to
    [Configure reference resource in the service mappings](#configure-reference-in-the-service-mappings).

### Determine whether the referenced resource is TF-based(Terraform) or DCL-based(Declarative Client Library)

Based on the KCC kind names of the references that you identified above, follow these
steps to identify if the referenced resource is DCL or TF-based:

1.  Search for the KCC kind name of the referenced resource on the
    [crds folder](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/master/crds).

1.  Open the CRD file containing the KCC kind name. Search for `dcl2crd` and
    `tf2crd` labels. There should be only one label for each CRD.

    1.  If the resource has the `cnrm.cloud.google.com/dcl2crd: "true"` label, it's a
        DCL-based resource.

    1.  If the resource has the `cnrm.cloud.google.com/tf2crd: "true"` label, it's a
        TF-based resource.

## Identify the reference field type

Next, you identify the type of the reference field. In most cases,
the reference field is a string field, but could be another type like a list.
Check the TF schema of the target resource in the TF provider to determine the
type of the reference field. The URL of file is in the format of
`https://github.com/hashicorp/terraform-provider-google-beta/blob/main/google-beta/services/[servicename]/resource_[tf_type_name_without_google_prefix].go`.
You can find a `Type` field for each argument, and two most common types are
`schema.TypeString` and `schema.TypeList`.

1.  If the field is a string, follow
    [Configure reference resource in the service mappings](#configure-reference-in-the-service-mappings)
    to continue configuring the resource reference.
1.  If the field is a list:

    1.  If the field is not required, mark it as an ignored field following
        [this example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/b746248cd5a9b30669380513de8fdc6b4c43018d/config/servicemappings/cloudbuild.yaml#L204).
    1.  If the field is required, create a GitHub question in this repo for
        further discussion.

1.  If the field isn't a string or a list, create a GitHub question in this repo
    for further discussion.

## Configure reference resource in the service mappings

Based on the information identified in the above sections, configure the service
mapping file by providing values for the required fields. Fields can be found in
[service mapping types](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/b746248cd5a9b30669380513de8fdc6b4c43018d/pkg/apis/core/v1alpha1/servicemapping_types.go#L242)

Here is an example configuration:
```
- tfField: disk.source_snapshot_encryption_key.kms_key_service_account
  key: kmsKeyServiceAccountRef
  description: |-
    The service account being used for the encryption request for the
    given KMS key. If absent, the Compute Engine default service account
    is used.
  gvk:
    kind: IAMServiceAccount
    version: v1beta1
    group: iam.cnrm.cloud.google.com
  targetField: email
```

1.  `tfField`: TF field which references a GCP resource. In the above
    example, it's `disk.source_snapshot_encryption_key.kms_key_service_account`.
    Note that it is snake case, and it should be a path if it has parent fields.
1.  `description`: description of the tfField, can be found in the
    [Terraform documentation](https://registry.terraform.io/providers/hashicorp/google-beta/latest/docs).
    If the referenced KCC kind is not supported, prepend `Only external field is
    supported to configure the reference.` and a newline in the description.
    Here is an
    [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/b746248cd5a9b30669380513de8fdc6b4c43018d/config/servicemappings/cloudbuild.yaml#L135).
1.  `key`: the key of the resource being referenced, the convention is to name
    it as `{tfFieldNameInCamelCase}Ref`. Note here we only need the TF field
    name, no need to include parent fields. Example name:
    `kmsKeyServiceAccountRef`.
1.  `targetField`: the referenced resource's Terraform field that will be
    extracted and set as the value of the tfField. In the example above, the
    value of the TF field
    `disk.source_snapshot_encryption_key.kms_key_service_account` (whose
    corresponding field
    `spec.disk.sourceSnapshotEncryptionKey.KmsKeyServiceAccountRef` is a
    reference field in KCC) should be the value of the TF field `email` in the
    corresponding TF type of the `IAMServiceAccount` KCC kind.
1.  `gvk`: Group,Version,Kind of the resource being referenced. Note that an
    alpha only resource cannot be referenced from a v1beta1 resource. After
    an alpha resource is promoted to beta, we can reference to the beta
    version of the resource.

    To identify alpha resources that should not be referenced by a stable
    v1beta1 resource:

    1.  Check the CRD of the referenced resource in
        [crds folder](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/master/crds).
    1.  Search for the value of `cnrm.cloud.google.com/stability-level` in the
        CRD, if it is `alpha`, then it is an alpha resource; otherwise, it
        contains a stable v1beta1 version.

1.  `dclBasedResource`: set to `true` if the referenced resource is a DCL-based
    resource; Do not need to set this field if it not a DCL-based resource (just
    like the example).
