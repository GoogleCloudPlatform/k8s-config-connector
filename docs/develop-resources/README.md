# Developing the Direct Resource Guide 

We are thrilled to introduce the new way to add a ConfigConnector resource: the direct resource guide!  

In the past few months, we have made a tremendous amount of efforts to make adding a new resource (or a new field) much faster and more manageable, we have changed some key resource reconciliation processs to be more reliable and Kubernetes-native. What's more, we have made a revolutionary change to the test driven development and PR review process to improve the test coverage for every single field. 

There are definitely more work need to be done, but we'd like to share the steps about adding the direct resource so that it can benefit more users. We will conitnue improving this guide and making the step changes to make it simpler and easier to use. Please stay tuned for the upcoming changes.   

# Contents

## Key steps and Exit Criteria 

* [1. Add MockGCP tests](./guides/1-add-mockgcp-tests.md)
* [2. Define API](./guides/2-define-apis.md)
* [3. Add KRM and API Mapper](./guides/3-add-mapper.md)
* [4. Add Controller](./guides/4-add-controller.md)
* [5. Releases](./guides/5-releases.md)

## Develop by scenraios

*Find out the resource's current status*
- If the resource is not in [ConfigConnector API Reference](https://cloud.google.com/config-connector/docs/reference/overview), it could be an Alpha resource. Check the list 


* Check if the 
Please follow this [guide](./scenarios/1-add-new-direct-resource.md) to add a Direct resource instead.
If there are any specific reasons you cannot use the Direct resource, please let us know by filing an issue.
* [Add a new resource](./scenarios/new-resource.md)
* [Add a new field](./scenarios/new-field.md)
* [Promote a Alpha Resource to Beta](./scenarios/alpha-to-beta.md)
* [Migrate a TF/DCL-based Resource to Direct (Alpha)](./scenarios/migrate-tf-resource-alpha.md)
* [Migrate a TF/DCL-based Resource to Direct (Beta)](./scenarios/migrate-tf-resource-beta.md)