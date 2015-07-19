package target

import (
	. "github.com/ghilbut/ygg.go/common"
)

type Connectee interface {
	Register(endpoint string) bool
	Unregister(endpoint string)
}

type ConnecteeDelegate interface {
	OnConnecteeStarted(connectee Connectee)
	OnConnecteeStopped()
	OnCtrlConnected(conn Connection)
	OnTargetConnected(conn Connection)
}
