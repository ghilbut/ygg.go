package target

import (
	. "github.com/ghilbut/ygg.go/common"
)

type Connectee interface {
	Start(delegate ConnecteeDelegate)
	Stop()
	Register(endpoint string) bool
	Unregister(endpoint string)
	HasEndpoint(endpoint string) bool
}

type ConnecteeDelegate interface {
	OnCtrlConnected(conn Connection)
	OnTargetConnected(conn Connection)
}
