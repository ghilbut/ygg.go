package fake

import (
	. "github.com/ghilbut/ygg.go/common"
	. "github.com/ghilbut/ygg.go/server/ctrl"
)

type FakeConnector struct {
	Connector

	conns map[string]Connection
}

func NewFakeConnector() *FakeConnector {

	connector := &FakeConnector{
		conns: make(map[string]Connection),
	}

	return connector
}

func (self *FakeConnector) Connect(ctrl *CtrlProxy) Connection {

	if conn, ok := self.conns[ctrl.Desc.Endpoint]; ok {
		delete(self.conns, ctrl.Desc.Endpoint)
		return conn
	}

	ctrl.Close()
	return nil
}

func (self *FakeConnector) SetTargetConnection(endpoint string, conn Connection) {
	if conn, ok := self.conns[endpoint]; ok {
		conn.Close()
	}

	self.conns[endpoint] = conn
}
