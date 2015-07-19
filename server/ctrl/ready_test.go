package ctrl_test

import (
	. "github.com/ghilbut/ygg.go/common"
	. "github.com/ghilbut/ygg.go/net"
	. "github.com/ghilbut/ygg.go/server/ctrl"
	. "github.com/ghilbut/ygg.go/test/mock/common"
	. "github.com/ghilbut/ygg.go/test/mock/server/ctrl"
	"github.com/golang/mock/gomock"
	"log"
	"testing"
)

//const kCtrlJson = "{ \"id\": \"A\", \"endpoint\": \"B\" }"
//const kTargetJson = "{ \"endpoint\": \"B\" }"

func Test_TargetReady_has_connection_after_set_connection(t *testing.T) {
	log.Println("######## [Test_TargetReady_has_connection_after_set_connection] ########")

	var conn Connection = NewLocalConnection()

	desc, _ := NewCtrlDesc(kCtrlJson)
	cproxy := NewCtrlProxy(conn, desc)

	ready := NewTargetReady()

	if ready.HasConnection(conn) {
		t.Fail()
	}

	ready.SetConnection(conn, cproxy)

	if !ready.HasConnection(conn) {
		t.Fail()
	}
}

func Test_TargetReady_clear_connection(t *testing.T) {
	log.Println("######## [Test_TargetReady_clear_connection] ########")

	var conn0 Connection = NewLocalConnection()
	var conn1 Connection = NewLocalConnection()
	var conn2 Connection = NewLocalConnection()

	desc, _ := NewCtrlDesc(kCtrlJson)
	cproxy0 := NewCtrlProxy(conn0.(*LocalConnection).Other(), desc)
	cproxy1 := NewCtrlProxy(conn1.(*LocalConnection).Other(), desc)
	cproxy2 := NewCtrlProxy(conn2.(*LocalConnection).Other(), desc)

	ready := NewTargetReady()
	ready.SetConnection(conn0, cproxy0)
	ready.SetConnection(conn1, cproxy1)
	ready.SetConnection(conn2, cproxy2)

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

func Test_TargetReady_remove_connection_when_it_is_closed(t *testing.T) {
	log.Println("######## [Test_TargetReady_remove_connection_when_it_is_closed] ########")

	var conn Connection = NewLocalConnection()

	desc, _ := NewCtrlDesc(kCtrlJson)
	cproxy0 := NewCtrlProxy(conn.(*LocalConnection).Other(), desc)

	ready := NewTargetReady()
	ready.SetConnection(conn, cproxy0)

	conn.Close()

	if ready.HasConnection(conn) {
		t.Fail()
	}
}

func Test_TargetReady_remove_connection_when_invalid_json_is_passed(t *testing.T) {
	log.Println("######## [Test_TargetReady_remove_connection_when_invalid_json_is_passed] ########")

	var lhs Connection = NewLocalConnection()
	var rhs = lhs.(*LocalConnection).Other()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockConnectionDelegate := NewMockConnectionDelegate(mockCtrl)
	mockConnectionDelegate.EXPECT().OnClosed(lhs).Times(3)
	lhs.BindDelegate(mockConnectionDelegate)

	mockTargetReadyDelegate := NewMockTargetReadyDelegate(mockCtrl)
	mockTargetReadyDelegate.EXPECT().OnTargetProxy(gomock.Any(), gomock.Any()).Times(0)

	desc, _ := NewCtrlDesc(kCtrlJson)
	cproxy := NewCtrlProxy(NewLocalConnection(), desc)

	ready := NewTargetReady()
	ready.Delegate = mockTargetReadyDelegate

	ready.SetConnection(rhs, cproxy)
	lhs.SendText("")
	if ready.HasConnection(rhs) {
		t.Fail()
	}

	ready.SetConnection(rhs, cproxy)
	lhs.SendText("{ \"key\": \"value")
	if ready.HasConnection(rhs) {
		t.Fail()
	}

	ready.SetConnection(rhs, cproxy)
	lhs.SendText("{}")
	if ready.HasConnection(rhs) {
		t.Fail()
	}
}

type _matcher struct {
	endpoint string
}

func (self *_matcher) Matches(x interface{}) bool {
	if proxy, ok := x.(*TargetProxy); ok {
		return proxy.Desc.Endpoint == self.endpoint
	}
	return false
}

func (self *_matcher) String() string {
	return "is target proxy"
}

func Test_TargetReady_ok(t *testing.T) {
	log.Println("######## [Test_TargetReady_ok] ########")

	var lhs Connection = NewLocalConnection()
	var rhs Connection = lhs.(*LocalConnection).Other()

	desc, _ := NewCtrlDesc(kCtrlJson)
	cproxy := NewCtrlProxy(NewLocalConnection(), desc)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockTargetReadyDelegate(mockCtrl)
	mockDelegate.EXPECT().OnTargetProxy(cproxy, &_matcher{"B"})

	ready := NewTargetReady()
	ready.Delegate = mockDelegate
	ready.SetConnection(rhs, cproxy)

	lhs.SendText(kTargetJson)
}
