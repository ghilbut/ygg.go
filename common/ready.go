package common

import (
	"github.com/ghilbut/ygg.go/debug"
	"log"
)

type CtrlReady struct {
	ConnectionDelegate

	readys   map[Connection]bool
	Delegate CtrlReadyDelegate
}

type CtrlReadyDelegate interface {
	OnCtrlProxy(proxy *CtrlProxy)
}

func NewCtrlReady() *CtrlReady {
	log.Printf("======== [CtrlReady][NewCtrlReady] ========")

	ready := &CtrlReady{
		readys:   make(map[Connection]bool),
		Delegate: nil,
	}

	return ready
}

func (self *CtrlReady) SetConnection(conn Connection) {
	log.Printf("======== [CtrlReady][SetConnection] ========")
	assert.True(!self.HasConnection(conn))

	self.readys[conn] = true
	conn.BindDelegate(self)
}

func (self *CtrlReady) HasConnection(conn Connection) bool {
	log.Printf("======== [CtrlReady][HasConnection] ========")
	assert.True(conn != nil)

	_, ok := self.readys[conn]
	return ok
}

func (self *CtrlReady) Clear() {
	log.Printf("======== [CtrlReady][Clear] ========")

	r := self.readys
	for conn, _ := range r {
		conn.Close()
	}
}

func (self *CtrlReady) OnText(conn Connection, text string) {
	log.Printf("======== [CtrlReady][OnText] ========")
	assert.True(self.HasConnection(conn))
	assert.True(self.Delegate != nil)

	var proxy *CtrlProxy = nil

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
		self.Delegate.OnCtrlProxy(proxy)
	}
}

func (self *CtrlReady) OnBinary(conn Connection, bytes []byte) {
	log.Printf("======== [CtrlReady][OnBinary] ========")
	assert.True(false)
}

func (self *CtrlReady) OnClosed(conn Connection) {
	log.Printf("======== [CtrlReady][OnClosed] ========")
	assert.True(self.HasConnection(conn))

	delete(self.readys, conn)
}
