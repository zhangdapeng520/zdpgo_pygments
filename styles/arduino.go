package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Arduino style.
var Arduino = Register(zdpgo_pygments.MustNewStyle("arduino", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Error:           "#a61717",
	zdpgo_pygments.Comment:         "#95a5a6",
	zdpgo_pygments.CommentPreproc:  "#728E00",
	zdpgo_pygments.Keyword:         "#728E00",
	zdpgo_pygments.KeywordConstant: "#00979D",
	zdpgo_pygments.KeywordPseudo:   "#00979D",
	zdpgo_pygments.KeywordReserved: "#00979D",
	zdpgo_pygments.KeywordType:     "#00979D",
	zdpgo_pygments.Operator:        "#728E00",
	zdpgo_pygments.Name:            "#434f54",
	zdpgo_pygments.NameBuiltin:     "#728E00",
	zdpgo_pygments.NameFunction:    "#D35400",
	zdpgo_pygments.NameOther:       "#728E00",
	zdpgo_pygments.LiteralNumber:   "#8A7B52",
	zdpgo_pygments.LiteralString:   "#7F8C8D",
	zdpgo_pygments.Background:      " bg:#ffffff",
}))
