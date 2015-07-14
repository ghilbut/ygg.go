package ctrl

import (
	. "github.com/ghilbut/ygg.go/common"
	"github.com/ghilbut/ygg.go/debug"
	"log"
)

type TargetReady struct {
	ConnectionDelegate

	readys      map[Connection]*CtrlProxy
	reverseFind map[*CtrlProxy]Connection
	Delegate    TargetReadyDelegate
}

type TargetReadyDelegate interface {
	OnTargetProxy(ctrl *CtrlProxy, target *TargetProxy)
}

func NewTargetReady() *TargetReady {
	log.Printf("======== [ctrl.TargetReady][NewTargetReady] ========")

	ready := &TargetReady{
		readys:      make(map[Connection]*CtrlProxy),
		reverseFind: make(map[*CtrlProxy]Connection),
		Delegate:    nil,
	}

	return ready
}

func (self *TargetReady) SetConnection(conn Connection, ctrl *CtrlProxy) {
	log.Printf("======== [ctrl.TargetReady][SetConnection] ========")
	assert.True(conn != nil)
	assert.True(ctrl != nil)
	assert.True(!self.HasConnection(conn))

	self.readys[conn] = ctrl
	self.reverseFind[ctrl] = conn
	conn.BindDelegate(self)
}

func (self *TargetReady) HasConnection(conn Connection) bool {
	log.Printf("======== [ctrl.TargetReady][HasConnection] ========")
	assert.True(conn != nil)

	ctrl, ok := self.readys[conn]
	if ok {
		check, ok := self.reverseFind[ctrl]
		assert.True(ok && check == conn)
	}
	return ok
}

func (self *TargetReady) Clear() {
	log.Printf("======== [ctrl.TargetReady][Clear] ========")

	for conn, _ := range self.readys {
		conn.Close()
	}
}

func (self *TargetReady) OnText(conn Connection, text string) {
	log.Printf("======== [ctrl.TargetReady][OnText] ========")
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
		ctrl := self.readys[conn]
		self.Delegate.OnTargetProxy(ctrl, proxy)
	}
}

func (self *TargetReady) OnBinary(conn Connection, bytes []byte) {
	log.Printf("======== [ctrl.TargetReady][OnBinary] ========")
	assert.True(false)
}

func (self *TargetReady) OnClosed(conn Connection) {
	log.Printf("======== [ctrl.TargetReady][OnClosed] ========")
	assert.True(self.HasConnection(conn))

	ctrl, _ := self.readys[conn]
	delete(self.readys, conn)
	delete(self.reverseFind, ctrl)
}
