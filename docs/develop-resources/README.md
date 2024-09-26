# Developing the Direct Resource Guide 

We are thrilled to introduce the new way to add a Config Connector resource: the direct resource guide!  

In the past few months, we have made a tremendous amount of efforts to make adding a new resource (or a new field) much faster and more manageable, we have changed some key resource reconciliation processs to be more reliable and Kubernetes-native. What's more, we have made a revolutionary change to the test driven development and PR review process to improve the test coverage for every single field. 

There are definitely more work need to be done, but we'd like to share the steps about adding the direct resource so that it can benefit more users. We will conitnue improving this guide and making the step changes to make it simpler and easier to use. Please stay tuned for the upcoming changes.

# Contents

## Introduction

* [Introduction](./guides/0-introduction.md)

## Key steps and Exit Criteria 

* [1. Add MockGCP tests](./guides/1-add-mockgcp-tests.md)
* [2. Define API](./guides/2-define-apis.md)
* [3. Add KRM and API Mapper](./guides/3-add-mapper.md)
* [4. Add Controller](./guides/4-add-controller.md)
* [5. Releases](./guides/5-releases.md)

## Develop by scenraios

You may only need to add a new field, or to promote an Alpha resource to Beta rather than writing an entirely new resource. To figure out the most suitable scenario, you need to know the resources current status:
- [ ] Check the latest [CRD](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/master/crds). If the resource exists but in Alpha, you won’t see it in  [Config Connector API Reference](https://cloud.google.com/config-connector/docs/reference/overview) until it is promoted to Beta. Please follow [promote alpha to beta](./scenarios/alpha-to-beta.md) guide.
- [ ] If no CRD is found,  you do need to add a new resource, please file an issue first to let us know. We can help make sure the resource is not already taken care of by someone else.
- [ ] We are migrating from the TF/DCL based resources to the Direct approach. That means we are holding off new PR reviews if they are using the TF/DCL based approach. Please let us know if you encounter any problems or any specific reasons that can only use the TF/DCL based approach.


* [Add a new resource](./scenarios/new-resource.md)
* [Add a new field](./scenarios/new-field.md)
* [Promote a Alpha Resource to Beta](./scenarios/alpha-to-beta.md)
* [Migrate a TF/DCL-based Resource to Direct (Alpha)](./scenarios/migrate-tf-resource-alpha.md)
* [Migrate a TF/DCL-based Resource to Direct (Beta)](./scenarios/migrate-tf-resource-beta.md)
