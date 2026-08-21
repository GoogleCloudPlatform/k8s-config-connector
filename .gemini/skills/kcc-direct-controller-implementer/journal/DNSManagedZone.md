# DNS Service Journal

### 2026-06-05 DNSManagedZone Direct Controller Implementation
- **Context**: Migration of `DNSManagedZone` under `dns.cnrm.cloud.google.com/v1beta1` to the Direct controller model.
- **Problem**: The underlying older-style client `google.golang.org/api/dns/v1` doesn't use protobufs, making `tags.DiffForTopLevelFields` non-applicable. Also, asynchronous operations are handled using `*api.Operation` for Updates/Patches, but `Create` returns `*api.ManagedZone` synchronously.
- **Solution**:
  1. Utilized `DNSManagedZoneSpec_FromAPI` to reconstruct the Spec representation of the actual GCP object, and compared individual updatable fields in Go using `reflect.DeepEqual`.
  2. Implemented `waitForDNSOp` utilizing `common.WaitForDoneOrTimeout` to poll `gcpClient.ManagedZoneOperations.Get` on update.
  3. Ensured `resource.Name = a.id.ManagedZone` and `desiredPb.Name = a.id.ManagedZone` are set during creation/update to prevent creating managed zones with an empty name.
- **Impact**: Clean direct reconciliation of DNSManagedZone, 100% test-suite validation against MockGCP, and seamless fallback-ready configuration in static config.
