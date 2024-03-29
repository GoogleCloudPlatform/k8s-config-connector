---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This file is automatically generated by Magic Modules and manual
#     changes will be clobbered when the file is regenerated.
#
#     Please read more about how to change this file in
#     .github/CONTRIBUTING.md.
#
# ----------------------------------------------------------------------------
subcategory: "Game Servers"
description: |-
  A Realm resource.
---

# google\_game\_services\_realm

A Realm resource.


To get more information about Realm, see:

* [API documentation](https://cloud.google.com/game-servers/docs/reference/rest/v1beta/projects.locations.realms)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/game-servers/docs)

## Example Usage - Game Service Realm Basic


```hcl
resource "google_game_services_realm" "default" {
  realm_id  = "tf-test-realm"
  time_zone = "EST"
  location  = "global"

  description = "one of the nine"
}
```

## Argument Reference

The following arguments are supported:


* `time_zone` -
  (Required)
  Required. Time zone where all realm-specific policies are evaluated. The value of
  this field must be from the IANA time zone database:
  https://www.iana.org/time-zones.

* `realm_id` -
  (Required)
  GCP region of the Realm.


- - -


* `labels` -
  (Optional)
  The labels associated with this realm. Each label is a key-value pair.

* `description` -
  (Optional)
  Human readable description of the realm.

* `location` -
  (Optional)
  Location of the Realm.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.


## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/realms/{{realm_id}}`

* `name` -
  The resource id of the realm, of the form:
  `projects/{project_id}/locations/{location}/realms/{realm_id}`. For
  example, `projects/my-project/locations/{location}/realms/my-realm`.

* `etag` -
  ETag of the resource.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


Realm can be imported using any of these accepted formats:

```
$ terraform import google_game_services_realm.default projects/{{project}}/locations/{{location}}/realms/{{realm_id}}
$ terraform import google_game_services_realm.default {{project}}/{{location}}/{{realm_id}}
$ terraform import google_game_services_realm.default {{location}}/{{realm_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
