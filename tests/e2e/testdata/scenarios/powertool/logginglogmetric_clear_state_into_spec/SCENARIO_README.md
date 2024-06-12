This scenario leverages the virtual kind `RunCLI` to mock the LoggingLogMetric
with unmanaged fields and run powertool to migrate `state-into-spec` from
`merge` to `absent`.

Command to run the test:
```
RUN_E2E=1 E2E_KUBE_TARGET=envtest E2E_GCP_TARGET=mock GOLDEN_REQUEST_CHECKS=1 \
  go test -test.count=1 -timeout 600s -v ./tests/e2e \
  -run TestE2EScript/scenarios/powertool/logginglogmetric_clear_state_into_spec
```

How to construct the test:

* script.yaml
  1. Follow the instructions [here](../README.md) to write a `RunCLI` test
     resource with `changeStateIntoSpecOptions.mockGetObject` set to `true`.
* mock_object_with_unmanaged_fields.yaml
  1. Create a DCL-based LoggingLogMetric resource with `state-into-spec: merge`.
     1. Note that the scenario test to migrate `state-into-spec` from `merge` to
        `absent`should be supported during the direct controller migration when
        the DCL-based reconciliation still works.
     2. If the test is added after the migration, you should use a previous
        version of KCC to get the test data.
  2. Update the underlying GCP resource using gcloud or Cloud Console.
  3. Do kubeclient.Get() and pretty print the result in JSON.
  4. Copy the output to `mock_object_with_unmanaged_fields.yaml` file.
* _cli-0-stdout.log
  1. Generate the golden file with `WRITE_GOLDEN_OUTPUT=1` when running the
     test.