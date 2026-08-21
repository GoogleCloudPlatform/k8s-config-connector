# 🎉 Config Connector Go Client: A Fresh Start!

We're excited to share that the Config Connector Go client is getting a major upgrade! We're moving to a more direct approach using `sigs.k8s.io/controller-runtime` and popular tools like `controller-gen` and `kubebuilder`. This means simpler, more reliable interactions with Config Connector resources.

Find the new APIs in the `./apis` directory and use the [Kubernetes dynamic client](https://pkg.go.dev/k8s.io/client-go/dynamic) for API calls.

We're excited about this new chapter for the Go client and believe it will significantly improve your workflow. Happy coding!  

## Regarding the existing `./pkg/clients/generated` directory

* It will remain in place for now, but please note it's an Alpha library and future plans may change (track progress in issue #3037).
* It's currently up-to-date with the 1.125 release, but won't receive further updates.
