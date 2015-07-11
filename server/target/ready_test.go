package target_test

import (
	. "github.com/ghilbut/ygg.go/common"
	. "github.com/ghilbut/ygg.go/server/target"
	. "github.com/ghilbut/ygg.go/test/fake"
	. "github.com/ghilbut/ygg.go/test/mock/common"
	. "github.com/ghilbut/ygg.go/test/mock/server/target"
	"github.com/golang/mock/gomock"
	"log"
	"testing"
)

const kJson = "{ \"endpoint\": \"A\" }"

func Test_TargetReady_has_connection_after_set_connection(t *testing.T) {
	log.Println("######## [Test_TargetReady_has_connection_after_set_connection] ########")

	var conn Connection = NewFakeConnection()

	ready := NewTargetReady()

	if ready.HasConnection(conn) {
		t.Fail()
	}

	ready.SetConnection(conn)

	if !ready.HasConnection(conn) {
		t.Fail()
	}
}

func Test_TargetReady_clear_connection(t *testing.T) {
	log.Println("######## [Test_TargetReady_clear_connection] ########")

	var conn0 Connection = NewFakeConnection()
	var conn1 Connection = NewFakeConnection()
	var conn2 Connection = NewFakeConnection()

	ready := NewTargetReady()
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

func Test_TargetReady_remove_connection_when_it_is_closed(t *testing.T) {
	log.Println("######## [Test_TargetReady_remove_connection_when_it_is_closed] ########")

	var conn Connection = NewFakeConnection()

	ready := NewTargetReady()
	ready.SetConnection(conn)

	conn.Close()

	if ready.HasConnection(conn) {
		t.Fail()
	}
}

func Test_TargetReady_remove_connection_when_invalid_json_is_passed(t *testing.T) {
	log.Println("######## [Test_TargetReady_remove_connection_when_invalid_json_is_passed] ########")

	var lhs Connection = NewFakeConnection()
	var rhs = lhs.(*FakeConnection).Other()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockConnectionDelegate := NewMockConnectionDelegate(mockCtrl)
	mockConnectionDelegate.EXPECT().OnClosed(lhs).Times(3)
	lhs.BindDelegate(mockConnectionDelegate)

	mockTargetReadyDelegate := NewMockTargetReadyDelegate(mockCtrl)
	mockTargetReadyDelegate.EXPECT().OnTargetProxy(gomock.Any()).Times(0)

	ready := NewTargetReady()
	ready.Delegate = mockTargetReadyDelegate

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

	var lhs Connection = NewFakeConnection()
	var rhs Connection = lhs.(*FakeConnection).Other()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockTargetReadyDelegate(mockCtrl)
	mockDelegate.EXPECT().OnTargetProxy(&_matcher{"A"})

	ready := NewTargetReady()
	ready.Delegate = mockDelegate
	ready.SetConnection(rhs)

	lhs.SendText(kJson)
}
