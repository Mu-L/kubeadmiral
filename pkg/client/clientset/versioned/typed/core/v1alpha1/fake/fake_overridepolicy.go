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

// FakeOverridePolicies implements OverridePolicyInterface
type FakeOverridePolicies struct {
	Fake *FakeCoreV1alpha1
	ns   string
}

var overridepoliciesResource = v1alpha1.SchemeGroupVersion.WithResource("overridepolicies")

var overridepoliciesKind = v1alpha1.SchemeGroupVersion.WithKind("OverridePolicy")

// Get takes name of the overridePolicy, and returns the corresponding overridePolicy object, and an error if there is any.
func (c *FakeOverridePolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OverridePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(overridepoliciesResource, c.ns, name), &v1alpha1.OverridePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OverridePolicy), err
}

// List takes label and field selectors, and returns the list of OverridePolicies that match those selectors.
func (c *FakeOverridePolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OverridePolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(overridepoliciesResource, overridepoliciesKind, c.ns, opts), &v1alpha1.OverridePolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OverridePolicyList{ListMeta: obj.(*v1alpha1.OverridePolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.OverridePolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested overridePolicies.
func (c *FakeOverridePolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(overridepoliciesResource, c.ns, opts))

}

// Create takes the representation of a overridePolicy and creates it.  Returns the server's representation of the overridePolicy, and an error, if there is any.
func (c *FakeOverridePolicies) Create(ctx context.Context, overridePolicy *v1alpha1.OverridePolicy, opts v1.CreateOptions) (result *v1alpha1.OverridePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(overridepoliciesResource, c.ns, overridePolicy), &v1alpha1.OverridePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OverridePolicy), err
}

// Update takes the representation of a overridePolicy and updates it. Returns the server's representation of the overridePolicy, and an error, if there is any.
func (c *FakeOverridePolicies) Update(ctx context.Context, overridePolicy *v1alpha1.OverridePolicy, opts v1.UpdateOptions) (result *v1alpha1.OverridePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(overridepoliciesResource, c.ns, overridePolicy), &v1alpha1.OverridePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OverridePolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOverridePolicies) UpdateStatus(ctx context.Context, overridePolicy *v1alpha1.OverridePolicy, opts v1.UpdateOptions) (*v1alpha1.OverridePolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(overridepoliciesResource, "status", c.ns, overridePolicy), &v1alpha1.OverridePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OverridePolicy), err
}

// Delete takes name of the overridePolicy and deletes it. Returns an error if one occurs.
func (c *FakeOverridePolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(overridepoliciesResource, c.ns, name, opts), &v1alpha1.OverridePolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOverridePolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(overridepoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.OverridePolicyList{})
	return err
}

// Patch applies the patch and returns the patched overridePolicy.
func (c *FakeOverridePolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OverridePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(overridepoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.OverridePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OverridePolicy), err
}
