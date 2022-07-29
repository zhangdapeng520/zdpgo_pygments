package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Xcode style.
var Xcode = Register(zdpgo_pygments.MustNewStyle("xcode", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Comment:           "#177500",
	zdpgo_pygments.CommentPreproc:    "#633820",
	zdpgo_pygments.LiteralString:     "#C41A16",
	zdpgo_pygments.LiteralStringChar: "#2300CE",
	zdpgo_pygments.Operator:          "#000000",
	zdpgo_pygments.Keyword:           "#A90D91",
	zdpgo_pygments.Name:              "#000000",
	zdpgo_pygments.NameAttribute:     "#836C28",
	zdpgo_pygments.NameClass:         "#3F6E75",
	zdpgo_pygments.NameFunction:      "#000000",
	zdpgo_pygments.NameBuiltin:       "#A90D91",
	zdpgo_pygments.NameBuiltinPseudo: "#5B269A",
	zdpgo_pygments.NameVariable:      "#000000",
	zdpgo_pygments.NameTag:           "#000000",
	zdpgo_pygments.NameDecorator:     "#000000",
	zdpgo_pygments.NameLabel:         "#000000",
	zdpgo_pygments.Literal:           "#1C01CE",
	zdpgo_pygments.LiteralNumber:     "#1C01CE",
	zdpgo_pygments.Error:             "#000000",
	zdpgo_pygments.Background:        " bg:#ffffff",
}))
