package ctrl_test

import (
	. "github.com/ghilbut/ygg.go/common"
	"github.com/ghilbut/ygg.go/common/ctrl"
	"github.com/ghilbut/ygg.go/common/target"
	. "github.com/ghilbut/ygg.go/server/ctrl"
	. "github.com/ghilbut/ygg.go/test/fake"
	"testing"
)

func Test_(t *testing.T) {

	var lhs Connection = NewFakeConnection()
	var rhs Connection = NewFakeConnection()

	const kCtrlJson = "{ \"id\": \"A\", \"endpoint\": \"B\" }"
	cdesc, _ := ctrl.NewDesc(kCtrlJson)
	cproxy := ctrl.NewProxy(lhs.(*FakeConnection).Other(), cdesc)

	adapter := NewBypassAdapter(ctrl)
	if adapter == nil {
		t.Fail()
	}

	const kTargetJson = "{ \"endpoint\": \"B\" }"
	tdesc, _ := target.NewDesc(kTargetJson)
	tproxy := target.NewProxy(tdesc, rhs.(*FakeConnection).Other())
	adpater.SetTarget(tproxy)
}
