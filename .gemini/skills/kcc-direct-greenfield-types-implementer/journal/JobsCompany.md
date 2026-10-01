# JobsCompany (CloudTalentSolutionCompany) Journal

## Context & Service Naming
The Cloud Talent Solution (Job Search API / `google.cloud.talent.v4` / `jobs.googleapis.com`) resource `Company` is implemented in KCC under the kind `CloudTalentSolutionCompany` in package `apis/cloudtalentsolution/v1alpha1` with group `cloudtalentsolution.cnrm.cloud.google.com`.

## Discrepancy in Automated Greenfield Detection
The automated gap analysis tool `hack/tools/greenfield/calculate_coverage.py` parses googleapis protobuf definitions where the service directory or proto path is identified as `jobs`. Without explicit service mapping in `services_acronyms.json` and `service_aliases` in `calculate_coverage.py`, it was predicting `JobsCompany` and marking it as missing from KCC.

## Solution & Enhancements
1. Added `"cloudtalentsolution": "CloudTalentSolution"`, `"jobs": "CloudTalentSolution"`, and `"talent": "CloudTalentSolution"` to `hack/tools/greenfield/services_acronyms.json`.
2. Added `"jobs": ["cloudtalentsolution"]` and `"talent": ["cloudtalentsolution"]` to `service_aliases` in `hack/tools/greenfield/calculate_coverage.py`.
3. Updated `CloudTalentSolutionCompanyRef.Normalize` in `apis/cloudtalentsolution/v1alpha1/cloudtalentsolutioncompany_reference.go` to directly invoke `refs.Normalize`, adhering to the standard direct reference normalization pattern without legacy fallback.
