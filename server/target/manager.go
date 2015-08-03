package target

import (
	. "github.com/ghilbut/ygg.go/common"
	"github.com/ghilbut/ygg.go/debug"
	"log"
)

type TargetManager struct {
	TargetBridgeDelegate
	CtrlReadyDelegate
	TargetReadyDelegate

	bridge            TargetBridge
	ctrlReady         *CtrlReady
	targetReady       *TargetReady
	endpointToAdapter map[string]Adapter
	adapterToEndpoint map[Adapter]string
}

func NewTargetManager() *TargetManager {
	log.Println("======== [TargetManager][NewTargetManager] ========")

	manager := &TargetManager{
		bridge:            nil,
		ctrlReady:         NewCtrlReady(),
		targetReady:       NewTargetReady(),
		endpointToAdapter: make(map[string]Adapter),
		adapterToEndpoint: make(map[Adapter]string),
	}

	manager.ctrlReady.Delegate = manager
	manager.targetReady.Delegate = manager
	return manager
}

func (self *TargetManager) OnTargetBridgeStarted(bridge TargetBridge) {
	log.Println("======== [TargetManager][OnTargetBridgeStarted] ========")
	assert.True(bridge != nil)

	self.bridge = bridge
}

func (self *TargetManager) OnTargetBridgeStopped() {
	log.Println("======== [TargetManager][OnTargetBridgeStoped] ========")

	self.ctrlReady.Clear()
	self.targetReady.Clear()

	for _, adapter := range self.endpointToAdapter {
		adapter.Close()
	}

	assert.True(len(self.adapterToEndpoint) == 0)

	self.bridge = nil
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

func (self *TargetManager) OnCtrlConnected(conn Connection) {
	log.Println("======== [TargetManager][OnCtrlConnected] ========")

	assert.True(conn != nil)
	self.ctrlReady.SetConnection(conn)
}

func (self *TargetManager) OnTargetConnected(conn Connection) {
	log.Println("======== [TargetManager][OnTargetConnected] ========")

	assert.True(conn != nil)
	self.targetReady.SetConnection(conn)
}

func (self *TargetManager) OnCtrlProxy(proxy *CtrlProxy) {
	log.Println("======== [TargetManager][OnCtrlProxy] ========")

	endpoint := proxy.Desc.Endpoint
	adapter, ok := self.endpointToAdapter[endpoint]
	if !ok {
		proxy.Close()
		return
	}

	adapter.SetCtrlProxy(proxy)
}

func (self *TargetManager) OnTargetProxy(proxy *TargetProxy) {
	log.Println("======== [TargetManager][OnTargetProxy] ========")

	endpoint := proxy.Desc.Endpoint

	if adapter, ok := self.endpointToAdapter[endpoint]; ok {
		adapter.Close()
	}

	adapter := NewManyToOneAdapter(proxy)
	if adapter == nil {
		proxy.Close()
		return
	}

	if !self.bridge.Register(endpoint) {
		proxy.Close()
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
	self.bridge.Unregister(endpoint)
	delete(self.adapterToEndpoint, adapter)
	delete(self.endpointToAdapter, endpoint)
}
