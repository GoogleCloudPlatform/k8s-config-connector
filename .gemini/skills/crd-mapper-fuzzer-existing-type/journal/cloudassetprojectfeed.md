# CloudAssetProjectFeed Direct KRM Types Journal

## Learnings & Key Observations

1. **Service Group vs API Directory**:
   - Resource `CloudAssetProjectFeed` has group `cloudasset.cnrm.cloud.google.com` and version `v1alpha1`.
   - The directory is `apis/cloudasset/v1alpha1/`, and `apis/cloudasset/generate.sh` orchestrates generation.

2. **Multiple `--service` Packages in `generate-mapper`**:
   - `CloudAssetProjectFeed` uses `Condition *Expr` which references proto message `google.type.Expr` in addition to `google.cloud.asset.v1.Feed`.
   - Adding `--service google.cloud.asset.v1,google.type` to `generate-mapper` in `apis/cloudasset/generate.sh` enables `generate-mapper` to automatically generate `Expr_v1alpha1_FromProto` / `Expr_v1alpha1_ToProto`.

3. **ProjectRef without `kind`**:
   - Baseline CRD `projectRef` schema has `external`, `name`, and `namespace`, but NOT `kind`.
   - Importing `"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"` as `refs` and using `refs.ProjectRef` ensures `projectRef` matches the baseline CRD schema without generating a `+ spec.projectRef.kind=string` diff.

4. **Required Fields in Sub-structs as Pointers**:
   - Defining scalar string fields (like `Expression` in `Expr`, `Topic` in `PubsubDestination`, and `FeedID` in `CloudAssetProjectFeedSpec`) as pointer `*string` with `// +required` tag allows `generate-mapper` to generate standard `direct.LazyPtr` / `direct.ValueOf` assignments cleanly without needing hand-coded mapping functions.
