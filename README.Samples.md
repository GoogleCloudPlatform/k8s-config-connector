# Overview

The following guidelines and principles will help you create samples that are
useful for Config Connector users. Good samples are important to the product and
should be taken as seriously as good code. By creating good samples we enable
our users to be successful more quickly.

# Sample Coverage

Every resource should have at least one basic sample and additional samples as
necessary.

## Basic Sample

The first sample should show as many fields as possible to enable users to see
all their options and delete where necessary. For this reason, it should contain
all fields which are not mutually exclusive into a subfolder.

## Other Samples

When a resource has multiple important use cases, then there should be multiple
samples. A key indicator of multiple use cases is mutually exclusive fields or
exclusive values for those fields. For example, with the `SQLInstance` resource,
the `databaseVersion` property selects different database vendors in MySQL or
PostgreSQL. However, in addition to changing the version property, one must also
select a specific `tier` value that is compatible with the version. SQL is such
an important use case to our users that it is worth having samples for both.

Another SQL example is replication. Replication is a deeply nested structure
that is complicated, and it is an important configuration for enterprise users.
However, replication requires two SQLInstances so it is not a great candidate to
be the only sample. For that reason a specific sample for replication with two
databases in it should be created.

# Directory Organization

In the [resources](config/samples/resources) folder each Resource should get its
own folder. The name of the folder should be its CRD Kind in all lowercase.

If the resource has a single sample, then it should be added to the folder
directly (without any sub-folder). If there are multiple samples, then each
sample should be in a sub-folder. The sub-folder name should be all lowercase
using dashes instead of spaces, and should be a grammatical description of the
sample type, with the literal resource name included. For example, for the
`SQLInstance` sample with two replicated MySQL instances, the folder name is
`mysql-sql-instance-with-replication`. The service name can be dropped from
folder names for readability, as in the `ComputeForwardingRule` sample folders
`global-forwarding-rule` and `regional-forwarding-rule`.

## File Naming

Files should be named `<service>_<api-version>_<kind>.yaml`. For example,
`pubsub_v1alpha2_pubsubtopic.yaml`.

## Sample Naming Convention

The `metadata.name` value of the sample resource should be in the format `<full
resource name>-sample` (all lowercase). If the name of the underlying GCP
resource name cannot be in this format, specify the GCP resource name using
`spec.resourceID` instead of `metadata.name` (e.g. see
[Service sample](config/samples/resources/service/serviceusage_v1beta1_service.yaml)).

In cases with multiple samples, the sample name should replicate the folder
name, but with the resource name moved to the front, separated from the
description by `-sample-` and with all other dashes removed. Even if the service
name wasn't present in the folder name, it should be included at the front of
the sample name. For example, the sample in
`mysql-sql-instance-with-replication` is named
`sqlinstance-sample-mysqlwithreplication`. Also, the sample in
`global-forwarding-rule` is named `computeforwardingrule-sample-global`.

In the case where multiple instances of the main resource are needed for a
complete sample, number the sample block starting from one, for example,
`computenetworkpeering-sample1` or `sqlinstance-sample1-mysqlwithreplication`
are followed by `computenetworkpeering-sample2` or
`sqlinstance-sample2-mysqlwithreplication` respectively.

# Dependency Resources

When a sample has a dependency resource then the dependency should be included
in the same directory. The dependency resource's declaration should be as brief
as possible and only the most necessary fields should be included. For example,
labels should be removed, fields with default values, etc. Multiple dependencies
of the same type should be defined in the same file, split by a YAML separation
line (`---`).

An example would be `PubSubSubscription` which depends on `PubSubTopic`. In this
case, a YAML file containing a `PubSubTopic` resource should be included in the
[pubsubsubscription](config/samples/resources/pubsubsubscription) folder.

## Dependency Naming Convention

The name of a dependency resource should be the name of the resource that
depends on it followed by `-dep`. For example, `PubSubSubscription` depends on
`PubSubTopic`. The Topic's name should be `pubsubsubscription-dep`.
Additionally, if there is a chain of dependencies, the name of each dependency
resource should be the name of the top-level resource that depends on it
followed by `-dep`. For example, `FilestoreBackup` depends on
`FilestoreInstance`, which in turn depends on `ComputeNetwork`. Both
`FilestoreInstance` and `ComputeNetwork` should be named `filestorebackup-dep`.
If there are multiple dependencies of the same type, then number the dep suffix,
starting from one: `-dep1`, `-dep2`, `-dep3`, etc.

If there are multiple samples for a given resource then the sample name should
be included as well to ensure uniqueness. For example, if `PubSubSubscription`
had two samples, named `pubsubsubscription-sample-first` and
`pubsubsubscription-sample-second` then the name of the dependency topics would
be`pubsubsubscription-dep-first` and `pubsubsubscription-dep-second`
respectively.

# Field Values

Field values should be generic. String values should not include project or
product names. For example, values like KCC, CNRM, etc,should be avoided.

## Project References

Resources that support a `spec.projectRef` field should always include that
field in the sample, e.g.:

```yaml
spec:
  projectRef:
    external: projects/${PROJECT_ID?}
```

**Note:** If the resource can be created in more than one type of hierarchical
resource (e.g. project, folder, organization), then you should have samples for
each case, and each sample should explicitly specify the hierarchical reference
like `spec.projectRef` in order to demonstrate how to create the resource under
different types of hierarchical resources.

**Note 2:** If the sample modifies the project itself (e.g.
`ResourceManagerPolicy`, `ComputeSharedVPCHostProject`, etc.), then set
`spec.projectRef` to point to a custom `Project` resource instead.

**Note 3:** You should create a custom `Project` resource if your test's target
resource and/or dependent resources are singleton GCP resources like
**Services** and **GKEHubFeatures**.

## Descriptive Context

Fields whose purpose is to provide more documentation (e.g. description fields,
title fields, etc.) should be filled with as much detail about the sample
context as possible. When there are no other descriptive fields, the name and
label fields of the resource should be used to provide any necessary context
instead.

Examples:

```yaml
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeNodeTemplate
metadata:
  name: computenodetemplate-sample-flexible
spec:
  ...
  description: Node template for sole tenant nodes running in us-central1, with 96vCPUs and any amount of memory on any machine type.
  ...
```

```yaml
apiVersion: iam.cnrm.cloud.google.com/v1beta1
kind: IAMPolicyMember
metadata:
  name: iampolicymember-sample-condition
spec:
  ...
  condition:
    title: expires_after_2019_12_31
    description: Expires at midnight of 2019-12-31
    expression: request.time < timestamp("2020-01-01T00:00:00Z")
  ...
```

# License

All sample files should have a license header. Copy the license header below
onto the top of your sample files and replace `[YEAR]` with the current year.

```
# Copyright [YEAR] Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
```
