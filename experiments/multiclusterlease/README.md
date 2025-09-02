## Multi-Cluster Lease

### Running leader election controller locally

```
$ make install

$ make manager

$ ./bin/manager --gcs-bucket=multiclusterlease-test --verbose=true

# check the global GCS bucket lease
$ gcloud storage ls --recursive gs://multiclusterlease-test
```

