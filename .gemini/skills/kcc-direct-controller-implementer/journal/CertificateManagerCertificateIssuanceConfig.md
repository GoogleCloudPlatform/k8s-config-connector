# CertificateManagerCertificateIssuanceConfig direct controller implementation journal

## Resource Details
- Kind: `CertificateManagerCertificateIssuanceConfig`
- Group: `certificatemanager.cnrm.cloud.google.com`
- Version: `v1alpha1`

## Discoveries & Learnings

### Immutability of CertificateIssuanceConfig
In GCP, the `CertificateIssuanceConfig` resource has no update API (no `UpdateCertificateIssuanceConfig` RPC exists on the `certificatemanager` client or service). The resource is completely immutable after creation.
Our direct controller implementation handles this by:
1. Returning a no-op/skip-update log statement in the `Update` adapter method.
2. Directly updating the Kubernetes `status` subresource based on the latest state of the resource in the API.

### Static Config Generator Discovery
The script `dev/tasks/generate_static_config.py` generates `pkg/controller/resourceconfig/static_config.go`. It has a discovery heuristic to identify which resources have `direct` controllers by scanning files matching `pkg/controller/direct/**/*_controller.go` for the pattern `:= &krm.Kind{` or `:= &krmv1beta1.Kind{`, and registering the GVK using `RegisterModel((krm|krmv1beta1).KindGVK)`.
Since `CertificateManagerCertificateIssuanceConfig` is in `v1alpha1`, it was imported with the alias `krmcertificatemanagerv1alpha1`. Because of this, the script initially skipped our controller and omitted it from the generated static configuration.
We resolved this by generalizing the python script's regexes to match any alias prefix starting with `krm` followed by optional letters or numbers (i.e. `krm[a-zA-Z0-9]*`). This makes the script fully robust for future migrations in `v1alpha1` or other version packages.
