This scenario meant to do the following:
- Apply a CC to instruct KCC to be in "Namespaced" mode.
- Apply a CCC to manage a namespace.
- Apply a KCC resource, in this case, ArtifactRegistryRepository.
- Update the resource with an unrelated annotation. Verify that this generates no GCP HTTP traffic.
- Update the resource to have annotation `cnrm.cloud.google.com/actuation-mode: "Paused"`.
- Update the resource spec field `description` to "description 2" while keeping the "Paused" annotation, using `TEST: APPLY-NO-WAIT`. Verify that zero HTTP calls are sent to GCP (i.e. no traffic is recorded/sent).
- Update the resource annotation `cnrm.cloud.google.com/actuation-mode: "Reconciling"` to resume actuation, using `TEST: WAIT-FOR-HTTP-REQUEST` with `VALUE_PRESENT: "description 2"`. Verify that the annotation change immediately triggers reconciliation against GCP.
