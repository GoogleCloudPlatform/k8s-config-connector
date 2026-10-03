This scenario is meant to test resource-level pause at creation and gating initial launch:
- Configure ConfigConnector and ConfigConnectorContext in "namespaced" mode.
- Create a KCC resource (ArtifactRegistryRepository) with annotation `cnrm.cloud.google.com/actuation-mode: "Paused"`. Observe no creation/actuation call to GCP while paused.
- Unpause the resource by setting the annotation `cnrm.cloud.google.com/actuation-mode: "Reconciling"`.
- Observe the resource is created and actuated on the cloud provider.
