I need to capture the logs from GCP for running a mockgcp test that I just created.  I then need to create a git commit.

For example, if I just created a script mockpubsub/testdata/topic/crud/script.yaml, then I should run

`WRITE_GOLDEN_OUTPUT=1 E2E_GCP_TARGET=real go test ./mockgcptests -run TestScripts/mockpubsub/testdata/topic/crud`

I would then run `git add mockpubsub/testdata/topic/crud/_http.log`, then `git commit` that with a commit message like "mockgcp: Capture golden output for mockpubsub/testdata/topic/crud"

For example, if I just created a script mockstorage/testdata/topic/bucket/script.yaml, then I should run

`WRITE_GOLDEN_OUTPUT=1 E2E_GCP_TARGET=real go test ./mockgcptests -run TestScripts/mockstorage/testdata/bucket/crud`

I would then run `git add mockstorage/testdata/bucket/crud/_http.log`, then `git commit` that with a commit message like "mockgcp: Capture golden output for mockstorage/testdata/bucket/crud"

Please capture the logs for the script I just created, called `mockworkflows/testdata/workflow/crud/script.yaml`.

When you are done, please output a JSON result like this:

{ "status": "success" }


If you have problems, please output a JSON result like this:

{ "status": "failure", "reason": "Fill in any information on why you could not complete the task" }