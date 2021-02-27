---
name: Bug
about: Report something that doesn't work how it should
labels: bug

---
## Checklist

- [ ] I did not find a related open issue.
- [ ] I did not find a solution in the troubleshooting guide: (https://cloud.google.com/config-connector/docs/troubleshooting)
- [ ] If this issue is time-sensitive, I have submitted a corresponding issue with [GCP support](https://cloud.google.com/support-hub).

## Bug Description

<!--- A clear and concise description of the bug. --->

## Additional Diagnostic Information

### Kubernetes Cluster Version

<!--
Run the following command to get the Kubernetes cluster version:

kubectl version --short
-->

### Config Connector Version

<!--
Run the following command to get the installed Config Connector version:

kubectl get ns cnrm-system -o jsonpath='{.metadata.annotations.cnrm\.cloud\.google\.com/version}'
-->

### Config Connector Mode

<!--
Run the following command to get the Config Connector Mode:

kubectl get ConfigConnector "configconnector.core.cnrm.cloud.google.com" -o=jsonpath="{@.spec.mode}"
-->

## Log Output

<!--
Include any relevant log output. See
  https://cloud.google.com/config-connector/docs/troubleshooting#logging

If the issue is for a specific resource, please include those logs.
-->

## Steps to Reproduce

### Steps to reproduce the issue
<!--- Steps needed to reproduce the issue. --->

### YAML snippets
<!--- YAML snippets needed to reproduce the issue. See the following as an example. --->

```yaml
apiVersion: pubsub.cnrm.cloud.google.com/v1beta1
kind: PubSubTopic
metadata:
  labels:
    label-one: "value-one"
  name: pubsubtopic-sample
```
