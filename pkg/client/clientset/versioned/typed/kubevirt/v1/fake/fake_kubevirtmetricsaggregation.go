/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	kubevirtv1 "github.com/MarSik/kubevirt-ssp-operator/pkg/apis/kubevirt/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubevirtMetricsAggregations implements KubevirtMetricsAggregationInterface
type FakeKubevirtMetricsAggregations struct {
	Fake *FakeKubevirtV1
	ns   string
}

var kubevirtmetricsaggregationsResource = schema.GroupVersionResource{Group: "kubevirt.io", Version: "v1", Resource: "kubevirtmetricsaggregations"}

var kubevirtmetricsaggregationsKind = schema.GroupVersionKind{Group: "kubevirt.io", Version: "v1", Kind: "KubevirtMetricsAggregation"}

// Get takes name of the kubevirtMetricsAggregation, and returns the corresponding kubevirtMetricsAggregation object, and an error if there is any.
func (c *FakeKubevirtMetricsAggregations) Get(name string, options v1.GetOptions) (result *kubevirtv1.KubevirtMetricsAggregation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kubevirtmetricsaggregationsResource, c.ns, name), &kubevirtv1.KubevirtMetricsAggregation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubevirtv1.KubevirtMetricsAggregation), err
}

// List takes label and field selectors, and returns the list of KubevirtMetricsAggregations that match those selectors.
func (c *FakeKubevirtMetricsAggregations) List(opts v1.ListOptions) (result *kubevirtv1.KubevirtMetricsAggregationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kubevirtmetricsaggregationsResource, kubevirtmetricsaggregationsKind, c.ns, opts), &kubevirtv1.KubevirtMetricsAggregationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kubevirtv1.KubevirtMetricsAggregationList{ListMeta: obj.(*kubevirtv1.KubevirtMetricsAggregationList).ListMeta}
	for _, item := range obj.(*kubevirtv1.KubevirtMetricsAggregationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubevirtMetricsAggregations.
func (c *FakeKubevirtMetricsAggregations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kubevirtmetricsaggregationsResource, c.ns, opts))

}

// Create takes the representation of a kubevirtMetricsAggregation and creates it.  Returns the server's representation of the kubevirtMetricsAggregation, and an error, if there is any.
func (c *FakeKubevirtMetricsAggregations) Create(kubevirtMetricsAggregation *kubevirtv1.KubevirtMetricsAggregation) (result *kubevirtv1.KubevirtMetricsAggregation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kubevirtmetricsaggregationsResource, c.ns, kubevirtMetricsAggregation), &kubevirtv1.KubevirtMetricsAggregation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubevirtv1.KubevirtMetricsAggregation), err
}

// Update takes the representation of a kubevirtMetricsAggregation and updates it. Returns the server's representation of the kubevirtMetricsAggregation, and an error, if there is any.
func (c *FakeKubevirtMetricsAggregations) Update(kubevirtMetricsAggregation *kubevirtv1.KubevirtMetricsAggregation) (result *kubevirtv1.KubevirtMetricsAggregation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kubevirtmetricsaggregationsResource, c.ns, kubevirtMetricsAggregation), &kubevirtv1.KubevirtMetricsAggregation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubevirtv1.KubevirtMetricsAggregation), err
}

// Delete takes name of the kubevirtMetricsAggregation and deletes it. Returns an error if one occurs.
func (c *FakeKubevirtMetricsAggregations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kubevirtmetricsaggregationsResource, c.ns, name), &kubevirtv1.KubevirtMetricsAggregation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubevirtMetricsAggregations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kubevirtmetricsaggregationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &kubevirtv1.KubevirtMetricsAggregationList{})
	return err
}

// Patch applies the patch and returns the patched kubevirtMetricsAggregation.
func (c *FakeKubevirtMetricsAggregations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *kubevirtv1.KubevirtMetricsAggregation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kubevirtmetricsaggregationsResource, c.ns, name, data, subresources...), &kubevirtv1.KubevirtMetricsAggregation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubevirtv1.KubevirtMetricsAggregation), err
}
