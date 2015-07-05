package common

type Connection struct {
	ConnectorDelegate

	conn     Connector
	delegate ConnectionDelegate
}

type ConnectionDelegate interface {
	OnText(conn *Connection, text string)
	OnBinary(conn *Connection, bytes []byte)
	OnClosed(conn *Connection)
}

func NewConnection(conn Connector) *Connection {

	self := &Connection{conn: conn, delegate: nil}
	conn.BindDelegate(self)

	return self
}

func (self *Connection) BindDelegate(delegate ConnectionDelegate) {
	self.delegate = delegate
}

func (self *Connection) UnbindDelegate() {
	self.delegate = nil
}

func (self *Connection) SendText(text string) {
	self.conn.SendText(text)
}

func (self *Connection) SendBinary(bytes []byte) {
	self.conn.SendBinary(bytes)
}

func (self *Connection) Close() {
	self.conn.Close()
}

func (self *Connection) OnText(text string) {
	d := self.delegate
	if d != nil {
		d.OnText(self, text)
	}
}

func (self *Connection) OnBinary(bytes []byte) {
	d := self.delegate
	if d != nil {
		d.OnBinary(self, bytes)
	}
}

func (self *Connection) OnClosed() {
	d := self.delegate
	if d != nil {
		d.OnClosed(self)
	}
}
