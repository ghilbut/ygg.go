package net_test

import (
	. "github.com/ghilbut/ygg.go/common"
	. "github.com/ghilbut/ygg.go/net"
	. "github.com/ghilbut/ygg.go/test/mock/common"
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_LocalConnection_has_pair(t *testing.T) {

	var lhs Connection = NewLocalConnection()
	var rhs Connection = lhs.(*LocalConnection).Other()

	if lhs == nil {
		t.Fail()
	}

	if rhs == nil {
		t.Fail()
	}
}

func Test_LocalConnection_send_text(t *testing.T) {

	var lhs Connection = NewLocalConnection()
	var rhs Connection = lhs.(*LocalConnection).Other()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnText(rhs, "A").Times(2)

	rhs.BindDelegate(mockDelegate)
	lhs.SendText("A")

	rhs.UnbindDelegate()
	lhs.SendText("A")

	rhs.BindDelegate(mockDelegate)
	lhs.SendText("A")
}

func Test_LocalConnection_recv_text(t *testing.T) {

	var lhs Connection = NewLocalConnection()
	var rhs Connection = lhs.(*LocalConnection).Other()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnText(lhs, "A").Times(2)

	lhs.BindDelegate(mockDelegate)
	rhs.SendText("A")

	lhs.UnbindDelegate()
	rhs.SendText("A")

	lhs.BindDelegate(mockDelegate)
	rhs.SendText("A")
}

func Test_LocalConnection_send_binary(t *testing.T) {

	var lhs Connection = NewLocalConnection()
	var rhs Connection = lhs.(*LocalConnection).Other()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnBinary(rhs, []byte{0x01, 0x02}).Times(2)

	rhs.BindDelegate(mockDelegate)
	lhs.SendBinary([]byte{0x01, 0x02})

	rhs.UnbindDelegate()
	lhs.SendBinary([]byte{0x01, 0x02})

	rhs.BindDelegate(mockDelegate)
	lhs.SendBinary([]byte{0x01, 0x02})
}

func Test_LocalConnection_recv_binary(t *testing.T) {

	var lhs Connection = NewLocalConnection()
	var rhs Connection = lhs.(*LocalConnection).Other()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnBinary(lhs, []byte{0x01, 0x02}).Times(2)

	lhs.BindDelegate(mockDelegate)
	rhs.SendBinary([]byte{0x01, 0x02})

	lhs.UnbindDelegate()
	rhs.SendBinary([]byte{0x01, 0x02})

	lhs.BindDelegate(mockDelegate)
	rhs.SendBinary([]byte{0x01, 0x02})
}

func Test_LocalConnection_close(t *testing.T) {

	var lhs Connection = NewLocalConnection()
	var rhs Connection = lhs.(*LocalConnection).Other()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnClosed(lhs).Times(2)
	mockDelegate.EXPECT().OnClosed(rhs).Times(2)

	lhs.BindDelegate(mockDelegate)
	rhs.BindDelegate(mockDelegate)
	lhs.Close()

	lhs.UnbindDelegate()
	rhs.UnbindDelegate()
	lhs.Close()

	lhs.BindDelegate(mockDelegate)
	rhs.BindDelegate(mockDelegate)
	lhs.Close()
}

func Test_LocalConnection_closed(t *testing.T) {

	var lhs Connection = NewLocalConnection()
	var rhs Connection = lhs.(*LocalConnection).Other()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnClosed(lhs).Times(2)
	mockDelegate.EXPECT().OnClosed(rhs).Times(2)

	lhs.BindDelegate(mockDelegate)
	rhs.BindDelegate(mockDelegate)
	rhs.Close()

	lhs.UnbindDelegate()
	rhs.UnbindDelegate()
	rhs.Close()

	lhs.BindDelegate(mockDelegate)
	rhs.BindDelegate(mockDelegate)
	rhs.Close()
}
