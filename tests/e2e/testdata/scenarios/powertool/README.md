A test-only virtual kind, `RunCLI`, is used to test the powertool in the scenario
tests.

Here is an example:

```
kind: RunCLI
args: [ "powertools", "force-set-field", "--namespace=${projectId}", "--kind=StorageBucket", "--name=storagebucket-${uniqueId}", "spec.location=EU" ]
```

There are two fields, `kind` and `args`.
* Value of `kind` must be `RunCLI` in order for the test framework to understand
  the syntax.
* Value of `args` is a list of strings that split the powertool command via
  white spaces. The example representation above means we want to test the
  following command:
  ```
  powertool force-set-field --namespace=${projectId} --kind=StorageBucket \
    --name=storagebucket-${uniqueId} spec.location=EU
  ```

To introduce more options for different sub-commands, you can add more fields to
indicate potential options about how to run the CLI. For example,

```
kind: RunCLI
args: [ "powertools", "change-state-into-spec", "--namespace=${projectId}", "--kind=LoggingLogMetric", "--name=logginglogmetric-merge-${uniqueId}" ]
changeStateIntoSpecOptions:
  mockGetObject: true  # File mock_object_with_unmanaged_fields.yaml must contain the mock object.
```

`changeStateIntoSpecOptions.mockGetObject` is set to `true`, which indicates
that the scenario test will mock the behavior of GetObject(). The mocked object
should be found in `mock_object_with_unmanaged_fields.yaml`.

Note: The current UI of the test is suboptimal. We should improve it over time
whenever we have a better design.