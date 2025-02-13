### Operator set up

```bash
$ export MOCKGCP_PATH = <absolute_path>
$ export APIS_PATH = <absolute_path>
$ export GEMINI_API_KEY = <your_key>
# run from main
$ go run main.go --apis-dir=$APIs_PATH --mockgcp-dir=$MOCKGCP_PATH
```

Prompt 1:
>>> Generate an APIs folder (with all the files broken down by file name) for the ApigeeEnvgroupAttachment resource
...
...