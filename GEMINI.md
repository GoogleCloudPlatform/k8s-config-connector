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
update to make to that primary resource.  We often have `dependencies.yaml`, which are other resources that we create in order to perform the test.  We
create the resources in `dependencies.yaml`, then the resource in `create.yaml`, then we run `update.yaml`.  We expect the resources to become "ready"
at each step of the test.

We capture the logs from the HTTP (and GRPC) traffic to GCP APIs.  This is compared against the "golden traffic" in the `_http.log` file.  We do
perform some normalization to remove volatile values, such as timestamps, server-generated identifiers and complicated hashes.

The core test here is TestAllInSeries under tests/e2e.  We normally run this test first against real GCP (env var `E2E_GCP_TARGET=real`),
write the `_http.log` (env var `WRITE_GOLDEN_OUTPUT=1`), and then commit this.
We then run the tests again against our mockgcp emulation/testing layer for GCP (env var `E2E_GCP_TARGET=mock`),
and often we have to improve our mockgcp layer or the normalization to get the results to be the same.
We have two scripts `hack/record-gcp` and `hack/compare-mock` to help streamline this process.


# Task-specific instructions

`docs/ai/promote-resource.md` describes how to promote a resource from alpha to beta.