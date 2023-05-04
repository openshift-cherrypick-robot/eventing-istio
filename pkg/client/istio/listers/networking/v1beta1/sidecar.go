/*
Copyright 2023 The Knative Authors

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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SidecarLister helps list Sidecars.
// All objects returned here must be treated as read-only.
type SidecarLister interface {
	// List lists all Sidecars in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Sidecar, err error)
	// Sidecars returns an object that can list and get Sidecars.
	Sidecars(namespace string) SidecarNamespaceLister
	SidecarListerExpansion
}

// sidecarLister implements the SidecarLister interface.
type sidecarLister struct {
	indexer cache.Indexer
}

// NewSidecarLister returns a new SidecarLister.
func NewSidecarLister(indexer cache.Indexer) SidecarLister {
	return &sidecarLister{indexer: indexer}
}

// List lists all Sidecars in the indexer.
func (s *sidecarLister) List(selector labels.Selector) (ret []*v1beta1.Sidecar, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Sidecar))
	})
	return ret, err
}

// Sidecars returns an object that can list and get Sidecars.
func (s *sidecarLister) Sidecars(namespace string) SidecarNamespaceLister {
	return sidecarNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SidecarNamespaceLister helps list and get Sidecars.
// All objects returned here must be treated as read-only.
type SidecarNamespaceLister interface {
	// List lists all Sidecars in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Sidecar, err error)
	// Get retrieves the Sidecar from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.Sidecar, error)
	SidecarNamespaceListerExpansion
}

// sidecarNamespaceLister implements the SidecarNamespaceLister
// interface.
type sidecarNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Sidecars in the indexer for a given namespace.
func (s sidecarNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.Sidecar, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Sidecar))
	})
	return ret, err
}

// Get retrieves the Sidecar from the indexer for a given namespace and name.
func (s sidecarNamespaceLister) Get(name string) (*v1beta1.Sidecar, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("sidecar"), name)
	}
	return obj.(*v1beta1.Sidecar), nil
}
