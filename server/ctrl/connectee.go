package ctrl

import (
	. "github.com/ghilbut/ygg.go/common"
)

type Connectee interface {
	Start(delegate ConnecteeDelegate)
	Stop()
}

type ConnecteeDelegate interface {
	OnCtrlConnected(conn Connection)
}
