/*

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

package v1alpha1

import (
	v1alpha1 "github.com/oam-dev/oam-go-sdk/apis/core.oam.dev/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ApplicationScopeLister helps list ApplicationScopes.
type ApplicationScopeLister interface {
	// List lists all ApplicationScopes in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApplicationScope, err error)
	// ApplicationScopes returns an object that can list and get ApplicationScopes.
	ApplicationScopes(namespace string) ApplicationScopeNamespaceLister
	ApplicationScopeListerExpansion
}

// applicationScopeLister implements the ApplicationScopeLister interface.
type applicationScopeLister struct {
	indexer cache.Indexer
}

// NewApplicationScopeLister returns a new ApplicationScopeLister.
func NewApplicationScopeLister(indexer cache.Indexer) ApplicationScopeLister {
	return &applicationScopeLister{indexer: indexer}
}

// List lists all ApplicationScopes in the indexer.
func (s *applicationScopeLister) List(selector labels.Selector) (ret []*v1alpha1.ApplicationScope, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApplicationScope))
	})
	return ret, err
}

// ApplicationScopes returns an object that can list and get ApplicationScopes.
func (s *applicationScopeLister) ApplicationScopes(namespace string) ApplicationScopeNamespaceLister {
	return applicationScopeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApplicationScopeNamespaceLister helps list and get ApplicationScopes.
type ApplicationScopeNamespaceLister interface {
	// List lists all ApplicationScopes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApplicationScope, err error)
	// Get retrieves the ApplicationScope from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApplicationScope, error)
	ApplicationScopeNamespaceListerExpansion
}

// applicationScopeNamespaceLister implements the ApplicationScopeNamespaceLister
// interface.
type applicationScopeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApplicationScopes in the indexer for a given namespace.
func (s applicationScopeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApplicationScope, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApplicationScope))
	})
	return ret, err
}

// Get retrieves the ApplicationScope from the indexer for a given namespace and name.
func (s applicationScopeNamespaceLister) Get(name string) (*v1alpha1.ApplicationScope, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("applicationscope"), name)
	}
	return obj.(*v1alpha1.ApplicationScope), nil
}
