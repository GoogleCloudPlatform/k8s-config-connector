# Step-by-step (scripted) tests

These tests run one-step-at-a-time, and allow for more complex steps
than some of our other yaml-driven tests.

The `script.yaml` file contains a set of kube objects, which are applied
in turn.  After each object is applied, we run some golden checks:

* We export the GCP object and we golden-compare to _exportNN.yaml
* We read the KRM object from the kubernetes cluster, and we golden-compare to _objectNN.yaml


We also support a few "special actions", which are triggered by setting
a top-level field `TEST` on the object:

* Setting `TEST: APPLY` is a no op as this is the default value for the TEST field.

* Setting `TEST: APPLY-NO-WAIT` will apply the object without waiting for the object
  to be marked as ready. We will also not export the object.

* Setting `TEST: APPLY-10-SEC` will apply an object that we know is going to fail. 
  i.e. immutable field is modified. Instead of keeping query the error resource, 
  we stop the test after 10s and capture the error log. This action can be used to 
  test the expected error state.

* Setting `TEST: PATCH-EXTERNALLY-MANAGED-FIELDS` will patch the object with
  externally-managed fields, i.e. all the patched fields will have manager
  `cnrm-controller-manager`. This is used to mock the situation when a resource
  has externally-managed fields to be cleaned up.

* Setting `TEST: READ-OBJECT` skips the apply; we read the current value of the
  object without changing it.

* Setting `TEST: READ-OBJECT-AND-COMPARE-SPEC` does what `TEST: READ-OBJECT`
  does and compares the spec of the golden objects of the current step and the
  step configured in `TARGET_STEP_FOR_READ_AND_COMPARE`.

* Setting `TEST: DELETE` will delete the KCC object and wait for the deletion
  to complete; it will automatically skip
  the GCP export and the kube export.  It suffices to set
  apiVersion / kind / namespace / name.

* Setting `TEST: DELETE-NO-WAIT` will delete the KCC object but not wait for
  deletion to complete; unlike `DELETE` it will try to do
  the GCP export and the kube export (the resource and object may still exist).
  It suffices to set apiVersion / kind / namespace / name.

* Setting `TEST: ABANDON` will delete the KCC object after setting the
  `cnrm.cloud.google.com/deletion-policy: abandon` annotation.  The KCC
  object will still be deleted from the kube-apiserver.  It suffices to set
  apiVersion / kind / namespace / name.

* Setting `TEST: WAIT-FOR-HTTP-REQUEST` along with `VALUE_PRESENT: your value` will apply the object
  and inspect the http log to check that the value in VALUE_PRESENT appears. The step will
  wait ~ seconds for that value to show up.

* Setting `WRITE-KUBE-OBJECT: false` will not export the KRM objects of the KCC resource.
