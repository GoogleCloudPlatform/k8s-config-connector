# Config Connector CLI (config-connector)

The `config-connector` CLI provides tools for managing Config Connector resources, such as exporting, previewing, and bulk-exporting resource definitions.

## Installation

The CLI can be installed as a component of the Google Cloud SDK (gcloud).

### Prerequisites

- [Google Cloud SDK (gcloud)](https://cloud.google.com/sdk/docs/install) installed and configured.

### Install via gcloud component

Run the following command to install the `config-connector` component:

```bash
gcloud components install config-connector
```

Once installed, the `config-connector` command should be available in your PATH.

## Commands

- [preview](../../pkg/cli/cmd/preview/README.md): Preview reconciliation behavior and validate migrations.
- **export**: Export existing GCP resources into Config Connector YAML.
- **bulk-export**: Export all GCP resources in a project, folder, or organization.
- **print-resources**: Print the list of resources supported by Config Connector.
- **version**: Print the version of the CLI.
- **licenses**: Print licenses for third-party dependencies.

## Usage

See the documentation for each command for specific usage details. For example, for the `preview` command, see [Previewing Reconciliation](../../pkg/cli/cmd/preview/README.md).
