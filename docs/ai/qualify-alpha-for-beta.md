# How to Qualify an Alpha Resource for Beta Promotion

This document outlines the process for making an alpha resource eligible for promotion to beta by ensuring it has full API test coverage.

## Background

The file `experiments/promoter/results/candidates.json` lists alpha resources that are candidates for promotion. Some of these resources may have the `apiCoverage` flag set to `false`.

A `false` value for `apiCoverage` indicates that while the resource has all the necessary components (APIs, mapper, controller, and mock GCP implementation), it lacks a comprehensive test suite that covers all the fields in its API.

## Steps to Qualify a Resource

To qualify a resource for beta promotion, you need to add a full test suite and then update its `apiCoverage` status.

### 1. Add a Full Test Suite

Follow the detailed instructions in [add-full-test-suite.md](experiments/promoter/tasks/add-full-test-suite.md) to create a complete test suite for the resource. This guide will walk you through:

-   Creating `create.yaml` and `update.yaml` test fixtures.
-   Verifying that these fixtures cover every field in the resource's CRD.
-   Recording the GCP traffic for both creation and update operations.
-   Verifying the mock GCP implementation against the recorded traffic.

### 2. Update API Coverage Status

Once you have successfully created and verified the full test suite, the final step is to update the resource's status in the `experiments/promoter/results/candidates.json` file.

Change the value of the `apiCoverage` field from `false` to `true` for the resource you have just tested. This signals that the resource now has complete API test coverage and is ready for the promotion process.
