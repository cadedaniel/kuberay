// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/ray-project/kuberay/ray-operator/apis/ray/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRayJobs implements RayJobInterface
type FakeRayJobs struct {
	Fake *FakeRayV1alpha1
	ns   string
}

var rayjobsResource = schema.GroupVersionResource{Group: "ray.io", Version: "v1alpha1", Resource: "rayjobs"}

var rayjobsKind = schema.GroupVersionKind{Group: "ray.io", Version: "v1alpha1", Kind: "RayJob"}

// Get takes name of the rayJob, and returns the corresponding rayJob object, and an error if there is any.
func (c *FakeRayJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RayJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rayjobsResource, c.ns, name), &v1alpha1.RayJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RayJob), err
}

// List takes label and field selectors, and returns the list of RayJobs that match those selectors.
func (c *FakeRayJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RayJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rayjobsResource, rayjobsKind, c.ns, opts), &v1alpha1.RayJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RayJobList{ListMeta: obj.(*v1alpha1.RayJobList).ListMeta}
	for _, item := range obj.(*v1alpha1.RayJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rayJobs.
func (c *FakeRayJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rayjobsResource, c.ns, opts))

}

// Create takes the representation of a rayJob and creates it.  Returns the server's representation of the rayJob, and an error, if there is any.
func (c *FakeRayJobs) Create(ctx context.Context, rayJob *v1alpha1.RayJob, opts v1.CreateOptions) (result *v1alpha1.RayJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rayjobsResource, c.ns, rayJob), &v1alpha1.RayJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RayJob), err
}

// Update takes the representation of a rayJob and updates it. Returns the server's representation of the rayJob, and an error, if there is any.
func (c *FakeRayJobs) Update(ctx context.Context, rayJob *v1alpha1.RayJob, opts v1.UpdateOptions) (result *v1alpha1.RayJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rayjobsResource, c.ns, rayJob), &v1alpha1.RayJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RayJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRayJobs) UpdateStatus(ctx context.Context, rayJob *v1alpha1.RayJob, opts v1.UpdateOptions) (*v1alpha1.RayJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(rayjobsResource, "status", c.ns, rayJob), &v1alpha1.RayJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RayJob), err
}

// Delete takes name of the rayJob and deletes it. Returns an error if one occurs.
func (c *FakeRayJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(rayjobsResource, c.ns, name, opts), &v1alpha1.RayJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRayJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rayjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RayJobList{})
	return err
}

// Patch applies the patch and returns the patched rayJob.
func (c *FakeRayJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RayJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rayjobsResource, c.ns, name, pt, data, subresources...), &v1alpha1.RayJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RayJob), err
}
