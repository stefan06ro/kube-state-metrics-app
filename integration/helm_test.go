// +build k8srequired

package integration

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/giantswarm/e2e-harness/pkg/framework"
	"github.com/giantswarm/microerror"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	resourceNamespace = "kube-system"
)

func TestHelm(t *testing.T) {
	channel := os.Getenv("CIRCLE_SHA1")

	err := framework.HelmCmd(fmt.Sprintf("registry install --wait quay.io/giantswarm/kubernetes-kube-state-metrics-chart:%s -n test-deploy", channel))
	if err != nil {
		t.Errorf("unexpected error during installation of the chart: %v", err)
	}
	defer framework.HelmCmd("delete test-deploy --purge")

	err = checkDeployment()
	if err != nil {
		t.Fatalf("deployment manifest is incorrect: %v", err)
	}

	err = framework.HelmCmd("test --debug --cleanup test-deploy")
	if err != nil {
		t.Errorf("unexpected error during test of the chart: %v", err)
	}
}

// TestMigration ensures that previously deployed resources are properly removed.
// It installs a chart with the same resources as kube-state-metrics with
// appropriate labels so that we can query for them. Then installs the
// kube-state-metrics chart and checks that the previous resources are removed
// and the ones for kube-state-metrics are in place.
func TestMigration(t *testing.T) {
	// Install legacy resources.
	err := framework.HelmCmd("install /e2e/fixtures/resources-chart -n resources")
	if err != nil {
		t.Fatalf("could not install resources chart: %v", err)
	}
	defer framework.HelmCmd("delete resources --purge")

	// Check legacy resources are present.
	err = checkResourcesPresent("app=kube-state-metrics,kind=legacy")
	if err != nil {
		t.Fatalf("could check legacy resources present: %v", err)
	}
	// Check managed resources are not present.
	err = checkResourcesNotPresent("app=kube-state-metrics,giantswarm.io/service-type=managed")
	if err != nil {
		t.Fatalf("could check managed resources not present: %v", err)
	}

	// Install kubernetes-kube-state-metrics-chart.
	channel := os.Getenv("CIRCLE_SHA1")
	err = framework.HelmCmd(fmt.Sprintf("registry install --wait quay.io/giantswarm/kubernetes-kube-state-metrics-chart:%s -n test-deploy", channel))
	if err != nil {
		t.Fatalf("could not install kubernetes-kube-state-metrics-chart: %v", err)
	}
	defer framework.HelmCmd("delete test-deploy --purge")

	// Check legacy resources are not present.
	err = checkResourcesNotPresent("app=kube-state-metrics,kind=legacy")
	if err != nil {
		t.Fatalf("could check legacy resources present: %v", err)
	}
	// Check managed resources are present.
	err = checkResourcesPresent("app=kube-state-metrics,giantswarm.io/service-type=managed")
	if err != nil {
		t.Fatalf("could check managed resources not present: %v", err)
	}
}

func checkResourcesPresent(labelSelector string) error {
	c := h.K8sClient()
	listOptions := metav1.ListOptions{
		LabelSelector: labelSelector,
	}

	d, err := c.Extensions().Deployments(resourceNamespace).List(listOptions)
	if err != nil {
		return microerror.Mask(err)
	}
	if len(d.Items) != 1 {
		return microerror.Newf("unexpected number of deployments, want 1, got %d", len(d.Items))
	}

	r, err := c.Rbac().ClusterRoles().List(listOptions)
	if err != nil {
		return microerror.Mask(err)
	}
	if len(r.Items) != 1 {
		return microerror.Newf("unexpected number of roles, want 1, got %d", len(r.Items))
	}

	rb, err := c.Rbac().ClusterRoleBindings().List(listOptions)
	if err != nil {
		return microerror.Mask(err)
	}
	if len(rb.Items) != 1 {
		return microerror.Newf("unexpected number of rolebindings, want 1, got %d", len(rb.Items))
	}

	s, err := c.Core().Services(resourceNamespace).List(listOptions)
	if err != nil {
		return microerror.Mask(err)
	}
	if len(s.Items) != 1 {
		return microerror.Newf("unexpected number of services, want 1, got %d", len(s.Items))
	}

	sa, err := c.Core().ServiceAccounts(resourceNamespace).List(listOptions)
	if err != nil {
		return microerror.Mask(err)
	}
	if len(sa.Items) != 1 {
		return microerror.Newf("unexpected number of serviceaccountss, want 1, got %d", len(sa.Items))
	}

	return nil
}

func checkResourcesNotPresent(labelSelector string) error {
	c := h.K8sClient()
	listOptions := metav1.ListOptions{
		LabelSelector: labelSelector,
	}

	d, err := c.Extensions().Deployments(resourceNamespace).List(listOptions)
	if err == nil && len(d.Items) > 0 {
		return microerror.New("expected error querying for deployments didn't happen")
	}
	if !apierrors.IsNotFound(err) {
		return microerror.Mask(err)
	}

	r, err := c.Rbac().ClusterRoles().List(listOptions)
	if err == nil && len(r.Items) > 0 {
		return microerror.New("expected error querying for roles didn't happen")
	}
	if !apierrors.IsNotFound(err) {
		return microerror.Mask(err)
	}

	rb, err := c.Rbac().ClusterRoleBindings().List(listOptions)
	if err == nil && len(rb.Items) > 0 {
		return microerror.New("expected error querying for rolebindings didn't happen")
	}
	if !apierrors.IsNotFound(err) {
		return microerror.Mask(err)
	}

	s, err := c.Core().Services(resourceNamespace).List(listOptions)
	if err == nil && len(s.Items) > 0 {
		return microerror.New("expected error querying for services didn't happen")
	}
	if !apierrors.IsNotFound(err) {
		return microerror.Mask(err)
	}

	sa, err := c.Core().ServiceAccounts(resourceNamespace).List(listOptions)
	if err == nil && len(sa.Items) > 0 {
		return microerror.New("expected error querying for serviceaccounts didn't happen")
	}
	if !apierrors.IsNotFound(err) {
		return microerror.Mask(err)
	}

	return nil
}

// checkDeployment ensures that key properties of the kube-state-metrics
// deployment are correct.
func checkDeployment() error {
	name := "kube-state-metrics"
	expectedMatchLabels := map[string]string{
		"app": "kube-state-metrics",
	}
	expectedLabels := map[string]string{
		"app": "kube-state-metrics",
		"giantswarm.io/service-type": "managed",
	}
	expectedReplicas := 1

	c := h.K8sClient()
	ds, err := c.Apps().Deployments(resourceNamespace).Get(name, metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		return microerror.Newf("could not find daemonset: '%s' %v", name, err)
	} else if err != nil {
		return microerror.Newf("unexpected error getting daemonset: %v", err)
	}

	// Check daemonset labels.
	if !reflect.DeepEqual(expectedLabels, ds.ObjectMeta.Labels) {
		return microerror.Newf("expected labels: %v got: %v", expectedLabels, ds.ObjectMeta.Labels)
	}

	// Check selector match labels.
	if !reflect.DeepEqual(expectedMatchLabels, ds.Spec.Selector.MatchLabels) {
		return microerror.Newf("expected match labels: %v got: %v", expectedMatchLabels, ds.Spec.Selector.MatchLabels)
	}

	// Check pod labels.
	if !reflect.DeepEqual(expectedLabels, ds.Spec.Template.ObjectMeta.Labels) {
		return microerror.Newf("expected pod labels: %v got: %v", expectedLabels, ds.Spec.Template.ObjectMeta.Labels)
	}

	// Check replica count.
	if *ds.Spec.Replicas != int32(expectedReplicas) {
		return microerror.Newf("expected replicas: %d got: %d", expectedReplicas, ds.Spec.Replicas)
	}

	return nil
}
