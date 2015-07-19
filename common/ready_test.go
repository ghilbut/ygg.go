package common_test

import (
	. "github.com/ghilbut/ygg.go/common"
	. "github.com/ghilbut/ygg.go/net"
	. "github.com/ghilbut/ygg.go/test/mock/common"
	"github.com/golang/mock/gomock"
	"log"
	"testing"
)

//const kJson = "{ \"id\": \"A\", \"endpoint\": \"B\" }"

func Test_CtrlReady_has_connection_after_set_connection(t *testing.T) {
	log.Println("######## [Test_CtrlReady_has_connection_after_set_connection] ########")

	var conn Connection = NewLocalConnection()

	ready := NewCtrlReady()

	if ready.HasConnection(conn) {
		t.Fail()
	}

	ready.SetConnection(conn)

	if !ready.HasConnection(conn) {
		t.Fail()
	}
}

func Test_CtrlReady_clear_connection(t *testing.T) {
	log.Println("######## [Test_CtrlReady_clear_connection] ########")

	var conn0 Connection = NewLocalConnection()
	var conn1 Connection = NewLocalConnection()
	var conn2 Connection = NewLocalConnection()

	ready := NewCtrlReady()
	ready.SetConnection(conn0)
	ready.SetConnection(conn1)
	ready.SetConnection(conn2)

	if !ready.HasConnection(conn0) ||
		!ready.HasConnection(conn1) ||
		!ready.HasConnection(conn2) {
		t.Fail()
	}

	ready.Clear()

	if ready.HasConnection(conn0) ||
		ready.HasConnection(conn1) ||
		ready.HasConnection(conn2) {
		t.Fail()
	}
}

func Test_CtrlReady_remove_connection_when_it_is_closed(t *testing.T) {
	log.Println("######## [Test_CtrlReady_remove_connection_when_it_is_closed] ########")

	var conn Connection = NewLocalConnection()

	ready := NewCtrlReady()
	ready.SetConnection(conn)

	conn.Close()

	if ready.HasConnection(conn) {
		t.Fail()
	}
}

func Test_CtrlReady_remove_connection_when_invalid_json_is_passed(t *testing.T) {
	log.Println("######## [Test_CtrlReady_remove_connection_when_invalid_json_is_passed] ########")

	var lhs Connection = NewLocalConnection()
	var rhs = lhs.(*LocalConnection).Other()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockConnectionDelegate := NewMockConnectionDelegate(mockCtrl)
	mockConnectionDelegate.EXPECT().OnClosed(lhs).Times(3)
	lhs.BindDelegate(mockConnectionDelegate)

	mockCtrlReadyDelegate := NewMockCtrlReadyDelegate(mockCtrl)
	mockCtrlReadyDelegate.EXPECT().OnCtrlProxy(gomock.Any()).Times(0)

	ready := NewCtrlReady()
	ready.Delegate = mockCtrlReadyDelegate

	ready.SetConnection(rhs)
	lhs.SendText("")
	if ready.HasConnection(rhs) {
		t.Fail()
	}

	ready.SetConnection(rhs)
	lhs.SendText("{ \"key\": \"value")
	if ready.HasConnection(rhs) {
		t.Fail()
	}

	ready.SetConnection(rhs)
	lhs.SendText("{}")
	if ready.HasConnection(rhs) {
		t.Fail()
	}
}

type _matcher struct {
	endpoint string
}

func (self *_matcher) Matches(x interface{}) bool {
	if proxy, ok := x.(*CtrlProxy); ok {
		return proxy.Desc.Endpoint == self.endpoint
	}
	return false
}

func (self *_matcher) String() string {
	return "is ctrl proxy"
}

func Test_CtrlReady_ok(t *testing.T) {
	log.Println("######## [Test_CtrlReady_ok] ########")

	var lhs Connection = NewLocalConnection()
	var rhs Connection = lhs.(*LocalConnection).Other()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockCtrlReadyDelegate(mockCtrl)
	mockDelegate.EXPECT().OnCtrlProxy(&_matcher{"B"})

	ready := NewCtrlReady()
	ready.Delegate = mockDelegate
	ready.SetConnection(rhs)

	lhs.SendText(kJson)
}
