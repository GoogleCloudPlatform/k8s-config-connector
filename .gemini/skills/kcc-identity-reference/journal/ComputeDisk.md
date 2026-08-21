# ComputeDisk Identity and Reference Journal

## Observations & Learnings

- **Zonal and Regional Resource Integration:** `ComputeDisk` represents both zonal and regional disks in GCP. Unlike resources that are strictly zonal, regional, or global, `ComputeDisk` can be either depending on the format of its `location` field (zones have 3 segments split by hyphens, e.g., `us-central1-a`, while regions have 2 segments, e.g., `us-central1`).
- **Dual Template Pattern via Struct Mapping:** Following reviewer feedback, instead of keeping a generic `Location` field and using hyphen-splitting within core identity formatting methods, we split `ComputeDiskIdentity` to have separate `Zone` and `Region` fields (where only one is populated).
- **Automatic Parse Binding:** We updated our registered `gcpurls.Template` templates to map `{zone}` and `{region}` placeholders directly. This causes `gcpurls.Template.Parse` to automatically bind the parsed path segments to the correct `Zone` or `Region` field on struct parsing.
- **Dynamic URI Formatting:** By leveraging checks against `i.Zone != ""`, we dynamically format the correct `String()` representation (using either zones or regions) and `ParentString()` representation, keeping the codebase incredibly clean and free of hyphen-splitting boilerplate.
- **No Schema Changes:** The schema of `ComputeDisk` has not been altered, complying strictly with safety protocols. We successfully cross-checked against `status.selfLink` which was already present.
