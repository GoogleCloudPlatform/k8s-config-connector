This scenario meant to do the following:
- Apply a CC to turn to instruct KCC to be in "Namespaced" mode
- Apply a CCC to manage a namespace
- Apply a KCC resource with a mutable field ("description"), in this case, ArtifactRegistryRepository
>  NOTE: This resources should be managed by the CCC we created
- Pause KCC by setting the CCC's object `actuationMode` to Paused.
- Apply the same KCC resource and change the "description" field. Observe no change to the resource on the cloud provider.
> NOTE: The way to assess that there is no change in the actual resource is by determining that the http log 
> file does not include any calls to the GCP provider. IOW, the log file is not present.
- "Un pause" KCC by setting the CCC's object `actuationMode` to Reconciling.
- Observe the resource from the cloud provider matches the previous intent.
