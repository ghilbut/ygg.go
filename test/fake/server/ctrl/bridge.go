package fake

import (
	. "github.com/ghilbut/ygg.go/common"
	"github.com/ghilbut/ygg.go/debug"
	. "github.com/ghilbut/ygg.go/server/ctrl"
	"log"
)

type FakeCtrlBridge struct {
	CtrlBridge

	Delegate CtrlBridgeDelegate
	conns    map[string]Connection
}

func NewFakeCtrlBridge() *FakeCtrlBridge {
	log.Println("======== [ctrl.FakeCtrlBridge][NewFakeCtrlBridge] ========")

	bridge := &FakeCtrlBridge{
		Delegate: nil,
		conns:    make(map[string]Connection),
	}

	return bridge
}

func (self *FakeCtrlBridge) Start() {
	log.Println("======== [ctrl.FakeCtrlBridge][Start] ========")
	assert.True(self.Delegate != nil)

	self.Delegate.OnCtrlBridgeStarted(self)
}

func (self *FakeCtrlBridge) Stop() {
	log.Println("======== [ctrl.FakeCtrlBridge][Stop] ========")
	assert.True(self.Delegate != nil)

	self.Delegate.OnCtrlBridgeStopped()
	self.Delegate = nil
}

func (self *FakeCtrlBridge) SetCtrlConnection(conn Connection) {
	log.Println("======== [ctrl.FakeCtrlBridge][SetCtrlConnection] ========")
	assert.True(conn != nil)

	if self.Delegate != nil {
		self.Delegate.OnCtrlConnected(conn)
	} else {
		conn.Close()
	}
}

func (self *FakeCtrlBridge) SetTargetConnection(endpoint string, conn Connection) {
	if conn, ok := self.conns[endpoint]; ok {
		conn.Close()
	}

	self.conns[endpoint] = conn
}

func (self *FakeCtrlBridge) Connect(ctrl *CtrlProxy) Connection {

	if conn, ok := self.conns[ctrl.Desc.Endpoint]; ok {
		delete(self.conns, ctrl.Desc.Endpoint)
		return conn
	}

	ctrl.Close()
	return nil
}
