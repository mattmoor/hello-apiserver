/*
Copyright 2017 The Kubernetes Authors.

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
package fake

import (
	experimental "github.com/mattmoor/hello-apiserver/pkg/apis/experimental"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBuilds implements BuildInterface
type FakeBuilds struct {
	Fake *FakeExperimental
	ns   string
}

var buildsResource = schema.GroupVersionResource{Group: "experimental", Version: "", Resource: "builds"}

var buildsKind = schema.GroupVersionKind{Group: "experimental", Version: "", Kind: "Build"}

func (c *FakeBuilds) Create(build *experimental.Build) (result *experimental.Build, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(buildsResource, c.ns, build), &experimental.Build{})

	if obj == nil {
		return nil, err
	}
	return obj.(*experimental.Build), err
}

func (c *FakeBuilds) Update(build *experimental.Build) (result *experimental.Build, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(buildsResource, c.ns, build), &experimental.Build{})

	if obj == nil {
		return nil, err
	}
	return obj.(*experimental.Build), err
}

func (c *FakeBuilds) UpdateStatus(build *experimental.Build) (*experimental.Build, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(buildsResource, "status", c.ns, build), &experimental.Build{})

	if obj == nil {
		return nil, err
	}
	return obj.(*experimental.Build), err
}

func (c *FakeBuilds) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(buildsResource, c.ns, name), &experimental.Build{})

	return err
}

func (c *FakeBuilds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(buildsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &experimental.BuildList{})
	return err
}

func (c *FakeBuilds) Get(name string, options v1.GetOptions) (result *experimental.Build, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(buildsResource, c.ns, name), &experimental.Build{})

	if obj == nil {
		return nil, err
	}
	return obj.(*experimental.Build), err
}

func (c *FakeBuilds) List(opts v1.ListOptions) (result *experimental.BuildList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(buildsResource, buildsKind, c.ns, opts), &experimental.BuildList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &experimental.BuildList{}
	for _, item := range obj.(*experimental.BuildList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested builds.
func (c *FakeBuilds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(buildsResource, c.ns, opts))

}

// Patch applies the patch and returns the patched build.
func (c *FakeBuilds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *experimental.Build, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(buildsResource, c.ns, name, data, subresources...), &experimental.Build{})

	if obj == nil {
		return nil, err
	}
	return obj.(*experimental.Build), err
}
