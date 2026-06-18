# Journal: CertificateManagerCertificateMapEntry Identity and Reference

## Observations

1. **Global Location and Parent Segments in CAI**:
   - The CAI definition for `CertificateManagerCertificateMapEntry` is:
     `//certificatemanager.googleapis.com/projects/{{PROJECT_ID}}/locations/global/certificateMaps/{{CERTIFICATE_MAP_NAME}}/certificateMapEntries/{{CERTIFICATE_MAP_ENTRY_NAME}}`.
   - The location segment is strictly hardcoded to `global`, matching its parent `CertificateManagerCertificateMap`.
   - The GCP URL Template is defined as:
     `projects/{project}/locations/global/certificateMaps/{certificatemap}/certificateMapEntries/{certificatemapentry}`.

2. **Dependent Parent Identity Integration**:
   - Defining the identity for `CertificateManagerCertificateMapEntry` required resolving and parsing its parent reference `MapRef`.
   - We normalized the `MapRef` using `MapRef.Normalize(...)` and parsed it via `.ParseExternalToIdentity()` to extract `CertificateMap` from the resulting `CertificateManagerCertificateMapIdentity`.

3. **No Schema Changes or Status Check Bypass**:
   - `CertificateManagerCertificateMapEntryStatus` does not contain `externalRef` or `name` fields.
   - We followed the strict safety rule of not modifying the schema to introduce these fields and bypassed status cross-checks within `GetIdentity`.
