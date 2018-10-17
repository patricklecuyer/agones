// Copyright 2018 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was autogenerated. Do not edit directly.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "agones.dev/agones/pkg/apis/stable/v1alpha1"
	scheme "agones.dev/agones/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GameServersGetter has a method to return a GameServerInterface.
// A group's client should implement this interface.
type GameServersGetter interface {
	GameServers(namespace string) GameServerInterface
}

// GameServerInterface has methods to work with GameServer resources.
type GameServerInterface interface {
	Create(*v1alpha1.GameServer) (*v1alpha1.GameServer, error)
	Update(*v1alpha1.GameServer) (*v1alpha1.GameServer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GameServer, error)
	List(opts v1.ListOptions) (*v1alpha1.GameServerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GameServer, err error)
	GameServerExpansion
}

// gameServers implements GameServerInterface
type gameServers struct {
	client rest.Interface
	ns     string
}

// newGameServers returns a GameServers
func newGameServers(c *StableV1alpha1Client, namespace string) *gameServers {
	return &gameServers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gameServer, and returns the corresponding gameServer object, and an error if there is any.
func (c *gameServers) Get(name string, options v1.GetOptions) (result *v1alpha1.GameServer, err error) {
	result = &v1alpha1.GameServer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gameservers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GameServers that match those selectors.
func (c *gameServers) List(opts v1.ListOptions) (result *v1alpha1.GameServerList, err error) {
	result = &v1alpha1.GameServerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gameservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gameServers.
func (c *gameServers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gameservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a gameServer and creates it.  Returns the server's representation of the gameServer, and an error, if there is any.
func (c *gameServers) Create(gameServer *v1alpha1.GameServer) (result *v1alpha1.GameServer, err error) {
	result = &v1alpha1.GameServer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gameservers").
		Body(gameServer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a gameServer and updates it. Returns the server's representation of the gameServer, and an error, if there is any.
func (c *gameServers) Update(gameServer *v1alpha1.GameServer) (result *v1alpha1.GameServer, err error) {
	result = &v1alpha1.GameServer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gameservers").
		Name(gameServer.Name).
		Body(gameServer).
		Do().
		Into(result)
	return
}

// Delete takes name of the gameServer and deletes it. Returns an error if one occurs.
func (c *gameServers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gameservers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gameServers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gameservers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched gameServer.
func (c *gameServers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GameServer, err error) {
	result = &v1alpha1.GameServer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gameservers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}