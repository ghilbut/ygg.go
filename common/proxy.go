package common

type CtrlProxy struct {
	ConnectionDelegate

	conn     Connection
	Desc     *CtrlDesc
	Delegate CtrlProxyDelegate
}

type CtrlProxyDelegate interface {
	OnText(proxy *CtrlProxy, text string)
	OnBinary(proxy *CtrlProxy, bytes []byte)
	OnClosed(proxy *CtrlProxy)
}

func NewCtrlProxy(conn Connection, desc *CtrlDesc) *CtrlProxy {

	if conn == nil {
		panic("connection is nil.")
	}

	if desc == nil {
		panic("desc is nil.")
	}

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
	if conn != self.conn {
		panic("conn is invalid connection.")
	}
	if d := self.Delegate; d != nil {
		d.OnText(self, text)
	}
}

func (self *CtrlProxy) OnBinary(conn Connection, bytes []byte) {
	if conn != self.conn {
		panic("conn is invalid connection.")
	}
	if d := self.Delegate; d != nil {
		d.OnBinary(self, bytes)
	}
}

func (self *CtrlProxy) OnClosed(conn Connection) {
	if conn != self.conn {
		panic("conn is invalid connection.")
	}
	if d := self.Delegate; d != nil {
		d.OnClosed(self)
	}
}

type TargetProxy struct {
	ConnectionDelegate

	conn     Connection
	desc     *TargetDesc
	Delegate TargetProxyDelegate
}

type TargetProxyDelegate interface {
	OnText(proxy *TargetProxy, text string)
	OnBinary(proxy *TargetProxy, bytes []byte)
	OnClosed(proxy *TargetProxy)
}

func NewTargetProxy(conn Connection, desc *TargetDesc) *TargetProxy {

	if conn == nil {
		panic("connection is nil.")
	}

	if desc == nil {
		panic("desc is nil.")
	}

	proxy := &TargetProxy{conn: conn, desc: desc}
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
	if conn != self.conn {
		panic("conn is invalid connection.")
	}
	if d := self.Delegate; d != nil {
		d.OnText(self, text)
	}
}

func (self *TargetProxy) OnBinary(conn Connection, bytes []byte) {
	if conn != self.conn {
		panic("conn is invalid connection.")
	}
	if d := self.Delegate; d != nil {
		d.OnBinary(self, bytes)
	}
}

func (self *TargetProxy) OnClosed(conn Connection) {
	if conn != self.conn {
		panic("conn is invalid connection.")
	}
	if d := self.Delegate; d != nil {
		d.OnClosed(self)
	}
}
