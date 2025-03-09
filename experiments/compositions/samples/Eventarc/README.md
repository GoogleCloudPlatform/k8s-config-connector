# Eventarc with GCS Trigger - KCC Composition

This directory provides a KCC Compositions approach to setting up an event-driven workflow that triggers a Cloud Workflow when objects are finalized in a Google Cloud Storage (GCS) bucket.

**Key Components**

* **KCC Composition:** Manages the creation of the Eventarc trigger and associated resources.
* **Custom Resource Definition (CRD):** Defines a CRD for `EventarcConfig` to simplify Eventarc trigger configuration.
* **Custom Resource (CR):** Provides a custom resource to specify the Eventarc trigger configuration.
* **Storage Notification:** Configures a `StorageNotification` to publish events from the GCS bucket to a Pub/Sub topic.
* **Service Account:** Creates a service account with necessary permissions for Eventarc and Pub/Sub.

**Prerequisites**

* **Kubernetes Cluster:** With Config Connector installed.
* **Service Account:** With IAM permissions to manage GCS buckets, Pub/Sub topics, Eventarc triggers, and Cloud Workflows.
* **Cloud Workflow:** The workflow you want to trigger.

**Why Getters and the Context API are Used**

This configuration utilizes KCC's `GetterConfiguration` and `Context` API to improve resource management and streamline deployments.

* **Getters (`GetterConfiguration`)**

    Getters allow the composition to extract values from resources within the cluster. In this case, they ensure that the `StorageNotification` resource is created only after the GCS bucket and Pub/Sub topic are fully available, preventing dependency errors.

* **Context API (`Context`)**

    The Context API provides a way to define contextual information, such as the project ID, that can be accessed by the composition. This avoids redundant specification of the project ID for each resource and promotes centralized configuration management. 
    *Note that in the `config-control` namespace, the Context API is prepopulated.*

**Deployment Steps**

1. **Apply `eventarcconfigs-crd.yaml`:** Creates the CRD (Custom Resource Definition) for EventarcConfig.

   ```bash
   kubectl apply -f eventarcconfigs-crd.yaml -n config-control
   
2. **Apply `eventarc-composition.yaml`:** Creates the KCC composition for EventarcConfig.

   ```bash
   kubectl apply -f eventarc-composition.yaml -n config-control

3. **Apply `facade.yaml`:** Creates the EventarcConfig custom resource, triggering the composition to create the Eventarc trigger, GCS bucket, Pub/Sub topic, and StorageNotification.

   ```bash
   kubectl apply -f facade.yaml -n config-control

To delete the resources, delete the YAML files in reverse order:

```bash
kubectl delete -f facade.yaml -n config-control
kubectl delete -f eventarc-composition.yaml -n config-control
kubectl delete -f eventarcconfigs-crd.yaml -n config-control
```