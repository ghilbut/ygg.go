package common

import (
	. "github.com/ghilbut/ygg.go/debug"
)

type Proxy interface {
	SendText(text string)
	SendBinary(bytes []byte)
	Close()
}

type CtrlProxy struct {
	Proxy
	ConnectionDelegate

	conn     Connection
	Desc     *CtrlDesc
	Delegate CtrlProxyDelegate
}

type CtrlProxyDelegate interface {
	OnCtrlText(proxy *CtrlProxy, text string)
	OnCtrlBinary(proxy *CtrlProxy, bytes []byte)
	OnCtrlClosed(proxy *CtrlProxy)
}

func NewCtrlProxy(conn Connection, desc *CtrlDesc) *CtrlProxy {
	Assert(conn != nil, "connection is nil.")
	Assert(desc != nil, "desc is nil.")

	proxy := &CtrlProxy{conn: conn, Desc: desc}
	conn.BindDelegate(proxy)

	return proxy
}

func (self *CtrlProxy) SendText(text string) {
	self.conn.SendText(text)
}

func (self *CtrlProxy) SendBinary(bytes []byte) {
	self.conn.SendBinary(bytes)
}

func (self *CtrlProxy) Close() {
	self.conn.Close()
}

func (self *CtrlProxy) OnText(conn Connection, text string) {
	Assert(conn == self.conn, "conn is invalid connection.")

	if d := self.Delegate; d != nil {
		d.OnCtrlText(self, text)
	}
}

func (self *CtrlProxy) OnBinary(conn Connection, bytes []byte) {
	Assert(conn == self.conn, "conn is invalid connection.")

	if d := self.Delegate; d != nil {
		d.OnCtrlBinary(self, bytes)
	}
}

func (self *CtrlProxy) OnClosed(conn Connection) {
	Assert(conn == self.conn, "conn is invalid connection.")

	if d := self.Delegate; d != nil {
		d.OnCtrlClosed(self)
	}
}

type TargetProxy struct {
	Proxy
	ConnectionDelegate

	conn     Connection
	Desc     *TargetDesc
	Delegate TargetProxyDelegate
}

type TargetProxyDelegate interface {
	OnTargetText(proxy *TargetProxy, text string)
	OnTargetBinary(proxy *TargetProxy, bytes []byte)
	OnTargetClosed(proxy *TargetProxy)
}

func NewTargetProxy(conn Connection, desc *TargetDesc) *TargetProxy {
	Assert(conn != nil, "connection is nil.")
	Assert(desc != nil, "desc is nil.")

	proxy := &TargetProxy{conn: conn, Desc: desc}
	conn.BindDelegate(proxy)

	return proxy
}

func (self *TargetProxy) SendText(text string) {
	self.conn.SendText(text)
}

func (self *TargetProxy) SendBinary(bytes []byte) {
	self.conn.SendBinary(bytes)
}

func (self *TargetProxy) Close() {
	self.conn.Close()
}

func (self *TargetProxy) OnText(conn Connection, text string) {
	Assert(conn == self.conn, "conn is invalid connection.")

	if d := self.Delegate; d != nil {
		d.OnTargetText(self, text)
	}
}

func (self *TargetProxy) OnBinary(conn Connection, bytes []byte) {
	Assert(conn == self.conn, "conn is invalid connection.")

	if d := self.Delegate; d != nil {
		d.OnTargetBinary(self, bytes)
	}
}

func (self *TargetProxy) OnClosed(conn Connection) {
	Assert(conn == self.conn, "conn is invalid connection.")

	if d := self.Delegate; d != nil {
		d.OnTargetClosed(self)
	}
}
