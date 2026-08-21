# Critical User Journeys

This page provides a guide to the most common configuration and management tasks in Config Connector.

## Configuring Operational Modes

*   **[Choose between Cluster and Namespace Mode](config-connector-mode.md)**: Understand the different ways to deploy and manage Config Connector in your cluster.
*   **[Configure Identity for a Namespace](config-connector-mode.md#identity-and-authentication)**: Set up the Google Service Account (GSA) used by Config Connector for a specific namespace.
*   **[Manage Billing and Quota](config-connector-mode.md#billing-and-quota)**: Control which Google Cloud project is billed for your resources and whose quota is consumed.

## Managing Reconciliation

*   **[Choose Controller Implementations (Direct vs. Terraform)](controller-configuration.md)**: Manually select between Direct, Terraform, or DCL controllers for specific resources.
*   **[Pause Reconciliation](pause.md)**: Temporarily stop Config Connector from managing one or more resources.
*   **[Pin to an Older Version (Backward Compatibility)](compatabilityversion.md)**: Run a specific namespace at an older version of Config Connector to ensure stability during upgrades.

## Advanced Configuration

*   **[Opt-in to Experimental Controllers](optin.md)**: Test new "Direct" controllers before they become the default.
*   **[Configure Rate Limiting](ratelimit.md)**: Adjust the rate at which Config Connector calls Google Cloud APIs.
*   **[Manage Container Resources](containerresource.md)**: Configure the memory and CPU limits for the Config Connector pods.

## Legacy and Deprecated Behaviors

*   **[Legacy Behavior Reference](legacy-behavior.md)**: Information on legacy annotations and configurations that are still supported but no longer recommended.
