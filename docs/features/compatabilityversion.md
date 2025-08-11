# Backward Compatibility Version

The backward compatibility version feature allows you to specify the version of Config Connector that a namespace should align with. This provides a mechanism to opt-out of new features or behavioral changes introduced in newer versions of Config Connector, ensuring stability and a predictable experience during upgrades.

By setting a specific version in the `ConfigConnectorContext` for a namespace, you instruct Config Connector to operate as if it were that older version, but only for the resources within that namespace.

This is particularly useful when you are managing a multi-tenant cluster and want to roll out upgrades incrementally, or if you need to temporarily pin a namespace to an older version to mitigate an issue or a breaking change.

## How it works

You can control this feature by setting the `spec.version` field in the `ConfigConnectorContext` object for a given namespace. When this field is set, the Config Connector controller for that namespace will adapt its behavior to match the specified version.

If the `spec.version` field is not set, the namespace will default to the behavior of the currently installed Config Connector version.

### Example

Here is an example of a `ConfigConnectorContext` resource that pins the namespace to version `1.126.0`:

```yaml
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnectorContext
metadata:
  name: configconnectorcontext.core.cnrm.cloud.google.com
  namespace: your-namespace
spec:
  version: "1.126.0"
  googleServiceAccount: "your-gsa@your-project.iam.gserviceaccount.com"
```

## Supported Versions

The range of supported backward compatibility versions depends on the version of Config Connector installed. Each version of Config Connector supports a specific range of older versions.

The following table shows an example of the supported compatibility versions for a given installed version of Config Connector.

| Installed KCC Version | Supported Compatibility Versions |
| --------------------- | -------------------------------- |
| 1.133.0               | 1.126.0 - 1.133.0                |
| 1.132.0               | 1.125.0 - 1.132.0                |
| 1.131.0               | 1.124.0 - 1.131.0                |
| 1.130.0               | 1.123.0 - 1.130.0                |
| 1.129.0               | 1.122.0 - 1.129.0                |
| 1.128.0               | 1.121.0 - 1.128.0                |
| 1.127.0               | 1.120.0 - 1.127.0                |
| 1.126.0               | 1.126.0                          |

**Note:** The table above is an example. Please refer to the official release notes for the exact supported version range for your installed version of Config Connector.

## Caveats

- Setting a version outside the supported range will result in an error.
- This feature should be used as a temporary measure to ensure stability during upgrades, not as a long-term solution to avoid upgrading.
- While this feature provides backward compatibility, it is always recommended to test your configurations against the latest version of Config Connector in a non-production environment before rolling out upgrades.
