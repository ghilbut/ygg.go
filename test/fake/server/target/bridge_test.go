package fake_test

import (
	. "github.com/ghilbut/ygg.go/common"
	. "github.com/ghilbut/ygg.go/net"
	. "github.com/ghilbut/ygg.go/server/target"
	. "github.com/ghilbut/ygg.go/test/fake/server/target"
	. "github.com/ghilbut/ygg.go/test/mock/common"
	. "github.com/ghilbut/ygg.go/test/mock/server/target"
	"github.com/golang/mock/gomock"
	"log"
	"testing"
)

type _NullTargetBridgeDelegate struct {
	TargetBridgeDelegate
}

func (self *_NullTargetBridgeDelegate) OnTargetBridgeStarted(c TargetBridge) {
}

func (self *_NullTargetBridgeDelegate) OnTargetBridgeStopped() {
}

func (self *_NullTargetBridgeDelegate) OnCtrlConnected(c Connection) {
}

func (self *_NullTargetBridgeDelegate) OnTargetConnected(c Connection) {
}

func Test_FakeTargetBridge_return_false_when_register_before_started(t *testing.T) {
	log.Println("######## [Test_FakeTargetBridge_return_false_when_register_failed_before_started] ########")

	bridge := NewFakeTargetBridge()

	if bridge.Register("A") {
		t.Fail()
	}

	if bridge.HasEndpoint("A") {
		t.Fail()
	}
}

func Test_FakeTargetBridge_return_true_when_register_on_running(t *testing.T) {
	log.Println("######## [Test_FakeTargetBridge_return_true_when_register_on_running] ########")

	bridge := NewFakeTargetBridge()
	bridge.Delegate = &_NullTargetBridgeDelegate{}

	bridge.Start()

	if bridge.HasEndpoint("A") {
		t.Fail()
	}

	if !bridge.Register("A") {
		t.Fail()
	}

	if !bridge.HasEndpoint("A") {
		t.Fail()
	}
}

func Test_FakeTargetBridge_return_false_when_register_after_stop(t *testing.T) {
	log.Println("######## [Test_FakeTargetBridge_return_false_when_register_after_stop] ########")

	bridge := NewFakeTargetBridge()
	bridge.Delegate = &_NullTargetBridgeDelegate{}

	bridge.Start()
	bridge.Register("A")

	bridge.Stop()

	if bridge.Register("A") {
		t.Fail()
	}

	if bridge.HasEndpoint("A") {
		t.Fail()
	}
}

func Test_FakeTargetBridge_return_false_when_endpoint_already_exists(t *testing.T) {
	log.Println("######## [Test_FakeTargetBridge_return_false_when_endpoint_already_exists] ########")

	bridge := NewFakeTargetBridge()
	bridge.Delegate = &_NullTargetBridgeDelegate{}

	bridge.Start()
	bridge.Register("A")

	if bridge.Register("A") {
		t.Fail()
	}

	if !bridge.HasEndpoint("A") {
		t.Fail()
	}
}

func Test_FakeTargetBridge_unregister(t *testing.T) {
	log.Println("######## [Test_FakeTargetBridge_unregister] ########")

	bridge := NewFakeTargetBridge()
	bridge.Delegate = &_NullTargetBridgeDelegate{}

	bridge.Start()
	bridge.Register("A")

	bridge.Unregister("A")

	if bridge.HasEndpoint("A") {
		t.Fail()
	}
}

func Test_FakeTargetBridge_delegate_connection_when_set_connection(t *testing.T) {
	log.Println("######## [Test_FakeTargetBridge_delegate_connection_when_set_connection] ########")

	conn := NewLocalConnection()

	bridge := NewFakeTargetBridge()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockTargetBridgeDelegate(mockCtrl)
	mockDelegate.EXPECT().OnTargetBridgeStarted(bridge)
	mockDelegate.EXPECT().OnCtrlConnected(conn)
	mockDelegate.EXPECT().OnTargetConnected(conn)

	bridge.Delegate = mockDelegate

	bridge.Start()
	bridge.SetCtrlConnection(conn)
	bridge.SetTargetConnection(conn)
}

func Test_FakeTargetBridge_close_connection_when_not_started(t *testing.T) {
	log.Println("######## [Test_FakeTargetBridge_close_connection_when_not_started] ########")

	ctrl := NewLocalConnection()
	target := NewLocalConnection()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnClosed(ctrl)
	mockDelegate.EXPECT().OnClosed(target)
	ctrl.BindDelegate(mockDelegate)
	target.BindDelegate(mockDelegate)

	bridge := NewFakeTargetBridge()
	bridge.SetCtrlConnection(ctrl)
	bridge.SetTargetConnection(target)
}
