# Bookstore

The bookstore application contains a web interface 'bookstore' that runs on top of
three micro-services: inventory, purchases, and users. The micro-services are backed by GCP
infrastructure consisting of a PubSub Subscription, PubSub Topic, and Spanner Instance.

## Prerequisites

Follow the prerequisite steps [here](../README.md#Prerequisites).

## Demo
All steps are run from this directory.

Create the infrastructure and deployments.

```
kubectl apply -f config/manifests/.
```

Show the web address of the application. Wait for the `EXTERNAL-IP` column to have a valid value.

```
kubectl get service booksfe
```

Launch the application in your browser

```
address=$(kubectl get service booksfe -o json | jq --raw-output '.status.loadBalancer.ingress[0].ip + ":" + (.spec.ports[0].port|tostring)')
sensible-browser ${address}
```

## Build

### Prerequisites
To build and deploy the images, the following binaries must be in your PATH:
- sed
- make
- docker
- kubectl
- kustomize
- gcloud

To build and run the images for yourself:

```
# Build the images
make docker-build

# Push them to Google Container Registry (GCR)
make docker-push

# Deploy your built images to your GKE cluster
make deploy
```
