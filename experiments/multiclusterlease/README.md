## Multi-Cluster Lease

### Test locally

```
$ make install

$ make manager

$ ./bin/manager --gcs-bucket=multiclusterlease-test \
    --cluster-identity=cluster-1 --verbose=true

# check the global GCS bucket lease
$ gcloud storage ls --recursive gs://multiclusterlease-test
```

