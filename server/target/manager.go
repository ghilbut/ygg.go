package target

import (
	. "github.com/ghilbut/ygg.go/common"
	. "github.com/ghilbut/ygg.go/debug"
	"log"
)

type Manager struct {
	ctrlReady   *CtrlReady
	targetReady *TargetReady
	adapters    map[string]Adapter
}

func NewManager() *Manager {

	manager := &Manager{
		ctrlReady:   NewCtrlReady(),
		targetReady: NewTargetReady(),
		adapters:    make(map[string]Adapter),
	}

	manager.ctrlReady.OnCtrlReadyProc = func(proxy *CtrlProxy) {
		log.Println("[OnCtrlReadyProc]")

		adapter, ok := manager.adapters[proxy.Desc.Endpoint]
		if !ok {
			proxy.Close()
			return
		}

		if !adapter.SetCtrlProxy(proxy) {
			proxy.Close()
		}
	}

	manager.targetReady.OnTargetReadyProc = func(proxy *TargetProxy) {
		log.Println("[OnTargetReadyProc")

		adapter := NewManyToOneAdapter(proxy)
		if adapter == nil {
			proxy.Close()
			return
		}

		endpoint := proxy.Desc.Endpoint

		if _, ok := manager.adapters[endpoint]; ok {
			//adapter.Close()
			return
		}

		manager.adapters[endpoint] = adapter
	}

	return manager
}

func (self *Manager) SetCtrlConnection(conn Connection) {
	Assert(conn != nil, "conn should not be nil.")
	self.ctrlReady.SetConnection(conn)
}

func (self *Manager) SetTargetConnection(conn Connection) {
	Assert(conn != nil, "conn should not be nil.")
	self.targetReady.SetConnection(conn)
}
