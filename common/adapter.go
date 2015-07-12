package common

type Adapter interface {
	BindDelegate(delegate AdapterDelegate)
	UnbindDelegate()
	SetCtrlProxy(proxy *CtrlProxy)
	Close()
}

type AdapterDelegate interface {
	OnAdapterClosed(adapter Adapter)
}
