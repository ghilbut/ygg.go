package fake

import (
	. "github.com/ghilbut/ygg.go/common"
	"github.com/ghilbut/ygg.go/debug"
	. "github.com/ghilbut/ygg.go/server/ctrl"
	"log"
)

type FakeConnectee struct {
	Connectee

	delegate ConnecteeDelegate
}

func NewFakeConnectee() *FakeConnectee {
	log.Println("======== [ctrl.FakeConnectee][NewFakeConnectee] ========")

	connectee := &FakeConnectee{
		delegate: nil,
	}

	return connectee
}

func (self *FakeConnectee) Start(delegate ConnecteeDelegate) {
	log.Println("======== [ctrl.FakeConnectee][Start] ========")
	assert.True(delegate != nil)

	self.delegate = delegate
}

func (self *FakeConnectee) Stop() {
	log.Println("======== [ctrl.FakeConnectee][Stop] ========")

	self.delegate = nil
}

func (self *FakeConnectee) SetCtrlConnection(conn Connection) {
	log.Println("======== [ctrl.FakeConnectee][SetCtrlConnection] ========")
	assert.True(conn != nil)

	if self.delegate != nil {
		self.delegate.OnCtrlConnected(conn)
	} else {
		conn.Close()
	}
}
