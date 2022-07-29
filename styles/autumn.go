package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Autumn style.
var Autumn = Register(zdpgo_pygments.MustNewStyle("autumn", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.TextWhitespace:      "#bbbbbb",
	zdpgo_pygments.Comment:             "italic #aaaaaa",
	zdpgo_pygments.CommentPreproc:      "noitalic #4c8317",
	zdpgo_pygments.CommentSpecial:      "italic #0000aa",
	zdpgo_pygments.Keyword:             "#0000aa",
	zdpgo_pygments.KeywordType:         "#00aaaa",
	zdpgo_pygments.OperatorWord:        "#0000aa",
	zdpgo_pygments.NameBuiltin:         "#00aaaa",
	zdpgo_pygments.NameFunction:        "#00aa00",
	zdpgo_pygments.NameClass:           "underline #00aa00",
	zdpgo_pygments.NameNamespace:       "underline #00aaaa",
	zdpgo_pygments.NameVariable:        "#aa0000",
	zdpgo_pygments.NameConstant:        "#aa0000",
	zdpgo_pygments.NameEntity:          "bold #800",
	zdpgo_pygments.NameAttribute:       "#1e90ff",
	zdpgo_pygments.NameTag:             "bold #1e90ff",
	zdpgo_pygments.NameDecorator:       "#888888",
	zdpgo_pygments.LiteralString:       "#aa5500",
	zdpgo_pygments.LiteralStringSymbol: "#0000aa",
	zdpgo_pygments.LiteralStringRegex:  "#009999",
	zdpgo_pygments.LiteralNumber:       "#009999",
	zdpgo_pygments.GenericHeading:      "bold #000080",
	zdpgo_pygments.GenericSubheading:   "bold #800080",
	zdpgo_pygments.GenericDeleted:      "#aa0000",
	zdpgo_pygments.GenericInserted:     "#00aa00",
	zdpgo_pygments.GenericError:        "#aa0000",
	zdpgo_pygments.GenericEmph:         "italic",
	zdpgo_pygments.GenericStrong:       "bold",
	zdpgo_pygments.GenericPrompt:       "#555555",
	zdpgo_pygments.GenericOutput:       "#888888",
	zdpgo_pygments.GenericTraceback:    "#aa0000",
	zdpgo_pygments.GenericUnderline:    "underline",
	zdpgo_pygments.Error:               "#F00 bg:#FAA",
	zdpgo_pygments.Background:          " bg:#ffffff",
}))
