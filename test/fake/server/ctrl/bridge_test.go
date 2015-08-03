package fake_test

import (
	. "github.com/ghilbut/ygg.go/common"
	. "github.com/ghilbut/ygg.go/net"
	. "github.com/ghilbut/ygg.go/test/fake/server/ctrl"
	. "github.com/ghilbut/ygg.go/test/mock/common"
	. "github.com/ghilbut/ygg.go/test/mock/server/ctrl"
	"github.com/golang/mock/gomock"
	"log"
	"testing"
)

const kCtrlJson = "{ \"id\" : \"A\", \"endpoint\" : \"B\" }"

func Test_FakeCtrlBridge_delegate_connection_when_set_connection(t *testing.T) {
	log.Println("######## [Test_ctrl.FakeCtrlBridge_delegate_connection_when_set_connection] ########")

	conn := NewLocalConnection()

	bridge := NewFakeCtrlBridge()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockCtrlBridgeDelegate(mockCtrl)
	mockDelegate.EXPECT().OnCtrlBridgeStarted(bridge)
	mockDelegate.EXPECT().OnCtrlConnected(conn)

	bridge.Delegate = mockDelegate
	bridge.Start()

	bridge.SetCtrlConnection(conn)
}

func Test_FakeCtrlBridge_close_connection_when_not_started(t *testing.T) {
	log.Println("######## [Test_ctrl.FakeCtrlBridge_close_connection_when_not_started] ########")

	ctrl := NewLocalConnection()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnClosed(ctrl)
	ctrl.BindDelegate(mockDelegate)

	bridge := NewFakeCtrlBridge()
	bridge.SetCtrlConnection(ctrl)
}

func Test_FakeCtrlBridge_close_old_connection_if_set_same_endpoint_again(t *testing.T) {
	log.Println("######## [Test_FakeCtrlBridge_close_old_connection_if_set_same_endpoint_again] ########")

	conn0 := NewLocalConnection()
	conn1 := NewLocalConnection()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnClosed(conn0)
	mockDelegate.EXPECT().OnClosed(conn1).Times(0)
	conn0.BindDelegate(mockDelegate)
	conn1.BindDelegate(mockDelegate)

	bridge := NewFakeCtrlBridge()
	bridge.SetTargetConnection("B", conn0)
	bridge.SetTargetConnection("B", conn1)
}

func Test_FakeCtrlBridge_close_ctrl_and_return_nil_when_there_is_no_matching_endpoint(t *testing.T) {
	log.Println("######## [Test_FakeCtrlBridge_close_ctrl_and_return_nil_when_there_is_no_matching_endpoint] ########")

	conn := NewLocalConnection()

	desc, _ := NewCtrlDesc(kCtrlJson)
	ctrl := NewCtrlProxy(conn, desc)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockCtrlProxyDelegate(mockCtrl)
	mockDelegate.EXPECT().OnCtrlClosed(ctrl)
	ctrl.Delegate = mockDelegate

	bridge := NewFakeCtrlBridge()

	if bridge.Connect(ctrl) != nil {
		t.Fail()
	}
}

func Test_FakeCtrlBridge_return_connection_when_there_is_matching_endpoint(t *testing.T) {
	log.Println("######## [Test_FakeCtrlBridge_return_connection_when_there_is_matching_endpoint] ########")

	conn := NewLocalConnection()

	bridge := NewFakeCtrlBridge()
	bridge.SetTargetConnection("B", conn)

	desc, _ := NewCtrlDesc(kCtrlJson)
	ctrl := NewCtrlProxy(conn, desc)

	if bridge.Connect(ctrl) != conn {
		t.Fail()
	}
}
