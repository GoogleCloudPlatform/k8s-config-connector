# Guide for Developing KCC Resources 

We are thrilled to introduce the guide for adding Config Connector resources and fields! Resources built using this guide are called "KCC Direct Controllers", because they are built directly against the service API. Developing a new resource (or a new field) is much faster and more manageable using this new Direct Controller approach than previous approaches. We have changed some key resource reconciliation processes to be more reliable and Kubernetes-native. We have also made a revolutionary change to the test driven development and PR review process to improve test coverage for every field in a resource. 

While there is more work to be done to further improve the process of adding KCC resources, we believe the guide is now ready to be shared broadly, such that Google developers, partners, and customers can add KCC resources. We will continue improving this guide to make it simpler and easier to develop new KCC resources. Please stay tuned for the upcoming changes.

# Contents

## Key steps and Exit Criteria 

* [1. Add MockGCP tests](./guides/1-add-mockgcp-tests.md)
* [2. Define API](./guides/2-define-apis.md)
* [3. Add KRM and API Mapper](./guides/3-add-mapper.md)
* [4. Add Controller](./guides/4-add-controller.md)
* [5. Releases](./guides/5-releases.md)

## Develop by scenraios

* [Add a new resource](./scenarios/new-resource.md)
* [Add a new field](./scenarios/new-field.md)
* [Promote a Alpha Resource to Beta](./scenarios/alpha-to-beta.md)
* [Migrate a TF/DCL-based Resource to Direct (Alpha)](./scenarios/migrate-tf-resource-alpha.md)
* [Migrate a TF/DCL-based Resource to Direct (Beta)](./scenarios/migrate-tf-resource-beta.md)
