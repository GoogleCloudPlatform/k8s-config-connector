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
