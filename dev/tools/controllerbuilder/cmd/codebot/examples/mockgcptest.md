I'm trying to create a test case for mockgcp.

A good test case for mockgcp does the basic operations on a GCP resource by using gcloud to create, list, describe and delete the resource.  It can also do a simple update.

For example, if asked to create a mockgcp test for CRUD operations on a pubsub topic (service=pubsub.googleapis.com, resource=topic), we create the file mockgcp/mockpubsub/testdata/topic/crud/script.yaml with the following contents:

```script.yaml
- exec: gcloud pubsub topics create test-${uniqueId}
- exec: gcloud pubsub topics describe test-${uniqueId}
- exec: gcloud pubsub topics delete test-${uniqueId}
```

Or to create mockgcp test for CRUD operations on a GCS bucket (service=storage.googleapis.com, resource=bucket) we create the file mockgcp/mockstorage/testdata/bucket/crud/script.yaml with the following contents:

```script.yaml
- exec: gcloud storage buckets create gs://test-${uniqueId}
- exec: gcloud storage buckets describe gs://test-${uniqueId}
- exec: gcloud storage buckets delete gs://test-${uniqueId}
```

You should use the CreateFile method to create the script.yaml file in the appropriate directory.  You can use ListFilesInWorkspace to make sure that you are creating a test in a new directory.
   
Please create a test case for mockgcp that tests the service `compute.googleapis.com` and the resource `network`
