# Journal: CertificateManagerCertificateMap Identity and Reference

## Observations

1. **Global Location in CAI Definition**:
   - The CAI definition for `CertificateManagerCertificateMap` is `//certificatemanager.googleapis.com/projects/{{PROJECT_ID}}/locations/global/certificateMaps/{{CERTIFICATE_MAP_NAME}}`.
   - The location segment in this resource is strictly hardcoded to `global`, rather than containing a placeholder variable like `{{LOCATION}}`.

2. **GCP URL Template Alignment**:
   - To align with the CAI nameFormat and satisfy the `TestRegisteredTemplatesMatchCAI` registration test without adding manual test exceptions under `pkg/gcpurls/registry_test.go`, we defined the template pattern exactly as:
     `projects/{project}/locations/global/certificateMaps/{certificatemap}`.
   - Consequently, the `CertificateManagerCertificateMapIdentity` struct only contains `Project` and `CertificateMap` fields, as the template does not contain a variable `{location}` placeholder.

3. **Status Check Bypass**:
   - `CertificateManagerCertificateMapStatus` does not contain `externalRef` or `name` fields.
   - As per the `SKILL.md` instructions, we did not modify the schema of `Status` and returned the parsed spec identity as-is in `GetIdentity`.
