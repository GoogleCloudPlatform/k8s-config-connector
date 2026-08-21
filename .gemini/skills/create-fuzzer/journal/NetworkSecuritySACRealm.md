# Fuzzer Journal: NetworkSecuritySACRealm

## Observations
- `NetworkSecuritySACRealm` is defined under `v1alpha1`, but the `networksecurity` directory also contains mappers/controllers for `v1beta1` resources (such as `AuthorizationPolicy`).
- Running the `generate-mapper` CLI tool with default arguments will write/generate mappers into `pkg/controller/direct/networksecurity/mapper.generated.go`, overwriting and deleting existing mappers for resources in different versions of the same API group (e.g., `v1beta1`'s `AuthorizationPolicy`).
- To avoid breaking existing generated files, it is highly recommended to hand-write mappers in a separate, dedicated resource file (e.g. `sacrealm_mappers.go`) in `pkg/controller/direct/<service>/` if multiple API versions share the same package directory and default filename.
- Since KRM types use file-scoped imports, we can safely define file-scoped aliases such as `krm` importing `"github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"` in `sacrealm_mappers.go` and `sacrealm_fuzzer.go` without any symbol or import collisions with files that map `v1beta1` types.
