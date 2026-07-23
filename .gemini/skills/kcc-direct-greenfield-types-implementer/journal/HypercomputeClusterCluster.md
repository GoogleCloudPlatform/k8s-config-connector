# HypercomputeClusterCluster Greenfield Types Implementation Journal

## Observations & Design Choices

1. **Schema Design**:
   - `HypercomputeClusterClusterSpec` represents a Greenfield implementation for `google.cloud.hypercomputecluster.v1.Cluster`.
   - The GCP API defines various resources including `network_resources`, `storage_resources`, `compute_resources`, and `orchestrator` (using Slurm).
   - In accordance with KCC guidelines, we mapped nested reference fields (such as `network` inside `NetworkReferenceObservedState`, `filestore` inside `FilestoreReferenceObservedState`, and `bucket` inside `BucketReferenceObservedState`) to their corresponding KCC strongly-typed reference models (e.g. `computerefs.ComputeNetworkRef`, `filestorev1beta1.FilestoreInstanceRef`, and `storagev1beta1.StorageBucketRef`) to enable seamless K8s controller integration.

2. **Avoiding Protoc Shadowing Issues**:
   - In order to prevent the shadowing issues described in prior CI runs (`Input is shadowed in the --proto_path`), we ensured that no duplicate proto files are checked in under `mockgcp/apis/google/cloud/hypercomputecluster/v1/`.
   - Instead, we relied on fetching the proto definition dynamically from `googleapis` by specifying `PROTO_SHA="cdc919ff596e263f2cc55a9780d2f74633da1ced"` inside `apis/hypercomputecluster/v1alpha1/generate.sh`.

3. **Identity & Reference Pattern**:
   - We verified and confirmed the template format `projects/{project}/locations/{location}/clusters/{cluster}` for `HypercomputeClusterClusterIdentity` under `apis/hypercomputecluster/v1alpha1/`.
   - We implemented identity unit tests in `hypercomputeclustercluster_identity_test.go` and verified they pass correctly.
