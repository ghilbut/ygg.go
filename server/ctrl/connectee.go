package ctrl

import (
	. "github.com/ghilbut/ygg.go/common"
)

type Connectee interface {
}

type ConnecteeDelegate interface {
	OnConnecteeStarted(connectee Connectee)
	OnConnecteeStopped()
	OnCtrlConnected(conn Connection)
}
