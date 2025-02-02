/*
Copyright 2021 Alibaba Cloud.

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
	v1alpha1 "github.com/AliyunContainerService/ai-models-on-ack/apis/scheduling.alibabacloud.com/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ResourcePolicyLister helps list ResourcePolicies.
// All objects returned here must be treated as read-only.
type ResourcePolicyLister interface {
	// List lists all ResourcePolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ResourcePolicy, err error)
	// ResourcePolicies returns an object that can list and get ResourcePolicies.
	ResourcePolicies(namespace string) ResourcePolicyNamespaceLister
	ResourcePolicyListerExpansion
}

// resourcePolicyLister implements the ResourcePolicyLister interface.
type resourcePolicyLister struct {
	listers.ResourceIndexer[*v1alpha1.ResourcePolicy]
}

// NewResourcePolicyLister returns a new ResourcePolicyLister.
func NewResourcePolicyLister(indexer cache.Indexer) ResourcePolicyLister {
	return &resourcePolicyLister{listers.New[*v1alpha1.ResourcePolicy](indexer, v1alpha1.Resource("resourcepolicy"))}
}

// ResourcePolicies returns an object that can list and get ResourcePolicies.
func (s *resourcePolicyLister) ResourcePolicies(namespace string) ResourcePolicyNamespaceLister {
	return resourcePolicyNamespaceLister{listers.NewNamespaced[*v1alpha1.ResourcePolicy](s.ResourceIndexer, namespace)}
}

// ResourcePolicyNamespaceLister helps list and get ResourcePolicies.
// All objects returned here must be treated as read-only.
type ResourcePolicyNamespaceLister interface {
	// List lists all ResourcePolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ResourcePolicy, err error)
	// Get retrieves the ResourcePolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ResourcePolicy, error)
	ResourcePolicyNamespaceListerExpansion
}

// resourcePolicyNamespaceLister implements the ResourcePolicyNamespaceLister
// interface.
type resourcePolicyNamespaceLister struct {
	listers.ResourceIndexer[*v1alpha1.ResourcePolicy]
}
