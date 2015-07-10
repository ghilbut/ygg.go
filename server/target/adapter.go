package target

import (
	. "github.com/ghilbut/ygg.go/common"
)

type ManyToOneAdapter struct {
	Adapter
	CtrlProxyDelegate
	TargetProxyDelegate

	ctrls  map[Proxy]bool
	target *TargetProxy
}

func NewManyToOneAdapter(proxy *TargetProxy) *ManyToOneAdapter {

	adapter := &ManyToOneAdapter{
		ctrls:  make(map[Proxy]bool),
		target: proxy,
	}
	proxy.Delegate = adapter

	return adapter
}

func (self *ManyToOneAdapter) SetCtrlProxy(proxy *CtrlProxy) bool {

	ctrls := self.ctrls

	if _, ok := ctrls[proxy]; ok {
		return false
	}

	ctrls[proxy] = true
	proxy.Delegate = self
	return true
}

func (self *ManyToOneAdapter) HasCtrlProxy(proxy *CtrlProxy) bool {
	_, ok := self.ctrls[proxy]
	return ok
}

func (self *ManyToOneAdapter) OnCtrlText(proxy *CtrlProxy, text string) {
	self.target.SendText(text)
}

func (self *ManyToOneAdapter) OnCtrlBinary(proxy *CtrlProxy, bytes []byte) {
	self.target.SendBinary(bytes)
}

func (self *ManyToOneAdapter) OnCtrlClosed(proxy *CtrlProxy) {
	if _, ok := self.ctrls[proxy]; ok {
		delete(self.ctrls, proxy)
	}
}

func (self *ManyToOneAdapter) OnTargetText(proxy *TargetProxy, text string) {
	for ctrl, _ := range self.ctrls {
		ctrl.SendText(text)
	}
}

func (self *ManyToOneAdapter) OnTargetBinary(proxy *TargetProxy, bytes []byte) {
	for ctrl, _ := range self.ctrls {
		ctrl.SendBinary(bytes)
	}
}

func (self *ManyToOneAdapter) OnTargetClosed(proxy *TargetProxy) {
	for ctrl, _ := range self.ctrls {
		ctrl.Close()
	}
}
