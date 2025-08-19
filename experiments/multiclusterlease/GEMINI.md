# Directory: experiments/multiclusterlease

This directory contains an experimental implementation of a multi-cluster leader election mechanism for Kubernetes.

## Overview

The `MultiClusterLease` CRD and controller provide a way for a single replica of a controller to be elected as a leader from a pool of candidates running across multiple Kubernetes clusters.

This is a powerful feature for building highly available and resilient control planes.

## Key Components

*   `MultiClusterLease` CRD: The API contract for declaring candidacy and observing the election outcome.
*   Election Controller: A decentralized controller that runs in each cluster and contends for a global lock.
*   Client-Side Library: A Go library that implements the standard `resourcelock.Interface` for easy integration with `controller-runtime`.

When you need to understand how KCC can be run in a multi-cluster configuration, this is the directory to study.

See also the root `GEMINI.md` and `experiments/GEMINI.md`.