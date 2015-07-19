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

type _NullConnecteeDelegate struct {
	ConnecteeDelegate
}

func (self *_NullConnecteeDelegate) OnConnecteeStarted(c Connectee) {
}

func (self *_NullConnecteeDelegate) OnConnecteeStopped() {
}

func (self *_NullConnecteeDelegate) OnCtrlConnected(c Connection) {
}

func (self *_NullConnecteeDelegate) OnTargetConnected(c Connection) {
}

func Test_FakeConnectee_return_false_when_register_before_started(t *testing.T) {
	log.Println("######## [Test_FakeConnectee_return_false_when_register_failed_before_started] ########")

	connectee := NewFakeConnectee()

	if connectee.Register("A") {
		t.Fail()
	}

	if connectee.HasEndpoint("A") {
		t.Fail()
	}
}

func Test_FakeConnectee_return_true_when_register_on_running(t *testing.T) {
	log.Println("######## [Test_FakeConnectee_return_true_when_register_on_running] ########")

	connectee := NewFakeConnectee()
	connectee.Delegate = &_NullConnecteeDelegate{}

	connectee.Start()

	if connectee.HasEndpoint("A") {
		t.Fail()
	}

	if !connectee.Register("A") {
		t.Fail()
	}

	if !connectee.HasEndpoint("A") {
		t.Fail()
	}
}

func Test_FakeConnectee_return_false_when_register_after_stop(t *testing.T) {
	log.Println("######## [Test_FakeConnectee_return_false_when_register_after_stop] ########")

	connectee := NewFakeConnectee()
	connectee.Delegate = &_NullConnecteeDelegate{}

	connectee.Start()
	connectee.Register("A")

	connectee.Stop()

	if connectee.Register("A") {
		t.Fail()
	}

	if connectee.HasEndpoint("A") {
		t.Fail()
	}
}

func Test_FakeConnectee_return_false_when_endpoint_already_exists(t *testing.T) {
	log.Println("######## [Test_FakeConnectee_return_false_when_endpoint_already_exists] ########")

	connectee := NewFakeConnectee()
	connectee.Delegate = &_NullConnecteeDelegate{}

	connectee.Start()
	connectee.Register("A")

	if connectee.Register("A") {
		t.Fail()
	}

	if !connectee.HasEndpoint("A") {
		t.Fail()
	}
}

func Test_FakeConnectee_unregister(t *testing.T) {
	log.Println("######## [Test_FakeConnectee_unregister] ########")

	connectee := NewFakeConnectee()
	connectee.Delegate = &_NullConnecteeDelegate{}

	connectee.Start()
	connectee.Register("A")

	connectee.Unregister("A")

	if connectee.HasEndpoint("A") {
		t.Fail()
	}
}

func Test_FakeConnectee_delegate_connection_when_set_connection(t *testing.T) {
	log.Println("######## [Test_FakeConnectee_delegate_connection_when_set_connection] ########")

	conn := NewLocalConnection()

	connectee := NewFakeConnectee()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnecteeDelegate(mockCtrl)
	mockDelegate.EXPECT().OnConnecteeStarted(connectee)
	mockDelegate.EXPECT().OnCtrlConnected(conn)
	mockDelegate.EXPECT().OnTargetConnected(conn)

	connectee.Delegate = mockDelegate

	connectee.Start()
	connectee.SetCtrlConnection(conn)
	connectee.SetTargetConnection(conn)
}

func Test_FakeConnectee_close_connection_when_not_started(t *testing.T) {
	log.Println("######## [Test_FakeConnectee_close_connection_when_not_started] ########")

	ctrl := NewLocalConnection()
	target := NewLocalConnection()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnClosed(ctrl)
	mockDelegate.EXPECT().OnClosed(target)
	ctrl.BindDelegate(mockDelegate)
	target.BindDelegate(mockDelegate)

	connectee := NewFakeConnectee()
	connectee.SetCtrlConnection(ctrl)
	connectee.SetTargetConnection(target)
}
