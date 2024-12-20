# Dataflow Flex Template Deployment on Kubernetes

This directory provides a KCC Compositions approach to deploying Dataflow Flex Templates on Kubernetes. It leverages a custom resource definition (CRD) and a facade file to simplify the deployment process.

**Key Components**

* **KCC Composition:** Manages the creation of the Dataflow Flex Template job and associated resources (e.g., a staging Cloud Storage bucket).
* **Custom Resource Definition (CRD):** Defines a CRD for `DataflowFlexTemplateConfig` to hold essential configuration parameters for the Dataflow job (e.g., project name, region).
* **Facade File:**  A simplified YAML file that references the CR to deploy the Dataflow job.

**Prerequisites**

* **Kubernetes Cluster:** With Config Connector installed.
* **Service Account:** With necessary IAM permissions to create and manage Dataflow jobs, Cloud Storage buckets, and other related resources.
* **Dataflow Flex Template:** The Flex Template you want to deploy, available in a Google Cloud Storage location.
* **Pipeline Definition:** Your pipeline definition file (e.g., a YAML file) stored in a Google Cloud Storage location. **Ensure you upload your `beam.yaml` pipeline definition file to a GCS bucket (`yamlPipelineFilePath`) before proceeding.**

**Deployment Steps**

1. **Apply `dataflowflextemplates-crd.yaml`:** Creates the CRD for `DataflowFlexTemplateConfig`.

   ```bash
   kubectl apply -f dataflowflextemplates-crd.yaml -n config-control

2. **Apply `dataflowflextemplates-composition.yaml`:** Creates the KCC Composition for managing the Dataflow deployment.

   ```bash
   kubectl apply -f dataflowflextemplates-composition.yaml -n config-control
   ```

3. **Apply facade.yaml:**  Applies the facade file, which triggers the Composition to create the Dataflow job and any associated resources.

   ```bash
   kubectl apply -f facade.yaml -n config-control
   ```

4. **Deletion Steps**

To delete the resources, delete the YAML files in reverse order:

   ```bash
   kubectl delete -f facade.yaml -n config-control
   kubectl delete -f dataflowflextemplates-composition.yaml -n config-control
   kubectl delete -f dataflowflextemplates-crd.yaml -n config-control
   ```
