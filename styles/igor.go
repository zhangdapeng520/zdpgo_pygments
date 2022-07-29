package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Igor style.
var Igor = Register(zdpgo_pygments.MustNewStyle("igor", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Comment:       "italic #FF0000",
	zdpgo_pygments.Keyword:       "#0000FF",
	zdpgo_pygments.NameFunction:  "#C34E00",
	zdpgo_pygments.NameDecorator: "#CC00A3",
	zdpgo_pygments.NameClass:     "#007575",
	zdpgo_pygments.LiteralString: "#009C00",
	zdpgo_pygments.Background:    " bg:#ffffff",
}))
