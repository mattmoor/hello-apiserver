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

// This file was automatically generated by informer-gen

package internalversion

import (
	experimental "github.com/mattmoor/hello-apiserver/pkg/apis/experimental"
	internalclientset "github.com/mattmoor/hello-apiserver/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "github.com/mattmoor/hello-apiserver/pkg/client/informers_generated/internalversion/internalinterfaces"
	internalversion "github.com/mattmoor/hello-apiserver/pkg/client/listers_generated/experimental/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// BuildInformer provides access to a shared informer and lister for
// Builds.
type BuildInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.BuildLister
}

type buildInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newBuildInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.Experimental().Builds(v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.Experimental().Builds(v1.NamespaceAll).Watch(options)
			},
		},
		&experimental.Build{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *buildInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&experimental.Build{}, newBuildInformer)
}

func (f *buildInformer) Lister() internalversion.BuildLister {
	return internalversion.NewBuildLister(f.Informer().GetIndexer())
}
