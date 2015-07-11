package common

import (
	"github.com/ghilbut/ygg.go/debug"
	"log"
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
	log.Println("======== [CtrlProxy][NewCtrlProxy] ========")
	assert.True(conn != nil)
	assert.True(desc != nil)

	proxy := &CtrlProxy{conn: conn, Desc: desc}
	conn.BindDelegate(proxy)

	return proxy
}

func (self *CtrlProxy) SendText(text string) {
	log.Println("======== [CtrlProxy][SendText] ========")

	self.conn.SendText(text)
}

func (self *CtrlProxy) SendBinary(bytes []byte) {
	log.Println("======== [CtrlProxy][SendBinary] ========")

	self.conn.SendBinary(bytes)
}

func (self *CtrlProxy) Close() {
	log.Println("======== [CtrlProxy][Close] ========")

	self.conn.Close()
}

func (self *CtrlProxy) OnText(conn Connection, text string) {
	log.Println("======== [CtrlProxy][OnText] ========")
	assert.True(conn == self.conn)

	if d := self.Delegate; d != nil {
		d.OnCtrlText(self, text)
	}
}

func (self *CtrlProxy) OnBinary(conn Connection, bytes []byte) {
	log.Println("======== [CtrlProxy][OnBinary] ========")
	assert.True(conn == self.conn)

	if d := self.Delegate; d != nil {
		d.OnCtrlBinary(self, bytes)
	}
}

func (self *CtrlProxy) OnClosed(conn Connection) {
	log.Println("======== [CtrlProxy][OnClosed] ========")
	assert.True(conn == self.conn)

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
	log.Println("======== [TargetProxy][NewTargetProxy] ========")
	assert.True(conn != nil)
	assert.True(desc != nil)

	proxy := &TargetProxy{conn: conn, Desc: desc}
	conn.BindDelegate(proxy)

	return proxy
}

func (self *TargetProxy) SendText(text string) {
	log.Println("======== [TargetProxy][SendText] ========")

	self.conn.SendText(text)
}

func (self *TargetProxy) SendBinary(bytes []byte) {
	log.Println("======== [TargetProxy][SendBinary] ========")

	self.conn.SendBinary(bytes)
}

func (self *TargetProxy) Close() {
	log.Println("======== [TargetProxy][Close] ========")

	self.conn.Close()
}

func (self *TargetProxy) OnText(conn Connection, text string) {
	log.Println("======== [TargetProxy][OnText] ========")
	assert.True(conn == self.conn)

	if d := self.Delegate; d != nil {
		d.OnTargetText(self, text)
	}
}

func (self *TargetProxy) OnBinary(conn Connection, bytes []byte) {
	log.Println("======== [TargetProxy][OnBinary] ========")
	assert.True(conn == self.conn)

	if d := self.Delegate; d != nil {
		d.OnTargetBinary(self, bytes)
	}
}

func (self *TargetProxy) OnClosed(conn Connection) {
	log.Println("======== [TargetProxy][OnClosed] ========")
	assert.True(conn == self.conn)

	if d := self.Delegate; d != nil {
		d.OnTargetClosed(self)
	}
}
