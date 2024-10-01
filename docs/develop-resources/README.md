# Developing the Direct Resource Guide 

We are thrilled to introduce the new way to add a Config Connector resource: the direct resource guide!  

In the past few months, we have made a tremendous amount of efforts to make adding a new resource (or a new field) much faster and more manageable, we have changed some key resource reconciliation processs to be more reliable and Kubernetes-native. What's more, we have made a revolutionary change to the test driven development and PR review process to improve the test coverage for every single field. 

There are definitely more work need to be done, but we'd like to share the steps about adding the direct resource so that it can benefit more users. We will conitnue improving this guide and making the step changes to make it simpler and easier to use. Please stay tuned for the upcoming changes.

Be sure to work through the [Contributing guide](../../CONTRIBUTING.md) before starting.

# Contents

## Introduction

* [Introduction](./guides/0-introduction.md)

## Key steps and Exit Criteria 

* [1. Add MockGCP tests](./guides/1-add-mockgcp-tests.md)
* [2. Define API](./guides/2-define-apis.md)
* [3. Add KRM and API Mapper](./guides/3-add-mapper.md)
* [4. Add Controller](./guides/4-add-controller.md)
* [5. Releases](./guides/5-releases.md)

## Develop by scenarios

To determine the best approach for adding support for the resource or field, please check the resource's current status:
**Check the CRD:** Examine the latest [CRDs](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/master/crds). The version can be found from the file name. If the resource exists but is in Alpha (and therefore not yet in the [Config Connector API Reference](https://cloud.google.com/config-connector/docs/reference/overview)), follow the [Alpha to Beta promotion guide](./scenarios/alpha-to-beta.md).  This may be the simplest solution.
**File an issue (if no CRD exists):** If no CRD is found, you'll likely need to create a new resource. Before starting this process, check if an issue is already filed. If not, please file an issue to avoid duplication of effort and allow us to coordinate effectively. If you are willing to take the resource yourself, please assign the issue. 
**Prioritize the Direct approach:** We are currently prioritizing the Direct approach over TF/DCL-based resource implementations.  If you encounter issues preventing the use of the Direct approach, please let us know in the filed issue so we can assist.

* [Add a new resource](./scenarios/new-resource.md)
* [Add a new field](./scenarios/new-field.md)
* [Promote a Alpha Resource to Beta](./scenarios/alpha-to-beta.md)
* [Migrate a TF/DCL-based Resource to Direct (Alpha)](./scenarios/migrate-tf-resource-alpha.md)
* [Migrate a TF/DCL-based Resource to Direct (Beta)](./scenarios/migrate-tf-resource-beta.md)
