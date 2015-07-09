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

func NewManyToOneAdapter(target *TargetProxy) *ManyToOneAdapter {

	adapter := &ManyToOneAdapter{
		ctrls:  make(map[Proxy]bool),
		target: target,
	}

	return adapter
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
