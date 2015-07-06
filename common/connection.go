package common

type Connection interface {
	BindDelegate(delegate ConnectionDelegate)
	UnbindDelegate()

	SendText(text string)
	SendBinary(bytes []byte)
	Close()
}

type ConnectionDelegate interface {
	OnText(conn Connection, text string)
	OnBinary(conn Connection, bytes []byte)
	OnClosed(conn Connection)
}
