# How to add mockgcp support for a resource

mockgcp is a mock implementation of the GCP APIs.  It allows easier testing of projects like Config Connector.

This document describes how to add support for a new service, by adding a test that uses gcloud.  Because we use gcloud, this process can be self-contained in mockgcp (we don't need to update Config Connector)

Because mockgcp operates indepdently of Config Connector, we refer to services and resources by their proto names, for example for the Config Connector resource ComputeInstance, the service is "compute" and the resource is "instance".

## Create a gcloud test

First begin by creating a test that uses gcloud, and write it to `mockgcp/mock<service>/testdata/<resource>/crud/script.yaml`

A good test case for mockgcp does the basic operations on a GCP resource by using gcloud to create, list, describe and delete the resource.  It can also do a simple update.

For example, if asked to create a mockgcp test for the gcloud commands under `gcloud pubsub topics`, we create the file `mockpubsub/testdata/topic/crud/script.yaml` with the following contents:

```script.yaml
- exec: gcloud pubsub topics create test-${uniqueId}
- exec: gcloud pubsub topics describe test-${uniqueId}
- exec: gcloud pubsub topics delete test-${uniqueId}
```

Or to create mockgcp test for the gcloud commands under `gcloud storage buckets` we create the file `mockstorage/testdata/bucket/crud/script.yaml` with the following contents:

```script.yaml
- exec: gcloud storage buckets create gs://test-${uniqueId}
- exec: gcloud storage buckets describe gs://test-${uniqueId}
- exec: gcloud storage buckets delete gs://test-${uniqueId}
```

Some hints:

* You can run the help command to see the available subcommands, for example you might run `gcloud pubsub topics --help`.
If you want to see the flags for any individual commands, you can run the help for them also, for example you might run `gcloud pubsub topics create --help`.

* You should run the help command for each command you output, to verify the flags and arguments of the commands.

* If you must specify a project, use the --project flag with this variable ${projectId}, for example `gcloud pubsub topics create test-${uniqueId} --project=${projectId}`.

* If the resource requires dependent resources, you should create them in the same script.yaml file.

## Creating a service stub and a normalizer

If this is a brand new service, you will also need to create a normalization function and register it.  Please copy the example from `mockgcp/mocksql/normalize.go` and the service from `mockgcp/mocksql/service.go`; comment out the normalization and the registration of proto services so they can act as placeholders for later.

You will also need to register the service in mockgcp/register.go by adding an anonymous import.

## Running the test and capturing the real GCP output

Now you can run the test and capture the real GCP output (it will be written to `mockgcp/mock<service>/testdata/<resource>/crud/_http.log`

Run `dev/tasks/record-gcp mock<service>/testdata/<resource>/crud` - that should succeed and write the _http.log

Note that the first time you write this, it will fail because it is updating the golden log.  If you run the test again, ideally the output will not change and the test should pass.

If the output changes, that means that GCP returns "volatile" results (like timestamps, or random identifiers).  Update `normalize.go` to replace these volatile values with placeholders.

Once the output is stable (the tests pass), then please create a commit that is something like "golden output for mock<service>/testdata/<resource>/cred"

## Running the test against mockgcp

Now run the test again against mockgcp (hint: `dev/tasks/compare-mock mockgcp/mock<service>/testdata/<resource>/crud`).

You will probably need to implement the resource so that the tests pass.  Create a file called `<resource>.go`, and implement the required proto methods.  Please look at `mockgcp/mockprivateca/capool.go` for an example.

You will also need to ensure that the GRPC services are registered in service.go, both in the `Register` function (which registers it for GRPC) and in `NewHTTPMux` (which registers it for HTTP).

If you run dev/tasks/compare-mock, ideally everything will now pass.  If so, commit with a message like "mockgcp: support for <service> <resource>"
