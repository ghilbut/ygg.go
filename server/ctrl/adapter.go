package ctrl

import (
	. "github.com/ghilbut/ygg.go/common"
	"github.com/ghilbut/ygg.go/debug"
	"log"
)

type OneToOneAdapter struct {
	Adapter
	CtrlProxyDelegate
	TargetProxyDelegate

	ctrl     *CtrlProxy
	target   *TargetProxy
	delegate AdapterDelegate
	q        chan bool
}

func NewOneToOneAdapter(proxy *TargetProxy) *OneToOneAdapter {
	log.Println("======== [OneToOneAdapter][NewOneToOneAdapter] ========")
	assert.True(proxy != nil)

	adapter := &OneToOneAdapter{
		ctrl:     nil,
		target:   proxy,
		delegate: nil,
		q:        make(chan bool),
	}

	if adapter != nil {
		go adapter.quit()
		adapter.target.Delegate = adapter
	}

	adapter.target.Delegate = adapter
	return adapter
}

func (self *OneToOneAdapter) quit() {
	log.Println("======== [OneToOneAdapter][quit][Wait] ========")

	<-self.q
	close(self.q)

	log.Println("======== [OneToOneAdapter][quit][Continue] ========")

	if self.ctrl != nil {
		self.ctrl.Close()
	}

	if self.target != nil {
		self.target.Close()
	}

	if self.delegate != nil {
		self.delegate.OnAdapterClosed(self)
	}
}

func (self *OneToOneAdapter) BindDelegate(delegate AdapterDelegate) {
	log.Println("======== [OneToOneAdapter][BindDelegate] ========")
	assert.True(delegate != nil)

	self.delegate = delegate
}

func (self *OneToOneAdapter) UnbindDelegate() {
	log.Println("======== [OneToOneAdapter][UnbindDelegate] ========")

	self.delegate = nil
}

func (self *OneToOneAdapter) SetCtrlProxy(proxy *CtrlProxy) {
	log.Println("======== [OneToOneAdapter][SetCtrlProxy] ========")

	assert.True(proxy != nil)
	assert.True(self.ctrl == nil)

	self.ctrl = proxy
	self.ctrl.Delegate = self
}

func (self *OneToOneAdapter) Close() {
	log.Println("======== [OneToOneAdapter][Close] ========")

	defer func() { recover() }()
	self.q <- true
}

func (self *OneToOneAdapter) OnCtrlText(proxy *CtrlProxy, text string) {
	log.Println("======== [OneToOneAdapter][OnCtrlText] ========")

	assert.True(proxy != nil)
	assert.True(proxy == self.ctrl)

	if self.target != nil {
		self.target.SendText(text)
	}
}

func (self *OneToOneAdapter) OnCtrlBinary(proxy *CtrlProxy, bytes []byte) {
	log.Println("======== [OneToOneAdapter][OnCtrlBinary] ========")

	assert.True(proxy != nil)
	assert.True(proxy == self.ctrl)

	if self.target != nil {
		self.target.SendBinary(bytes)
	}
}

func (self *OneToOneAdapter) OnCtrlClosed(proxy *CtrlProxy) {
	log.Println("======== [OneToOneAdapter][OnCtrlClosed] ========")

	assert.True(proxy != nil)
	assert.True(proxy == self.ctrl)

	self.ctrl.Delegate = nil
	self.ctrl = nil

	defer func() { recover() }()
	self.q <- true
}

func (self *OneToOneAdapter) OnTargetText(proxy *TargetProxy, text string) {
	log.Println("======== [OneToOneAdapter][OnTargetText] ========")

	assert.True(proxy != nil)
	assert.True(proxy == self.target)

	if self.ctrl != nil {
		self.ctrl.SendText(text)
	}
}

func (self *OneToOneAdapter) OnTargetBinary(proxy *TargetProxy, bytes []byte) {
	log.Println("======== [OneToOneAdapter][OnTargetBinary] ========")

	assert.True(proxy != nil)
	assert.True(proxy == self.target)

	if self.ctrl != nil {
		self.ctrl.SendBinary(bytes)
	}
}

func (self *OneToOneAdapter) OnTargetClosed(proxy *TargetProxy) {
	log.Println("======== [OneToOneAdapter][OnTargetClosed] ========")

	assert.True(proxy != nil)
	assert.True(proxy == self.target)

	self.target.Delegate = nil
	self.target = nil

	defer func() { recover() }()
	self.q <- true
}
