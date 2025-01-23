I'm working on a test for my project.  I have captured the golden output from gcloud running against real GCP.  Now I need to verify that the golden output from the test is identical when run against my mocks.

For example, if I am working on a script testdata/topic/crud/script.yaml, then I should run the tests against the mocks by executing this shell command:

`WRITE_GOLDEN_OUTPUT=1 E2E_GCP_TARGET=mock go test ./mockgcptests -run TestScripts/mockpubsub/testdata/topic/crud`

I would then run `git diff`.  There should be no differences.  If there are differences, I either need to update the normalization functions, or I need to fix my mocks.

As a second example, if I am working on a script mockstorage/testdata/topic/bucket/script.yaml, then I should run the tests against the mocks by executing this shell command:

`WRITE_GOLDEN_OUTPUT=1 E2E_GCP_TARGET=mock go test ./mockgcptests -run TestScripts/mockstorage/testdata/bucket/crud`

I would then run `git diff`; again there should be no differences.

When fixing differences:

* Do not change the golden file (_http.log) directly, that should only be updated by running the test.
* If a field is missing a default value, look for a method named like 'populateDefaults' and try to add code to populate a default value.
* If a method or whole service needs to be implemented, create a subtask with CreateSubtask.
* If there are other differences you don't know how to handle, call the "CreateSubtask" function with a description of the differences that you think should be fixed along with any suggestions on how to fix them.

Please try to run the tests with the mocks for the script I just created, called `mockgcp/mockcompute/testdata/address/crud/script.yaml`.  I think the relevant implementation code is in mockcompute/regionaladdress.go