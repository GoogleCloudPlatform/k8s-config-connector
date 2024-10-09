# Direct Resource Reference Guide


# TL;DR

Referencing GCP objects should define a API field `spec.<Kind>Ref(s)` to reference the dependency of another GCP resource.  


# Basic rules

If a Config Connector resource depends on another resource, it should define a `spec.<Kind>Ref` field, even if the dependent does not have a Config Connector resource yet (in which case only External is required).

If a Config Connector resource needs to depend on a list of resources of a single Kind, it should define a `spec.<Kind>Refs `field.


```yaml
spec:
  # A single dependency to network 
  computeNetworkRef: 
    name: sample-network  

  # A list of dependencies of the same kind 
  projectRefs:
    name: Config Connector-project1
    name: Config Connector-project2
    name: Config Connector-project3 
```

# Naming

As a naming convention, the resource reference field should be `<Kind>Ref(s)`.

* `Kind` is the dependency’s Kind with a lowercase first letter, i.e. `projectRef`
* `Ref` can be singular or plural, depending on the number of dependencies, i.e. `projectRef` refers to a single project, `projectRefs` refers to a list of projects


# API Rule


```yaml
type <KindRef> struct {
  // +optional   
  External string   `json:"external,omitempty"`
  // +optional  
  Name string       `json:"name,omitempty"`
  // +optional
  Namespace string  `json:"name,omitempty"`
}
```

# Validation


## Rule 1: Config Connector level only

The reference validation can be either CRD validation or Config Connector controller check. It should not require GCP calls.


## Rule 2: Required fields

* If the reference does not have a corresponding Config Connector Kind yet, the `.<Kind>Ref.external` is required. Note: the `.<Kind>Ref` itself can be optional.
* If the reference has a corresponding Config Connector Kind, the `<Kind>Ref.external` and `<Kind>Ref.name` are `oneOf` required. Note: the `.<Kind>Ref` itself can be optional.

## Rule 3:  External

The `external` should be in the format of the asset inventory without the service domain.


* i.e. `computeNetworkRef.external` should be in the form of `projects/<projectID>/global/networks/<networkID>`


## Rule 4:  Namespace

If the referenced Config Connector object is cluster scoped or in the same namespace, the referenced `<Kind>Ref.namespace` is optional and should use `default`

If the referenced Config Connector object is namespace scoped but not in the same namespace, the referenced `<Kind>Ref.namespace` is **required **to avoid Config Connector ambiguity and customer errors.


## Rule 5:  Errors

Config Connector has a predefined `k8s.ReferenceNotFound` error that should be used when the referenced Config Connector object is not found.


# Same Kind references

A list of references of the same kind can introduce many problems if not handled well. 

Config Connector should have strict validations for those resources _according to the real usage_. This will make Config Connector survive in the long run to avoid backward compatibility overhead.  


## Rule 1: Avoid mixed form

If the GCP service expects each reference to be unique, Config Connector should require using either `<Kind>Refs[].external` or  `<Kind>Refs[].name`, but **not** a mix of the two types. This gives sanity uniqueness checks without too much user overhead.


```yaml
spec:
  projectRefs:
    name: Config Connector-project1
    name: Config Connector-project2
    name: Config Connector-project3 
```


Or


```yaml
spec:
  projectRefs:
    external: projects/gcp1
    external: projects/gcp2
    external: projects/gcp3
```


## Rule 2: Form switch allowed

Config Connector shall allow users to change between  `<Kind>Refs[].external `and  `<Kind>Refs[].name`


## Rule 3: Unique `external`

If `<Kind>Refs[].external` is used, Config Connector shall only validate the uniqueness of the string values, but not check any GCP level requirements. 


## Rule 4: Non-unique `name`

Config Connector does not (yet) have a good handle on the uniqueness of the Config Connector objects and their corresponding GCP resources. 

If `<Kind>Refs[].name` is used, Config Connector shall **not** validate the uniqueness of the namespace/name value, but check the uniqueness of the GCP resources from the corresponding Config Connector objects from [externalRef field](external-reference.md).


## Rule 5: Ordering

Config Connector shall send the **exact same order** of `<Kind>Refs[] `to GCP service, unless sorting is preferred by the GCP service.

A change to the `<Kind>Refs[] `order shall **not** trigger a new GCP call if only the order is changed but not the real content, unless the order matters to GCP services. If Config Connector cannot make the decision, skip this check. (open to discuss, I see some real use cases here) 


# Code Style

The code of adding a reference should be placed in `<kind>_reference.go` file under  `apis/<service>/<version>/`.


# Backward Compatibility

For TF-based or DCL-based Beta resources, Config Connector shall keep their original CRD and behavior when migrating to the Direct Resource.


## Generic Reference with Kind

Some Config Connector resources support a generic reference that requires `Kind`. 

When migrating to the Direct Resource, we should treat the `resourceRef.Kind` as `<Kind>Ref`, all other rules apply.


```yaml
spec:
   resourceRef:
      kind: Project
      name: gcp1
```

Some other thoughts: a conversion to change `resourceRef` to `<Kind>Ref`; warn that `spec.resourceRef` as deprecated in `status`


## Ambiguous `external` usage 

Some Config Connector resources make the `resourceKind.external` to serve different usages.

For example, the `computeForwardingRule` has the `ComputeAddress` reference which `external` allows IP address value like `8.8.8.8 `(search `spec.ipAddress.addressRef.external` in this [page](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computeforwardingrule)) 


### Rule 1

We will continue supporting the existing bebavior when migrating to the Direct Resource. 


### Rule 2

Each of those ambiguous usages requires special handling, we shall limit those special handling code inside the corresponding Direct Resource resource code base and mark them with “legacy” comment to avoid repeating the bad design. 
