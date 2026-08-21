# Skill: Fix MockGCP Diffs

## Overview

When the golden tests for K8s Config Connector mock output diverge from real GCP output, we need to inspect the discrepancies and fix either the mock implementation or the normalizers. This ensures that `hack/compare-mock` runs cleanly and accurately represents GCP API contracts.

## How to use

1. Look closely at the `compare-mock` HTTP log differences (typically mock on left `=>` real on right).
2. For missing default values (e.g., `<missing> => REGIONAL`), add a `populateDefaultsFor<Resource>` function to the mock service's file (e.g. `mockgcp/mockcompute/networksv1.go`). Make sure it is called on `Insert` and `Get`.
3. For generated IDs or volatile values (e.g. IPs, resource URLs) where real GCP generates dynamically but mockgcp outputs something static, you need to update the normalizer `mockgcp/mock<service>/normalize.go`.
4. Run `hack/compare-mock "fixtures/^<testname>$"` to see the diff and overwrite `_http.log`.
5. Run `git diff` on the test fixtures to ensure that the golden `_http.log` accurately replaces volatile data with placeholder variables (e.g., `${ipAddress}`).
6. Certain operation metadata values (e.g. `done: <missing> => false`) can be safely ignored as mock operations are generally simpler.