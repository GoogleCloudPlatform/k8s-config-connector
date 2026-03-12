When a resource has a server-generated ID, we need to store the generated ID into the kubernetes object,
so that we can associate the GCP object with the kubernetes object on subsequent reconciliations.

Legacy controllers would write this ID to spec.resourceID,
which is the same field that users can specify if they want to adopt an existing resource.
However, this makes the field ownership of the spec.resourceID field unclear: did the user specify it, or did KCC write it back?

For greenfield controllers, we are introducing status.externalRef.
All new controllers should write the identity to status.externalRef,
so this field will both consistently give the GCP resource ID, as well as indicating whether a resource is under KCC control.

The question is whether we should also write to spec.resourceID.
To avoid a behavioural change, when moving from terraform to direct, we will write to both status.externalRef _and_ spec.resourceID.
We want to avoid two problems:

* It should be possible to create a GCP resource with the direct controller and then revert back to the terraform controller (without creating a new GCP resource).

* Tooling or other controllers should continue to see the resourceID in the existing spec.resourceID field.

Even if we could address the first problem (for example by teaching terraform to read status.externalRef),
the only way we have to solve the user expectations problem is with time (a "deprecation" period).

Instead, we accept that although we would prefer never to write to spec, we will continue
to write spec.resourceID for server-generated-id resources that we are migrating from DCL or Terraform to Direct.
We do not intend to write to spec for non-server-generated-id resources, nor for greenfield Direct controllers (unless there is a very compelling and documented reason).
This edge case for writing to spec should not be used as an excuse for writing any other spec fields.

We may be able to stop writing spec.resourceID as part of "v1";
we should start communicating now that status.externalRef is the canonical location for the GCP identity
(and that eventually spec.resourceID will no longer be populated).
We should add this to the documentation for the resources.
