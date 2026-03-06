**Title: KCC Project Direction and Focus**

To the KCC Community,

This is an update on the future direction and focus of the Config Connector (KCC) project. As KCC is now a mature project, we are shifting our execution model to prioritize long-term stability, API completeness, and performance.

Our primary objectives are:

*   **Complete API Coverage:** Our main priority is to ensure KCC supports all generally-available GCP services and their corresponding fields.
*   **Stability and Performance:** We will focus on enhancing the scalability and performance of KCC for large-scale deployments, alongside a continued emphasis on bug fixes to ensure the reliability of the core controllers.

**Approach to New Features:**

To maintain the stability of the core project, we will be more conservative about introducing large new features directly into KCC. Significant new capabilities, such as advanced High Availability/Disaster Recovery or dry-run functionality, will be developed as extensions or in separate, related projects. This strategy keeps the KCC core stable and focused, while still allowing for innovation.

**Our Public Roadmap and Feature Request Process:**

We are committed to a fully transparent development process. **GitHub Issues and Milestones are the single source of truth for our roadmap and delivery timelines.** To ensure all requests are tracked and prioritized effectively against our goals of stability and API coverage, **all new feature requests must be submitted as a GitHub issue.**

An item is not considered planned work until it is reflected in a public milestone. This process provides a consistent, up-to-date view for everyone and ensures we can focus engineering time on the priorities visible in the roadmap.

Thank you for your continued support of KCC.

The KCC Team

---

### The direct Config Connector guide is ready

We launched a major improvement to develope the Config Connector resources. This approach significantly enhances the reliability of the Config Connector object reconciliation and provides a more Kubernetes-native developing experience. Learn more in [the Direct resource development guide](./docs/develop-resources/).

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