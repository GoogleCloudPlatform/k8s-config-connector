# Eventarc with GCS Trigger - KCC Composition

This repository provides a KCC Compositions approach to setting up an event-driven workflow that triggers a Cloud Workflow when objects are finalized in a Google Cloud Storage (GCS) bucket.

**Key Components**

* **KCC Composition:** Manages the creation of the Eventarc trigger and associated resources.
* **Custom Resource Definition (CRD):** Defines a CRD for `EventarcConfig` to simplify Eventarc trigger configuration.
* **Custom Resource (CR):** Provides a custom resource to specify the Eventarc trigger configuration.
* **Storage Notification:** Configures a `StorageNotification` to publish events from the GCS bucket to a Pub/Sub topic.
* **Service Account:** Creates a service account with necessary permissions for Eventarc and Pub/Sub.

**Prerequisites**

* **Google Cloud Project:**  With necessary APIs enabled (Cloud Storage, Pub/Sub, Eventarc, Cloud Workflows).
* **Kubernetes Cluster:** With Config Connector installed.
* **Service Account:** With IAM permissions to manage GCS buckets, Pub/Sub topics, Eventarc triggers, and Cloud Workflows.
* **Cloud Workflow:** The workflow you want to trigger.

**Deployment Steps**

1. **Apply `sa.yaml`:** Creates the service account and grants it necessary permissions.

   ```bash
   kubectl apply -f sa.yaml -n config-control
  
2. **Apply `eventarcconfigs-crd.yaml`:** Creates the CRD (Custom Resource Definition) for EventarcConfig.

   ```bash
   kubectl apply -f eventarcconfigs-crd.yaml -n config-control
   
3. **Apply `eventarc-composition.yaml`:** Creates the KCC composition and the CRD for EventarcConfig.

   ```bash
   kubectl apply -f eventarc-composition.yaml -n config-control

4. **Apply `facade.yaml`:** Creates the EventarcConfig custom resource, triggering the composition to create the Eventarc trigger, GCS bucket, Pub/Sub topic, and StorageNotification.

   ```bash
   kubectl apply -f facade.yaml -n config-control

To delete the resources, delete the YAML files in reverse order:

```bash
kubectl delete -f facade.yaml -n config-control
kubectl delete -f eventarc-composition.yaml -n config-control
kubectl delete -f eventarcconfigs-crd.yaml -n config-control
kubectl delete -f sa.yaml -n config-control
```