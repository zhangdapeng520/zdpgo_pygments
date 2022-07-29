package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Abap style.
var Abap = Register(zdpgo_pygments.MustNewStyle("abap", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Comment:        "italic #888",
	zdpgo_pygments.CommentSpecial: "#888",
	zdpgo_pygments.Keyword:        "#00f",
	zdpgo_pygments.OperatorWord:   "#00f",
	zdpgo_pygments.Name:           "#000",
	zdpgo_pygments.LiteralNumber:  "#3af",
	zdpgo_pygments.LiteralString:  "#5a2",
	zdpgo_pygments.Error:          "#F00",
	zdpgo_pygments.Background:     " bg:#ffffff",
}))
