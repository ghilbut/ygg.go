package ctrl

import (
	. "github.com/ghilbut/ygg.go/common"
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
	r := self.readys
	if _, ok := r[conn]; ok {
		panic("[ctrl.TargetReady] connection is already exists.")
	}

	r[conn] = true
	conn.BindDelegate(self)
}

func (self *TargetReady) HasConnection(conn Connection) bool {
	_, ok := self.readys[conn]
	return ok
}

func (self *TargetReady) Clear() {
	r := self.readys
	for conn, _ := range r {
		conn.Close()
	}
}

func (self *TargetReady) OnText(conn Connection, text string) {

	var proxy *TargetProxy = nil
	readys := self.readys
	OnTargetReadyProc := self.OnTargetReadyProc

	defer func() {
		if proxy == nil {
			conn.Close()
		}
	}()

	if _, ok := readys[conn]; !ok {
		panic("[ctrl.TargetReady] there is no matched connection.")
	}

	if OnTargetReadyProc == nil {
		panic("[ctrl.TargetReady] OnTargetReadyProc callback is nil.")
	}

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
	// nothing
}

func (self *TargetReady) OnClosed(conn Connection) {
	r := self.readys
	if _, ok := r[conn]; !ok {
		panic("[ctrl.TargetReady] there is no matched connection.")
	}

	delete(r, conn)
}
