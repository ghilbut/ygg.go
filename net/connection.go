package net

import (
	. "github.com/ghilbut/ygg.go/common"
)

type LocalConnection struct {
	Connection
	ConnectionDelegate

	other    *LocalConnection
	delegate ConnectionDelegate
}

func NewLocalConnection() *LocalConnection {

	self := &LocalConnection{}
	other := &LocalConnection{}

	self.other = other
	other.other = self

	return self
}

func (self *LocalConnection) Other() *LocalConnection {
	return self.other
}

func (self *LocalConnection) BindDelegate(delegate ConnectionDelegate) {
	self.delegate = delegate
}

func (self *LocalConnection) UnbindDelegate() {
	self.delegate = nil
}

func (self *LocalConnection) SendText(text string) {
	o := self.other
	if d := o.delegate; d != nil {
		d.OnText(o, text)
	}
}

func (self *LocalConnection) SendBinary(bytes []byte) {
	o := self.other
	if d := o.delegate; d != nil {
		d.OnBinary(o, bytes)
	}
}

func (self *LocalConnection) Close() {
	if d := self.delegate; d != nil {
		d.OnClosed(self)
	}
	o := self.other
	if d := o.delegate; d != nil {
		d.OnClosed(o)
	}
}
