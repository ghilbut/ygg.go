package common

type _OnCtrlReadyProc func(*CtrlProxy)

type CtrlReady struct {
	ConnectionDelegate

	readys          map[Connection]bool
	OnCtrlReadyProc _OnCtrlReadyProc
}

func NewCtrlReady() *CtrlReady {

	ready := &CtrlReady{
		readys:          make(map[Connection]bool),
		OnCtrlReadyProc: nil,
	}

	return ready
}

func (self *CtrlReady) SetConnection(conn Connection) {
	r := self.readys
	if _, ok := r[conn]; ok {
		panic("[ctrl.CtrlReady] connection is already exists.")
	}

	r[conn] = true
	conn.BindDelegate(self)
}

func (self *CtrlReady) HasConnection(conn Connection) bool {
	_, ok := self.readys[conn]
	return ok
}

func (self *CtrlReady) Clear() {
	r := self.readys
	for conn, _ := range r {
		conn.Close()
	}
}

func (self *CtrlReady) OnText(conn Connection, text string) {

	var proxy *CtrlProxy = nil
	readys := self.readys
	OnCtrlReadyProc := self.OnCtrlReadyProc

	defer func() {
		if proxy == nil {
			conn.Close()
		}
	}()

	if _, ok := readys[conn]; !ok {
		panic("[ctrl.CtrlReady] there is no matched connection.")
	}

	if OnCtrlReadyProc == nil {
		panic("[ctrl.CtrlReady] OnCtrlReadyProc callback is nil.")
	}

	desc, _ := NewCtrlDesc(text)
	if desc == nil {
		return
	}

	proxy = NewCtrlProxy(conn, desc)
	if proxy != nil {
		OnCtrlReadyProc(proxy)
	}
}

func (self *CtrlReady) OnBinary(conn Connection, bytes []byte) {
	// nothing
}

func (self *CtrlReady) OnClosed(conn Connection) {
	r := self.readys
	if _, ok := r[conn]; !ok {
		panic("[ctrl.CtrlReady] there is no matched connection.")
	}

	delete(r, conn)
}
