# Manifest Storage

The templates and the expanded manifests need a durable storage for processing and traceability. We have a few choices:
1. Inline inside CRDs
2. Object Storage (rook.io, minio, GCS, …)
3. Shared File System (rook.io, nfs, …)

## Options

### Inline
The Manifests are embedded in the `CRD_T` instance.  The limitations with inline includes:
1. Access via KRM API only, requires k8s authentication and k8s client.
2. Upper limit on template size
3. Upper limit on expanded manifest size

For POC we could use the inline option. Beyond POC we would need to consider a more scalable option.

### Shared File System
A shared file system makes programming simple and allows faster iterations since most languages and tools understand POSIX APIs. But this is typically limited to the cluster/pods where the FS is mounted. For external systems to interface with the manifests we would need to mount the filesystem which may be infeasible in many situations (Think Pantheon, …)

### Object Storage
Object Storage makes it scalable compared to Inline, but requires client libs to read and write into the storage. This option allows external systems to interface with the manifests as well, making it more flexible than Shared File System.

## POC

For POC we are starting with inline option.