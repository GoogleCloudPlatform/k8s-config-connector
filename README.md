# GCP Config Connector

Config Connector is a Kubernetes add-on that allows customers to manage GCP
resources, such as Cloud Spanner or Cloud Storage, through your cluster's API.

With Config Connector, now you can describe GCP resources declaratively using
Kubernetes-style configuration. Config Connector will create any new GCP
resources and update any existing ones to the state specified by your
configuration, and continuously makes sure GCP is kept in sync. The same
resource model is the basis of Istio, Knative, Kubernetes, and the Google Cloud
Services Platform.

As a result, developers can manage their whole application, including both its
Kubernetes components as well as any GCP dependencies, using the same
configuration, and -- more importantly -- *tooling*. For example, the same
customization or templating tool can be used to manage test vs. production
versions of an application across both Kubernetes and GCP.

This repository contains full Config Connector source code. This inlcudes
controllers, CRDs, install bundles, and sample resource configurations.

## Usage

See https://cloud.google.com/config-connector/docs/overview.

For simple starter examples, see the
[Resource reference](https://cloud.google.com/config-connector/docs/reference/overview)
and
[Cloud Foundation Toolkit Config Connector Solutions](https://github.com/GoogleCloudPlatform/cloud-foundation-toolkit/tree/master/config-connector/solutions).

## Building Config Connector

### Recommended Operating System

-   Ubuntu (18.04/20.04)
-   Debian (9/10/11)

### Software requirements

-   [go 1.19+]
-   [git]
-   [make]
-   [jq]
-   [kubebuilder 2.3.1]
-   [kustomize 3.5.4]
-   [kube-apiserver 1.21.0]

### Set up your environment

#### Option 1: Set up an environment in a fresh VM (recommended)

1.  Create an Ubuntu 20.04
    [VM](https://cloud.google.com/compute/docs/create-linux-vm-instance) on
    Google Cloud.

1.  Open an SSH connection to the VM.

1.  Create a new directory for GoogleCloudPlatform open source projects if it
    does not exist.

    ```shell
    mkdir -p ~/go/src/github.com/GoogleCloudPlatform
    ```

1.  Update apt and install build-essential.

    ```shell
    sudo apt-get update
    sudo apt install build-essential
    ```

1.  Clone the source code.

    ```shell
    cd ~/go/src/github.com/GoogleCloudPlatform
    git clone https://github.com/GoogleCloudPlatform/k8s-config-connector
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

1.  Set up a GKE cluster for testing purposes.

    NOTE: `gcp-setup.sh` assumes the VM you are running it from is in a GCP
    project which does not already have a GKE cluster with Config Connector
    already set up.

    ```shell
    ./gcp-setup.sh
    ```

#### Option 2: Set up an environment manually yourself

1.  Install all [required dependencies](#software-requirements)

1.  Add all required dependencies to your `$PATH`.

1.  Set up a [GOPATH](http://golang.org/doc/code.html#GOPATH).

1.  Add `$GOPATH/bin` to your `$PATH`.

1.  Clone the repository:

    ```shell
    cd $GOPATH/src/github.com/GoogleCloudPlatform
    git clone https://github.com/GoogleCloudPlatform/k8s-config-connector
    ```

### Build the source code

1.  Enter the source code directory:

    ```shell
    cd $GOPATH/src/github.com/GoogleCloudPlatform/k8s-config-connector
    ```

1.  Build the controller:

    ```shell
    make manager
    ```

1.  Build the CRDs:

    ```shell
    make manifests
    ```

1.  Build the config-connector CLI tool:

    ```shell
    make config-connector
    ```

### Create a Resource

1.  Enable Artifact Registry for your project.

    ```shell
    gcloud services enable artifactregistry.googleapis.com
    ```

1.  Create a Docker repository. You may need to wait ~10-15 minutes to let
    your cluster get set up after running `make deploy`.

    ```shell
    cd $GOPATH/src/github.com/GoogleCloudPlatform/k8s-config-connector
    kubectl apply -f config/samples/resources/artifactregistryrepository/artifactregistry_v1beta1_artifactregistryrepository.yaml
    ```

1.  Wait a few minutes and then make sure your repository exists in GCP.

    ```shell
    gcloud artifacts repositories list
    ```

    If you see a repository, then your cluster is properly functioning and
    actuating K8s resources onto GCP.

### Make a Code Change

At this point, your cluster is running a CNRM Controller Manager image built on
your system. Let's make a code change to verify that you are ready to start
development.

1.  Edit
    $GOPATH/src/github.com/GoogleCloudPlatform/k8s-config-connector/cmd/manager/main.go.
    Insert the `log.Printf(...)` statement below on the first line of the
    `main()` function.

    ```shell
    package manager

    func main() {
        log.Printf("I have finished the getting started guide.")
        ...
    }
    ```

1.  Build and deploy your change, force a pull of the container image.

    ```shell
    make deploy-controller && kubectl delete pods --namespace cnrm-system --all
    ```

1.  Verify your new log statement is on the first line of the logs for the CNRM
    Controller Manager pod.

    ```shell
    kubectl --namespace cnrm-system logs cnrm-controller-manager-0
    ```

## Contributing to Config Connector

Please refer to our [contribution guide] for more details.

[go 1.19+]: https://go.dev/doc/install
[git]: https://docs.github.com/en/get-started/quickstart/set-up-git
[make]: https://www.gnu.org/software/make/
[jq]: https://stedolan.github.io/jq/
[kubebuilder 2.3.1]: https://github.com/kubernetes-sigs/kubebuilder/releases/tag/v2.3.1
[kustomize 3.5.4]: https://github.com/kubernetes-sigs/kustomize/releases/tag/kustomize%2Fv3.5.4
[kube-apiserver 1.21.0]: https://dl.k8s.io/v1.21.0/bin/linux/amd64/kube-apiserver
[contribution guide]: CONTRIBUTING.md
