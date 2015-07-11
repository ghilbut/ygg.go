package common

import (
	"github.com/ghilbut/ygg.go/debug"
)

type _OnCtrlReadyProc func(*CtrlProxy)

type CtrlReady struct {
	ConnectionDelegate

	readys          map[Connection]bool
	OnCtrlReadyProc _OnCtrlReadyProc
}

func NewCtrlReady() *CtrlReady {

	ready := &CtrlReady{
		readys:          make(map[Connection]bool),
		OnCtrlReadyProc: nil,
	}

	return ready
}

func (self *CtrlReady) SetConnection(conn Connection) {
	assert.False(self.HasConnection(conn))

	self.readys[conn] = true
	conn.BindDelegate(self)
}

func (self *CtrlReady) HasConnection(conn Connection) bool {
	_, ok := self.readys[conn]
	return ok
}

func (self *CtrlReady) Clear() {
	r := self.readys
	for conn, _ := range r {
		conn.Close()
	}
}

func (self *CtrlReady) OnText(conn Connection, text string) {
	assert.True(self.HasConnection(conn))
	assert.True(self.OnCtrlReadyProc != nil)

	var proxy *CtrlProxy = nil
	OnCtrlReadyProc := self.OnCtrlReadyProc

	defer func() {
		if proxy == nil {
			conn.Close()
		}
	}()

	desc, _ := NewCtrlDesc(text)
	if desc == nil {
		return
	}

	proxy = NewCtrlProxy(conn, desc)
	if proxy != nil {
		OnCtrlReadyProc(proxy)
	}
}

func (self *CtrlReady) OnBinary(conn Connection, bytes []byte) {
	assert.True(false)
}

func (self *CtrlReady) OnClosed(conn Connection) {
	assert.True(self.HasConnection(conn))
	delete(self.readys, conn)
}
