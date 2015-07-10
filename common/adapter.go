package common

type Adapter interface {
	SetCtrlProxy(proxy *CtrlProxy) bool
	HasCtrlProxy(proxy *CtrlProxy) bool
}
