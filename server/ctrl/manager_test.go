package ctrl_test

import (
	. "github.com/ghilbut/ygg.go/common"
	. "github.com/ghilbut/ygg.go/server/ctrl"
	. "github.com/ghilbut/ygg.go/test/fake"
	. "github.com/ghilbut/ygg.go/test/fake/server/ctrl"
	. "github.com/ghilbut/ygg.go/test/mock/common"
	"github.com/golang/mock/gomock"
	"log"
	"testing"
)

// const kCtrlJson = "{ \"id\": \"A0\", \"endpoint\": \"B\" }"
// const kCtrlJson = "{ \"id\": \"A1\", \"endpoint\": \"B\" }"
// const kTargetJson = "{ \"endpoint\": \"B\" }"

// const kText = "Message"

// var kBytes = []byte{0x01, 0x02}

func Test_CtrlManager_send_text(t *testing.T) {
	log.Println("######## [Test_CtrlManager_send_text] ########")

	var ctrl Connection = NewFakeConnection()
	var target Connection = NewFakeConnection()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnText(ctrl, kText)
	ctrl.BindDelegate(mockDelegate)

	connectee := NewFakeConnectee()
	connector := NewFakeConnector()
	manager := NewCtrlManager(connectee, connector)
	manager.Start()

	connector.SetTargetConnection("B", target.(*FakeConnection).Other())
	connectee.SetCtrlConnection(ctrl.(*FakeConnection).Other())
	ctrl.SendText(kCtrlJson)
	target.SendText(kTargetJson)

	target.SendText(kText)
}

func Test_CtrlManager_recv_text(t *testing.T) {
	log.Println("######## [Test_CtrlManager_recv_text] ########")

	var ctrl Connection = NewFakeConnection()
	var target Connection = NewFakeConnection()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnText(target, kText)
	target.BindDelegate(mockDelegate)

	connectee := NewFakeConnectee()
	connector := NewFakeConnector()
	manager := NewCtrlManager(connectee, connector)
	manager.Start()

	connector.SetTargetConnection("B", target.(*FakeConnection).Other())
	connectee.SetCtrlConnection(ctrl.(*FakeConnection).Other())
	ctrl.SendText(kCtrlJson)
	target.SendText(kTargetJson)

	ctrl.SendText(kText)
}

func Test_CtrlManager_send_binary(t *testing.T) {
	log.Println("######## [Test_CtrlManager_notify_binary] ########")

	var ctrl Connection = NewFakeConnection()
	var target Connection = NewFakeConnection()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnBinary(ctrl, kBytes)
	ctrl.BindDelegate(mockDelegate)

	connectee := NewFakeConnectee()
	connector := NewFakeConnector()
	manager := NewCtrlManager(connectee, connector)
	manager.Start()

	connector.SetTargetConnection("B", target.(*FakeConnection).Other())
	connectee.SetCtrlConnection(ctrl.(*FakeConnection).Other())
	ctrl.SendText(kCtrlJson)
	target.SendText(kTargetJson)

	target.SendBinary(kBytes)
}

func Test_CtrlManager_recv_binary(t *testing.T) {
	log.Println("######## [Test_CtrlManager_recv_binary] ########")

	var ctrl Connection = NewFakeConnection()
	var target Connection = NewFakeConnection()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnBinary(target, kBytes)
	target.BindDelegate(mockDelegate)

	connectee := NewFakeConnectee()
	connector := NewFakeConnector()
	manager := NewCtrlManager(connectee, connector)
	manager.Start()

	connector.SetTargetConnection("B", target.(*FakeConnection).Other())
	connectee.SetCtrlConnection(ctrl.(*FakeConnection).Other())
	ctrl.SendText(kCtrlJson)
	target.SendText(kTargetJson)

	ctrl.SendBinary(kBytes)
}

func Test_CtrlManager_remove_adapter_when_target_is_closed(t *testing.T) {
	log.Println("######## [Test_CtrlManager_remove_adapter_when_target_is_closed] ########")

	var ctrl Connection = NewFakeConnection()
	var target Connection = NewFakeConnection()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnClosed(ctrl)
	mockDelegate.EXPECT().OnClosed(target)
	ctrl.BindDelegate(mockDelegate)
	target.BindDelegate(mockDelegate)

	connectee := NewFakeConnectee()
	connector := NewFakeConnector()
	manager := NewCtrlManager(connectee, connector)
	manager.Start()

	connector.SetTargetConnection("B", target.(*FakeConnection).Other())
	connectee.SetCtrlConnection(ctrl.(*FakeConnection).Other())
	ctrl.SendText(kCtrlJson)
	target.SendText(kTargetJson)

	target.Close()
}

func Test_CtrlManager_remove_adapter_when_set_endpoint_which_alreay_exists(t *testing.T) {
	log.Println("######## [Test_CtrlManager_remove_adapter_when_set_endpoint_which_alreay_exists] ########")

	var target0 Connection = NewFakeConnection()
	var target1 Connection = NewFakeConnection()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnClosed(target0)
	mockDelegate.EXPECT().OnClosed(target1).Times(0)
	target0.BindDelegate(mockDelegate)
	target1.BindDelegate(mockDelegate)

	connectee := NewFakeConnectee()
	connector := NewFakeConnector()
	manager := NewCtrlManager(connectee, connector)
	manager.Start()

	connector.SetTargetConnection("B", target0.(*FakeConnection).Other())
	target0.SendText(kTargetJson)

	connector.SetTargetConnection("B", target1.(*FakeConnection).Other())
	target1.SendText(kTargetJson)
}

func Test_CtrlManager_remove_adapter_when_manager_is_stopped(t *testing.T) {
	log.Println("######## [Test_CtrlManager_remove_adapter_when_target_is_stopped] ########")

	var ctrl Connection = NewFakeConnection()
	var target Connection = NewFakeConnection()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnClosed(ctrl)
	mockDelegate.EXPECT().OnClosed(target)
	ctrl.BindDelegate(mockDelegate)
	target.BindDelegate(mockDelegate)

	connectee := NewFakeConnectee()
	connector := NewFakeConnector()
	manager := NewCtrlManager(connectee, connector)
	manager.Start()

	connector.SetTargetConnection("B", target.(*FakeConnection).Other())
	connectee.SetCtrlConnection(ctrl.(*FakeConnection).Other())
	ctrl.SendText(kCtrlJson)
	target.SendText(kTargetJson)

	manager.Stop()
}
