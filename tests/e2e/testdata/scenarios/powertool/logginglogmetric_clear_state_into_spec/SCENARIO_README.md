This scenario leverages the test step `PATCH-EXTERNALLY-MANAGED-FIELDS` to add
externally-managed fields into the LoggingLogMetric resource and run powertool
to migrate `state-into-spec` from `merge` to `absent`.

Command to run the test:
```
RUN_E2E=1 E2E_KUBE_TARGET=envtest E2E_GCP_TARGET=mock GOLDEN_REQUEST_CHECKS=1 \
  go test -test.count=1 -timeout 600s -v ./tests/e2e \
  -run TestE2EScript/scenarios/powertool/logginglogmetric_clear_state_into_spec
```

How to construct the test:

*   script.yaml
    1.  Add a YAML configuration for a LoggingLogMetric resource. Note that
        `cnrm.cloud.google.com/state-into-spec: merge` annotation must be
        configured.
    2.  Add a YAML configuration to patch externally-managed fields into the same
        resource.
        1.  Ensure the YAML configuration has the same GVK, name, and/or
            namespace as the first one.
        2.  Add the externally managed fields under spec. Note that you should
            only configure the fields that you want to manage externally. Don't
            add any other fields even if they are required in the CRD schema.
        3.  Add `TEST: PATCH-EXTERNALLY-MANAGED-FIELDS` before this YAML.
    3.  With the virtual kind, `RunCLI`, use the powertool to update the
        `cnrm.cloud.google.com/state-into-spec` annotation from `merge` to
        `absent` and to remove the externally-managed fields.
    4.  Use `TEST: READ-OBJECT-AND-COMPARE-SPEC` and
        `TARGET_STEP_FOR_READ_AND_COMPARE: 1` to read the kube object of
        LoggingLogMetric and compare it with the kube object in step 1 to ensure
        the powertool has successfully removed all the externally-managed
        fields.

*   Files started with `_`
    1.  Generate the golden files with `WRITE_GOLDEN_OUTPUT=1` when running the
        test.
    2.  Ensure `_cli-2-stdout.log` only contains changes for the
        `cnrm.cloud.google.com/state-into-spec` annotation and the
        externally-managed fields.
