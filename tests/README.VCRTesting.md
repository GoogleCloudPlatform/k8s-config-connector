# VCR Testing

VCR are tests run against real GCP to record http interactions and replay them
in future runs. With replay no resources are actually being created/deleted and
no real GCP interactions are made, allows for efficient and reliable testing of
KCC resources.

VCR is controlled via the environment variables E2E_GCP_TARGET and VCR_MODE. To
run VCR test, E2E_GCP_TARGET need to be set to `vcr`.

## Record Mode

We pre-captured vcr logs by set VCR_MODE to `record`. tests runs against real
GCP, vcr logs are created based on the actual API calls and saved under each
test folder, e.g.
pkg/test/resourcefixture/testdata/basic/serviceusage/v1beta1/service/_vcr_cassettes.

Example command:
```
E2E_KUBE_TARGET=envtest RUN_E2E=1 E2E_GCP_TARGET=vcr \
  VCR_MODE=record go test -timeout 3600s -v ./tests/e2e \
  -run TestAllInSeries/fixtures/^[targettest]$
```

Replace `[targettest]` with the direct folder name of the test case, i.e. the
last section of the path to the testdata. E.g. it should be `storagebucket` for
testdata under
`pkg/test/resourcefixture/testdata/basic/storage/v1beta1/storagebucket`.

Test name can be a common substring that can be used by many tests, using regex
`^[targettest]$` ensures that only the specified test will run.

One sub-folder and three files will be generated:
```
├── [targettest]
|   ├── _vcr_cassettes
|       └── tf.yaml
|       └── dcl.yaml
|       └── oauth.yaml
|   ...
```
Requests from different http clients will be saved into different files.

## Replay Mode

We run vcr test as a part of pre-submit job by set VCR_MODE to `replay`. to
playback existing cassettes. Replay mode might fail if there's no existing vcr
cassette to replay.

Example command: `E2E_KUBE_TARGET=envtest RUN_E2E=1 E2E_GCP_TARGET=vcr
VCR_MODE=replay go test -timeout 3600s -v ./tests/e2e -run
TestAllInSeries/fixtures/^[targettest]$`

## Use VCR testing to verify observed state

To verify the observed state is configured properly, we should use VCR testing
and golden object comparison.

Here are the steps to generate the test data:

1. Run the following command to record the vcr cassettes. The test should pass.

    ```
    E2E_KUBE_TARGET=envtest RUN_E2E=1 E2E_GCP_TARGET=vcr VCR_MODE=record \
      go test -timeout 3600s -v ./tests/e2e \
      -run TestAllInSeries/fixtures/^[targettest]$
    ```

1. Run the following command to verify the vcr cassettes. The test should pass.

    ```
    E2E_KUBE_TARGET=envtest RUN_E2E=1 E2E_GCP_TARGET=vcr VCR_MODE=replay \
      go test -timeout 600s -v ./tests/e2e \
      -run TestAllInSeries/fixtures/^[targettest]$
    ```

1.  Delete the existing file with name
    `_generated_object_[targettest].golden.yaml` under the targettest folder.

1.  Run the following command to record the golden object. The test should pass.
    ```
    E2E_KUBE_TARGET=envtest RUN_E2E=1 E2E_GCP_TARGET=vcr VCR_MODE=replay \
      GOLDEN_OBJECT_CHECKS=1 WRITE_GOLDEN_OUTPUT=1 \
      go test -timeout 600s -v ./tests/e2e \
      -run TestAllInSeries/fixtures/^[targettest]$
    ```

1.  Run the following command to verify the vcr cassettes. The test should pass.

    ```
    E2E_KUBE_TARGET=envtest RUN_E2E=1 E2E_GCP_TARGET=vcr VCR_MODE=replay \
      GOLDEN_OBJECT_CHECKS=1 \
      go test -timeout 600s -v ./tests/e2e \
      -run TestAllInSeries/fixtures/^[targettest]$
    ```
