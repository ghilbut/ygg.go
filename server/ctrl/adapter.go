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

func NewOneToOneAdapter(proxy *TargetProxy) Adapter {
	log.Println("======== [OneToOneAdapter][NewOneToOneAdapter] ========")

	adapter := &OneToOneAdapter{
		ctrl:     nil,
		target:   proxy,
		delegate: nil,
		q:        make(chan bool),
	}

	go func(adapter *OneToOneAdapter) {
		log.Println("======== [OneToOneAdapter][Closing][Wait] ========")

		<-adapter.q
		close(adapter.q)

		log.Println("======== [OneToOneAdapter][Closing][Continue] ========")

		if ctrl := adapter.ctrl; ctrl != nil {
			adapter.ctrl = nil
			ctrl.Delegate = nil
			ctrl.Close()
		}

		if target := adapter.target; target != nil {
			adapter.target = nil
			target.Delegate = nil
			target.Close()
		}

		if adapter.delegate != nil {
			adapter.delegate.OnClosed(adapter)
		}
	}(adapter)

	adapter.target.Delegate = adapter
	return adapter
}

func (self *OneToOneAdapter) BindDelegate(delegate AdapterDelegate) {
	log.Println("======== [OneToOneAdapter][BindDelegate] ========")

	self.delegate = delegate
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

	//	defer recover()
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

	//	defer recover()
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

	//	defer recover()
	self.q <- true
}
