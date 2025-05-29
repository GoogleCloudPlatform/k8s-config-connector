# How to Contribute

We'd love to accept your patches and contributions to this project. We use this
GitHub project as our primary source of truth and the main development
repository for Config Connector. The source code in this project is also
mirrored to internal Google repository for the purposes of releases.

## Contributor License Agreement

Contributions to this project must be accompanied by a Contributor License
Agreement. You (or your employer) retain the copyright to your contribution;
this simply gives us permission to use and redistribute your contributions as
part of the project. Head over to <https://cla.developers.google.com/> to see
your current agreements on file or to sign a new one.

You generally only need to submit a CLA once, so if you've already submitted one
(even if it was for a different project), you probably don't need to do it
again.

## Code reviews

All submissions, including submissions by project members, require review. We
use GitHub pull requests for this purpose. Consult
[GitHub Help](https://help.github.com/articles/about-pull-requests/) for more
information on using pull requests.

## Community Guidelines

This project follows
[Google's Open Source Community Guidelines](https://opensource.google.com/conduct/).

## Set up your DEV environment

You need to set up your own DEV environment before contributing to this project.

We follow the typical contribution flow similar to most OSS projects on GitHub.

### Configure Variables

Export the `GITHUB_USERNAME` environment variable which will be used in subsequent
steps.

```
export GITHUB_USERNAME=YOUR_USERNAME
```

### Fork and pull

We follow the
[Fork and pull model](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/getting-started/about-collaborative-development-models#fork-and-pull-model)
in GitHub.

You need to first fork this repository, and you can later on open a pull request
to propose changes from your own fork to the **master** branch in this source
repository.

GitHub provides detailed instructions in
[Fork a repo](https://docs.github.com/en/get-started/quickstart/fork-a-repo). In
summary, you perform the follow steps to get your fork ready:

1.  Set up Git and authentication with GitHub.com.

    https://docs.github.com/en/get-started/quickstart/set-up-git

2.  Fork the `k8s-config-connector` repo. Instructions below assumes you also
    name your fork as `k8s-config-connector`. If you use a different name for
    the fork, you should replace the commands with the right name.

    https://docs.github.com/en/get-started/quickstart/fork-a-repo#forking-a-repository

3.  Clone your forked repo to your dev machine.

    https://docs.github.com/en/get-started/quickstart/fork-a-repo#cloning-your-forked-repository

    We recommend you to create the local clone under the path
    `~/go/src/github.com/$GITHUB_USERNAME`. This will help to avoid a few known
    build frictions related to generated code.

    ```shell
    mkdir -p ~/go/src/github.com/$GITHUB_USERNAME
    cd ~/go/src/github.com/$GITHUB_USERNAME
    git clone https://github.com/$GITHUB_USERNAME/k8s-config-connector   # If you use ssh key auth, this will be git@github.com:$GITHUB_USERNAME/k8s-config-connector.git
    ```

### Set up your environment

> Note: Some of the environment setup scripts will try to write to the bash `~/.profile` and source it. You may need to modify those lines to suit your shell environment.

Once you have cloned your forked repo, you can use some helper scripts in the
repo to quickly set up a local dev environment.

1.  Make sure you have [gcloud](https://cloud.google.com/sdk/docs/install)
    installed and configured with a default GCP project. Confirm that you
    have the role of either an editor or owner in this project.

1.  Make sure you have at least 30 GB of free disk size.

1.  Update apt and install build-essential.

    ```shell
    sudo apt-get update
    sudo apt install build-essential
    ```

1.  Change to environment-setup directory.

    ```shell
    cd ~/go/src/github.com/$GITHUB_USERNAME/k8s-config-connector/scripts/environment-setup
    ```

1.  Set up sudoless Docker.

    ```shell
    ./docker-setup.sh
    ```

1.  Exit your current session, then SSH back into the VM. Then run the
    following to ensure you have set up sudoless docker correctly:

    ```shell
    docker run hello-world
    ```

1.  Install Golang.

    ```shell
    cd ~/go/src/github.com/$GITHUB_USERNAME/k8s-config-connector/scripts/environment-setup
    ./golang-setup.sh
    source ~/.profile
    ```

1.  Install other build dependencies.

    ```shell
    ./repo-setup.sh
    source ~/.profile
    ```

1.  Set up a GKE cluster for testing purposes. This script takes **a long time**
    to run, it assumes there is a GKE cluster named "cnrm-dev" in your default
    GCP project configured through gcloud, and creates one if it doesn't exist.

    If you prefer to use an existing GKE cluster, you can modify `CLUSTER_NAME`
    in the script and use the existing cluster name instead, which will reduce
    the time it takes. Make sure the existing GKE cluster has
    [workload identity](https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity#enable)
    enabled.

    ```shell
    ./gcp-setup.sh
    ```
1.  _(Optional)_ Verify that worload identity federation is [setup correctly](https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity#verify_the_setup).

1.  Now that you have everything set up, you can build your own images and then
    deploy the Config Connector CRDs and workloads (including controller
    manager, webhooks, etc...) into your test GKE cluster.

    1.  Note deploying 300+ CRDs into your test cluster can take **a long time**
        to complete. If you are only testing/fixing issues for a few CRDs. You
        can instead just apply the CRDs you are going to work on. As an example,
        we want to deploy CRD `ArtifactRegistryRepositories` because we want to
        validate creation of this resource in the next step. So we can do:

        ```shell
        cd ~/go/src/github.com/$GITHUB_USERNAME/k8s-config-connector
        make manifests
        kubectl apply -f config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_artifactregistryrepositories.artifactregistry.cnrm.cloud.google.com.yaml
        ```

    1.  We need to install the following two CRDs as they are hard dependencies
        to reconcile all the other supported CRDs:
        ```shell
        kubectl apply -f operator/config/crd/base/bases/core.cnrm.cloud.google.com_configconnectors.yaml
        kubectl apply -f operator/config/crd/base/bases/core.cnrm.cloud.google.com_configconnectorcontexts.yaml
        ```

    1.  Then we build/push the locally built images and deploy the workloads
        using the command below:

        ```shell
        make deploy-controller
        ```

    1. If you want to install config connector on a brand new GKE cluster, the following command will install all CRDs, locally build, push and deploy all workloads to a standard GKE cluster.

        ```shell
        make deploy-kcc-standard
        make install
        ```
    
        For autopilot clusters, please use the following command.

        ```shell
        make deploy-kcc-autopilot
        make install
        ```

### Validate your environment

The script `gcp-setup.sh` annotates your `default` namespace in the GKE cluster
with a
[project-id](https://cloud.google.com/config-connector/docs/how-to/organizing-resources/project-scoped-resources#annotate_namespace_configuration)
annotation equal to your default GCP project id in gcloud. This enables Config
Connector to create GCP resources in that default GCP project. We can validate
by creating an Artifact Registry resource through Config Connector.

1.  Enable Artifact Registry for your project.

    ```shell
    gcloud services enable artifactregistry.googleapis.com
    ```

2.  Create a GCP ArtifactRegistryRepository resource. You can check if the
    workloads are ready by: `kubectl get pods -n cnrm-system`

    Then you can create a new ArtifactRegistryRepository resource:

    ```shell
    kubectl apply -f config/samples/resources/artifactregistryrepository/artifactregistry_v1beta1_artifactregistryrepository.yaml
    ```

3.  Wait a few minutes and then make sure your repository exists in GCP.

    ```shell
    gcloud artifacts repositories list
    ```

    If you see a repository named `artifactregistryrepository-sample`, then your
    cluster is properly functioning and actuating K8s resources onto GCP.

### Setup Troubleshooting

#### Looking for error logs

You can look for error logs by checking the controller logs following the [troubleshooting](https://cloud.google.com/config-connector/docs/troubleshooting#check-controller-logs).

#### Pods fail to pull image
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

#### Sample Artifact Registry is not created
Make sure that the `cnrm.cloud.google.com/project-id` annotation is replaced with your PROJECT_ID in the sample "artifactregistry_v1beta1_artifactregistryrepository.yaml". More detail can be found in [documentation](https://cloud.google.com/config-connector/docs/how-to/organizing-resources/project-scoped-resources).

#### Error getting ConfigConnectorContext object

```
kubectl apply -f operator/config/crd/base/bases/core.cnrm.cloud.google.com_configconnectors.yaml
kubectl apply -f operator/config/crd/base/bases/core.cnrm.cloud.google.com_configconnectorcontexts.yaml
make deploy-controller && kubectl delete pods --namespace cnrm-system --all
```

### Make a Code Change

At this point, your cluster is running a CNRM Controller Manager image built on
your system. Let's make a code change to verify that you are ready to start
development.

Edit cmd/manager/main.go in your local repository. Insert the `log.Printf(...)`
statement below on the first line of the `main()` function.

```shell
package manager

func main() {
    log.Printf("I have finished the getting started guide.")
    ...
}
```

To apply the change, you can either deploy the container image into the GKE
Cluster, or run the Controller Manager directly as a local executable.

#### Build and Deploy the Controller Manager into the GKE Cluster

Build and deploy your change, force a pull of the container image.

```
make deploy-controller && kubectl delete pods --namespace cnrm-system --all
```

Verify your new log statement is on the first line of the logs for the CNRM
Controller Manager pod.

```
kubectl --namespace cnrm-system logs cnrm-controller-manager-0
```

#### Build and Run the Controller Manager locally

If you don't want to deploy the controller manager into your dev cluster, you
can run it locally on your dev machine with the steps below.

1.  Get credentials for the cnrm-controller-manager service account.

    First, you need to create a long-lived API token.

    ```shell
    kubectl -n cnrm-system apply -f - <<EOF
    apiVersion: v1
    kind: Secret
    metadata:
      name: cnrm-controller-manager-secret
      annotations:
        kubernetes.io/service-account.name: cnrm-controller-manager
    type: kubernetes.io/service-account-token
    EOF
    ```

    Then, create a kubeconfig using the API token.

    ```shell
    set -o errexit

    kubectx=$(kubectl config current-context)
    server=$(kubectl config view -o jsonpath="{.clusters[?(@.name==\"${kubectx}\")].cluster.server}")
    clusterName='cnrm-dev'
    namespace='cnrm-system'
    serviceAccount='cnrm-controller-manager'
    secretName='cnrm-controller-manager-secret'
    ca=$(kubectl --namespace="$namespace" get secret/"$secretName" -o=jsonpath='{.data.ca\.crt}')
    token=$(kubectl --namespace="$namespace" get secret/"$secretName" -o=jsonpath='{.data.token}' | base64 --decode)

    cat << EOF >> ~/.kube/cnrm-dev-controller-manager
    ---
    apiVersion: v1
    kind: Config
    clusters:
      - name: ${clusterName}
        cluster:
          certificate-authority-data: ${ca}
          server: ${server}
    contexts:
      - name: ${serviceAccount}@${clusterName}
        context:
          cluster: ${clusterName}
          namespace: ${namespace}
          user: ${serviceAccount}
    users:
      - name: ${serviceAccount}
        user:
          token: ${token}
    current-context: ${serviceAccount}@${clusterName}
    EOF
    ```

2.  `kubectl edit statefulset cnrm-controller-manager -n cnrm-system` and scale
    down the replica to 0.

3.  Run `KUBECONFIG=~/.kube/cnrm-dev-controller-manager make run` and inspect the output logs.

#### Test your changes

If you are adding a new resource, you need to follow the steps in [NewResourceFromTerraform.md](README.ChangingTerraform.md)
to make code changes, add test data, and run the tests for your resource.

If you are working on a existing resource, test yaml should exist under
./pkg/test/resourcefixture/testdata/basic, you can run the test command directly
to make sure the test can still pass. Example command:

```bash
# Export the environment variables needed in the dynamic tests if you haven't done it.
TEST_FOLDER_ID=123456789 go test -v -tags=integration ./pkg/controller/dynamic/ -test.run TestCreateNoChangeUpdateDelete -run-tests cloudschedulerjob -timeout 900s
```
Replace `cloudschedulerjob` with your test target.

### Submit a Pull Request

At this point you already knows how to make changes and verify it in your local
dev environment. When you have tested your change and are ready to submit a PR,
you can first validate the change locally:

```
make ready-pr
```

You can then commit your change and make a pull request. See [GitHub's contributing
to projects: making and pushing changes](https://docs.github.com/en/get-started/quickstart/contributing-to-projects#making-and-pushing-changes).
