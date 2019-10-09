# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project's packages adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [v0.6.0]

### Changed

- Upgraded to kube-state-metrics [new release 1.8.0](https://github.com/kubernetes/kube-state-metrics/releases/tag/v1.8.0)
- Fix templating for image registry.

## [v0.5.0]

### Changed

- Migrated to be deployed via an app CR not a chartconfig CR.

## [v0.4.0] 

### Changed

- Upgraded to kube state metric [new release 1.7.2](https://github.com/kubernetes/kube-state-metrics/releases/tag/v1.7.2)

## [v0.3.4] 

### Changed

- Increase memory resources from 300Mi to 350Mi.

## [v0.3.3]

### Changed

- Change prioty class to `giantswarm-critical`.

## [v0.3.2]

### Changed

- Add missing RBAC resources to cluster role.

## [v0.3.1]

### Changed

- Add `apps` api group in RBAC for deployments, daemonsets and replicas.

## [v0.3.0]

### Changed

- Upgraded to kube state metric [new release 1.6.0](https://github.com/kubernetes/kube-state-metrics/releases/tag/v1.6.0)
- Tunned the addon resizer for bigger clusters.

[0.4.0]: https://github.com/giantswarm/kubernetes-kube-state-metrics/pull/49
[0.3.4]: https://github.com/giantswarm/kubernetes-kube-state-metrics/pull/47
[0.3.1]: https://github.com/giantswarm/kubernetes-kube-state-metrics/pull/43
[0.3.0]: https://github.com/giantswarm/kubernetes-kube-state-metrics/pull/40
