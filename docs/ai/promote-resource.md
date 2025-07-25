# Promoting a resource

To promote a resource from v1alpha1 to v1beta1:

* Locate the file containing the type, for example apis/cloudbuild/v1alpha1/workerpool_types.go.  Typically the format is `apis/<service>/<version>/<kind>_types.go`.

* Locate the controller for the type, for example `pkg/controller/direct/cloudbuild/workerpool_controller.go`.  Typically the format is `pkg/controller/direct/<sevrvice>/<kind>_controller.go.

* Add the new type to the `generate.sh` script under v1beta1, and remove the type from `v1alpha1`.  If this is the first type being promoted for the service you will need to create a `generate.sh` script - use the `v1alpha1` generate.sh script as a reference. 

* Move the types file to `apis/<service>/v1beta1` and update the package declaration from v1alpha1 to v1beta1.
 mv apis/backupdr/v1alpha1/backupplan_types.go apis/backupdr/v1beta1/ && mv                       │
 │   apis/backupdr/v1alpha1/backupplan_identity.go apis/backupdr/v1beta1/ && mv                       │
 │   apis/backupdr/v1alpha1/backupplan_reference.go apis/backupdr/v1beta1/                            │
 │  
* Update the controller import for the type to import the `v1beta1` package (because the type has moved)

* Make the v1beta1 generate.sh script executable (`chmod +x`) and then run it.