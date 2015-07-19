package fake_test

import (
	. "github.com/ghilbut/ygg.go/test/fake"
	. "github.com/ghilbut/ygg.go/test/fake/server/ctrl"
	. "github.com/ghilbut/ygg.go/test/mock/common"
	. "github.com/ghilbut/ygg.go/test/mock/server/ctrl"
	"github.com/golang/mock/gomock"
	"log"
	"testing"
)

func Test_FakeConnectee_delegate_connection_when_set_connection(t *testing.T) {
	log.Println("######## [Test_ctrl.FakeConnectee_delegate_connection_when_set_connection] ########")

	conn := NewFakeConnection()

	connectee := NewFakeConnectee()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnecteeDelegate(mockCtrl)
	mockDelegate.EXPECT().OnConnecteeStarted(connectee)
	mockDelegate.EXPECT().OnCtrlConnected(conn)

	connectee.Delegate = mockDelegate
	connectee.Start()

	connectee.SetCtrlConnection(conn)
}

func Test_FakeConnectee_close_connection_when_not_started(t *testing.T) {
	log.Println("######## [Test_ctrl.FakeConnectee_close_connection_when_not_started] ########")

	ctrl := NewFakeConnection()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnClosed(ctrl)
	ctrl.BindDelegate(mockDelegate)

	connectee := NewFakeConnectee()
	connectee.SetCtrlConnection(ctrl)
}
