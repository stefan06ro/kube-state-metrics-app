module github.com/giantswarm/kube-state-metrics-app

go 1.14

require (
	github.com/giantswarm/apptest v0.12.0
	github.com/giantswarm/backoff v0.2.0
	github.com/giantswarm/microerror v0.3.0
	github.com/giantswarm/micrologger v0.5.0
	k8s.io/apimachinery v0.20.10
)

replace github.com/coreos/etcd v3.3.10+incompatible => github.com/coreos/etcd v3.3.25+incompatible
