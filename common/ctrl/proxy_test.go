package ctrl_test

import (
	. "github.com/ghilbut/ygg.go/common"
	. "github.com/ghilbut/ygg.go/common/ctrl"
	. "github.com/ghilbut/ygg.go/test/fake"
	. "github.com/ghilbut/ygg.go/test/mock"
	. "github.com/ghilbut/ygg.go/test/mock/common/ctrl"
	"github.com/golang/mock/gomock"
	"testing"
)

const kJson = "{ \"id\": \"A\", \"endpoint\": \"B\" }"

func Test_return_instance_with_endpoint_value(t *testing.T) {

	var lhs Connection = NewFakeConnection()
	if lhs == nil {
		t.Fail()
	}

	desc, _ := NewDesc(kJson)
	if desc == nil {
		t.Fail()
	}

	proxy := NewProxy(lhs.(*FakeConnection).Other(), desc)

	if proxy == nil {
		t.Fail()
	}

	if proxy.Delegate != nil {
		t.Fail()
	}
}

func Test_panic_when_connection_is_nil(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	desc, _ := NewDesc(kJson)
	if desc == nil {
		t.Fail()
	}

	NewProxy(nil, desc)
}

func Test_panic_when_desc_is_nil(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	var lhs Connection = NewFakeConnection()
	if lhs == nil {
		t.Fail()
	}

	NewProxy(lhs.(*FakeConnection).Other(), nil)
}

func Test_send_text(t *testing.T) {

	var lhs Connection = NewFakeConnection()
	desc, _ := NewDesc(kJson)
	proxy := NewProxy(lhs.(*FakeConnection).Other(), desc)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnText(lhs, "A")
	lhs.BindDelegate(mockDelegate)

	proxy.SendText("A")
}

func Test_recv_text(t *testing.T) {

	var lhs Connection = NewFakeConnection()
	desc, _ := NewDesc(kJson)
	proxy := NewProxy(lhs.(*FakeConnection).Other(), desc)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockProxyDelegate(mockCtrl)
	mockDelegate.EXPECT().OnText(proxy, "A").Times(2)

	proxy.Delegate = mockDelegate
	lhs.SendText("A")

	proxy.Delegate = nil
	lhs.SendText("A")

	proxy.Delegate = mockDelegate
	lhs.SendText("A")
}

func Test_panic_when_recv_text_from_external_connection(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	var lhs Connection = NewFakeConnection()
	var rhs Connection = lhs.(*FakeConnection).Other()
	desc, _ := NewDesc(kJson)
	proxy := NewProxy(rhs, desc)

	lhs.BindDelegate(proxy)
	rhs.SendText("A")
}

func Test_send_binary(t *testing.T) {

	var lhs Connection = NewFakeConnection()
	desc, _ := NewDesc(kJson)
	proxy := NewProxy(lhs.(*FakeConnection).Other(), desc)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnBinary(lhs, []byte{0x01, 0x02})
	lhs.BindDelegate(mockDelegate)

	proxy.SendBinary([]byte{0x01, 0x02})
}

func Test_recv_binary(t *testing.T) {

	var lhs Connection = NewFakeConnection()
	desc, _ := NewDesc(kJson)
	proxy := NewProxy(lhs.(*FakeConnection).Other(), desc)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockProxyDelegate(mockCtrl)
	mockDelegate.EXPECT().OnBinary(proxy, []byte{0x01, 0x02}).Times(2)

	proxy.Delegate = mockDelegate
	lhs.SendBinary([]byte{0x01, 0x02})

	proxy.Delegate = nil
	lhs.SendBinary([]byte{0x01, 0x02})

	proxy.Delegate = mockDelegate
	lhs.SendBinary([]byte{0x01, 0x02})
}

func Test_panic_when_recv_binary_from_external_connection(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	var lhs Connection = NewFakeConnection()
	var rhs Connection = lhs.(*FakeConnection).Other()
	desc, _ := NewDesc(kJson)
	proxy := NewProxy(rhs, desc)

	lhs.BindDelegate(proxy)
	rhs.SendBinary([]byte{0x01, 0x02})
}

func Test_close(t *testing.T) {

	var lhs Connection = NewFakeConnection()
	desc, _ := NewDesc(kJson)
	proxy := NewProxy(lhs.(*FakeConnection).Other(), desc)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnClosed(lhs)
	lhs.BindDelegate(mockDelegate)

	proxy.Close()
}

func Test_closed(t *testing.T) {

	var lhs Connection = NewFakeConnection()
	desc, _ := NewDesc(kJson)
	proxy := NewProxy(lhs.(*FakeConnection).Other(), desc)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockProxyDelegate(mockCtrl)
	mockDelegate.EXPECT().OnClosed(proxy).Times(2)

	proxy.Delegate = mockDelegate
	lhs.Close()

	proxy.Delegate = nil
	lhs.Close()

	proxy.Delegate = mockDelegate
	lhs.Close()
}

func Test_panic_when_closed_from_external_connection(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	var lhs Connection = NewFakeConnection()
	var rhs Connection = lhs.(*FakeConnection).Other()
	desc, _ := NewDesc(kJson)
	proxy := NewProxy(rhs, desc)

	lhs.BindDelegate(proxy)
	rhs.Close()
}
