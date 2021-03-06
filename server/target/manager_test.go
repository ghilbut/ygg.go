package target_test

import (
	. "github.com/ghilbut/ygg.go/common"
	. "github.com/ghilbut/ygg.go/net"
	. "github.com/ghilbut/ygg.go/server/target"
	. "github.com/ghilbut/ygg.go/test/fake/server/target"
	. "github.com/ghilbut/ygg.go/test/mock/common"
	"github.com/golang/mock/gomock"
	"log"
	"testing"
)

// const kCtrlA0Json = "{ \"id\": \"A0\", \"endpoint\": \"B\" }"
// const kCtrlA1Json = "{ \"id\": \"A1\", \"endpoint\": \"B\" }"
// const kTargetJson = "{ \"endpoint\": \"B\" }"

// const kText = "Message"

// var kBytes = []byte{0x01, 0x02}

func Test_TargetManager_notify_text(t *testing.T) {
	log.Println("######## [Test_TargetManager_notify_text] ########")

	var ctrlA0 Connection = NewLocalConnection()
	var ctrlA1 Connection = NewLocalConnection()
	var target Connection = NewLocalConnection()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnText(ctrlA0, kText)
	mockDelegate.EXPECT().OnText(ctrlA1, kText)
	ctrlA0.BindDelegate(mockDelegate)
	ctrlA1.BindDelegate(mockDelegate)

	manager := NewTargetManager()
	bridge := NewFakeTargetBridge()
	bridge.Delegate = manager
	bridge.Start()

	bridge.SetTargetConnection(target.(*LocalConnection).Other())
	target.SendText(kTargetJson)

	bridge.SetCtrlConnection(ctrlA0.(*LocalConnection).Other())
	ctrlA0.SendText(kCtrlA0Json)
	bridge.SetCtrlConnection(ctrlA1.(*LocalConnection).Other())
	ctrlA1.SendText(kCtrlA1Json)

	target.SendText(kText)
}

func Test_TargetManager_recv_text(t *testing.T) {
	log.Println("######## [Test_TargetManager_recv_text] ########")

	var ctrlA0 Connection = NewLocalConnection()
	var ctrlA1 Connection = NewLocalConnection()
	var target Connection = NewLocalConnection()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnText(target, kText).Times(2)
	target.BindDelegate(mockDelegate)

	manager := NewTargetManager()
	bridge := NewFakeTargetBridge()
	bridge.Delegate = manager
	bridge.Start()

	bridge.SetTargetConnection(target.(*LocalConnection).Other())
	target.SendText(kTargetJson)

	bridge.SetCtrlConnection(ctrlA0.(*LocalConnection).Other())
	ctrlA0.SendText(kCtrlA0Json)
	bridge.SetCtrlConnection(ctrlA1.(*LocalConnection).Other())
	ctrlA1.SendText(kCtrlA1Json)

	ctrlA0.SendText(kText)
	ctrlA1.SendText(kText)
}

func Test_TargetManager_notify_binary(t *testing.T) {
	log.Println("######## [Test_TargetManager_notify_binary] ########")

	var ctrlA0 Connection = NewLocalConnection()
	var ctrlA1 Connection = NewLocalConnection()
	var target Connection = NewLocalConnection()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnBinary(ctrlA0, kBytes)
	mockDelegate.EXPECT().OnBinary(ctrlA1, kBytes)
	ctrlA0.BindDelegate(mockDelegate)
	ctrlA1.BindDelegate(mockDelegate)

	manager := NewTargetManager()
	bridge := NewFakeTargetBridge()
	bridge.Delegate = manager
	bridge.Start()

	bridge.SetTargetConnection(target.(*LocalConnection).Other())
	target.SendText(kTargetJson)

	bridge.SetCtrlConnection(ctrlA0.(*LocalConnection).Other())
	ctrlA0.SendText(kCtrlA0Json)
	bridge.SetCtrlConnection(ctrlA1.(*LocalConnection).Other())
	ctrlA1.SendText(kCtrlA1Json)

	target.SendBinary(kBytes)
}

func Test_TargetManager_recv_binary(t *testing.T) {
	log.Println("######## [Test_TargetManager_recv_binary] ########")

	var ctrlA0 Connection = NewLocalConnection()
	var ctrlA1 Connection = NewLocalConnection()
	var target Connection = NewLocalConnection()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnBinary(target, kBytes).Times(2)
	target.BindDelegate(mockDelegate)

	manager := NewTargetManager()
	bridge := NewFakeTargetBridge()
	bridge.Delegate = manager
	bridge.Start()

	bridge.SetTargetConnection(target.(*LocalConnection).Other())
	target.SendText(kTargetJson)

	bridge.SetCtrlConnection(ctrlA0.(*LocalConnection).Other())
	ctrlA0.SendText(kCtrlA0Json)
	bridge.SetCtrlConnection(ctrlA1.(*LocalConnection).Other())
	ctrlA1.SendText(kCtrlA1Json)

	ctrlA0.SendBinary(kBytes)
	ctrlA1.SendBinary(kBytes)
}

func Test_TargetManager_remove_adapter_when_target_is_closed(t *testing.T) {
	log.Println("######## [Test_TargetManager_remove_adapter_when_target_is_closed] ########")

	var ctrlA0 Connection = NewLocalConnection()
	var ctrlA1 Connection = NewLocalConnection()
	var target Connection = NewLocalConnection()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnClosed(gomock.Any()).Times(2)
	ctrlA0.BindDelegate(mockDelegate)
	ctrlA1.BindDelegate(mockDelegate)

	manager := NewTargetManager()
	bridge := NewFakeTargetBridge()
	bridge.Delegate = manager
	bridge.Start()

	bridge.SetTargetConnection(target.(*LocalConnection).Other())
	target.SendText(kTargetJson)

	bridge.SetCtrlConnection(ctrlA0.(*LocalConnection).Other())
	ctrlA0.SendText(kCtrlA0Json)
	bridge.SetCtrlConnection(ctrlA1.(*LocalConnection).Other())
	ctrlA1.SendText(kCtrlA1Json)

	target.Close()

	if manager.HasEndpoint("B") {
		t.Fail()
	}
}

func Test_TargetManager_remove_adapter_when_set_endpoint_which_alreay_exists(t *testing.T) {
	log.Println("######## [Test_TargetManager_remove_adapter_when_set_endpoint_which_alreay_exists] ########")

	var target0 Connection = NewLocalConnection()
	var target1 Connection = NewLocalConnection()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnClosed(target0)
	mockDelegate.EXPECT().OnClosed(target1).Times(0)
	target0.BindDelegate(mockDelegate)
	target1.BindDelegate(mockDelegate)

	manager := NewTargetManager()
	bridge := NewFakeTargetBridge()
	bridge.Delegate = manager
	bridge.Start()

	bridge.SetTargetConnection(target0.(*LocalConnection).Other())
	target0.SendText(kTargetJson)

	bridge.SetTargetConnection(target1.(*LocalConnection).Other())
	target1.SendText(kTargetJson)
}

func Test_TargetManager_remove_adapter_when_manager_is_stopped(t *testing.T) {
	log.Println("######## [Test_TargetManager_remove_adapter_when_target_is_stopped] ########")

	var ctrlA0 Connection = NewLocalConnection()
	var ctrlA1 Connection = NewLocalConnection()
	var target Connection = NewLocalConnection()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnClosed(ctrlA0)
	mockDelegate.EXPECT().OnClosed(ctrlA1)
	mockDelegate.EXPECT().OnClosed(target)
	ctrlA0.BindDelegate(mockDelegate)
	ctrlA1.BindDelegate(mockDelegate)
	target.BindDelegate(mockDelegate)

	manager := NewTargetManager()
	bridge := NewFakeTargetBridge()
	bridge.Delegate = manager
	bridge.Start()

	bridge.SetTargetConnection(target.(*LocalConnection).Other())
	target.SendText(kTargetJson)

	bridge.SetCtrlConnection(ctrlA0.(*LocalConnection).Other())
	ctrlA0.SendText(kCtrlA0Json)
	bridge.SetCtrlConnection(ctrlA1.(*LocalConnection).Other())
	ctrlA1.SendText(kCtrlA1Json)

	bridge.Stop()

	if manager.HasEndpoint("B") {
		t.Fail()
	}
}
