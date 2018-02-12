package pkg

import (
	"k8s.io/client-go/util/workqueue"
	"k8s.io/client-go/tools/cache"
	"k8s.io/apimachinery/pkg/util/runtime"
	"fmt"
	"k8s.io/apimachinery/pkg/util/wait"
	"time"
	core "k8s.io/api/core/v1"
)

type Controller struct {
	indexer cache.Indexer
	queue workqueue.RateLimitingInterface
	informer cache.Controller
}

func NewController(queue workqueue.RateLimitingInterface, indexer cache.Indexer, informer cache.Controller) *Controller  {
	return &Controller{
		indexer:indexer,
		queue:queue,
		informer:informer,
	}
}

func (c *Controller) Run(thrediness int, stopCh chan struct{})  {
	defer runtime.HandleCrash()

	defer c.queue.ShutDown()
	fmt.Println("Starting pod controller")
	go c.informer.Run(stopCh)

	if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("waiting timeout for caches sync"))
		return
	}

	for i := 0; i<thrediness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	<-stopCh
	fmt.Println("Stopping pod controller")
}

func (c *Controller) runWorker()  {
	for c.processNextItem() {}
}

func (c *Controller) processNextItem() bool {

	key, shutdown := c.queue.Get()
	if shutdown {
		return false
	}
	defer c.queue.Done(key)

	err := c.syncToStdout(key.(string))

	c.handleErr(err, key)
	return true

}

func (c *Controller) syncToStdout(key string) error  {
	obj, exists, err := c.indexer.GetByKey(key)
	if err != nil {
		fmt.Println(fmt.Sprintf("key %v is not found on store. err = %v", key, err))
		return err
	}
	if !exists {
		fmt.Println(fmt.Sprintf("pod %s not exists", key))
	}else {
		fmt.Println(fmt.Sprintf("sync/add/update pod %v", obj.(*core.Pod).GetName()))
	}
	return nil

}

func (c *Controller) handleErr(err error, key interface{})  {

}