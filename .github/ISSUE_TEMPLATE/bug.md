---
name: Bug
about: Report something that doesn't work how it should
title: ''
labels: bug
assignees: ''

---

**Describe the bug**
A clear and concise description of what the bug is.

**ConfigConnector Version**
Run the following command to get the current ConfigConnector version
```bash 
kubectl get ns cnrm-system -o jsonpath='{.metadata.annotations.cnrm\.cloud\.google\.com/version}' 
```

**To Reproduce**
*Steps to reproduce the behavior:*

*YAML snippets:*
<!--- Please provide the YAML snippets to reproduce the issue. See the following YAML sample as an example --->

```yaml
apiVersion: pubsub.cnrm.cloud.google.com/v1beta1
kind: PubSubTopic
metadata:
  labels:
    label-one: "value-one"
  name: pubsubtopic-sample
```

**Expected behavior**
A clear and concise description of what you expected to happen.
