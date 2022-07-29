package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Borland style.
var Borland = Register(zdpgo_pygments.MustNewStyle("borland", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.TextWhitespace:    "#bbbbbb",
	zdpgo_pygments.Comment:           "italic #008800",
	zdpgo_pygments.CommentPreproc:    "noitalic #008080",
	zdpgo_pygments.CommentSpecial:    "noitalic bold",
	zdpgo_pygments.LiteralString:     "#0000FF",
	zdpgo_pygments.LiteralStringChar: "#800080",
	zdpgo_pygments.LiteralNumber:     "#0000FF",
	zdpgo_pygments.Keyword:           "bold #000080",
	zdpgo_pygments.OperatorWord:      "bold",
	zdpgo_pygments.NameTag:           "bold #000080",
	zdpgo_pygments.NameAttribute:     "#FF0000",
	zdpgo_pygments.GenericHeading:    "#999999",
	zdpgo_pygments.GenericSubheading: "#aaaaaa",
	zdpgo_pygments.GenericDeleted:    "bg:#ffdddd #000000",
	zdpgo_pygments.GenericInserted:   "bg:#ddffdd #000000",
	zdpgo_pygments.GenericError:      "#aa0000",
	zdpgo_pygments.GenericEmph:       "italic",
	zdpgo_pygments.GenericStrong:     "bold",
	zdpgo_pygments.GenericPrompt:     "#555555",
	zdpgo_pygments.GenericOutput:     "#888888",
	zdpgo_pygments.GenericTraceback:  "#aa0000",
	zdpgo_pygments.GenericUnderline:  "underline",
	zdpgo_pygments.Error:             "bg:#e3d2d2 #a61717",
	zdpgo_pygments.Background:        " bg:#ffffff",
}))
