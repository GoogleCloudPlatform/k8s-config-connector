
1. Define some variables that describe the command and API you want to mock:

```
export REPO_ROOT="$(git rev-parse --show-toplevel)"
export GCLOUD_COMMAND="gcloud workflows"
```
   
1. Ask codebot to generate the mockgcp test script:

```
cd ${REPO_ROOT}/mockgcp/
cat ${REPO_ROOT}/dev/tools/controllerbuilder/cmd/codebot/examples/mockgcp-test-case/1-create-test.md | \
    envsubst 'GCLOUD_COMMAND=${GCLOUD_COMMAND}' | \
    codebot
```

1. ....