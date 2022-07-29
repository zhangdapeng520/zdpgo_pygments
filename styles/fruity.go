package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Fruity style.
var Fruity = Register(zdpgo_pygments.MustNewStyle("fruity", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.TextWhitespace:    "#888888",
	zdpgo_pygments.Background:        "#ffffff bg:#111111",
	zdpgo_pygments.GenericOutput:     "#444444 bg:#222222",
	zdpgo_pygments.Keyword:           "#fb660a bold",
	zdpgo_pygments.KeywordPseudo:     "nobold",
	zdpgo_pygments.LiteralNumber:     "#0086f7 bold",
	zdpgo_pygments.NameTag:           "#fb660a bold",
	zdpgo_pygments.NameVariable:      "#fb660a",
	zdpgo_pygments.Comment:           "#008800 bg:#0f140f italic",
	zdpgo_pygments.NameAttribute:     "#ff0086 bold",
	zdpgo_pygments.LiteralString:     "#0086d2",
	zdpgo_pygments.NameFunction:      "#ff0086 bold",
	zdpgo_pygments.GenericHeading:    "#ffffff bold",
	zdpgo_pygments.KeywordType:       "#cdcaa9 bold",
	zdpgo_pygments.GenericSubheading: "#ffffff bold",
	zdpgo_pygments.NameConstant:      "#0086d2",
	zdpgo_pygments.CommentPreproc:    "#ff0007 bold",
}))
