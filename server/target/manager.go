package target

import (
	. "github.com/ghilbut/ygg.go/common"
	"log"
)

type Manager struct {
	ConnectionDelegate

	ctrls    map[Connection]bool
	targets  map[Connection]bool
	adapters map[string]Adapter
}

func NewManager() *Manager {

	manager := &Manager{
		ctrls:    make(map[Connection]bool),
		targets:  make(map[Connection]bool),
		adapters: make(map[string]Adapter),
	}

	return manager
}

func (self *Manager) SetCtrlConnection(conn Connection) {

	ctrls := self.ctrls

	if _, ok := ctrls[conn]; ok {
		panic("[target.Manager] can not set ctrl connection again which already exists.")
	}

	ctrls[conn] = true
	conn.BindDelegate(self)
}

func (self *Manager) SetTargetConnection(conn Connection) {

	targets := self.targets

	if _, ok := targets[conn]; ok {
		panic("[target.Manager] can not set target connection again which already exists.")
	}

	targets[conn] = true
	conn.BindDelegate(self)
}

func (self *Manager) setCtrlProxy(conn Connection, text string) {

	ctrls := self.ctrls
	adapters := self.adapters

	defer func() {
		delete(ctrls, conn)
	}()

	desc, _ := NewCtrlDesc(text)
	if desc == nil {
		conn.Close()
		return
	}

	proxy := NewCtrlProxy(conn, desc)
	if proxy == nil {
		conn.Close()
		return
	}

	adapter, ok := adapters[desc.Endpoint]
	if !ok {
		proxy.Close()
		return
	}

	if !adapter.SetCtrlProxy(proxy) {
		proxy.Close()
	}
}

func (self *Manager) setTargetAdapter(conn Connection, text string) {

	targets := self.targets
	adapters := self.adapters

	defer func() {
		delete(targets, conn)
	}()

	desc, _ := NewTargetDesc(text)
	if desc == nil {
		conn.Close()
		return
	}

	proxy := NewTargetProxy(conn, desc)
	if proxy == nil {
		conn.Close()
		return
	}

	adapter := NewManyToOneAdapter(proxy)
	if adapter == nil {
		proxy.Close()
		return
	}

	endpoint := desc.Endpoint

	if _, ok := adapters[endpoint]; ok {
		//adapter.Close()
		return
	}

	adapters[desc.Endpoint] = adapter
}

func (self *Manager) OnText(conn Connection, text string) {
	log.Println("[target.Manager]", "[OnText]", text)

	if _, ok := self.ctrls[conn]; ok {
		log.Println("[target.Manager]", "[OnText]", "setCtrlProxy(...)")

		self.setCtrlProxy(conn, text)
		return
	}

	if _, ok := self.targets[conn]; ok {
		log.Println("[target.Manager]", "[OnText]", "setTargetAdapter(...)")

		self.setTargetAdapter(conn, text)
		return
	}

	panic("[target.Manager][OnText] connection is not exists.")
}

func (self *Manager) OnBinary(conn Connection, bytes []byte) {
	panic("[target.Manager][OnBinary] not used on this step.")
}

func (self *Manager) OnClosed(conn Connection) {
	log.Println("[target.Manager]", "[OnClose]")

	ctrls := self.ctrls
	if _, ok := ctrls[conn]; ok {
		delete(ctrls, conn)
		return
	}

	targets := self.targets
	if _, ok := targets[conn]; ok {
		delete(targets, conn)
		return
	}
}
