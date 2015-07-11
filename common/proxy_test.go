package common_test

import (
	. "github.com/ghilbut/ygg.go/common"
	. "github.com/ghilbut/ygg.go/test/fake"
	. "github.com/ghilbut/ygg.go/test/mock/common"
	"github.com/golang/mock/gomock"
	"log"
	"testing"
)

const kJson = "{ \"id\": \"A\", \"endpoint\": \"B\" }"

func Test_CtrlProxy_return_instance_with_endpoint_value(t *testing.T) {
	log.Println("######## [Test_CtrlProxy_return_instance_with_endpoint_value] ########")

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

func Test_CtrlProxy_send_text(t *testing.T) {
	log.Println("######## [Test_CtrlProxy_send_text] ########")

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
	log.Println("######## [Test_CtrlProxy_recv_text] ########")

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

func Test_CtrlProxy_send_binary(t *testing.T) {
	log.Println("######## [Test_CtrlProxy_send_binary] ########")

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
	log.Println("######## [Test_CtrlProxy_recv_binary] ########")

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

func Test_CtrlProxy_close(t *testing.T) {
	log.Println("######## [Test_CtrlProxy_close] ########")

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
	log.Println("######## [Test_CtrlProxy_closed] ########")

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

func Test_TargetProxy_return_instance_with_endpoint_value(t *testing.T) {
	log.Println("######## [Test_TargetProxy_return_instance_with_endpoint_value] ########")

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

func Test_TargetProxy_send_text(t *testing.T) {
	log.Println("######## [Test_TargetProxy_send_text] ########")

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
	log.Println("######## [Test_TargetProxy_recv_text] ########")

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

func Test_TargetProxy_send_binary(t *testing.T) {
	log.Println("######## [Test_TargetProxy_send_binary] ########")

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
	log.Println("######## [Test_TargetProxy_recv_binary] ########")

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

func Test_TargetProxy_close(t *testing.T) {
	log.Println("######## [Test_TargetProxy_close] ########")

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
	log.Println("######## [Test_TargetProxy_closed] ########")

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
