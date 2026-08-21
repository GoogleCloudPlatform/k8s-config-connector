For the LoggingLogMetric we want to treat `spec.metricDescriptor.labels[].valueType` and `spec.metricDescriptor.valueType` set to "STRING" or "" as the same value!

In other words, we expect no further `POST`s onto the cloud provider if we change one of the values from "STRING" to "" or vice versa. Note `GET`s are find for the underlying controllers to get the "live" state of the API object.