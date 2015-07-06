package target

import (
	. "github.com/ghilbut/ygg.go/common"
)

type Proxy struct {
	ConnectionDelegate

	desc     *Desc
	conn     Connection
	Delegate ProxyDelegate
}

type ProxyDelegate interface {
	OnText(proxy *Proxy, text string)
	OnBinary(proxy *Proxy, bytes []byte)
	OnClosed(proxy *Proxy)
}

func NewProxy(desc *Desc, conn Connection) *Proxy {

	if desc == nil {
		panic("desc is nil.")
	}

	if conn == nil {
		panic("connection is nil.")
	}

	proxy := &Proxy{desc: desc, conn: conn}
	conn.BindDelegate(proxy)

	return proxy
}

func (self *Proxy) SendText(text string) {
	self.conn.SendText(text)
}

func (self *Proxy) SendBinary(bytes []byte) {
	self.conn.SendBinary(bytes)
}

func (self *Proxy) Close() {
	self.conn.Close()
}

func (self *Proxy) OnText(conn Connection, text string) {
	if conn != self.conn {
		panic("conn is invalid connection.")
	}
	if d := self.Delegate; d != nil {
		d.OnText(self, text)
	}
}

func (self *Proxy) OnBinary(conn Connection, bytes []byte) {
	if conn != self.conn {
		panic("conn is invalid connection.")
	}
	if d := self.Delegate; d != nil {
		d.OnBinary(self, bytes)
	}

}

func (self *Proxy) OnClosed(conn Connection) {
	if conn != self.conn {
		panic("conn is invalid connection.")
	}
	if d := self.Delegate; d != nil {
		d.OnClosed(self)
	}

}
