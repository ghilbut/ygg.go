package target

import (
	. "github.com/ghilbut/ygg.go/common"
	"github.com/ghilbut/ygg.go/debug"
	"log"
)

type TargetReady struct {
	ConnectionDelegate

	readys   map[Connection]bool
	Delegate TargetReadyDelegate
}

type TargetReadyDelegate interface {
	OnTargetProxy(proxy *TargetProxy)
}

func NewTargetReady() *TargetReady {
	log.Printf("======== [target.TargetReady][NewTargetReady] ========")

	ready := &TargetReady{
		readys:   make(map[Connection]bool),
		Delegate: nil,
	}

	return ready
}

func (self *TargetReady) SetConnection(conn Connection) {
	log.Printf("======== [target.TargetReady][SetConnection] ========")
	assert.True(conn != nil)
	assert.True(!self.HasConnection(conn))

	self.readys[conn] = true
	conn.BindDelegate(self)
}

func (self *TargetReady) HasConnection(conn Connection) bool {
	log.Printf("======== [target.TargetReady][HasConnection] ========")
	assert.True(conn != nil)

	_, ok := self.readys[conn]
	return ok
}

func (self *TargetReady) Clear() {
	log.Printf("======== [target.TargetReady][Clear] ========")

	for conn, _ := range self.readys {
		conn.Close()
	}
}

func (self *TargetReady) OnText(conn Connection, text string) {
	log.Printf("======== [target.TargetReady][OnText] ========")
	assert.True(self.HasConnection(conn))
	assert.True(self.Delegate != nil)

	var proxy *TargetProxy = nil

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
		delete(self.readys, conn)
		self.Delegate.OnTargetProxy(proxy)
	}
}

func (self *TargetReady) OnBinary(conn Connection, bytes []byte) {
	log.Printf("======== [target.TargetReady][OnBinary] ========")
	assert.True(false)
}

func (self *TargetReady) OnClosed(conn Connection) {
	log.Printf("======== [target.TargetReady][OnClosed] ========")
	assert.True(self.HasConnection(conn))

	delete(self.readys, conn)
}
