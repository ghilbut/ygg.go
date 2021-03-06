package ctrl_test

import (
	. "github.com/ghilbut/ygg.go/common"
	. "github.com/ghilbut/ygg.go/net"
	. "github.com/ghilbut/ygg.go/server/ctrl"
	. "github.com/ghilbut/ygg.go/test/mock/common"
	"github.com/golang/mock/gomock"
	"log"
	"testing"
)

const kCtrlJson = "{ \"id\": \"A\", \"endpoint\": \"B\" }"
const kTargetJson = "{ \"endpoint\": \"B\" }"

const kText = "Message"

var kBytes = []byte{0x01, 0x02}

func Test_OneToOneAdapter_send_text(t *testing.T) {
	log.Println("######## [Test_OneToOneAdapter_send_text] ########")

	var ctrl Connection = NewLocalConnection()
	var target Connection = NewLocalConnection()

	cdesc, _ := NewCtrlDesc(kCtrlJson)
	cproxy := NewCtrlProxy(ctrl.(*LocalConnection).Other(), cdesc)

	tdesc, _ := NewTargetDesc(kTargetJson)
	tproxy := NewTargetProxy(target.(*LocalConnection).Other(), tdesc)

	adapter := NewOneToOneAdapter(tproxy)
	adapter.SetCtrlProxy(cproxy)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockConnDelegate := NewMockConnectionDelegate(mockCtrl)
	mockConnDelegate.EXPECT().OnText(ctrl, kText)
	ctrl.BindDelegate(mockConnDelegate)

	target.SendText(kText)
}

func Test_OneToOneAdapter_recv_text(t *testing.T) {
	log.Println("######## [Test_OneToOneAdapter_recv_text] ########")

	var ctrl Connection = NewLocalConnection()
	var target Connection = NewLocalConnection()

	cdesc, _ := NewCtrlDesc(kCtrlJson)
	cproxy := NewCtrlProxy(ctrl.(*LocalConnection).Other(), cdesc)

	tdesc, _ := NewTargetDesc(kTargetJson)
	tproxy := NewTargetProxy(target.(*LocalConnection).Other(), tdesc)

	adapter := NewOneToOneAdapter(tproxy)
	adapter.SetCtrlProxy(cproxy)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockConnDelegate := NewMockConnectionDelegate(mockCtrl)
	mockConnDelegate.EXPECT().OnText(target, kText)
	target.BindDelegate(mockConnDelegate)

	ctrl.SendText(kText)
}

func Test_OneToOneAdapter_send_binary(t *testing.T) {
	log.Println("######## [Test_OneToOneAdapter_send_text] ########")

	var ctrl Connection = NewLocalConnection()
	var target Connection = NewLocalConnection()

	cdesc, _ := NewCtrlDesc(kCtrlJson)
	cproxy := NewCtrlProxy(ctrl.(*LocalConnection).Other(), cdesc)

	tdesc, _ := NewTargetDesc(kTargetJson)
	tproxy := NewTargetProxy(target.(*LocalConnection).Other(), tdesc)

	adapter := NewOneToOneAdapter(tproxy)
	adapter.SetCtrlProxy(cproxy)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockConnDelegate := NewMockConnectionDelegate(mockCtrl)
	mockConnDelegate.EXPECT().OnBinary(ctrl, kBytes)
	ctrl.BindDelegate(mockConnDelegate)

	target.SendBinary(kBytes)
}

func Test_OneToOneAdapter_recv_binary(t *testing.T) {
	log.Println("######## [Test_OneToOneAdapter_recv_binary] ########")

	var ctrl Connection = NewLocalConnection()
	var target Connection = NewLocalConnection()

	cdesc, _ := NewCtrlDesc(kCtrlJson)
	cproxy := NewCtrlProxy(ctrl.(*LocalConnection).Other(), cdesc)

	tdesc, _ := NewTargetDesc(kTargetJson)
	tproxy := NewTargetProxy(target.(*LocalConnection).Other(), tdesc)

	adapter := NewOneToOneAdapter(tproxy)
	adapter.SetCtrlProxy(cproxy)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockConnDelegate := NewMockConnectionDelegate(mockCtrl)
	mockConnDelegate.EXPECT().OnBinary(target, kBytes)
	target.BindDelegate(mockConnDelegate)

	ctrl.SendBinary(kBytes)
}

func Test_OneToOneAdapter_close_ctrl(t *testing.T) {
	log.Println("######## [Test_OneToOneAdapter_close_ctrl] ########")

	var ctrl Connection = NewLocalConnection()
	var target Connection = NewLocalConnection()

	cdesc, _ := NewCtrlDesc(kCtrlJson)
	cproxy := NewCtrlProxy(ctrl.(*LocalConnection).Other(), cdesc)

	tdesc, _ := NewTargetDesc(kTargetJson)
	tproxy := NewTargetProxy(target.(*LocalConnection).Other(), tdesc)

	adapter := NewOneToOneAdapter(tproxy)
	adapter.SetCtrlProxy(cproxy)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockConnDelegate := NewMockConnectionDelegate(mockCtrl)
	mockConnDelegate.EXPECT().OnClosed(ctrl)
	mockConnDelegate.EXPECT().OnClosed(target)
	ctrl.BindDelegate(mockConnDelegate)
	target.BindDelegate(mockConnDelegate)

	mockAdapterDelegate := NewMockAdapterDelegate(mockCtrl)
	mockAdapterDelegate.EXPECT().OnAdapterClosed(adapter)
	adapter.BindDelegate(mockAdapterDelegate)

	ctrl.Close()
}

func Test_OneToOneAdapter_close_target(t *testing.T) {
	log.Println("######## [Test_OneToOneAdapter_close_target] ########")

	var ctrl Connection = NewLocalConnection()
	var target Connection = NewLocalConnection()

	cdesc, _ := NewCtrlDesc(kCtrlJson)
	cproxy := NewCtrlProxy(ctrl.(*LocalConnection).Other(), cdesc)

	tdesc, _ := NewTargetDesc(kTargetJson)
	tproxy := NewTargetProxy(target.(*LocalConnection).Other(), tdesc)

	adapter := NewOneToOneAdapter(tproxy)
	adapter.SetCtrlProxy(cproxy)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockConnDelegate := NewMockConnectionDelegate(mockCtrl)
	mockConnDelegate.EXPECT().OnClosed(ctrl)
	mockConnDelegate.EXPECT().OnClosed(target)
	ctrl.BindDelegate(mockConnDelegate)
	target.BindDelegate(mockConnDelegate)

	mockAdapterDelegate := NewMockAdapterDelegate(mockCtrl)
	mockAdapterDelegate.EXPECT().OnAdapterClosed(adapter)
	adapter.BindDelegate(mockAdapterDelegate)

	target.Close()
}

func Test_OneToOneAdapter_close_adapter(t *testing.T) {
	log.Println("######## [Test_OneToOneAdapter_close_adapter] ########")

	var ctrl Connection = NewLocalConnection()
	var target Connection = NewLocalConnection()

	cdesc, _ := NewCtrlDesc(kCtrlJson)
	cproxy := NewCtrlProxy(ctrl.(*LocalConnection).Other(), cdesc)

	tdesc, _ := NewTargetDesc(kTargetJson)
	tproxy := NewTargetProxy(target.(*LocalConnection).Other(), tdesc)

	adapter := NewOneToOneAdapter(tproxy)
	adapter.SetCtrlProxy(cproxy)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockConnDelegate := NewMockConnectionDelegate(mockCtrl)
	mockConnDelegate.EXPECT().OnClosed(ctrl)
	mockConnDelegate.EXPECT().OnClosed(target)
	ctrl.BindDelegate(mockConnDelegate)
	target.BindDelegate(mockConnDelegate)

	mockAdapterDelegate := NewMockAdapterDelegate(mockCtrl)
	mockAdapterDelegate.EXPECT().OnAdapterClosed(adapter)
	adapter.BindDelegate(mockAdapterDelegate)

	adapter.Close()
}
