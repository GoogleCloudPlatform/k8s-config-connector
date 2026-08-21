# Promote an Alpha Resource to Beta

You want to first figure out if the resource is Terraform-based, DCL-based or direct. Check the labels in the resource' [CRD](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/master/crds). 

- If the resource has `cnrm.cloud.google.com/tf2crd` label, it is Terraform-based
- If the resource has `cnrm.cloud.google.com/dcl2crd` label, it is DCL-based
- If the resource has neither label, it is direct resource already.

## From Terraform/DCL-based resource to Beta

For Terraform-based or DCL-based resource, you shall migrate them to [the direct approach (Alpha)](./migrate-tf-resource-alpha.md) first, and then promote to Beta.

## From direct Alpha to Beta

For direct resource, you can follow [deep-dives 5](../deep-dives/5-releases.md#52-bump-from-v1alpha1-to-v1beta1) to promote the resource.

