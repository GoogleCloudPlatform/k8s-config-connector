# AppEngineFirewallRule MockGCP Alignment Journal

- **Date:** October 22, 2026
- **Resource Kind:** `AppEngineFirewallRule`
- **GCP Service:** `appengine.googleapis.com` (`google.appengine.v1.Firewall`)

## Overview

No mock previously existed for `appengine.googleapis.com` or `AppEngineFirewallRule` in `mockgcp/`. We implemented the `mockappengine` service using `httptogrpc` backed by `cloud.google.com/go/appengine/apiv1/appenginepb`.

## Real GCP Recording & Findings

1. **App Engine Application Dependency**:
   - In real GCP, App Engine firewall rules live under an App Engine application: `apps/{app_id}/firewall/ingressRules/{priority}`.
   - The GCP project must have an App Engine application created (`apps/{project}`).
2. **Default Rule Behavior**:
   - In real GCP, every App Engine application is initialized with a default rule:
     - Priority: `2147483647` (`math.MaxInt32`)
     - Action: `ALLOW`
     - Source Range: `*`
     - Description: `The default action.`
   - Deleting the default rule returns an `INVALID_ARGUMENT` error (`Cannot delete the default rule`).
3. **Specifying Project in Test Fixtures**:
   - In KCC's service mapping for `google_app_engine_firewall_rule`, `idTemplate: "apps/{{project}}/firewall/ingressRules/{{priority}}"`.
   - Because the service mapping does not specify container bindings for project, `spec.project: ${projectId}` must be explicitly configured in KRM manifests (`create.yaml` and `update.yaml`).
4. **Proto Route Parameter Binding**:
   - In `google/appengine/v1/appengine.proto`, `CreateIngressRule` is mapped as `/v1/{parent=apps/*}/firewall/ingressRules`. Consequently, the bound `parent` field in `CreateIngressRuleRequest` is `apps/{app_id}` rather than the full collection path `apps/{app_id}/firewall/ingressRules`. The mock parser was designed to support both formats.

## Implementation Details

- **Service Registration**: Registered `mockappengine` in `mockgcp/register.go` and added `appengine.googleapis.com` to `mockgcp/mockserviceusage/knownservices.go`.
- **Firewall Server**: Implemented `GetIngressRule`, `CreateIngressRule`, `UpdateIngressRule`, `DeleteIngressRule`, `ListIngressRules`, and `BatchUpdateIngressRules` in `mockgcp/mockappengine/firewallrules.go`.
- **Harness & Ratcheting**: Removed `AppEngineFirewallRule` from ratcheting exclusion in `tests/e2e/ratcheting.go` and enabled the kind in `config/tests/samples/create/harness.go`.

## Validation

- Recorded authentic real GCP golden logs (`_http.log`) and golden objects with `hack/record-gcp "fixtures/^appenginefirewallrulebasic$"`.
- Ran `hack/compare-mock "fixtures/^appenginefirewallrulebasic$"` to generate `_http_mock.log` and verify full mock reconciliation.
- Executed `TestGoldenLogAlignment` and `TestRealHTTPLogsDoNotContainMockGCP`, passing with 0 errors and 0 diffs.
