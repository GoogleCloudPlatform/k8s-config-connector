# Directory: pkg/dcl

This directory contains the Config Connector controllers that are based on the Declarative Client Library (DCL).

DCL is a library that provides a declarative interface for managing GCP resources. The controllers in this directory use DCL to create, update, and delete GCP resources.

The goal is to eventually migrate all controllers to the "direct" approach in `pkg/controllers`, but many DCL-based controllers are still in use.

When working with a resource that is managed by a DCL-based controller, you will need to understand how DCL works to understand the controller's logic.

See also the root `GEMINI.md` and `pkg/GEMINI.md`.
