package pkg

import (
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/util/workqueue"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Kube struct {
	Client kubernetes.Interface
}

func InitializeClient(kubeconfig, master string) (*Kube, error) {
	conf, err := clientcmd.BuildConfigFromFlags(master, kubeconfig)
	if err != nil {
		return nil, err
	}
	client, err := kubernetes.NewForConfig(conf)
	if err != nil {
		return nil, err
	}
	return &Kube{Client:client}, nil
}

func (k *Kube) Run()  {
	podListWatcher := cache.NewListWatchFromClient(k.Client.CoreV1().RESTClient(), "pods", core.NamespaceDefault, fields.Everything())

	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	indexer, informer := cache.NewIndexerInformer(podListWatcher, &core.Pod{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(newObj)
			if err == nil {
				queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
	}, cache.Indexers{})

	controller := NewController(queue, indexer, informer)

	indexer.Add(&core.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "testpod",
			Namespace: core.NamespaceDefault,
		},
	})

	stop := make(chan struct{})
	defer close(stop)
	go controller.Run(1, stop)


	select {}


}