# EPIC: Kubernetes Config Connector (KCC) Support in Sovereign Cloud & Isolated Universes

**Status:** Draft for Review • September 2026  
**Author:** Frank Currie, Product Management  
**Target Milestone:** KCC v1.158.0 / v1.159.0  
**Affected Repositories:** `GoogleCloudPlatform/k8s-config-connector`, `hashicorp/terraform-provider-google-beta`

---

## Executive Summary

Kubernetes Config Connector (KCC) automates the lifecycle of Google Cloud infrastructure through Kubernetes-native custom resources. However, when deployed into isolated Sovereign Cloud partitions (such as the Google Sovereign Cloud Dogfood in `custom.universe.goog` or France/S3NS in `custom.universe.goog`), KCC fails systematically across both authentication and actuation layers. 

Every outbound API call from KCC reconcilers routes to commercial Google Cloud endpoints (`*.googleapis.com`). Because Sovereign Cloud workloads authenticate via regional Workload Identity pools (e.g., `<project>.<partition>.svc.id.goog`) backed by isolated Security Token Service (STS) issuers, commercial Google Cloud API gateways reject these tokens with `HTTP 401 Unauthorized (ACCESS_TOKEN_TYPE_UNSUPPORTED)`. Concurrently, KCC deploys its entire catalog of 295 Custom Resource Definitions (CRDs), despite Sovereign Cloud regions offering only a bounded subset of 33 core services. This uncurated installation imposes unnecessary memory and etcd overhead while registering controllers for APIs that cannot resolve.

This epic defines the engineering blueprint to generalize KCC for multi-universe deployment across all sovereign partitions, air-gapped environments, and custom cloud endpoints without regional hardcoding.

---

## Architectural Problem Analysis

Empirical validation on a live Sovereign Cloud cluster (`sample-cluster` in project `partition:sample-project`, region `u-region-1`, universe `custom.universe.goog`) surfaced four structural defects:

1. **Terraform Provider Shim Bypasses Universe Domain Configuration:**  
   KCC instantiates the embedded Google Terraform provider (`tfschema.Provider`) within `pkg/tf/provider/provider.go`. While the underlying Terraform provider exposes a `universe_domain` attribute, KCC's wrapper struct `Config` omits this field, and `New(ctx, config)` never passes `universe_domain` into the provider schema configuration map. Consequently, all Terraform-based reconcilers (such as `PubSubTopic`, `StorageBucket`, and `ComputeAddress`) default strictly to `googleapis.com`.

2. **Direct Reconcilers Contain Hardcoded Endpoint Strings:**  
   Direct controllers implemented under `pkg/controller/direct/*` construct gRPC/REST clients using string-formatted endpoint literals (e.g., `fmt.Sprintf("%s-assuredworkloads.googleapis.com:443", location)` or literal `"containeranalysis.googleapis.com:443"`). These bypass both `GOOGLE_CLOUD_UNIVERSE_DOMAIN` and client option injectors.

3. **Workload Identity & Service Account Identity Mapping Failures:**  
   In sovereign partitions, project IDs are qualified with partition prefixes (e.g., `partition:sample-project`). Google IAM formats the resulting Service Account emails with dot notation rather than colon notation (`sa@sample-project.partition.iam.gserviceaccount.com`). KCC's internal IAM parser reconstructs emails assuming standard commercial syntax (`sa@partition:sample-project.iam.gserviceaccount.com`), which contains an illegal RFC 5322 domain character (`:`) and triggers immediate API rejection.

4. **Static Monolithic CRD Deployment vs. Bounded Regional Capabilities:**  
   KCC operator unconditionally registers 295 CRDs during bootstrap. Telemetry from `serviceusage.googleapis.com` in `u-region-1` confirms that only 33 APIs are available in this sovereign boundary. Deploying 262 unsupported CRDs causes continuous reconciliation retry loops, metric export failures, and etcd bloat.

---

## Target Architecture

```
                               +--------------------------------------------+
                               |         ConfigConnector Controller         |
                               +--------------------------------------------+
                                                     |
                 +-----------------------------------+-----------------------------------+
                 |                                                                       |
                 v                                                                       v
   +---------------------------+                                           +---------------------------+
   |   Terraform Reconciler    |                                           |     Direct Reconciler     |
   | (pkg/tf/provider/provider)|                                           |   (pkg/controller/direct) |
   +---------------------------+                                           +---------------------------+
                 |                                                                       |
                 | Reads GOOGLE_CLOUD_UNIVERSE_DOMAIN                                    | Resolves Endpoint via
                 | Sets cfgMap["universe_domain"]                                        | gcp.GetEndpoint(svc, universe)
                 v                                                                       v
   +---------------------------------------------------------------------------------------------------+
   |                                Sovereign API Gateway Enpoint                                      |
   |                          https://<service>.<UNIVERSE_DOMAIN>/v1/...                               |
   |                                 (e.g., custom.universe.goog)                                   |
   +---------------------------------------------------------------------------------------------------+
```

---

## Child Issues & Delivery Plan

| Issue ID | Title | Component | Target Scope |
| :--- | :--- | :--- | :--- |
| **#1** | Propagate `universe_domain` to Terraform Provider Shim & DCL Client | `pkg/tf/provider`, `pkg/controller/kccmanager` | Universal universe domain propagation |
| **#2** | Replace Hardcoded `googleapis.com` Literals in Direct Controllers | `pkg/controller/direct/*`, `pkg/gcp` | Endpoint builder utility for gRPC/REST |
| **#3** | Resolve Partition-Qualified Project IDs in IAM Identity Parsing | `pkg/gcp/project.go`, `pkg/controller/iam` | Partitioned service account email formatting |
| **#4** | Dynamic Service Usage Discovery & Regional CRD Gating | `operator/pkg/discovery`, `operator/config` | Selective CRD registration based on enabled APIs |

---

## Non-Goals & Costly Refusals

1. **No Regional Hardcoding:** Under no circumstances will Sovereign Cloud Germany domain names (`custom.universe.goog`), project prefixes (`eu0:`), or regional identities be hardcoded into KCC. All logic must consume `GOOGLE_CLOUD_UNIVERSE_DOMAIN` or the `universeDomain` spec field dynamically.
2. **No Fallback to Unauthenticated Access:** If STS tokens fail to exchange or if an endpoint cannot be verified against the declared universe domain, KCC must fail closed and report a descriptive `UpdateFailed` condition rather than retrying commercial endpoints.
3. **No Breaking Changes for Public Cloud Users:** Default behavior in the absence of `GOOGLE_CLOUD_UNIVERSE_DOMAIN` or `universeDomain` must remain 100% backward-compatible with `googleapis.com`.
