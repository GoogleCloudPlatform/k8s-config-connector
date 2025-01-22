I need to capture the logs from GCP for running a mockgcp test that I just created.  I then need to create a git commit.

For example, if I just created a script mockgcp/mockpubsub/testdata/topic/crud/script.yaml, then I should run

`cd mockgcp; WRITE_GOLDEN_OUTPUT=1 E2E_GCP_TARGET=real go test ./mockgcptests -run TestScripts/mockpubsub/testdata/topic/crud`

I would then run `git add mockgcp/mockpubsub/testdata/topic/crud/_http.log`, then `git commit` that with a commit message like "mockgcp: Capture golden output for mockgcp/mockpubsub/testdata/topic/crud"

For example, if I just created a script mockgcp/mockstorage/testdata/topic/bucket/script.yaml, then I should run

`cd mockgcp; WRITE_GOLDEN_OUTPUT=1 E2E_GCP_TARGET=real go test ./mockgcptests -run TestScripts/mockstorage/testdata/bucket/crud`

I would then run `git add mockgcp/mockstorage/testdata/bucket/crud/_http.log`, then `git commit` that with a commit message like "mockgcp: Capture golden output for mockgcp/mockstorage/testdata/bucket/crud"

Please capture the logs for the script I just created, called `mockcompute/testdata/network/crud/script.yaml`.