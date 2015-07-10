package common

import (
	. "github.com/ghilbut/ygg.go/debug"
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
	r := self.readys

	_, ok := r[conn]
	Assert(!ok, "[ctrl.CtrlReady] connection is already exists.")

	r[conn] = true
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

	var proxy *CtrlProxy = nil
	readys := self.readys
	OnCtrlReadyProc := self.OnCtrlReadyProc

	defer func() {
		if proxy == nil {
			conn.Close()
		}
	}()

	_, ok := readys[conn]
	Assert(ok, "[ctrl.CtrlReady] there is no matched connection.")
	Assert(OnCtrlReadyProc != nil, "[ctrl.CtrlReady] OnCtrlReadyProc callback is nil.")

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
	// nothing
}

func (self *CtrlReady) OnClosed(conn Connection) {
	r := self.readys

	_, ok := r[conn]
	Assert(ok, "[ctrl.CtrlReady] there is no matched connection.")

	delete(r, conn)
}
