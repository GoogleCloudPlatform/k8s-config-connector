# Step-by-step (scripted) tests

These tests run one-step-at-a-time, and allow for more complex steps
than some of our other yaml-driven tests.

The `script.yaml` file contains a set of kube objects, which are applied
in turn.  After each object is applied, we run some golden checks:

* We export the GCP object and we golden-compare to _exportNN.yaml
* We read the KRM object from the kuberneters cluster, and we golden-compare to _objectNN.yaml


We also support a few "special actions", which are triggered by setting
a top-level field `TEST` on the object:

* Setting TEST: APPLY is a no op as this is the default value for the TEST field.

* Setting TEST: APPLY-NO-WAIT will apply the object without waiting for the object
  to be makred as ready. We will also not export the object.

* Setting `TEST: DELETE` will delete the KCC object and wait for the deltion
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

* Setting `TEST: CLOUD-APPLY` will apply the KCC object, if supported, directly onto the
cloud provider. For now GCP is the only supported provider.

* Setting `TEST: WAIT-FOR-HTTP-REQUEST`along with `VALUE_PRESENT: your value` will apply the object
  and inspect the http log to check that the value in VALUE_PRESENT appears. The step will
  wait ~ seconds for that value to show up.

* Setting `WRITE-KUBE-OBJECT: false` will not export the KRM objects of the KCC resource.