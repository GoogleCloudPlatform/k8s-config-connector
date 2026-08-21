# Direct vs Terraform vs DCL backed reconcilers

KCC supports multiple resources, and has different implementation strategies for some for historical reasons.
We are moving all resources to Direct controllers, but that will take time.
Some resources are implements using Terraform or DCL, but we want to do that safely without causing any interruption to our users.

The canonical source of information as to whether a resource uses Direct/Terraform/DCL is stored by a label on the CRD.
Our canonical source of CRDs is `config/crds/resources`, for example StorateBucket is at
`config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_storagebuckets.storage.cnrm.cloud.google.com.yaml`

If there is a label `cnrm.cloud.google.com/tf2crd: "true"`, then the resource uses Terraform for reconciliation.
If there is a label `cnrm.cloud.google.com/dcl2crd: "true"`, then the resource uses DCL for reconciliation.
If neither label is present, the resource uses a "Direct" controller.
