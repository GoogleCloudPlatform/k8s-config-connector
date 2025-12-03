When a resource has a server-generated ID, we need to store the generated ID into the object.

Legacy controllers would write back to spec.resourceID, which is the same field that users can specify if they want to adopt
an existing resource.  However, this makes the field ownership of the spec.resourceID field unclear: did the user specify it, or did KCC write it back?

Instead, we now write to status.externalRef.  The downside is that this is technically a breaking change between terraform and direct controllers,
so if we create an object with the direct controller we can't go back to using the terraform controller (it will create a second object).
We might also break user workflows if they are expecting to see the ID in the existing spec.resourceID field.

We could instead write to both spec.resourceID _and_ status.externalRef.  If we did this then we could move between terraform and direct freely,
and we would not break existing workflows.  The downside is that we do carry forward the field ownership ambiguity.
Perhaps we could address this as part of moving from v1beta1 to v1.

When we move a server-generated-id controller from terraform/DCL to direct, and stop writing to spec.resourceID,
we should document in the release notes the two gotchas:

* Can't switch freely between legacy and direct controllers

* Look to the status.externalRef field for the server-generated ID.