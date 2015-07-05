package common

type Connector interface {
	BindDelegate(delegate ConnectorDelegate)
	UnbindDelegate()

	SendText(text string)
	SendBinary(bytes []byte)
	Close()
}

type ConnectorDelegate interface {
	OnText(text string)
	OnBinary(bytes []byte)
	OnClosed()
}
