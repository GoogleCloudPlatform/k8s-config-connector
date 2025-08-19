# Directory: pkg/controllers

This directory is the heart of Config Connector. It contains the implementation of the resource controllers that manage the lifecycle of GCP resources.

## Structure

The controllers are organized by GCP service. Each controller is responsible for a single KCC resource type.

A controller's main job is to watch for changes to its corresponding KCC resource and then create, update, or delete the underlying GCP resource to match the desired state.

This directory contains the "direct" controllers, which make calls to the Google Cloud SDKs. Other controllers are based on Terraform (`pkg/tf`) or DCL (`pkg/dcl`).

When you need to understand the logic for how a particular GCP resource is managed, you will need to look at the corresponding controller in this directory.

See also the root `GEMINI.md` and `pkg/GEMINI.md`.
