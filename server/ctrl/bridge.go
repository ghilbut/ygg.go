package ctrl

import (
	. "github.com/ghilbut/ygg.go/common"
)

type CtrlBridge interface {
	Connect(ctrl *CtrlProxy) Connection
}

type CtrlBridgeDelegate interface {
	OnCtrlBridgeStarted(bridge CtrlBridge)
	OnCtrlBridgeStopped()
	OnCtrlConnected(conn Connection)
}
