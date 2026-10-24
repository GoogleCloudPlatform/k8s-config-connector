# AppEngineFirewallRule Migration Journal

## Overview
Migrated `AppEngineFirewallRule` from the legacy Terraform-based reconciler to the `direct` reconciler architecture.

## Key Observations and Decisions
1. **GAPIC REST Client Parent Formatting for Create:**
   - In GAPIC's `firewallRESTClient.CreateIngressRule`, the endpoint URL template is `/v1/{parent=apps/*}/firewall/ingressRules`. GAPIC internally appends `/firewall/ingressRules` to `req.GetParent()`.
   - As a result, `Parent` in `CreateIngressRuleRequest` must be set to `apps/{app}`, whereas `GetIngressRule`, `UpdateIngressRule`, and `DeleteIngressRule` expect the full resource path `apps/{app}/firewall/ingressRules/{priority}` as `Name`.

2. **Enum Serialization and GAPIC URL Encoding:**
   - The GAPIC REST client adds `%24alt=json%3Benum-encoding%3Dint` to query parameters and marshals enums as integers (`ALLOW` = 1, `DENY` = 2) in request JSON bodies, whereas the legacy Terraform provider serialized string enum values with `alt=json`.
   - MockGCP supports both representations seamlessly.

3. **Brownfield Dual-Testing Alignment:**
   - `pkg/controller/resourceconfig/static_config.go` was updated to add `k8s.ReconcilerTypeDirect` to `SupportedControllers` for `AppEngineFirewallRule` while maintaining `DefaultController: k8s.ReconcilerTypeTerraform`.
   - Golden comparison files (`_http_mock.diff`, `_final_object.diff`, `_http_old_controller_mock.log`, `_final_object_old_controller.golden.yaml`) capture the exact behavioral comparison between the direct and legacy controllers.
   - All tests pass hermetically against MockGCP.
