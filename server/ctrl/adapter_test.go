package ctrl_test

import (
	. "github.com/ghilbut/ygg.go/common"
	. "github.com/ghilbut/ygg.go/server/ctrl"
	. "github.com/ghilbut/ygg.go/test/fake"
	. "github.com/ghilbut/ygg.go/test/mock/common"
	"github.com/golang/mock/gomock"
	"log"
	"testing"
)

const kCtrlJson = "{ \"id\": \"A\", \"endpoint\": \"B\" }"
const kTargetJson = "{ \"endpoint\": \"B\" }"

const kText = "Message"

var kBytes = []byte{0x01, 0x02}

func Test_OneToOneAdapter_close_ctrl(t *testing.T) {
	log.Println("######## [Test_OneToOneAdapter_close_ctrl] ########")

	var ctrl Connection = NewFakeConnection()
	var target Connection = NewFakeConnection()

	cdesc, _ := NewCtrlDesc(kCtrlJson)
	cproxy := NewCtrlProxy(ctrl.(*FakeConnection).Other(), cdesc)

	tdesc, _ := NewTargetDesc(kTargetJson)
	tproxy := NewTargetProxy(target.(*FakeConnection).Other(), tdesc)

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
	mockAdapterDelegate.EXPECT().OnClosed(adapter)
	adapter.BindDelegate(mockAdapterDelegate)

	ctrl.Close()
}

func Test_OneToOneAdapter_close_target(t *testing.T) {
	log.Println("######## [Test_OneToOneAdapter_close_target] ########")

	var ctrl Connection = NewFakeConnection()
	var target Connection = NewFakeConnection()

	cdesc, _ := NewCtrlDesc(kCtrlJson)
	cproxy := NewCtrlProxy(ctrl.(*FakeConnection).Other(), cdesc)

	tdesc, _ := NewTargetDesc(kTargetJson)
	tproxy := NewTargetProxy(target.(*FakeConnection).Other(), tdesc)

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
	mockAdapterDelegate.EXPECT().OnClosed(adapter)
	adapter.BindDelegate(mockAdapterDelegate)

	target.Close()
}

func Test_OneToOneAdapter_close_adapter(t *testing.T) {
	log.Println("######## [Test_OneToOneAdapter_close_adapter] ########")

	var ctrl Connection = NewFakeConnection()
	var target Connection = NewFakeConnection()

	cdesc, _ := NewCtrlDesc(kCtrlJson)
	cproxy := NewCtrlProxy(ctrl.(*FakeConnection).Other(), cdesc)

	tdesc, _ := NewTargetDesc(kTargetJson)
	tproxy := NewTargetProxy(target.(*FakeConnection).Other(), tdesc)

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
	mockAdapterDelegate.EXPECT().OnClosed(adapter)
	adapter.BindDelegate(mockAdapterDelegate)

	adapter.Close()
}
