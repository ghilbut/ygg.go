package target

import (
	. "github.com/ghilbut/ygg.go/common"
	"github.com/ghilbut/ygg.go/debug"
)

type _OnTargetReadyProc func(*TargetProxy)

type TargetReady struct {
	ConnectionDelegate

	readys            map[Connection]bool
	OnTargetReadyProc _OnTargetReadyProc
}

func NewTargetReady() *TargetReady {

	ready := &TargetReady{
		readys:            make(map[Connection]bool),
		OnTargetReadyProc: nil,
	}

	return ready
}

func (self *TargetReady) SetConnection(conn Connection) {
	assert.True(!self.HasConnection(conn))

	self.readys[conn] = true
	conn.BindDelegate(self)
}

func (self *TargetReady) HasConnection(conn Connection) bool {
	_, ok := self.readys[conn]
	return ok
}

func (self *TargetReady) Clear() {
	for conn, _ := range self.readys {
		conn.Close()
	}
}

func (self *TargetReady) OnText(conn Connection, text string) {
	assert.True(self.HasConnection(conn))
	assert.True(self.OnTargetReadyProc != nil)

	var proxy *TargetProxy = nil
	OnTargetReadyProc := self.OnTargetReadyProc

	defer func() {
		if proxy == nil {
			conn.Close()
		}
	}()

	desc, _ := NewTargetDesc(text)
	if desc == nil {
		return
	}

	proxy = NewTargetProxy(conn, desc)
	if proxy != nil {
		OnTargetReadyProc(proxy)
	}
}

func (self *TargetReady) OnBinary(conn Connection, bytes []byte) {
	assert.True(false)
}

func (self *TargetReady) OnClosed(conn Connection) {
	assert.True(self.HasConnection(conn))
	delete(self.readys, conn)
}
