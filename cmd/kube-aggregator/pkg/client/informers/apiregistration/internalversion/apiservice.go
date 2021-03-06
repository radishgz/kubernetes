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

// This file was automatically generated by informer-gen with arguments: --input-dirs=[k8s.io/kubernetes/cmd/kube-aggregator/pkg/apis/apiregistration,k8s.io/kubernetes/cmd/kube-aggregator/pkg/apis/apiregistration/v1alpha1] --internal-clientset-package=k8s.io/kubernetes/cmd/kube-aggregator/pkg/client/clientset_generated/internalclientset --listers-package=k8s.io/kubernetes/cmd/kube-aggregator/pkg/client/listers --output-package=k8s.io/kubernetes/cmd/kube-aggregator/pkg/client/informers --versioned-clientset-package=k8s.io/kubernetes/cmd/kube-aggregator/pkg/client/clientset_generated/clientset

package internalversion

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	apiregistration "k8s.io/kubernetes/cmd/kube-aggregator/pkg/apis/apiregistration"
	internalclientset "k8s.io/kubernetes/cmd/kube-aggregator/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "k8s.io/kubernetes/cmd/kube-aggregator/pkg/client/informers/internalinterfaces"
	internalversion "k8s.io/kubernetes/cmd/kube-aggregator/pkg/client/listers/apiregistration/internalversion"
	cache "k8s.io/kubernetes/pkg/client/cache"
	time "time"
)

// APIServiceInformer provides access to a shared informer and lister for
// APIServices.
type APIServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.APIServiceLister
}

type aPIServiceInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newAPIServiceInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.Apiregistration().APIServices().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.Apiregistration().APIServices().Watch(options)
			},
		},
		&apiregistration.APIService{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *aPIServiceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InternalInformerFor(&apiregistration.APIService{}, newAPIServiceInformer)
}

func (f *aPIServiceInformer) Lister() internalversion.APIServiceLister {
	return internalversion.NewAPIServiceLister(f.Informer().GetIndexer())
}
