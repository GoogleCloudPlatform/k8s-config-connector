I need to add support for a new field to the ContainerCluster KCC resource; the ContainerCluster resource uses terraform.

I want to follow the example of this PR: https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/5815

Please download the diff as needed (which you can do by looking at https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/5815.diff)

That PR added support for this field:

terraform field name: enable_cilium_clusterwide_network_policy
terraform resource name: google_container_cluster
KCC resource name: ContainerCluster
KCC field path: spec.networkConfig.enableCiliumClusterwideNetworkPolicy

In this case, I want to add support for this field;

terraform field name: enable_nested_virtualization
terraform resource name: google_container_cluster, google_container_node_pool
KCC resource name: ContainerCluster, ContainerNodePool
KCC field path: spec.nodeConfig.advancedMachineFeatures.enableNestedVirtualization

Please get the correct terraform implementation from the latest terraform implementation at https://github.com/hashicorp/terraform-provider-google/blob/main/google/services/container/node_config.go

As in the previous PR, please only add support for the new field to third_party/github.com/hashicorp/terraform-provider-google-beta/google-beta/services/container/node_config.go .  Do not add support for other missing fields.

Please also extend the test `pkg/test/resourcefixture/testdata/basic/container/v1beta1/containercluster/containercluster` to include setting the new field (in create.yaml) and updating the new field (in update.yaml)


Here is the plan I propose you follow:

* Extract the needed changes to support the new field enable_nested_virtualization from https://github.com/hashicorp/terraform-provider-google/blob/main/google/services/container/node_config.go

* Apply those changes to our code, in third_party/github.com/hashicorp/terraform-provider-google-beta/google-beta/services/container/node_config.go .  Commit the changes with a message like "terraform: add support for enable_nested_virtualization to google_container_cluster and google_container_node_pool"

* Update mockgcp to include support for updating the relevant proto field, in mockgcp/mockcontainer/cluster.go and mockgcp/mockcontainer/nodepool.go.  Commit the changes with a message like "mockgcp: add support for enableNestedVirtualization to container cluster / nodePool"

* Update the terraform service mapping if needed.  Commit the changes with a message like "feature: add support for enableNestedVirtualization to ContainerCluster and ContainerNodePool"

* Run `make ready-pr` to update the CRDs etc.  Commit the changes with a message like "autogen: generation for enableNestedVirtualization in ContainerCluster and ContainerNodePool"

* Update the tests to verify setting and updating enableNestedVirtualization.  Commit with a message like "tests: verify enableNestedVirtualization in ContainerCluster and ContainerNodePool"

* Run `hack/compare-mock fixtures/container` to update the golden output.  Commit with a message like "autogen: golden output for enableNestedVirtualization support"