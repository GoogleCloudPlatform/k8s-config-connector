# Service Catalog Sample - Cloud SQL (MySQL)

This sample demonstrates how to build a simple Kubernetes web service with a CloudSQL backend.

## Prerequisites

Follow the prerequisite steps [here](../README.md#Prerequisites).

## Build the Musicians Image and Deploy

```
make
```

Run the command below until the deployment's `AVAILABLE` column to have a value of `1`.

```shell
kubectl get deployment musicians
```

Run the command below until the `cloudsql-user-service` has a valid value for `EXTERNAL-IP`
```
kubectl get service cloudsql-user-service
```

Find the external IP address of the Kubernetes load balancer and assign it to `${IP}`

```shell
IP=$(kubectl get service cloudsql-user-service -o=jsonpath='{.status.loadBalancer.ingress[0].ip}')
```

## Access the Application

Run through teh following `curl` commands to create a table, add four rows to the table, and query the table.

```shell
# Create (or reset) musicians table:
curl -X POST http://${IP}/reset

# Query the musician information (will return empty result):
curl http://${IP}/musicians

# Create new musicians:
curl http://${IP}/musicians -d '{"name": "John", "instrument": "Guitar"}'
curl http://${IP}/musicians -d '{"name": "Paul", "instrument": "Bass Guitar"}'
curl http://${IP}/musicians -d '{"name": "Ringo", "instrument": "Drums"}'
curl http://${IP}/musicians -d '{"name": "George", "instrument": "Lead Guitar"}'

# Query the musicians again:
curl http://${IP}/musicians
```

## Cleanup

Destroy the resources created for this sample.

```shell
make destroy
```