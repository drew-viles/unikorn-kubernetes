/*
Copyright 2022-2024 EscherCloud.

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

	v1alpha1 "github.com/eschercloudai/unikorn/pkg/apis/unikorn/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeControlPlaneApplicationBundles implements ControlPlaneApplicationBundleInterface
type FakeControlPlaneApplicationBundles struct {
	Fake *FakeUnikornV1alpha1
}

var controlplaneapplicationbundlesResource = v1alpha1.SchemeGroupVersion.WithResource("controlplaneapplicationbundles")

var controlplaneapplicationbundlesKind = v1alpha1.SchemeGroupVersion.WithKind("ControlPlaneApplicationBundle")

// Get takes name of the controlPlaneApplicationBundle, and returns the corresponding controlPlaneApplicationBundle object, and an error if there is any.
func (c *FakeControlPlaneApplicationBundles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ControlPlaneApplicationBundle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(controlplaneapplicationbundlesResource, name), &v1alpha1.ControlPlaneApplicationBundle{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ControlPlaneApplicationBundle), err
}

// List takes label and field selectors, and returns the list of ControlPlaneApplicationBundles that match those selectors.
func (c *FakeControlPlaneApplicationBundles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ControlPlaneApplicationBundleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(controlplaneapplicationbundlesResource, controlplaneapplicationbundlesKind, opts), &v1alpha1.ControlPlaneApplicationBundleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ControlPlaneApplicationBundleList{ListMeta: obj.(*v1alpha1.ControlPlaneApplicationBundleList).ListMeta}
	for _, item := range obj.(*v1alpha1.ControlPlaneApplicationBundleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested controlPlaneApplicationBundles.
func (c *FakeControlPlaneApplicationBundles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(controlplaneapplicationbundlesResource, opts))
}

// Create takes the representation of a controlPlaneApplicationBundle and creates it.  Returns the server's representation of the controlPlaneApplicationBundle, and an error, if there is any.
func (c *FakeControlPlaneApplicationBundles) Create(ctx context.Context, controlPlaneApplicationBundle *v1alpha1.ControlPlaneApplicationBundle, opts v1.CreateOptions) (result *v1alpha1.ControlPlaneApplicationBundle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(controlplaneapplicationbundlesResource, controlPlaneApplicationBundle), &v1alpha1.ControlPlaneApplicationBundle{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ControlPlaneApplicationBundle), err
}

// Update takes the representation of a controlPlaneApplicationBundle and updates it. Returns the server's representation of the controlPlaneApplicationBundle, and an error, if there is any.
func (c *FakeControlPlaneApplicationBundles) Update(ctx context.Context, controlPlaneApplicationBundle *v1alpha1.ControlPlaneApplicationBundle, opts v1.UpdateOptions) (result *v1alpha1.ControlPlaneApplicationBundle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(controlplaneapplicationbundlesResource, controlPlaneApplicationBundle), &v1alpha1.ControlPlaneApplicationBundle{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ControlPlaneApplicationBundle), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeControlPlaneApplicationBundles) UpdateStatus(ctx context.Context, controlPlaneApplicationBundle *v1alpha1.ControlPlaneApplicationBundle, opts v1.UpdateOptions) (*v1alpha1.ControlPlaneApplicationBundle, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(controlplaneapplicationbundlesResource, "status", controlPlaneApplicationBundle), &v1alpha1.ControlPlaneApplicationBundle{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ControlPlaneApplicationBundle), err
}

// Delete takes name of the controlPlaneApplicationBundle and deletes it. Returns an error if one occurs.
func (c *FakeControlPlaneApplicationBundles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(controlplaneapplicationbundlesResource, name, opts), &v1alpha1.ControlPlaneApplicationBundle{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeControlPlaneApplicationBundles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(controlplaneapplicationbundlesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ControlPlaneApplicationBundleList{})
	return err
}

// Patch applies the patch and returns the patched controlPlaneApplicationBundle.
func (c *FakeControlPlaneApplicationBundles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ControlPlaneApplicationBundle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(controlplaneapplicationbundlesResource, name, pt, data, subresources...), &v1alpha1.ControlPlaneApplicationBundle{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ControlPlaneApplicationBundle), err
}
