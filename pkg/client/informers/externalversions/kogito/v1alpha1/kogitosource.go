/*
Copyright 2021 The Knative Authors

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	kogitov1alpha1 "knative.dev/eventing-kogito/pkg/apis/kogito/v1alpha1"
	versioned "knative.dev/eventing-kogito/pkg/client/clientset/versioned"
	internalinterfaces "knative.dev/eventing-kogito/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "knative.dev/eventing-kogito/pkg/client/listers/kogito/v1alpha1"
)

// KogitoSourceInformer provides access to a shared informer and lister for
// KogitoSources.
type KogitoSourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.KogitoSourceLister
}

type kogitoSourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKogitoSourceInformer constructs a new informer for KogitoSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKogitoSourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKogitoSourceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKogitoSourceInformer constructs a new informer for KogitoSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKogitoSourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KogitoV1alpha1().KogitoSources(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KogitoV1alpha1().KogitoSources(namespace).Watch(context.TODO(), options)
			},
		},
		&kogitov1alpha1.KogitoSource{},
		resyncPeriod,
		indexers,
	)
}

func (f *kogitoSourceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKogitoSourceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kogitoSourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kogitov1alpha1.KogitoSource{}, f.defaultInformer)
}

func (f *kogitoSourceInformer) Lister() v1alpha1.KogitoSourceLister {
	return v1alpha1.NewKogitoSourceLister(f.Informer().GetIndexer())
}
