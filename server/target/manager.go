package target

import (
	. "github.com/ghilbut/ygg.go/common"
	"github.com/ghilbut/ygg.go/debug"
	"log"
)

type TargetManager struct {
	TargetReadyDelegate

	ctrlReady         *CtrlReady
	targetReady       *TargetReady
	endpointToAdapter map[string]Adapter
	adapterToEndpoint map[Adapter]string
}

func NewTargetManager() *TargetManager {

	manager := &TargetManager{
		ctrlReady:         NewCtrlReady(),
		targetReady:       NewTargetReady(),
		endpointToAdapter: make(map[string]Adapter),
		adapterToEndpoint: make(map[Adapter]string),
	}

	manager.targetReady.Delegate = manager

	manager.ctrlReady.OnCtrlReadyProc = func(proxy *CtrlProxy) {
		log.Println("======== [][OnCtrlReadyProc] ========")

		adapter, ok := manager.endpointToAdapter[proxy.Desc.Endpoint]
		if !ok {
			proxy.Close()
			return
		}

		adapter.SetCtrlProxy(proxy)
		proxy.Close()
	}

	return manager
}

func (self *TargetManager) SetCtrlConnection(conn Connection) {
	log.Println("======== [TargetManager][SetCtrlConnection] ========")

	assert.True(conn != nil)
	self.ctrlReady.SetConnection(conn)
}

func (self *TargetManager) SetTargetConnection(conn Connection) {
	log.Println("======== [TargetManager][SetTargetConnection] ========")

	assert.True(conn != nil)
	self.targetReady.SetConnection(conn)
}

func (self *TargetManager) HasAdapter(adapter Adapter) bool {
	log.Println("======== [TargetManager][HasAdapter] ========")
	assert.True(adapter != nil)

	endpoint, ok := self.adapterToEndpoint[adapter]
	_, check := self.endpointToAdapter[endpoint]
	assert.True(ok == check)
	return ok
}

func (self *TargetManager) HasEndpoint(endpoint string) bool {
	log.Println("======== [TargetManager][HasEndpoint] ========")

	adapter, ok := self.endpointToAdapter[endpoint]
	_, check := self.adapterToEndpoint[adapter]
	assert.True(ok == check)
	return ok
}

func (self *TargetManager) OnTargetProxy(proxy *TargetProxy) {
	log.Println("======== [TargetManager][OnTargetReadyProc] ========")

	adapter := NewManyToOneAdapter(proxy)
	if adapter == nil {
		proxy.Close()
		return
	}

	endpoint := proxy.Desc.Endpoint

	if _, ok := self.endpointToAdapter[endpoint]; ok {
		//adapter.Close()
		return
	}

	self.endpointToAdapter[endpoint] = adapter
	self.adapterToEndpoint[adapter] = endpoint
	adapter.BindDelegate(self)
}

func (self *TargetManager) OnAdapterClosed(adapter Adapter) {
	log.Println("======== [TargetManager][OnAdapterClosed] ========")
	assert.True(self.HasAdapter(adapter))

	adapter.UnbindDelegate()
	endpoint, _ := self.adapterToEndpoint[adapter]
	delete(self.adapterToEndpoint, adapter)
	delete(self.endpointToAdapter, endpoint)
}
