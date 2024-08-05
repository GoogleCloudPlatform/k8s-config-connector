### ** Coming soon: a new way to write the controller **

We are developing a faster, more Kubernetes native and more reliable Config Connector reconciliation approach(Direct Controller).
Thus, we decide to hold off any PRs about adding new TF-based or DCL-based resources until mid Q4 2024. 
We will share more guidance about resource coverage in Q4 2024. Please stay tuned!


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
configuration, and more importantly **tooling**. For example, the same
customization or templating tool can be used to manage test vs. production
versions of an application across both Kubernetes and GCP.

This repository contains full Config Connector source code. This includes
controllers, CRDs, install bundles, and sample resource configurations.

## Usage

See https://cloud.google.com/config-connector/docs/overview.

See
[Choosing an installation type](https://cloud.google.com/config-connector/docs/concepts/installation-types)
to decide how you want to install Config Connector.

For simple starter examples, see the
[Resource reference](https://cloud.google.com/config-connector/docs/reference/overview).

## Contributing to Config Connector

Please refer to our [contribution guide](CONTRIBUTING.md) for more details.

