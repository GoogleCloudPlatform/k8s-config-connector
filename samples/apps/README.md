# Apps

This directory contains sample applications.

## Prerequisites

Before trying any of the samples you must run through the following prerequisites to get your environment configured.

In addition, it is assumed your system has the following.
* cnrm
* gcloud
* kustomize


Create a GKE cluster with CNRM support.

```
cnrm init
```

Install `jq`.

```
sudo apt-get install jq
```

Create service account.

```
PROJECT=$(gcloud config get-value project)
SA_NAME=cnrm-demo
gcloud iam service-accounts create ${SA_NAME}
SA_EMAIL=${SA_NAME}@${PROJECT}.iam.gserviceaccount.com
gcloud projects add-iam-policy-binding $PROJECT  --member serviceAccount:${SA_EMAIL}  --role roles/editor
```

Create a key for your service account and download it to a file.

```
KEY_PATH=/tmp/key.json
gcloud iam service-accounts keys create --iam-account ${SA_EMAIL} ${KEY_PATH}
```

Inject service account key into your cluster.

```
kubectl create secret generic gcp-key --from-file=key.json=${KEY_PATH}
```

## Bookstore

The [bookstore](bookstore) is a simple application with a web user interface.

## Musicians

The [musicians](musicians) sample is a simple web service which uses a Cloud SQL database.