# GCP Config Connector Samples

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

This repository contains sample applications and resources for use with Config
Connector.

For more information about Config Connector, see
https://cloud.google.com/config-connector/docs/overview.

