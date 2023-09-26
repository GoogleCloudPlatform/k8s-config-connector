# How to Contribute

We'd love to accept your patches and contributions to this project. Source code
in this project is currently being mirrored from internal Google repository. The
team is currently focusing on:

*   Adding more documentation on design, architecture and development flow.
*   Migrating required infrastracture and tooling to this open source project.

As a result please expect low maintainer resourcing to review pull requests and
take contributions. The team will conduct PR reviews and try to upstream
completed PRs on a best-effort basis.

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

### Fork and pull

We follow the
[Fork and pull model](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/getting-started/about-collaborative-development-models#fork-and-pull-model)
in GitHub.

You need to first fork this repository, and you can later open a pull request
to propose changes from your own fork to the **master** branch in this source
repository.

GitHub provides detailed instructions in
[Fork a repo](https://docs.github.com/en/get-started/quickstart/fork-a-repo).

If you're doing this for the first time, the easiest way to get going is with the `gh` CLI tool:

1. Install the `gh` CLI tool following the [instructions](https://github.com/cli/cli#installation)

1. Clone the repo.  We recommend you to create the local clone under the path
    `~/go/src/github.com/GoogleCloudPlatform`. 

    ```shell
    gh repo clone github.com/GoogleCloudPlatform/k8s-config-connector ~/go/src/github.com/GoogleCloudPlatform/k8s-config-connector
    cd ~/go/src/github.com/GoogleCloudPlatform/k8s-config-connector
    ```

1. Create a fork of the repo in your github account (and set up the remotes):

   ```shell
   gh repo fork --remote
   ```

1. You can now inspect the remotes with `git remote -v`, you should see a remote
   named `upstream` that is the main project, and a remote named `origin` that is your fork.
   You will typically work locally, push branches to your fork (`origin`), and then send
   Pull-Requests (PRs) to the main project (`upstream`) when they are ready.


### Run tests against a mock environment

KCC includes some tests that run against a [mock](https://en.wikipedia.org/wiki/Mock_object) version of GCP.  You should be able
to run these tests at this point (though you will need to have [go installed](https://go.dev/doc/install) ).

To run just one test:

```
# Download and set up envtest, a local test version of kubernetes apiserver
export KUBEBUILDER_ASSETS=$(go run sigs.k8s.io/controller-runtime/tools/setup-envtest@master use -p path)

RUN_E2E=1 E2E_KUBE_TARGET=envtest E2E_GCP_TARGET=mock go test ./tests/e2e/ -test.count=1 -v -run TestAllInSeries/samples/privatecacapool
```

This test takes about a minute (after compilation) - and we're working on making it much faster.  mockgcp coverage is currently pretty sparse,
but we are working on adding coverage for more of the GCP APIs.

Successful output should look like this:

```
...
--- PASS: TestAllInSeries (50.71s)
    --- PASS: TestAllInSeries/samples (50.71s)
        --- PASS: TestAllInSeries/samples/privatecacapool (48.92s)
PASS
ok  	github.com/GoogleCloudPlatform/k8s-config-connector/tests/e2e	51.313s
```

### Set up for testing against a real environment

Once you have cloned your forked repo, you can use some helper scripts in the
repo to quickly set up a local dev environment.

1.  Make sure you have [gcloud](https://cloud.google.com/sdk/docs/install)
    installed and configured with a default GCP project.

1.  Make sure you have at least 30 GB of free disk size.

1.  Update apt and install build-essential.

    ```shell
    sudo apt-get update
    sudo apt install build-essential
    ```

1.  Change to environment-setup directory.

    ```shell
    cd ~/go/src/github.com/GoogleCloudPlatform/k8s-config-connector/scripts/environment-setup
    ```

1.  Set up sudoless Docker.

    ```shell
    ./docker-setup.sh
    ```

1.  Exit your current session, then SSH back in to the VM. Then run the
    following to ensure you have set up sudoless docker correctly:

    ```shell
    docker run hello-world
    ```

1.  Install Golang.

    ```shell
    cd ~/go/src/github.com/GoogleCloudPlatform/k8s-config-connector/scripts/environment-setup
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

1.  Now that you have everything set up, you can build your own images and then
    deploy the Config Connector CRDs and workloads (including controller
    manager, webhooks, etc...) into your test GKE cluster. Note deploying 300+
    CRDs into your test cluster can take **a long time** to complete. If you are
    only testing/fixing issues for a few CRDs. You can instead just apply the
    CRDs you are going to work on. As an example, we want to deploy CRD
    `ArtifactRegistryRepositories` because we want to validate creation of this
    resource in the next step. So we can do:

    ```shell
    cd ~/go/src/github.com/GoogleCloudPlatform/k8s-config-connector
    make manifests
    kubectl apply -f config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_artifactregistryrepositories.artifactregistry.cnrm.cloud.google.com.yaml
    ```

    And then we build/push the locally built images and deploy the workloads
    using the command below:

    ```shell
    make deploy-controller
    ```

### Validate your environment

The script `gcp-setup.sh` annotates your `default` namespace in the GKE cluster
with a
[project-id](https://cloud.google.com/config-connector/docs/how-to/organizing-resources/project-scoped-resources#annotate_namespace_configuration)
annotation equals to your default GCP project id in gcloud. This enables Config
Connector to create GCP resources in that default GCP project. We can validate
by creating an Artifact Registry resource through Config Connector.

1.  Enable Artifact Registry for your project.

    ```shell
    gcloud services enable artifactregistry.googleapis.com
    ```

1.  Create a GCP ArtifactRegistryRepository resource. You can check if the
    workloads are ready by: `kubectl get pods -n cnrm-system`

    Then you can create a new ArtifactRegistryRepository resource:

    ```shell
    kubectl apply -f config/samples/resources/artifactregistryrepository/artifactregistry_v1beta1_artifactregistryrepository.yaml
    ```

1.  Wait a few minutes and then make sure your repository exists in GCP.

    ```shell
    gcloud artifacts repositories list
    ```

    If you see a repository named `artifactregistryrepository-sample`, then your
    cluster is properly functioning and actuating K8s resources onto GCP.

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

1.  `kubectl edit statefulset cnrm-controller-manager -n cnrm-system` and scale
    down the replica to 0.
2.  Run `make run` and inspect the output logs.

### Submit a Pull Request

At this point you already knows how to make changes and verify it in your local
dev environment. When you have tested your change and are ready to submit a PR,
you can first validate the change locally:

```
make ready-pr
```

You can then commit your change and make a pull request. See more details
[here](https://docs.github.com/en/get-started/quickstart/contributing-to-projects#making-and-pushing-changes).
