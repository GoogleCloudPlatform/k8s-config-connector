# Journal: DNSManagedZone Identity and Reference Support

## Observations

1. **Schema Integrity and Compatibility:**
   - The baseline CRD for `DNSManagedZone` does not contain `spec.projectRef` or other parent reference fields.
   - To adhere strictly to "Do Not Change the Schema" mandates, we avoided adding any new fields (such as `spec.projectRef`) to the Spec or Status structs.
   - Project ID resolution is performed gracefully using `refs.ResolveProjectID` dynamically in `getIdentityFromDNSManagedZoneSpec`.

2. **Template Variables & gcpurls.Template:**
   - The canonical `cloudassetinventory_names.jsonl` contains:
     `{"resourceType": "dns.googleapis.com/ManagedZone", "nameFormats": ["//dns.googleapis.com/projects/{{PROJECT_ID}}/managedZones/{{ZONE_NUMBER}}"]}`
   - The template format was mapped to: `projects/{project}/managedZones/{managedZone}`.
   - Both variables (`project` and `managedZone`) align with the field names (`Project` and `ManagedZone`) when both are lowercased, avoiding any template/parsing panics at initialization.

3. **Reuse of DNSManagedZoneGVK:**
   - The `DNSManagedZoneGVK` variable is already defined in `managedzone_types.go`.
   - Reusing `DNSManagedZoneGVK` instead of redeclaring it in `dnsmanagedzone_reference.go` prevents naming conflicts and keeps the codebase DRY.
