### [2026-10-06] NetAppActiveDirectory Greenfield Direct Controller
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for Greenfield resource `NetAppActiveDirectory`.
- **Problem**:
  1. `netBiosPrefix` in GCP NetApp Volumes must be at most 10 characters (otherwise GCP returns `400: Invalid netbios prefix, length must not be longer than 10 characters`).
  2. GCP defaults `organizational_unit` to `"CN=Computers"` on the server side when omitted.
- **Solution**:
  1. Used concise static prefixes (`adnetbios` / `maxnetbios`) for fixture test configurations.
  2. Handled server default for `organizational_unit` in `compareActiveDirectory` to prevent unwanted diffs and extra updates during re-reconciliation.
- **Impact**: All minimal and maximal fixtures successfully verified, recorded against real GCP, and passed re-reconciliation testing.
