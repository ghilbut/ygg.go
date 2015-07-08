package ctrl

import (
	"github.com/ghilbut/ygg.go/common/ctrl"
	"github.com/ghilbut/ygg.go/common/target"
)

type Adapter interface {
	SetTarget(proxy *target.Proxy)
	OnCtrlText(ctrl *ctrl.Proxy, text string)
	OnCtrlBinary(ctrl *ctrl.Proxy, bytes []byte)
	OnCtrlClosed(ctrl *ctrl.Proxy)
	OnTargetText(target *target.Proxy, text string)
	OnTargetBinary(target *target.Proxy, bytes []byte)
	OnTargetClosed(target *target.Proxy)
}

type _CtrlBinder struct {
	ctrl.ProxyDelegate

	proxy    *ctrl.Proxy
	Delegate Adapter
}

func (self *_CtrlBinder) SendText(text string) {
	self.proxy.SendText(text)
}

func (self *_CtrlBinder) SendBinary(bytes []byte) {
	self.proxy.SendBinary(bytes)
}

func (self *_CtrlBinder) Close() {
	self.proxy.Close()
}

func (self *_CtrlBinder) OnText(proxy *ctrl.Proxy, text string) {
	self.Delegate.OnCtrlText(proxy, text)
}

func (self *_CtrlBinder) OnBinary(proxy *ctrl.Proxy, bytes []byte) {
	self.Delegate.OnCtrlBinary(proxy, bytes)
}

func (self *_CtrlBinder) OnClosed(proxy *ctrl.Proxy) {
	self.Delegate.OnCtrlClosed(proxy)
}

type _TargetBinder struct {
	target.ProxyDelegate

	proxy    *target.Proxy
	Delegate Adapter
}

func (self *_TargetBinder) SendText(text string) {
	self.proxy.SendText(text)
}

func (self *_TargetBinder) SendBinary(bytes []byte) {
	self.proxy.SendBinary(bytes)
}

func (self *_TargetBinder) Close() {
	self.proxy.Close()
}

func (self *_TargetBinder) OnText(proxy *target.Proxy, text string) {
	self.Delegate.OnTargetText(proxy, text)
}

func (self *_TargetBinder) OnBinary(proxy *target.Proxy, bytes []byte) {
	self.Delegate.OnTargetBinary(proxy, bytes)
}

func (self *_TargetBinder) OnClosed(proxy *target.Proxy) {
	self.Delegate.OnTargetClosed(proxy)
}

type BypassAdapter struct {
	Adapter

	cbinder *_CtrlBinder
	tbinder *_TargetBinder
}

func NewBypassAdapter(ctrl *ctrl.Proxy) Adapter {

	self := &BypassAdapter{
		cbinder: &_CtrlBinder{proxy: ctrl, Delegate: nil},
		tbinder: nil,
	}

	self.cbinder.Delegate = self

	return self
}

func (self *BypassAdapter) SetTarget(proxy *target.Proxy) {
	self.tbinder = &_TargetBinder{proxy: proxy, Delegate: nil}
	self.tbinder.Delegate = self
}

func (self *BypassAdapter) OnCtrlText(ctrl *ctrl.Proxy, text string) {
	if t := self.tbinder; t != nil {
		t.SendText(text)
	}
}

func (self *BypassAdapter) OnCtrlBinary(ctrl *ctrl.Proxy, bytes []byte) {
	if t := self.tbinder; t != nil {
		t.SendBinary(bytes)
	}

}

func (self *BypassAdapter) OnCtrlClosed(ctrl *ctrl.Proxy) {
	self.cbinder = nil
	if t := self.tbinder; t != nil {
		t.Close()
	}
}

func (self *BypassAdapter) OnTargetText(target *target.Proxy, text string) {
	if c := self.cbinder; c != nil {
		c.SendText(text)
	}

}

func (self *BypassAdapter) OnTargetBinary(target *target.Proxy, bytes []byte) {
	if c := self.cbinder; c != nil {
		c.SendBinary(bytes)
	}
}

func (self *BypassAdapter) OnTargetClosed(target *target.Proxy) {
	self.tbinder = nil
	if c := self.cbinder; c != nil {
		c.Close()
	}
}
