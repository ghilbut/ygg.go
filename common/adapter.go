package common

type Adapter interface {
	BindDelegate(delegate AdapterDelegate)
	SetCtrlProxy(proxy *CtrlProxy)
	Close()
}

type AdapterDelegate interface {
	OnClosed(adapter Adapter)
}
