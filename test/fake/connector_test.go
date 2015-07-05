package fake_test

import (
	. "github.com/ghilbut/ygg.go/common"
	. "github.com/ghilbut/ygg.go/test/fake"
	"github.com/ghilbut/ygg.go/test/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_fake_connector_has_pair(t *testing.T) {

	var lhs Connector = NewFakeConnector()
	var rhs Connector = lhs.(*FakeConnector).Other()

	if lhs == nil {
		t.Fail()
	}

	if rhs == nil {
		t.Fail()
	}
}

func Test_send_text(t *testing.T) {

	var lhs Connector = NewFakeConnector()
	var rhs Connector = lhs.(*FakeConnector).Other()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := mock.NewMockConnectorDelegate(mockCtrl)
	mockDelegate.EXPECT().OnText("A").Times(2)

	lhs.BindDelegate(mockDelegate)
	rhs.BindDelegate(mockDelegate)
	lhs.SendText("A")
	rhs.SendText("A")
}

func Test_send_binary(t *testing.T) {
	var lhs Connector = NewFakeConnector()
	var rhs Connector = lhs.(*FakeConnector).Other()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := mock.NewMockConnectorDelegate(mockCtrl)
	mockDelegate.EXPECT().OnBinary([]byte{0x01, 0x02}).Times(2)

	lhs.BindDelegate(mockDelegate)
	rhs.BindDelegate(mockDelegate)
	lhs.SendBinary([]byte{0x01, 0x02})
	rhs.SendBinary([]byte{0x01, 0x02})
}

func Test_close(t *testing.T) {
	var lhs Connector = NewFakeConnector()
	var rhs Connector = lhs.(*FakeConnector).Other()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := mock.NewMockConnectorDelegate(mockCtrl)
	mockDelegate.EXPECT().OnClosed().Times(4)

	lhs.BindDelegate(mockDelegate)
	rhs.BindDelegate(mockDelegate)
	lhs.Close()
	rhs.Close()
}
