# ContainerCluster remove-default-node-pool Reconciliation Behavior

## Overview
This document summarizes the architectural discussion and proposed redesign around how Config Connector (KCC) handles the `ContainerCluster` resource when the `cnrm.cloud.google.com/remove-default-node-pool: "true"` annotation is applied. 

Specifically, it addresses a recent fix ([PR 6906](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/6906)) that introduced aggressive, silent mutation of user intents to bypass strict Terraform provider validation rules, and proposes a shift back to a declarative, fail-fast model with an explicit workaround.

## The Core Issues

When provisioning GKE clusters, users often want to delete the default node pool and use dedicated pools. This is done via the `remove-default-node-pool: "true"` annotation. However, this intersects with two complex behaviors:

### 1. KCC Self-Sabotage (Drift Prevention)
Even if a user correctly omits `nodeVersion` from their YAML manifest, KCC's internal "managed fields" logic reads the actual cluster state from GCP. It sees the running master's `nodeVersion` and automatically injects it into the desired state (`config`) before calling Terraform to prevent drift.

*TODO*: Verify the conflict. I don't recall seeing such an issue while it sounds like it can happen all the time.

**The Conflict:** Terraform's GKE provider has strict validation rules: if `remove_default_node_pool = true`, specifying `node_version` is completely forbidden and throws a hard validation error during the `Plan` (Diff) phase. Because KCC auto-injected the field, KCC essentially deadlocked itself.

### 2. The Private Cluster Bootstrap Edge Case (GitOps)
When creating a highly secured private cluster, GCP must still temporarily spin up the default node pool to bootstrap the control plane. In strict VPCs, these temporary nodes *must* have specific Network Tags to communicate with the master. 

**The Conflict:** 
* To add network tags, the user *must* specify `nodeConfig` in their initial KRM manifest. Terraform allows this during the initial **Create**.
* However, on subsequent **Updates** (e.g., KCC's regular reconciliation an hour later), Terraform strictly forbids `nodeConfig` if the default pool has been removed. 
* For GitOps users (e.g., ArgoCD) with static YAML files, their manifest permanently contains `nodeConfig`. This causes perpetual validation errors on every sync after creation.

## The Recent "Norm" Fix
A recent commit attempted to solve both problems by acting as a "smart" middleman. 

During the KRM-to-Terraform translation phase, the fix:
1. Kept `nodeConfig` during initial creation.
2. **Unconditionally stripped** `nodeVersion` and `nodeConfig` from the configuration on all subsequent updates if `remove-default-node-pool: "true"` was set.

## Why the "Norm" is Flawed
While the recent fix stops the deadlock, it establishes a dangerous architectural "Norm": **KCC silently correcting and guessing user intent.**

In a declarative system, ambiguity should fail fast. If a user explicitly specifies `nodeVersion: 1.25` or `nodeConfig` alongside `remove-default-node-pool: "true"`, KCC cannot know if:
1. They made a simple copy-paste mistake in their manifest.
2. They are intentionally trying to use a static GitOps workflow for a private cluster.

By unconditionally deleting the user's explicit fields before Terraform sees them, KCC masks the underlying validation errors and hides the user's mistakes. The system effectively says, "I know you asked for this contradictory state, but I will silently change your request to make it work."

## The Proposed "Workaround" Architecture
To realign with Kubernetes declarative principles ("What you see in the YAML is what you get, or it fails"), the behavior should be redesigned from a hidden "Norm" to an explicit "Workaround."

### 1. Stop Silently Stripping User Input
Revert the logic that unconditionally deletes user-provided `nodeVersion` and `nodeConfig`. If a user explicitly provides these fields in an invalid context, pass them to Terraform. Terraform will throw a validation error, forcing the user to correct their manifest.

### 2. Fix KCC Self-Sabotage (Only Strip Unset Fields)

*TODO*: Again, verify if it is needed.

KCC must only strip fields that *it injected itself*. Before calculating the diff, KCC should check if `nodeVersion` or `nodeConfig` were originally present in the user's applied spec. If they were *not* set by the user (meaning KCC injected them during the state-overlay phase), KCC should strip them out. This prevents the auto-injection deadlock without touching user intent.

### 3. Explicit Opt-In Workaround for GitOps
To support the GitOps private cluster edge case without breaking the declarative model, introduce an explicit opt-in annotation. 

For example: `cnrm.cloud.google.com/strip-default-node-pool-config-on-update: "true"`

If—and only if—the user applies this annotation, KCC will implement the workaround: "I will pass `nodeConfig` during initial creation, but I will actively strip it from the desired state on subsequent updates."

This forces the user to declare their intent: *"I know this YAML is contradictory for updates, but I require it for bootstrapping. Please apply the workaround."* It eliminates ambiguity and keeps KCC from guessing.