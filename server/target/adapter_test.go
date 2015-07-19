package target_test

import (
	. "github.com/ghilbut/ygg.go/common"
	. "github.com/ghilbut/ygg.go/net"
	. "github.com/ghilbut/ygg.go/server/target"
	. "github.com/ghilbut/ygg.go/test/mock/common"
	"github.com/golang/mock/gomock"
	"log"
	"testing"
)

const kCtrlA0Json = "{ \"id\": \"A0\", \"endpoint\": \"B\" }"
const kCtrlA1Json = "{ \"id\": \"A1\", \"endpoint\": \"B\" }"
const kTargetJson = "{ \"endpoint\": \"B\" }"

const kText = "Message"

var kBytes = []byte{0x01, 0x02}

func Test_ManyToOneAdapter_target_notify_text_to_ctrls(t *testing.T) {
	log.Println("######## [Test_ManyToOneAdapter_target_notify_text_to_ctrls] ########")

	var lhs0 Connection = NewLocalConnection()
	var lhs1 Connection = NewLocalConnection()
	var rhs Connection = NewLocalConnection()

	c0desc, _ := NewCtrlDesc(kCtrlA0Json)
	c0proxy := NewCtrlProxy(lhs0.(*LocalConnection).Other(), c0desc)
	c1desc, _ := NewCtrlDesc(kCtrlA1Json)
	c1proxy := NewCtrlProxy(lhs1.(*LocalConnection).Other(), c1desc)

	tdesc, _ := NewTargetDesc(kTargetJson)
	tproxy := NewTargetProxy(rhs.(*LocalConnection).Other(), tdesc)

	var adapter Adapter = NewManyToOneAdapter(tproxy)
	adapter.SetCtrlProxy(c0proxy)
	adapter.SetCtrlProxy(c1proxy)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	const text = "Message"
	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnText(lhs0, text)
	mockDelegate.EXPECT().OnText(lhs1, text)

	lhs0.BindDelegate(mockDelegate)
	lhs1.BindDelegate(mockDelegate)
	rhs.SendText(text)
}

func Test_ManyToOneAdapter_target_recv_text_from_ctrls(t *testing.T) {
	log.Println("######## [Test_ManyToOneAdapter_target_recv_text_from_ctrls] ########")

	var lhs0 Connection = NewLocalConnection()
	var lhs1 Connection = NewLocalConnection()
	var rhs Connection = NewLocalConnection()

	c0desc, _ := NewCtrlDesc(kCtrlA0Json)
	c0proxy := NewCtrlProxy(lhs0.(*LocalConnection).Other(), c0desc)
	c1desc, _ := NewCtrlDesc(kCtrlA1Json)
	c1proxy := NewCtrlProxy(lhs1.(*LocalConnection).Other(), c1desc)

	tdesc, _ := NewTargetDesc(kTargetJson)
	tproxy := NewTargetProxy(rhs.(*LocalConnection).Other(), tdesc)

	var adapter Adapter = NewManyToOneAdapter(tproxy)
	adapter.SetCtrlProxy(c0proxy)
	adapter.SetCtrlProxy(c1proxy)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	const text = "Message"
	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnText(rhs, text).Times(2)

	rhs.BindDelegate(mockDelegate)
	lhs0.SendText(text)
	lhs1.SendText(text)
}

func Test_ManyToOneAdapter_target_notify_binary_to_ctrls(t *testing.T) {
	log.Println("######## [Test_ManyToOneAdapter_target_notify_binary_to_ctrls] ########")

	var lhs0 Connection = NewLocalConnection()
	var lhs1 Connection = NewLocalConnection()
	var rhs Connection = NewLocalConnection()

	c0desc, _ := NewCtrlDesc(kCtrlA0Json)
	c0proxy := NewCtrlProxy(lhs0.(*LocalConnection).Other(), c0desc)
	c1desc, _ := NewCtrlDesc(kCtrlA1Json)
	c1proxy := NewCtrlProxy(lhs1.(*LocalConnection).Other(), c1desc)

	tdesc, _ := NewTargetDesc(kTargetJson)
	tproxy := NewTargetProxy(rhs.(*LocalConnection).Other(), tdesc)

	var adapter Adapter = NewManyToOneAdapter(tproxy)
	adapter.SetCtrlProxy(c0proxy)
	adapter.SetCtrlProxy(c1proxy)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	const text = "Message"
	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnBinary(lhs0, kBytes)
	mockDelegate.EXPECT().OnBinary(lhs1, kBytes)

	lhs0.BindDelegate(mockDelegate)
	lhs1.BindDelegate(mockDelegate)
	rhs.SendBinary(kBytes)
}

func Test_ManyToOneAdapter_target_recv_binary_from_ctrls(t *testing.T) {
	log.Println("######## [Test_ManyToOneAdapter_target_recv_binary_from_ctrls] ########")

	var lhs0 Connection = NewLocalConnection()
	var lhs1 Connection = NewLocalConnection()
	var rhs Connection = NewLocalConnection()

	c0desc, _ := NewCtrlDesc(kCtrlA0Json)
	c0proxy := NewCtrlProxy(lhs0.(*LocalConnection).Other(), c0desc)
	c1desc, _ := NewCtrlDesc(kCtrlA1Json)
	c1proxy := NewCtrlProxy(lhs1.(*LocalConnection).Other(), c1desc)

	tdesc, _ := NewTargetDesc(kTargetJson)
	tproxy := NewTargetProxy(rhs.(*LocalConnection).Other(), tdesc)

	var adapter Adapter = NewManyToOneAdapter(tproxy)
	adapter.SetCtrlProxy(c0proxy)
	adapter.SetCtrlProxy(c1proxy)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	const text = "Message"
	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnBinary(rhs, kBytes).Times(2)

	rhs.BindDelegate(mockDelegate)
	lhs0.SendBinary(kBytes)
	lhs1.SendBinary(kBytes)
}

func Test_ManyToOneAdapter_close_all_ctrls_when_target_closed(t *testing.T) {
	log.Println("######## [Test_ManyToOneAdapter_close_all_ctrls_when_target_closed] ########")

	var lhs0 Connection = NewLocalConnection()
	var lhs1 Connection = NewLocalConnection()
	var rhs Connection = NewLocalConnection()

	c0desc, _ := NewCtrlDesc(kCtrlA0Json)
	c0proxy := NewCtrlProxy(lhs0.(*LocalConnection).Other(), c0desc)
	c1desc, _ := NewCtrlDesc(kCtrlA1Json)
	c1proxy := NewCtrlProxy(lhs1.(*LocalConnection).Other(), c1desc)

	tdesc, _ := NewTargetDesc(kTargetJson)
	tproxy := NewTargetProxy(rhs.(*LocalConnection).Other(), tdesc)

	var adapter Adapter = NewManyToOneAdapter(tproxy)
	adapter.SetCtrlProxy(c0proxy)
	adapter.SetCtrlProxy(c1proxy)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	const text = "Message"
	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnClosed(lhs0)
	mockDelegate.EXPECT().OnClosed(lhs1)

	lhs0.BindDelegate(mockDelegate)
	lhs1.BindDelegate(mockDelegate)
	rhs.Close()
}

func Test_ManyToOneAdapter_remove_ctrl_proxy_after_closed(t *testing.T) {
	log.Println("######## [Test_ManyToOneAdapter_remove_ctrl_proxy_after_closed] ########")

	var lhs0 Connection = NewLocalConnection()
	var lhs1 Connection = NewLocalConnection()
	var rhs Connection = NewLocalConnection()

	c0desc, _ := NewCtrlDesc(kCtrlA0Json)
	c0proxy := NewCtrlProxy(lhs0.(*LocalConnection).Other(), c0desc)
	c1desc, _ := NewCtrlDesc(kCtrlA1Json)
	c1proxy := NewCtrlProxy(lhs1.(*LocalConnection).Other(), c1desc)

	tdesc, _ := NewTargetDesc(kTargetJson)
	tproxy := NewTargetProxy(rhs.(*LocalConnection).Other(), tdesc)

	adapter := NewManyToOneAdapter(tproxy)
	adapter.SetCtrlProxy(c0proxy)
	adapter.SetCtrlProxy(c1proxy)

	if !adapter.HasCtrlProxy(c0proxy) {
		t.Fail()
	}

	c0proxy.Close()
	if adapter.HasCtrlProxy(c0proxy) {
		t.Fail()
	}

	if !adapter.HasCtrlProxy(c1proxy) {
		t.Fail()
	}

	c1proxy.Close()
	if adapter.HasCtrlProxy(c1proxy) {
		t.Fail()
	}
}
