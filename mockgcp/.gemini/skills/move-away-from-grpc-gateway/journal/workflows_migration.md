## workflows migration
- Moved mockgcp workflows to httptogrpc.
- Both workflows and workflowsexecution generated protos were removed.
- mockworkflowexecution was already using httptogrpc but its generated protos were still present in mockgcp/generated.
- Updated mockworkflows to use cloud.google.com/go/workflows/apiv1/workflowspb.
- Removed RewriteError as it's not supported by httptogrpc, confirmed tests still pass.