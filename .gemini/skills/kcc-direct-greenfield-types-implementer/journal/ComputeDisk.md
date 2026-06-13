# ComputeDisk Direct KRM Types Implementation Journal

## Observations & Quirks

- **Zonal vs Regional resource**: `ComputeDisk` is unique because a single Config Connector Kind maps to two different GCP APIs and path formats: Zonal `compute.googleapis.com/Disk` (`projects/{project}/zones/{location}/disks/{disk}`) and Regional `compute.googleapis.com/RegionDisk` (`projects/{project}/regions/{location}/disks/{disk}`).
- **Template Reflection Constraint**: `gcpurls.Template` uses reflection on the target identity struct. If any field in the struct is not of type `string` or another `struct` (such as a `bool` like `IsRegional`), it will panic with `field "IsRegional" in struct *reflect.rtype is not a string or a struct`.
- **Simplifying Identity Structs**: To prevent the reflection panic, avoid adding non-string, non-struct helper fields (like booleans) directly inside the identity struct. Instead, infer these attributes dynamically via struct method receivers (e.g. `i.IsRegional()` by checking if the location is zonal or regional).
- **Core k8s SecretKeyRef Package**: For resources using customer-supplied encryption keys with `SecretKeyRef` fields, the correct `SecretKeyRef` type is in `github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1` (aliased as `k8sv1alpha1`), rather than `pkg/apis/k8s/v1alpha1` which only contains `Condition`.
