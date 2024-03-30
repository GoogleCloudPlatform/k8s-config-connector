# Example e2e testing for Jinja expander with inline templates

## Pre-requisites
Build the images first:
cd expanders/jinja-expander
make docker-build
make docker-push

cd manifest/inline
make docker-build
make docker-push

## Running test
```
make test
```

The Job autodeletes in 60s.
If the job fails to pull images refer the Troubleshooting section below.


## Troubleshooting

### Job pod failing to start with image pull backoff
Job Pod was failing to start.

```
Events:
  Type     Reason     Age                 From               Message
  ----     ------     ----                ----               -------
  Normal   Scheduled  16m                 default-scheduler  Successfully assigned default/inline-jinja-1-lq57j to gke-cs-hncs4-default-pool-bccba734-wzcr
  Normal   Pulling    14m (x4 over 16m)   kubelet            Pulling image "gcr.io/cdcs-test/manifests-inline:latest"
  Warning  Failed     14m (x4 over 16m)   kubelet            Failed to pull image "gcr.io/cdcs-test/manifests-inline:latest": rpc error: code = Unknown desc = failed to pull and unpack image "gcr.io/cdcs-test/manifests-inline:latest": failed to resolve reference "gcr.io/cdcs-test/manifests-inline:latest": unexpected status from HEAD request to https://gcr.io/v2/cdcs-test/manifests-inline/manifests/latest: 403 Forbidden
  Warning  Failed     14m (x4 over 16m)   kubelet            Error: ErrImagePull
  Warning  Failed     14m (x6 over 16m)   kubelet            Error: ImagePullBackOff
  Normal   BackOff    77s (x64 over 16m)  kubelet            Back-off pulling image "gcr.io/cdcs-test/manifests-inline:latest"

```


Followed instructions from https://cloud.google.com/kubernetes-engine/docs/troubleshooting#gke_ar_permission_error- under `Container Registry` tab.

```
❯ gcloud iam service-accounts list
DISPLAY NAME                            EMAIL                                               DISABLED
Compute Engine default service account  167907884474-compute@developer.gserviceaccount.com  False
❯       gsutil ls
gs://artifacts.cdcs-test.appspot.com/
❯ gs://artifacts.cdcs-test.appspot.com/
❯ gsutil iam ch \
serviceAccount:167907884474-compute@developer.gserviceaccount.com:roles/storage.objectViewer \
  gs://artifacts.cdcs-test.appspot.com/

```

Then deleted the job's pod.
It was recreated by job.

```
Events:
  Type    Reason     Age   From               Message
  ----    ------     ----  ----               -------
  Normal  Scheduled  91s   default-scheduler  Successfully assigned default/inline-jinja-1-zdgvn to gke-cs-hncs4-default-pool-bccba734-wzcr
  Normal  Pulling    91s   kubelet            Pulling image "gcr.io/cdcs-test/manifests-inline:latest"
  Normal  Pulled     89s   kubelet            Successfully pulled image "gcr.io/cdcs-test/manifests-inline:latest" in 1.824681905s (1.824696886s including waiting)
  Normal  Created    89s   kubelet            Created container copyin
  Normal  Started    89s   kubelet            Started container copyin

```
