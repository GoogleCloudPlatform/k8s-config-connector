I'm trying to create a test case for mockgcp.

A good test case for mockgcp does the basic operations on a GCP resource by using gcloud to create, list, describe and delete the resource.  It can also do a simple update.

For example, if asked to create a mockgcp test for the gcloud commands under `gcloud pubsub topics`, we create the file mockpubsub/testdata/topic/crud/script.yaml with the following contents:

```script.yaml
- exec: gcloud pubsub topics create test-${uniqueId}
- exec: gcloud pubsub topics describe test-${uniqueId}
- exec: gcloud pubsub topics delete test-${uniqueId}
```

Or to create mockgcp test for the gcloud commands under `gcloud storage buckets` we create the file mockstorage/testdata/bucket/crud/script.yaml with the following contents:

```script.yaml
- exec: gcloud storage buckets create gs://test-${uniqueId}
- exec: gcloud storage buckets describe gs://test-${uniqueId}
- exec: gcloud storage buckets delete gs://test-${uniqueId}
```

Some hints:

* You should use the CreateFile method to create the script.yaml file in the appropriate directory.  You can use ListFilesInWorkspace to make sure that you are creating a test in a new directory.

* You can run the help command to see the available subcommands, for example you might run `gcloud pubsub topics --help`.  If you want to see the flags for any individual commands, you can run the help for them also, for example you might run `gcloud pubsub topics create --help`.

* You should run the help command for each command you output, to verify the flags and arguments of the commands.

Please create a test case for the gcloud commands under `${GCLOUD_COMMAND}`

When you have completed, please output the name of the test script you have created, in a JSON format like this:

{ "path_to_created_test": "mockstorage/testdata/bucket/crud/script.yaml" }