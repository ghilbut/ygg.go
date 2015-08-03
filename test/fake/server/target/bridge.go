package fake

import (
	. "github.com/ghilbut/ygg.go/common"
	"github.com/ghilbut/ygg.go/debug"
	. "github.com/ghilbut/ygg.go/server/target"
	"log"
)

type FakeTargetBridge struct {
	TargetBridge

	endpoints map[string]bool
	Delegate  TargetBridgeDelegate
}

func NewFakeTargetBridge() *FakeTargetBridge {
	log.Println("======== [FakeTargetBridge][NewFakeTargetBridge] ========")

	bridge := &FakeTargetBridge{
		endpoints: make(map[string]bool),
		Delegate:  nil,
	}

	return bridge
}

func (self *FakeTargetBridge) Start() {
	log.Println("======== [FakeTargetBridge][Start] ========")
	assert.True(self.Delegate != nil)

	self.Delegate.OnTargetBridgeStarted(self)
}

func (self *FakeTargetBridge) Stop() {
	log.Println("======== [FakeTargetBridge][Stop] ========")

	for endpoint, _ := range self.endpoints {
		delete(self.endpoints, endpoint)
	}

	self.Delegate.OnTargetBridgeStopped()
	self.Delegate = nil
}

func (self *FakeTargetBridge) Register(endpoint string) bool {
	log.Println("======== [FakeTargetBridge][Register] ========")
	assert.True(len(endpoint) > 0)

	if self.Delegate == nil {
		return false
	}

	if _, ok := self.endpoints[endpoint]; ok {
		return false
	}

	self.endpoints[endpoint] = true
	return true
}

func (self *FakeTargetBridge) Unregister(endpoint string) {
	log.Println("======== [FakeTargetBridge][Unregister] ========")
	assert.True(len(endpoint) > 0)

	if _, ok := self.endpoints[endpoint]; ok {
		delete(self.endpoints, endpoint)
	}
}

func (self *FakeTargetBridge) HasEndpoint(endpoint string) bool {
	log.Println("======== [FakeTargetBridge][HasEndpoint] ========")

	_, ok := self.endpoints[endpoint]
	return ok
}

func (self *FakeTargetBridge) SetCtrlConnection(conn Connection) {
	log.Println("======== [FakeTargetBridge][SetCtrlConnection] ========")
	assert.True(conn != nil)

	if self.Delegate != nil {
		self.Delegate.OnCtrlConnected(conn)
	} else {
		conn.Close()
	}
}

func (self *FakeTargetBridge) SetTargetConnection(conn Connection) {
	log.Println("======== [FakeTargetBridge][SetTargetConnection] ========")
	assert.True(conn != nil)

	if self.Delegate != nil {
		self.Delegate.OnTargetConnected(conn)
	} else {
		conn.Close()
	}
}
