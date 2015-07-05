package common_test

import (
	. "github.com/ghilbut/ygg.go/common"
	. "github.com/ghilbut/ygg.go/test/fake"
	"github.com/ghilbut/ygg.go/test/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_send_text(t *testing.T) {

	var lhs Connector = NewFakeConnector()
	var rhs Connector = lhs.(*FakeConnector).Other()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := mock.NewMockConnectorDelegate(mockCtrl)
	mockDelegate.EXPECT().OnText("A")
	rhs.BindDelegate(mockDelegate)

	conn := NewConnection(lhs)
	conn.SendText("A")
}

func Test_recv_text(t *testing.T) {

	var lhs Connector = NewFakeConnector()
	var rhs Connector = lhs.(*FakeConnector).Other()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	conn := NewConnection(lhs)

	mockDelegate := mock.NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnText(conn, "A")

	conn.BindDelegate(mockDelegate)
	rhs.SendText("A")
}

func Test_send_binary(t *testing.T) {

	var lhs Connector = NewFakeConnector()
	var rhs Connector = lhs.(*FakeConnector).Other()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := mock.NewMockConnectorDelegate(mockCtrl)
	mockDelegate.EXPECT().OnBinary([]byte{0x01, 0x02})
	rhs.BindDelegate(mockDelegate)

	conn := NewConnection(lhs)
	conn.SendBinary([]byte{0x01, 0x02})
}

func Test_recv_binary(t *testing.T) {

	var lhs Connector = NewFakeConnector()
	var rhs Connector = lhs.(*FakeConnector).Other()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	conn := NewConnection(lhs)

	mockDelegate := mock.NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnBinary(conn, []byte{0x01, 0x02})

	conn.BindDelegate(mockDelegate)
	rhs.SendBinary([]byte{0x01, 0x02})
}

func Test_close(t *testing.T) {

	var lhs Connector = NewFakeConnector()
	var rhs Connector = lhs.(*FakeConnector).Other()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	conn := NewConnection(lhs)

	{
		mockDelegate := mock.NewMockConnectionDelegate(mockCtrl)
		mockDelegate.EXPECT().OnClosed(conn)
		conn.BindDelegate(mockDelegate)
	}

	{
		mockDelegate := mock.NewMockConnectorDelegate(mockCtrl)
		mockDelegate.EXPECT().OnClosed()
		rhs.BindDelegate(mockDelegate)
	}

	conn.Close()
}

func Test_closed(t *testing.T) {

	var lhs Connector = NewFakeConnector()
	var rhs Connector = lhs.(*FakeConnector).Other()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	conn := NewConnection(lhs)

	{
		mockDelegate := mock.NewMockConnectorDelegate(mockCtrl)
		mockDelegate.EXPECT().OnClosed()
		rhs.BindDelegate(mockDelegate)
	}

	{
		mockDelegate := mock.NewMockConnectionDelegate(mockCtrl)
		mockDelegate.EXPECT().OnClosed(conn)
		conn.BindDelegate(mockDelegate)
	}

	rhs.Close()
}
