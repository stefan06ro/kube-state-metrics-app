// +build k8srequired

package key

func AppName() string {
	return "kube-state-metrics"
}

func CRName() string {
	return "kube-state-metrics-app"
}

func DefaultCatalogName() string {
	return "default"
}

func DefaultTestCatalogName() string {
	return "default-test"
}

func Namespace() string {
	return "giantswarm"
}
