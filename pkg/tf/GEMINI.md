# Directory: pkg/tf

This directory contains the Config Connector controllers that are based on the Terraform provider for Google.

Many of the older KCC controllers are implemented by wrapping the Terraform provider. This allowed KCC to support a large number of GCP resources quickly.

The long-term goal is to migrate these controllers to the "direct" approach in `pkg/controllers`, but many Terraform-based controllers are still in use.

When working with a resource that is managed by a Terraform-based controller, you will need to understand how Terraform works to understand the controller's logic.

See also the root `GEMINI.md` and `pkg/GEMINI.md`.
