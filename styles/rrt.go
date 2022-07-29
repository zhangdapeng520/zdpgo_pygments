package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Rrt style.
var Rrt = Register(zdpgo_pygments.MustNewStyle("rrt", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.CommentPreproc:      "#e5e5e5",
	zdpgo_pygments.Comment:             "#00ff00",
	zdpgo_pygments.KeywordType:         "#ee82ee",
	zdpgo_pygments.Keyword:             "#ff0000",
	zdpgo_pygments.LiteralNumber:       "#ff6600",
	zdpgo_pygments.LiteralStringSymbol: "#ff6600",
	zdpgo_pygments.LiteralString:       "#87ceeb",
	zdpgo_pygments.NameFunction:        "#ffff00",
	zdpgo_pygments.NameConstant:        "#7fffd4",
	zdpgo_pygments.NameVariable:        "#eedd82",
	zdpgo_pygments.Background:          "#f8f8f2 bg:#000000",
}))
