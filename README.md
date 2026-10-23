# GCP Config Connector

Config Connector is a Kubernetes add-on that allows customers to manage GCP
resources, such as Cloud Spanner or Cloud Storage, through your cluster's API.

With Config Connector, you can describe GCP resources declaratively using
Kubernetes-style configuration. Config Connector will create new GCP
resources and update existing ones to the state specified by your
configuration, continuously ensuring GCP is kept in sync. The same
resource model is the basis of Istio, Knative, Kubernetes, and the Google Cloud
Services Platform.

As a result, developers can manage their entire application, including both its
Kubernetes components and any GCP dependencies, using the same
configuration and tooling. For example, the same customization or templating
tool can be used to manage test vs. production versions of an application across
both Kubernetes and GCP.

This repository contains the full Config Connector source code, including
controllers, CRDs, installation bundles, and sample resource configurations.

## Architecture & Controller Engines

Config Connector reconciles Kubernetes Resource Model (KRM) custom resources to Google Cloud APIs across three controller engines:

1. **Direct Controllers** (`pkg/controller/direct/`, `apis/`): Modern, Go SDK-native controllers backed by hermetic in-memory tests via MockGCP. Direct controllers provide robust status reporting and a pure Kubernetes-native developer experience. All new features and greenfield resources are built using the Direct approach.
2. **Terraform-Based Controllers** (`third_party/.../terraform-provider-google-beta/`, `config/servicemappings/`): Brownfield controllers leveraging vendored Terraform Google Beta provider code.
3. **DCL-Based Controllers** (`pkg/dcl/`, `declarative-resource-client-library`): Brownfield controllers leveraging the Declarative Resource Client Library.

### Ongoing Migration to Direct Controllers

We are actively migrating legacy Terraform- and DCL-based controllers to Direct Controllers. During this migration, resources maintain 100% field parity and schema compatibility while cutting over to direct Go client SDK reconciliation and hermetic testing.

## Usage

See the official documentation at https://cloud.google.com/config-connector/docs/overview.

- [Config Connector CLI](./docs/cli/README.md): Documentation for the `config-connector` command-line tool.
- [Managing Config Connector with GitOps (ArgoCD)](./docs/guides/gitops-argocd.md): Instructions on deploying and upgrading via GitOps.
- [Choosing an installation type](https://cloud.google.com/config-connector/docs/concepts/installation-types): Guidance on installing Config Connector in cluster mode vs. namespaced mode.
- [Resource reference](https://cloud.google.com/config-connector/docs/reference/overview): Examples and reference specs for supported resources.

## Developer Documentation

The following guides describe how to build, test, and contribute to Config Connector:

- [CONTRIBUTING.md](CONTRIBUTING.md): Contributor onboarding guide, covering hermetic local development workflows across all three tracks (Direct, Legacy TF/DCL, and Direct Migration).
- [Direct Resource Development Guide](./docs/develop-resources/): In-depth deep dives and scenarios for developing Direct resources and controllers.
- [Adding a New Resource](./docs/develop-resources/scenarios/new-resource.md): Step-by-step guidance for adding new Config Connector resources.

## Contributing to Config Connector

Please refer to our [contribution guide](CONTRIBUTING.md) for detailed instructions on getting started with local development and submitting pull requests.
