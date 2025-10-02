This is the Config Connector project, also known as KCC.

KCC is a set of kubernetes controllers for managing Google Cloud Platform (GCP) resources.  It is OSS under the Apache 2 license.

Each GCP resource maps to a different CRD and controller.

For example, GCP Storage Buckets is managed by the StorageBucket CRD.  The group for StorageBucket is storage.cnrm.cloud.google.com.

KCC has been running for many years, and the older controllers wrap the terraform provider for google, or a library called DCL.
Newer controllers follow the more traditional kubernetes controller pattern, leveraging controller-runtime and making calls to the google cloud SDKs.  We call this approach the "direct" approach.
We are gradually trying to migrate all controllers to the "direct" approach, because the code is much simpler to understand.

However, KCC has a lot of existing users using it at scale.  We want to ensure that the same KCC yaml produces the same GCP resources,
i.e. we do not want to break existing users.  For this reason we must be careful when replacing terraform or DCL controllers with direct controllers.
We have a large and growing test-suite, containing KCC yaml for descibing GCP resources.
We have a mock layer for GCP, so that we can run this test suite without requiring a real GCP account; this lets us inject faults and can be much faster for slow resources.
We can then also run these tests hermetically, which is very handy for running tests against github - we do not need a GCP account.

# GCP Projects and Namespaces

KCC can manage resources in multiple GCP projects.  Typically a platform team will run KCC in a central "platfrom" cluster,
and app teams will each have their own GCP project, and each app team GCP project will be managed in its own kubernetes namespace.

By default, KCC will use the namespace as the GCP project name.  This can be tweaked by setting the `cnrm.cloud.google.com/project-id` annotation
either on a KCC object or on a namespace.  In general though, things work well if there is a 1:1 correspondence between kube namespaces and GCP projects.

# Namespace mode and Cluster mode

KCC has two modes of operation: namespace mode and cluster mode.

In cluster mode, we run one instance of the KCC controller binary for the whole cluster.  It watches for instances of the KCC CRDs in all namespaces,
and creates/updates/deletes the corresponding GCP resources.  Because it is a single instance, it runs as one kubernetes ServiceAccount and a single
GCP ServiceAccount (typically using Workload Identity, but we can also configure a GCP serviceaccount key).

In namespace mode, we run one instance of the KCC controller binary for each "enabled" namespace.  Each instance only watches for KCC CRDs instances
in that namespace.  This lets us run with a kubernetes ServiceAccount per namespace, as well as a GCP ServiceAccount per namespace.  This is more secure,
and also is easier to scale.

There are two CRDs that control the behaviour: ConfigConnector is a cluster-scoped CRD that controls cluster-scoped options.  In particular:
* spec.mode determines whether we run in cluster-mode or namespace-mode.

When running in namespace mode, a namespace is enabled by creating a instance of the ConfigConnectorContext CRD in that namespace.  This acts
as the trigger for watching that namespace, and also allows configuration of things like the GCP ServiceAccount to use for that namespace.

We often abbreviate ConfigConnectorContext to CCC or "triple-C".

# Resources and Controllers

Each resource is represented by a file under `config/crds/resources`.
You can extract the name of the resource by running `cat <file> | yq '.spec.names.kind'` on the file.

A top-level parent controller routes reconciliation to one of three underlying controllers: Terraform (TF), DCL, or Direct. The controller is selected using the following order of precedence:

1.  **Resource Annotation (deprecated):** A resource can specify a controller directly using the annotation `cnrm.cloud.google.com/reconciler: direct`. This is supported for backward compatibility, but its use is discouraged and it will be deprecated in the future.

2.  **ConfigConnectorContext Override:** The `ConfigConnectorContext` resource allows for overriding the controller for a specific resource `GroupKind` using the `spec.experiments.controllerOverrides` field.

3.  **Static Configuration:** A static map in `pkg/controller/resourceconfig/static_config.go` defines the default and supported controllers for each resource. This is the default mechanism if no overrides are specified.

Direct controllers can be found under `pkg/controller/direct`.
The controller will have a file name ending in `_controller.go`.
The controller will call `RegisterModel` using a KRM containing the resource name and ending in GVK.

# Resource Status

Config Connector updates the "status" field to reflect the current state of the resource. To check if a resource is ready, 
inspect its "status.condition":

1. Ready: The resource is successfully reconciled when "status.condition.status" is set to "True" and "status.condition.reason" is "UpToDate".
2. Not Ready: (todo)
3. Error: If "status.condition.status" is "False", the resource is not ready. the "message" and "reason" fields under "status.condition"
will provide additional information.

on the resource's status.

# Resource References

In Config Connector, a resource reference is a mechanism for defining dependencies between resources within Kubernetes configuration. 
This simplifies management by allowing one resource to point to other resources, which Config Connector then resolves its dependencies
automatically. 

To specify resource references in the primary resource's yaml configuration Spec, the reference field's name is the
referenced resource's short name followed by "Ref" suffix. For example:
The reference to a PubSubTopic is "topicRef"; The reference to a StorageBucket is "bucketRef".

There are three primary ways to reference to another resource:

1. Use the "name" field to point to another Config Connector managed resource located in the same Kubernetes namespace.
2. Use both "name" and "namespace" fields to point to another Config Connector managed resource located in a different Kubernetes namespace.
3. use the "external" field to point to a pre-existing Google Cloud resource not managed by Config Connector.



# Options

We have an emerging pattern for configuring options.  The "state-into-spec" option was an early option to demonstrate the pattern.
When we want to configure the behaviour of a KCC object - and in particular when we want to change behaviour - we will often first
make the behaviour opt-in by supporting an annotation on the object.  Because it is opt-in, we do not break existing users,
but we still can unblock the use-case and get user feedback.  As it becomes more concrete, we can add corresponding fields to the ConfigConnectorContext
and the ConfigConnector CRD.  Because the platform team often controls the ConfigConnector and ConfigConnectorContext objects,
this lets the platform team change the default behaviour of KCC without requiring all their app-teams to opt-in on each of their resources.

We typically do not set a default value for these CCC and CC fields, and later if we want to change the opt-in behaviour to be opt-in,
we can set the default to the opt-in value.  We typically continue to allow users to explicitly set the opt-out value, so that
they can have the old behaviour for as long as they want, particularly if the old feature is easy to support, just not recommended.

This strategy lets us introduce new features with minimal risk of breaking users, it lets us get feedback, and it later lets us change the default
behaviour while still giving users a way to opt-out back into the old behaviour.

# Testing Strategy

We use a lot of golden testing.  We have a set of test fixtures rooted in `pkg/test/resourcefixture/testdata/basic`.  They are in directories, often
`<service_name>/<version>/<kind>/<testname>` (for example `pkg/test/resourcefixture/testdata/basic/storage/v1beta1/storagebucket/storagebucketsoftdelete`),
but we have not been 100% consistent on this.

Within a test directory, we typically have `create.yaml` which describes the primary resource that we are testing.  We have `update.yaml`, which describes an
update to make to that primary resource. If the primary resource's configuration contains reference fields, we need a `dependencies.yaml`, which contains 
all dependency resources that are referenced by the primary resource. We create the resources in `dependencies.yaml`, then the resource in `create.yaml`, 
then we run `update.yaml`. We expect the resources to become "ready" at each step of the test.

We capture the logs from the HTTP (and GRPC) traffic to GCP APIs.  This is compared against the "golden traffic" in the `_http.log` file.  We do
perform some normalization to remove volatile values, such as timestamps, server-generated identifiers and complicated hashes.

The core test here is TestAllInSeries under tests/e2e.  We normally run this test first against real GCP (env var `E2E_GCP_TARGET=real`),
write the `_http.log` (env var `WRITE_GOLDEN_OUTPUT=1`), and then commit this.
We then run the tests again against our mockgcp emulation/testing layer for GCP (env var `E2E_GCP_TARGET=mock`),
and often we have to improve our mockgcp layer or the normalization to get the results to be the same.
We have two scripts `hack/record-gcp` and `hack/compare-mock` to help streamline this process.

# Github Issues

When asked to work with github issues, use the `gh issue` tool to read/update issues.

# Import Alias Convention

When promoting a resource from `v1alpha1` to `v1beta1`, we should keep `krm` as the import alias for `v1alpha1` and use `krmv1beta1` for `v1beta1`. This is to minimize the code changes.


# Task-Specific Docs

* `docs/ai/qualify-alpha-for-beta.md` shares tips on how to qualify alpha resources for beta promotion.
* `docs/ai/how-to-promote-resource.md` shares tips on how to promote alpha resources to beta.
* `docs/ai/add-missing-field.md` describes how to add a missing field, for example when the GCP service adds a new field.
* `docs/ai/create-crd-for-existing-terraform-resource.md` describes how to create a CRD for an existing terraform resource.
* `docs/ai/github-workflow.md` describes how to generate github workflows.