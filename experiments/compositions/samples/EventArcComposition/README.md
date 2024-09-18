# Eventarc Trigger with GCS and Cloud Workflow - KCC Composition

This repository contains Kubernetes manifests and KCC compositions to set up an event-driven workflow where object finalization events in a Google Cloud Storage (GCS) bucket trigger a Cloud Workflow.

## Components

The configuration is divided into three files:

* **`gcs.yaml`:**
    * A KCC composition (`cors-bucket`) that uses a custom resource definition (CRD) (`CRBucket`) to create a GCS bucket with lifecycle rules, versioning, and CORS configuration.
    * Includes the CRD definition for `CRBucket`.
    * Defines a `CRBucket` custom resource to configure the bucket.
* **`pubsub.yaml`:**
    * A KCC composition (`pubsub-topic-composition`) that uses a CRD (`CRPubSubTopic`) to create a Pub/Sub topic with labels and schema settings.
    * Includes the CRD definition for `CRPubSubTopic`.
    * Defines a `CRPubSubTopic` custom resource to configure the topic.
* **`eventarc.yaml`:**
    * A KCC composition (`eventarc-trigger-composition`) that uses a CRD (`CREventarcTrigger`) to create an Eventarc trigger that listens for events on the Pub/Sub topic and invokes a Cloud Workflow.
    * Includes the CRD definition for `CREventarcTrigger`.
    * Defines a `CREventarcTrigger` custom resource to configure the trigger.
    * Defines a `StorageNotification` resource that connects the GCS bucket to the Pub/Sub topic.

## Prerequisites

* **Google Cloud Project:** A Google Cloud project with the necessary APIs enabled (Cloud Storage, Pub/Sub, Eventarc, Cloud Workflows).
* **Kubernetes Cluster:** A Kubernetes cluster with Config Connector or the required CRDs for GCS, Pub/Sub, and Eventarc installed.
* **Service Account:** A service account with the necessary IAM permissions to create and manage the resources (GCS buckets, Pub/Sub topics, Eventarc triggers, and Cloud Workflows).
* **Cloud Workflow:** A Cloud Workflow that you want to trigger.

## Deployment

1. **`kubectl` cmds**
```
kubectl apply -f gcs.yaml -n config-control
kubectl apply -f pubsub.yaml -n config-control
kubectl apply -f eventarc.yaml -n config-control
```

2. **Cleanup**
```
kubectl delete -f gcs.yaml -n config-control
kubectl delete -f pubsub.yaml -n config-control
kubectl delete -f eventarc.yaml -n config-control
```