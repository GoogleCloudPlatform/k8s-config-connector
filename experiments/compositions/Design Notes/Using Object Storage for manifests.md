# Object storage for manifests

The object location (path) is created by the Allotrope Manager and then passed on to the Expander pods as part of the pod bring-up.
The expander pods should read the manifests from that location and then expand them and write it back to the same location.

## Object Storage provider

Options:
1. Minio
2. Rook
3. GCS + Fuse (https://cloud.google.com/kubernetes-engine/docs/concepts/storage-overview#fuse)

## Folder structure (path)
We could consider having a folder structure that shows the expansion path OR use a single folder that the expanders would write to.


Option 1:

expansionCRD.name/template.crd.name/in       << manager copies the template+params here
expansionCRD.name/template.crd.name/expander1.name << expander1 reads from in/ & expands here
expansionCRD.name/template.crd.name/expander2.name << expander2 reads from expander1 & expands here
expansionCRD.name/template.crd.name/expander3.name << expander3 reads from expander2 & expands here
expansionCRD.name/template.crd.name/out      << manager copies from expander3 to here

Option 2:

expansionCRD.name/template.crd.name/in       << manager copies the template+params here
expansionCRD.name/template.crd.name/out   << all expanders read from in/ & out/ and continue to expand in out/



