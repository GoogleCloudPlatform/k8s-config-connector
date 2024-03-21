# Setup Troubleshooting

The following guidelines provide solution to common setup issues.

## Pods fail to pull image
When the cluster is created without providing a service account, a Compute Engine service account is created for the cluster. Users must grant the service account permission to pull images from the project registry.

1.  Find the Compute Engine service account.

    ```shell
    gcloud iam service-accounts list | grep "Compute Engine default service account"
    ```

2.  Grant service account read permission.

    ```shell
    gcloud projects add-iam-policy-binding [PROJECT_ID] \
        --member="[SERVICE_ACCOUNT]"
        --role="roles/storage.objectViewer"
    ```

## Sample Artifact Registry is not created
Make sure that the `cnrm.cloud.google.com/project-id` annotation is replaced with your PROJECT_ID in the sample "artifactregistry_v1beta1_artifactregistryrepository.yaml". More detail can be found in [documentation](https://cloud.google.com/config-connector/docs/how-to/organizing-resources/project-scoped-resources).