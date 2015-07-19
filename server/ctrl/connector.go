package ctrl

import (
	. "github.com/ghilbut/ygg.go/common"
)

type Connector interface {
	Connect(ctrl *CtrlProxy) Connection
}
