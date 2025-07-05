package helloworld

import (
	"context"

	clusterv1alpha1 "github.com/karmada-io/karmada/pkg/apis/cluster/v1alpha1"
	workv1alpha2 "github.com/karmada-io/karmada/pkg/apis/work/v1alpha2"
	"github.com/karmada-io/karmada/pkg/scheduler/framework"
	"k8s.io/klog/v2"
)

const (
	// Name is the name of the plugin used in the plugin registry and configurations.
	Name = "HelloWorld"
)

// HelloWorld is a simple plugin that implements the Karmada FilterPlugin interface.
type HelloWorld struct {
}

var _ framework.FilterPlugin = &HelloWorld{}

func New() (framework.Plugin, error) {
	klog.Infof("HelloWorld Plugin: Factory 'New' function called.")
	return &HelloWorld{}, nil
}

func (p *HelloWorld) Name() string {
	return Name
}

// Filter implements the filtering logic of the TestFilter plugin.
func (p *HelloWorld) Filter(ctx context.Context,
	bindingSpec *workv1alpha2.ResourceBindingSpec, bindingStatus *workv1alpha2.ResourceBindingStatus, cluster *clusterv1alpha1.Cluster) *framework.Result {

	// Get the name and namespace of the resource that is being scheduled.
	// We get this from the 'Resource' field within the binding's spec.
	resourceName := bindingSpec.Resource.Name
	resourceNamespace := bindingSpec.Resource.Namespace

	// We just print a log message to prove that our plugin was called,
	// and to verify that we are getting the correct name and namespace.
	klog.Infof(
		" HelloWorld Filter Plugin: Successfully called for resource '%s/%s' on cluster '%s'",
		resourceNamespace,
		resourceName,
		cluster.Name,
	)

	return framework.NewResult(framework.Success)
}
