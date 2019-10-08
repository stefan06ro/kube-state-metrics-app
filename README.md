[![CircleCI](https://circleci.com/gh/giantswarm/kube-state-metrics-app.svg?style=svg)](https://circleci.com/gh/giantswarm/kube-state-metrics-app)

# kube-state-metrics Helm Chart

Helm Chart for kube-state-metrics in Tenant Clusters.

* Installs the [kube-state-metrics agent].

## Deployment

* Managed by [app-operator].
* Production releases are stored in the [default-catalog].
* WIP releases are stored in the [default-test-catalog].

## Installing the Chart

To install the chart locally:

```bash
$ git clone https://github.com/giantswarm/kube-state-metrics-ap.git
$ cd kube-state-metrics-app
$ helm install helm/kube-state-metrics-app
```

Provide a custom `values.yaml`:

```bash
$ helm install kube-state-metrics-app -f values.yaml
```

## Release Process

* Ensure CHANGELOG.md is up to date.
* Create a new GitHub release with the version e.g. `v0.1.0` and link the
changelog entry.
* This will push a new git tag and trigger a new tarball to be pushed to the
[default-catalog].  
* Update [cluster-operator] with the new version.

[app-operator]: https://github.com/giantswarm/app-operator
[cluster-operator]: https://github.com/giantswarm/cluster-operator
[default-catalog]: https://github.com/giantswarm/default-catalog
[default-test-catalog]: https://github.com/giantswarm/default-test-catalog
[kube-state-metrics agent]: https://github.com/kubernetes/kube-state-metrics