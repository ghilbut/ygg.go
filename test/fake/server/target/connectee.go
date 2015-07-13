package fake

import (
	. "github.com/ghilbut/ygg.go/common"
	"github.com/ghilbut/ygg.go/debug"
	. "github.com/ghilbut/ygg.go/server/target"
	"log"
)

type FakeConnectee struct {
	Connectee

	endpoints map[string]bool
	delegate  ConnecteeDelegate
}

func NewFakeConnectee() *FakeConnectee {
	log.Println("======== [FakeConnectee][NewFakeConnectee] ========")

	connectee := &FakeConnectee{
		endpoints: make(map[string]bool),
		delegate:  nil,
	}

	return connectee
}

func (self *FakeConnectee) Start(delegate ConnecteeDelegate) {
	log.Println("======== [FakeConnectee][Start] ========")
	assert.True(delegate != nil)

	self.delegate = delegate
}

func (self *FakeConnectee) Stop() {
	log.Println("======== [FakeConnectee][Stop] ========")

	for endpoint, _ := range self.endpoints {
		delete(self.endpoints, endpoint)
	}

	self.delegate = nil
}

func (self *FakeConnectee) Register(endpoint string) bool {
	log.Println("======== [FakeConnectee][Register] ========")
	assert.True(len(endpoint) > 0)

	if self.delegate == nil {
		return false
	}

	if _, ok := self.endpoints[endpoint]; ok {
		return false
	}

	self.endpoints[endpoint] = true
	return true
}

func (self *FakeConnectee) Unregister(endpoint string) {
	log.Println("======== [FakeConnectee][Unregister] ========")
	assert.True(len(endpoint) > 0)

	if _, ok := self.endpoints[endpoint]; ok {
		delete(self.endpoints, endpoint)
	}
}

func (self *FakeConnectee) HasEndpoint(endpoint string) bool {
	log.Println("======== [FakeConnectee][HasEndpoint] ========")

	_, ok := self.endpoints[endpoint]
	return ok
}

func (self *FakeConnectee) SetCtrlConnection(conn Connection) {
	log.Println("======== [FakeConnectee][SetCtrlConnection] ========")
	assert.True(conn != nil)

	if self.delegate != nil {
		self.delegate.OnCtrlConnected(conn)
	} else {
		conn.Close()
	}
}

func (self *FakeConnectee) SetTargetConnection(conn Connection) {
	log.Println("======== [FakeConnectee][SetTargetConnection] ========")
	assert.True(conn != nil)

	if self.delegate != nil {
		self.delegate.OnTargetConnected(conn)
	} else {
		conn.Close()
	}
}
