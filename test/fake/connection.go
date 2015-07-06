package fake

import (
	. "github.com/ghilbut/ygg.go/common"
)

type FakeConnection struct {
	Connection
	ConnectionDelegate

	other    *FakeConnection
	delegate ConnectionDelegate
}

func NewFakeConnection() *FakeConnection {

	self := &FakeConnection{}
	other := &FakeConnection{}

	self.other = other
	other.other = self

	return self
}

func (self *FakeConnection) Other() *FakeConnection {
	return self.other
}

func (self *FakeConnection) BindDelegate(delegate ConnectionDelegate) {
	self.delegate = delegate
}

func (self *FakeConnection) UnbindDelegate() {
	self.delegate = nil
}

func (self *FakeConnection) SendText(text string) {
	o := self.other
	if d := o.delegate; d != nil {
		d.OnText(o, text)
	}
}

func (self *FakeConnection) SendBinary(bytes []byte) {
	o := self.other
	if d := o.delegate; d != nil {
		d.OnBinary(o, bytes)
	}
}

func (self *FakeConnection) Close() {
	if d := self.delegate; d != nil {
		d.OnClosed(self)
	}
	o := self.other
	if d := o.delegate; d != nil {
		d.OnClosed(o)
	}
}
