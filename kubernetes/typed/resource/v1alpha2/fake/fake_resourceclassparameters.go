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
	"context"
	json "encoding/json"
	"fmt"

	v1alpha2 "k8s.io/api/resource/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	resourcev1alpha2 "k8s.io/client-go/applyconfigurations/resource/v1alpha2"
	testing "k8s.io/client-go/testing"
)

// FakeResourceClassParameters implements ResourceClassParametersInterface
type FakeResourceClassParameters struct {
	Fake *FakeResourceV1alpha2
	ns   string
}

var resourceclassparametersResource = v1alpha2.SchemeGroupVersion.WithResource("resourceclassparameters")

var resourceclassparametersKind = v1alpha2.SchemeGroupVersion.WithKind("ResourceClassParameters")

// Get takes name of the resourceClassParameters, and returns the corresponding resourceClassParameters object, and an error if there is any.
func (c *FakeResourceClassParameters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.ResourceClassParameters, err error) {
	emptyResult := &v1alpha2.ResourceClassParameters{}
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resourceclassparametersResource, c.ns, name), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.ResourceClassParameters), err
}

// List takes label and field selectors, and returns the list of ResourceClassParameters that match those selectors.
func (c *FakeResourceClassParameters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.ResourceClassParametersList, err error) {
	emptyResult := &v1alpha2.ResourceClassParametersList{}
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resourceclassparametersResource, resourceclassparametersKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.ResourceClassParametersList{ListMeta: obj.(*v1alpha2.ResourceClassParametersList).ListMeta}
	for _, item := range obj.(*v1alpha2.ResourceClassParametersList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resourceClassParameters.
func (c *FakeResourceClassParameters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resourceclassparametersResource, c.ns, opts))

}

// Create takes the representation of a resourceClassParameters and creates it.  Returns the server's representation of the resourceClassParameters, and an error, if there is any.
func (c *FakeResourceClassParameters) Create(ctx context.Context, resourceClassParameters *v1alpha2.ResourceClassParameters, opts v1.CreateOptions) (result *v1alpha2.ResourceClassParameters, err error) {
	emptyResult := &v1alpha2.ResourceClassParameters{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resourceclassparametersResource, c.ns, resourceClassParameters), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.ResourceClassParameters), err
}

// Update takes the representation of a resourceClassParameters and updates it. Returns the server's representation of the resourceClassParameters, and an error, if there is any.
func (c *FakeResourceClassParameters) Update(ctx context.Context, resourceClassParameters *v1alpha2.ResourceClassParameters, opts v1.UpdateOptions) (result *v1alpha2.ResourceClassParameters, err error) {
	emptyResult := &v1alpha2.ResourceClassParameters{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resourceclassparametersResource, c.ns, resourceClassParameters), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.ResourceClassParameters), err
}

// Delete takes name of the resourceClassParameters and deletes it. Returns an error if one occurs.
func (c *FakeResourceClassParameters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(resourceclassparametersResource, c.ns, name, opts), &v1alpha2.ResourceClassParameters{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResourceClassParameters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resourceclassparametersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.ResourceClassParametersList{})
	return err
}

// Patch applies the patch and returns the patched resourceClassParameters.
func (c *FakeResourceClassParameters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.ResourceClassParameters, err error) {
	emptyResult := &v1alpha2.ResourceClassParameters{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resourceclassparametersResource, c.ns, name, pt, data, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.ResourceClassParameters), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied resourceClassParameters.
func (c *FakeResourceClassParameters) Apply(ctx context.Context, resourceClassParameters *resourcev1alpha2.ResourceClassParametersApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha2.ResourceClassParameters, err error) {
	if resourceClassParameters == nil {
		return nil, fmt.Errorf("resourceClassParameters provided to Apply must not be nil")
	}
	data, err := json.Marshal(resourceClassParameters)
	if err != nil {
		return nil, err
	}
	name := resourceClassParameters.Name
	if name == nil {
		return nil, fmt.Errorf("resourceClassParameters.Name must be provided to Apply")
	}
	emptyResult := &v1alpha2.ResourceClassParameters{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resourceclassparametersResource, c.ns, *name, types.ApplyPatchType, data), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.ResourceClassParameters), err
}
