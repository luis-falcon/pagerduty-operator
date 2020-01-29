// Code generated by main. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/openshift/hive/pkg/apis/hive/v1"
	scheme "github.com/openshift/hive/pkg/client/clientset-generated/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SelectorSyncSetsGetter has a method to return a SelectorSyncSetInterface.
// A group's client should implement this interface.
type SelectorSyncSetsGetter interface {
	SelectorSyncSets() SelectorSyncSetInterface
}

// SelectorSyncSetInterface has methods to work with SelectorSyncSet resources.
type SelectorSyncSetInterface interface {
	Create(*v1.SelectorSyncSet) (*v1.SelectorSyncSet, error)
	Update(*v1.SelectorSyncSet) (*v1.SelectorSyncSet, error)
	UpdateStatus(*v1.SelectorSyncSet) (*v1.SelectorSyncSet, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.SelectorSyncSet, error)
	List(opts metav1.ListOptions) (*v1.SelectorSyncSetList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.SelectorSyncSet, err error)
	SelectorSyncSetExpansion
}

// selectorSyncSets implements SelectorSyncSetInterface
type selectorSyncSets struct {
	client rest.Interface
}

// newSelectorSyncSets returns a SelectorSyncSets
func newSelectorSyncSets(c *HiveV1Client) *selectorSyncSets {
	return &selectorSyncSets{
		client: c.RESTClient(),
	}
}

// Get takes name of the selectorSyncSet, and returns the corresponding selectorSyncSet object, and an error if there is any.
func (c *selectorSyncSets) Get(name string, options metav1.GetOptions) (result *v1.SelectorSyncSet, err error) {
	result = &v1.SelectorSyncSet{}
	err = c.client.Get().
		Resource("selectorsyncsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SelectorSyncSets that match those selectors.
func (c *selectorSyncSets) List(opts metav1.ListOptions) (result *v1.SelectorSyncSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SelectorSyncSetList{}
	err = c.client.Get().
		Resource("selectorsyncsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested selectorSyncSets.
func (c *selectorSyncSets) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("selectorsyncsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a selectorSyncSet and creates it.  Returns the server's representation of the selectorSyncSet, and an error, if there is any.
func (c *selectorSyncSets) Create(selectorSyncSet *v1.SelectorSyncSet) (result *v1.SelectorSyncSet, err error) {
	result = &v1.SelectorSyncSet{}
	err = c.client.Post().
		Resource("selectorsyncsets").
		Body(selectorSyncSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a selectorSyncSet and updates it. Returns the server's representation of the selectorSyncSet, and an error, if there is any.
func (c *selectorSyncSets) Update(selectorSyncSet *v1.SelectorSyncSet) (result *v1.SelectorSyncSet, err error) {
	result = &v1.SelectorSyncSet{}
	err = c.client.Put().
		Resource("selectorsyncsets").
		Name(selectorSyncSet.Name).
		Body(selectorSyncSet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *selectorSyncSets) UpdateStatus(selectorSyncSet *v1.SelectorSyncSet) (result *v1.SelectorSyncSet, err error) {
	result = &v1.SelectorSyncSet{}
	err = c.client.Put().
		Resource("selectorsyncsets").
		Name(selectorSyncSet.Name).
		SubResource("status").
		Body(selectorSyncSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the selectorSyncSet and deletes it. Returns an error if one occurs.
func (c *selectorSyncSets) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("selectorsyncsets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *selectorSyncSets) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("selectorsyncsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched selectorSyncSet.
func (c *selectorSyncSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.SelectorSyncSet, err error) {
	result = &v1.SelectorSyncSet{}
	err = c.client.Patch(pt).
		Resource("selectorsyncsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
