# kubebuilder-declarative-pattern

kubebuilder-declarative-pattern provides a set of tools for building declarative cluster operators with kubebuilder. Declarative operators provide a fast path to orchestrating Kubernetes deployments to enable domain experts to focus on their component instead of re-answering questions like 'How do I get this YAML into the cluster?' or 'How do I update it?'.

## Development

### Running Smoke Tests

Smoke tests are provided to ensure basic functionality of the framework against example operators. They should be run as part of significant code changes. The tests require a running Kubernetes cluster to be targeted from the local machine and write access to a GCR bucket.

```bash
cd hack
IMG=<a writeable image path, eg, gcr.io/my-project/controller:latest> go run smoketest.go
```

## Documentation

- [Building an Operator (walkthrough)](./docs/addon/walkthrough/README.md)
- [Pattern Documentation](https://godoc.org/sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns)
- [Managing Addons with Operators (Video, KubeCon'18)](https://www.youtube.com/watch?v=LPejvfBR5_w)

## Community, discussion, contribution, and support

Learn how to engage with the Kubernetes community on the [community page](http://kubernetes.io/community/).

You can reach the maintainers of this project at:

- [Slack](http://slack.k8s.io/)
- [Mailing List](https://groups.google.com/forum/#!forum/kubernetes-dev)

### Code of conduct

Participation in the Kubernetes community is governed by the [Kubernetes Code of Conduct](code-of-conduct.md).

[owners]: https://git.k8s.io/community/contributors/guide/owners.md
[Creative Commons 4.0]: https://git.k8s.io/website/LICENSE
