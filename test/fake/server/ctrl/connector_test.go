package fake_test

import (
	. "github.com/ghilbut/ygg.go/common"
	. "github.com/ghilbut/ygg.go/net"
	. "github.com/ghilbut/ygg.go/test/fake/server/ctrl"
	. "github.com/ghilbut/ygg.go/test/mock/common"
	"github.com/golang/mock/gomock"
	"log"
	"testing"
)

const kCtrlJson = "{ \"id\" : \"A\", \"endpoint\" : \"B\" }"

func Test_FakeConnector_close_old_connection_if_set_same_endpoint_again(t *testing.T) {
	log.Println("######## [Test_FakeConnector_close_old_connection_if_set_same_endpoint_again] ########")

	conn0 := NewLocalConnection()
	conn1 := NewLocalConnection()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnClosed(conn0)
	mockDelegate.EXPECT().OnClosed(conn1).Times(0)
	conn0.BindDelegate(mockDelegate)
	conn1.BindDelegate(mockDelegate)

	connector := NewFakeConnector()
	connector.SetTargetConnection("B", conn0)
	connector.SetTargetConnection("B", conn1)
}

func Test_FakeConnector_close_ctrl_and_return_nil_when_there_is_no_matching_endpoint(t *testing.T) {
	log.Println("######## [Test_FakeConnector_close_ctrl_and_return_nil_when_there_is_no_matching_endpoint] ########")

	conn := NewLocalConnection()

	desc, _ := NewCtrlDesc(kCtrlJson)
	ctrl := NewCtrlProxy(conn, desc)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockCtrlProxyDelegate(mockCtrl)
	mockDelegate.EXPECT().OnCtrlClosed(ctrl)
	ctrl.Delegate = mockDelegate

	connector := NewFakeConnector()

	if connector.Connect(ctrl) != nil {
		t.Fail()
	}
}

func Test_FakeConnector_return_connection_when_there_is_matching_endpoint(t *testing.T) {
	log.Println("######## [Test_FakeConnector_return_connection_when_there_is_matching_endpoint] ########")

	conn := NewLocalConnection()

	connector := NewFakeConnector()
	connector.SetTargetConnection("B", conn)

	desc, _ := NewCtrlDesc(kCtrlJson)
	ctrl := NewCtrlProxy(conn, desc)

	if connector.Connect(ctrl) != conn {
		t.Fail()
	}
}
