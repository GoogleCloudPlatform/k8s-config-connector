# WordPress

The WordPress application demonstrates how you can configure a WordPress site powered by GCP MySQL database and using Workload Identity for authentication.

## Prerequisites

1. Create or identify a GCP project.
1. Create or identify a GKE cluster where Config Connector has not yet been installed.
1. [Enable Workload Identity](https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity#enable_workload_identity_on_a_new_cluster) on the cluster where you will install Config Connector.
1. Follow the steps [here](https://cloud.google.com/config-connector/docs/how-to/install-upgrade-uninstall) to install Config Connector

## Steps with Helm

All steps are run from this directory.

1. [Install Helm](https://helm.sh/docs/using_helm/)
1. Review and update the values in `./charts/wordpress-gcp/values.yaml` .
1. Validate and install the sample with Helm

    ```bash
    # validate your chart
    helm lint ./charts/wordpress-gcp/ --set google.projectId=[PROJECT_ID]

    # check the output of your chart
    helm template ./charts/wordpress-gcp/ --set google.projectId=[PROJECT_ID]

    # install your chart
    helm install ./charts/wordpress-gcp/ --set google.projectId=[PROJECT_ID]
    ```

1.  Check the status of your database by running `kubectl describe sqlinstance wp-db`. Once the database is created, obtain the external IP address of your WordPress application by checking `kubectl get svc wordpress-external`. Navigate to this address and validate that you see WordPress installation page.

1. Clean up the installation:

    ```bash
    # list Helm releases
    helm list
    
    # delete release
    helm delete [release_name]

## Steps with kustomize

1. [Install kustomize](https://github.com/kubernetes-sigs/kustomize/blob/master/docs/INSTALL.md)
1. Review and update the values in `overlays/production/`. Note how patches are used to update the values.
1. Install the sample with kustomize

    ```bash
    kustomize build ./kustomize/overlays/production | kubectl apply -f -
    ```

1.  Check the status of your database by running `kubectl describe sqlinstance wp-db`. Once the database is created, obtain the external IP address of your wordpress application by checking `kubectl get svc wordpress-external`. Navigate to this address and validate that you see WordPress installation page.
1. Uninstall the sample:

    ```bash
    kustomize build ./kustomize/overlays/production | kubectl delete -f -
    ```