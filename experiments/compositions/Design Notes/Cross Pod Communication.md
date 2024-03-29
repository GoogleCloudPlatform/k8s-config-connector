# Cross Pod Communication

We need a way to pass manifests and parameters between the Expander pods and the Allotrope manager pods.

## Options

### Option 1: Config Map

Load template and params in a configmap and pass it on to the expander pods.
The expander pods write a result configmap that is used by the manager.

ConfigMaps limit the size of template and params. It could work for small manifests. POC maybe.

### Option 2: grpc

Similar to KPT we could use GRPC to communicate b/w the manager and the expander pods.
The manager would wait synchronously for response from the expander pods.

### Option 3: CRD

Similar to Configmap, use a custom CRD for this purpose.
This option has the same limitations as the Configmap option.

### Option 4: Shared File System

We use a shared file system provider that is mounted on the manger as well as the expander pods.
The communication is via file locks or explicit handshakes via k8s API or direct calls.

### Option 5: Object Storage

We use an object storage to post the manifests and the expanders would consume it from there.
Results are written back to the object storage.
Communication can be via k8s API (CRDs).

## POC

For the POC we are using Option 3 (inline in CRD).
`manifests/inline` implement this option as a sidecar.