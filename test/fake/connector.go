package fake

import (
	. "github.com/ghilbut/ygg.go/common"
)

type FakeConnector struct {
	Connector
	ConnectorDelegate

	other    *FakeConnector
	delegate ConnectorDelegate
}

func NewFakeConnector() *FakeConnector {

	self := &FakeConnector{}
	other := &FakeConnector{}

	self.other = other
	other.other = self

	return self
}

func (self *FakeConnector) Other() *FakeConnector {
	return self.other
}

func (self *FakeConnector) BindDelegate(delegate ConnectorDelegate) {
	self.delegate = delegate
}

func (self *FakeConnector) UnbindDelegate() {
	self.delegate = nil
}

func (self *FakeConnector) SendText(text string) {
	o := self.other
	d := o.delegate
	if d != nil {
		d.OnText(text)
	}
}

func (self *FakeConnector) SendBinary(bytes []byte) {
	o := self.other
	d := o.delegate
	if d != nil {
		d.OnBinary(bytes)
	}
}

func (self *FakeConnector) Close() {
	o := self.other
	d := o.delegate
	if d != nil {
		d.OnClosed()
	}
}
