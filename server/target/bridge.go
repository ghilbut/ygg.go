package target

import (
	. "github.com/ghilbut/ygg.go/common"
)

type TargetBridge interface {
	Register(endpoint string) bool
	Unregister(endpoint string)
}

type TargetBridgeDelegate interface {
	OnTargetBridgeStarted(bridge TargetBridge)
	OnTargetBridgeStopped()
	OnCtrlConnected(conn Connection)
	OnTargetConnected(conn Connection)
}
