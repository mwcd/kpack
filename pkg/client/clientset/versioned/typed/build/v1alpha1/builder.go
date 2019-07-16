/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/pivotal/build-service-system/pkg/apis/build/v1alpha1"
	scheme "github.com/pivotal/build-service-system/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BuildersGetter has a method to return a BuilderInterface.
// A group's client should implement this interface.
type BuildersGetter interface {
	Builders(namespace string) BuilderInterface
}

// BuilderInterface has methods to work with Builder resources.
type BuilderInterface interface {
	Create(*v1alpha1.Builder) (*v1alpha1.Builder, error)
	Update(*v1alpha1.Builder) (*v1alpha1.Builder, error)
	UpdateStatus(*v1alpha1.Builder) (*v1alpha1.Builder, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Builder, error)
	List(opts v1.ListOptions) (*v1alpha1.BuilderList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Builder, err error)
	BuilderExpansion
}

// builders implements BuilderInterface
type builders struct {
	client rest.Interface
	ns     string
}

// newBuilders returns a Builders
func newBuilders(c *BuildV1alpha1Client, namespace string) *builders {
	return &builders{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the builder, and returns the corresponding builder object, and an error if there is any.
func (c *builders) Get(name string, options v1.GetOptions) (result *v1alpha1.Builder, err error) {
	result = &v1alpha1.Builder{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("builders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Builders that match those selectors.
func (c *builders) List(opts v1.ListOptions) (result *v1alpha1.BuilderList, err error) {
	result = &v1alpha1.BuilderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("builders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested builders.
func (c *builders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("builders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a builder and creates it.  Returns the server's representation of the builder, and an error, if there is any.
func (c *builders) Create(builder *v1alpha1.Builder) (result *v1alpha1.Builder, err error) {
	result = &v1alpha1.Builder{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("builders").
		Body(builder).
		Do().
		Into(result)
	return
}

// Update takes the representation of a builder and updates it. Returns the server's representation of the builder, and an error, if there is any.
func (c *builders) Update(builder *v1alpha1.Builder) (result *v1alpha1.Builder, err error) {
	result = &v1alpha1.Builder{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("builders").
		Name(builder.Name).
		Body(builder).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *builders) UpdateStatus(builder *v1alpha1.Builder) (result *v1alpha1.Builder, err error) {
	result = &v1alpha1.Builder{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("builders").
		Name(builder.Name).
		SubResource("status").
		Body(builder).
		Do().
		Into(result)
	return
}

// Delete takes name of the builder and deletes it. Returns an error if one occurs.
func (c *builders) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("builders").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *builders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("builders").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched builder.
func (c *builders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Builder, err error) {
	result = &v1alpha1.Builder{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("builders").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
