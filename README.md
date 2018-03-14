[![CircleCI](https://circleci.com/gh/giantswarm/kubernetes-kube-state-metrics.svg?style=svg&circle-token=1d6a6248b1d64bd698c7b68801a879ecc9e185f8)](https://circleci.com/gh/giantswarm/kubernetes-kube-state-metrics)

# kube-state-metrics Helm Chart
Helm Chart for kube-state-metrics in Guest Clusters

* Installs the [kube-state-metrics agent](https://github.com/kubernetes/kube-state-metrics).

## Installing the Chart

To install the chart locally:

```bash
$ git clone https://github.com/giantswarm/kubernetes-kube-state-metrics.git
$ cd kubernetes-kube-state-metrics
$ helm install kubernetes-kube-state-metrics/helm/kubernetes-kube-state-metrics-chart
```

## Configuration

| Parameter               | Description                                     | Default                                 |
|-------------------------|-------------------------------------------------|-----------------------------------------|
| `name`                  | The name of the service                         | `kube-state-metrics`                    |
| `namespace`             | The namespaces the services runs in             | `kube-system`                           |
| `image.repository`      | The image repository to pull from               | `quay.io/giantswarm/kube-state-metrics` |
| `image.tag`             | The image tag to pull from                      | `v1.2.0`                                |
| `replicas`              | The number of replicas of the container         | `1`                                     |
| `port`                  | The port of the container                       | `10301`                                 |
| `resources`             | kube-state-metrics resource requests and limits | `cpu:50m  - memory:75Mi`                |
| `test.image.repository` | The test image repository to pull from          | `quay.io/giantswarm/alpine-testing`     |
| `test.image.tag`        | The test image tag to pull from                 | `0.1.0`                                 |
