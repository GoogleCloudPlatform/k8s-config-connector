# DNSRecordSet Migration Journal

## Overview
Successfully implemented the direct controller for DNSRecordSet, replacing the legacy Terraform provider. 

## Key Observations & Quirks
1. **API Selection**: We used the `ResourceRecordSets` API (GET/CREATE/PATCH/DELETE) which is synchronous and does not return long-running operations.
2. **Changes API vs ResourceRecordSets**: The legacy Terraform controller uses GCP's bulk `Changes` API (`/changes` endpoint) which was not mocked in `mockgcp`. Since the direct controller uses the standard REST endpoint (`/rrsets/{name}/{type}`), we were able to run mock E2E tests successfully. We explicitly skipped the legacy TF fallback test in `unified_test.go` under MockGCP because of this `/changes` API dependency discrepancy.
3. **Reference Mapping**: `RrdatasRefs` is used in KRM to refer to other objects (like `ComputeAddress`), which is mapped to `Rrdatas` in the API. We updated the hand-written mapper `dnsrecordset_mappers.go` to handle reference resolution for both top-level fields and nested `routingPolicy` fields (geo, backupGeo, wrr).
