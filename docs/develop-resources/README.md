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

* [1. Add a new Direct resource](./scenarios/1-new-direct.md)
* [2. Add a new TF/DCL-based resource](./scenarios/2-new-tf.md)
* [3. Add a new field to Direct resource](./scenarios/3-new-field-to-direct.md)
* [4. Add a new field to TF/DCL-based resource](./scenarios/4-new-field-to-tf.md)
* [5. Bump a TF/DCL-based Alpha Resource to Direct Beta](./scenarios/5-tf-alpha-to-direct.md)
* [6. Bump a TF/DCL-based Beta  Resource to Direct Beta](./scenarios/6-tf-beta-to-direct.md)