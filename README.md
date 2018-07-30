[![CircleCI](https://circleci.com/gh/giantswarm/kubernetes-kube-state-metrics.svg?style=svg&circle-token=1d6a6248b1d64bd698c7b68801a879ecc9e185f8)](https://circleci.com/gh/giantswarm/kubernetes-kube-state-metrics)

# kube-state-metrics Helm Chart
Helm Chart for kube-state-metrics in Guest Clusters.

* Installs the [kube-state-metrics agent](https://github.com/kubernetes/kube-state-metrics).

## Installing the Chart

To install the chart locally:

```bash
$ git clone https://github.com/giantswarm/kubernetes-kube-state-metrics.git
$ cd kubernetes-kube-state-metrics
$ helm install helm/kubernetes-kube-state-metrics-chart
```

Provide a custom `values.yaml`:

```bash
$ helm install kubernetes-kube-state-metrics-chart -f values.yaml
```

Deployment to Guest Clusters is handled by [chart-operator](https://github.com/giantswarm/chart-operator).
