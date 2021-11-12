package gamedef

import (
	"github.com/jiy1012/cellnet"
	_ "github.com/jiy1012/cellnet/codec/json"
	"github.com/jiy1012/cellnet/util"
	"github.com/davyxu/goobjfmt"
	"reflect"
)

type TestEchoJsonACK struct {
	Content string
}

func (m *TestEchoJsonACK) String() string { return goobjfmt.CompactTextString(m) }

func init() {

	// coredef.proto
	cellnet.RegisterMessageMeta("json", "gamedef.TestEchoJsonACK", reflect.TypeOf((*TestEchoJsonACK)(nil)).Elem(), util.StringHash("gamedef.TestEchoJsonACK"))
}
