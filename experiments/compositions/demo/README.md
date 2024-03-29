# Demo 

## Environment setup

For the demo we would need:
1. A K8s cluster with KCC installed 
2. Dev machine with gcloud, kubectl
3. kubeconfig pointing to the k8s cluster

`poc_setup.sh` is an automated way of setting up an isolated project, creating `Config Controller` instance in the project. 
You could use your own cluster with KCC enabled in it.

## Starting

Make sure to edit the `demo/composition-appteam.yaml` to use your projects, folder and billing accounts.
In terminal 1, `cd demo/` and run `./setup.sh` to install the necessary CRDs and 2 Compositions
In terminal 2, `cd composition/` and run the controller using `make run`.

## Scenario 1

In terminal 1, `kubectl apply -f appteam1.yaml`.
You should see a project `clearing-service` being created in the folder. As well as a GCS bucket in the project.
This sets up the namespace for the project, installs KCC in namespace mode and creates a project for that team/tenant.

## Scenario 2

In terminal 1, `kubectl apply -f cloudsql1.yaml`.
In the `clearing-service` project this creates a CloudSQL instance in primary-secondary mode along with necessary KMS Keys.

