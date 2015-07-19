package fake

import (
	. "github.com/ghilbut/ygg.go/common"
	"github.com/ghilbut/ygg.go/debug"
	. "github.com/ghilbut/ygg.go/server/ctrl"
	"log"
)

type FakeConnectee struct {
	Connectee

	Delegate ConnecteeDelegate
}

func NewFakeConnectee() *FakeConnectee {
	log.Println("======== [ctrl.FakeConnectee][NewFakeConnectee] ========")

	connectee := &FakeConnectee{
		Delegate: nil,
	}

	return connectee
}

func (self *FakeConnectee) Start() {
	log.Println("======== [ctrl.FakeConnectee][Start] ========")
	assert.True(self.Delegate != nil)

	self.Delegate.OnConnecteeStarted(self)
}

func (self *FakeConnectee) Stop() {
	log.Println("======== [ctrl.FakeConnectee][Stop] ========")
	assert.True(self.Delegate != nil)

	self.Delegate.OnConnecteeStopped()
	self.Delegate = nil
}

func (self *FakeConnectee) SetCtrlConnection(conn Connection) {
	log.Println("======== [ctrl.FakeConnectee][SetCtrlConnection] ========")
	assert.True(conn != nil)

	if self.Delegate != nil {
		self.Delegate.OnCtrlConnected(conn)
	} else {
		conn.Close()
	}
}
