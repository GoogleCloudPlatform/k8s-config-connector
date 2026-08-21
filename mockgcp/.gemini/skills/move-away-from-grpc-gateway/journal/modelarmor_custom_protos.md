# ModelArmor Custom Protos

During the migration of `modelarmor`, it was noted that the protos were referenced in `mockgcp/Makefile` as being in `./third_party/googleapis/mockgcp/cloud/modelarmor/v1/service.proto`. However, these files were not easily locatable in the workspace during the task, possibly due to them being part of the ephemeral `third_party/googleapis` directory created during build.

The service was successfully migrated to use the official `cloud.google.com/go/modelarmor/apiv1/modelarmorpb` package.
