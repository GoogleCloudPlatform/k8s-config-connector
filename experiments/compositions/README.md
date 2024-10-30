# Compositions

## Getting Started

Suggest creating a script to set up your environment.
`export REGISTRY=gcr.io/$(gcloud config get-value project)`

Make sure you `gcloud auth login`
Create a Config Controller cluster.
You then need to disable the 3rd party controller protections.
https://cloud.google.com/anthos-config-management/docs/how-to/stopping-policy-controller#gcloud-configmanagement
To suspend Policy Controller:kubectl edit validatingwebhookconfigurations.admissionregistration.k8s.io gatekeeper-validating-webhook-configurationDelete the webhooks field and everything underneath it and remove the `policycontroller.configmanagement.gke.io/managed-by-operator` label if it exists.

To BUILD the Composition
cd to the composition directory
You need to manually increment the version with each iteration for the image.
`make docker-build docker-push IMG=${REGISTRY}/composition:0.0.1`
`make deploy IMG=${REGISTRY}/composition:0.0.1`

To BUILD the facade example
make manifests
kubectl apply -f composition_role.yaml
kubectl apply -f composition_role_binding.yaml

SAMPLES
each of composition and facade has a sample directory with sample yaml to test
the controllers.


**Note: It can take up to 5 minutes for the `InputReconciler` to create resources accessible by `kubectl` commands. Please wait for that duration to examine if resources have been created as intended.**
