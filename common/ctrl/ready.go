package ctrl

import (
	. "github.com/ghilbut/ygg.go/common"
)

type _OnReadyProc func(*Proxy)

type Ready struct {
	ConnectionDelegate

	readys      map[Connection]bool
	OnReadyProc _OnReadyProc
}

func NewReady() *Ready {

	ready := &Ready{
		readys:      make(map[Connection]bool),
		OnReadyProc: nil,
	}

	return ready
}

func (self *Ready) SetConnection(conn Connection) {
	r := self.readys
	if _, ok := r[conn]; ok {
		panic("[ctrl.Ready] connection is already exists.")
	}

	r[conn] = true
	conn.BindDelegate(self)
}

func (self *Ready) HasConnection(conn Connection) bool {
	_, ok := self.readys[conn]
	return ok
}

func (self *Ready) Clear() {
	r := self.readys
	for conn, _ := range r {
		conn.Close()
	}
}

func (self *Ready) OnText(conn Connection, text string) {

	var proxy *Proxy = nil
	readys := self.readys
	OnReadyProc := self.OnReadyProc

	defer func() {
		if proxy == nil {
			conn.Close()
		}
	}()

	if _, ok := readys[conn]; !ok {
		panic("[ctrl.Ready] there is no matched connection.")
	}

	if OnReadyProc == nil {
		panic("[ctrl.Ready] OnReadyProc callback is nil.")
	}

	desc, _ := NewDesc(text)
	if desc == nil {
		return
	}

	proxy = NewProxy(conn, desc)
	if proxy != nil {
		OnReadyProc(proxy)
	}
}

func (self *Ready) OnBinary(conn Connection, bytes []byte) {
	// nothing
}

func (self *Ready) OnClosed(conn Connection) {
	r := self.readys
	if _, ok := r[conn]; !ok {
		panic("[ctrl.Ready] there is no matched connection.")
	}

	delete(r, conn)
}
