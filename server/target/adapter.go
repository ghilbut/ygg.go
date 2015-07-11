package target

import (
	. "github.com/ghilbut/ygg.go/common"
	"github.com/ghilbut/ygg.go/debug"
	"log"
)

type ManyToOneAdapter struct {
	Adapter
	CtrlProxyDelegate
	TargetProxyDelegate

	ctrls    map[*CtrlProxy]bool
	target   *TargetProxy
	delegate AdapterDelegate
	q        chan bool
}

func NewManyToOneAdapter(proxy *TargetProxy) *ManyToOneAdapter {
	log.Println("======== [ManyToOneAdapter][NewManyToOneAdapter] ========")

	adapter := &ManyToOneAdapter{
		ctrls:    make(map[*CtrlProxy]bool),
		target:   proxy,
		delegate: nil,
		q:        make(chan bool),
	}

	go func() {
		log.Println("======== [ManyToOneAdapter][Closing][Wait] ========")

		<-adapter.q
		close(adapter.q)

		log.Println("======== [ManyToOneAdapter][Closing][Continue] ========")

		for ctrl := range adapter.ctrls {
			delete(adapter.ctrls, ctrl)
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
	}()

	adapter.target.Delegate = adapter
	return adapter
}

func (self *ManyToOneAdapter) BindDelegate(delegate AdapterDelegate) {
	log.Println("======== [ManyToOneAdapter][BindDelegate] ========")

	self.delegate = delegate
}

func (self *ManyToOneAdapter) SetCtrlProxy(proxy *CtrlProxy) {
	log.Println("======== [ManyToOneAdapter][SetCtrlProxy] ========")

	assert.True(proxy != nil)
	assert.False(self.HasCtrlProxy(proxy))

	self.ctrls[proxy] = true
	proxy.Delegate = self
}

func (self *ManyToOneAdapter) HasCtrlProxy(proxy *CtrlProxy) bool {
	log.Println("======== [ManyToOneAdapter][HasCtrlProxy] ========")

	assert.True(proxy != nil)

	_, ok := self.ctrls[proxy]
	return ok
}

func (self *ManyToOneAdapter) Close() {
	log.Println("======== [ManyToOneAdapter][Close] ========")

	//	defer recover()
	self.q <- true
}

func (self *ManyToOneAdapter) OnCtrlText(proxy *CtrlProxy, text string) {
	log.Println("======== [ManyToOneAdapter][OnCtrlText] ========")

	assert.True(proxy != nil)
	assert.True(self.HasCtrlProxy(proxy))

	self.target.SendText(text)
}

func (self *ManyToOneAdapter) OnCtrlBinary(proxy *CtrlProxy, bytes []byte) {
	log.Println("======== [ManyToOneAdapter][OnCtrlBinary] ========")

	assert.True(proxy != nil)
	assert.True(self.HasCtrlProxy(proxy))

	self.target.SendBinary(bytes)
}

func (self *ManyToOneAdapter) OnCtrlClosed(proxy *CtrlProxy) {
	log.Println("======== [ManyToOneAdapter][OnCtrlClosed] ========")

	assert.True(proxy != nil)
	assert.True(self.HasCtrlProxy(proxy))

	delete(self.ctrls, proxy)
}

func (self *ManyToOneAdapter) OnTargetText(proxy *TargetProxy, text string) {
	log.Println("======== [ManyToOneAdapter][OnTargetText] ========")

	assert.True(proxy != nil)
	assert.True(proxy == self.target)

	for ctrl, _ := range self.ctrls {
		ctrl.SendText(text)
	}
}

func (self *ManyToOneAdapter) OnTargetBinary(proxy *TargetProxy, bytes []byte) {
	log.Println("======== [ManyToOneAdapter][OnTargetBinary] ========")

	assert.True(proxy != nil)
	assert.True(proxy == self.target)

	for ctrl, _ := range self.ctrls {
		ctrl.SendBinary(bytes)
	}
}

func (self *ManyToOneAdapter) OnTargetClosed(proxy *TargetProxy) {
	log.Println("======== [ManyToOneAdapter][OnTargetClosed] ========")

	assert.True(proxy != nil)
	assert.True(proxy == self.target)

	self.q <- true
}
