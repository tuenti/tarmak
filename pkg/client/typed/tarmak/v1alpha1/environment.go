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

package v1alpha1

import (
	v1alpha1 "github.com/jetstack/tarmak/pkg/apis/tarmak/v1alpha1"
	scheme "github.com/jetstack/tarmak/pkg/client/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EnvironmentsGetter has a method to return a EnvironmentInterface.
// A group's client should implement this interface.
type EnvironmentsGetter interface {
	Environments(namespace string) EnvironmentInterface
}

// EnvironmentInterface has methods to work with Environment resources.
type EnvironmentInterface interface {
	Create(*v1alpha1.Environment) (*v1alpha1.Environment, error)
	Update(*v1alpha1.Environment) (*v1alpha1.Environment, error)
	UpdateStatus(*v1alpha1.Environment) (*v1alpha1.Environment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Environment, error)
	List(opts v1.ListOptions) (*v1alpha1.EnvironmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Environment, err error)
	EnvironmentExpansion
}

// environments implements EnvironmentInterface
type environments struct {
	client rest.Interface
	ns     string
}

// newEnvironments returns a Environments
func newEnvironments(c *TarmakV1alpha1Client, namespace string) *environments {
	return &environments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a environment and creates it.  Returns the server's representation of the environment, and an error, if there is any.
func (c *environments) Create(environment *v1alpha1.Environment) (result *v1alpha1.Environment, err error) {
	result = &v1alpha1.Environment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("environments").
		Body(environment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a environment and updates it. Returns the server's representation of the environment, and an error, if there is any.
func (c *environments) Update(environment *v1alpha1.Environment) (result *v1alpha1.Environment, err error) {
	result = &v1alpha1.Environment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("environments").
		Name(environment.Name).
		Body(environment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclientstatus=false comment above the type to avoid generating UpdateStatus().

func (c *environments) UpdateStatus(environment *v1alpha1.Environment) (result *v1alpha1.Environment, err error) {
	result = &v1alpha1.Environment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("environments").
		Name(environment.Name).
		SubResource("status").
		Body(environment).
		Do().
		Into(result)
	return
}

// Delete takes name of the environment and deletes it. Returns an error if one occurs.
func (c *environments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("environments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *environments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("environments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the environment, and returns the corresponding environment object, and an error if there is any.
func (c *environments) Get(name string, options v1.GetOptions) (result *v1alpha1.Environment, err error) {
	result = &v1alpha1.Environment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("environments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Environments that match those selectors.
func (c *environments) List(opts v1.ListOptions) (result *v1alpha1.EnvironmentList, err error) {
	result = &v1alpha1.EnvironmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("environments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested environments.
func (c *environments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("environments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched environment.
func (c *environments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Environment, err error) {
	result = &v1alpha1.Environment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("environments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
