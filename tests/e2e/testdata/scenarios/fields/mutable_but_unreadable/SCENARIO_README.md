For the LoggingLogMetric resource the metadata fields are:

```yaml
    cnrm.cloud.google.com/mutable-but-unreadable-fields: '{"spec":{"metricDescriptor":{"launchStage":"EARLY_ACCESS","metadata":{"ingestDelay":"1s","samplePeriod":"1s"}}}}'
```

What we want to see is that if we update the spec (the user intent) of our LoggingLogMetric resource WITH the same values for our mutable but unreadable fields, that update DOES NOT cause an underlying cloud provider API call. That is, there is no http log for the apply.