// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/kubewharf/kubeadmiral/pkg/apis/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSchedulerPluginWebhookConfigurations implements SchedulerPluginWebhookConfigurationInterface
type FakeSchedulerPluginWebhookConfigurations struct {
	Fake *FakeCoreV1alpha1
}

var schedulerpluginwebhookconfigurationsResource = v1alpha1.SchemeGroupVersion.WithResource("schedulerpluginwebhookconfigurations")

var schedulerpluginwebhookconfigurationsKind = v1alpha1.SchemeGroupVersion.WithKind("SchedulerPluginWebhookConfiguration")

// Get takes name of the schedulerPluginWebhookConfiguration, and returns the corresponding schedulerPluginWebhookConfiguration object, and an error if there is any.
func (c *FakeSchedulerPluginWebhookConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SchedulerPluginWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(schedulerpluginwebhookconfigurationsResource, name), &v1alpha1.SchedulerPluginWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SchedulerPluginWebhookConfiguration), err
}

// List takes label and field selectors, and returns the list of SchedulerPluginWebhookConfigurations that match those selectors.
func (c *FakeSchedulerPluginWebhookConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SchedulerPluginWebhookConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(schedulerpluginwebhookconfigurationsResource, schedulerpluginwebhookconfigurationsKind, opts), &v1alpha1.SchedulerPluginWebhookConfigurationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SchedulerPluginWebhookConfigurationList{ListMeta: obj.(*v1alpha1.SchedulerPluginWebhookConfigurationList).ListMeta}
	for _, item := range obj.(*v1alpha1.SchedulerPluginWebhookConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested schedulerPluginWebhookConfigurations.
func (c *FakeSchedulerPluginWebhookConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(schedulerpluginwebhookconfigurationsResource, opts))
}

// Create takes the representation of a schedulerPluginWebhookConfiguration and creates it.  Returns the server's representation of the schedulerPluginWebhookConfiguration, and an error, if there is any.
func (c *FakeSchedulerPluginWebhookConfigurations) Create(ctx context.Context, schedulerPluginWebhookConfiguration *v1alpha1.SchedulerPluginWebhookConfiguration, opts v1.CreateOptions) (result *v1alpha1.SchedulerPluginWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(schedulerpluginwebhookconfigurationsResource, schedulerPluginWebhookConfiguration), &v1alpha1.SchedulerPluginWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SchedulerPluginWebhookConfiguration), err
}

// Update takes the representation of a schedulerPluginWebhookConfiguration and updates it. Returns the server's representation of the schedulerPluginWebhookConfiguration, and an error, if there is any.
func (c *FakeSchedulerPluginWebhookConfigurations) Update(ctx context.Context, schedulerPluginWebhookConfiguration *v1alpha1.SchedulerPluginWebhookConfiguration, opts v1.UpdateOptions) (result *v1alpha1.SchedulerPluginWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(schedulerpluginwebhookconfigurationsResource, schedulerPluginWebhookConfiguration), &v1alpha1.SchedulerPluginWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SchedulerPluginWebhookConfiguration), err
}

// Delete takes name of the schedulerPluginWebhookConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeSchedulerPluginWebhookConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(schedulerpluginwebhookconfigurationsResource, name, opts), &v1alpha1.SchedulerPluginWebhookConfiguration{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSchedulerPluginWebhookConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(schedulerpluginwebhookconfigurationsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SchedulerPluginWebhookConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched schedulerPluginWebhookConfiguration.
func (c *FakeSchedulerPluginWebhookConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SchedulerPluginWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(schedulerpluginwebhookconfigurationsResource, name, pt, data, subresources...), &v1alpha1.SchedulerPluginWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SchedulerPluginWebhookConfiguration), err
}
