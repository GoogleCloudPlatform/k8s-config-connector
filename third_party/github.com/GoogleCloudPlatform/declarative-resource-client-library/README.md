# Declarative Client Library

*This is not an officially supported Google product.* If you are looking for
a way to provision Google Cloud resources programmatically, please use the
offical [Google Cloud Client Libraries](https://cloud.google.com/apis/docs/cloud-client-libraries).

## About the DCL

The Declarative Client Library (DCL) is a Go library that provides a
declarative configuration interface for Google Cloud resources on top of
their existing imperative APIs.

Each resource type has four primary methods:

- *Get*: Takes a resource's identity field (such as region and name) and
returns the declarative interface of the resource as it currently exists.
- *List*: List all of the resources of a given type and returns their
declarative interfaces.
- *Apply*: Make GCP's state match the given resource, creating the resource
if it doesn't exists, updating the resource to the desired state if necessary,
or doing nothing if the resource exists and is in the desired state.
- *Delete*: Delete the given resource.

The DCL exists to support open source provisioning tools like the
[Terraform Google Cloud Platform Provider](https://www.terraform.io/docs/providers/google/index.html).
For bug reports and feature requests, please visit these tools' GitHub
repos directly.