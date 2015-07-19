package ctrl

import (
	. "github.com/ghilbut/ygg.go/common"
	"github.com/ghilbut/ygg.go/debug"
	"log"
)

type CtrlManager struct {
	ConnecteeDelegate
	CtrlReadyDelegate
	TargetReadyDelegate

	connectee   Connectee
	connector   Connector
	ctrlReady   *CtrlReady
	targetReady *TargetReady
	adapters    map[Adapter]bool
}

func NewCtrlManager(connector Connector) *CtrlManager {
	log.Println("======== [CtrlManager][NewCtrlManager] ========")
	assert.True(connector != nil)

	manager := &CtrlManager{
		connectee:   nil,
		connector:   connector,
		ctrlReady:   NewCtrlReady(),
		targetReady: NewTargetReady(),
		adapters:    make(map[Adapter]bool),
	}

	manager.ctrlReady.Delegate = manager
	manager.targetReady.Delegate = manager
	return manager
}

func (self *CtrlManager) OnConnecteeStarted(connectee Connectee) {
	log.Println("======== [CtrlManager][OnConnecteeStarted] ========")
	assert.True(connectee != nil)

	self.connectee = connectee
}

func (self *CtrlManager) OnConnecteeStopped() {
	log.Println("======== [CtrlManager][OnConnecteeStopped] ========")
	assert.True(self.connectee != nil)

	self.ctrlReady.Clear()
	self.targetReady.Clear()

	for adapter, _ := range self.adapters {
		adapter.Close()
	}

	assert.True(len(self.adapters) == 0)

	self.connectee = nil

}

func (self *CtrlManager) HasAdapter(adapter Adapter) bool {
	log.Println("======== [CtrlManager][HasAdapter] ========")
	assert.True(adapter != nil)

	_, ok := self.adapters[adapter]
	return ok
}

func (self *CtrlManager) OnCtrlConnected(conn Connection) {
	log.Println("======== [CtrlManager][OnCtrlConnected] ========")

	assert.True(conn != nil)
	self.ctrlReady.SetConnection(conn)
}

func (self *CtrlManager) OnCtrlProxy(proxy *CtrlProxy) {
	log.Println("======== [CtrlManager][OnCtrlProxy] ========")

	conn := self.connector.Connect(proxy)
	if conn != nil {
		self.targetReady.SetConnection(conn, proxy)
	} else {
		proxy.Close()
	}
}

func (self *CtrlManager) OnTargetProxy(ctrl *CtrlProxy, target *TargetProxy) {
	log.Println("======== [CtrlManager][OnTargetProxy] ========")

	adapter := NewOneToOneAdapter(target)
	if adapter == nil {
		ctrl.Close()
		target.Close()
		return
	}

	adapter.SetCtrlProxy(ctrl)

	self.adapters[adapter] = true
	adapter.BindDelegate(self)
}

func (self *CtrlManager) OnAdapterClosed(adapter Adapter) {
	log.Println("======== [CtrlManager][OnAdapterClosed] ========")
	assert.True(self.HasAdapter(adapter))

	adapter.UnbindDelegate()
	delete(self.adapters, adapter)
}
