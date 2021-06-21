module github.com/giantswarm/kube-state-metrics-app

go 1.14

require (
	github.com/giantswarm/apptest v0.11.0
	github.com/giantswarm/backoff v0.2.0
	github.com/giantswarm/microerror v0.3.0
	github.com/giantswarm/micrologger v0.5.0
	github.com/stretchr/testify v1.6.1 // indirect
	k8s.io/apiextensions-apiserver v0.19.2 // indirect
	k8s.io/apimachinery v0.19.2
	sigs.k8s.io/controller-runtime v0.6.5 // indirect
)

replace github.com/coreos/etcd v3.3.10+incompatible => github.com/coreos/etcd v3.3.25+incompatible
