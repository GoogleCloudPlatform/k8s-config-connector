# ComputeDisk Identity and Reference Journal

## Observations & Learnings

- **Zonal and Regional Resource Integration:** `ComputeDisk` represents both zonal and regional disks in GCP. Unlike resources that are strictly zonal, regional, or global, `ComputeDisk` can be either depending on the format of its `location` field (zones have 3 segments split by hyphens, e.g., `us-central1-a`, while regions have 2 segments, e.g., `us-central1`).
- **Dual Template Pattern:** To support both formats under the `IdentityV2` pattern, we registered two separate `gcpurls.Template` formats (`ZonalComputeDiskIdentityFormat` and `RegionalComputeDiskIdentityFormat`) and implemented a custom check `IsZonal()` based on the hyphen-count of the location value.
- **Dynamic URI Formatting:** By leveraging `IsZonal()`, we dynamically construct the correct `String()` representation (using either zones or regions) and `ParentString()` representation, ensuring that zonal and regional disks are handled transparently.
- **No Schema Changes:** The schema of `ComputeDisk` has not been altered, complying strictly with safety protocols. We successfully cross-checked against `status.selfLink` which was already present.
