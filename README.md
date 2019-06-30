# GCP Config Connector Alpha User Guide

*GCP Config Connector is currently in closed alpha. To gain access, please email [gcp-config-connector-alpha@google.com](mailto:gcp-config-connector-alpha@google.com) with your account information (user/service account email).*

## Introducing GCP Config Connector
GCP Config Connector is a Kubernetes add-on that allows customers to manage GCP resources, such as Cloud Spanner or Cloud Storage, through your cluster's API.

With Config Connector, now you can describe GCP resources declaratively using Kubernetes-style configuration. Config Connector will create any new GCP resources and update any existing ones to the state specified by your configuration, and continuously makes sure GCP is kept in sync. The same resource model is the basis of Istio, Knative, Kubernetes, and the Google Cloud Services Platform.

As a result, developers can manage their whole application, including both its Kubernetes components as well as any GCP dependencies, using the same configuration, and -- more importantly -- *tooling*. For example, the same customization or templating tool can be used to manage test vs. production versions of an application across both Kubernetes and GCP.

## Config Connector Value Proposition

Currently, developers of Kubernetes applications that are accessing GCP resources need to use two separate sets of tooling for managing their deployments. They can use Kubernetes-native tools, such as `kubectl`, for resources that are native to their Kubernetes cluster, but they must then use a totally separate set of deployment tools to manage any other cloud resources. Having these two disjoint approaches reduces developer velocity, introduces additional workflow dependencies, and often requires manual steps, making the overall process more error prone. This added complexity often prevents developers from efficiently inter-changing between using in-cluster and cloud resources across different environments, e.g. using in-cluster MySQL DB for testing and GCP Cloud SQL for production.

Config Connector makes it possible to unify your resource management workflow by creating a single declarative specification spanning both Kubernetes and GCP.  This workflow consolidation accelerates release cycles, makes the deployment process less error prone, and enables both the easy interchanging of in-cluster and cloud resources as well as the creation just-in-time environments that utilize hosted cloud services. These benefits result in cost savings from using ephemeral environments, as well as savings on training costs for DevOps.

For developers of new Kubernetes applications who are deciding whether or not they should use GCP resources, Config Connector offers a way to capture the desired state of the whole application using a single YAML configuration that can be spun up with a single `kubectl` command, lowering the high barrier to entry when considering mixed in-cluster+cloud applications.

For existing Kubernetes applications that are currently using additional tools (such as Cloud Deployment Manager or Terraform) to deploy cloud resources, Config Connector can reduce the number of dependencies in your resource management workflow, saving infrastructure operation and training costs.

## Getting Started with Config Connector

There are two ways to install ConfigConnector Alpha on your Kubernetes Cluster: by downloading tarball archive or by installing GKE cluster with Add-on enabled. After completing prerequistite steps below either follow [tarball installation instructions](#Instructions-for-Tarball-Installation) or [GKE Add-on installation instructions](#Instructions-for-GKE-Addon-Installation).

### Prerequisites
- `kubectl`
- `gcloud` with a configuration pointed to a project you have the **Owner** role on
- If your identity does not have the `cluster-admin` role in your cluster, run
  the following:
  ```bash
  kubectl create clusterrolebinding cluster-admin-binding \
    --clusterrole cluster-admin \
    --user $(gcloud config get-value account)
  ```

### Instructions for Tarball Installation

1. Download a tarball with the installation configuration:
    ```bash
    # Download the archive
    ACCESS_TOKEN=$(gcloud auth print-access-token)
    curl \
      -X GET \
      -H "Authorization: Bearer ${ACCESS_TOKEN}" \
      -sLO \
      --location-trusted \
      https://us-central1-cnrm-eap.cloudfunctions.net\
    /download/latest/infra/install-bundle.tar.gz
    
    # Extract the contents
    tar -xvzf install-bundle.tar.gz
    ```
    
    *If the tarball extraction fails, you are likely not whitelisted for the Config Connector alpha. Please reach out to
    [gcp-config-connector-alpha@google.com](mailto:gcp-config-connector-alpha@google.com) with the account to whitelist.*

1. Install the infrastructure in your cluster:
    ```bash
    kubectl apply -f install-bundle/
    ```
    **NOTE:** If you previously installed Config Connector alpha and wish to
    upgrade, newer versions of the `CustomResourceDefinitions` may cause some of
    the resources to not be applied successfully. You will need to delete the
    previous version of the CRD via `kubectl delete crd [NAME]`, and then rerun
    the installation command above. *This will delete any resources of that
    particular type*.


### Instructions for GKE Addon Installation
1. Work with your Google Contact to whitelist your GCP project to use ConfigConnector. You will need to provide the project id.
1. Use instructions (here)[https://docs.google.com/document/d/1X3SWfLS9J1QuYxJTmcMEnApMAmgPc7WX4Gy8X_ePDes] to config service account that will allow using alpha API.
1. Ensure that you're using the latest version of `gcloud`. To update run `gcloud components update`.
1. Use the following gcloud command to create Kubernetes cluster with ConfigConnector addon enabled.

    ```bash
    CLUSTER_NAME=[my new cluster name]
    ZONE=[my zone]
    gcloud alpha container clusters create ${CLUSTER_NAME} --zone=${ZONE} --addons=ConfigConnector --release-channel rapid
    ```
1. After the cluster is created, verify that ConfigConnector-specific resource definitions were installed by running 
   ```bash
   kubectl get crds
   ```
   you will see CRDs with `cnrm` in their name, such as `sqldatabases.sql.cnrm.cloud.google.com`.

## Config Connector Scenarios

Config Connector Alpha includes two sample applications that demonstrate two key usage scenarios. The first sample explores the eventually-consistent, `kubectl apply -f` scenario, where your entire infrastructure may be defined declaratively in a single file and deployed with a single command. The second sample uses `kustomize`, an example of tooling in the Kubernetes ecosystem, to use the same set of configuration files to deploy an application with a SQL database dependency first using an in-cluster MySQL database, and finally using a Cloud SQL database instance.

### Prerequisites

1. In order to create Google Cloud resources, Config Connector needs an IAM service account to authenticate as. Create a service
account, create a key for the service account, and inject it in the cluster:
    ```bash
    # Replace [PROJECT_ID] with your Google Cloud project's ID
    PROJECT_ID=[PROJECT_ID]
 
    SA_EMAIL="cnrm-system@${PROJECT_ID}.iam.gserviceaccount.com"
    gcloud iam service-accounts create cnrm-system --project ${PROJECT_ID}
 
    # By default, we will give the service account the project Owner role for
    # this project. You may configure this role
    # to be any that you wish. Keep in mind that Config Connector can only manage resources
    # it has the necessary permissions for.
    ROLE="roles/owner"
    gcloud projects add-iam-policy-binding ${PROJECT_ID} \
      --member "serviceAccount:${SA_EMAIL}" \
      --role "${ROLE}"
 
    gcloud iam service-accounts keys create --iam-account "${SA_EMAIL}" ./key.json
    kubectl create secret generic gcp-key --from-file ./key.json --namespace cnrm-system
 
    # Clean up the key.json file from your local system
    rm ./key.json
    ```

1. Before running any of the sample applications, you need to create an identity that can be used by the demo applications and make it available as a secret in your cluster. You may use a different name than the one provided or reuse a service account created previously if it has the Editor role on your project.

    Create an identity that sample applications will use to run:

    ```bash
    SA_EMAIL=cnrm-application-demo@${PROJECT_ID}.iam.gserviceaccount.com

    # NOTE: Creating a service account will fail if it already exists; skip to
    # the key creation below if you already created it previously.
    gcloud iam service-accounts create cnrm-application-demo \
        --project ${PROJECT_ID}
    gcloud projects add-iam-policy-binding $PROJECT_ID \
        --member serviceAccount:${SA_EMAIL} \
        --role roles/editor
    ```

    Create and inject a key for the service account into the cluster:

    ```bash
    gcloud iam service-accounts keys create \
        --iam-account ${SA_EMAIL} \
        ./key.json
    kubectl create secret generic gcp-key --from-file=key.json
```

3. Enable the APIs for the relevant GCP services:
    ```bash
    gcloud services enable pubsub.googleapis.com --project ${PROJECT_ID}
    gcloud services enable spanner.googleapis.com --project ${PROJECT_ID}
    gcloud services enable sqladmin.googleapis.com --project ${PROJECT_ID}
    ```

4. Config Connector conceptually binds a Kubernetes namespace to a GCP project. Either set your default namespace to a namespace whose name matches the GCP project ID you wish to persist resources in, or add the `cnrm.cloud.google.com/project-id` annotation to your namespace:

    ```bash
    kubectl annotate namespace default \
      "cnrm.cloud.google.com/project-id=${PROJECT_ID}" \
      --overwrite
    ```

### Downloading the Samples

Download a tarball with all the sample applications and example YAML files:
```bash
# Download the archive
ACCESS_TOKEN=$(gcloud auth print-access-token)
curl \
  -X GET \
  -H "Authorization: Bearer ${ACCESS_TOKEN}" \
  -sLO \
  --location-trusted \
  https://us-central1-cnrm-eap.cloudfunctions.net\
/download/latest/samples/samples.tar.gz

# Extract the contents
tar -xvzf samples.tar.gz
```

### Kubernetes Microservice Application Backed by GCP Infrastructure - *Bookstore Sample*

The bookstore application is a Kubernetes application consisting of four in-cluster hosted microservices: an Angular web frontend, and three Node.js backend services - *inventory*, *purchases*, and *users*. In addition, these backend micro-services make use of a Cloud Spanner instance, as well as a Cloud Pub/Sub topic and subscription. This sample illustrates how you can deploy a Kubernetes application that uses cloud resources with a single `kubectl` command.

Make sure to follow the steps in the Prerequisites section above.

To run the bookstore app, `kubectl apply` the release configuration YAML file:

```bash
kubectl apply -f samples/apps/bookstore/config/release-configuration.yaml
```

To connect to the web app, wait for the bookstore to be ready:

```bash
# Wait until the database is populated with books, and the bookstore web
# app is accepting traffic...
./samples/apps/bookstore/scripts/populate-bookstore.sh

# Get the external IP of the web app
EXTERNAL_IP=$(kubectl get svc booksfe -o json | \
jq --raw-output ".status.loadBalancer.ingress[0].ip")

# Navigate to the application’s external IP
echo "Navigate to http://${EXTERNAL_IP}"
```

.. and then navigate in your web browser to `http://${EXTERNAL_IP}`

Try navigating to Spanner in the Cloud Console, finding the Bookstore Spanner instance, and deleting some books from the
`inventory-database` books table. Refresh the web app; they should disappear!

Once you've bought a book, you can check the message published on your Pub/Sub topic by running the following:

```bash
 gcloud pubsub subscriptions pull --auto-ack cnrm-subscription
```

Delete all the resources:

```bash
kubectl delete -f samples/apps/bookstore/config/release-configuration.yaml
```

### Enabling Deployment Portability between In-Cluster MySQL and GCP Cloud SQL - *Musicians Sample*

This sample shows how Config Connector enables just-in-time spin-up and tear-down of Kubernetes applications with cloud dependencies. In this sample, we use `kustomize` to easily swap between an in-cluster MySQL database for testing and a Cloud SQL instance for production use, using the same set of configuration files. In the past, this would have involved significant out-of-band configuration, but with Config Connector, cloud resources can now seamlessly integrate into your existing Kubernetes deployment pipeline.

The musicians application consists of a Go web service that is backed by a Cloud SQL Instance. The service can create a
single table whose schema consists of two columns: *name* and *instrument*. 
 
Make sure to follow the steps in the Prerequisites section, along with this additional prerequisite: 
- Install [kustomize](https://github.com/kubernetes-sigs/kustomize). Kustomize will allow us to specify base resources
  for our configuration, and on top of that we can specify patches to modify them or add new ones. Kustomize overlays
  allow us to run this demo in two modes: using GCP Cloud SQL and in-cluster MySQL.
  
#### In-cluster MySQL

You can view the kustomization file, which controls what patches to overlay with kustomize, here:
`samples/apps/musicians/manifests/overlays/test/kustomization.yaml`

Run kustomize to patch the application to use in-cluster MySQL:
```bash
kustomize build samples/apps/musicians/manifests/overlays/test/ | 
kubectl apply -f -
```

Make sure the application’s web service is ready and get the IP address. It might take a few minutes. 
```bash
./samples/apps/musicians/scripts/test-service-wait.sh

# Get the address of the web service
IP=$(kubectl get service musicians-test \
-o=jsonpath='{.status.loadBalancer.ingress[0].ip}') 
```

#### Testing the app

The application is now provisioned. Run through the following script to test it:

```bash
# Create (or reset) musicians table
curl -X POST http://${IP}/reset

# Query the musician information (will return empty result)
curl http://${IP}/musicians

# Create new musicians
curl http://${IP}/musicians -d '{"name": "John", "instrument": "Guitar"}'
curl http://${IP}/musicians -d '{"name": "Paul", "instrument": "Bass Guitar"}'
curl http://${IP}/musicians -d '{"name": "Ringo", "instrument": "Drums"}'
curl http://${IP}/musicians -d '{"name": "George", "instrument": "Lead Guitar"}'

# Query the musicians again
curl http://${IP}/musicians
```

Delete all the resources:
```bash
kustomize build samples/apps/musicians/manifests/overlays/test/ | 
kubectl delete -f -
```

#### Cloud SQL

***NOTE**: Cloud SQL Instance names are reserved for upwards of two months after deletion, so if you’re running this demo
multiple times, you should choose a new Cloud SQL Instance name each time. See the Troubleshooting section for more
info.*

Use kustomize to change the number of replicas and replace in-cluster MySQL with Cloud SQL as the backend. You can view
the kustomization file, which controls the behavior, at
`samples/apps/musicians/manifests/overlays/prod/kustomization.yaml`.

**NOTE**: If your namespace name is not the same as your GCP project's ID, update the `PROJECT_ID` environment variable to the desired project in `samples/apps/musicians/manifests/overlays/prod/musicians-deployment-project-patch.yaml`.

```bash
# Run kustomize and apply the generated deployment
kustomize build samples/apps/musicians/manifests/overlays/prod/ | 
kubectl apply -f - 

# View how many replicas were made
# (Note that the pods will not come up successfully until the Cloud SQL
# instance is ready)
kubectl get pods
```

Wait for all the resources to become ready. This might take a few minutes.

```bash
./samples/apps/musicians/scripts/prod-service-wait.sh

# Get the address of the web service
IP=$(kubectl get service musicians-prod \
-o=jsonpath='{.status.loadBalancer.ingress[0].ip}') 
```

Go through the same steps as [above](#testing-the-app) to test the app. 

Delete all the resources:

```bash
kustomize build samples/apps/musicians/manifests/overlays/prod/ | 
kubectl delete --wait=false -f - # This will clean up asynchronously
```

### Sample YAML Files

Sample YAML files for each of the supported Config Connector resource types may be found in the `samples/resources` directory. You
can create resources from the sample configs using `kubectl apply`. 

Using these samples as a starting point, you may explore further management of GCP resources using Config Connector:
- Define new instances in yaml files
- Using kubectl to create, modify, and delete them
- Managing multiple resources at a time, including dependencies among them
- Use these resources in a broader context of Kubernetes managed infrastructure or applications, e.g., point an
  application using in-cluster MySQL to Google Cloud SQL

### Create Your Own Configuration

To explore the resources included with Config Connector and all supported parameters, and to create your own configuration, you can view the custom resource definitions installed in your cluster:
```bash
# List all available custom resources
kubectl get customresourcedefinition

# Get detailed information about Cloud SQL instance resource (for example):
kubectl get customresourcedefinition \
  sqlinstances.sql.cnrm.cloud.google.com -oyaml
```

In the custom resource definition, notice the spec and status schemas expressed in OpenAPIV3 format. To view schema of
the resource spec, you can run:
```bash
# View the schema of resource spec (Cloud SQL Instance):
kubectl get customresourcedefinition \
  sqlinstances.sql.cnrm.cloud.google.com -ojson \
  | jq --raw-output '.spec.validation.openAPIV3Schema.properties.spec'
```

The spec schema shows all supported (writable) properties, and matches the underlying API with the following exception:
- Fields which reference other resources are named `xyzRef` and use Kubernetes-style references. For example:
    ```yaml
    apiVersion: pubsub.cnrm.cloud.google.com/v1alpha1
    kind: PubSubSubscription
    metadata:
      name: pubsubsubscription-sample
    spec:
      topicRef:
        name: pubsubtopic-sample
      # ...
    ```

You can find more details on the meaning of specific properties in the API documentation for all supported resources:

| Custom Resource Name | Documentation Link |
| -------------------- | ------------------ |
| bigtableclusters.bigtable.cnrm.cloud.google.com | [Bigtable Cluster](https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.clusters) |
| bigtableinstances.bigtable.cnrm.cloud.google.com | [Bigtable Instance](https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances) |
| bigquerydatasets.bigquery.cnrm.cloud.google.com | [BigQuery Dataset](https://cloud.google.com/bigquery/docs/reference/rest/v2/datasets) |
| iampolicies.iam.cnrm.cloud.google.com | [IAM Policy](https://cloud.google.com/iam/reference/rest/v1/Policy) |
| pubsubsubscriptions.pubsub.cnrm.cloud.google.com | [Pub/Sub Subscription](https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.subscriptions) |
| pubsubtopics.pubsub.cnrm.cloud.google.com | [Pub/Sub Topic](https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.topics) |
| redisinstances.redis.cnrm.cloud.google.com | [Cloud Memorystore for Redis](https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances) |
| spannerinstances.spanner.cnrm.cloud.google.com | [Spanner Instance](https://cloud.google.com/spanner/docs/reference/rest/v1/projects.instances) |
| sqlinstances.sql.cnrm.cloud.google.com | [Cloud SQL Instance](https://cloud.google.com/sql/docs/mysql/admin-api/v1beta4/instances) |
| sqldatabases.sql.cnrm.cloud.google.com | [Cloud SQL Database](https://cloud.google.com/sql/docs/mysql/admin-api/v1beta4/databases) |
| storagebuckets.storage.cnrm.cloud.google.com | [Cloud Storage Bucket](https://cloud.google.com/storage/docs/json_api/v1/buckets) |
| storagebucketaccesscontrols.storage.cnrm.cloud.google.com | [Cloud Storage Bucket Access Control](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls) |
| storagedefaultobjectaccesscontrols.storage.cnrm.cloud.google.com | [Cloud Storage Default Object Access Control](https://cloud.google.com/storage/docs/json_api/v1/defaultObjectAccessControls) |

A convenient way to bootstrap your configuration from an existing resource may be using the API Explorer which allows
you to get JSON representation of an existing resource. The JSON representation can be easily converted to a Config Connector YAML
config file based on the custom resource definition schema.

#### Using Existing Resources
Even if you have already created a GCP resource using command line or Google Cloud Console you can manage it using Config Connector.
To do so, create a YAML configuration for your pre-existing resource (make sure to match the resource name and values of
properties that you do not wish to modify) and `kubectl apply` the configuration:

```bash
kubectl apply -f existing_resource.yaml
```

Config Connector will assume control of the resource and update it to match the declared configuration.

Note: if you `kubectl delete` the configuration, Config Connector will delete the underlying GCP resource.

## Config Connector Resource Functionality

GCP services managed by Config Connector feel native to Kubernetes; thus, all the behavior you would expect from native Kubernetes resources is replicated with Config Connector-managed resources, such as eventual consistency, declarative updates, status, and events.

### Eventual Consistency

Config Connector resources are eventually consistent; you may create resources with dependencies in any order, and the system will eventually reconcile the state to your declared intent. For instance, if you create a PubSubSubscription resource before the corresponding PubSubTopic resource, the API server will accept the resource, but the Config Connector controller for PubSubSubscription will wait until the associated topic exists and then create the subscription. This allows for resources and their dependencies to be expressed side-by-side as configuration files, and the system will reconcile the ordering.

The length of time that the system can be an inconsistent state depends on the count and types of resources that you create. Config Connector’s controller uses the Kubernetes watch API to notice changes in resources' declared specification. Usually the controller can actuate changes within seconds. However, the underlying GCP APIs can greatly increase the duration of inconsistency. For example, if you create a single `PubSubTopic`, the system will reach consistency within seconds. However, if you create a `SQLInstance` and a dependent `SQLDatabase`, the system will be inconsistent for upwards of 20 minutes due to the duration of Cloud SQL operations.

### Namespacing

Config Connector resources are namespaced, with namespaces mapping to underlying GCP projects. By default, the namespace name is expected to match the project ID. Cluster admins may map a namespace to an arbitrary GCP project by specifying the desired project ID using the `cnrm.cloud.google.com/project-id` annotation on the Kubernetes `Namespace` resource.

### RBAC & Permissions

Control-plane access control for Config Connector resources is handled by the cluster’s Role-Based Access Control (RBAC) system. Users with the appropriate RBAC permissions to create the Config Connector Kubernetes resources will have indirect permission to create the associated GCP resources. 

For data-plane access control, GCP IAM Service Accounts and permissions are used, so that workloads can continue using the GCP tools and client libraries that they use today. Config Connector includes an `IAMPolicy` resource that can attach IAM roles to supported resources.

For example, to attach IAM roles to a Cloud Pub/Sub topic `pubsubtopic-sample`, try creating the following resource:
```yaml
apiVersion: iam.cnrm.cloud.google.com/v1alpha1
kind: IAMPolicy
metadata:
  name: iampolicy-sample
spec:
  resourceRef:
    apiVersion: pubsub.cnrm.cloud.google.com/v1alpha1
    kind: PubSubTopic
    name: pubsubtopic-sample
  bindings:
    - role: roles/pubsub.admin
      members:
        - user:me@myownpersonaldomain.com
```

If you wish to set project-level IAM role bindings, simply specify `kind: Project` in `resourceRef` field, with no `apiVersion` or `name`.

### Status and Events

Kubernetes resources differentiate between the concept of user-settable fields and read-only fields. The spec clause contains settable fields and its purpose is to allow user to specify desired state of the resource. The status clause contains fields that report current state of the resource.

In addition, Kubernetes controllers log important status changes as Event resources.

Config Connector adheres to both of these Kubernetes standards. After creating a resource, you may view its status and any corresponding events:
```bash
# e.g. “kubectl describe pubsubtopics my-topic”
kubectl describe [KIND] [NAME]
```

To view all the events that occurred in a particular namespace (for both Config Connector and non-Config Connector resources), run:
```bash
kubectl get events --namespace [NAMESPACE]
```

### Resource Abandonment

If you wish to delete a Config Connector resource but not delete the underlying GCP resource, add the following annotation to the desired Config Connector resource before deleting:
```yaml
...
metadata:
  annotations:
    cnrm.cloud.google.com/deletion-policy: abandon
...
```

## Uninstalling ConfigConnector

To easiest way to ensure that ConfigConnector and all its associated resources are deleted is to delete Kubernetes cluster. Please refer to [Resource Abandonment](#Resource-Abandonment) section if you wish to delete a Config Connector resource but not delete the underlying GCP resource.

### Uninstalling on GKE

```bash
gcloud container clusters delete [Cluster Name]
```

### For non-GKE installations
Please refer to your Kubernetes instance vendor documentation to how to delete the underlying Kubernetes cluster.


## Troubleshooting
**Q**: The installation bundle or samples archive seems to be corrupted.

**A**: This is likely an authentication issue. Run the following to determine whether it was an error returned from the server:
```bash
cat [FILENAME]
```

If the contents are an HTML document with a message saying “Authentication required.”, then the server did not receive a valid auth bearer token. Ensure that the following commands print out an access token of the form `ya29.XXXXXXX...`: 
```bash
ACCESS_TOKEN=$(gcloud auth print-access-token)
echo ${ACCESS_TOKEN}
```

If it does, then try running the `curl` command again. If it still does not work, you are likely not a member of the Config Connector Alpha Google Group.

**Q**: The Cloud SQL instance is not provisioning properly on subsequent runs of the musicians demo.

**A**: Cloud SQL Instance names are reserved for upwards of two months after deletion, so if you’re running this demo multiple times, you should choose a new Cloud SQL Instance name each time. To replace the instance name in the configuration files with a new one, run the following:
```bash
OLD_SQL_NAME=musicians-demo # This is the current instance name being used
NEW_SQL_NAME=musicians-demo-2 # Replace this with your desired new name
sed -i'' -e "s/${OLD_SQL_NAME}/${NEW_SQL_NAME}/" samples/apps/musicians/manifests/overlays/prod/musicians-deployment-project-patch.yaml samples/apps/musicians/manifests/overlays/prod/cloud-sql-instance.yaml samples/apps/musicians/scripts/prod-service-wait.sh
```

## Known Issues and Limitations
**Scale**: All resources are managed by a single cnrm-controller-manager controller running in the cnrm-system namespace of your cluster. It will be able to handle a reasonable amount of changes, but extreme scenarios, involving thousands of resources, will increase the duration between the time of change being applied and time actuated on GCP.

**Orphaned Resources**: In the event that the cnrm-controller-manager controller crashes, GCP resources that it was in the middle of creating has a possibility of being orphaned. When the controller restarts, for each KRM resource, it will attempt to find the corresponding GCP resource. In cases of eventual consistency, or for resources that do not allow CNRM to control the name, the cnrm-controller-manager will create a new GCP resource.

**Cloud SQL Instance Name Reservations**: CloudSQL Instance Names are reserved, even after their associated instance is deleted, for upwards of two months. To work around this issue, make sure to change the name of the SQL Instance in your YAML manifest after deleting the associated SQLInstance resource. 

**Resource Coverage**: During the alpha timeframe, only a select number of APIs and resources are supported. Coverage will be scaled out over time.

## FAQ

**Is there a cost to using Config Connector?**

There is no cost to using Config Connector beyond the cost of the cluster it runs in. Customers may incur costs for using the GCP resources they create (i.e. Spanner, etc.) 

## Feedback

If you encounter an issue, or have any questions, please contact the development team:
[gcc-feedback@google.com](mailto:gcc-feedback@google.com). Feedback always welcome!

We are also looking for feedback to help inform our future direction:
- What tools and processes are you currently using to manage GCP services?
- Are there key API features missing?
- Which APIs and resources would you like us to support, in priority order?
- Would you want to have a way of simplifying or restricting resources? Can you provide examples of what you would like
  to see?
