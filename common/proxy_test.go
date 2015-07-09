package common_test

import (
	. "github.com/ghilbut/ygg.go/common"
	. "github.com/ghilbut/ygg.go/test/fake"
	. "github.com/ghilbut/ygg.go/test/mock/common"
	"github.com/golang/mock/gomock"
	"testing"
)

const kJson = "{ \"id\": \"A\", \"endpoint\": \"B\" }"

func Test_CtrlProxy_return_instance_with_endpoint_value(t *testing.T) {

	var lhs Connection = NewFakeConnection()
	if lhs == nil {
		t.Fail()
	}

	desc, _ := NewCtrlDesc(kJson)
	if desc == nil {
		t.Fail()
	}

	proxy := NewCtrlProxy(lhs.(*FakeConnection).Other(), desc)

	if proxy == nil {
		t.Fail()
	}

	if proxy.Delegate != nil {
		t.Fail()
	}
}

func Test_CtrlProxy_panic_when_connection_is_nil(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	desc, _ := NewCtrlDesc(kJson)
	if desc == nil {
		t.Fail()
	}

	NewCtrlProxy(nil, desc)
}

func Test_CtrlProxy_panic_when_desc_is_nil(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	var lhs Connection = NewFakeConnection()
	if lhs == nil {
		t.Fail()
	}

	NewCtrlProxy(lhs.(*FakeConnection).Other(), nil)
}

func Test_CtrlProxy_send_text(t *testing.T) {

	var lhs Connection = NewFakeConnection()
	desc, _ := NewCtrlDesc(kJson)
	proxy := NewCtrlProxy(lhs.(*FakeConnection).Other(), desc)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnText(lhs, "A")
	lhs.BindDelegate(mockDelegate)

	proxy.SendText("A")
}

func Test_CtrlProxy_recv_text(t *testing.T) {

	var lhs Connection = NewFakeConnection()
	desc, _ := NewCtrlDesc(kJson)
	proxy := NewCtrlProxy(lhs.(*FakeConnection).Other(), desc)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockCtrlProxyDelegate(mockCtrl)
	mockDelegate.EXPECT().OnCtrlText(proxy, "A").Times(2)

	proxy.Delegate = mockDelegate
	lhs.SendText("A")

	proxy.Delegate = nil
	lhs.SendText("A")

	proxy.Delegate = mockDelegate
	lhs.SendText("A")
}

func Test_CtrlProxy_panic_when_recv_text_from_external_connection(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	var lhs Connection = NewFakeConnection()
	var rhs Connection = lhs.(*FakeConnection).Other()
	desc, _ := NewCtrlDesc(kJson)
	proxy := NewCtrlProxy(rhs, desc)

	lhs.BindDelegate(proxy)
	rhs.SendText("A")
}

func Test_CtrlProxy_send_binary(t *testing.T) {

	var lhs Connection = NewFakeConnection()
	desc, _ := NewCtrlDesc(kJson)
	proxy := NewCtrlProxy(lhs.(*FakeConnection).Other(), desc)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnBinary(lhs, []byte{0x01, 0x02})
	lhs.BindDelegate(mockDelegate)

	proxy.SendBinary([]byte{0x01, 0x02})
}

func Test_CtrlProxy_recv_binary(t *testing.T) {

	var lhs Connection = NewFakeConnection()
	desc, _ := NewCtrlDesc(kJson)
	proxy := NewCtrlProxy(lhs.(*FakeConnection).Other(), desc)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockCtrlProxyDelegate(mockCtrl)
	mockDelegate.EXPECT().OnCtrlBinary(proxy, []byte{0x01, 0x02}).Times(2)

	proxy.Delegate = mockDelegate
	lhs.SendBinary([]byte{0x01, 0x02})

	proxy.Delegate = nil
	lhs.SendBinary([]byte{0x01, 0x02})

	proxy.Delegate = mockDelegate
	lhs.SendBinary([]byte{0x01, 0x02})
}

func Test_CtrlProxy_panic_when_recv_binary_from_external_connection(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	var lhs Connection = NewFakeConnection()
	var rhs Connection = lhs.(*FakeConnection).Other()
	desc, _ := NewCtrlDesc(kJson)
	proxy := NewCtrlProxy(rhs, desc)

	lhs.BindDelegate(proxy)
	rhs.SendBinary([]byte{0x01, 0x02})
}

func Test_CtrlProxy_close(t *testing.T) {

	var lhs Connection = NewFakeConnection()
	desc, _ := NewCtrlDesc(kJson)
	proxy := NewCtrlProxy(lhs.(*FakeConnection).Other(), desc)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnClosed(lhs)
	lhs.BindDelegate(mockDelegate)

	proxy.Close()
}

func Test_CtrlProxy_closed(t *testing.T) {

	var lhs Connection = NewFakeConnection()
	desc, _ := NewCtrlDesc(kJson)
	proxy := NewCtrlProxy(lhs.(*FakeConnection).Other(), desc)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockCtrlProxyDelegate(mockCtrl)
	mockDelegate.EXPECT().OnCtrlClosed(proxy).Times(2)

	proxy.Delegate = mockDelegate
	lhs.Close()

	proxy.Delegate = nil
	lhs.Close()

	proxy.Delegate = mockDelegate
	lhs.Close()
}

func Test_CtrlProxy_panic_when_closed_from_external_connection(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	var lhs Connection = NewFakeConnection()
	var rhs Connection = lhs.(*FakeConnection).Other()
	desc, _ := NewCtrlDesc(kJson)
	proxy := NewCtrlProxy(rhs, desc)

	lhs.BindDelegate(proxy)
	rhs.Close()
}

func Test_TargetProxy_return_instance_with_endpoint_value(t *testing.T) {

	desc, _ := NewTargetDesc(kJson)
	if desc == nil {
		t.Fail()
	}

	var lhs Connection = NewFakeConnection()
	if lhs == nil {
		t.Fail()
	}

	proxy := NewTargetProxy(lhs.(*FakeConnection).Other(), desc)

	if proxy == nil {
		t.Fail()
	}

	if proxy.Delegate != nil {
		t.Fail()
	}
}

func Test_TargetProxy_panic_when_connection_is_nil(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	desc, _ := NewTargetDesc(kJson)
	if desc == nil {
		t.Fail()
	}

	NewTargetProxy(nil, desc)
}

func Test_TargetProxy_panic_when_desc_is_nil(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	var lhs Connection = NewFakeConnection()
	if lhs == nil {
		t.Fail()
	}

	NewTargetProxy(lhs.(*FakeConnection).Other(), nil)
}

func Test_TargetProxy_send_text(t *testing.T) {

	var lhs Connection = NewFakeConnection()
	desc, _ := NewTargetDesc(kJson)
	proxy := NewTargetProxy(lhs.(*FakeConnection).Other(), desc)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnText(lhs, "A")
	lhs.BindDelegate(mockDelegate)

	proxy.SendText("A")
}

func Test_TargetProxy_recv_text(t *testing.T) {

	var lhs Connection = NewFakeConnection()
	desc, _ := NewTargetDesc(kJson)
	proxy := NewTargetProxy(lhs.(*FakeConnection).Other(), desc)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockTargetProxyDelegate(mockCtrl)
	mockDelegate.EXPECT().OnTargetText(proxy, "A").Times(2)

	proxy.Delegate = mockDelegate
	lhs.SendText("A")

	proxy.Delegate = nil
	lhs.SendText("A")

	proxy.Delegate = mockDelegate
	lhs.SendText("A")
}

func Test_TargetProxy_panic_when_recv_text_from_external_connection(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	var lhs Connection = NewFakeConnection()
	var rhs Connection = lhs.(*FakeConnection).Other()
	desc, _ := NewTargetDesc(kJson)
	proxy := NewTargetProxy(rhs, desc)

	lhs.BindDelegate(proxy)
	rhs.SendText("A")
}

func Test_TargetProxy_send_binary(t *testing.T) {

	var lhs Connection = NewFakeConnection()
	desc, _ := NewTargetDesc(kJson)
	proxy := NewTargetProxy(lhs.(*FakeConnection).Other(), desc)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnBinary(lhs, []byte{0x01, 0x02})
	lhs.BindDelegate(mockDelegate)

	proxy.SendBinary([]byte{0x01, 0x02})
}

func Test_TargetProxy_recv_binary(t *testing.T) {

	var lhs Connection = NewFakeConnection()
	desc, _ := NewTargetDesc(kJson)
	proxy := NewTargetProxy(lhs.(*FakeConnection).Other(), desc)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockTargetProxyDelegate(mockCtrl)
	mockDelegate.EXPECT().OnTargetBinary(proxy, []byte{0x01, 0x02}).Times(2)

	proxy.Delegate = mockDelegate
	lhs.SendBinary([]byte{0x01, 0x02})

	proxy.Delegate = nil
	lhs.SendBinary([]byte{0x01, 0x02})

	proxy.Delegate = mockDelegate
	lhs.SendBinary([]byte{0x01, 0x02})
}

func Test_TargetProxy_panic_when_recv_binary_from_external_connection(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	var lhs Connection = NewFakeConnection()
	var rhs Connection = lhs.(*FakeConnection).Other()
	desc, _ := NewTargetDesc(kJson)
	proxy := NewTargetProxy(rhs, desc)

	lhs.BindDelegate(proxy)
	rhs.SendBinary([]byte{0x01, 0x02})
}

func Test_TargetProxy_close(t *testing.T) {

	var lhs Connection = NewFakeConnection()
	desc, _ := NewTargetDesc(kJson)
	proxy := NewTargetProxy(lhs.(*FakeConnection).Other(), desc)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnClosed(lhs)
	lhs.BindDelegate(mockDelegate)

	proxy.Close()
}

func Test_TargetProxy_closed(t *testing.T) {

	var lhs Connection = NewFakeConnection()
	desc, _ := NewTargetDesc(kJson)
	proxy := NewTargetProxy(lhs.(*FakeConnection).Other(), desc)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockTargetProxyDelegate(mockCtrl)
	mockDelegate.EXPECT().OnTargetClosed(proxy).Times(2)

	proxy.Delegate = mockDelegate
	lhs.Close()

	proxy.Delegate = nil
	lhs.Close()

	proxy.Delegate = mockDelegate
	lhs.Close()
}

func Test_TargetProxy_panic_when_closed_from_external_connection(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	var lhs Connection = NewFakeConnection()
	var rhs Connection = lhs.(*FakeConnection).Other()
	desc, _ := NewTargetDesc(kJson)
	proxy := NewTargetProxy(rhs, desc)

	lhs.BindDelegate(proxy)
	rhs.Close()
}
