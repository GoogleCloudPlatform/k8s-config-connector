# Directory: hack

This directory contains shell scripts for development and automation tasks. These scripts are not part of the final product but are essential for the development workflow.

## Key Scripts

*   `record-gcp`: This script is used to record GCP API traffic for a test case. It runs the test against a real GCP project and saves the HTTP traffic to a `_http.log` file. Usage: `WRITE_GOLDEN_OUTPUT=1 E2E_GCP_TARGET=real go test ...`
*   `compare-mock`: This script is used to run a test against the mock GCP environment and compare the generated traffic with the golden `_http.log` file. It is used to ensure that the mock GCP layer is behaving correctly. To regenerate the golden logs for a specific area, you can run `hack/compare_mock fixtures/<area>`.
*   `update-codegen.sh`: This script runs the code generators, such as `controller-gen`, to update the generated code (e.g., deepcopy functions, CRDs).
*   `verify-*.sh`: These scripts are used in CI to verify that the code is correctly formatted, generated, and linted.

When you need to perform common development tasks, check this directory for a helpful script.

See also the root `GEMINI.md` and the `dev/GEMINI.md`.
