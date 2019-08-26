# kube-state-metrics Helm Chart
Helm Chart for kube-state-metrics in Tenant Clusters.

* Installs the [kube-state-metrics agent](https://github.com/kubernetes/kube-state-metrics).

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

Deployment to Tenant Clusters is handled by [app-operator](https://github.com/giantswarm/app-operator).
