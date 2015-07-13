package fake_test

import (
	. "github.com/ghilbut/ygg.go/common"
	. "github.com/ghilbut/ygg.go/server/target"
	. "github.com/ghilbut/ygg.go/test/fake"
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

func (self *_NullConnecteeDelegate) OnCtrlConnected(conn Connection) {
}

func (self *_NullConnecteeDelegate) OnTargetConnected(conn Connection) {
}

func Test_FakeConnectee_return_false_when_register_before_started(t *testing.T) {
	log.Println("######## [Test_FakeConnectee_return_false_when_register_failed_before_started] ########")

	var connectee Connectee = NewFakeConnectee()

	if connectee.Register("A") {
		t.Fail()
	}

	if connectee.HasEndpoint("A") {
		t.Fail()
	}
}

func Test_FakeConnectee_return_true_when_register_on_running(t *testing.T) {
	log.Println("######## [Test_FakeConnectee_return_true_when_register_on_running] ########")

	var connectee Connectee = NewFakeConnectee()
	delegate := &_NullConnecteeDelegate{}

	connectee.Start(delegate)

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

	var connectee Connectee = NewFakeConnectee()
	delegate := &_NullConnecteeDelegate{}
	connectee.Start(delegate)
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

	delegate := &_NullConnecteeDelegate{}

	var connectee Connectee = NewFakeConnectee()
	connectee.Start(delegate)
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

	var connectee Connectee = NewFakeConnectee()
	delegate := &_NullConnecteeDelegate{}
	connectee.Start(delegate)
	connectee.Register("A")

	connectee.Unregister("A")

	if connectee.HasEndpoint("A") {
		t.Fail()
	}
}

func Test_FakeConnectee_delegate_connection_when_set_connection(t *testing.T) {
	log.Println("######## [Test_FakeConnectee_delegate_connection_when_set_connection] ########")

	conn := NewFakeConnection()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnecteeDelegate(mockCtrl)
	mockDelegate.EXPECT().OnCtrlConnected(conn)
	mockDelegate.EXPECT().OnTargetConnected(conn)

	connectee := NewFakeConnectee()
	connectee.Start(mockDelegate)
	connectee.SetCtrlConnection(conn)
	connectee.SetTargetConnection(conn)
}

func Test_FakeConnectee_close_connection_when_not_started(t *testing.T) {
	log.Println("######## [Test_FakeConnectee_close_connection_when_not_started] ########")

	ctrl := NewFakeConnection()
	target := NewFakeConnection()

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
